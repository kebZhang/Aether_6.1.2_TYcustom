// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// AUN3Indication 9.11.3.104
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// AUN3REG Row, sBit, len = [0, 0], 1, 1
type AUN3Indication struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewAUN3Indication(iei uint8) (x *AUN3Indication) {
	x = &AUN3Indication{}
	x.SetIei(iei)
	x.Len = 1
	return x
}

func (a *AUN3Indication) GetIei() (iei uint8) { return a.Iei }
func (a *AUN3Indication) SetIei(iei uint8)    { a.Iei = iei }
func (a *AUN3Indication) GetLen() (n uint8)   { return a.Len }
func (a *AUN3Indication) SetLen(n uint8)      { a.Len = n }

func (a *AUN3Indication) GetAUN3REG() uint8 {
	return a.Octet & GetBitMask(1, 0)
}

func (a *AUN3Indication) SetAUN3REG(aun3reg uint8) {
	a.Octet = (a.Octet & 0xFE) | (aun3reg & 0x01)
}
