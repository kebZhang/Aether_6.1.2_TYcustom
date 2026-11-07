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

func TestNasTypeNewAuthenticationFailureParameter(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultAuthenticationFailureParameterTable = []NasTypeIeiData{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, nasMessage.AuthenticationFailureAuthenticationFailureParameterType},
}

func TestNasTypeAuthenticationFailureParameterGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeAuthenticationResultAuthenticationFailureParameterTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthenticationResultAuthenticationFailureParameterLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAuthenticationFailureParameterGetSetLen(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeAuthenticationResultAuthenticationFailureParameterLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeOctetData struct {
	inLen uint8
	in    [14]uint8
	out   [14]uint8
}

var nasTypeOctetTable = []nasTypeOctetData{
	{14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationFailureParameterGetSetOctet(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeOctetTable {
		a.SetLen(table.inLen)
		a.SetAuthenticationFailureParameter(table.in)
		if !reflect.DeepEqual(table.out, a.GetAuthenticationFailureParameter()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAuthenticationFailureParameter())
		}
	}
}

type testAuthenticationFailureParameterDataTemplate struct {
	in  nasType.AuthenticationFailureParameter
	out nasType.AuthenticationFailureParameter
}

var authenticationFailureParameterTestData = []nasType.AuthenticationFailureParameter{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, 14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationFailureParameterExpectedTestData = []nasType.AuthenticationFailureParameter{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, 14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationFailureParameterTestTable = []testAuthenticationFailureParameterDataTemplate{
	{authenticationFailureParameterTestData[0], authenticationFailureParameterExpectedTestData[0]},
}

func TestNasTypeAuthenticationFailureParameter(t *testing.T) {
	for i, table := range authenticationFailureParameterTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetAuthenticationFailureParameter(table.in.Octet)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
