// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PagingRestriction 9.11.3.77
// PagingRestriction Row, sBit, len = [0, INF], 8 , INF
type PagingRestriction struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewPagingRestriction(iei uint8) (x *PagingRestriction) {
	x = &PagingRestriction{}
	x.SetIei(iei)
	return x
}

func (a *PagingRestriction) GetIei() (iei uint8) {
	return a.Iei
}

func (a *PagingRestriction) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *PagingRestriction) GetLen() (length uint16) {
	return a.Len
}

func (a *PagingRestriction) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *PagingRestriction) GetPagingRestriction() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *PagingRestriction) SetPagingRestriction(contents []uint8) {
	copy(a.Buffer, contents)
}
