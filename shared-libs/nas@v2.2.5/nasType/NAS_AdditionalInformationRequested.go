// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// AdditionalInformationRequested 9.11.3.12A
// AdditionalInformationRequested Row, sBit, len = [0, INF], 8 , INF
type AdditionalInformationRequested struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewAdditionalInformationRequested(iei uint8) (x *AdditionalInformationRequested) {
	x = &AdditionalInformationRequested{}
	x.SetIei(iei)
	return x
}

func (a *AdditionalInformationRequested) GetIei() (iei uint8) {
	return a.Iei
}

func (a *AdditionalInformationRequested) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *AdditionalInformationRequested) GetLen() (length uint16) {
	return a.Len
}

func (a *AdditionalInformationRequested) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *AdditionalInformationRequested) GetAdditionalInformationRequested() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *AdditionalInformationRequested) SetAdditionalInformationRequested(contents []uint8) {
	copy(a.Buffer, contents)
}
