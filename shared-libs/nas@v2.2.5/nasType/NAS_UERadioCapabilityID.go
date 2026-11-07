// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// UERadioCapabilityID 9.11.3.68
// UERadioCapabilityID Row, sBit, len = [0, INF], 8 , INF
type UERadioCapabilityID struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewUERadioCapabilityID(iei uint8) (x *UERadioCapabilityID) {
	x = &UERadioCapabilityID{}
	x.SetIei(iei)
	return x
}

func (a *UERadioCapabilityID) GetIei() (iei uint8) {
	return a.Iei
}

func (a *UERadioCapabilityID) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *UERadioCapabilityID) GetLen() (length uint16) {
	return a.Len
}

func (a *UERadioCapabilityID) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *UERadioCapabilityID) GetUERadioCapabilityID() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *UERadioCapabilityID) SetUERadioCapabilityID(contents []uint8) {
	copy(a.Buffer, contents)
}
