// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import "fmt"

const (
	ULNASTransportRequestTypeInitialRequest              uint8 = 1
	ULNASTransportRequestTypeExistingPduSession          uint8 = 2
	ULNASTransportRequestTypeInitialEmergencyRequest     uint8 = 3
	ULNASTransportRequestTypeExistingEmergencyPduSession uint8 = 4
	ULNASTransportRequestTypeModificationRequest         uint8 = 5
	ULNASTransportRequestTypeReserved                    uint8 = 7
)

const (
	PayloadContainerTypeN1SMInfo          uint8 = 0x01
	PayloadContainerTypeSMS               uint8 = 0x02
	PayloadContainerTypeLPP               uint8 = 0x03
	PayloadContainerTypeSOR               uint8 = 0x04
	PayloadContainerTypeUEPolicy          uint8 = 0x05
	PayloadContainerTypeUEParameterUpdate uint8 = 0x06
	PayloadContainerTypeMultiplePayload   uint8 = 0x0f
)

const (
	Cause5GSMOperatorDeterminedBarring                                   uint8 = 0x08
	Cause5GSMInsufficientResources                                       uint8 = 0x1a
	Cause5GSMMissingOrUnknownDNN                                         uint8 = 0x1b
	Cause5GSMUnknownPDUSessionType                                       uint8 = 0x1c
	Cause5GSMUserAuthenticationOrAuthorizationFailed                     uint8 = 0x1d
	Cause5GSMRequestRejectedUnspecified                                  uint8 = 0x1f
	Cause5GSMServiceOptionNotSupported                                   uint8 = 0x20
	Cause5GSMRequestedServiceOptionNotSubscribed                         uint8 = 0x21
	Cause5GSMServiceOptionTemporarilyOutOfOrder                          uint8 = 0x22
	Cause5GSMPTIAlreadyInUse                                             uint8 = 0x23
	Cause5GSMRegularDeactivation                                         uint8 = 0x24
	Cause5GSM5GSQoSNotAccepted                                           uint8 = 0x25
	Cause5GSMNetworkFailure                                              uint8 = 0x26
	Cause5GSMReactivationRequested                                       uint8 = 0x27
	Cause5GSMSemanticErrorInTheTFTOperation                              uint8 = 0x29
	Cause5GSMSyntacticalErrorInTheTFTOperation                           uint8 = 0x2a
	Cause5GSMInvalidPDUSessionIdentity                                   uint8 = 0x2b
	Cause5GSMSemanticErrorsInPacketFilter                                uint8 = 0x2c
	Cause5GSMSyntacticalErrorInPacketFilter                              uint8 = 0x2d
	Cause5GSMOutOfLADNServiceArea                                        uint8 = 0x2e
	Cause5GSMPTIMismatch                                                 uint8 = 0x2f
	Cause5GSMPDUSessionTypeIPv4OnlyAllowed                               uint8 = 0x32
	Cause5GSMPDUSessionTypeIPv6OnlyAllowed                               uint8 = 0x33
	Cause5GSMPDUSessionDoesNotExist                                      uint8 = 0x36
	Cause5GSMPDUSessionTypeIPv4v6OnlyAllowed                             uint8 = 0x39
	Cause5GSMPDUSessionTypeUnstructuredOnlyAllowed                       uint8 = 0x3a
	Cause5GSMUnsupported5QIValue                                         uint8 = 0x3b
	Cause5GSMPDUSessionTypeEthernetOnlyAllowed                           uint8 = 0x3d
	Cause5GSMInsufficientResourcesForSpecificSliceAndDNN                 uint8 = 0x43
	Cause5GSMNotSupportedSSCMode                                         uint8 = 0x44
	Cause5GSMInsufficientResourcesForSpecificSlice                       uint8 = 0x45
	Cause5GSMMissingOrUnknownDNNInASlice                                 uint8 = 0x46
	Cause5GSMInvalidPTIValue                                             uint8 = 0x51
	Cause5GSMMaximumDataRatePerUEForUserPlaneIntegrityProtectionIsTooLow uint8 = 0x52
	Cause5GSMSemanticErrorInTheQoSOperation                              uint8 = 0x53
	Cause5GSMSyntacticalErrorInTheQoSOperation                           uint8 = 0x54
	Cause5GSMInvalidMappedEPSBearerIdentity                              uint8 = 0x55
	Cause5GSMUASServicesNotAllowed                                       uint8 = 0x56
	Cause5GSMSemanticallyIncorrectMessage                                uint8 = 0x5f
	Cause5GSMInvalidMandatoryInformation                                 uint8 = 0x60
	Cause5GSMMessageTypeNonExistentOrNotImplemented                      uint8 = 0x61
	Cause5GSMMessageTypeNotCompatibleWithTheProtocolState                uint8 = 0x62
	Cause5GSMInformationElementNonExistentOrNotImplemented               uint8 = 0x63
	Cause5GSMConditionalIEError                                          uint8 = 0x64
	Cause5GSMMessageNotCompatibleWithTheProtocolState                    uint8 = 0x65
	Cause5GSMProtocolErrorUnspecified                                    uint8 = 0x6f
)

