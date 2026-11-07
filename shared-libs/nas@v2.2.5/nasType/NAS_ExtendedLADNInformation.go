// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ExtendedLADNInformation 9.11.3.96
// ExtendedLADNInformation Row, sBit, len = [0, INF], 8 , INF
type ExtendedLADNInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewExtendedLADNInformation(iei uint8) (x *ExtendedLADNInformation) {
	x = &ExtendedLADNInformation{}
	x.SetIei(iei)
	return x
}

func (a *ExtendedLADNInformation) GetIei() (iei uint8) {
	return a.Iei
}

func (a *ExtendedLADNInformation) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ExtendedLADNInformation) GetLen() (length uint16) {
	return a.Len
}

func (a *ExtendedLADNInformation) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *ExtendedLADNInformation) GetExtendedLADNInformation() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *ExtendedLADNInformation) SetExtendedLADNInformation(contents []uint8) {
	copy(a.Buffer, contents)
}
