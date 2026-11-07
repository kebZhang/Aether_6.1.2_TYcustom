// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// EquivalentSNPNs 9.11.3.92
type EquivalentSNPNs struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewEquivalentSNPNs(iei uint8) (x *EquivalentSNPNs) {
	x = &EquivalentSNPNs{}
	x.SetIei(iei)
	return x
}

func (a *EquivalentSNPNs) GetIei() (iei uint8)    { return a.Iei }
func (a *EquivalentSNPNs) SetIei(iei uint8)       { a.Iei = iei }
func (a *EquivalentSNPNs) GetLen() (length uint8) { return a.Len }
func (a *EquivalentSNPNs) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *EquivalentSNPNs) GetEquivalentSNPNs() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *EquivalentSNPNs) SetEquivalentSNPNs(v []uint8) { copy(a.Buffer, v) }
