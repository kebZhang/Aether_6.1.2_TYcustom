// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ExtendedDRXParameters 9.11.3.26A
// ExtendedDRXParameters Row, sBit, len = [0, INF], 8 , INF
type ExtendedDRXParameters struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewExtendedDRXParameters(iei uint8) (x *ExtendedDRXParameters) {
	x = &ExtendedDRXParameters{}
	x.SetIei(iei)
	return x
}

func (a *ExtendedDRXParameters) GetIei() (iei uint8) {
	return a.Iei
}

func (a *ExtendedDRXParameters) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ExtendedDRXParameters) GetLen() (length uint16) {
	return a.Len
}

func (a *ExtendedDRXParameters) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ExtendedDRXParameters) GetExtendedDRXParameters() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *ExtendedDRXParameters) SetExtendedDRXParameters(contents []uint8) {
	copy(a.Buffer, contents)
}
