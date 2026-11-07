// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// MAPDUSessionInformation 9.11.3.31A
// Iei Row, sBit, len = [0, 0], 8 , 4
// MAPSI Row, sBit, len = [0, 0], 4 , 4
type MAPDUSessionInformation struct {
	Octet uint8
}

func NewMAPDUSessionInformation(iei uint8) (x *MAPDUSessionInformation) {
	x = &MAPDUSessionInformation{}
	x.SetIei(iei)
	return x
}

// MAPDUSessionInformation 9.11.3.31A
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *MAPDUSessionInformation) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

// MAPDUSessionInformation 9.11.3.31A
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *MAPDUSessionInformation) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

// MAPDUSessionInformation 9.11.3.31A
// MAPSI Row, sBit, len = [0, 0], 4 , 4
func (a *MAPDUSessionInformation) GetMAPSI() (mapsi uint8) {
	return a.Octet & GetBitMask(4, 0)
}

// MAPDUSessionInformation 9.11.3.31A
// MAPSI Row, sBit, len = [0, 0], 4 , 4
func (a *MAPDUSessionInformation) SetMAPSI(mapsi uint8) {
	a.Octet = (a.Octet & 0xF0) | (mapsi & 0x0F)
}
