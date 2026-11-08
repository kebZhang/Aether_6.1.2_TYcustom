// SPDX-License-Identifier: Apache-2.0
//
// TYcustom: open sbiConnsPerPeer HTTP/2 connections from this NF to every peer
// NF instead of one, and spread the requests over them round-robin.
//
// Every outgoing HTTP request of this NF goes through http.DefaultTransport:
// the openapi clients (cfg.HTTPClient = &http.Client{}) and every http.Client
// in this NF leave Transport nil. Replacing it here, before main runs, changes
// the connection count towards all peers without touching the openapi library.
//
// Each slot is an independent clone of the stock transport, so each keeps its
// own connection pool and opens its own TCP+TLS connection to a peer. Inside a
// slot the stock behaviour is unchanged: when a connection reaches the
// server's max concurrent streams (250) another one is dialed, and an idle
// connection is closed by the peer's server after 60 s.
//
// Set sbiConnsPerPeer = 1 to get back the stock behaviour (one slot).

package main

import (
	"net/http"
	"sync"
	"sync/atomic"
)

const sbiConnsPerPeer = 8

type multiConnTransport struct {
	slots   []http.RoundTripper
	cursors sync.Map // peer host:port -> *atomic.Uint32
}

func newMultiConnTransport(n int) *multiConnTransport {
	base := http.DefaultTransport.(*http.Transport)
	t := &multiConnTransport{slots: make([]http.RoundTripper, n)}
	for i := range t.slots {
		t.slots[i] = base.Clone()
	}
	return t
}

func (t *multiConnTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	c, ok := t.cursors.Load(req.URL.Host)
	if !ok {
		c, _ = t.cursors.LoadOrStore(req.URL.Host, new(atomic.Uint32))
	}
	i := c.(*atomic.Uint32).Add(1) % uint32(len(t.slots))
	return t.slots[i].RoundTrip(req)
}

func init() {
	http.DefaultTransport = newMultiConnTransport(sbiConnsPerPeer)
}
