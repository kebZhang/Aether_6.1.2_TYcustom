// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// UpdatedPEIPSAssistanceInformation 9.11.3.80
type UpdatedPEIPSAssistanceInformation struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewUpdatedPEIPSAssistanceInformation(iei uint8) (x *UpdatedPEIPSAssistanceInformation) {
	x = &UpdatedPEIPSAssistanceInformation{}
	x.SetIei(iei)
	return x
}

func (a *UpdatedPEIPSAssistanceInformation) GetIei() (iei uint8)    { return a.Iei }
func (a *UpdatedPEIPSAssistanceInformation) SetIei(iei uint8)       { a.Iei = iei }
func (a *UpdatedPEIPSAssistanceInformation) GetLen() (length uint8) { return a.Len }
func (a *UpdatedPEIPSAssistanceInformation) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *UpdatedPEIPSAssistanceInformation) GetUpdatedPEIPSAssistanceInformation() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *UpdatedPEIPSAssistanceInformation) SetUpdatedPEIPSAssistanceInformation(v []uint8) {
	copy(a.Buffer, v)
}
