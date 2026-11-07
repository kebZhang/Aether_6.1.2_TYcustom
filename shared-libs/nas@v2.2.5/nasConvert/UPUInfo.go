// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"
	"fmt"
)

func UpuAckToModels(buf []uint8) (string, error) {
	if len(buf) != 17 || buf[0] != 0x01 {
		return "", fmt.Errorf("NAS UPU Ack is not valid")
	}
	return hex.EncodeToString(buf[1:]), nil
}
