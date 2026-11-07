// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasType"
)

type PDUSessionEstablishmentRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity
	nasType.IntegrityProtectionMaximumDataRate
	*nasType.PDUSessionType
	*nasType.SSCMode
	*nasType.Capability5GSM
	*nasType.MaximumNumberOfSupportedPacketFilters
	*nasType.AlwaysonPDUSessionRequested
	*nasType.SMPDUDNRequestContainer
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionEstablishmentRequest(iei uint8) (pDUSessionEstablishmentRequest *PDUSessionEstablishmentRequest) {
	pDUSessionEstablishmentRequest = &PDUSessionEstablishmentRequest{}
	return pDUSessionEstablishmentRequest
}

const (
	PDUSessionEstablishmentRequestPDUSessionTypeType                        uint8 = 0x09
	PDUSessionEstablishmentRequestSSCModeType                               uint8 = 0x0A
	PDUSessionEstablishmentRequestCapability5GSMType                        uint8 = 0x28
	PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType uint8 = 0x55
	PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType           uint8 = 0x0B
	PDUSessionEstablishmentRequestSMPDUDNRequestContainerType               uint8 = 0x39
	PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType  uint8 = 0x7B
)

func (a *PDUSessionEstablishmentRequest) EncodePDUSessionEstablishmentRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet); err != nil {
		return
	}
	if a.PDUSessionType != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionType.Octet); err != nil {
			return
		}
	}
	if a.SSCMode != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.SSCMode.Octet); err != nil {
			return
		}
	}
	if a.Capability5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
			return
		}
	}
	if a.MaximumNumberOfSupportedPacketFilters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet); err != nil {
			return
		}
	}
	if a.AlwaysonPDUSessionRequested != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionRequested.Octet); err != nil {
			return
		}
	}
	if a.SMPDUDNRequestContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.SMPDUDNRequestContainer.Buffer); err != nil {
			return
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		logger.NasMsgLog.Infoln("Encode ExtendedProtocolConfigurationOptions in EncodePDUSessionEstablishmentRequest")
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return
		}
	}
}

func (a *PDUSessionEstablishmentRequest) DecodePDUSessionEstablishmentRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet); err != nil {
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
		case PDUSessionEstablishmentRequestPDUSessionTypeType:
			a.PDUSessionType = nasType.NewPDUSessionType(ieiN)
			a.PDUSessionType.Octet = ieiN
		case PDUSessionEstablishmentRequestSSCModeType:
			a.SSCMode = nasType.NewSSCMode(ieiN)
			a.SSCMode.Octet = ieiN
		case PDUSessionEstablishmentRequestCapability5GSMType:
			a.Capability5GSM = nasType.NewCapability5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GSM.Len); err != nil {
				return
			}
			a.Capability5GSM.SetLen(a.Capability5GSM.GetLen())
			if a.Capability5GSM.GetLen() > uint8(len(a.Capability5GSM.Octet)) {
				a.Capability5GSM = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return
			}
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
				return
			}
		case PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType:
			a.MaximumNumberOfSupportedPacketFilters = nasType.NewMaximumNumberOfSupportedPacketFilters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet); err != nil {
				return
			}
		case PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType:
			a.AlwaysonPDUSessionRequested = nasType.NewAlwaysonPDUSessionRequested(ieiN)
			a.AlwaysonPDUSessionRequested.Octet = ieiN
		case PDUSessionEstablishmentRequestSMPDUDNRequestContainerType:
			a.SMPDUDNRequestContainer = nasType.NewSMPDUDNRequestContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SMPDUDNRequestContainer.Len); err != nil {
				return
			}
			a.SMPDUDNRequestContainer.SetLen(a.SMPDUDNRequestContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.Buffer[:a.SMPDUDNRequestContainer.GetLen()]); err != nil {
				return
			}
		case PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return
			}
		default:
		}
	}
}
