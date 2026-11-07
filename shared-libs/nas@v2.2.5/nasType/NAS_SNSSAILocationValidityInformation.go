// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// SNSSAILocationValidityInformation 9.11.3.100
type SNSSAILocationValidityInformation struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewSNSSAILocationValidityInformation(iei uint8) (x *SNSSAILocationValidityInformation) {
	x = &SNSSAILocationValidityInformation{}
	x.SetIei(iei)
	return x
}

func (a *SNSSAILocationValidityInformation) GetIei() (iei uint8)    { return a.Iei }
func (a *SNSSAILocationValidityInformation) SetIei(iei uint8)       { a.Iei = iei }
func (a *SNSSAILocationValidityInformation) GetLen() (length uint8) { return a.Len }
func (a *SNSSAILocationValidityInformation) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *SNSSAILocationValidityInformation) GetSNSSAILocationValidityInformation() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *SNSSAILocationValidityInformation) SetSNSSAILocationValidityInformation(v []uint8) {
	copy(a.Buffer, v)
}
