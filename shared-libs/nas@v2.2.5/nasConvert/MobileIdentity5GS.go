// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"
	"fmt"
	"math/bits"
	"strconv"
	"strings"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/omec-project/openapi/v2/models"
)

func GetTypeOfIdentity(buf byte) uint8 {
	return buf & 0x07
}

// TS 24.501 9.11.3.4
// suci(imsi) =
// "suci-0-${mcc}-${mnc}-${routingIndentifier}-${protectionScheme}-${homeNetworkPublicKeyIdentifier}-${schemeOutput}"
// suci(nai) = NAI string returned by NaiToString, for example "nai-1-..."
func SuciToString(buf []byte) (suci string, plmnId string, err error) {
	var mcc, mnc, routingInd, protectionScheme, homeNetworkPublicKeyIdentifier, schemeOutput string

	if len(buf) == 0 {
		return "", "", fmt.Errorf("empty SUCI buffer")
	}

	supiFormat := (buf[0] & 0xf0) >> 4
	if supiFormat == nasMessage.SupiFormatNai {
		suci, err = NaiToString(buf)
		return suci, "", err
	}

	if len(buf) < 9 {
		return "", "", fmt.Errorf("invalid SUCI buffer length: %d", len(buf))
	}

	// Encode buf to SUCI in supi format "IMSI"

	// Plmn(MCC + MNC)
	mccDigit3 := (buf[2] & 0x0f)
	tmpBytes := []byte{bits.RotateLeft8(buf[1], 4), (mccDigit3 << 4)}
	mcc = hex.EncodeToString(tmpBytes)
	mcc = mcc[:3] // remove rear 0

	mncDigit3 := (buf[2] & 0xf0) >> 4
	tmpBytes = []byte{bits.RotateLeft8(buf[3], 4), mncDigit3 << 4}
	mnc = hex.EncodeToString(tmpBytes)
	if mnc[2] == 'f' {
		mnc = mnc[:2] // mnc is 2 digit -> remove 'f'
	} else {
		mnc = mnc[:3] // mnc is 3 digit -> remove rear 0
	}
	plmnId = mcc + mnc

	// Routing Indicator
	var routingIndBytes []byte
	routingIndBytes = append(routingIndBytes, bits.RotateLeft8(buf[4], 4))
	routingIndBytes = append(routingIndBytes, bits.RotateLeft8(buf[5], 4))
	routingInd = hex.EncodeToString(routingIndBytes)

	if idx := strings.Index(routingInd, "f"); idx != -1 {
		routingInd = routingInd[0:idx]
	}

	// Protection Scheme
	protectionScheme = fmt.Sprintf("%x", buf[6]) // convert byte to hex string without leading 0s

	// Home Network Public Key Indentifier
	homeNetworkPublicKeyIdentifier = fmt.Sprintf("%d", buf[7])

	// Scheme output
	if protectionScheme == strconv.Itoa(nasMessage.ProtectionSchemeNullScheme) {
		// MSIN
		var msinBytes []byte
		for i := 8; i < len(buf); i++ {
			msinBytes = append(msinBytes, bits.RotateLeft8(buf[i], 4))
		}
		schemeOutput = hex.EncodeToString(msinBytes)
		if len(schemeOutput) > 0 && schemeOutput[len(schemeOutput)-1] == 'f' {
			schemeOutput = schemeOutput[:len(schemeOutput)-1]
		}
	} else {
		schemeOutput = hex.EncodeToString(buf[8:])
	}

	suci = strings.Join([]string{
		"suci", "0", mcc, mnc, routingInd, protectionScheme, homeNetworkPublicKeyIdentifier,
		schemeOutput,
	}, "-")
	return suci, plmnId, nil
}

func NaiToString(buf []byte) (nai string, err error) {
	if len(buf) < 2 {
		return "", fmt.Errorf("invalid NAI buffer")
	}

	prefix := "nai"
	naiBytes := buf[1:]
	naiStr := hex.EncodeToString(naiBytes)
	nai = strings.Join([]string{prefix, "1", naiStr}, "-")
	return nai, nil
}

// nasType: TS 24.501 9.11.3.4
func GutiToString(buf []byte) (guami models.Guami, guti string) {
	if len(buf) != 11 {
		logger.ConvertLog.Errorf("invalid GUTI buffer length: %d", len(buf))
		return
	}

	plmnID := PlmnIDToString(buf[1:4])
	amfID := hex.EncodeToString(buf[4:7])
	tmsi5G := hex.EncodeToString(buf[7:])

	guami.PlmnId = models.PlmnIdNid{
		Mcc: plmnID[:3],
		Mnc: plmnID[3:],
	}
	guami.AmfId = amfID
	guti = plmnID + amfID + tmsi5G
	return
}

