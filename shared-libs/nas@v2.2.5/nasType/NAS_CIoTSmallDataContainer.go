// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// CIoTSmallDataContainer 9.11.3.18B
type CIoTSmallDataContainer struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewCIoTSmallDataContainer(iei uint8) (x *CIoTSmallDataContainer) {
	x = &CIoTSmallDataContainer{}
	x.SetIei(iei)
	return x
}

func (a *CIoTSmallDataContainer) GetIei() (iei uint8)    { return a.Iei }
func (a *CIoTSmallDataContainer) SetIei(iei uint8)       { a.Iei = iei }
func (a *CIoTSmallDataContainer) GetLen() (length uint8) { return a.Len }
func (a *CIoTSmallDataContainer) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *CIoTSmallDataContainer) GetCIoTSmallDataContainer() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *CIoTSmallDataContainer) SetCIoTSmallDataContainer(v []uint8) { copy(a.Buffer, v) }
