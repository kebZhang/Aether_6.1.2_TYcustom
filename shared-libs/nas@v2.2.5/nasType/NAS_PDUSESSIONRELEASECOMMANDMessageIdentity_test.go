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

func TestNasTypeNewPDUSESSIONRELEASECOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMMANDMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONRELEASECOMMANDMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONRELEASECOMMANDMessageIdentityTable = []nasTypePDUSESSIONRELEASECOMMANDMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypePDUSESSIONRELEASECOMMANDMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONRELEASECOMMANDMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type PDUSESSIONRELEASECOMMANDMessageIdentityTestDataTemplate struct {
	in  nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
	out nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
}

var pDUSESSIONRELEASECOMMANDMessageIdentityTestData = []nasType.PDUSESSIONRELEASECOMMANDMessageIdentity{
	{0x03},
}

var pDUSESSIONRELEASECOMMANDMessageIdentityExpectedTestData = []nasType.PDUSESSIONRELEASECOMMANDMessageIdentity{
	{0x03},
}

var pDUSESSIONRELEASECOMMANDMessageIdentityTable = []PDUSESSIONRELEASECOMMANDMessageIdentityTestDataTemplate{
	{pDUSESSIONRELEASECOMMANDMessageIdentityTestData[0], pDUSESSIONRELEASECOMMANDMessageIdentityExpectedTestData[0]},
}

func TestNasTypePDUSESSIONRELEASECOMMANDMessageIdentity(t *testing.T) {
	for _, table := range pDUSESSIONRELEASECOMMANDMessageIdentityTable {
		a := nasType.NewPDUSESSIONRELEASECOMMANDMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
