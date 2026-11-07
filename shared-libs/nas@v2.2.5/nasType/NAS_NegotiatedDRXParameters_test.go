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

func TestNasTypeNewNegotiatedDRXParameters(t *testing.T) {
	a := nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeNegotiatedDRXParametersRegistrationRequestAdditionalGUTITable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptNegotiatedDRXParametersType, nasMessage.RegistrationAcceptNegotiatedDRXParametersType},
}

func TestNasTypeNegotiatedDRXParametersGetSetIei(t *testing.T) {
	a := nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
	for _, table := range nasTypeNegotiatedDRXParametersRegistrationRequestAdditionalGUTITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeNegotiatedDRXParametersLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeNegotiatedDRXParametersGetSetLen(t *testing.T) {
	a := nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
	for _, table := range nasTypeNegotiatedDRXParametersLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeNegotiatedDRXParametersDRXValueData struct {
	in  uint8
	out uint8
}

var nasTypeNegotiatedDRXParametersDRXValueTable = []nasTypeNegotiatedDRXParametersDRXValueData{
	{0x0f, 0x0f},
}

func TestNasTypeNegotiatedDRXParametersGetSetDRXValue(t *testing.T) {
	a := nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
	for _, table := range nasTypeNegotiatedDRXParametersDRXValueTable {
		a.SetDRXValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetDRXValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetDRXValue())
		}
	}
}

type testNegotiatedDRXParametersDataTemplate struct {
	inIei       uint8
	inLen       uint8
	inDRXValue  uint8
	outIei      uint8
	outLen      uint8
	outDRXValue uint8
}

var testNegotiatedDRXParametersTestTable = []testNegotiatedDRXParametersDataTemplate{
	{
		nasMessage.RegistrationAcceptNegotiatedDRXParametersType, 2, 0x0f,
		nasMessage.RegistrationAcceptNegotiatedDRXParametersType, 2, 0x0f,
	},
}

func TestNasTypeNegotiatedDRXParameters(t *testing.T) {
	for i, table := range testNegotiatedDRXParametersTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)

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
