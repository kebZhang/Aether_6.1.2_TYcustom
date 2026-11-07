// Copyright 2019 free5GC.org
// SPDX-FileCopyrightText: 2025 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
type DNN struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewDNN(iei uint8) (dNN *DNN) {
	dNN = &DNN{}
	dNN.SetIei(iei)
	return dNN
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) GetIei() (iei uint8) {
	return a.Iei
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) SetIei(iei uint8) {
	a.Iei = iei
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) GetLen() (length uint8) {
	return a.Len
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) GetDNN() (dNN []uint8) {
	if len(a.Buffer) == 0 {
		return nil
	}
	dnn := new(Dnn)
	if err := dnn.UnmarshalBinary(a.Buffer); err != nil {
		return nil
	}
	return *dnn
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) SetDNN(dNN []uint8) {
	tmp := Dnn(dNN)
	dnn, err := tmp.MarshalBinary()
	if err != nil {
		return
	}
	a.Buffer = dnn
	a.Len = uint8(len(a.Buffer))
}
