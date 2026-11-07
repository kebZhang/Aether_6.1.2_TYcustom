// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// REMOTEUEREPORTRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type REMOTEUEREPORTRESPONSEMessageIdentity struct {
	Octet uint8
}

func NewREMOTEUEREPORTRESPONSEMessageIdentity() (x *REMOTEUEREPORTRESPONSEMessageIdentity) {
	x = &REMOTEUEREPORTRESPONSEMessageIdentity{}
	return x
}

// REMOTEUEREPORTRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *REMOTEUEREPORTRESPONSEMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// REMOTEUEREPORTRESPONSEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *REMOTEUEREPORTRESPONSEMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
