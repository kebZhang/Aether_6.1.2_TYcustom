// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// Non3GPPPathSwitchingInformation 9.11.3.102
// Non3GPPPathSwitchingInformation Row, sBit, len = [0, INF], 8 , INF
type Non3GPPPathSwitchingInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewNon3GPPPathSwitchingInformation(iei uint8) (x *Non3GPPPathSwitchingInformation) {
	x = &Non3GPPPathSwitchingInformation{}
	x.SetIei(iei)
	return x
}

func (a *Non3GPPPathSwitchingInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *Non3GPPPathSwitchingInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *Non3GPPPathSwitchingInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *Non3GPPPathSwitchingInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *Non3GPPPathSwitchingInformation) GetNon3GPPPathSwitchingInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *Non3GPPPathSwitchingInformation) SetNon3GPPPathSwitchingInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
