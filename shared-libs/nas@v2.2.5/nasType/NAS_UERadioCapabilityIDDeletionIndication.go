// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// UERadioCapabilityIDDeletionIndication 9.11.3.69
// UERadioCapabilityIDDeletionIndication Row, sBit, len = [0, INF], 8 , INF
type UERadioCapabilityIDDeletionIndication struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewUERadioCapabilityIDDeletionIndication(iei uint8) (x *UERadioCapabilityIDDeletionIndication) {
	x = &UERadioCapabilityIDDeletionIndication{}
	x.SetIei(iei)
	return x
}

func (a *UERadioCapabilityIDDeletionIndication) GetIei() (iei uint8) {
	return a.Iei
}

func (a *UERadioCapabilityIDDeletionIndication) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *UERadioCapabilityIDDeletionIndication) GetLen() (length uint16) {
	return a.Len
}

func (a *UERadioCapabilityIDDeletionIndication) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *UERadioCapabilityIDDeletionIndication) GetUERadioCapabilityIDDeletionIndication() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *UERadioCapabilityIDDeletionIndication) SetUERadioCapabilityIDDeletionIndication(contents []uint8) {
	copy(a.Buffer, contents)
}
