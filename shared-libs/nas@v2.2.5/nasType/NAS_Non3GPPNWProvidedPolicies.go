// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// Non3GPPNWProvidedPolicies 9.11.3.36A
// Iei Row, sBit, len = [0, 0], 8 , 4
// Value Row, sBit, len = [0, 0], 4 , 4
type Non3GPPNWProvidedPolicies struct {
	Octet uint8
}

func NewNon3GPPNWProvidedPolicies(iei uint8) (non3GPPNWProvidedPolicies *Non3GPPNWProvidedPolicies) {
	non3GPPNWProvidedPolicies = &Non3GPPNWProvidedPolicies{}
	non3GPPNWProvidedPolicies.SetIei(iei)
	return non3GPPNWProvidedPolicies
}

func (a *Non3GPPNWProvidedPolicies) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

func (a *Non3GPPNWProvidedPolicies) SetIei(iei uint8) {
	a.Octet = (a.Octet & 15) + ((iei & 15) << 4)
}

func (a *Non3GPPNWProvidedPolicies) GetValue() (value uint8) {
	return a.Octet & GetBitMask(4, 0)
}

func (a *Non3GPPNWProvidedPolicies) SetValue(value uint8) {
	a.Octet = (a.Octet & 240) + (value & 15)
}
