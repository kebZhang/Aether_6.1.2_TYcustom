// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// N5GCIndication 9.11.3.72
// Iei Row, sBit, len = [0, 0], 8 , 4
// N5GCREG Row, sBit, len = [0, 0], 1 , 1
type N5GCIndication struct {
	Octet uint8
}

func NewN5GCIndication(iei uint8) (n5GCIndication *N5GCIndication) {
	n5GCIndication = &N5GCIndication{}
	n5GCIndication.SetIei(iei)
	return n5GCIndication
}

// N5GCIndication 9.11.3.72
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *N5GCIndication) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

// N5GCIndication 9.11.3.72
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *N5GCIndication) SetIei(iei uint8) {
	a.Octet = (a.Octet & 15) + ((iei & 15) << 4)
}

// N5GCIndication 9.11.3.72
// N5GCREG Row, sBit, len = [0, 0], 1 , 1
func (a *N5GCIndication) GetN5GCREG() (n5GCREG uint8) {
	return a.Octet & GetBitMask(1, 0)
}

// N5GCIndication 9.11.3.72
// N5GCREG Row, sBit, len = [0, 0], 1 , 1
func (a *N5GCIndication) SetN5GCREG(n5GCREG uint8) {
	a.Octet = (a.Octet & 0xFE) | (n5GCREG & 0x01)
}
