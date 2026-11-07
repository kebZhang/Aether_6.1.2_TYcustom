// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nas

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasMessage"
)

// Message TODO：description
type Message struct {
	SecurityHeader
	*GmmMessage
	*GsmMessage
}

// SecurityHeader TODO：description
type SecurityHeader struct {
	ProtocolDiscriminator     uint8
	SecurityHeaderType        uint8
	MessageAuthenticationCode uint32
	SequenceNumber            uint8
}

const (
	SecurityHeaderTypePlainNas                                                 uint8 = 0x00
	SecurityHeaderTypeIntegrityProtected                                       uint8 = 0x01
	SecurityHeaderTypeIntegrityProtectedAndCiphered                            uint8 = 0x02
	SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext            uint8 = 0x03
	SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext uint8 = 0x04
)

// NewMessage TODO:desc
func NewMessage() *Message {
	Message := &Message{}
	return Message
}

// NewGmmMessage TODO:desc
func NewGmmMessage() *GmmMessage {
	GmmMessage := &GmmMessage{}
	return GmmMessage
}

// NewGmmMessage TODO:desc
func NewGsmMessage() *GsmMessage {
	GsmMessage := &GsmMessage{}
	return GsmMessage
}

// GmmHeader Octet1 protocolDiscriminator securityHeaderType
//           Octet2 MessageType

type GmmHeader struct {
	Octet [3]uint8
}

type GsmHeader struct {
	Octet [4]uint8
}

// GetMessageType 9.8
func (a *GmmHeader) GetMessageType() (messageType uint8) {
	messageType = a.Octet[2]
	return messageType
}

// GetMessageType 9.8
func (a *GmmHeader) SetMessageType(messageType uint8) {
	a.Octet[2] = messageType
}

func (a *GmmHeader) GetExtendedProtocolDiscriminator() uint8 {
	return a.Octet[0]
}

func (a *GmmHeader) SetExtendedProtocolDiscriminator(epd uint8) {
	a.Octet[0] = epd
}

func (a *GsmHeader) GetExtendedProtocolDiscriminator() uint8 {
	return a.Octet[0]
}

func (a *GsmHeader) SetExtendedProtocolDiscriminator(epd uint8) {
	a.Octet[0] = epd
}

// GetMessageType 9.8
func (a *GsmHeader) GetMessageType() (messageType uint8) {
	messageType = a.Octet[3]
	return messageType
}

// GetMessageType 9.8
func (a *GsmHeader) SetMessageType(messageType uint8) {
	a.Octet[3] = messageType
}

func GetEPD(byteArray []byte) uint8 {
	return byteArray[0]
}

func GetSecurityHeaderType(byteArray []byte) uint8 {
	return byteArray[1]
}

type GmmMessage struct {
	GmmHeader
	*nasMessage.AuthenticationRequest                            // 8.2.1
	*nasMessage.AuthenticationResponse                           // 8.2.2
	*nasMessage.AuthenticationResult                             // 8.2.3
	*nasMessage.AuthenticationFailure                            // 8.2.4
	*nasMessage.AuthenticationReject                             // 8.2.5
	*nasMessage.RegistrationRequest                              // 8.2.6
	*nasMessage.RegistrationAccept                               // 8.2.7
	*nasMessage.RegistrationComplete                             // 8.2.8
	*nasMessage.RegistrationReject                               // 8.2.9
	*nasMessage.ULNASTransport                                   // 8.2.10
	*nasMessage.DLNASTransport                                   // 8.2.11
	*nasMessage.DeregistrationRequestUEOriginatingDeregistration // 8.2.12
	*nasMessage.DeregistrationAcceptUEOriginatingDeregistration  // 8.2.13
	*nasMessage.DeregistrationRequestUETerminatedDeregistration  // 8.2.14
	*nasMessage.DeregistrationAcceptUETerminatedDeregistration   // 8.2.15
	*nasMessage.ServiceRequest                                   // 8.2.16
	*nasMessage.ServiceAccept                                    // 8.2.17
	*nasMessage.ServiceReject                                    // 8.2.18
	*nasMessage.ConfigurationUpdateCommand                       // 8.2.19
	*nasMessage.ConfigurationUpdateComplete                      // 8.2.20
	*nasMessage.IdentityRequest                                  // 8.2.21
	*nasMessage.IdentityResponse                                 // 8.2.22
	*nasMessage.Notification                                     // 8.2.23
	*nasMessage.NotificationResponse                             // 8.2.24
	*nasMessage.SecurityModeCommand                              // 8.2.25
	*nasMessage.SecurityModeComplete                             // 8.2.26
	*nasMessage.SecurityModeReject                               // 8.2.27
	*nasMessage.SecurityProtected5GSNASMessage                   // 8.2.28
	*nasMessage.Status5GMM                                       // 8.2.29
	*nasMessage.ControlPlaneServiceRequest                       // 8.2.30
	*nasMessage.NetworkSliceSpecificAuthenticationCommand        // 8.2.31
	*nasMessage.NetworkSliceSpecificAuthenticationComplete       // 8.2.32
	*nasMessage.NetworkSliceSpecificAuthenticationResult         // 8.2.33
	*nasMessage.RelayKeyRequest                                  // 8.2.34
	*nasMessage.RelayKeyAccept                                   // 8.2.35
	*nasMessage.RelayKeyReject                                   // 8.2.36
	*nasMessage.RelayAuthenticationRequest                       // 8.2.37
	*nasMessage.RelayAuthenticationResponse                      // 8.2.38
}

