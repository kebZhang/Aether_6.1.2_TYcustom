// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RELAYAUTHENTICATIONRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RELAYAUTHENTICATIONRESPONSEMessageIdentity struct {
	Octet uint8
}

func NewRELAYAUTHENTICATIONRESPONSEMessageIdentity() (x *RELAYAUTHENTICATIONRESPONSEMessageIdentity) {
	x = &RELAYAUTHENTICATIONRESPONSEMessageIdentity{}
	return x
}

// RELAYAUTHENTICATIONRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYAUTHENTICATIONRESPONSEMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RELAYAUTHENTICATIONRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYAUTHENTICATIONRESPONSEMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
