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

func TestNasTypeNewAuthenticationFailureMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationFailureMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeMessageType struct {
	in  uint8
	out uint8
}

var nasTypeMessageTypeTable = []nasTypeMessageType{
	{0x03, 0x03},
}

func TestNasTypeGetSetMessageType(t *testing.T) {
	a := nasType.NewAuthenticationFailureMessageIdentity()
	for _, table := range nasTypeMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type AuthenticationFailureMessageIdentityTestDataTemplate struct {
	in  nasType.AuthenticationFailureMessageIdentity
	out nasType.AuthenticationFailureMessageIdentity
}

var authenticationFailureMessageIdentityTestData = []nasType.AuthenticationFailureMessageIdentity{
	{0x03},
}

var authenticationFailureMessageIdentityExpectedTestData = []nasType.AuthenticationFailureMessageIdentity{
	{0x03},
}

var authenticationFailureMessageIdentityTable = []AuthenticationFailureMessageIdentityTestDataTemplate{
	{authenticationFailureMessageIdentityTestData[0], authenticationFailureMessageIdentityExpectedTestData[0]},
}

func TestNasTypeAuthenticationFailureMessageIdentity(t *testing.T) {
	for i, table := range authenticationFailureMessageIdentityTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationFailureMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