const (
	MsgTypeRegistrationRequest                              uint8 = 65
	MsgTypeRegistrationAccept                               uint8 = 66
	MsgTypeRegistrationComplete                             uint8 = 67
	MsgTypeRegistrationReject                               uint8 = 68
	MsgTypeDeregistrationRequestUEOriginatingDeregistration uint8 = 69
	MsgTypeDeregistrationAcceptUEOriginatingDeregistration  uint8 = 70
	MsgTypeDeregistrationRequestUETerminatedDeregistration  uint8 = 71
	MsgTypeDeregistrationAcceptUETerminatedDeregistration   uint8 = 72
	MsgTypeServiceRequest                                   uint8 = 76
	MsgTypeServiceReject                                    uint8 = 77
	MsgTypeServiceAccept                                    uint8 = 78
	MsgTypeConfigurationUpdateCommand                       uint8 = 84
	MsgTypeConfigurationUpdateComplete                      uint8 = 85
	MsgTypeAuthenticationRequest                            uint8 = 86
	MsgTypeAuthenticationResponse                           uint8 = 87
	MsgTypeAuthenticationReject                             uint8 = 88
	MsgTypeAuthenticationFailure                            uint8 = 89
	MsgTypeAuthenticationResult                             uint8 = 90
	MsgTypeIdentityRequest                                  uint8 = 91
	MsgTypeIdentityResponse                                 uint8 = 92
	MsgTypeSecurityModeCommand                              uint8 = 93
	MsgTypeSecurityModeComplete                             uint8 = 94
	MsgTypeSecurityModeReject                               uint8 = 95
	MsgTypeStatus5GMM                                       uint8 = 100
	MsgTypeNotification                                     uint8 = 101
	MsgTypeNotificationResponse                             uint8 = 102
	MsgTypeULNASTransport                                   uint8 = 103
	MsgTypeDLNASTransport                                   uint8 = 104
	MsgTypeControlPlaneServiceRequest                       uint8 = 79
	MsgTypeNetworkSliceSpecificAuthenticationCommand        uint8 = 80
	MsgTypeNetworkSliceSpecificAuthenticationComplete       uint8 = 81
	MsgTypeNetworkSliceSpecificAuthenticationResult         uint8 = 82
	MsgTypeRelayKeyRequest                                  uint8 = 105
	MsgTypeRelayKeyAccept                                   uint8 = 106
	MsgTypeRelayKeyReject                                   uint8 = 107
	MsgTypeRelayAuthenticationRequest                       uint8 = 108
	MsgTypeRelayAuthenticationResponse                      uint8 = 109
)

