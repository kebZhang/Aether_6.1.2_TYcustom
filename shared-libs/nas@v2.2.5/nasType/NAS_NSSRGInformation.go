// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NSSRGInformation 9.11.3.82
// NSSRGInformation Row, sBit, len = [0, INF], 8 , INF
type NSSRGInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewNSSRGInformation(iei uint8) (x *NSSRGInformation) {
	x = &NSSRGInformation{}
	x.SetIei(iei)
	return x
}

func (a *NSSRGInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *NSSRGInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *NSSRGInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *NSSRGInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *NSSRGInformation) GetNSSRGInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *NSSRGInformation) SetNSSRGInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
