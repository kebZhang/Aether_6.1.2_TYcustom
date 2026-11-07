// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// EPSBearerContextStatus 9.11.3.23A
// EPSBearerContextStatus Row, sBit, len = [0, INF], 8 , INF
type EPSBearerContextStatus struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewEPSBearerContextStatus(iei uint8) (x *EPSBearerContextStatus) {
	x = &EPSBearerContextStatus{}
	x.SetIei(iei)
	return x
}

func (a *EPSBearerContextStatus) GetIei() (iei uint8) {
	return a.Iei
}

func (a *EPSBearerContextStatus) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *EPSBearerContextStatus) GetLen() (length uint16) {
	return a.Len
}

func (a *EPSBearerContextStatus) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *EPSBearerContextStatus) GetEPSBearerContextStatus() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *EPSBearerContextStatus) SetEPSBearerContextStatus(contents []uint8) {
	copy(a.Buffer, contents)
}
