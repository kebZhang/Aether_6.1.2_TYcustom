// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"
	"fmt"

	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/omec-project/openapi/v2"
	"github.com/omec-project/openapi/v2/models"
)

// TS 24.501 9.11.3.37
func RequestedNssaiToModels(nasNssai *nasType.RequestedNSSAI) ([]models.MappingOfSnssai, error) {
	var requestNssai []models.MappingOfSnssai

	buf := nasNssai.GetSNSSAIValue()
	lengthOfBuf := int(nasNssai.GetLen())
	if lengthOfBuf > len(buf) {
		return nil, fmt.Errorf("requested NSSAI length %d exceeds buffer length %d", lengthOfBuf, len(buf))
	}
	// enforce the advertised length so trailing bytes beyond it are never parsed
	buf = buf[:lengthOfBuf]

	offset := 0
	for offset < lengthOfBuf {
		lengthOfSnssaiContents := buf[offset]
		if snssai, err := snssaiToModels(lengthOfSnssaiContents, buf[offset:]); err != nil {
			return nil, err
		} else {
			requestNssai = append(requestNssai, snssai)
			// lengthOfSnssaiContents is 1 byte
			offset += int(lengthOfSnssaiContents + 1)
		}
	}

	return requestNssai, nil
}

// TS 24.501 9.11.2.8, Length & value part of S-NSSAI IE
func snssaiToModels(lengthOfSnssaiContents uint8, buf []byte) (models.MappingOfSnssai, error) {
	snssai := models.MappingOfSnssai{}

	if int(lengthOfSnssaiContents)+1 > len(buf) {
		return snssai, fmt.Errorf("S-NSSAI contents length %d exceeds remaining buffer", lengthOfSnssaiContents)
	}

	switch lengthOfSnssaiContents {
	case 0x01: // SST
		snssai.ServingSnssai = models.Snssai{
			Sst: int32(buf[1]),
		}
		return snssai, nil
	case 0x02: // SST and mapped HPLMN SST
		snssai.ServingSnssai = models.Snssai{
			Sst: int32(buf[1]),
		}
		snssai.HomeSnssai = models.Snssai{
			Sst: int32(buf[2]),
		}
		return snssai, nil
	case 0x04: // SST and SD
		snssai.ServingSnssai = models.Snssai{
			Sst: int32(buf[1]),
			Sd:  openapi.PtrString(hex.EncodeToString(buf[2:5])),
		}
		return snssai, nil
	case 0x05: // SST, SD and mapped HPLMN SST
		snssai.ServingSnssai = models.Snssai{
			Sst: int32(buf[1]),
			Sd:  openapi.PtrString(hex.EncodeToString(buf[2:5])),
		}
		snssai.HomeSnssai = models.Snssai{
			Sst: int32(buf[5]),
		}
		return snssai, nil
	case 0x08: // SST, SD, mapped HPLMN SST and mapped HPLMN SD
		snssai.ServingSnssai = models.Snssai{
			Sst: int32(buf[1]),
			Sd:  openapi.PtrString(hex.EncodeToString(buf[2:5])),
		}
		snssai.HomeSnssai = models.Snssai{
			Sst: int32(buf[5]),
			Sd:  openapi.PtrString(hex.EncodeToString(buf[6:9])),
		}
		return snssai, nil
	default:
		return snssai, fmt.Errorf("invalid length of S-NSSAI contents: %d", lengthOfSnssaiContents)
	}
}

func RejectedNssaiToNas(rejectedNssaiInPlmn []models.Snssai, rejectedNssaiInTa []models.Snssai) nasType.RejectedNSSAI {
	var rejectedNssaiNas nasType.RejectedNSSAI

	var byteArray []uint8
	for _, rejectedSnssai := range rejectedNssaiInPlmn {
		byteArray = append(byteArray, RejectedSnssaiToNas(rejectedSnssai,
			nasMessage.RejectedSnssaiCauseNotAvailableInCurrentPlmn)...)
	}
	for _, rejectedSnssai := range rejectedNssaiInTa {
		byteArray = append(byteArray, RejectedSnssaiToNas(rejectedSnssai,
			nasMessage.RejectedSnssaiCauseNotAvailableInCurrentRegistrationArea)...)
	}

	rejectedNssaiNas.SetLen(uint8(len(byteArray)))
	rejectedNssaiNas.SetRejectedNSSAIContents(byteArray)
	return rejectedNssaiNas
}
