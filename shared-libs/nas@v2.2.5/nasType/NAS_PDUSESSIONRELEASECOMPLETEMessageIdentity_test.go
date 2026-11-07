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

func TestNasTypeNewPDUSESSIONRELEASECOMPLETEMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONRELEASECOMPLETEMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONRELEASECOMPLETEMessageIdentityTable = []nasTypePDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypePDUSESSIONRELEASECOMPLETEMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()
	for _, table := range nasTypePDUSESSIONRELEASECOMPLETEMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type PDUSESSIONRELEASECOMPLETEMessageIdentityTestDataTemplate struct {
	in  nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity
	out nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityTestData = []nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityExpectedTestData = []nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityTable = []PDUSESSIONRELEASECOMPLETEMessageIdentityTestDataTemplate{
	{PDUSESSIONRELEASECOMPLETEMessageIdentityTestData[0], PDUSESSIONRELEASECOMPLETEMessageIdentityExpectedTestData[0]},
}

func TestNasTypePDUSESSIONRELEASECOMPLETEMessageIdentity(t *testing.T) {
	for _, table := range PDUSESSIONRELEASECOMPLETEMessageIdentityTable {
		a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
