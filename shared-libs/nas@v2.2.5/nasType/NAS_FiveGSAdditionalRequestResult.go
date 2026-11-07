// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// FiveGSAdditionalRequestResult 9.11.3.81
// FiveGSAdditionalRequestResult Row, sBit, len = [0, INF], 8 , INF
type FiveGSAdditionalRequestResult struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewFiveGSAdditionalRequestResult(iei uint8) (x *FiveGSAdditionalRequestResult) {
	x = &FiveGSAdditionalRequestResult{}
	x.SetIei(iei)
	return x
}

func (a *FiveGSAdditionalRequestResult) GetIei() (iei uint8) {
	return a.Iei
}

func (a *FiveGSAdditionalRequestResult) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *FiveGSAdditionalRequestResult) GetLen() (length uint16) {
	return a.Len
}

func (a *FiveGSAdditionalRequestResult) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *FiveGSAdditionalRequestResult) GetFiveGSAdditionalRequestResult() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *FiveGSAdditionalRequestResult) SetFiveGSAdditionalRequestResult(contents []uint8) {
	copy(a.Buffer, contents)
}
