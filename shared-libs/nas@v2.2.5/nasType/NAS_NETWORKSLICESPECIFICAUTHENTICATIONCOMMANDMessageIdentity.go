// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity struct {
	Octet uint8
}

func NewNETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity() (x *NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity) {
	x = &NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity{}
	return x
}

// NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
