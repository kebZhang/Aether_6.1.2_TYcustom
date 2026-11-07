// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// Fivegsmcongestionreattemptindicator 9.11.4.21
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// CATBO Row, sBit, len = [0, 0], 2, 1
// ABO Row, sBit, len = [0, 0], 1, 1
type Fivegsmcongestionreattemptindicator struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewFivegsmcongestionreattemptindicator(iei uint8) (x *Fivegsmcongestionreattemptindicator) {
	x = &Fivegsmcongestionreattemptindicator{}
	x.SetIei(iei)
	return x
}

func (a *Fivegsmcongestionreattemptindicator) GetIei() (iei uint8)    { return a.Iei }
func (a *Fivegsmcongestionreattemptindicator) SetIei(iei uint8)       { a.Iei = iei }
func (a *Fivegsmcongestionreattemptindicator) GetLen() (length uint8) { return a.Len }
func (a *Fivegsmcongestionreattemptindicator) SetLen(length uint8)    { a.Len = length }
func (a *Fivegsmcongestionreattemptindicator) GetCATBO() uint8 {
	return a.Octet & GetBitMask(2, 1) >> 1
}

func (a *Fivegsmcongestionreattemptindicator) SetCATBO(catbo uint8) {
	a.Octet = (a.Octet & 0xFD) + ((catbo & 1) << 1)
}

func (a *Fivegsmcongestionreattemptindicator) GetABO() uint8 {
	return a.Octet & GetBitMask(1, 0)
}

func (a *Fivegsmcongestionreattemptindicator) SetABO(abo uint8) {
	a.Octet = (a.Octet & 0xFE) + (abo & 1)
}
