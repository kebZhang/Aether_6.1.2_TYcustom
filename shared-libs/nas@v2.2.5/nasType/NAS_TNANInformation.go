// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// TNANInformation 9.11.3.94
type TNANInformation struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewTNANInformation(iei uint8) (x *TNANInformation) {
	x = &TNANInformation{}
	x.SetIei(iei)
	return x
}

func (a *TNANInformation) GetIei() (iei uint8)    { return a.Iei }
func (a *TNANInformation) SetIei(iei uint8)       { a.Iei = iei }
func (a *TNANInformation) GetLen() (length uint8) { return a.Len }
func (a *TNANInformation) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *TNANInformation) GetTNANInformation() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *TNANInformation) SetTNANInformation(v []uint8) { copy(a.Buffer, v) }
