// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RELAYAUTHENTICATIONREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RELAYAUTHENTICATIONREQUESTMessageIdentity struct {
	Octet uint8
}

func NewRELAYAUTHENTICATIONREQUESTMessageIdentity() (x *RELAYAUTHENTICATIONREQUESTMessageIdentity) {
	x = &RELAYAUTHENTICATIONREQUESTMessageIdentity{}
	return x
}

// RELAYAUTHENTICATIONREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYAUTHENTICATIONREQUESTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RELAYAUTHENTICATIONREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RELAYAUTHENTICATIONREQUESTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
