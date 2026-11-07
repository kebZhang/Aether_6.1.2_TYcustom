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

func TestNasTypeNewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentRequest, nas.MsgTypePDUSessionEstablishmentRequest},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
