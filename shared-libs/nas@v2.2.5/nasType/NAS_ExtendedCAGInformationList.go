// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ExtendedCAGInformationList 9.11.3.86
// ExtendedCAGInformationList Row, sBit, len = [0, INF], 8 , INF
type ExtendedCAGInformationList struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewExtendedCAGInformationList(iei uint8) (x *ExtendedCAGInformationList) {
	x = &ExtendedCAGInformationList{}
	x.SetIei(iei)
	return x
}

func (a *ExtendedCAGInformationList) GetIei() (iei uint8) {
	return a.Iei
}

func (a *ExtendedCAGInformationList) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ExtendedCAGInformationList) GetLen() (length uint16) {
	return a.Len
}

func (a *ExtendedCAGInformationList) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ExtendedCAGInformationList) GetExtendedCAGInformationList() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *ExtendedCAGInformationList) SetExtendedCAGInformationList(contents []uint8) {
	copy(a.Buffer, contents)
}
