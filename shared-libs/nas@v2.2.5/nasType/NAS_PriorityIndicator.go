// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PriorityIndicator 9.11.3.91
// Iei Row, sBit, len = [0, 0], 8 , 4
// MPSI Row, sBit, len = [0, 0], 1 , 1
type PriorityIndicator struct {
	Octet uint8
}

func NewPriorityIndicator(iei uint8) (x *PriorityIndicator) {
	x = &PriorityIndicator{}
	x.SetIei(iei)
	return x
}

func (a *PriorityIndicator) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

func (a *PriorityIndicator) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

func (a *PriorityIndicator) GetMPSI() (mpsi uint8) {
	return a.Octet & GetBitMask(1, 0)
}

func (a *PriorityIndicator) SetMPSI(mpsi uint8) {
	a.Octet = (a.Octet & 0xFE) | (mpsi & 0x01)
}
