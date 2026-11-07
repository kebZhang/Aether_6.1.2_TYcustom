// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type AuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationRequestMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.ABBA
	*nasType.AuthenticationParameterRAND
	*nasType.AuthenticationParameterAUTN
	*nasType.EAPMessage
}

func NewAuthenticationRequest(iei uint8) (authenticationRequest *AuthenticationRequest) {
	authenticationRequest = &AuthenticationRequest{}
	return authenticationRequest
}

const (
	AuthenticationRequestAuthenticationParameterRANDType uint8 = 0x21
	AuthenticationRequestAuthenticationParameterAUTNType uint8 = 0x20
	AuthenticationRequestEAPMessageType                  uint8 = 0x78
)

func (a *AuthenticationRequest) EncodeAuthenticationRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer); err != nil {
		return
	}
	if a.AuthenticationParameterRAND != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterRAND.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet); err != nil {
			return
		}
	}
	if a.AuthenticationParameterAUTN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()]); err != nil {
			return
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
			return
		}
	}
}

func (a *AuthenticationRequest) DecodeAuthenticationRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
		return
	}
	a.ABBA.SetLen(a.ABBA.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Buffer); err != nil {
		return
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return
		}
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		switch tmpIeiN {
		case AuthenticationRequestAuthenticationParameterRANDType:
			a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet); err != nil {
				return
			}
		case AuthenticationRequestAuthenticationParameterAUTNType:
			a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterAUTN.Len); err != nil {
				return
			}
			a.AuthenticationParameterAUTN.SetLen(a.AuthenticationParameterAUTN.GetLen())
			if a.AuthenticationParameterAUTN.GetLen() > uint8(len(a.AuthenticationParameterAUTN.Octet)) {
				a.AuthenticationParameterAUTN = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return
			}
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()]); err != nil {
				return
			}
		case AuthenticationRequestEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return
			}
		default:
		}
	}
}
