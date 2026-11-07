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

func TestNasTypeNewPDUSessionID(t *testing.T) {
	a := nasType.NewPDUSessionID()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionIDULNASTransportOldPDUSessionIDTypeTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportOldPDUSessionIDType, nasMessage.ULNASTransportOldPDUSessionIDType},
}

func TestNasTypePDUSessionIDGetSetIei(t *testing.T) {
	a := nasType.NewPDUSessionID()
	for _, table := range nasTypePDUSessionIDULNASTransportOldPDUSessionIDTypeTable {
		a.SetPDUSessionID(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUSessionID()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPDUSessionID())
		}
	}
}

type nasTypePDUSessionIDPduSessionIdentity2ValueData struct {
	in  uint8
	out uint8
}

var nasTypePDUSessionIDPduSessionIdentity2ValueTable = []nasTypePDUSessionIDPduSessionIdentity2ValueData{
	{0xff, 0xff},
}

func TestNasTypePDUSessionIDGetSetPduSessionIdentity2Value(t *testing.T) {
	a := nasType.NewPDUSessionID()
	for _, table := range nasTypePDUSessionIDPduSessionIdentity2ValueTable {
		a.SetPDUSessionID(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUSessionID()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPDUSessionID())
		}
	}
}

type testPDUSessionIDDataTemplate struct {
	inPduSessionIdentity2Value  uint8
	outPduSessionIdentity2Value uint8
}

var testPDUSessionIDTestTable = []testPDUSessionIDDataTemplate{
	{0x0f, 0x0f},
}

func TestNasTypePDUSessionID(t *testing.T) {
	for i, table := range testPDUSessionIDTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPDUSessionID()
		a.SetPDUSessionID(table.inPduSessionIdentity2Value)

		if !reflect.DeepEqual(table.outPduSessionIdentity2Value, a.GetPDUSessionID()) {
			t.Errorf("in(%v): out %v, actual %x", table.inPduSessionIdentity2Value, table.outPduSessionIdentity2Value, a.GetPDUSessionID())
		}
	}
}