func MessageName(code uint8) string {
	switch code {
	case MsgTypeRegistrationRequest:
		return "RegistrationRequest"
	case MsgTypeRegistrationAccept:
		return "RegistrationAccept"
	case MsgTypeRegistrationComplete:
		return "RegistrationComplete"
	case MsgTypeRegistrationReject:
		return "RegistrationReject"
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		return "DeregistrationRequestUEOriginatingDeregistration"
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		return "DeregistrationAcceptUEOriginatingDeregistration"
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		return "DeregistrationRequestUETerminatedDeregistration"
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		return "DeregistrationAcceptUETerminatedDeregistration"
	case MsgTypeServiceRequest:
		return "ServiceRequest"
	case MsgTypeServiceReject:
		return "ServiceReject"
	case MsgTypeServiceAccept:
		return "ServiceAccept"
	case MsgTypeConfigurationUpdateCommand:
		return "ConfigurationUpdateCommand"
	case MsgTypeConfigurationUpdateComplete:
		return "ConfigurationUpdateComplete"
	case MsgTypeAuthenticationRequest:
		return "AuthenticationRequest"
	case MsgTypeAuthenticationResponse:
		return "AuthenticationResponse"
	case MsgTypeAuthenticationReject:
		return "AuthenticationReject"
	case MsgTypeAuthenticationFailure:
		return "AuthenticationFailure"
	case MsgTypeAuthenticationResult:
		return "AuthenticationResult"
	case MsgTypeIdentityRequest:
		return "IdentityRequest"
	case MsgTypeIdentityResponse:
		return "IdentityResponse"
	case MsgTypeSecurityModeCommand:
		return "SecurityModeCommand"
	case MsgTypeSecurityModeComplete:
		return "SecurityModeComplete"
	case MsgTypeSecurityModeReject:
		return "SecurityModeReject"
	case MsgTypeStatus5GMM:
		return "Status5GMM"
	case MsgTypeNotification:
		return "Notification"
	case MsgTypeNotificationResponse:
		return "NotificationResponse"
	case MsgTypeULNASTransport:
		return "ULNASTransport"
	case MsgTypeDLNASTransport:
		return "DLNASTransport"
	case MsgTypeControlPlaneServiceRequest:
		return "ControlPlaneServiceRequest"
	case MsgTypeNetworkSliceSpecificAuthenticationCommand:
		return "NetworkSliceSpecificAuthenticationCommand"
	case MsgTypeNetworkSliceSpecificAuthenticationComplete:
		return "NetworkSliceSpecificAuthenticationComplete"
	case MsgTypeNetworkSliceSpecificAuthenticationResult:
		return "NetworkSliceSpecificAuthenticationResult"
	case MsgTypeRelayKeyRequest:
		return "RelayKeyRequest"
	case MsgTypeRelayKeyAccept:
		return "RelayKeyAccept"
	case MsgTypeRelayKeyReject:
		return "RelayKeyReject"
	case MsgTypeRelayAuthenticationRequest:
		return "RelayAuthenticationRequest"
	case MsgTypeRelayAuthenticationResponse:
		return "RelayAuthenticationResponse"
	case MsgTypeServiceLevelAuthenticationCommand:
		return "ServiceLevelAuthenticationCommand"
	case MsgTypeServiceLevelAuthenticationComplete:
		return "ServiceLevelAuthenticationComplete"
	case MsgTypeRemoteUEReport:
		return "RemoteUEReport"
	case MsgTypeRemoteUEReportResponse:
		return "RemoteUEReportResponse"
	default:
		return fmt.Sprintf("Unknown message type: %d", code)
	}
}

func (a *Message) PlainNasDecode(byteArray *[]byte) error {
	epd := GetEPD(*byteArray)
	switch epd {
	case nasMessage.Epd5GSMobilityManagementMessage:
		return a.GmmMessageDecode(byteArray)
	case nasMessage.Epd5GSSessionManagementMessage:
		return a.GsmMessageDecode(byteArray)
	}
	return fmt.Errorf("extended Protocol Discriminator[%d] is not allowed in Nas Message Decode", epd)
}

func (a *Message) PlainNasEncode() ([]byte, error) {
	data := new(bytes.Buffer)
	if a.GmmMessage != nil {
		err := a.GmmMessageEncode(data)
		return data.Bytes(), err
	} else if a.GsmMessage != nil {
		err := a.GsmMessageEncode(data)
		return data.Bytes(), err
	}
	return nil, fmt.Errorf("Gmm/Gsm Message are both empty in Nas Message Encode")
}

