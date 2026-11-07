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

func TestNasTypeNewRQTimerValue(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionReleaseCompleteRQTimerValueTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType},
}

func TestNasTypeRQTimerValueGetSetIei(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypePDUSessionReleaseCompleteRQTimerValueTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeRQTimerValueUintTable = []NasTypeLenuint8Data{
	{0x01, 0x01},
}

func TestNasTypeRQTimerValueGetSetUint(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypeRQTimerValueUintTable {
		a.SetUnit(table.in)
		if !reflect.DeepEqual(table.out, a.GetUnit()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUnit())
		}
	}
}

type nasTypeRQTimerValueTimerValueData struct {
	in  uint8
	out uint8
}

var nasTypeRQTimerValueTimerValueTable = []nasTypeRQTimerValueTimerValueData{
	{0x01, 0x01},
}

func TestNasTypeRQTimerValueGetSetTimerValue(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypeRQTimerValueTimerValueTable {
		a.SetTimerValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetTimerValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTimerValue())
		}
	}
}

type testRQTimerValueDataTemplate struct {
	inUnit       uint8
	inTimerValue uint8
	in           nasType.RQTimerValue
	out          nasType.RQTimerValue
}

var rQTimerValueTestData = []nasType.RQTimerValue{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, 0x01},
}

var rQTimerValueExpectedTestData = []nasType.RQTimerValue{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, 0x21},
}

var rQTimerValueTestTable = []testRQTimerValueDataTemplate{
	{0x01, 0x01, rQTimerValueTestData[0], rQTimerValueExpectedTestData[0]},
}

func TestNasTypeRQTimerValue(t *testing.T) {
	for _, table := range rQTimerValueTestTable {
		a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)

		a.SetIei(table.in.GetIei())
		a.SetUnit(table.inUnit)
		a.SetTimerValue(table.inTimerValue)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
