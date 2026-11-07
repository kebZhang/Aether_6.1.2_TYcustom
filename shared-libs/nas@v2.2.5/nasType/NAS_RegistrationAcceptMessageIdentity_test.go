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

func TestNasTypeNewRegistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewRegistrationAcceptMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeRegistrationAcceptMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationAcceptMessageIdentityTable = []nasTypeRegistrationAcceptMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeRegistrationAcceptMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewRegistrationAcceptMessageIdentity()
	for _, table := range nasTypeRegistrationAcceptMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type RegistrationAcceptMessageIdentityTestDataTemplate struct {
	in  nasType.RegistrationAcceptMessageIdentity
	out nasType.RegistrationAcceptMessageIdentity
}

var RegistrationAcceptMessageIdentityTestData = []nasType.RegistrationAcceptMessageIdentity{
	{0x03},
}

var RegistrationAcceptMessageIdentityExpectedTestData = []nasType.RegistrationAcceptMessageIdentity{
	{0x03},
}

var RegistrationAcceptMessageIdentityTable = []RegistrationAcceptMessageIdentityTestDataTemplate{
	{RegistrationAcceptMessageIdentityTestData[0], RegistrationAcceptMessageIdentityExpectedTestData[0]},
}

func TestNasTypeRegistrationAcceptMessageIdentity(t *testing.T) {
	for _, table := range RegistrationAcceptMessageIdentityTable {
		a := nasType.NewRegistrationAcceptMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
