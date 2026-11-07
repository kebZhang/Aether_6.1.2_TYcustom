// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// TruncatedFiveGSTMSIConfiguration 9.11.3.70
// TruncatedFiveGSTMSIConfiguration Row, sBit, len = [0, INF], 8 , INF
type TruncatedFiveGSTMSIConfiguration struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewTruncatedFiveGSTMSIConfiguration(iei uint8) (x *TruncatedFiveGSTMSIConfiguration) {
	x = &TruncatedFiveGSTMSIConfiguration{}
	x.SetIei(iei)
	return x
}

func (a *TruncatedFiveGSTMSIConfiguration) GetIei() (iei uint8) {
	return a.Iei
}

func (a *TruncatedFiveGSTMSIConfiguration) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *TruncatedFiveGSTMSIConfiguration) GetLen() (length uint16) {
	return a.Len
}

func (a *TruncatedFiveGSTMSIConfiguration) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *TruncatedFiveGSTMSIConfiguration) GetTruncatedFiveGSTMSIConfiguration() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *TruncatedFiveGSTMSIConfiguration) SetTruncatedFiveGSTMSIConfiguration(contents []uint8) {
	copy(a.Buffer, contents)
}
