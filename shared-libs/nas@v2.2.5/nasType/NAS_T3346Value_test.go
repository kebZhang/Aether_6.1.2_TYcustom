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

func TestNasTypeNewT3346Value(t *testing.T) {
	a := nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeT3346ValueGetSetIei(t *testing.T) {
	a := nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

func TestNasTypeT3346ValueGetSetLen(t *testing.T) {
	a := nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeT3346ValueGPRSTimer2ValueData struct {
	in  uint8
	out uint8
}

var nasTypeT3346ValueGPRSTimer2ValueTable = []NasTypeT3346ValueGPRSTimer2ValueData{
	{0x2, 0x2},
}

func TestNasTypeT3346ValueGetSetGPRSTimer2Value(t *testing.T) {
	a := nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
	for _, table := range nasTypeT3346ValueGPRSTimer2ValueTable {
		a.SetGPRSTimer2Value(table.in)
		if !reflect.DeepEqual(table.out, a.GetGPRSTimer2Value()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetGPRSTimer2Value())
		}
	}
}

type testT3346ValueDataTemplate struct {
	in  nasType.T3346Value
	out nasType.T3346Value
}

var T3346ValueTestData = []nasType.T3346Value{
	{nasMessage.RegistrationRejectT3346ValueType, 1, 0x05},
}

var T3346ValueExpectedData = []nasType.T3346Value{
	{nasMessage.RegistrationRejectT3346ValueType, 1, 0x05},
}

var T3346ValueDataTestTable = []testT3346ValueDataTemplate{
	{T3346ValueTestData[0], T3346ValueExpectedData[0]},
}

func TestNasTypeT3346Value(t *testing.T) {
	for _, table := range T3346ValueDataTestTable {
		a := nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetGPRSTimer2Value(0x05)
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
