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

func TestNasTypeNewSpareHalfOctetAndAccessType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndAccessType()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeAccessType struct {
	in  uint8
	out uint8
}

var nasTypeAccessTypeTable = []nasTypeAccessType{
	{0x03, 0x03},
}

func TestNasTypeGetSetAccessType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndAccessType()
	for _, table := range nasTypeAccessTypeTable {
		a.SetAccessType(table.in)
		if !reflect.DeepEqual(table.out, a.GetAccessType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAccessType())
		}
	}
}

type AccessTypeAndSpareHalfOctetTestDataTemplate struct {
	in  nasType.SpareHalfOctetAndAccessType
	out nasType.SpareHalfOctetAndAccessType
}

var accessTypeAndSpareHalfOctetTestData = []nasType.SpareHalfOctetAndAccessType{
	{0x03},
}

var accessTypeAndSpareHalfOctetExpectedTestData = []nasType.SpareHalfOctetAndAccessType{
	{0x03},
}

var accessTypeAndSpareHalfOctetTable = []AccessTypeAndSpareHalfOctetTestDataTemplate{
	{accessTypeAndSpareHalfOctetTestData[0], accessTypeAndSpareHalfOctetExpectedTestData[0]},
}

func TestNasTypeAccessTypeAndSpareHalfOctet(t *testing.T) {
	for i, table := range accessTypeAndSpareHalfOctetTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSpareHalfOctetAndAccessType()

		a.SetAccessType(table.in.GetAccessType())
		if !reflect.DeepEqual(table.out.GetAccessType(), a.GetAccessType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetAccessType(), a.GetAccessType())
		}
	}
}
