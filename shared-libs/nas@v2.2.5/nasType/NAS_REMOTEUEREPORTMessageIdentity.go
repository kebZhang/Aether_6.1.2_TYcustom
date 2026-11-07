// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// REMOTEUEREPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type REMOTEUEREPORTMessageIdentity struct {
	Octet uint8
}

func NewREMOTEUEREPORTMessageIdentity() (x *REMOTEUEREPORTMessageIdentity) {
	x = &REMOTEUEREPORTMessageIdentity{}
	return x
}

// REMOTEUEREPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *REMOTEUEREPORTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// REMOTEUEREPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *REMOTEUEREPORTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
