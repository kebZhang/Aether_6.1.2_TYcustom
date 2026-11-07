// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// MasterSessionKey 9.11.3.107
type MasterSessionKey struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewMasterSessionKey(iei uint8) (x *MasterSessionKey) {
	x = &MasterSessionKey{}
	x.SetIei(iei)
	return x
}

func (a *MasterSessionKey) GetIei() (iei uint8)     { return a.Iei }
func (a *MasterSessionKey) SetIei(iei uint8)        { a.Iei = iei }
func (a *MasterSessionKey) GetLen() (length uint16) { return a.Len }
func (a *MasterSessionKey) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *MasterSessionKey) GetMasterSessionKey() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}
func (a *MasterSessionKey) SetMasterSessionKey(v []uint8) { copy(a.Buffer, v) }
