// SPDX-FileCopyrightText: 2025 Intel Corporation
// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasTestpacket

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasConvert"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/omec-project/openapi/v2/models"
)

const (
	PDUSesModiReq    string = "PDU Session Modification Request"
	PDUSesModiCmp    string = "PDU Session Modification Complete"
	PDUSesModiCmdRej string = "PDU Session Modification Command Reject"
	PDUSesRelReq     string = "PDU Session Release Request"
	PDUSesRelCmp     string = "PDU Session Release Complete"
	PDUSesRelRej     string = "PDU Session Release Reject"
	PDUSesAuthCmp    string = "PDU Session Authentication Complete"
)

func GetRegistrationRequest(
	registrationType uint8,
	mobileIdentity nasType.MobileIdentity5GS,
	requestedNSSAI *nasType.RequestedNSSAI,
	ueSecurityCapability *nasType.UESecurityCapability,
	capability5GMM *nasType.Capability5GMM,
	nasMessageContainer []uint8,
	uplinkDataStatus *nasType.UplinkDataStatus,
) []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeRegistrationRequest)

	registrationRequest := nasMessage.NewRegistrationRequest(0)
	registrationRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	registrationRequest.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	registrationRequest.SetSpareHalfOctet(0x00)
	registrationRequest.SetMessageType(nas.MsgTypeRegistrationRequest)
	registrationRequest.SetTSC(nasMessage.TypeOfSecurityContextFlagNative)
	registrationRequest.NgksiAndRegistrationType5GS.SetNasKeySetIdentifiler(0x7)
	registrationRequest.SetFOR(1)
	registrationRequest.SetRegistrationType5GS(registrationType)
	registrationRequest.MobileIdentity5GS = mobileIdentity

	registrationRequest.UESecurityCapability = ueSecurityCapability
	registrationRequest.Capability5GMM = capability5GMM
	registrationRequest.RequestedNSSAI = requestedNSSAI
	registrationRequest.UplinkDataStatus = uplinkDataStatus

	if nasMessageContainer != nil {
		registrationRequest.NASMessageContainer = nasType.NewNASMessageContainer(
			nasMessage.RegistrationRequestNASMessageContainerType)
		registrationRequest.NASMessageContainer.SetLen(uint16(len(nasMessageContainer)))
		registrationRequest.SetNASMessageContainerContents(nasMessageContainer)
	}

	m.RegistrationRequest = registrationRequest

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetPduSessionEstablishmentRequest(pduSessionId uint8) []byte {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionEstablishmentRequest)

	pduSessionEstablishmentRequest := nasMessage.NewPDUSessionEstablishmentRequest(0)
	pduSessionEstablishmentRequest.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSSessionManagementMessage)
	pduSessionEstablishmentRequest.SetMessageType(nas.MsgTypePDUSessionEstablishmentRequest)
	pduSessionEstablishmentRequest.SetPDUSessionID(pduSessionId)
	pduSessionEstablishmentRequest.SetPTI(0x01)
	pduSessionEstablishmentRequest.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(0xff)
	pduSessionEstablishmentRequest.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(0xff)

	pduSessionEstablishmentRequest.PDUSessionType = nasType.NewPDUSessionType(nasMessage.PDUSessionEstablishmentRequestPDUSessionTypeType)
	pduSessionEstablishmentRequest.SetPDUSessionTypeValue(uint8(0x01)) // IPv4 type

	pduSessionEstablishmentRequest.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(
		nasMessage.PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType)
	protocolConfigurationOptions := nasConvert.NewProtocolConfigurationOptions()
	protocolConfigurationOptions.AddIPAddressAllocationViaNASSignallingUL()
	protocolConfigurationOptions.AddDNSServerIPv4AddressRequest()
	protocolConfigurationOptions.AddDNSServerIPv6AddressRequest()
	pcoContents := protocolConfigurationOptions.Marshal()
	pcoContentsLength := len(pcoContents)
	pduSessionEstablishmentRequest.ExtendedProtocolConfigurationOptions.SetLen(uint16(pcoContentsLength))
	pduSessionEstablishmentRequest.SetExtendedProtocolConfigurationOptionsContents(pcoContents)

	m.PDUSessionEstablishmentRequest = pduSessionEstablishmentRequest

	data := new(bytes.Buffer)
	err := m.GsmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetUlNasTransport_PduSessionEstablishmentRequest(pduSessionId uint8, requestType uint8, dnnString string,
	sNssai *models.Snssai,
) []byte {
	pduSessionEstablishmentRequest := GetPduSessionEstablishmentRequest(pduSessionId)

	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeULNASTransport)

	ulNasTransport := nasMessage.NewULNASTransport(0)
	ulNasTransport.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	ulNasTransport.SetMessageType(nas.MsgTypeULNASTransport)
	ulNasTransport.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	ulNasTransport.PduSessionID2Value = new(nasType.PduSessionID2Value)
	ulNasTransport.PduSessionID2Value.SetIei(nasMessage.ULNASTransportPduSessionID2ValueType)
	ulNasTransport.SetPduSessionID2Value(pduSessionId)
	ulNasTransport.RequestType = new(nasType.RequestType)
	ulNasTransport.RequestType.SetIei(nasMessage.ULNASTransportRequestTypeType)
	ulNasTransport.SetRequestTypeValue(requestType)
	if dnnString != "" {
		dnn := []byte(dnnString)
		ulNasTransport.DNN = new(nasType.DNN)
		ulNasTransport.DNN.SetIei(nasMessage.ULNASTransportDNNType)
		ulNasTransport.DNN.SetLen(uint8(len(dnn)))
		ulNasTransport.SetDNN(dnn)
	}
	if sNssai != nil {
		var sdTemp [3]uint8
		sd, err := hex.DecodeString(sNssai.GetSd())
		if err != nil {
			logger.NasMsgLog.Errorf("sNssai decode error: %+v", err)
		}
		copy(sdTemp[:], sd)
		ulNasTransport.SNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportSNSSAIType)
		ulNasTransport.SNSSAI.SetLen(4)
		ulNasTransport.SetSST(uint8(sNssai.Sst))
		ulNasTransport.SetSD(sdTemp)
	}

	ulNasTransport.SetPayloadContainerType(nasMessage.PayloadContainerTypeN1SMInfo)
	ulNasTransport.PayloadContainer.SetLen(uint16(len(pduSessionEstablishmentRequest)))
	ulNasTransport.SetPayloadContainerContents(pduSessionEstablishmentRequest)

	m.ULNASTransport = ulNasTransport

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetPduSessionReleaseRequest(pduSessionId uint8) []byte {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionReleaseRequest)

	pduSessionReleaseRequest := nasMessage.NewPDUSessionReleaseRequest(0)
	pduSessionReleaseRequest.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSSessionManagementMessage)
	pduSessionReleaseRequest.SetMessageType(nas.MsgTypePDUSessionReleaseRequest)
	pduSessionReleaseRequest.SetPDUSessionID(pduSessionId)
	pduSessionReleaseRequest.SetPTI(0x01)

	m.PDUSessionReleaseRequest = pduSessionReleaseRequest

	data := new(bytes.Buffer)
	err := m.GsmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetPduSessionReleaseComplete(pduSessionId uint8) []byte {
	m := nas.NewMessage()
	m.GsmMessage = nas.NewGsmMessage()
	m.GsmHeader.SetMessageType(nas.MsgTypePDUSessionReleaseComplete)

	pduSessionReleaseComplete := nasMessage.NewPDUSessionReleaseComplete(0)
	pduSessionReleaseComplete.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSSessionManagementMessage)
	pduSessionReleaseComplete.SetMessageType(nas.MsgTypePDUSessionReleaseComplete)
	pduSessionReleaseComplete.SetPDUSessionID(pduSessionId)
	pduSessionReleaseComplete.SetPTI(0x01)

	m.PDUSessionReleaseComplete = pduSessionReleaseComplete

	data := new(bytes.Buffer)
	err := m.GsmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetAuthenticationResponse(authenticationResponseParam []uint8, eapMsg string) []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeAuthenticationResponse)

	authenticationResponse := nasMessage.NewAuthenticationResponse(0)
	authenticationResponse.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	authenticationResponse.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	authenticationResponse.SetSpareHalfOctet(0)
	authenticationResponse.SetMessageType(nas.MsgTypeAuthenticationResponse)

	if len(authenticationResponseParam) > 0 {
		authenticationResponse.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(
			nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
		authenticationResponse.AuthenticationResponseParameter.SetLen(uint8(len(authenticationResponseParam)))
		copy(authenticationResponse.AuthenticationResponseParameter.Octet[:], authenticationResponseParam[0:16])
	} else if eapMsg != "" {
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
		if err != nil {
			logger.NasMsgLog.Warnf("EAP decode error: %+v", err)
		}
		authenticationResponse.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationResponseEAPMessageType)
		authenticationResponse.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		authenticationResponse.SetEAPMessage(rawEapMsg)
	}

	m.AuthenticationResponse = authenticationResponse

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetRegistrationComplete(sorTransparentContainer []uint8) []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeRegistrationComplete)

	registrationComplete := nasMessage.NewRegistrationComplete(0)
	registrationComplete.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	registrationComplete.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	registrationComplete.SetSpareHalfOctet(0)
	registrationComplete.SetMessageType(nas.MsgTypeRegistrationComplete)

	if sorTransparentContainer != nil {
		registrationComplete.SORTransparentContainer = nasType.NewSORTransparentContainer(
			nasMessage.RegistrationCompleteSORTransparentContainerType)
		registrationComplete.SetLen(uint16(len(sorTransparentContainer)))
		registrationComplete.SetSORContent(sorTransparentContainer)
	}

	m.RegistrationComplete = registrationComplete

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

