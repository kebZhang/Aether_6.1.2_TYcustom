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

func TestNasTypeNewAuthenticationParameterAUTN(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultAuthenticationParameterAUTNTable = []NasTypeIeiData{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, nasMessage.AuthenticationRequestAuthenticationParameterAUTNType},
}

func TestNasTypeAuthenticationParameterAUTNGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationResultAuthenticationParameterAUTNTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthenticationResultAuthenticationParameterAUTNLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAuthenticationParameterAUTNGetSetLen(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationResultAuthenticationParameterAUTNLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeAuthenticationParameterAUTNOctetData struct {
	inLen uint8
	in    [16]uint8
	out   [16]uint8
}

var nasTypeAuthenticationParameterAUTNOctetTable = []nasTypeAuthenticationParameterAUTNOctetData{
	{16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationParameterAUTNGetSetAUTN(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationParameterAUTNOctetTable {
		a.SetLen(table.inLen)
		a.SetAUTN(table.in)
		if !reflect.DeepEqual(table.out, a.GetAUTN()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAUTN())
		}
	}
}

type testAuthenticationParameterAUTNDataTemplate struct {
	in  nasType.AuthenticationParameterAUTN
	out nasType.AuthenticationParameterAUTN
}

var authenticationParameterAUTNTestData = []nasType.AuthenticationParameterAUTN{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterAUTNExpectedTestData = []nasType.AuthenticationParameterAUTN{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterAUTNTestTable = []testAuthenticationParameterAUTNDataTemplate{
	{authenticationParameterAUTNTestData[0], authenticationParameterAUTNExpectedTestData[0]},
}

func TestNasTypeAuthenticationParameterAUTN(t *testing.T) {
	for i, table := range authenticationParameterAUTNTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetAUTN(table.in.Octet)

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
