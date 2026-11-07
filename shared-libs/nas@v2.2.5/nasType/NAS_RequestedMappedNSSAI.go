// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RequestedMappedNSSAI 9.11.3.31B
// RequestedMappedNSSAI Row, sBit, len = [0, INF], 8 , INF
type RequestedMappedNSSAI struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewRequestedMappedNSSAI(iei uint8) (x *RequestedMappedNSSAI) {
	x = &RequestedMappedNSSAI{}
	x.SetIei(iei)
	return x
}

func (a *RequestedMappedNSSAI) GetIei() (iei uint8) {
	return a.Iei
}

func (a *RequestedMappedNSSAI) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *RequestedMappedNSSAI) GetLen() (length uint16) {
	return a.Len
}

func (a *RequestedMappedNSSAI) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *RequestedMappedNSSAI) GetRequestedMappedNSSAI() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *RequestedMappedNSSAI) SetRequestedMappedNSSAI(contents []uint8) {
	copy(a.Buffer, contents)
}
