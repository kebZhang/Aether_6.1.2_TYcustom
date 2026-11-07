// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/openapi/v2/models"
)

func PDUSessionTypeToModels(nasPduSessType uint8) (pduSessType models.PduSessionType) {
	switch nasPduSessType {
	case nasMessage.PDUSessionTypeIPv4:
		pduSessType = models.PDUSESSIONTYPE_IPV4
	case nasMessage.PDUSessionTypeIPv6:
		pduSessType = models.PDUSESSIONTYPE_IPV6
	case nasMessage.PDUSessionTypeIPv4IPv6:
		pduSessType = models.PDUSESSIONTYPE_IPV4_V6
	case nasMessage.PDUSessionTypeUnstructured:
		pduSessType = models.PDUSESSIONTYPE_UNSTRUCTURED
	case nasMessage.PDUSessionTypeEthernet:
		pduSessType = models.PDUSESSIONTYPE_ETHERNET
	}

	return
}

func ModelsToPDUSessionType(pduSessType models.PduSessionType) (nasPduSessType uint8) {
	switch pduSessType {
	case models.PDUSESSIONTYPE_IPV4:
		nasPduSessType = nasMessage.PDUSessionTypeIPv4
	case models.PDUSESSIONTYPE_IPV6:
		nasPduSessType = nasMessage.PDUSessionTypeIPv6
	case models.PDUSESSIONTYPE_IPV4_V6:
		nasPduSessType = nasMessage.PDUSessionTypeIPv4IPv6
	case models.PDUSESSIONTYPE_UNSTRUCTURED:
		nasPduSessType = nasMessage.PDUSessionTypeUnstructured
	case models.PDUSESSIONTYPE_ETHERNET:
		nasPduSessType = nasMessage.PDUSessionTypeEthernet
	}
	return
}
