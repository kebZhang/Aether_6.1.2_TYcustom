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

func TestNasTypeNewRejectedNSSAI(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultRejectedNSSAITable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, nasMessage.RegistrationAcceptRejectedNSSAIType},
}

func TestNasTypeRejectedNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRejectedNSSAITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}

		// if a.GetIei() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetIei())
		// }
	}
}

var nasTypeAuthenticationResultRejectedNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeRejectedNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRejectedNSSAILenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
		// if a.GetLen() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetLen())
		// }
	}
}

type nasTypeRejectedNSSAIContentsData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeRejectedNSSAIContentsTable = []nasTypeRejectedNSSAIContentsData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeRejectedNSSAIGetSetRejectedNSSAIContents(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeRejectedNSSAIContentsTable {
		a.SetLen(table.inLen)
		a.SetRejectedNSSAIContents(table.in)
		if !reflect.DeepEqual(table.out, a.GetRejectedNSSAIContents()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetRejectedNSSAIContents())
		}
	}
}

type testRejectedNSSAIDataTemplate struct {
	in  nasType.RejectedNSSAI
	out nasType.RejectedNSSAI
}

var RejectedNSSAITestData = []nasType.RejectedNSSAI{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, 2, []byte{0x00, 0x00}},
}

var RejectedNSSAIExpectedTestData = []nasType.RejectedNSSAI{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, 2, []byte{0x00, 0x00}},
}

var RejectedNSSAITestTable = []testRejectedNSSAIDataTemplate{
	{RejectedNSSAITestData[0], RejectedNSSAIExpectedTestData[0]},
}

func TestNasTypeRejectedNSSAI(t *testing.T) {
	for i, table := range RejectedNSSAITestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetRejectedNSSAIContents(table.in.Buffer)

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
