// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// WUSAssistanceInformation 9.11.3.71
// WUSAssistanceInformation Row, sBit, len = [0, INF], 8 , INF
type WUSAssistanceInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewWUSAssistanceInformation(iei uint8) (x *WUSAssistanceInformation) {
	x = &WUSAssistanceInformation{}
	x.SetIei(iei)
	return x
}

func (a *WUSAssistanceInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *WUSAssistanceInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *WUSAssistanceInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *WUSAssistanceInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *WUSAssistanceInformation) GetWUSAssistanceInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *WUSAssistanceInformation) SetWUSAssistanceInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
