// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RELAYKEYREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RELAYKEYREQUESTMessageIdentity struct {
	Octet uint8
}

func NewRELAYKEYREQUESTMessageIdentity() (x *RELAYKEYREQUESTMessageIdentity) {
	x = &RELAYKEYREQUESTMessageIdentity{}
	return x
}

// RELAYKEYREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYREQUESTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RELAYKEYREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYREQUESTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
