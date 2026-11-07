// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// CONTROLPLANESERVICEREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type CONTROLPLANESERVICEREQUESTMessageIdentity struct {
	Octet uint8
}

func NewCONTROLPLANESERVICEREQUESTMessageIdentity() (cONTROLPLANESERVICEREQUESTMessageIdentity *CONTROLPLANESERVICEREQUESTMessageIdentity) {
	cONTROLPLANESERVICEREQUESTMessageIdentity = &CONTROLPLANESERVICEREQUESTMessageIdentity{}
	return cONTROLPLANESERVICEREQUESTMessageIdentity
}

// CONTROLPLANESERVICEREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *CONTROLPLANESERVICEREQUESTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// CONTROLPLANESERVICEREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *CONTROLPLANESERVICEREQUESTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
