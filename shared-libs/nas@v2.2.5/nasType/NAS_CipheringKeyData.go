// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// CipheringKeyData 9.11.3.18C
// CipheringKeyData Row, sBit, len = [0, INF], 8 , INF
type CipheringKeyData struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewCipheringKeyData(iei uint8) (x *CipheringKeyData) {
	x = &CipheringKeyData{}
	x.SetIei(iei)
	return x
}

func (a *CipheringKeyData) GetIei() (iei uint8) {
	return a.Iei
}

func (a *CipheringKeyData) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *CipheringKeyData) GetLen() (length uint16) {
	return a.Len
}

func (a *CipheringKeyData) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *CipheringKeyData) GetCipheringKeyData() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *CipheringKeyData) SetCipheringKeyData(contents []uint8) {
	copy(a.Buffer, contents)
}
