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

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationCommand, nas.MsgTypePDUSessionModificationCommand},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
