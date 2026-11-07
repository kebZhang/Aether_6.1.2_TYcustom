// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// EthernetHeaderCompressionConfiguration 9.11.4.28
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// CIDLen Row, sBit, len = [0, 0], 2, 2
type EthernetHeaderCompressionConfiguration struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewEthernetHeaderCompressionConfiguration(iei uint8) (x *EthernetHeaderCompressionConfiguration) {
	x = &EthernetHeaderCompressionConfiguration{}
	x.SetIei(iei)
	return x
}

func (a *EthernetHeaderCompressionConfiguration) GetIei() (iei uint8)    { return a.Iei }
func (a *EthernetHeaderCompressionConfiguration) SetIei(iei uint8)       { a.Iei = iei }
func (a *EthernetHeaderCompressionConfiguration) GetLen() (length uint8) { return a.Len }
func (a *EthernetHeaderCompressionConfiguration) SetLen(length uint8)    { a.Len = length }
func (a *EthernetHeaderCompressionConfiguration) GetCIDLen() (cidLen uint8) {
	return a.Octet & GetBitMask(2, 0)
}

func (a *EthernetHeaderCompressionConfiguration) SetCIDLen(cidLen uint8) {
	a.Octet = (a.Octet & 0xFC) | (cidLen & 0x03)
}
