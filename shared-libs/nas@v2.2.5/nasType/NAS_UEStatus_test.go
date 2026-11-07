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

func TestNasTypeNewUEStatus(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeUEStatusGetSetIei(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

func TestNasTypeUEStatusGetSetLen(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeUEStatusN1ModeRegData struct {
	in  uint8
	out uint8
}

var nasTypeUEStatusN1ModeRegTable = []NasTypeUEStatusN1ModeRegData{
	{0x01, 0x01},
}

func TestNasTypeUEStatusGetSetN1ModeReg(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeUEStatusN1ModeRegTable {
		a.SetN1ModeReg(table.in)
		if !reflect.DeepEqual(table.out, a.GetN1ModeReg()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetN1ModeReg())
		}
	}
}

type NasTypeUEStatusS1ModeRegData struct {
	in  uint8
	out uint8
}

var nasTypeUEStatusS1ModeRegTable = []NasTypeUEStatusS1ModeRegData{
	{0x01, 0x01},
}

func TestNasTypeUEStatusGetSetS1ModeReg(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeUEStatusS1ModeRegTable {
		a.SetS1ModeReg(table.in)
		if !reflect.DeepEqual(table.out, a.GetS1ModeReg()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetS1ModeReg())
		}
	}
}

type testUEStatusDataTemplate struct {
	in  nasType.UEStatus
	out nasType.UEStatus
}

var UEStatusTestData = []nasType.UEStatus{
	{nasMessage.RegistrationRequestUEStatusType, 1, 0x05},
}

var UEStatusExpectedData = []nasType.UEStatus{
	{nasMessage.RegistrationRequestUEStatusType, 1, 0x03},
}

var UEStatusDataTestTable = []testUEStatusDataTemplate{
	{UEStatusTestData[0], UEStatusExpectedData[0]},
}

func TestNasTypeUEStatus(t *testing.T) {
	for _, table := range UEStatusDataTestTable {
		a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetN1ModeReg(0x01)
		a.SetS1ModeReg(0x01)
		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
