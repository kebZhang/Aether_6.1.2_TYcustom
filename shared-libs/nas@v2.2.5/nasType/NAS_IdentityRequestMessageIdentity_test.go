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

type nasTypeIdentityRequestMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityRequestMessageIdentityTable = []nasTypeIdentityRequestMessageIdentityData{
	{nas.MsgTypeIdentityRequest, nas.MsgTypeIdentityRequest},
}

func TestNasTypeNewIdentityRequestMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityRequestMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeIdentityRequestMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityRequestMessageIdentity()
	for _, table := range nasTypeIdentityRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
