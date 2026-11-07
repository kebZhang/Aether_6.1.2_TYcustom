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

func TestNasTypeNewPduSessionID2Value(t *testing.T) {
	a := nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionIDULNASTransportPduSessionID2ValueTypeTypeTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportPduSessionID2ValueType, nasMessage.ULNASTransportPduSessionID2ValueType},
}

func TestNasTypePduSessionID2ValueGetSetIei(t *testing.T) {
	a := nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
	for _, table := range nasTypePDUSessionIDULNASTransportPduSessionID2ValueTypeTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypePDUSessionIDPduSessionID2ValueData struct {
	in  uint8
	out uint8
}

var nasTypePduSessionIdentity2ValueTable = []nasTypePDUSessionIDPduSessionID2ValueData{
	{0xff, 0xff},
}

func TestNasTypeGetSetPduSessionIdentity2Value(t *testing.T) {
	a := nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
	for _, table := range nasTypePduSessionIdentity2ValueTable {
		a.SetPduSessionID2Value(table.in)
		if !reflect.DeepEqual(table.out, a.GetPduSessionID2Value()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPduSessionID2Value())
		}
	}
}

type testPduSessionIdentity2ValueDataTemplate struct {
	inIei                       uint8
	inPduSessionIdentity2Value  uint8
	outIei                      uint8
	outPduSessionIdentity2Value uint8
}

var testPduSessionIdentity2ValueTestTable = []testPduSessionIdentity2ValueDataTemplate{
	{
		nasMessage.ULNASTransportPduSessionID2ValueType, 0x0f,
		nasMessage.ULNASTransportPduSessionID2ValueType, 0x0f,
	},
}

func TestNasTypePDUSessionID2Value(t *testing.T) {
	for i, table := range testPduSessionIdentity2ValueTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
		a.SetIei(table.inIei)
		a.SetPduSessionID2Value(table.inPduSessionIdentity2Value)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outPduSessionIdentity2Value, a.GetPduSessionID2Value()) {
			t.Errorf("in(%v): out %v, actual %x", table.inPduSessionIdentity2Value, table.outPduSessionIdentity2Value, a.GetPduSessionID2Value())
		}
	}
}
