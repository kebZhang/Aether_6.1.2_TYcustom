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

func TestNasTypeNewRequestedDRXParameters(t *testing.T) {
	a := nasType.NewRequestedDRXParameters(nasMessage.RegistrationRequestRequestedDRXParametersType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRequestedDRXParametersServiceRejectT3346ValueTypeTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestRequestedDRXParametersType, nasMessage.RegistrationRequestRequestedDRXParametersType},
}

func TestNasTypeRequestedDRXParametersGetSetIei(t *testing.T) {
	a := nasType.NewRequestedDRXParameters(nasMessage.RegistrationRequestRequestedDRXParametersType)
	for _, table := range nasTypeRequestedDRXParametersServiceRejectT3346ValueTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeRequestedDRXParametersLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeRequestedDRXParametersGetSetLen(t *testing.T) {
	a := nasType.NewRequestedDRXParameters(nasMessage.RegistrationRequestRequestedDRXParametersType)
	for _, table := range nasTypeRequestedDRXParametersLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeRequestedDRXParametersDRXValueData struct {
	in  uint8
	out uint8
}

var nasTypeRequestedDRXParametersDRXValueTable = []nasTypeRequestedDRXParametersDRXValueData{
	{0x0f, 0x0f},
}

func TestNasTypeRequestedDRXParametersGetSetGPRSTimer2Value(t *testing.T) {
	a := nasType.NewRequestedDRXParameters(nasMessage.RegistrationRequestRequestedDRXParametersType)
	for _, table := range nasTypeRequestedDRXParametersDRXValueTable {
		a.SetDRXValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetDRXValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetDRXValue())
		}
	}
}

type testRequestedDRXParametersDataTemplate struct {
	inIei       uint8
	inLen       uint8
	inDRXValue  uint8
	outIei      uint8
	outLen      uint8
	outDRXValue uint8
}

var testRequestedDRXParametersTestTable = []testRequestedDRXParametersDataTemplate{
	{
		nasMessage.RegistrationRequestRequestedDRXParametersType, 2, 0x0f,
		nasMessage.RegistrationRequestRequestedDRXParametersType, 2, 0x0f,
	},
}

func TestNasTypeRequestedDRXParameters(t *testing.T) {
	for i, table := range testRequestedDRXParametersTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRequestedDRXParameters(nasMessage.RegistrationRequestRequestedDRXParametersType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetDRXValue(table.inDRXValue)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outDRXValue, a.GetDRXValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.inDRXValue, table.outDRXValue, a.GetDRXValue())
		}
	}
}
