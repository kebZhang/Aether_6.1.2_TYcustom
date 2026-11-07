// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// IPHeaderCompressionConfiguration 9.11.4.24
type IPHeaderCompressionConfiguration struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewIPHeaderCompressionConfiguration(iei uint8) (x *IPHeaderCompressionConfiguration) {
	x = &IPHeaderCompressionConfiguration{}
	x.SetIei(iei)
	return x
}

func (a *IPHeaderCompressionConfiguration) GetIei() (iei uint8)    { return a.Iei }
func (a *IPHeaderCompressionConfiguration) SetIei(iei uint8)       { a.Iei = iei }
func (a *IPHeaderCompressionConfiguration) GetLen() (length uint8) { return a.Len }
func (a *IPHeaderCompressionConfiguration) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *IPHeaderCompressionConfiguration) GetIPHeaderCompressionConfiguration() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *IPHeaderCompressionConfiguration) SetIPHeaderCompressionConfiguration(v []uint8) {
	copy(a.Buffer, v)
}
