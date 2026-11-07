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

func TestNasTypeNewRegistrationRequestMessageIdentity(t *testing.T) {
	a := nasType.NewRegistrationRequestMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeRegistrationRequestMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationRequestMessageIdentityTable = []nasTypeRegistrationRequestMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeRegistrationRequestMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewRegistrationRequestMessageIdentity()
	for _, table := range nasTypeRegistrationRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type RegistrationRequestMessageIdentityTestDataTemplate struct {
	in  nasType.RegistrationRequestMessageIdentity
	out nasType.RegistrationRequestMessageIdentity
}

var RegistrationRequestMessageIdentityTestData = []nasType.RegistrationRequestMessageIdentity{
	{0x03},
}

var RegistrationRequestMessageIdentityExpectedTestData = []nasType.RegistrationRequestMessageIdentity{
	{0x03},
}

var RegistrationRequestMessageIdentityTable = []RegistrationRequestMessageIdentityTestDataTemplate{
	{RegistrationRequestMessageIdentityTestData[0], RegistrationRequestMessageIdentityExpectedTestData[0]},
}

func TestNasTypeRegistrationRequestMessageIdentity(t *testing.T) {
	for _, table := range RegistrationRequestMessageIdentityTable {
		a := nasType.NewRegistrationRequestMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
