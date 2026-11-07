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

type nasTypeRejectMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeRejectMessageIdentityTable = []nasTypeRejectMessageIdentityData{
	{nasMessage.PDUSessionEstablishmentRejectEAPMessageType, nasMessage.PDUSessionEstablishmentRejectEAPMessageType},
}

func TestNasTypeNewAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	for _, table := range nasTypeRejectMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