func (a *Message) GmmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GmmMessage = NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GmmHeader); err != nil {
		return err
	}
	switch a.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		a.RegistrationRequest = nasMessage.NewRegistrationRequest(MsgTypeRegistrationRequest)
		a.DecodeRegistrationRequest(byteArray)
	case MsgTypeRegistrationAccept:
		a.RegistrationAccept = nasMessage.NewRegistrationAccept(MsgTypeRegistrationAccept)
		a.DecodeRegistrationAccept(byteArray)
	case MsgTypeRegistrationComplete:
		a.RegistrationComplete = nasMessage.NewRegistrationComplete(MsgTypeRegistrationComplete)
		a.DecodeRegistrationComplete(byteArray)
	case MsgTypeRegistrationReject:
		a.RegistrationReject = nasMessage.NewRegistrationReject(MsgTypeRegistrationReject)
		a.DecodeRegistrationReject(byteArray)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.DeregistrationRequestUEOriginatingDeregistration = nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(MsgTypeDeregistrationRequestUEOriginatingDeregistration)
		a.DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.DeregistrationAcceptUEOriginatingDeregistration = nasMessage.NewDeregistrationAcceptUEOriginatingDeregistration(MsgTypeDeregistrationAcceptUEOriginatingDeregistration)
		a.DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.DeregistrationRequestUETerminatedDeregistration = nasMessage.NewDeregistrationRequestUETerminatedDeregistration(MsgTypeDeregistrationRequestUETerminatedDeregistration)
		a.DecodeDeregistrationRequestUETerminatedDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.DeregistrationAcceptUETerminatedDeregistration = nasMessage.NewDeregistrationAcceptUETerminatedDeregistration(MsgTypeDeregistrationAcceptUETerminatedDeregistration)
		a.DecodeDeregistrationAcceptUETerminatedDeregistration(byteArray)
	case MsgTypeServiceRequest:
		a.ServiceRequest = nasMessage.NewServiceRequest(MsgTypeServiceRequest)
		a.DecodeServiceRequest(byteArray)
	case MsgTypeServiceReject:
		a.ServiceReject = nasMessage.NewServiceReject(MsgTypeServiceReject)
		a.DecodeServiceReject(byteArray)
	case MsgTypeServiceAccept:
		a.ServiceAccept = nasMessage.NewServiceAccept(MsgTypeServiceAccept)
		a.DecodeServiceAccept(byteArray)
	case MsgTypeConfigurationUpdateCommand:
		a.ConfigurationUpdateCommand = nasMessage.NewConfigurationUpdateCommand(MsgTypeConfigurationUpdateCommand)
		a.DecodeConfigurationUpdateCommand(byteArray)
	case MsgTypeConfigurationUpdateComplete:
		a.ConfigurationUpdateComplete = nasMessage.NewConfigurationUpdateComplete(MsgTypeConfigurationUpdateComplete)
		a.DecodeConfigurationUpdateComplete(byteArray)
	case MsgTypeAuthenticationRequest:
		a.AuthenticationRequest = nasMessage.NewAuthenticationRequest(MsgTypeAuthenticationRequest)
		a.DecodeAuthenticationRequest(byteArray)
	case MsgTypeAuthenticationResponse:
		a.AuthenticationResponse = nasMessage.NewAuthenticationResponse(MsgTypeAuthenticationResponse)
		a.DecodeAuthenticationResponse(byteArray)
	case MsgTypeAuthenticationReject:
		a.AuthenticationReject = nasMessage.NewAuthenticationReject(MsgTypeAuthenticationReject)
		a.DecodeAuthenticationReject(byteArray)
	case MsgTypeAuthenticationFailure:
		a.AuthenticationFailure = nasMessage.NewAuthenticationFailure(MsgTypeAuthenticationFailure)
		a.DecodeAuthenticationFailure(byteArray)
	case MsgTypeAuthenticationResult:
		a.AuthenticationResult = nasMessage.NewAuthenticationResult(MsgTypeAuthenticationResult)
		a.DecodeAuthenticationResult(byteArray)
	case MsgTypeIdentityRequest:
		a.IdentityRequest = nasMessage.NewIdentityRequest(MsgTypeIdentityRequest)
		a.DecodeIdentityRequest(byteArray)
	case MsgTypeIdentityResponse:
		a.IdentityResponse = nasMessage.NewIdentityResponse(MsgTypeIdentityResponse)
		a.DecodeIdentityResponse(byteArray)
	case MsgTypeSecurityModeCommand:
		a.SecurityModeCommand = nasMessage.NewSecurityModeCommand(MsgTypeSecurityModeCommand)
		a.DecodeSecurityModeCommand(byteArray)
	case MsgTypeSecurityModeComplete:
		a.SecurityModeComplete = nasMessage.NewSecurityModeComplete(MsgTypeSecurityModeComplete)
		a.DecodeSecurityModeComplete(byteArray)
	case MsgTypeSecurityModeReject:
		a.SecurityModeReject = nasMessage.NewSecurityModeReject(MsgTypeSecurityModeReject)
		a.DecodeSecurityModeReject(byteArray)
	case MsgTypeStatus5GMM:
		a.Status5GMM = nasMessage.NewStatus5GMM(MsgTypeStatus5GMM)
		a.DecodeStatus5GMM(byteArray)
	case MsgTypeNotification:
		a.Notification = nasMessage.NewNotification(MsgTypeNotification)
		a.DecodeNotification(byteArray)
	case MsgTypeNotificationResponse:
		a.NotificationResponse = nasMessage.NewNotificationResponse(MsgTypeNotificationResponse)
		a.DecodeNotificationResponse(byteArray)
	case MsgTypeULNASTransport:
		a.ULNASTransport = nasMessage.NewULNASTransport(MsgTypeULNASTransport)
		a.DecodeULNASTransport(byteArray)
	case MsgTypeDLNASTransport:
		a.DLNASTransport = nasMessage.NewDLNASTransport(MsgTypeDLNASTransport)
		a.DecodeDLNASTransport(byteArray)
	case MsgTypeControlPlaneServiceRequest:
		a.ControlPlaneServiceRequest = nasMessage.NewControlPlaneServiceRequest(MsgTypeControlPlaneServiceRequest)
		a.DecodeControlPlaneServiceRequest(byteArray)
	case MsgTypeNetworkSliceSpecificAuthenticationCommand:
		a.NetworkSliceSpecificAuthenticationCommand = nasMessage.NewNetworkSliceSpecificAuthenticationCommand(MsgTypeNetworkSliceSpecificAuthenticationCommand)
		a.DecodeNetworkSliceSpecificAuthenticationCommand(byteArray)
	case MsgTypeNetworkSliceSpecificAuthenticationComplete:
		a.NetworkSliceSpecificAuthenticationComplete = nasMessage.NewNetworkSliceSpecificAuthenticationComplete(MsgTypeNetworkSliceSpecificAuthenticationComplete)
		a.DecodeNetworkSliceSpecificAuthenticationComplete(byteArray)
	case MsgTypeNetworkSliceSpecificAuthenticationResult:
		a.NetworkSliceSpecificAuthenticationResult = nasMessage.NewNetworkSliceSpecificAuthenticationResult(MsgTypeNetworkSliceSpecificAuthenticationResult)
		a.DecodeNetworkSliceSpecificAuthenticationResult(byteArray)
	case MsgTypeRelayKeyRequest:
		a.RelayKeyRequest = nasMessage.NewRelayKeyRequest(MsgTypeRelayKeyRequest)
		a.DecodeRelayKeyRequest(byteArray)
	case MsgTypeRelayKeyAccept:
		a.RelayKeyAccept = nasMessage.NewRelayKeyAccept(MsgTypeRelayKeyAccept)
		a.DecodeRelayKeyAccept(byteArray)
	case MsgTypeRelayKeyReject:
		a.RelayKeyReject = nasMessage.NewRelayKeyReject(MsgTypeRelayKeyReject)
		a.DecodeRelayKeyReject(byteArray)
	case MsgTypeRelayAuthenticationRequest:
		a.RelayAuthenticationRequest = nasMessage.NewRelayAuthenticationRequest(MsgTypeRelayAuthenticationRequest)
		a.DecodeRelayAuthenticationRequest(byteArray)
	case MsgTypeRelayAuthenticationResponse:
		a.RelayAuthenticationResponse = nasMessage.NewRelayAuthenticationResponse(MsgTypeRelayAuthenticationResponse)
		a.DecodeRelayAuthenticationResponse(byteArray)
	default:
		return fmt.Errorf("NAS decode fail: MsgType[%d] does not exist in GMM Message", a.GmmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GmmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		a.EncodeRegistrationRequest(buffer)
	case MsgTypeRegistrationAccept:
		a.EncodeRegistrationAccept(buffer)
	case MsgTypeRegistrationComplete:
		a.EncodeRegistrationComplete(buffer)
	case MsgTypeRegistrationReject:
		a.EncodeRegistrationReject(buffer)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.EncodeDeregistrationRequestUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.EncodeDeregistrationRequestUETerminatedDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.EncodeDeregistrationAcceptUETerminatedDeregistration(buffer)
	case MsgTypeServiceRequest:
		a.EncodeServiceRequest(buffer)
	case MsgTypeServiceReject:
		a.EncodeServiceReject(buffer)
	case MsgTypeServiceAccept:
		a.EncodeServiceAccept(buffer)
	case MsgTypeConfigurationUpdateCommand:
		a.EncodeConfigurationUpdateCommand(buffer)
	case MsgTypeConfigurationUpdateComplete:
		a.EncodeConfigurationUpdateComplete(buffer)
	case MsgTypeAuthenticationRequest:
		a.EncodeAuthenticationRequest(buffer)
	case MsgTypeAuthenticationResponse:
		a.EncodeAuthenticationResponse(buffer)
	case MsgTypeAuthenticationReject:
		a.EncodeAuthenticationReject(buffer)
	case MsgTypeAuthenticationFailure:
		a.EncodeAuthenticationFailure(buffer)
	case MsgTypeAuthenticationResult:
		a.EncodeAuthenticationResult(buffer)
	case MsgTypeIdentityRequest:
		a.EncodeIdentityRequest(buffer)
	case MsgTypeIdentityResponse:
		a.EncodeIdentityResponse(buffer)
	case MsgTypeSecurityModeCommand:
		a.EncodeSecurityModeCommand(buffer)
	case MsgTypeSecurityModeComplete:
		a.EncodeSecurityModeComplete(buffer)
	case MsgTypeSecurityModeReject:
		a.EncodeSecurityModeReject(buffer)
	case MsgTypeStatus5GMM:
		a.EncodeStatus5GMM(buffer)
	case MsgTypeNotification:
		a.EncodeNotification(buffer)
	case MsgTypeNotificationResponse:
		a.EncodeNotificationResponse(buffer)
	case MsgTypeULNASTransport:
		a.EncodeULNASTransport(buffer)
	case MsgTypeDLNASTransport:
		a.EncodeDLNASTransport(buffer)
	case MsgTypeControlPlaneServiceRequest:
		a.EncodeControlPlaneServiceRequest(buffer)
	case MsgTypeNetworkSliceSpecificAuthenticationCommand:
		a.EncodeNetworkSliceSpecificAuthenticationCommand(buffer)
	case MsgTypeNetworkSliceSpecificAuthenticationComplete:
		a.EncodeNetworkSliceSpecificAuthenticationComplete(buffer)
	case MsgTypeNetworkSliceSpecificAuthenticationResult:
		a.EncodeNetworkSliceSpecificAuthenticationResult(buffer)
	case MsgTypeRelayKeyRequest:
		a.EncodeRelayKeyRequest(buffer)
	case MsgTypeRelayKeyAccept:
		a.EncodeRelayKeyAccept(buffer)
	case MsgTypeRelayKeyReject:
		a.EncodeRelayKeyReject(buffer)
	case MsgTypeRelayAuthenticationRequest:
		a.EncodeRelayAuthenticationRequest(buffer)
	case MsgTypeRelayAuthenticationResponse:
		a.EncodeRelayAuthenticationResponse(buffer)
	default:
		return fmt.Errorf("NAS encode fail: MsgType[%d] does not exist in GMM Message", a.GmmHeader.GetMessageType())
	}
	return nil
}

