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

func TestNasTypeNewPDUAddress(t *testing.T) {
	a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUAddressPDUSessionEstablishmentAcceptPDUAddressTypeTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentAcceptPDUAddressType, nasMessage.PDUSessionEstablishmentAcceptPDUAddressType},
}

func TestNasTypePDUAddressGetSetIei(t *testing.T) {
	a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
	for _, table := range nasTypePDUAddressPDUSessionEstablishmentAcceptPDUAddressTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypePDUAddressLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypePDUAddressGetSetLen(t *testing.T) {
	a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
	for _, table := range nasTypePDUAddressLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypePDUAddressPDUSessionTypeValueData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypePDUAddressPDUSessionTypeValueTable = []nasTypePDUAddressPDUSessionTypeValueData{
	{2, 0xff, 0x07},
}

func TestNasTypePDUAddressGetSetPDUSessionTypeValue(t *testing.T) {
	a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
	for _, table := range nasTypePDUAddressPDUSessionTypeValueTable {
		a.SetLen(table.inLen)
		a.SetPDUSessionTypeValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUSessionTypeValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPDUSessionTypeValue())
		}
	}
}

type nasTypePDUAddressPDUAddressInformationData struct {
	inLen uint8
	in    [12]uint8
	out   [12]uint8
}

var nasTypePDUAddressPDUAddressInformationTable = []nasTypePDUAddressPDUAddressInformationData{
	{12, [12]uint8{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}, [12]uint8{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}},
}

func TestNasTypePDUAddressGetSetPDUAddressInformation(t *testing.T) {
	a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
	for _, table := range nasTypePDUAddressPDUAddressInformationTable {
		a.SetLen(table.inLen)
		a.SetPDUAddressInformation(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUAddressInformation()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPDUAddressInformation())
		}
	}
}

type testPDUAddressDataTemplate struct {
	inIei                    uint8
	inLen                    uint8
	inPDUSessionTypeValue    uint8
	inPDUAddressInformation  [12]uint8
	outIei                   uint8
	outLen                   uint8
	outPDUSessionTypeValue   uint8
	outPDUAddressInformation [12]uint8
}

var testPDUAddressTestTable = []testPDUAddressDataTemplate{
	{
		nasMessage.PDUSessionEstablishmentAcceptPDUAddressType, 12, 0x07,
		[12]uint8{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f},
		nasMessage.PDUSessionEstablishmentAcceptPDUAddressType, 12, 0x07,
		[12]uint8{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f},
	},
}

func TestNasTypePDUAddress(t *testing.T) {
	for i, table := range testPDUAddressTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetPDUSessionTypeValue(table.inPDUSessionTypeValue)
		a.SetPDUAddressInformation(table.inPDUAddressInformation)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outPDUSessionTypeValue, a.GetPDUSessionTypeValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.inPDUSessionTypeValue, table.outPDUSessionTypeValue, a.GetPDUSessionTypeValue())
		}
		if !reflect.DeepEqual(table.outPDUAddressInformation, a.GetPDUAddressInformation()) {
			t.Errorf("in(%v): out %v, actual %x", table.inPDUAddressInformation, table.outPDUAddressInformation, a.GetPDUAddressInformation())
		}
	}
}
