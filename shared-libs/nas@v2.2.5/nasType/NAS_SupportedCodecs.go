// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// SupportedCodecs 9.11.3.51A
// SupportedCodecs Row, sBit, len = [0, INF], 8 , INF
type SupportedCodecs struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewSupportedCodecs(iei uint8) (x *SupportedCodecs) {
	x = &SupportedCodecs{}
	x.SetIei(iei)
	return x
}

func (a *SupportedCodecs) GetIei() (iei uint8) {
	return a.Iei
}

func (a *SupportedCodecs) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *SupportedCodecs) GetLen() (length uint16) {
	return a.Len
}

func (a *SupportedCodecs) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *SupportedCodecs) GetSupportedCodecs() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *SupportedCodecs) SetSupportedCodecs(contents []uint8) {
	copy(a.Buffer, contents)
}
