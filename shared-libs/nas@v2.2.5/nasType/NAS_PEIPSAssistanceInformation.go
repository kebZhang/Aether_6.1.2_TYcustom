// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PEIPSAssistanceInformation 9.11.3.80
// PEIPSAssistanceInformation Row, sBit, len = [0, INF], 8 , INF
type PEIPSAssistanceInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewPEIPSAssistanceInformation(iei uint8) (x *PEIPSAssistanceInformation) {
	x = &PEIPSAssistanceInformation{}
	x.SetIei(iei)
	return x
}

func (a *PEIPSAssistanceInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *PEIPSAssistanceInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *PEIPSAssistanceInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *PEIPSAssistanceInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PEIPSAssistanceInformation) GetPEIPSAssistanceInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *PEIPSAssistanceInformation) SetPEIPSAssistanceInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
