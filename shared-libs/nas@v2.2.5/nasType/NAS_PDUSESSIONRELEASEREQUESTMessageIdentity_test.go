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

func TestNasTypeNewPDUSESSIONRELEASEREQUESTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONRELEASEREQUESTMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONRELEASEREQUESTMessageIdentityTable = []nasTypePDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypePDUSESSIONRELEASEREQUESTMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()
	for _, table := range nasTypePDUSESSIONRELEASEREQUESTMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type PDUSESSIONRELEASEREQUESTMessageIdentityTestDataTemplate struct {
	in  nasType.PDUSESSIONRELEASEREQUESTMessageIdentity
	out nasType.PDUSESSIONRELEASEREQUESTMessageIdentity
}

var PDUSESSIONRELEASEREQUESTMessageIdentityTestData = []nasType.PDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASEREQUESTMessageIdentityExpectedTestData = []nasType.PDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASEREQUESTMessageIdentityTable = []PDUSESSIONRELEASEREQUESTMessageIdentityTestDataTemplate{
	{PDUSESSIONRELEASEREQUESTMessageIdentityTestData[0], PDUSESSIONRELEASEREQUESTMessageIdentityExpectedTestData[0]},
}

func TestNasTypePDUSESSIONRELEASEREQUESTMessageIdentity(t *testing.T) {
	for _, table := range PDUSESSIONRELEASEREQUESTMessageIdentityTable {
		a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