const (
	Cause5GMMIllegalUE                                      uint8 = 0x03
	Cause5GMMPEINotAccepted                                 uint8 = 0x05
	Cause5GMMIllegalME                                      uint8 = 0x06
	Cause5GMM5GSServicesNotAllowed                          uint8 = 0x07
	Cause5GMMUEIdentityCannotBeDerivedByTheNetwork          uint8 = 0x09
	Cause5GMMImplicitlyDeregistered                         uint8 = 0x0a
	Cause5GMMPLMNNotAllowed                                 uint8 = 0x0b
	Cause5GMMTrackingAreaNotAllowed                         uint8 = 0x0c
	Cause5GMMRoamingNotAllowedInThisTrackingArea            uint8 = 0x0d
	Cause5GMMNoSuitableCellsInTrackingArea                  uint8 = 0x0f
	Cause5GMMMACFailure                                     uint8 = 0x14
	Cause5GMMSynchFailure                                   uint8 = 0x15
	Cause5GMMCongestion                                     uint8 = 0x16
	Cause5GMMUESecurityCapabilitiesMismatch                 uint8 = 0x17
	Cause5GMMSecurityModeRejectedUnspecified                uint8 = 0x18
	Cause5GMMNon5GAuthenticationUnacceptable                uint8 = 0x1a
	Cause5GMMN1ModeNotAllowed                               uint8 = 0x1b
	Cause5GMMRestrictedServiceArea                          uint8 = 0x1c
	Cause5GMMLADNNotAvailable                               uint8 = 0x2b
	Cause5GMMMaximumNumberOfPDUSessionsReached              uint8 = 0x41
	Cause5GMMInsufficientResourcesForSpecificSliceAndDNN    uint8 = 0x43
	Cause5GMMInsufficientResourcesForSpecificSlice          uint8 = 0x45
	Cause5GMMngKSIAlreadyInUse                              uint8 = 0x47
	Cause5GMMNon3GPPAccessTo5GCNNotAllowed                  uint8 = 0x48
	Cause5GMMServingNetworkNotAuthorized                    uint8 = 0x49
	Cause5GMMPayloadWasNotForwarded                         uint8 = 0x5a
	Cause5GMMDNNNotSupportedOrNotSubscribedInTheSlice       uint8 = 0x5b
	Cause5GMMInsufficientUserPlaneResourcesForThePDUSession uint8 = 0x5c
	Cause5GMMSemanticallyIncorrectMessage                   uint8 = 0x5f
	Cause5GMMInvalidMandatoryInformation                    uint8 = 0x60
	Cause5GMMMessageTypeNonExistentOrNotImplemented         uint8 = 0x61
	Cause5GMMMessageTypeNotCompatibleWithTheProtocolState   uint8 = 0x62
	Cause5GMMInformationElementNonExistentOrNotImplemented  uint8 = 0x63
	Cause5GMMConditionalIEError                             uint8 = 0x64
	Cause5GMMMessageNotCompatibleWithTheProtocolState       uint8 = 0x65
	Cause5GMMProtocolErrorUnspecified                       uint8 = 0x6f
)

// TS 24.501 9.11.3.7
const (
	RegistrationType5GSInitialRegistration          uint8 = 0x01
	RegistrationType5GSMobilityRegistrationUpdating uint8 = 0x02
	RegistrationType5GSPeriodicRegistrationUpdating uint8 = 0x03
	RegistrationType5GSEmergencyRegistration        uint8 = 0x04
	RegistrationType5GSReserved                     uint8 = 0x07
)

// TS 24.501 9.11.3.7
const (
	FollowOnRequestNoPending uint8 = 0x00
	FollowOnRequestPending   uint8 = 0x01
)

const (
	MobileIdentity5GSTypeNoIdentity uint8 = 0x00
	MobileIdentity5GSTypeSuci       uint8 = 0x01
	MobileIdentity5GSType5gGuti     uint8 = 0x02
	MobileIdentity5GSTypeImei       uint8 = 0x03
	MobileIdentity5GSType5gSTmsi    uint8 = 0x04
	MobileIdentity5GSTypeImeisv     uint8 = 0x05
)

// TS 24.501 9.11.3.2A
const (
	DRXValueNotSpecified  uint8 = 0x00
	DRXcycleParameterT32  uint8 = 0x01
	DRXcycleParameterT64  uint8 = 0x02
	DRXcycleParameterT128 uint8 = 0x03
	DRXcycleParameterT256 uint8 = 0x04
)

