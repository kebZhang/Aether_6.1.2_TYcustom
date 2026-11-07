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

func TestNasTypeNewULNASTRANSPORTMessageIdentity(t *testing.T) {
	a := nasType.NewULNASTRANSPORTMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeULNASTRANSPORTMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeULNASTRANSPORTMessageIdentityTable = []nasTypeULNASTRANSPORTMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeULNASTRANSPORTMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewULNASTRANSPORTMessageIdentity()
	for _, table := range nasTypeULNASTRANSPORTMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type ULNASTRANSPORTMessageIdentityTestDataTemplate struct {
	in  nasType.ULNASTRANSPORTMessageIdentity
	out nasType.ULNASTRANSPORTMessageIdentity
}

var ULNASTRANSPORTMessageIdentityTestData = []nasType.ULNASTRANSPORTMessageIdentity{
	{0x03},
}

var ULNASTRANSPORTMessageIdentityExpectedTestData = []nasType.ULNASTRANSPORTMessageIdentity{
	{0x03},
}

var ULNASTRANSPORTMessageIdentityTable = []ULNASTRANSPORTMessageIdentityTestDataTemplate{
	{ULNASTRANSPORTMessageIdentityTestData[0], ULNASTRANSPORTMessageIdentityExpectedTestData[0]},
}

func TestNasTypeULNASTRANSPORTMessageIdentity(t *testing.T) {
	for _, table := range ULNASTRANSPORTMessageIdentityTable {
		a := nasType.NewULNASTRANSPORTMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
