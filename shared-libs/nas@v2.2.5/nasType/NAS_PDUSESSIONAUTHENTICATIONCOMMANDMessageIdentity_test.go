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

func TestNasTypeNewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType{
	{nas.MsgTypePDUSessionAuthenticationCommand, nas.MsgTypePDUSessionAuthenticationCommand},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
