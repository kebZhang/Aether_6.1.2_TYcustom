// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ProSeRelayTransactionIdentity 9.11.3.88
// PRTI Row, sBit, len = [0, 0], 8 , 8
type ProSeRelayTransactionIdentity struct {
	Octet uint8
}

func NewProSeRelayTransactionIdentity() (proSeRelayTransactionIdentity *ProSeRelayTransactionIdentity) {
	proSeRelayTransactionIdentity = &ProSeRelayTransactionIdentity{}
	return proSeRelayTransactionIdentity
}

// ProSeRelayTransactionIdentity 9.11.3.88
// PRTI Row, sBit, len = [0, 0], 8 , 8
func (a *ProSeRelayTransactionIdentity) GetProSeRelayTransactionIdentityValue() (pRTI uint8) {
	return a.Octet
}

// ProSeRelayTransactionIdentity 9.11.3.88
// PRTI Row, sBit, len = [0, 0], 8 , 8
func (a *ProSeRelayTransactionIdentity) SetProSeRelayTransactionIdentityValue(pRTI uint8) {
	a.Octet = pRTI
}
