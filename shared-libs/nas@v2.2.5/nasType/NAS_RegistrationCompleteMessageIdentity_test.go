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

func TestNasTypeNewRegistrationCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewRegistrationCompleteMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeRegistrationCompleteMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationCompleteMessageIdentityTable = []nasTypeRegistrationCompleteMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeRegistrationCompleteMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewRegistrationCompleteMessageIdentity()
	for _, table := range nasTypeRegistrationCompleteMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type RegistrationCompleteMessageIdentityTestDataTemplate struct {
	in  nasType.RegistrationCompleteMessageIdentity
	out nasType.RegistrationCompleteMessageIdentity
}

var RegistrationCompleteMessageIdentityTestData = []nasType.RegistrationCompleteMessageIdentity{
	{0x03},
}

var RegistrationCompleteMessageIdentityExpectedTestData = []nasType.RegistrationCompleteMessageIdentity{
	{0x03},
}

var RegistrationCompleteMessageIdentityTable = []RegistrationCompleteMessageIdentityTestDataTemplate{
	{RegistrationCompleteMessageIdentityTestData[0], RegistrationCompleteMessageIdentityExpectedTestData[0]},
}

func TestNasTypeRegistrationCompleteMessageIdentity(t *testing.T) {
	for _, table := range RegistrationCompleteMessageIdentityTable {
		a := nasType.NewRegistrationCompleteMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
