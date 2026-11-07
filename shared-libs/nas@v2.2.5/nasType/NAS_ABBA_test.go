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

func TestNasTypeNewABBA(t *testing.T) {
	a := nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultABBATable = []NasTypeIeiData{
	{nasMessage.AuthenticationResultABBAType, nasMessage.AuthenticationResultABBAType},
}

func TestNasTypeABBAGetSetIei(t *testing.T) {
	a := nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
	for _, table := range nasTypeAuthenticationResultABBATable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}

		// if a.GetIei() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetIei())
		// }
	}
}

var nasTypeAuthenticationResultABBALenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeABBAGetSetLen(t *testing.T) {
	a := nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
	for _, table := range nasTypeAuthenticationResultABBALenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
		// if a.GetLen() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetLen())
		// }
	}
}

type nasTypeContentData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeContentTable = []nasTypeContentData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeABBAGetSetContent(t *testing.T) {
	a := nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
	for _, table := range nasTypeContentTable {
		a.SetLen(table.inLen)
		a.SetABBAContents(table.in)
		if !reflect.DeepEqual(table.out, a.GetABBAContents()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetABBAContents())
		}
	}
}

type testABBADataTemplate struct {
	in  nasType.ABBA
	out nasType.ABBA
}

var aBBATestData = []nasType.ABBA{
	{nasMessage.AuthenticationResultABBAType, 2, []byte{0x00, 0x00}},
}

var aBBAExpectedTestData = []nasType.ABBA{
	{nasMessage.AuthenticationResultABBAType, 2, []byte{0x00, 0x00}},
}

var aBBATestTable = []testABBADataTemplate{
	{aBBATestData[0], aBBAExpectedTestData[0]},
}

func TestNasTypeABBA(t *testing.T) {
	for i, table := range aBBATestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewABBA(nasMessage.AuthenticationResultABBAType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetABBAContents(table.in.Buffer)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Buffer, a.Buffer) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
		}
	}
}
