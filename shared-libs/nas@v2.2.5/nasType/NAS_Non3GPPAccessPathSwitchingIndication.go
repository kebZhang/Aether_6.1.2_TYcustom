// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// Non3GPPAccessPathSwitchingIndication 9.11.3.99
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// NAPS Row, sBit, len = [0, 0], 1 , 1
type Non3GPPAccessPathSwitchingIndication struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewNon3GPPAccessPathSwitchingIndication(iei uint8) (x *Non3GPPAccessPathSwitchingIndication) {
	x = &Non3GPPAccessPathSwitchingIndication{}
	x.SetIei(iei)
	return x
}

func (a *Non3GPPAccessPathSwitchingIndication) GetIei() (iei uint8)    { return a.Iei }
func (a *Non3GPPAccessPathSwitchingIndication) SetIei(iei uint8)       { a.Iei = iei }
func (a *Non3GPPAccessPathSwitchingIndication) GetLen() (length uint8) { return a.Len }
func (a *Non3GPPAccessPathSwitchingIndication) SetLen(length uint8)    { a.Len = length }
func (a *Non3GPPAccessPathSwitchingIndication) GetNAPS() uint8 {
	return a.Octet & GetBitMask(1, 0)
}

func (a *Non3GPPAccessPathSwitchingIndication) SetNAPS(naps uint8) {
	a.Octet = (a.Octet & 0xFE) | (naps & 0x01)
}
