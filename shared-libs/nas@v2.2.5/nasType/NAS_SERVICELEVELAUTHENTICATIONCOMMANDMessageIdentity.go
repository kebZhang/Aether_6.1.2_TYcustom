// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity struct {
	Octet uint8
}

func NewSERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity() (x *SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity) {
	x = &SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity{}
	return x
}

// SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
