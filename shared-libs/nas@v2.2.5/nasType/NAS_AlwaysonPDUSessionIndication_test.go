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

var PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput uint8 = 0x08

func TestNasTypeNewAlwaysonPDUSessionIndication(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionIndicationTable = []NasTypeIeiData{
	{PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput, 0x08},
}

func TestNasTypeAlwaysonPDUSessionIndicationGetSetIei(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType)
	for _, table := range nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionIndicationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeAlwaysonPDUSessionIndicationAPSI struct {
	in  uint8
	out uint8
}

var nasTypeAlwaysonPDUSessionIndicationAPSITable = []nasTypeAlwaysonPDUSessionIndicationAPSI{
	{0x01, 0x01},
}

func TestNasTypeAlwaysonPDUSessionIndicationGetSetAPSI(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType)
	for _, table := range nasTypeAlwaysonPDUSessionIndicationAPSITable {
		a.SetAPSI(table.in)
		if !reflect.DeepEqual(table.out, a.GetAPSI()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAPSI())
		}
	}
}

type testAlwaysonPDUSessionIndicationDataTemplate struct {
	in  nasType.AlwaysonPDUSessionIndication
	out nasType.AlwaysonPDUSessionIndication
}

var alwaysonPDUSessionIndicationTestData = []nasType.AlwaysonPDUSessionIndication{
	{(0x80 + 0x01)},
}

var alwaysonPDUSessionIndicationExpectedTestData = []nasType.AlwaysonPDUSessionIndication{
	{(0x80 + 0x01)},
}

var alwaysonPDUSessionIndicationTestTable = []testAlwaysonPDUSessionIndicationDataTemplate{
	{alwaysonPDUSessionIndicationTestData[0], alwaysonPDUSessionIndicationExpectedTestData[0]},
}

func TestNasTypeAlwaysonPDUSessionIndication(t *testing.T) {
	for _, table := range alwaysonPDUSessionIndicationTestTable {
		a := nasType.NewAlwaysonPDUSessionIndication(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)

		a.SetIei(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)
		a.SetAPSI(table.in.GetAPSI())

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
