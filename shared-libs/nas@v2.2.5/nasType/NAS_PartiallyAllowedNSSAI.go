// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PartiallyAllowedNSSAI 9.11.3.103
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type PartiallyAllowedNSSAI struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewPartiallyAllowedNSSAI(iei uint8) (x *PartiallyAllowedNSSAI) {
	x = &PartiallyAllowedNSSAI{}
	x.SetIei(iei)
	return x
}

func (a *PartiallyAllowedNSSAI) GetIei() (iei uint8)     { return a.Iei }
func (a *PartiallyAllowedNSSAI) SetIei(iei uint8)        { a.Iei = iei }
func (a *PartiallyAllowedNSSAI) GetLen() (length uint16) { return a.Len }
func (a *PartiallyAllowedNSSAI) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PartiallyAllowedNSSAI) GetSNSSAIValue() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *PartiallyAllowedNSSAI) SetSNSSAIValue(v []uint8) { copy(a.Buffer, v) }
