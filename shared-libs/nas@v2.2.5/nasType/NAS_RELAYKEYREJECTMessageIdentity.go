// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RELAYKEYREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RELAYKEYREJECTMessageIdentity struct {
	Octet uint8
}

func NewRELAYKEYREJECTMessageIdentity() (x *RELAYKEYREJECTMessageIdentity) {
	x = &RELAYKEYREJECTMessageIdentity{}
	return x
}

// RELAYKEYREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYREJECTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RELAYKEYREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYREJECTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
