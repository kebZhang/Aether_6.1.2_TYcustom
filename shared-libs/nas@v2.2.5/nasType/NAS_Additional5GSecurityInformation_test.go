// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewAdditional5GSecurityInformation(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36) // security mode command message
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSecurityModeCommandAdditional5GSecurityInformationTable = []NasTypeIeiData{
	{0x36, 0x36},
}

func TestNasTypeAdditional5GSecurityInformationGetSetIei(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeSecurityModeCommandAdditional5GSecurityInformationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeSecurityModeCommandAdditional5GSecurityInformationLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAdditional5GSecurityInformationGetSetLen(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeSecurityModeCommandAdditional5GSecurityInformationLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeRINMRuint8Data struct {
	in  uint8
	out uint8
}

var nasTypeAdditional5GSecurityInformationRINMR = []NasTypeRINMRuint8Data{
	{0x1, 0x1},
	{0x0, 0x0},
}

func TestNasTypeAdditional5GSecurityInformationGetSetRINMR(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeAdditional5GSecurityInformationRINMR {
		a.SetRINMR(table.in)
		if !reflect.DeepEqual(table.out, a.GetRINMR()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRINMR())
		}
	}
}

type NasTypeHDPuint8Data struct {
	in  uint8
	out uint8
}

var nasTypeAdditional5GSecurityInformationHDP = []NasTypeHDPuint8Data{
	{0x1, 0x1},
	{0x0, 0x0},
}

func TestNasTypeAdditional5GSecurityInformationGetSetHDP(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeAdditional5GSecurityInformationHDP {
		a.SetHDP(table.in)
		if !reflect.DeepEqual(table.out, a.GetHDP()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetHDP())
		}
	}
}
