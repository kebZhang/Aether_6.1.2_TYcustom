// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// FiveGSMNetworkFeatureSupport 9.11.4.18
type FiveGSMNetworkFeatureSupport struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewFiveGSMNetworkFeatureSupport(iei uint8) (x *FiveGSMNetworkFeatureSupport) {
	x = &FiveGSMNetworkFeatureSupport{}
	x.SetIei(iei)
	return x
}

func (a *FiveGSMNetworkFeatureSupport) GetIei() (iei uint8)    { return a.Iei }
func (a *FiveGSMNetworkFeatureSupport) SetIei(iei uint8)       { a.Iei = iei }
func (a *FiveGSMNetworkFeatureSupport) GetLen() (length uint8) { return a.Len }
func (a *FiveGSMNetworkFeatureSupport) SetLen(length uint8) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

func (a *FiveGSMNetworkFeatureSupport) GetFiveGSMNetworkFeatureSupport() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *FiveGSMNetworkFeatureSupport) SetFiveGSMNetworkFeatureSupport(v []uint8) {
	copy(a.Buffer, v)
}