func GutiToNas(guti string) nasType.GUTI5G {
	var gutiNas nasType.GUTI5G

	gutiNas.SetLen(11)
	gutiNas.SetSpare(0)
	gutiNas.SetSpare2(15)
	gutiNas.SetTypeOfIdentity(nasMessage.MobileIdentity5GSType5gGuti)

	var mcc1, mcc2, mcc3, mnc1, mnc2, mnc3 int
	if mcc1Tmp, err := strconv.Atoi(string(guti[0])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc1 error: %+v", err)
	} else {
		mcc1 = mcc1Tmp
	}
	if mcc2Tmp, err := strconv.Atoi(string(guti[1])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc2 error: %+v", err)
	} else {
		mcc2 = mcc2Tmp
	}
	if mcc3Tmp, err := strconv.Atoi(string(guti[2])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc3 error: %+v", err)
	} else {
		mcc3 = mcc3Tmp
	}
	if mnc1Tmp, err := strconv.Atoi(string(guti[3])); err != nil {
		logger.ConvertLog.Warnf("atoi mnc1 error: %+v", err)
	} else {
		mnc1 = mnc1Tmp
	}
	if mnc2Tmp, err := strconv.Atoi(string(guti[4])); err != nil {
		logger.ConvertLog.Warnf("atoi mnc2 error: %+v", err)
	} else {
		mnc2 = mnc2Tmp
	}
	mnc3 = 0x0f
	amfId := ""
	tmsi := ""
	if len(guti) == 20 {
		if mnc3Tmp, err := strconv.Atoi(string(guti[5])); err != nil {
			logger.ConvertLog.Warnf("atoi guti error: %+v", err)
		} else {
			mnc3 = mnc3Tmp
		}
		amfId = guti[6:12]
		tmsi = guti[12:]
	} else {
		amfId = guti[5:11]
		tmsi = guti[11:]
	}
	gutiNas.SetMCCDigit1(uint8(mcc1))
	gutiNas.SetMCCDigit2(uint8(mcc2))
	gutiNas.SetMCCDigit3(uint8(mcc3))
	gutiNas.SetMNCDigit1(uint8(mnc1))
	gutiNas.SetMNCDigit2(uint8(mnc2))
	gutiNas.SetMNCDigit3(uint8(mnc3))

	amfRegionId, amfSetId, amfPointer := AmfIdToNas(amfId)
	gutiNas.SetAMFRegionID(amfRegionId)
	gutiNas.SetAMFSetID(amfSetId)
	gutiNas.SetAMFPointer(amfPointer)
	if tmsiBytes, err := hex.DecodeString(tmsi); err != nil {
		logger.ConvertLog.Warnf("decode TMSI failed: %+v", err)
	} else {
		copy(gutiNas.Octet[7:11], tmsiBytes[:])
	}
	return gutiNas
}

// PEI: ^(imei-[0-9]{15}|imeisv-[0-9]{16}|.+)$
func PeiToString(pei []byte) string {
	if len(pei) == 0 {
		logger.ConvertLog.Errorln("empty PEI")
		return ""
	}

	const imei_prefix = "imei"
	const imeisv_prefix = "imeisv"
	var digits strings.Builder

	prefix := ""

	lower := pei[0] & 0x0F
	upper := (pei[0] & 0xF0) >> 4

	typeIdentity := lower & 0x07
	switch typeIdentity {
	case 3:
		prefix = imei_prefix
	case 5:
		prefix = imeisv_prefix
	case 6: // mac-address
		logger.ConvertLog.Errorf("unsupported type of identity: %d", typeIdentity)
		return prefix
	default:
		logger.ConvertLog.Errorf("invalid type of identity: %d", typeIdentity)
		return prefix
	}

	if upper <= 9 {
		digits.WriteByte('0' + upper)
	} else {
		logger.ConvertLog.Errorf("invalid value/character 0x%x", upper)
		return ""
	}

	for i, b := range pei[1:] {
		lower = b & 0x0F
		upper = (b & 0xF0) >> 4

		if lower <= 9 {
			digits.WriteByte('0' + lower)
		} else {
			logger.ConvertLog.Errorf("invalid value/character 0x%x", lower)
			return ""
		}

		if upper <= 9 {
			digits.WriteByte('0' + upper)
		} else if i == len(pei)-2 && upper == 0x0f {
			// last digit is filler
		} else {
			logger.ConvertLog.Errorf("invalid value/character 0x%x", upper)
			return ""
		}
	}

	digitStr := digits.String()

	// Validation for IMEI and IMEISV
	digitStrLen := len(digitStr)
	if (prefix == imei_prefix && digitStrLen != 15) || (prefix == imeisv_prefix && digitStrLen != 16) {
		logger.ConvertLog.Errorf("invalid %s length: %d", prefix, digitStrLen)
		return ""
	}

	// Validate TAC and SNR using the Luhn formula
	if (prefix == imei_prefix) && !isValidLuhn(digitStr) {
		logger.ConvertLog.Errorf("invalid TAC/SNR in %s: %s", prefix, digitStr)
		return ""
	}

	return prefix + "-" + digitStr
}

// Validate a string using the Luhn formula
func isValidLuhn(input string) bool {
	sum := 0
	alt := false
	for _, char := range input[:14] {
		n := int(char - '0')
		if alt {
			n *= 2
			if n > 9 {
				n = n/10 + n%10
			}
		}
		sum += n
		alt = !alt
	}
	cd := (10 - sum%10) % 10
	logger.ConvertLog.Debugf("Luhn sum: %d, and cd: %d", sum, cd)
	return cd == int(input[len(input)-1]-'0')
}
