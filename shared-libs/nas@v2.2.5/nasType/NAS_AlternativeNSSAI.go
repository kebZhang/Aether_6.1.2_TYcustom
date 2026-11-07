// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// AlternativeNSSAI 9.11.3.97
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type AlternativeNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewAlternativeNSSAI(iei uint8) (alternativeNSSAI *AlternativeNSSAI) {
	alternativeNSSAI = &AlternativeNSSAI{}
	alternativeNSSAI.SetIei(iei)
	return alternativeNSSAI
}

func (a *AlternativeNSSAI) GetIei() (iei uint8) {
	return a.Iei
}

func (a *AlternativeNSSAI) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *AlternativeNSSAI) GetLen() (length uint8) {
	return a.Len
}

func (a *AlternativeNSSAI) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *AlternativeNSSAI) GetSNSSAIValue() (sNSSAIValue []uint8) {
	sNSSAIValue = make([]uint8, len(a.Buffer))
	copy(sNSSAIValue, a.Buffer)
	return sNSSAIValue
}

func (a *AlternativeNSSAI) SetSNSSAIValue(sNSSAIValue []uint8) {
	copy(a.Buffer, sNSSAIValue)
}
