// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// AdditionalConfigurationIndication 9.11.3.74
// Iei Row, sBit, len = [0, 0], 8 , 4
// SCMR Row, sBit, len = [0, 0], 1 , 1
type AdditionalConfigurationIndication struct {
	Octet uint8
}

func NewAdditionalConfigurationIndication(iei uint8) (x *AdditionalConfigurationIndication) {
	x = &AdditionalConfigurationIndication{}
	x.SetIei(iei)
	return x
}

func (a *AdditionalConfigurationIndication) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

func (a *AdditionalConfigurationIndication) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

func (a *AdditionalConfigurationIndication) GetSCMR() (scmr uint8) {
	return a.Octet & GetBitMask(1, 0)
}

func (a *AdditionalConfigurationIndication) SetSCMR(scmr uint8) {
	a.Octet = (a.Octet & 0xFE) | (scmr & 0x01)
}
