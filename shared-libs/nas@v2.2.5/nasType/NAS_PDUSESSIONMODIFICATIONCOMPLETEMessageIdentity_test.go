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

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationComplete, nas.MsgTypePDUSessionModificationComplete},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
