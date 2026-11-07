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

func TestNasTypeSpareHalfOctetAndSecurityHeaderType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndSecurityHeaderType()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeSecurityHeaderTypeAndSpareHalfOctetData struct {
	inSecurityHeader  uint8
	inSpareHalfOctet  uint8
	outSecurityHeader uint8
	outSpareHalfOctet uint8
}

var nasTypeSecurityHeaderTypeAndSpareHalfOctetTable = []nasTypeSecurityHeaderTypeAndSpareHalfOctetData{
	{0x8, 0x1, 0x8, 0x01},
}

func TestNasTypeGetSetSpareHalfOctetAndSecurityHeaderType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndSecurityHeaderType()
	for _, table := range nasTypeSecurityHeaderTypeAndSpareHalfOctetTable {
		a.SetSecurityHeaderType(table.inSecurityHeader)
		if !reflect.DeepEqual(table.outSecurityHeader, a.GetSecurityHeaderType()) {
			t.Errorf("Not equal: expected %v, got %v", table.outSecurityHeader, a.GetSecurityHeaderType())
		}
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		if !reflect.DeepEqual(table.outSpareHalfOctet, a.GetSpareHalfOctet()) {
			t.Errorf("Not equal: expected %v, got %v", table.outSpareHalfOctet, a.GetSpareHalfOctet())
		}
	}
}