// TS 24.501 9.11.3.32
const (
	TypeOfSecurityContextFlagNative uint8 = 0x00
	TypeOfSecurityContextFlagMapped uint8 = 0x01
)

// TS 24.501 9.11.3.32
const (
	NasKeySetIdentifierNoKeyIsAvailable int32 = 0x07
)

// TS 24.501 9.11.3.11
const (
	AccessType3GPP    uint8 = 0x01
	AccessTypeNon3GPP uint8 = 0x02
	AccessTypeBoth    uint8 = 0x03
)

// TS 24.501 9.11.3.50
const (
	ServiceTypeSignalling                uint8 = 0x00
	ServiceTypeData                      uint8 = 0x01
	ServiceTypeMobileTerminatedServices  uint8 = 0x02
	ServiceTypeEmergencyServices         uint8 = 0x03
	ServiceTypeEmergencyServicesFallback uint8 = 0x04
	ServiceTypeHighPriorityAccess        uint8 = 0x05
)

// TS 24.501 9.11.3.20
const (
	ReRegistrationNotRequired uint8 = 0x00
	ReRegistrationRequired    uint8 = 0x01
)

// TS 24.501 9.11.3.28 TS 24.008 10.5.5.10
const (
	IMEISVNotRequested uint8 = 0x00
	IMEISVRequested    uint8 = 0x01
)

// TS 24.501 9.11.3.6
const (
	RegistrationResult5GS3GPPAccess           uint8 = 0x01
	RegistrationResult5GSNon3GPPAccess        uint8 = 0x02
	RegistrationResult5GS3GPPandNon3GPPAccess uint8 = 0x03
)

// TS 24.501 9.11.3.6
const (
	SMSOverNasNotAllowed uint8 = 0x00
	SMSOverNasAllowed    uint8 = 0x01
)

// TS 24.501 9.11.3.46
const (
	SnssaiNotAvailableInCurrentPlmn             uint8 = 0x00
	SnssaiNotAvailableInCurrentRegistrationArea uint8 = 0x01
)

// TS 24.008 10.5.7.4a
const (
	GPRSTimer3UnitMultiplesOf10Minutes uint8 = 0x00
	GPRSTimer3UnitMultiplesOf1Hour     uint8 = 0x01
	GPRSTimer3UnitMultiplesOf10Hours   uint8 = 0x02
	GPRSTimer3UnitMultiplesOf2Seconds  uint8 = 0x03
	GPRSTimer3UnitMultiplesOf30Seconds uint8 = 0x04
	GPRSTimer3UnitMultiplesOf1Minute   uint8 = 0x05
)

// TS 24.501 9.11.3.9A
const (
	NGRanRadioCapabilityUpdateNotNeeded uint8 = 0x00
	NGRanRadioCapabilityUpdateNeeded    uint8 = 0x01
)

// TS 24.501 9.11.3.49
const (
	AllowedTypeAllowedArea    uint8 = 0x00
	AllowedTypeNonAllowedArea uint8 = 0x01
)

// TS 24.501 9.11.3.46
const (
	RejectedSnssaiCauseNotAvailableInCurrentPlmn             uint8 = 0x00
	RejectedSnssaiCauseNotAvailableInCurrentRegistrationArea uint8 = 0x01
)

// TS 24.501 9.11.4.10
const (
	PDUSessionTypeIPv4         uint8 = 0x01
	PDUSessionTypeIPv6         uint8 = 0x02
	PDUSessionTypeIPv4IPv6     uint8 = 0x03
	PDUSessionTypeUnstructured uint8 = 0x04
	PDUSessionTypeEthernet     uint8 = 0x05
)

// TS 24.501 9.11.3.4
const (
	SupiFormatImsi uint8 = 0x00
	SupiFormatNai  uint8 = 0x01
)

// TS 24.501 9.11.3.4
const (
	ProtectionSchemeNullScheme    int = 0
	ProtectionSchemeECIESProfileA int = 1
	ProtectionSchemeECIESProfileB int = 2
)

