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

func TestNasTypeNewLADNIndication(t *testing.T) {
	a := nasType.NewLADNIndication(nasMessage.RegistrationRequestLADNIndicationType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationRequestLADNIndicationTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestLADNIndicationType, nasMessage.RegistrationRequestLADNIndicationType},
}

func TestNasTypeLADNIndicationGetSetIei(t *testing.T) {
	a := nasType.NewLADNIndication(nasMessage.RegistrationRequestLADNIndicationType)
	for _, table := range nasTypeRegistrationRequestLADNIndicationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeLADNIndicationLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeLADNIndicationGetSetLen(t *testing.T) {
	a := nasType.NewLADNIndication(nasMessage.RegistrationRequestLADNIndicationType)
	for _, table := range nasTypeLADNIndicationLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeLADNIndicationLADNDNNValueData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeLADNIndicationLADNDNNValueTable = []nasTypeLADNIndicationLADNDNNValueData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeLADNIndicationGetSetLADNDNNValue(t *testing.T) {
	a := nasType.NewLADNIndication(nasMessage.RegistrationRequestLADNIndicationType)
	for _, table := range nasTypeLADNIndicationLADNDNNValueTable {
		a.SetLen(table.inLen)
		a.SetLADNDNNValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetLADNDNNValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLADNDNNValue())
		}
	}
}

type testLADNIndicationDataTemplate struct {
	inIei           uint8
	inLen           uint16
	inLADNDNNValue  []uint8
	outIei          uint8
	outLen          uint16
	outLADNDNNValue []uint8
}

var testLADNIndicationTestTable = []testLADNIndicationDataTemplate{
	{
		nasMessage.RegistrationRequestLADNIndicationType, 2,
		[]uint8{0xff, 0xff},
		nasMessage.RegistrationRequestLADNIndicationType, 2,
		[]uint8{0xff, 0xff},
	},
}

func TestNasTypeLADNIndication(t *testing.T) {
	for i, table := range testLADNIndicationTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewLADNIndication(nasMessage.RegistrationRequestLADNIndicationType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetLADNDNNValue(table.inLADNDNNValue)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outLADNDNNValue, a.GetLADNDNNValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.inLADNDNNValue, table.outLADNDNNValue, a.GetLADNDNNValue())
		}
	}
}
