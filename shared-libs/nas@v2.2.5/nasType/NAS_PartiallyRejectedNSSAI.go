// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PartiallyRejectedNSSAI 9.11.3.103
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type PartiallyRejectedNSSAI struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewPartiallyRejectedNSSAI(iei uint8) (x *PartiallyRejectedNSSAI) {
	x = &PartiallyRejectedNSSAI{}
	x.SetIei(iei)
	return x
}

func (a *PartiallyRejectedNSSAI) GetIei() (iei uint8)     { return a.Iei }
func (a *PartiallyRejectedNSSAI) SetIei(iei uint8)        { a.Iei = iei }
func (a *PartiallyRejectedNSSAI) GetLen() (length uint16) { return a.Len }
func (a *PartiallyRejectedNSSAI) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PartiallyRejectedNSSAI) GetSNSSAIValue() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *PartiallyRejectedNSSAI) SetSNSSAIValue(v []uint8) { copy(a.Buffer, v) }