// TS 24.501 Table 9.11.4.14.1
const (
	SessionAMBRUnitNotUsed uint8 = 0x00
	SessionAMBRUnit1Kbps   uint8 = 0x01
	SessionAMBRUnit4Kbps   uint8 = 0x02
	SessionAMBRUnit16Kbps  uint8 = 0x03
	SessionAMBRUnit64Kbps  uint8 = 0x04
	SessionAMBRUnit256Kbps uint8 = 0x05
	SessionAMBRUnit1Mbps   uint8 = 0x06
	SessionAMBRUnit4Mbps   uint8 = 0x07
	SessionAMBRUnit16Mbps  uint8 = 0x08
	SessionAMBRUnit64Mbps  uint8 = 0x09
	SessionAMBRUnit256Mbps uint8 = 0x0A
	SessionAMBRUnit1Gbps   uint8 = 0x0B
	SessionAMBRUnit4Gbps   uint8 = 0x0C
	SessionAMBRUnit16Gbps  uint8 = 0x0D
	SessionAMBRUnit64Gbps  uint8 = 0x0E
	SessionAMBRUnit256Gbps uint8 = 0x0F
	SessionAMBRUnit1Tbps   uint8 = 0x10
	SessionAMBRUnit4Tbps   uint8 = 0x11
	SessionAMBRUnit16Tbps  uint8 = 0x12
	SessionAMBRUnit64Tbps  uint8 = 0x13
	SessionAMBRUnit256Tbps uint8 = 0x14
	SessionAMBRUnit1Pbps   uint8 = 0x15
	SessionAMBRUnit4Pbps   uint8 = 0x16
	SessionAMBRUnit16Pbps  uint8 = 0x17
	SessionAMBRUnit64Pbps  uint8 = 0x18
	SessionAMBRUnit256Pbps uint8 = 0x19
)

// TS 24.008 10.5.6.3
const (
	PCSCFIPv6AddressRequestUL                                     uint16 = 0x0001
	IMCNSubsystemSignalingFlagUL                                  uint16 = 0x0002
	DNSServerIPv6AddressRequestUL                                 uint16 = 0x0003
	NotSupportedUL                                                uint16 = 0x0004
	MSSupportOfNetworkRequestedBearerControlIndicatorUL           uint16 = 0x0005
	DSMIPv6HomeAgentAddressRequestUL                              uint16 = 0x0007
	DSMIPv6HomeNetworkPrefixRequestUL                             uint16 = 0x0008
	DSMIPv6IPv4HomeAgentAddressRequestUL                          uint16 = 0x0009
	IPAddressAllocationViaNASSignallingUL                         uint16 = 0x000a
	IPv4AddressAllocationViaDHCPv4UL                              uint16 = 0x000b
	PCSCFIPv4AddressRequestUL                                     uint16 = 0x000c
	DNSServerIPv4AddressRequestUL                                 uint16 = 0x000d
	MSISDNRequestUL                                               uint16 = 0x000e
	IFOMSupportRequestUL                                          uint16 = 0x000f
	IPv4LinkMTURequestUL                                          uint16 = 0x0010
	MSSupportOfLocalAddressInTFTIndicatorUL                       uint16 = 0x0011
	PCSCFReSelectionSupportUL                                     uint16 = 0x0012
	NBIFOMRequestIndicatorUL                                      uint16 = 0x0013
	NBIFOMModeUL                                                  uint16 = 0x0014
	NonIPLinkMTURequestUL                                         uint16 = 0x0015
	APNRateControlSupportIndicatorUL                              uint16 = 0x0016
	UEStatus3GPPPSDataOffUL                                       uint16 = 0x0017
	ReliableDataServiceRequestIndicatorUL                         uint16 = 0x0018
	AdditionalAPNRateControlForExceptionDataSupportIndicatorUL    uint16 = 0x0019
	PDUSessionIDUL                                                uint16 = 0x001a
	EthernetFramePayloadMTURequestUL                              uint16 = 0x0020
	UnstructuredLinkMTURequestUL                                  uint16 = 0x0021
	I5GSMCauseValueUL                                             uint16 = 0x0022 // 5GSMCauseValueUL
	QoSRulesWithTheLengthOfTwoOctetsSupportIndicatorUL            uint16 = 0x0023
	QoSFlowDescriptionsWithTheLengthOfTwoOctetsSupportIndicatorUL uint16 = 0x0024
	LinkControlProtocolUL                                         uint16 = 0xc021
	PushAccessControlProtocolUL                                   uint16 = 0xc023
	ChallengeHandshakeAuthenticationProtocolUL                    uint16 = 0xc223
	InternetProtocolControlProtocolUL                             uint16 = 0x8021
)

