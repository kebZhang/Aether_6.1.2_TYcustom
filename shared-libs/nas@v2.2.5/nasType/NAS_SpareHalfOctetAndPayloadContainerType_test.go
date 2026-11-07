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

func TestNasTypeNewSpareHalfOctetAndPayloadContainerType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndPayloadContainerType()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePayloadContainerTypeAndSparePayloadContainerType struct {
	in  uint8
	out uint8
}

var nasTypePayloadContainerTypeAndSparePayloadContainerTypeTable = []nasTypePayloadContainerTypeAndSparePayloadContainerType{
	{0x0f, 0x0f},
}

func TestNasTypeGetSetPayloadSpareHalfOctetAndPayloadContainerType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndPayloadContainerType()
	for _, table := range nasTypePayloadContainerTypeAndSparePayloadContainerTypeTable {
		a.SetPayloadContainerType(table.in)
		if !reflect.DeepEqual(table.out, a.GetPayloadContainerType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPayloadContainerType())
		}
	}
}
