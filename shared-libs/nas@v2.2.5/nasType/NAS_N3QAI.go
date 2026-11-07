// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// N3QAI 9.11.4.36
type N3QAI struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewN3QAI(iei uint8) (x *N3QAI) {
	x = &N3QAI{}
	x.SetIei(iei)
	return x
}

func (a *N3QAI) GetIei() (iei uint8)     { return a.Iei }
func (a *N3QAI) SetIei(iei uint8)        { a.Iei = iei }
func (a *N3QAI) GetLen() (length uint16) { return a.Len }
func (a *N3QAI) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *N3QAI) GetN3QAI() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *N3QAI) SetN3QAI(v []uint8) { copy(a.Buffer, v) }
