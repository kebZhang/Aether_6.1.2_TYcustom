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

func TestNasTypeNewSecurityModeCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewSecurityModeCompleteMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeSecurityModeCompleteMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeSecurityModeCompleteMessageIdentityTable = []nasTypeSecurityModeCompleteMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeSecurityModeCompleteMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewSecurityModeCompleteMessageIdentity()
	for _, table := range nasTypeSecurityModeCompleteMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type SecurityModeCompleteMessageIdentityTestDataTemplate struct {
	in  nasType.SecurityModeCompleteMessageIdentity
	out nasType.SecurityModeCompleteMessageIdentity
}

var SecurityModeCompleteMessageIdentityTestData = []nasType.SecurityModeCompleteMessageIdentity{
	{0x03},
}

var SecurityModeCompleteMessageIdentityExpectedTestData = []nasType.SecurityModeCompleteMessageIdentity{
	{0x03},
}

var SecurityModeCompleteMessageIdentityTable = []SecurityModeCompleteMessageIdentityTestDataTemplate{
	{SecurityModeCompleteMessageIdentityTestData[0], SecurityModeCompleteMessageIdentityExpectedTestData[0]},
}

func TestNasTypeSecurityModeCompleteMessageIdentity(t *testing.T) {
	for _, table := range SecurityModeCompleteMessageIdentityTable {
		a := nasType.NewSecurityModeCompleteMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
