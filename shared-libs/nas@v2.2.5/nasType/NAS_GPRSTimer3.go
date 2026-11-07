// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// GPRSTimer3 9.11.2.5
// GPRSTimer3 Row, sBit, len = [0, INF], 8 , INF
type GPRSTimer3 struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewGPRSTimer3(iei uint8) (x *GPRSTimer3) {
	x = &GPRSTimer3{}
	x.SetIei(iei)
	return x
}

func (a *GPRSTimer3) GetIei() (iei uint8) {
	return a.Iei
}

func (a *GPRSTimer3) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *GPRSTimer3) GetLen() (length uint16) {
	return a.Len
}

func (a *GPRSTimer3) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *GPRSTimer3) GetGPRSTimer3() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *GPRSTimer3) SetGPRSTimer3(contents []uint8) {
	copy(a.Buffer, contents)
}
