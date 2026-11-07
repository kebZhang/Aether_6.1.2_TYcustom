// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RELAYKEYACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RELAYKEYACCEPTMessageIdentity struct {
	Octet uint8
}

func NewRELAYKEYACCEPTMessageIdentity() (x *RELAYKEYACCEPTMessageIdentity) {
	x = &RELAYKEYACCEPTMessageIdentity{}
	return x
}

// RELAYKEYACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYACCEPTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RELAYKEYACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYKEYACCEPTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
