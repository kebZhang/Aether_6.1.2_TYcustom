// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// MobileStationClassmark2 9.11.3.31C
// MobileStationClassmark2 Row, sBit, len = [0, INF], 8 , INF
type MobileStationClassmark2 struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewMobileStationClassmark2(iei uint8) (x *MobileStationClassmark2) {
	x = &MobileStationClassmark2{}
	x.SetIei(iei)
	return x
}

func (a *MobileStationClassmark2) GetIei() (iei uint8) {
	return a.Iei
}

func (a *MobileStationClassmark2) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *MobileStationClassmark2) GetLen() (length uint16) {
	return a.Len
}

func (a *MobileStationClassmark2) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *MobileStationClassmark2) GetMobileStationClassmark2() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *MobileStationClassmark2) SetMobileStationClassmark2(contents []uint8) {
	copy(a.Buffer, contents)
}
