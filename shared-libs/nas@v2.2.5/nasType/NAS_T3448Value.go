// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// T3448Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
type T3448Value struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewT3448Value(iei uint8) (t3448Value *T3448Value) {
	t3448Value = &T3448Value{}
	t3448Value.SetIei(iei)
	return t3448Value
}

// T3448Value 9.11.2.4
// Iei Row, sBit, len = [], 8, 8
func (a *T3448Value) GetIei() (iei uint8) {
	return a.Iei
}

// T3448Value 9.11.2.4
// Iei Row, sBit, len = [], 8, 8
func (a *T3448Value) SetIei(iei uint8) {
	a.Iei = iei
}

// T3448Value 9.11.2.4
// Len Row, sBit, len = [], 8, 8
func (a *T3448Value) GetLen() (length uint8) {
	return a.Len
}

// T3448Value 9.11.2.4
// Len Row, sBit, len = [], 8, 8
func (a *T3448Value) SetLen(length uint8) {
	a.Len = length
}

// T3448Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
func (a *T3448Value) GetGPRSTimer2Value() (gPRSTimer2Value uint8) {
	return a.Octet
}

// T3448Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
func (a *T3448Value) SetGPRSTimer2Value(gPRSTimer2Value uint8) {
	a.Octet = gPRSTimer2Value
}
