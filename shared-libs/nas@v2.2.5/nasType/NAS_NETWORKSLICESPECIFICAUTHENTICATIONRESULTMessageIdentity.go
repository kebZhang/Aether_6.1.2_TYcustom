// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity struct {
	Octet uint8
}

func NewNETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity() (x *NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity) {
	x = &NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity{}
	return x
}

// NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
