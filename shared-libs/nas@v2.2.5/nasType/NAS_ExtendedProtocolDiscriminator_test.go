// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewExtendedProtocolDiscriminatort(t *testing.T) {
	a := nasType.NewExtendedProtocolDiscriminator()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeExtendedProtocolDiscriminatorData struct {
	in  uint8
	out uint8
}

var nasTypeExtendedProtocolDiscriminatorTable = []nasTypeExtendedProtocolDiscriminatorData{
	{2, 2},
}

func TestNasTypeGetSetExtendedProtocolDiscriminator(t *testing.T) {
	a := nasType.NewExtendedProtocolDiscriminator()
	for _, table := range nasTypeExtendedProtocolDiscriminatorTable {
		a.SetExtendedProtocolDiscriminator(table.in)
		if !reflect.DeepEqual(table.out, a.GetExtendedProtocolDiscriminator()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetExtendedProtocolDiscriminator())
		}
	}
}