type GsmMessage struct {
	GsmHeader
	*nasMessage.PDUSessionEstablishmentRequest      // 8.3.1
	*nasMessage.PDUSessionEstablishmentAccept       // 8.3.2
	*nasMessage.PDUSessionEstablishmentReject       // 8.3.3
	*nasMessage.PDUSessionAuthenticationCommand     // 8.3.4
	*nasMessage.PDUSessionAuthenticationComplete    // 8.3.5
	*nasMessage.PDUSessionAuthenticationResult      // 8.3.6
	*nasMessage.PDUSessionModificationRequest       // 8.3.7
	*nasMessage.PDUSessionModificationReject        // 8.3.8
	*nasMessage.PDUSessionModificationCommand       // 8.3.9
	*nasMessage.PDUSessionModificationComplete      // 8.3.10
	*nasMessage.PDUSessionModificationCommandReject // 8.3.11
	*nasMessage.PDUSessionReleaseRequest            // 8.3.12
	*nasMessage.PDUSessionReleaseReject             // 8.3.13
	*nasMessage.PDUSessionReleaseCommand            // 8.3.14
	*nasMessage.PDUSessionReleaseComplete           // 8.3.15
	*nasMessage.Status5GSM                          // 8.3.16
	*nasMessage.ServiceLevelAuthenticationCommand   // 8.3.17
	*nasMessage.ServiceLevelAuthenticationComplete  // 8.3.18
	*nasMessage.RemoteUEReport                      // 8.3.19
	*nasMessage.RemoteUEReportResponse              // 8.3.20
}

