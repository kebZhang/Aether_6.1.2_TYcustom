// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ReAttemptIndicator 9.11.4.17
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// RATC Row, sBit, len = [0, 0], 2, 1
// EPLMNC Row, sBit, len = [0, 0], 1, 1
type ReAttemptIndicator struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewReAttemptIndicator(iei uint8) (x *ReAttemptIndicator) {
	x = &ReAttemptIndicator{}
	x.SetIei(iei)
	return x
}

func (a *ReAttemptIndicator) GetIei() (iei uint8)    { return a.Iei }
func (a *ReAttemptIndicator) SetIei(iei uint8)       { a.Iei = iei }
func (a *ReAttemptIndicator) GetLen() (length uint8) { return a.Len }
func (a *ReAttemptIndicator) SetLen(length uint8)    { a.Len = length }
func (a *ReAttemptIndicator) GetRATC() uint8 {
	return a.Octet & GetBitMask(2, 1) >> 1
}

func (a *ReAttemptIndicator) SetRATC(ratc uint8) {
	a.Octet = (a.Octet & 253) + ((ratc & 1) << 1)
}

func (a *ReAttemptIndicator) GetEPLMNC() uint8 {
	return a.Octet & GetBitMask(1, 0)
}

func (a *ReAttemptIndicator) SetEPLMNC(eplmnc uint8) {
	a.Octet = (a.Octet & 254) + (eplmnc & 1)
}
