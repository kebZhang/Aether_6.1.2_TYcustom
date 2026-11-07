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

func TestNasTypeNewServiceAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewServiceAcceptMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeServiceAcceptMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeServiceAcceptMessageIdentityTable = []nasTypeServiceAcceptMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeServiceAcceptMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewServiceAcceptMessageIdentity()
	for _, table := range nasTypeServiceAcceptMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type ServiceAcceptMessageIdentityTestDataTemplate struct {
	in  nasType.ServiceAcceptMessageIdentity
	out nasType.ServiceAcceptMessageIdentity
}

var ServiceAcceptMessageIdentityTestData = []nasType.ServiceAcceptMessageIdentity{
	{0x03},
}

var ServiceAcceptMessageIdentityExpectedTestData = []nasType.ServiceAcceptMessageIdentity{
	{0x03},
}

var ServiceAcceptMessageIdentityTable = []ServiceAcceptMessageIdentityTestDataTemplate{
	{ServiceAcceptMessageIdentityTestData[0], ServiceAcceptMessageIdentityExpectedTestData[0]},
}

func TestNasTypeServiceAcceptMessageIdentity(t *testing.T) {
	for _, table := range ServiceAcceptMessageIdentityTable {
		a := nasType.NewServiceAcceptMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
