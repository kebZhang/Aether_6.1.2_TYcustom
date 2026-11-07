// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// SNSSAITimeValidityInformation 9.11.3.101
type SNSSAITimeValidityInformation struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewSNSSAITimeValidityInformation(iei uint8) (x *SNSSAITimeValidityInformation) {
	x = &SNSSAITimeValidityInformation{}
	x.SetIei(iei)
	return x
}

func (a *SNSSAITimeValidityInformation) GetIei() (iei uint8)    { return a.Iei }
func (a *SNSSAITimeValidityInformation) SetIei(iei uint8)       { a.Iei = iei }
func (a *SNSSAITimeValidityInformation) GetLen() (length uint8) { return a.Len }
func (a *SNSSAITimeValidityInformation) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *SNSSAITimeValidityInformation) GetSNSSAITimeValidityInformation() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *SNSSAITimeValidityInformation) SetSNSSAITimeValidityInformation(v []uint8) {
	copy(a.Buffer, v)
}
