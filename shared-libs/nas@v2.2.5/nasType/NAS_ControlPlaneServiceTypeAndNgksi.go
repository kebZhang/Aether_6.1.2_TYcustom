// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ControlPlaneServiceTypeAndNgksi 9.11.3.18D 9.11.3.32
// ControlPlaneServiceType Row, sBit, len = [0, 0], 8 , 4
// TSC Row, sBit, len = [0, 0], 4 , 1
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 3 , 3
type ControlPlaneServiceTypeAndNgksi struct {
	Octet uint8
}

func NewControlPlaneServiceTypeAndNgksi() (controlPlaneServiceTypeAndNgksi *ControlPlaneServiceTypeAndNgksi) {
	controlPlaneServiceTypeAndNgksi = &ControlPlaneServiceTypeAndNgksi{}
	return controlPlaneServiceTypeAndNgksi
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.18D
// ControlPlaneServiceType Row, sBit, len = [0, 0], 8 , 4
func (a *ControlPlaneServiceTypeAndNgksi) GetControlPlaneServiceType() (controlPlaneServiceType uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.18D
// ControlPlaneServiceType Row, sBit, len = [0, 0], 8 , 4
func (a *ControlPlaneServiceTypeAndNgksi) SetControlPlaneServiceType(controlPlaneServiceType uint8) {
	a.Octet = (a.Octet & 15) + ((controlPlaneServiceType & 15) << 4)
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.32
// TSC Row, sBit, len = [0, 0], 4 , 1
func (a *ControlPlaneServiceTypeAndNgksi) GetTSC() (tSC uint8) {
	return a.Octet & GetBitMask(4, 3) >> 3
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.32
// TSC Row, sBit, len = [0, 0], 4 , 1
func (a *ControlPlaneServiceTypeAndNgksi) SetTSC(tSC uint8) {
	a.Octet = (a.Octet & 247) + ((tSC & 1) << 3)
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.32
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 3 , 3
func (a *ControlPlaneServiceTypeAndNgksi) GetNasKeySetIdentifiler() (nasKeySetIdentifiler uint8) {
	return a.Octet & GetBitMask(3, 0)
}

// ControlPlaneServiceTypeAndNgksi 9.11.3.32
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 3 , 3
func (a *ControlPlaneServiceTypeAndNgksi) SetNasKeySetIdentifiler(nasKeySetIdentifiler uint8) {
	a.Octet = (a.Octet & 248) + (nasKeySetIdentifiler & 7)
}
