// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewEAPMessage(t *testing.T) {
	a := nasType.NewEAPMessage(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationRequestEAPMessageIeiTable = []NasTypeIeiData{
	{0, 0},
}

func TestNasTypeEAPMessageGetSetIei(t *testing.T) {
	a := nasType.NewEAPMessage(0)
	for _, table := range nasTypeAuthenticationRequestEAPMessageIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeEAPMessageLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeEAPMessageGetSetLen(t *testing.T) {
	a := nasType.NewEAPMessage(0)
	for _, table := range nasTypeEAPMessageLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypetEAPMessageData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeEAPMessageTable = []nasTypetEAPMessageData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeEAPMessageGetSetEAPMessage(t *testing.T) {
	a := nasType.NewEAPMessage(0)
	for _, table := range nasTypeEAPMessageTable {
		a.SetLen(table.inLen)
		a.SetEAPMessage(table.in)
		if !reflect.DeepEqual(table.out, a.GetEAPMessage()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetEAPMessage())
		}
	}
}

type testEAPDataTemplate struct {
	in  nasType.EAPMessage
	out nasType.EAPMessage
}

var EAPMessageTestData = []nasType.EAPMessage{
	{0, 2, []byte{0x00, 0x00}}, // AuthenticationResult
}

var EAPMessageExpectedTestData = []nasType.EAPMessage{
	{0, 2, []byte{0x00, 0x00}}, // AuthenticationResult
}

var EAPMessageTestTable = []testEAPDataTemplate{
	{EAPMessageTestData[0], EAPMessageExpectedTestData[0]},
}

func TestNasTypeEAPMessage(t *testing.T) {
	for i, table := range EAPMessageTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewEAPMessage(0) // AuthenticationResult

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetEAPMessage(table.in.Buffer)

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
