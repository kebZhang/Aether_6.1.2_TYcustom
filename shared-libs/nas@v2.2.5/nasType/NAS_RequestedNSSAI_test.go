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

func TestNasTypeNewRequestedNSSAI(t *testing.T) {
	a := nasType.NewRequestedNSSAI(nasMessage.RegistrationRequestRequestedNSSAIType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultRequestedNSSAITable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestRequestedNSSAIType, nasMessage.RegistrationRequestRequestedNSSAIType},
}

func TestNasTypeRequestedNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewRequestedNSSAI(nasMessage.RegistrationRequestRequestedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRequestedNSSAITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthenticationResultRequestedNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeRequestedNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewRequestedNSSAI(nasMessage.RegistrationRequestRequestedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRequestedNSSAILenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeRequestedNSSAIData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeRequestedNSSAITable = []nasTypeRequestedNSSAIData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeRequestedNSSAIGetSetContent(t *testing.T) {
	a := nasType.NewRequestedNSSAI(nasMessage.RegistrationRequestRequestedNSSAIType)
	for _, table := range nasTypeRequestedNSSAITable {
		a.SetLen(table.inLen)
		a.SetSNSSAIValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetSNSSAIValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetSNSSAIValue())
		}
	}
}

type testRequestedNSSAIDataTemplate struct {
	in  nasType.RequestedNSSAI
	out nasType.RequestedNSSAI
}

var RequestedNSSAITestData = []nasType.RequestedNSSAI{
	{nasMessage.RegistrationRequestRequestedNSSAIType, 2, []byte{0x01, 0x02}},
}

var RequestedNSSAIExpectedTestData = []nasType.RequestedNSSAI{
	{nasMessage.RegistrationRequestRequestedNSSAIType, 2, []byte{0x01, 0x02}},
}

var RequestedNSSAITestTable = []testRequestedNSSAIDataTemplate{
	{RequestedNSSAITestData[0], RequestedNSSAIExpectedTestData[0]},
}

func TestNasTypeRequestedNSSAI(t *testing.T) {
	for i, table := range RequestedNSSAITestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRequestedNSSAI(nasMessage.RegistrationRequestRequestedNSSAIType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetSNSSAIValue(table.in.Buffer)

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