const (
	MsgTypePDUSessionEstablishmentRequest      uint8 = 193
	MsgTypePDUSessionEstablishmentAccept       uint8 = 194
	MsgTypePDUSessionEstablishmentReject       uint8 = 195
	MsgTypePDUSessionAuthenticationCommand     uint8 = 197
	MsgTypePDUSessionAuthenticationComplete    uint8 = 198
	MsgTypePDUSessionAuthenticationResult      uint8 = 199
	MsgTypePDUSessionModificationRequest       uint8 = 201
	MsgTypePDUSessionModificationReject        uint8 = 202
	MsgTypePDUSessionModificationCommand       uint8 = 203
	MsgTypePDUSessionModificationComplete      uint8 = 204
	MsgTypePDUSessionModificationCommandReject uint8 = 205
	MsgTypePDUSessionReleaseRequest            uint8 = 209
	MsgTypePDUSessionReleaseReject             uint8 = 210
	MsgTypePDUSessionReleaseCommand            uint8 = 211
	MsgTypePDUSessionReleaseComplete           uint8 = 212
	MsgTypeStatus5GSM                          uint8 = 214
	MsgTypeServiceLevelAuthenticationCommand   uint8 = 216
	MsgTypeServiceLevelAuthenticationComplete  uint8 = 217
	MsgTypeRemoteUEReport                      uint8 = 218
	MsgTypeRemoteUEReportResponse              uint8 = 219
)

