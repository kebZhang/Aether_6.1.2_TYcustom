// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ReceivedMBSContainer 9.11.4.31
type ReceivedMBSContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewReceivedMBSContainer(iei uint8) (x *ReceivedMBSContainer) {
	x = &ReceivedMBSContainer{}
	x.SetIei(iei)
	return x
}

func (a *ReceivedMBSContainer) GetIei() (iei uint8)     { return a.Iei }
func (a *ReceivedMBSContainer) SetIei(iei uint8)        { a.Iei = iei }
func (a *ReceivedMBSContainer) GetLen() (length uint16) { return a.Len }
func (a *ReceivedMBSContainer) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ReceivedMBSContainer) GetReceivedMBSContainer() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *ReceivedMBSContainer) SetReceivedMBSContainer(v []uint8) { copy(a.Buffer, v) }
