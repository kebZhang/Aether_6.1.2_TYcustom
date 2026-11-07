// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"github.com/omec-project/openapi/v2/models"
)

// TS 24.501 9.11.3.29 LADN indication: a sequence of length-prefixed DNN values
func LadnToModels(buf []uint8) (dnnValues []string) {
	// at most 8 LADN DNN values are considered; remaining octets are ignored
	const maxLadnDnnValues = 8

	for bufOffset := 0; bufOffset < len(buf) && len(dnnValues) < maxLadnDnnValues; {
		lenOfDnn := int(buf[bufOffset])
		// reject a zero length (would never advance bufOffset) or one that overruns buf,
		// accounting for the length octet itself
		if lenOfDnn == 0 || bufOffset+1+lenOfDnn > len(buf) {
			break
		}
		dnn := string(buf[bufOffset+1 : bufOffset+1+lenOfDnn])
		dnnValues = append(dnnValues, dnn)
		bufOffset += 1 + lenOfDnn
	}

	return
}

// TS 24.501 9.11.3.30 LADN information: one [DNN][tracking area identity list] entry
func LadnToNas(dnn string, taiLists []models.Tai) (ladnNas []uint8) {
	dnnNas := []byte(dnn)

	ladnNas = append(ladnNas, uint8(len(dnnNas)))
	ladnNas = append(ladnNas, dnnNas...)

	taiListNas := TaiListToNas(taiLists)
	ladnNas = append(ladnNas, uint8(len(taiListNas)))
	ladnNas = append(ladnNas, taiListNas...)
	return
}
