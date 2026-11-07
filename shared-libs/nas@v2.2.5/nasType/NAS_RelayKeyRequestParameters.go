// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RelayKeyRequestParameters 9.11.3.89
// RelayKeyRequestParameters Row, sBit, len = [0, INF], 8 , INF
type RelayKeyRequestParameters struct {
	Len    uint16
	Buffer []uint8
}

func NewRelayKeyRequestParameters() (relayKeyRequestParameters *RelayKeyRequestParameters) {
	relayKeyRequestParameters = &RelayKeyRequestParameters{}
	return relayKeyRequestParameters
}

// RelayKeyRequestParameters 9.11.3.89
// Len Row, sBit, len = [], 8, 16
func (a *RelayKeyRequestParameters) GetLen() (length uint16) {
	return a.Len
}

// RelayKeyRequestParameters 9.11.3.89
// Len Row, sBit, len = [], 8, 16
func (a *RelayKeyRequestParameters) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

// RelayKeyRequestParameters 9.11.3.89
// RelayKeyRequestParameters Row, sBit, len = [0, INF], 8 , INF
func (a *RelayKeyRequestParameters) GetRelayKeyRequestParameters() (relayKeyRequestParameters []uint8) {
	relayKeyRequestParameters = make([]uint8, len(a.Buffer))
	copy(relayKeyRequestParameters, a.Buffer)
	return relayKeyRequestParameters
}

// RelayKeyRequestParameters 9.11.3.89
// RelayKeyRequestParameters Row, sBit, len = [0, INF], 8 , INF
func (a *RelayKeyRequestParameters) SetRelayKeyRequestParameters(relayKeyRequestParameters []uint8) {
	copy(a.Buffer, relayKeyRequestParameters)
}
