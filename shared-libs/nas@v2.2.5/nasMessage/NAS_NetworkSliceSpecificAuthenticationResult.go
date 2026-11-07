// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationResult 8.2.33
type NetworkSliceSpecificAuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationResult(iei uint8) (networkSliceSpecificAuthenticationResult *NetworkSliceSpecificAuthenticationResult) {
	networkSliceSpecificAuthenticationResult = &NetworkSliceSpecificAuthenticationResult{}
	return networkSliceSpecificAuthenticationResult
}

func (a *NetworkSliceSpecificAuthenticationResult) EncodeNetworkSliceSpecificAuthenticationResult(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *NetworkSliceSpecificAuthenticationResult) DecodeNetworkSliceSpecificAuthenticationResult(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
		return
	}
	if a.SNSSAI.GetLen() > uint8(len(a.SNSSAI.Octet)) {
		a.SNSSAI.SetLen(0) // SNSSAI is mandatory (not a pointer) - reset it so a later Encode cannot re-panic on it
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
		return
	}
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}