// TS 24.008 10.5.6.3
const (
	PCSCFIPv6AddressDL                                   uint16 = 0x0001
	IMCNSubsystemSignalingFlagDL                         uint16 = 0x0002
	DNSServerIPv6AddressDL                               uint16 = 0x0003
	PolicyControlRejectionCodeDL                         uint16 = 0x0004
	SelectedBearerControlModeDL                          uint16 = 0x0005
	DSMIPv6HomeAgentAddressDL                            uint16 = 0x0007
	DSMIPv6HomeNetworkPrefixDL                           uint16 = 0x0008
	DSMIPv6IPv4HomeAgentAddressDL                        uint16 = 0x0009
	PCSCFIPv4AddressDL                                   uint16 = 0x000c
	DNSServerIPv4AddressDL                               uint16 = 0x000d
	MSISDNDL                                             uint16 = 0x000e
	IFOMSupportDL                                        uint16 = 0x000f
	IPv4LinkMTUDL                                        uint16 = 0x0010
	NetworkSupportOfLocaladdressInTFTIndicatorDL         uint16 = 0x0011
	NBIFOMAcceptedIndicatorDL                            uint16 = 0x0013
	NBIFOMModeDL                                         uint16 = 0x0014
	NonIPLinkMTUDL                                       uint16 = 0x0015
	APNRateControlParametersDL                           uint16 = 0x0016
	Indication3GPPPSDataOffSupportDL                     uint16 = 0x0017
	ReliableDataServiceAcceptedIndicatorDL               uint16 = 0x0018
	AdditionalAPNRateControlForExceptionDataParametersDL uint16 = 0x0019
	SNSSAIDL                                             uint16 = 0x001b
	QoSRulesDL                                           uint16 = 0x001c
	SessionAMBRDL                                        uint16 = 0x001d
	PDUSessionAddressLifetimeDL                          uint16 = 0x001e
	QoSFlowDescriptions                                  uint16 = 0x001f
	EthernetFramePayloadMTU                              uint16 = 0x0020
	UnstructuredLinkMTU                                  uint16 = 0x0021
	QoSRulesWithTheLengthOfTwoOctets                     uint16 = 0x0023
	QoSFlowDescriptionsWithTheLengthOfTwoOctets          uint16 = 0x0024
)