// TS 24.501 8.2.26.
func GetSecurityModeComplete(nasMessageContainer []uint8) []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeSecurityModeComplete)

	securityModeComplete := nasMessage.NewSecurityModeComplete(0)
	securityModeComplete.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	// TODO: modify security header type if need security protected
	securityModeComplete.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	securityModeComplete.SetSpareHalfOctet(0)
	securityModeComplete.SetMessageType(nas.MsgTypeSecurityModeComplete)

	securityModeComplete.IMEISV = nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	securityModeComplete.IMEISV.SetLen(9)
	securityModeComplete.SetOddEvenIdic(0)
	securityModeComplete.SetTypeOfIdentity(nasMessage.MobileIdentity5GSTypeImeisv)
	securityModeComplete.SetIdentityDigit1(1)
	securityModeComplete.SetIdentityDigitP_1(1)
	securityModeComplete.SetIdentityDigitP(1)
	securityModeComplete.SetIdentityDigitP_15(15)

	if nasMessageContainer != nil {
		securityModeComplete.NASMessageContainer = nasType.NewNASMessageContainer(
			nasMessage.SecurityModeCompleteNASMessageContainerType)
		securityModeComplete.NASMessageContainer.SetLen(uint16(len(nasMessageContainer)))
		securityModeComplete.SetNASMessageContainerContents(nasMessageContainer)
	}

	m.SecurityModeComplete = securityModeComplete

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetDeregistrationRequest(accessType uint8, switchOff uint8, ngKsi uint8,
	mobileIdentity5GS nasType.MobileIdentity5GS,
) []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeDeregistrationRequestUEOriginatingDeregistration)

	deregistrationRequest := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
	deregistrationRequest.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationRequest.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	deregistrationRequest.SetSpareHalfOctet(0)
	deregistrationRequest.SetMessageType(
		nas.MsgTypeDeregistrationRequestUEOriginatingDeregistration)

	deregistrationRequest.SetAccessType(accessType)
	deregistrationRequest.SetSwitchOff(switchOff)
	deregistrationRequest.SetReRegistrationRequired(0)
	deregistrationRequest.SetTSC(ngKsi)
	deregistrationRequest.SetNasKeySetIdentifiler(ngKsi)
	deregistrationRequest.MobileIdentity5GS.SetLen(mobileIdentity5GS.GetLen())
	deregistrationRequest.SetMobileIdentity5GSContents(mobileIdentity5GS.GetMobileIdentity5GSContents())

	m.DeregistrationRequestUEOriginatingDeregistration = deregistrationRequest

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetDeregistrationAccept() []byte {
	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration)

	deregistrationAccept := nasMessage.NewDeregistrationAcceptUETerminatedDeregistration(0)
	deregistrationAccept.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationAccept.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	deregistrationAccept.SetSpareHalfOctet(0)
	deregistrationAccept.SetMessageType(
		nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration)

	m.DeregistrationAcceptUETerminatedDeregistration = deregistrationAccept

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetUlNasTransport_PduSessionReleaseRequest(pduSessionId uint8) []byte {
	pduSessionReleaseRequest := GetPduSessionReleaseRequest(pduSessionId)

	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeULNASTransport)

	ulNasTransport := nasMessage.NewULNASTransport(0)
	ulNasTransport.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	ulNasTransport.SetMessageType(nas.MsgTypeULNASTransport)
	ulNasTransport.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	ulNasTransport.PduSessionID2Value = new(nasType.PduSessionID2Value)
	ulNasTransport.PduSessionID2Value.SetIei(nasMessage.ULNASTransportPduSessionID2ValueType)
	ulNasTransport.SetPduSessionID2Value(pduSessionId)

	ulNasTransport.SetPayloadContainerType(nasMessage.PayloadContainerTypeN1SMInfo)
	ulNasTransport.PayloadContainer.SetLen(uint16(len(pduSessionReleaseRequest)))
	ulNasTransport.SetPayloadContainerContents(pduSessionReleaseRequest)

	m.ULNASTransport = ulNasTransport

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}

