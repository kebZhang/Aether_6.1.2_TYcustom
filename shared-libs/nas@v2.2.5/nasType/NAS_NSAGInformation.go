// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NSAGInformation 9.11.3.87
// NSAGInformation Row, sBit, len = [0, INF], 8 , INF
type NSAGInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewNSAGInformation(iei uint8) (x *NSAGInformation) {
	x = &NSAGInformation{}
	x.SetIei(iei)
	return x
}

func (a *NSAGInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *NSAGInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *NSAGInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *NSAGInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *NSAGInformation) GetNSAGInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *NSAGInformation) SetNSAGInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
