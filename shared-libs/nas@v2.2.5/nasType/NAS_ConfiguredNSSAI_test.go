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

func TestNasTypeNewConfiguredNSSAI(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationRequestConfiguredNSSAIIeiTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, nasMessage.ConfigurationUpdateCommandAllowedNSSAIType},
}

func TestNasTypeConfiguredNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	for _, table := range nasTypeAuthenticationRequestConfiguredNSSAIIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeConfiguredNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeConfiguredNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	for _, table := range nasTypeConfiguredNSSAILenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypetConfiguredNSSAISNSSAIValue struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeConfiguredNSSAISNSSAIValueTable = []nasTypetConfiguredNSSAISNSSAIValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeConfiguredNSSAIGetSetSNSSAIValue(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(0)
	for _, table := range nasTypeConfiguredNSSAISNSSAIValueTable {
		a.SetLen(table.inLen)
		a.SetSNSSAIValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetSNSSAIValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetSNSSAIValue())
		}
	}
}

type testConfiguredNSSAIDataTemplate struct {
	in  nasType.ConfiguredNSSAI
	out nasType.ConfiguredNSSAI
}

var configuredNSSAITestData = []nasType.ConfiguredNSSAI{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, 2, []byte{0x00, 0x00}},
}

var configuredNSSAIExpectedData = []nasType.ConfiguredNSSAI{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, 2, []byte{0x00, 0x00}},
}

var configuredNSSAITestTable = []testConfiguredNSSAIDataTemplate{
	{configuredNSSAITestData[0], configuredNSSAIExpectedData[0]},
}

func TestNasTypeConfiguredNSSAI(t *testing.T) {
	for i, table := range configuredNSSAITestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewConfiguredNSSAI(0)

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
