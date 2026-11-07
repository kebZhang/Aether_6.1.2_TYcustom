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

func TestNasTypeNewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentAccept, nas.MsgTypePDUSessionEstablishmentAccept},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