func GetUlNasTransport_PduSessionReleaseComplete(pduSessionId uint8, requestType uint8, dnnString string,
	sNssai *models.Snssai,
) []byte {
	pduSessionReleaseRequest := GetPduSessionReleaseComplete(pduSessionId)

	m := nas.NewMessage()
	m.GmmMessage = nas.NewGmmMessage()
	m.GmmHeader.SetMessageType(nas.MsgTypeULNASTransport)

	ulNasTransport := nasMessage.NewULNASTransport(0)
	ulNasTransport.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	ulNasTransport.SetMessageType(nas.MsgTypeULNASTransport)
	ulNasTransport.SetExtendedProtocolDiscriminator(
		nasMessage.Epd5GSMobilityManagementMessage)
	ulNasTransport.PduSessionID2Value = new(nasType.PduSessionID2Value)
	ulNasTransport.PduSessionID2Value.SetIei(nasMessage.ULNASTransportPduSessionID2ValueType)
	ulNasTransport.SetPduSessionID2Value(pduSessionId)
	ulNasTransport.RequestType = new(nasType.RequestType)
	ulNasTransport.RequestType.SetIei(nasMessage.ULNASTransportRequestTypeType)
	ulNasTransport.SetRequestTypeValue(requestType)
	if dnnString != "" {
		dnn := []byte(dnnString)
		ulNasTransport.DNN = new(nasType.DNN)
		ulNasTransport.DNN.SetIei(nasMessage.ULNASTransportDNNType)
		ulNasTransport.DNN.SetLen(uint8(len(dnn)))
		ulNasTransport.SetDNN(dnn)
	}
	if sNssai != nil {
		var sdTemp [3]uint8
		sd, err := hex.DecodeString(sNssai.GetSd())
		if err != nil {
			logger.NasMsgLog.Warnf("sNssai SD decode error: %+v", err)
		}
		copy(sdTemp[:], sd)
		ulNasTransport.SNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportSNSSAIType)
		ulNasTransport.SNSSAI.SetLen(4)
		ulNasTransport.SetSST(uint8(sNssai.Sst))
		ulNasTransport.SetSD(sdTemp)
	}

	ulNasTransport.SetPayloadContainerType(nasMessage.PayloadContainerTypeN1SMInfo)
	ulNasTransport.PayloadContainer.SetLen(uint16(len(pduSessionReleaseRequest)))
	ulNasTransport.SetPayloadContainerContents(pduSessionReleaseRequest)

	m.ULNASTransport = ulNasTransport

	data := new(bytes.Buffer)
	err := m.GmmMessageEncode(data)
	if err != nil {
		logger.NasMsgLog.Errorln(err.Error())
	}

	return data.Bytes()
}
