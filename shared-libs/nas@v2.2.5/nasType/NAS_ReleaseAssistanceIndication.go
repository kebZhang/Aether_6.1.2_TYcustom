// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ReleaseAssistanceIndication 9.11.3.46A
// Iei Row, sBit, len = [0, 0], 8 , 4
// PDDEI Row, sBit, len = [0, 0], 2 , 2
type ReleaseAssistanceIndication struct {
	Octet uint8
}

func NewReleaseAssistanceIndication(iei uint8) (x *ReleaseAssistanceIndication) {
	x = &ReleaseAssistanceIndication{}
	x.SetIei(iei)
	return x
}

// ReleaseAssistanceIndication 9.11.3.46A
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *ReleaseAssistanceIndication) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

// ReleaseAssistanceIndication 9.11.3.46A
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *ReleaseAssistanceIndication) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

// ReleaseAssistanceIndication 9.11.3.46A
// PDDEI Row, sBit, len = [0, 0], 2 , 2
func (a *ReleaseAssistanceIndication) GetPDDEI() (pDDEI uint8) {
	return a.Octet & GetBitMask(2, 0)
}

// ReleaseAssistanceIndication 9.11.3.46A
// PDDEI Row, sBit, len = [0, 0], 2 , 2
func (a *ReleaseAssistanceIndication) SetPDDEI(pDDEI uint8) {
	a.Octet = (a.Octet & 0xFC) | (pDDEI & 0x03)
}
