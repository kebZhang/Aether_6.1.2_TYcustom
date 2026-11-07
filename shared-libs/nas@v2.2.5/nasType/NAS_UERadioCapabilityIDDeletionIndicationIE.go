// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// UERadioCapabilityIDDeletionIndicationIE 9.11.3.69
// Iei Row, sBit, len = [0, 0], 8 , 4
// DelInd Row, sBit, len = [0, 0], 4 , 4
type UERadioCapabilityIDDeletionIndicationIE struct {
	Octet uint8
}

func NewUERadioCapabilityIDDeletionIndicationIE(iei uint8) (x *UERadioCapabilityIDDeletionIndicationIE) {
	x = &UERadioCapabilityIDDeletionIndicationIE{}
	x.SetIei(iei)
	return x
}

// UERadioCapabilityIDDeletionIndicationIE 9.11.3.69
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *UERadioCapabilityIDDeletionIndicationIE) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

// UERadioCapabilityIDDeletionIndicationIE 9.11.3.69
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *UERadioCapabilityIDDeletionIndicationIE) SetIei(iei uint8) {
	a.Octet = (a.Octet & 15) + ((iei & 15) << 4)
}

// UERadioCapabilityIDDeletionIndicationIE 9.11.3.69
// DelInd Row, sBit, len = [0, 0], 4 , 4
func (a *UERadioCapabilityIDDeletionIndicationIE) GetDeletionIndicationValue() (delInd uint8) {
	return a.Octet & GetBitMask(4, 0)
}

// UERadioCapabilityIDDeletionIndicationIE 9.11.3.69
// DelInd Row, sBit, len = [0, 0], 4 , 4
func (a *UERadioCapabilityIDDeletionIndicationIE) SetDeletionIndicationValue(delInd uint8) {
	a.Octet = (a.Octet & 240) + (delInd & 15)
}
