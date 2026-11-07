// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PLMNIdentityWithDisasterCondition 9.11.3.85
// PLMNIdentityWithDisasterCondition Row, sBit, len = [0, INF], 8 , INF
type PLMNIdentityWithDisasterCondition struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewPLMNIdentityWithDisasterCondition(iei uint8) (x *PLMNIdentityWithDisasterCondition) {
	x = &PLMNIdentityWithDisasterCondition{}
	x.SetIei(iei)
	return x
}

func (a *PLMNIdentityWithDisasterCondition) GetIei() (iei uint8) {
	return a.Iei
}

func (a *PLMNIdentityWithDisasterCondition) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *PLMNIdentityWithDisasterCondition) GetLen() (length uint16) {
	return a.Len
}

func (a *PLMNIdentityWithDisasterCondition) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PLMNIdentityWithDisasterCondition) GetPLMNIdentityWithDisasterCondition() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *PLMNIdentityWithDisasterCondition) SetPLMNIdentityWithDisasterCondition(contents []uint8) {
	copy(a.Buffer, contents)
}
