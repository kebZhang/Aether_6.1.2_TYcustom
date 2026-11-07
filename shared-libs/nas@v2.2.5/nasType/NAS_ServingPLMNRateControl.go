// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ServingPLMNRateControl 9.11.4.20
type ServingPLMNRateControl struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewServingPLMNRateControl(iei uint8) (x *ServingPLMNRateControl) {
	x = &ServingPLMNRateControl{}
	x.SetIei(iei)
	return x
}

func (a *ServingPLMNRateControl) GetIei() (iei uint8)    { return a.Iei }
func (a *ServingPLMNRateControl) SetIei(iei uint8)       { a.Iei = iei }
func (a *ServingPLMNRateControl) GetLen() (length uint8) { return a.Len }
func (a *ServingPLMNRateControl) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ServingPLMNRateControl) GetServingPLMNRateControl() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *ServingPLMNRateControl) SetServingPLMNRateControl(v []uint8) { copy(a.Buffer, v) }
