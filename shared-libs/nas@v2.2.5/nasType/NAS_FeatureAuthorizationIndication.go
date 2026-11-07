// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// FeatureAuthorizationIndication 9.11.3.105
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// MBSRAI Row, sBit, len = [0, 0], 2, 2
type FeatureAuthorizationIndication struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewFeatureAuthorizationIndication(iei uint8) (x *FeatureAuthorizationIndication) {
	x = &FeatureAuthorizationIndication{}
	x.SetIei(iei)
	return x
}

func (a *FeatureAuthorizationIndication) GetIei() (iei uint8)    { return a.Iei }
func (a *FeatureAuthorizationIndication) SetIei(iei uint8)       { a.Iei = iei }
func (a *FeatureAuthorizationIndication) GetLen() (length uint8) { return a.Len }
func (a *FeatureAuthorizationIndication) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *FeatureAuthorizationIndication) GetFeatureAuthorizationIndication() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *FeatureAuthorizationIndication) SetFeatureAuthorizationIndication(v []uint8) {
	copy(a.Buffer, v)
}

func (a *FeatureAuthorizationIndication) GetMBSRAI() uint8 {
	if len(a.Buffer) == 0 {
		return 0
	}
	return a.Buffer[0] & GetBitMask(2, 0)
}

func (a *FeatureAuthorizationIndication) SetMBSRAI(mbsrai uint8) {
	if len(a.Buffer) > 0 {
		a.Buffer[0] = (a.Buffer[0] & 0xFC) | (mbsrai & 0x03)
	}
}
