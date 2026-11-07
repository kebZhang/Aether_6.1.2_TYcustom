// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NID 9.11.3.79
// NID Row, sBit, len = [0, INF], 8 , INF
type NID struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewNID(iei uint8) (x *NID) {
	x = &NID{}
	x.SetIei(iei)
	return x
}

func (a *NID) GetIei() (iei uint8) {
	return a.Iei
}

func (a *NID) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *NID) GetLen() (length uint16) {
	return a.Len
}

func (a *NID) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *NID) GetNID() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *NID) SetNID(contents []uint8) {
	copy(a.Buffer, contents)
}
