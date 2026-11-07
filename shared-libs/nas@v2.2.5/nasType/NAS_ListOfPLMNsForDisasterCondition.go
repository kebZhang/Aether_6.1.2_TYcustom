// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ListOfPLMNsForDisasterCondition 9.11.3.83
// ListOfPLMNsForDisasterCondition Row, sBit, len = [0, INF], 8 , INF
type ListOfPLMNsForDisasterCondition struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewListOfPLMNsForDisasterCondition(iei uint8) (x *ListOfPLMNsForDisasterCondition) {
	x = &ListOfPLMNsForDisasterCondition{}
	x.SetIei(iei)
	return x
}

func (a *ListOfPLMNsForDisasterCondition) GetIei() (iei uint8) {
	return a.Iei
}

func (a *ListOfPLMNsForDisasterCondition) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ListOfPLMNsForDisasterCondition) GetLen() (length uint16) {
	return a.Len
}

func (a *ListOfPLMNsForDisasterCondition) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ListOfPLMNsForDisasterCondition) GetListOfPLMNsForDisasterCondition() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *ListOfPLMNsForDisasterCondition) SetListOfPLMNsForDisasterCondition(contents []uint8) {
	copy(a.Buffer, contents)
}
