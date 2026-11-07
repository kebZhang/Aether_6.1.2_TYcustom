// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// UERequestType 9.11.3.76
// UERequestType Row, sBit, len = [0, INF], 8 , INF
type UERequestType struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewUERequestType(iei uint8) (x *UERequestType) {
	x = &UERequestType{}
	x.SetIei(iei)
	return x
}

func (a *UERequestType) GetIei() (iei uint8) {
	return a.Iei
}

func (a *UERequestType) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *UERequestType) GetLen() (length uint16) {
	return a.Len
}

func (a *UERequestType) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *UERequestType) GetUERequestType() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *UERequestType) SetUERequestType(contents []uint8) {
	copy(a.Buffer, contents)
}
