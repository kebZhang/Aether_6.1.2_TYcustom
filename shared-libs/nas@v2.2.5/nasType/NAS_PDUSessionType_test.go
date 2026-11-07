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

var PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput uint8 = 0x09

func TestNasTypeNewPDUSessionType(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentRequestPDUSessionTypeTable = []NasTypeIeiData{
	{PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput, 0x09},
}

func TestNasTypePDUSessionTypeGetSetIei(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestPDUSessionTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypePDUSessionEstablishmentRequestPDUSessionTypeSpareTable = []NasTypeLenuint8Data{
	{0x1, 0x1},
}

func TestNasTypePDUSessionTypeGetSetSpare(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestPDUSessionTypeSpareTable {
		a.SetSpare(table.in)
		if !reflect.DeepEqual(table.out, a.GetSpare()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSpare())
		}
	}
}

var nasTypePDUSessionTypeValue = []NasTypeLenuint8Data{
	{0x0, 0x0},
	{0x1, 0x1},
	{0x2, 0x2},
	{0x3, 0x3},
	{0x4, 0x4},
	{0x5, 0x5},
}

func TestNasTypePDUSessionTypeGetSetPDUSessionTypeValue(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionTypeValue {
		a.SetPDUSessionTypeValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUSessionTypeValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPDUSessionTypeValue())
		}
	}
}

type testPDUSessionTypeDataTemplate struct {
	inPDUSessionTypeValue uint8
	in                    nasType.PDUSessionType
	out                   nasType.PDUSessionType
}

var pDUSessionTypeTestData = []nasType.PDUSessionType{
	{nasMessage.PDUSessionEstablishmentRequestPDUSessionTypeType},
}

var pDUSessionTypeExpectedData = []nasType.PDUSessionType{
	{(0x90 + 0x01)},
}

var pDUSessionTypeTestTable = []testPDUSessionTypeDataTemplate{
	{0x01, pDUSessionTypeTestData[0], pDUSessionTypeExpectedData[0]},
}

func TestNasTypePDUSessionType(t *testing.T) {
	for _, table := range pDUSessionTypeTestTable {
		a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)

		a.SetIei(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
		a.SetPDUSessionTypeValue(table.inPDUSessionTypeValue)

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
