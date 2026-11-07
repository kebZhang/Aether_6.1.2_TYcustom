// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/omec-project/openapi/v2/models"
)

func SpareHalfOctetAndNgksiToModels(ngKsiNas nasType.SpareHalfOctetAndNgksi) (ngKsiModels models.NgKsi) {
	switch ngKsiNas.GetTSC() {
	case nasMessage.TypeOfSecurityContextFlagNative:
		ngKsiModels.Tsc = models.SCTYPE_NATIVE
	case nasMessage.TypeOfSecurityContextFlagMapped:
		ngKsiModels.Tsc = models.SCTYPE_MAPPED
	}

	ngKsiModels.Ksi = int32(ngKsiNas.GetNasKeySetIdentifiler())
	return
}

func SpareHalfOctetAndNgksiToNas(ngKsiModels models.NgKsi) (ngKsiNas nasType.SpareHalfOctetAndNgksi) {
	switch ngKsiModels.Tsc {
	case models.SCTYPE_NATIVE:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagNative)
	case models.SCTYPE_MAPPED:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagMapped)
	}

	ngKsiNas.SetSpareHalfOctet(0)
	ngKsiNas.SetNasKeySetIdentifiler(uint8(ngKsiModels.Ksi))
	return
}
