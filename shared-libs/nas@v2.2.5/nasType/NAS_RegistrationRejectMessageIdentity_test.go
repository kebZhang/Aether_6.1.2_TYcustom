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

func TestNasTypeNewRegistrationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewRegistrationRejectMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeRegistrationRejectMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationRejectMessageIdentityTable = []nasTypeRegistrationRejectMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeRegistrationRejectMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewRegistrationRejectMessageIdentity()
	for _, table := range nasTypeRegistrationRejectMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type RegistrationRejectMessageIdentityTestDataTemplate struct {
	in  nasType.RegistrationRejectMessageIdentity
	out nasType.RegistrationRejectMessageIdentity
}

var RegistrationRejectMessageIdentityTestData = []nasType.RegistrationRejectMessageIdentity{
	{0x03},
}

var RegistrationRejectMessageIdentityExpectedTestData = []nasType.RegistrationRejectMessageIdentity{
	{0x03},
}

var RegistrationRejectMessageIdentityTable = []RegistrationRejectMessageIdentityTestDataTemplate{
	{RegistrationRejectMessageIdentityTestData[0], RegistrationRejectMessageIdentityExpectedTestData[0]},
}

func TestNasTypeRegistrationRejectMessageIdentity(t *testing.T) {
	for _, table := range RegistrationRejectMessageIdentityTable {
		a := nasType.NewRegistrationRejectMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
