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

func TestNasTypeNewNotificationMessageIdentity(t *testing.T) {
	a := nasType.NewNotificationMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeNotificationMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypeNotificationMessageIdentityMessageTypeTable = []nasTypeNotificationMessageIdentityMessageType{
	{nas.MsgTypeNotification, nas.MsgTypeNotification},
}

func TestNasTypeGetSetNotificationMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewNotificationMessageIdentity()
	for _, table := range nasTypeNotificationMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
