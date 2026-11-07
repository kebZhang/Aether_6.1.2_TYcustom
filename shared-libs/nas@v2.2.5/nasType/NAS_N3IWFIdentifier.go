// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// N3IWFIdentifier 9.11.3.93
type N3IWFIdentifier struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewN3IWFIdentifier(iei uint8) (x *N3IWFIdentifier) {
	x = &N3IWFIdentifier{}
	x.SetIei(iei)
	return x
}

func (a *N3IWFIdentifier) GetIei() (iei uint8)    { return a.Iei }
func (a *N3IWFIdentifier) SetIei(iei uint8)       { a.Iei = iei }
func (a *N3IWFIdentifier) GetLen() (length uint8) { return a.Len }
func (a *N3IWFIdentifier) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *N3IWFIdentifier) GetN3IWFIdentifier() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *N3IWFIdentifier) SetN3IWFIdentifier(v []uint8) { copy(a.Buffer, v) }
