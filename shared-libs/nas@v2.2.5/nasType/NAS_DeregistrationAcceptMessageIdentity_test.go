// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/nasType"
)

type nasTypeDeregistrationAcceptMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationAcceptMessageIdentityTable = []nasTypeDeregistrationAcceptMessageIdentityData{
	{nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration, nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration},
}

func TestNasTypeNewDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	for _, table := range nasTypeDeregistrationAcceptMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
