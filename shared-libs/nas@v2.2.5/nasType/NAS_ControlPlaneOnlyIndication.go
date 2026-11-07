// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ControlPlaneOnlyIndication 9.11.4.23
// Iei Row, sBit, len = [0, 0], 8 , 4
// CPOI Row, sBit, len = [0, 0], 1 , 1
type ControlPlaneOnlyIndication struct {
	Octet uint8
}

func NewControlPlaneOnlyIndication(iei uint8) (x *ControlPlaneOnlyIndication) {
	x = &ControlPlaneOnlyIndication{}
	x.SetIei(iei)
	return x
}

func (a *ControlPlaneOnlyIndication) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

func (a *ControlPlaneOnlyIndication) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

func (a *ControlPlaneOnlyIndication) GetCPOI() (cpoi uint8) {
	return a.Octet & GetBitMask(1, 0)
}

func (a *ControlPlaneOnlyIndication) SetCPOI(cpoi uint8) {
	a.Octet = (a.Octet & 0xFE) | (cpoi & 0x01)
}
