// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ATSSSContainer 9.11.4.22
type ATSSSContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewATSSSContainer(iei uint8) (x *ATSSSContainer) {
	x = &ATSSSContainer{}
	x.SetIei(iei)
	return x
}

func (a *ATSSSContainer) GetIei() (iei uint8) { return a.Iei }
func (a *ATSSSContainer) SetIei(iei uint8)    { a.Iei = iei }
func (a *ATSSSContainer) GetLen() (n uint16)  { return a.Len }
func (a *ATSSSContainer) SetLen(n uint16) {
	a.Len = n
	a.Buffer = make([]uint8, a.Len)
}

func (a *ATSSSContainer) GetATSSSContainer() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *ATSSSContainer) SetATSSSContainer(v []uint8) { copy(a.Buffer, v) }
