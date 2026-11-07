// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PendingNSSAI 9.11.3.37
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type PendingNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewPendingNSSAI(iei uint8) (pendingNSSAI *PendingNSSAI) {
	pendingNSSAI = &PendingNSSAI{}
	pendingNSSAI.SetIei(iei)
	return pendingNSSAI
}

func (a *PendingNSSAI) GetIei() (iei uint8) {
	return a.Iei
}

func (a *PendingNSSAI) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *PendingNSSAI) GetLen() (length uint8) {
	return a.Len
}

func (a *PendingNSSAI) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PendingNSSAI) GetSNSSAIValue() (sNSSAIValue []uint8) {
	sNSSAIValue = make([]uint8, len(a.Buffer))
	copy(sNSSAIValue, a.Buffer)
	return sNSSAIValue
}

func (a *PendingNSSAI) SetSNSSAIValue(sNSSAIValue []uint8) {
	copy(a.Buffer, sNSSAIValue)
}
