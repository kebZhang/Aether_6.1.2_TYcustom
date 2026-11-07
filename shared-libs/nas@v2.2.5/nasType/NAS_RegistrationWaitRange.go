// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RegistrationWaitRange 9.11.3.84
// RegistrationWaitRange Row, sBit, len = [0, INF], 8 , INF
type RegistrationWaitRange struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewRegistrationWaitRange(iei uint8) (x *RegistrationWaitRange) {
	x = &RegistrationWaitRange{}
	x.SetIei(iei)
	return x
}

func (a *RegistrationWaitRange) GetIei() (iei uint8) {
	return a.Iei
}

func (a *RegistrationWaitRange) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *RegistrationWaitRange) GetLen() (length uint16) {
	return a.Len
}

func (a *RegistrationWaitRange) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *RegistrationWaitRange) GetRegistrationWaitRange() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *RegistrationWaitRange) SetRegistrationWaitRange(contents []uint8) {
	copy(a.Buffer, contents)
}
