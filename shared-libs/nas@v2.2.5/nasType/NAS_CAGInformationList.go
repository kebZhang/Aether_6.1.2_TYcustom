// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// CAGInformationList 9.11.3.18A
// CAGInformationList Row, sBit, len = [0, INF], 8 , INF
type CAGInformationList struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewCAGInformationList(iei uint8) (x *CAGInformationList) {
	x = &CAGInformationList{}
	x.SetIei(iei)
	return x
}

func (a *CAGInformationList) GetIei() (iei uint8) {
	return a.Iei
}

func (a *CAGInformationList) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *CAGInformationList) GetLen() (length uint16) {
	return a.Len
}

func (a *CAGInformationList) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *CAGInformationList) GetCAGInformationList() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *CAGInformationList) SetCAGInformationList(contents []uint8) {
	copy(a.Buffer, contents)
}