func Cause5GMMToString(cause uint8) string {
	switch cause {
	case Cause5GMMIllegalUE:
		return fmt.Sprintf("Illegal UE (%d)", Cause5GMMIllegalUE)
	case Cause5GMMPEINotAccepted:
		return fmt.Sprintf("PEI not accepted (%d)", Cause5GMMPEINotAccepted)
	case Cause5GMMIllegalME:
		return fmt.Sprintf("Illegal ME (%d)", Cause5GMMIllegalME)
	case Cause5GMM5GSServicesNotAllowed:
		return fmt.Sprintf("5GS services not allowed (%d)", Cause5GMM5GSServicesNotAllowed)
	case Cause5GMMUEIdentityCannotBeDerivedByTheNetwork:
		return fmt.Sprintf("UE identity cannot be derived by the network (%d)",
			Cause5GMMUEIdentityCannotBeDerivedByTheNetwork)
	case Cause5GMMImplicitlyDeregistered:
		return fmt.Sprintf("Implicitly deregistered (%d)", Cause5GMMImplicitlyDeregistered)
	case Cause5GMMPLMNNotAllowed:
		return fmt.Sprintf("PLMN not allowed (%d)", Cause5GMMPLMNNotAllowed)
	case Cause5GMMTrackingAreaNotAllowed:
		return fmt.Sprintf("Tracking area not allowed (%d)", Cause5GMMTrackingAreaNotAllowed)
	case Cause5GMMRoamingNotAllowedInThisTrackingArea:
		return fmt.Sprintf("Roaming not allowed in this tracking area (%d)", Cause5GMMRoamingNotAllowedInThisTrackingArea)
	case Cause5GMMNoSuitableCellsInTrackingArea:
		return fmt.Sprintf("No suitable cells in tracking area (%d)", Cause5GMMNoSuitableCellsInTrackingArea)
	case Cause5GMMMACFailure:
		return fmt.Sprintf("MAC failure (%d)", Cause5GMMMACFailure)
	case Cause5GMMSynchFailure:
		return fmt.Sprintf("Synch failure (%d)", Cause5GMMSynchFailure)
	case Cause5GMMCongestion:
		return fmt.Sprintf("Congestion (%d)", Cause5GMMCongestion)
	case Cause5GMMUESecurityCapabilitiesMismatch:
		return fmt.Sprintf("UE security capabilities mismatch (%d)", Cause5GMMUESecurityCapabilitiesMismatch)
	case Cause5GMMSecurityModeRejectedUnspecified:
		return fmt.Sprintf("Security mode rejected, upspecified (%d)", Cause5GMMSecurityModeRejectedUnspecified)
	case Cause5GMMNon5GAuthenticationUnacceptable:
		return fmt.Sprintf("Non 5G authentication unacceptable (%d)", Cause5GMMNon5GAuthenticationUnacceptable)
	case Cause5GMMN1ModeNotAllowed:
		return fmt.Sprintf("N1 mode not allowed (%d)", Cause5GMMN1ModeNotAllowed)
	case Cause5GMMRestrictedServiceArea:
		return fmt.Sprintf("Restricted service area (%d)", Cause5GMMRestrictedServiceArea)
	case Cause5GMMLADNNotAvailable:
		return fmt.Sprintf("LADN not available (%d)", Cause5GMMLADNNotAvailable)
	case Cause5GMMMaximumNumberOfPDUSessionsReached:
		return fmt.Sprintf("Maximum number of PDU sessions reached (%d)", Cause5GMMMaximumNumberOfPDUSessionsReached)
	case Cause5GMMInsufficientResourcesForSpecificSliceAndDNN:
		return fmt.Sprintf("Insufficient resources for specific slice and DNN (%d)",
			Cause5GMMInsufficientResourcesForSpecificSliceAndDNN)
	case Cause5GMMInsufficientResourcesForSpecificSlice:
		return fmt.Sprintf("Insufficient resources for specific slice (%d)", Cause5GMMInsufficientResourcesForSpecificSlice)
	case Cause5GMMngKSIAlreadyInUse:
		return fmt.Sprintf("ngKSI already in use (%d)", Cause5GMMngKSIAlreadyInUse)
	case Cause5GMMNon3GPPAccessTo5GCNNotAllowed:
		return fmt.Sprintf("Non-3GPP access to 5GCN not allowed (%d)", Cause5GMMNon3GPPAccessTo5GCNNotAllowed)
	case Cause5GMMServingNetworkNotAuthorized:
		return fmt.Sprintf("Serving network not authorized (%d)", Cause5GMMServingNetworkNotAuthorized)
	case Cause5GMMPayloadWasNotForwarded:
		return fmt.Sprintf("Payload was not forwarded (%d)", Cause5GMMPayloadWasNotForwarded)
	case Cause5GMMDNNNotSupportedOrNotSubscribedInTheSlice:
		return fmt.Sprintf("DNN not supported or not subscribed in the slice (%d)",
			Cause5GMMDNNNotSupportedOrNotSubscribedInTheSlice)
	case Cause5GMMInsufficientUserPlaneResourcesForThePDUSession:
		return fmt.Sprintf("Insufficient user plane resources for the PDU session (%d)",
			Cause5GMMInsufficientUserPlaneResourcesForThePDUSession)
	case Cause5GMMSemanticallyIncorrectMessage:
		return fmt.Sprintf("Semantically incorrect message (%d)", Cause5GMMSemanticallyIncorrectMessage)
	case Cause5GMMInvalidMandatoryInformation:
		return fmt.Sprintf("Invalid mandatory information (%d)", Cause5GMMInvalidMandatoryInformation)
	case Cause5GMMMessageTypeNonExistentOrNotImplemented:
		return fmt.Sprintf("Message type non existent or not implemented (%d)",
			Cause5GMMMessageTypeNonExistentOrNotImplemented)
	case Cause5GMMMessageTypeNotCompatibleWithTheProtocolState:
		return fmt.Sprintf("Message type not compatible with the protocol state (%d)",
			Cause5GMMMessageTypeNotCompatibleWithTheProtocolState)
	case Cause5GMMInformationElementNonExistentOrNotImplemented:
		return fmt.Sprintf("Information element non existent or not implemented (%d)",
			Cause5GMMInformationElementNonExistentOrNotImplemented)
	case Cause5GMMConditionalIEError:
		return fmt.Sprintf("Conditional IE error (%d)", Cause5GMMConditionalIEError)
	case Cause5GMMMessageNotCompatibleWithTheProtocolState:
		return fmt.Sprintf("Message not compatible with the protocol state (%d)",
			Cause5GMMMessageNotCompatibleWithTheProtocolState)
	case Cause5GMMProtocolErrorUnspecified:
		return fmt.Sprintf("Protocol error unspecified (%d)", Cause5GMMProtocolErrorUnspecified)
	default:
		return ""
	}
}

