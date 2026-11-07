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

func TestNasTypeNewLADNInformation(t *testing.T) {
	a := nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationRequestLADNInformationTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandLADNInformationType, nasMessage.ConfigurationUpdateCommandLADNInformationType},
}

func TestNasTypeLADNInformationGetSetIei(t *testing.T) {
	a := nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
	for _, table := range nasTypeRegistrationRequestLADNInformationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeLADNInformationLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeLADNInformationGetSetLen(t *testing.T) {
	a := nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
	for _, table := range nasTypeLADNInformationLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeLADNInformationLADNDData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeLADNInformationLADNDTable = []nasTypeLADNInformationLADNDData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeLADNInformationGetSetLADND(t *testing.T) {
	a := nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
	for _, table := range nasTypeLADNInformationLADNDTable {
		a.SetLen(table.inLen)
		a.SetLADND(table.in)
		if !reflect.DeepEqual(table.out, a.GetLADND()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLADND())
		}
	}
}

type testLADNInformationDataTemplate struct {
	inIei    uint8
	inLen    uint16
	inLADND  []uint8
	outIei   uint8
	outLen   uint16
	outLADND []uint8
}

var testLADNInformationTestTable = []testLADNInformationDataTemplate{
	{
		nasMessage.ConfigurationUpdateCommandLADNInformationType, 2,
		[]uint8{0xff, 0xff},
		nasMessage.ConfigurationUpdateCommandLADNInformationType, 2,
		[]uint8{0xff, 0xff},
	},
}

func TestNasTypeLADNInformation(t *testing.T) {
	for i, table := range testLADNInformationTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetLADND(table.inLADND)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outLADND, a.GetLADND()) {
			t.Errorf("in(%v): out %v, actual %x", table.inLADND, table.outLADND, a.GetLADND())
		}
	}
}
