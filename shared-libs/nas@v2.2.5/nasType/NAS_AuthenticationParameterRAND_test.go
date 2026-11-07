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

func TestNasTypeNewAuthenticationParameterRAND(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationRequestAuthenticationParameterRANDTable = []NasTypeIeiData{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, nasMessage.AuthenticationRequestAuthenticationParameterRANDType},
}

func TestNasTypeAuthenticationParameterRANDGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	for _, table := range nasTypeAuthenticationRequestAuthenticationParameterRANDTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeAuthenticationParameterRANDOctetData struct {
	in  [16]uint8
	out [16]uint8
}

var nasTypeAuthenticationParameterRANDOctetTable = []nasTypeAuthenticationParameterRANDOctetData{
	{[16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationParameterRANDGetSetRANDValue(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	for _, table := range nasTypeAuthenticationParameterRANDOctetTable {
		a.SetRANDValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetRANDValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRANDValue())
		}
	}
}

type testAuthenticationParameterRANDDataTemplate struct {
	in  nasType.AuthenticationParameterRAND
	out nasType.AuthenticationParameterRAND
}

var authenticationParameterRANDTestData = []nasType.AuthenticationParameterRAND{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterRANDExpectedTestData = []nasType.AuthenticationParameterRAND{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterRANDTestTable = []testAuthenticationParameterRANDDataTemplate{
	{authenticationParameterRANDTestData[0], authenticationParameterRANDExpectedTestData[0]},
}

func TestNasTypeAuthenticationParameterRAND(t *testing.T) {
	for i, table := range authenticationParameterRANDTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)

		a.SetIei(table.in.GetIei())
		a.SetRANDValue(table.in.Octet)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