// Cause5GSMToString renders a 5GSM cause value as its TS 24.501 clause 9.11.4.2 description
// followed by the value itself. It returns an empty string for a value the specification does
// not define, which is also how Cause5GMMToString reports an unknown 5GMM cause.
//
// 5GSM and 5GMM causes are separate registers that share numeric values - 69 is "insufficient
// resources for specific slice" in both, while 91 is a 5GMM cause with no 5GSM counterpart - so
// they need separate functions. A cause from the wrong register written into the 5GSM cause IE
// is decoded by the UE as 31 "request rejected, unspecified" per clause 9.11.4.2, which makes it
// silent; keeping the two apart is what lets a caller check the register it actually means.
func Cause5GSMToString(cause uint8) string {
	switch cause {
	case Cause5GSMOperatorDeterminedBarring:
		return fmt.Sprintf("Operator determined barring (%d)", Cause5GSMOperatorDeterminedBarring)
	case Cause5GSMInsufficientResources:
		return fmt.Sprintf("Insufficient resources (%d)", Cause5GSMInsufficientResources)
	case Cause5GSMMissingOrUnknownDNN:
		return fmt.Sprintf("Missing or unknown DNN (%d)", Cause5GSMMissingOrUnknownDNN)
	case Cause5GSMUnknownPDUSessionType:
		return fmt.Sprintf("Unknown PDU session type (%d)", Cause5GSMUnknownPDUSessionType)
	case Cause5GSMUserAuthenticationOrAuthorizationFailed:
		return fmt.Sprintf("User authentication or authorization failed (%d)", Cause5GSMUserAuthenticationOrAuthorizationFailed)
	case Cause5GSMRequestRejectedUnspecified:
		return fmt.Sprintf("Request rejected, unspecified (%d)", Cause5GSMRequestRejectedUnspecified)
	case Cause5GSMServiceOptionNotSupported:
		return fmt.Sprintf("Service option not supported (%d)", Cause5GSMServiceOptionNotSupported)
	case Cause5GSMRequestedServiceOptionNotSubscribed:
		return fmt.Sprintf("Requested service option not subscribed (%d)", Cause5GSMRequestedServiceOptionNotSubscribed)
	case Cause5GSMPTIAlreadyInUse:
		return fmt.Sprintf("PTI already in use (%d)", Cause5GSMPTIAlreadyInUse)
	case Cause5GSMRegularDeactivation:
		return fmt.Sprintf("Regular deactivation (%d)", Cause5GSMRegularDeactivation)
	case Cause5GSM5GSQoSNotAccepted:
		return fmt.Sprintf("5GS QoS not accepted (%d)", Cause5GSM5GSQoSNotAccepted)
	case Cause5GSMNetworkFailure:
		return fmt.Sprintf("Network failure (%d)", Cause5GSMNetworkFailure)
	case Cause5GSMReactivationRequested:
		return fmt.Sprintf("Reactivation requested (%d)", Cause5GSMReactivationRequested)
	case Cause5GSMSemanticErrorInTheTFTOperation:
		return fmt.Sprintf("Semantic error in the TFT operation (%d)", Cause5GSMSemanticErrorInTheTFTOperation)
	case Cause5GSMSyntacticalErrorInTheTFTOperation:
		return fmt.Sprintf("Syntactical error in the TFT operation (%d)", Cause5GSMSyntacticalErrorInTheTFTOperation)
	case Cause5GSMInvalidPDUSessionIdentity:
		return fmt.Sprintf("Invalid PDU session identity (%d)", Cause5GSMInvalidPDUSessionIdentity)
	case Cause5GSMSemanticErrorsInPacketFilter:
		return fmt.Sprintf("Semantic errors in packet filter(s) (%d)", Cause5GSMSemanticErrorsInPacketFilter)
	case Cause5GSMSyntacticalErrorInPacketFilter:
		return fmt.Sprintf("Syntactical error in packet filter(s) (%d)", Cause5GSMSyntacticalErrorInPacketFilter)
	case Cause5GSMOutOfLADNServiceArea:
		return fmt.Sprintf("Out of LADN service area (%d)", Cause5GSMOutOfLADNServiceArea)
	case Cause5GSMPTIMismatch:
		return fmt.Sprintf("PTI mismatch (%d)", Cause5GSMPTIMismatch)
	case Cause5GSMPDUSessionTypeIPv4OnlyAllowed:
		return fmt.Sprintf("PDU session type IPv4 only allowed (%d)", Cause5GSMPDUSessionTypeIPv4OnlyAllowed)
	case Cause5GSMPDUSessionTypeIPv6OnlyAllowed:
		return fmt.Sprintf("PDU session type IPv6 only allowed (%d)", Cause5GSMPDUSessionTypeIPv6OnlyAllowed)
	case Cause5GSMPDUSessionDoesNotExist:
		return fmt.Sprintf("PDU session does not exist (%d)", Cause5GSMPDUSessionDoesNotExist)
	case Cause5GSMPDUSessionTypeIPv4v6OnlyAllowed:
		return fmt.Sprintf("PDU session type IPv4v6 only allowed (%d)", Cause5GSMPDUSessionTypeIPv4v6OnlyAllowed)
	case Cause5GSMPDUSessionTypeUnstructuredOnlyAllowed:
		return fmt.Sprintf("PDU session type Unstructured only allowed (%d)", Cause5GSMPDUSessionTypeUnstructuredOnlyAllowed)
	case Cause5GSMUnsupported5QIValue:
		return fmt.Sprintf("Unsupported 5QI value (%d)", Cause5GSMUnsupported5QIValue)
	case Cause5GSMPDUSessionTypeEthernetOnlyAllowed:
		return fmt.Sprintf("PDU session type Ethernet only allowed (%d)", Cause5GSMPDUSessionTypeEthernetOnlyAllowed)
	case Cause5GSMInsufficientResourcesForSpecificSliceAndDNN:
		return fmt.Sprintf("Insufficient resources for specific slice and DNN (%d)", Cause5GSMInsufficientResourcesForSpecificSliceAndDNN)
	case Cause5GSMNotSupportedSSCMode:
		return fmt.Sprintf("Not supported SSC mode (%d)", Cause5GSMNotSupportedSSCMode)
	case Cause5GSMInsufficientResourcesForSpecificSlice:
		return fmt.Sprintf("Insufficient resources for specific slice (%d)", Cause5GSMInsufficientResourcesForSpecificSlice)
	case Cause5GSMMissingOrUnknownDNNInASlice:
		return fmt.Sprintf("Missing or unknown DNN in a slice (%d)", Cause5GSMMissingOrUnknownDNNInASlice)
	case Cause5GSMInvalidPTIValue:
		return fmt.Sprintf("Invalid PTI value (%d)", Cause5GSMInvalidPTIValue)
	case Cause5GSMMaximumDataRatePerUEForUserPlaneIntegrityProtectionIsTooLow:
		return fmt.Sprintf("Maximum data rate per UE for user-plane integrity protection is too low (%d)", Cause5GSMMaximumDataRatePerUEForUserPlaneIntegrityProtectionIsTooLow)
	case Cause5GSMSemanticErrorInTheQoSOperation:
		return fmt.Sprintf("Semantic error in the QoS operation (%d)", Cause5GSMSemanticErrorInTheQoSOperation)
	case Cause5GSMSyntacticalErrorInTheQoSOperation:
		return fmt.Sprintf("Syntactical error in the QoS operation (%d)", Cause5GSMSyntacticalErrorInTheQoSOperation)
	case Cause5GSMInvalidMappedEPSBearerIdentity:
		return fmt.Sprintf("Invalid mapped EPS bearer identity (%d)", Cause5GSMInvalidMappedEPSBearerIdentity)
	case Cause5GSMUASServicesNotAllowed:
		return fmt.Sprintf("UAS services not allowed (%d)", Cause5GSMUASServicesNotAllowed)
	case Cause5GSMSemanticallyIncorrectMessage:
		return fmt.Sprintf("Semantically incorrect message (%d)", Cause5GSMSemanticallyIncorrectMessage)
	case Cause5GSMInvalidMandatoryInformation:
		return fmt.Sprintf("Invalid mandatory information (%d)", Cause5GSMInvalidMandatoryInformation)
	case Cause5GSMMessageTypeNonExistentOrNotImplemented:
		return fmt.Sprintf("Message type non-existent or not implemented (%d)", Cause5GSMMessageTypeNonExistentOrNotImplemented)
	case Cause5GSMMessageTypeNotCompatibleWithTheProtocolState:
		return fmt.Sprintf("Message type not compatible with the protocol state (%d)", Cause5GSMMessageTypeNotCompatibleWithTheProtocolState)
	case Cause5GSMInformationElementNonExistentOrNotImplemented:
		return fmt.Sprintf("Information element non-existent or not implemented (%d)", Cause5GSMInformationElementNonExistentOrNotImplemented)
	case Cause5GSMConditionalIEError:
		return fmt.Sprintf("Conditional IE error (%d)", Cause5GSMConditionalIEError)
	case Cause5GSMMessageNotCompatibleWithTheProtocolState:
		return fmt.Sprintf("Message not compatible with the protocol state (%d)", Cause5GSMMessageNotCompatibleWithTheProtocolState)
	case Cause5GSMProtocolErrorUnspecified:
		return fmt.Sprintf("Protocol error, unspecified (%d)", Cause5GSMProtocolErrorUnspecified)
	default:
		// Cause5GSMServiceOptionTemporarilyOutOfOrder (34) reaches this arm deliberately. The
		// constant is kept for compatibility, but neither table 9.11.4.2.1 nor annex B of
		// TS 24.501 V18.13.0 defines 34 for 5GSM, and a UE decodes an undefined cause as 31.
		// Rendering it would report a value that must not be sent as though it were sendable.
		return ""
	}
}

// TS 24.501 9.11.3.18D
const (
	ControlPlaneServiceTypeMobileOriginatingRequest  uint8 = 0x00
	ControlPlaneServiceTypeMobileTerminatingRequest  uint8 = 0x01
	ControlPlaneServiceTypeEmergencyServices         uint8 = 0x02
	ControlPlaneServiceTypeEmergencyServicesFallback uint8 = 0x03
)