func (a *Message) GsmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GsmMessage = NewGsmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GsmHeader); err != nil {
		return err
	}
	switch a.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		a.PDUSessionEstablishmentRequest = nasMessage.NewPDUSessionEstablishmentRequest(MsgTypePDUSessionEstablishmentRequest)
		a.DecodePDUSessionEstablishmentRequest(byteArray)
	case MsgTypePDUSessionEstablishmentAccept:
		a.PDUSessionEstablishmentAccept = nasMessage.NewPDUSessionEstablishmentAccept(MsgTypePDUSessionEstablishmentAccept)
		a.DecodePDUSessionEstablishmentAccept(byteArray)
	case MsgTypePDUSessionEstablishmentReject:
		a.PDUSessionEstablishmentReject = nasMessage.NewPDUSessionEstablishmentReject(MsgTypePDUSessionEstablishmentReject)
		a.DecodePDUSessionEstablishmentReject(byteArray)
	case MsgTypePDUSessionAuthenticationCommand:
		a.PDUSessionAuthenticationCommand = nasMessage.NewPDUSessionAuthenticationCommand(MsgTypePDUSessionAuthenticationCommand)
		a.DecodePDUSessionAuthenticationCommand(byteArray)
	case MsgTypePDUSessionAuthenticationComplete:
		a.PDUSessionAuthenticationComplete = nasMessage.NewPDUSessionAuthenticationComplete(MsgTypePDUSessionAuthenticationComplete)
		a.DecodePDUSessionAuthenticationComplete(byteArray)
	case MsgTypePDUSessionAuthenticationResult:
		a.PDUSessionAuthenticationResult = nasMessage.NewPDUSessionAuthenticationResult(MsgTypePDUSessionAuthenticationResult)
		a.DecodePDUSessionAuthenticationResult(byteArray)
	case MsgTypePDUSessionModificationRequest:
		a.PDUSessionModificationRequest = nasMessage.NewPDUSessionModificationRequest(MsgTypePDUSessionModificationRequest)
		a.DecodePDUSessionModificationRequest(byteArray)
	case MsgTypePDUSessionModificationReject:
		a.PDUSessionModificationReject = nasMessage.NewPDUSessionModificationReject(MsgTypePDUSessionModificationReject)
		a.DecodePDUSessionModificationReject(byteArray)
	case MsgTypePDUSessionModificationCommand:
		a.PDUSessionModificationCommand = nasMessage.NewPDUSessionModificationCommand(MsgTypePDUSessionModificationCommand)
		a.DecodePDUSessionModificationCommand(byteArray)
	case MsgTypePDUSessionModificationComplete:
		a.PDUSessionModificationComplete = nasMessage.NewPDUSessionModificationComplete(MsgTypePDUSessionModificationComplete)
		a.DecodePDUSessionModificationComplete(byteArray)
	case MsgTypePDUSessionModificationCommandReject:
		a.PDUSessionModificationCommandReject = nasMessage.NewPDUSessionModificationCommandReject(MsgTypePDUSessionModificationCommandReject)
		a.DecodePDUSessionModificationCommandReject(byteArray)
	case MsgTypePDUSessionReleaseRequest:
		a.PDUSessionReleaseRequest = nasMessage.NewPDUSessionReleaseRequest(MsgTypePDUSessionReleaseRequest)
		a.DecodePDUSessionReleaseRequest(byteArray)
	case MsgTypePDUSessionReleaseReject:
		a.PDUSessionReleaseReject = nasMessage.NewPDUSessionReleaseReject(MsgTypePDUSessionReleaseReject)
		a.DecodePDUSessionReleaseReject(byteArray)
	case MsgTypePDUSessionReleaseCommand:
		a.PDUSessionReleaseCommand = nasMessage.NewPDUSessionReleaseCommand(MsgTypePDUSessionReleaseCommand)
		a.DecodePDUSessionReleaseCommand(byteArray)
	case MsgTypePDUSessionReleaseComplete:
		a.PDUSessionReleaseComplete = nasMessage.NewPDUSessionReleaseComplete(MsgTypePDUSessionReleaseComplete)
		a.DecodePDUSessionReleaseComplete(byteArray)
	case MsgTypeStatus5GSM:
		a.Status5GSM = nasMessage.NewStatus5GSM(MsgTypeStatus5GSM)
		a.DecodeStatus5GSM(byteArray)
	case MsgTypeServiceLevelAuthenticationCommand:
		a.ServiceLevelAuthenticationCommand = nasMessage.NewServiceLevelAuthenticationCommand(MsgTypeServiceLevelAuthenticationCommand)
		a.DecodeServiceLevelAuthenticationCommand(byteArray)
	case MsgTypeServiceLevelAuthenticationComplete:
		a.ServiceLevelAuthenticationComplete = nasMessage.NewServiceLevelAuthenticationComplete(MsgTypeServiceLevelAuthenticationComplete)
		a.DecodeServiceLevelAuthenticationComplete(byteArray)
	case MsgTypeRemoteUEReport:
		a.RemoteUEReport = nasMessage.NewRemoteUEReport(MsgTypeRemoteUEReport)
		a.DecodeRemoteUEReport(byteArray)
	case MsgTypeRemoteUEReportResponse:
		a.RemoteUEReportResponse = nasMessage.NewRemoteUEReportResponse(MsgTypeRemoteUEReportResponse)
		a.DecodeRemoteUEReportResponse(byteArray)
	default:
		return fmt.Errorf("NAS decode fail: MsgType[%d] does not exist in GSM Message", a.GsmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GsmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		a.EncodePDUSessionEstablishmentRequest(buffer)
	case MsgTypePDUSessionEstablishmentAccept:
		a.EncodePDUSessionEstablishmentAccept(buffer)
	case MsgTypePDUSessionEstablishmentReject:
		a.EncodePDUSessionEstablishmentReject(buffer)
	case MsgTypePDUSessionAuthenticationCommand:
		a.EncodePDUSessionAuthenticationCommand(buffer)
	case MsgTypePDUSessionAuthenticationComplete:
		a.EncodePDUSessionAuthenticationComplete(buffer)
	case MsgTypePDUSessionAuthenticationResult:
		a.EncodePDUSessionAuthenticationResult(buffer)
	case MsgTypePDUSessionModificationRequest:
		a.EncodePDUSessionModificationRequest(buffer)
	case MsgTypePDUSessionModificationReject:
		a.EncodePDUSessionModificationReject(buffer)
	case MsgTypePDUSessionModificationCommand:
		a.EncodePDUSessionModificationCommand(buffer)
	case MsgTypePDUSessionModificationComplete:
		a.EncodePDUSessionModificationComplete(buffer)
	case MsgTypePDUSessionModificationCommandReject:
		a.EncodePDUSessionModificationCommandReject(buffer)
	case MsgTypePDUSessionReleaseRequest:
		a.EncodePDUSessionReleaseRequest(buffer)
	case MsgTypePDUSessionReleaseReject:
		a.EncodePDUSessionReleaseReject(buffer)
	case MsgTypePDUSessionReleaseCommand:
		a.EncodePDUSessionReleaseCommand(buffer)
	case MsgTypePDUSessionReleaseComplete:
		a.EncodePDUSessionReleaseComplete(buffer)
	case MsgTypeStatus5GSM:
		a.EncodeStatus5GSM(buffer)
	case MsgTypeServiceLevelAuthenticationCommand:
		a.EncodeServiceLevelAuthenticationCommand(buffer)
	case MsgTypeServiceLevelAuthenticationComplete:
		a.EncodeServiceLevelAuthenticationComplete(buffer)
	case MsgTypeRemoteUEReport:
		a.EncodeRemoteUEReport(buffer)
	case MsgTypeRemoteUEReportResponse:
		a.EncodeRemoteUEReportResponse(buffer)
	default:
		return fmt.Errorf("NAS encode fail: MsgType[%d] does not exist in GSM Message", a.GsmHeader.GetMessageType())
	}
	return nil
}
