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

type ConfigurationUpdateComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ConfigurationUpdateCompleteMessageIdentity
}

func NewConfigurationUpdateComplete(iei uint8) (configurationUpdateComplete *ConfigurationUpdateComplete) {
	configurationUpdateComplete = &ConfigurationUpdateComplete{}
	return configurationUpdateComplete
}

func (a *ConfigurationUpdateComplete) EncodeConfigurationUpdateComplete(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ConfigurationUpdateCompleteMessageIdentity.Octet); err != nil {
		return
	}
}

func (a *ConfigurationUpdateComplete) DecodeConfigurationUpdateComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ConfigurationUpdateCompleteMessageIdentity.Octet); err != nil {
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
		default:
		}
	}
}
