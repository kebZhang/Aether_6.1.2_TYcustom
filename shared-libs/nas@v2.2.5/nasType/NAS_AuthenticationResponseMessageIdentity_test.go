// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

type nasTypeResponseMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeResponseMessageIdentityTable = []nasTypeResponseMessageIdentityData{
	{nasMessage.AuthenticationResponseEAPMessageType, nasMessage.AuthenticationResponseEAPMessageType},
}

func TestNasTypeNewAuthenticationResponseMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationResponseMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetAuthenticationResponseMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationResponseMessageIdentity()
	for _, table := range nasTypeResponseMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
