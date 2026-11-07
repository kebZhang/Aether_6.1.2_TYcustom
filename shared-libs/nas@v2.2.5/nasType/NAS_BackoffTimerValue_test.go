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

func TestNasTypeNewBackoffTimerValue(t *testing.T) {
	a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationRequestBackoffTimerValueIeiTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptT3512ValueType, nasMessage.RegistrationAcceptT3512ValueType},
}

func TestNasTypeBackoffTimerValueGetSetIei(t *testing.T) {
	a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeAuthenticationRequestBackoffTimerValueIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeBackoffTimerValueLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeBackoffTimerValueGetSetLen(t *testing.T) {
	a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeBackoffTimerValueLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeBackoffTimerValueUintTimerValue struct {
	in  uint8
	out uint8
}

var nasTypeBackoffTimerValueUintTimerValueTable = []nasTypeBackoffTimerValueUintTimerValue{
	{0x07, 0x07},
}

func TestNasTypeBackoffTimerValueGetSetUintTimerValue(t *testing.T) {
	a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeBackoffTimerValueUintTimerValueTable {
		a.SetUnitTimerValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetUnitTimerValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUnitTimerValue())
		}
	}
}

type nasTypeBackoffTimerValueTimerValue struct {
	in  uint8
	out uint8
}

var nasTypeBackoffTimerValueTimerValueTable = []nasTypeBackoffTimerValueTimerValue{
	{0x07, 0x07},
}

func TestNasTypeBackoffTimerValueGetSetTimerValue(t *testing.T) {
	a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
	for _, table := range nasTypeBackoffTimerValueTimerValueTable {
		a.SetTimerValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetTimerValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTimerValue())
		}
	}
}

type testBackoffTimerValueDataTemplate struct {
	inUnitTimerValue uint8
	inTimerValue     uint8
	in               nasType.BackoffTimerValue
	out              nasType.BackoffTimerValue
}

var BackoffTimerValueTestData = []nasType.BackoffTimerValue{
	{nasMessage.RegistrationAcceptT3512ValueType, 1, 0xff},
}

var BackoffTimerValueExpectedData = []nasType.BackoffTimerValue{
	{nasMessage.RegistrationAcceptT3512ValueType, 1, 0xff},
}

var BackoffTimerValueDataTestTable = []testBackoffTimerValueDataTemplate{
	{0x07, 0x1F, BackoffTimerValueTestData[0], BackoffTimerValueExpectedData[0]},
}

func TestNasTypeBackoffTimer(t *testing.T) {
	for _, table := range BackoffTimerValueDataTestTable {
		a := nasType.NewBackoffTimerValue(nasMessage.RegistrationAcceptT3512ValueType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetUnitTimerValue(table.inUnitTimerValue)
		a.SetTimerValue(table.inTimerValue)
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
