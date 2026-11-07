// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// TimeDuration 9.11.2.19
// TimeDuration Row, sBit, len = [0, INF], 8 , INF
type TimeDuration struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewTimeDuration(iei uint8) (x *TimeDuration) {
	x = &TimeDuration{}
	x.SetIei(iei)
	return x
}

func (a *TimeDuration) GetIei() (iei uint8) {
	return a.Iei
}

func (a *TimeDuration) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *TimeDuration) GetLen() (length uint16) {
	return a.Len
}

func (a *TimeDuration) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *TimeDuration) GetTimeDuration() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *TimeDuration) SetTimeDuration(contents []uint8) {
	copy(a.Buffer, contents)
}
