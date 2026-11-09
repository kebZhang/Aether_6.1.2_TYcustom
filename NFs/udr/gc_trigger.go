// SPDX-License-Identifier: Apache-2.0
//
// TYcustom noGC: GET :6060/gc runs one full GC on demand, so an experiment
// that turns GC off on the pod (GOGC=off) can still collect between runs.
//
// This file changes no GC setting. Whether this NF collects on its own is
// decided only by the pod environment (GOGC, GOMEMLIMIT); the reply below
// only reports what the runtime is using.
//
// runtime.GC returns only after mark and sweep are finished, so the reply
// comes after the whole GC. Reply = plain "key value" lines:
//   gogc         GOGC in effect in this process: "off" or the percent
//   auto_gc      GCs since the process started that were not forced
//                (NumGC - NumForcedGC); stays 0 while GC is off
//   heap_before  HeapAlloc before this GC: what the last run left on the heap
//   heap_after   HeapAlloc after this GC
//
// Nothing else is served on :6060; without requests the listener only waits.

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/metrics"
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/gc", forceGC)
	go func() {
		if err := http.ListenAndServe(":6060", mux); err != nil {
			fmt.Fprintln(os.Stderr, "TYcustom gc trigger:", err)
		}
	}()
}

func forceGC(w http.ResponseWriter, _ *http.Request) {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	runtime.GC()
	runtime.ReadMemStats(&after)

	s := []metrics.Sample{{Name: "/gc/gogc:percent"}} // GOGC=off reads as -1
	metrics.Read(s)
	gogc := "off"
	if p := int64(s[0].Value.Uint64()); p >= 0 {
		gogc = fmt.Sprint(p)
	}

	fmt.Fprintf(w, "gogc %s\nauto_gc %d\nheap_before %d\nheap_after %d\n",
		gogc, after.NumGC-after.NumForcedGC, before.HeapAlloc, after.HeapAlloc)
}
