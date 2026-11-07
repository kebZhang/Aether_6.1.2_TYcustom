// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RANTimingSynchronization 9.11.3.95
type RANTimingSynchronization struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewRANTimingSynchronization(iei uint8) (x *RANTimingSynchronization) {
	x = &RANTimingSynchronization{}
	x.SetIei(iei)
	return x
}

func (a *RANTimingSynchronization) GetIei() (iei uint8)    { return a.Iei }
func (a *RANTimingSynchronization) SetIei(iei uint8)       { a.Iei = iei }
func (a *RANTimingSynchronization) GetLen() (length uint8) { return a.Len }
func (a *RANTimingSynchronization) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *RANTimingSynchronization) GetRANTimingSynchronization() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *RANTimingSynchronization) SetRANTimingSynchronization(v []uint8) { copy(a.Buffer, v) }
