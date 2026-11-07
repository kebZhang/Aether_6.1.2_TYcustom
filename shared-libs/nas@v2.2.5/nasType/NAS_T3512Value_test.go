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

func TestNasTypeNewT3512Value(t *testing.T) {
	a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeT3512ValueGetSetIei(t *testing.T) {
	a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

func TestNasTypeT3512ValueGetSetLen(t *testing.T) {
	a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeT3512ValueUintValueData struct {
	in  uint8
	out uint8
}

var nasTypeT3512ValueUintValueTable = []NasTypeT3512ValueUintValueData{
	{0x1, 0x01},
}

func TestNasTypeT3512ValueGetSetUintValue(t *testing.T) {
	a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeT3512ValueUintValueTable {
		a.SetUnit(table.in)
		if !reflect.DeepEqual(table.out, a.GetUnit()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUnit())
		}
	}
}

type NasTypeT3512ValueTimerValueData struct {
	in  uint8
	out uint8
}

var nasTypeT3512ValueTimerValueTable = []NasTypeT3512ValueTimerValueData{
	{0x1, 0x01},
}

func TestNasTypeT3512ValueGetSetTimerValue(t *testing.T) {
	a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeT3512ValueTimerValueTable {
		a.SetTimerValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetTimerValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTimerValue())
		}
	}
}

type testT3512ValueDataTemplate struct {
	in  nasType.T3512Value
	out nasType.T3512Value
}

var T3512ValueTestData = []nasType.T3512Value{
	{nasMessage.RegistrationAcceptT3512ValueType, 1, 0x00},
}

var T3512ValueExpectedData = []nasType.T3512Value{
	{nasMessage.RegistrationAcceptT3512ValueType, 1, 0x21},
}

var T3512ValueDataTestTable = []testT3512ValueDataTemplate{
	{T3512ValueTestData[0], T3512ValueExpectedData[0]},
}

func TestNasTypeT3512Value(t *testing.T) {
	for _, table := range T3512ValueDataTestTable {
		a := nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetUnit(0x01)
		a.SetTimerValue(0x01)
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
