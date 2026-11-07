// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity struct {
	Octet uint8
}

func NewSERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity() (x *SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity) {
	x = &SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity{}
	return x
}

// SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
