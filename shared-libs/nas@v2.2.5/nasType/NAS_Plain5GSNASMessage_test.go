// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}
