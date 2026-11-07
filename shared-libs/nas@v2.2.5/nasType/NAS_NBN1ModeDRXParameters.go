// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NBN1ModeDRXParameters 9.11.3.73
// NBN1ModeDRXParameters Row, sBit, len = [0, INF], 8 , INF
type NBN1ModeDRXParameters struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewNBN1ModeDRXParameters(iei uint8) (x *NBN1ModeDRXParameters) {
	x = &NBN1ModeDRXParameters{}
	x.SetIei(iei)
	return x
}

func (a *NBN1ModeDRXParameters) GetIei() (iei uint8) {
	return a.Iei
}

func (a *NBN1ModeDRXParameters) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *NBN1ModeDRXParameters) GetLen() (length uint16) {
	return a.Len
}

func (a *NBN1ModeDRXParameters) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *NBN1ModeDRXParameters) GetNBN1ModeDRXParameters() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *NBN1ModeDRXParameters) SetNBN1ModeDRXParameters(contents []uint8) {
	copy(a.Buffer, contents)
}
