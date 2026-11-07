// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity struct {
	Octet uint8
}

func NewNETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity() (x *NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity) {
	x = &NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity{}
	return x
}

// NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
