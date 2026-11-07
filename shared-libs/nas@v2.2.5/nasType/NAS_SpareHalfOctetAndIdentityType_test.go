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

type nasTypeIdentityTypeAndSpareHalfOctetData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityTypeAndSpareHalfOctetTable = []nasTypeIdentityTypeAndSpareHalfOctetData{
	{0x07, 0x07},
}

func TestNasTypeNewSpareHalfOctetAndIdentityType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndIdentityType()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeIdentityTypeAndSpareHalfOctet(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndIdentityType()
	for _, table := range nasTypeIdentityTypeAndSpareHalfOctetTable {
		a.SetTypeOfIdentity(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfIdentity()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfIdentity())
		}
	}
}
