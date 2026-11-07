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

func TestNasTypeNewSecurityModeCommandMessageIdentity(t *testing.T) {
	a := nasType.NewSecurityModeCommandMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeSecurityModeCommandMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeSecurityModeCommandMessageIdentityTable = []nasTypeSecurityModeCommandMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeSecurityModeCommandMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewSecurityModeCommandMessageIdentity()
	for _, table := range nasTypeSecurityModeCommandMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type SecurityModeCommandMessageIdentityTestDataTemplate struct {
	in  nasType.SecurityModeCommandMessageIdentity
	out nasType.SecurityModeCommandMessageIdentity
}

var SecurityModeCommandMessageIdentityTestData = []nasType.SecurityModeCommandMessageIdentity{
	{0x03},
}

var SecurityModeCommandMessageIdentityExpectedTestData = []nasType.SecurityModeCommandMessageIdentity{
	{0x03},
}

var SecurityModeCommandMessageIdentityTable = []SecurityModeCommandMessageIdentityTestDataTemplate{
	{SecurityModeCommandMessageIdentityTestData[0], SecurityModeCommandMessageIdentityExpectedTestData[0]},
}

func TestNasTypeSecurityModeCommandMessageIdentity(t *testing.T) {
	for _, table := range SecurityModeCommandMessageIdentityTable {
		a := nasType.NewSecurityModeCommandMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
