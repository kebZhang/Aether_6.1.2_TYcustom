// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RelayKeyResponseParameters 9.11.3.90
// RelayKeyResponseParameters Row, sBit, len = [0, INF], 8 , INF
type RelayKeyResponseParameters struct {
	Len    uint16
	Buffer []uint8
}

func NewRelayKeyResponseParameters() (relayKeyResponseParameters *RelayKeyResponseParameters) {
	relayKeyResponseParameters = &RelayKeyResponseParameters{}
	return relayKeyResponseParameters
}

// RelayKeyResponseParameters 9.11.3.90
// Len Row, sBit, len = [], 8, 16
func (a *RelayKeyResponseParameters) GetLen() (length uint16) {
	return a.Len
}

// RelayKeyResponseParameters 9.11.3.90
// Len Row, sBit, len = [], 8, 16
func (a *RelayKeyResponseParameters) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

// RelayKeyResponseParameters 9.11.3.90
// RelayKeyResponseParameters Row, sBit, len = [0, INF], 8 , INF
func (a *RelayKeyResponseParameters) GetRelayKeyResponseParameters() (relayKeyResponseParameters []uint8) {
	relayKeyResponseParameters = make([]uint8, len(a.Buffer))
	copy(relayKeyResponseParameters, a.Buffer)
	return relayKeyResponseParameters
}

// RelayKeyResponseParameters 9.11.3.90
// RelayKeyResponseParameters Row, sBit, len = [0, INF], 8 , INF
func (a *RelayKeyResponseParameters) SetRelayKeyResponseParameters(relayKeyResponseParameters []uint8) {
	copy(a.Buffer, relayKeyResponseParameters)
}
