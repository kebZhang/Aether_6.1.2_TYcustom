// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ExtendedRejectedNSSAI 9.11.3.75
// ExtendedRejectedNSSAI Row, sBit, len = [0, INF], 8 , INF
type ExtendedRejectedNSSAI struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewExtendedRejectedNSSAI(iei uint8) (x *ExtendedRejectedNSSAI) {
	x = &ExtendedRejectedNSSAI{}
	x.SetIei(iei)
	return x
}

func (a *ExtendedRejectedNSSAI) GetIei() (iei uint8) {
	return a.Iei
}

func (a *ExtendedRejectedNSSAI) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ExtendedRejectedNSSAI) GetLen() (length uint16) {
	return a.Len
}

func (a *ExtendedRejectedNSSAI) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ExtendedRejectedNSSAI) GetExtendedRejectedNSSAI() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *ExtendedRejectedNSSAI) SetExtendedRejectedNSSAI(contents []uint8) {
	copy(a.Buffer, contents)
}
