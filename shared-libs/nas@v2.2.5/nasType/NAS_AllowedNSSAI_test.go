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

func TestNasTypeNewAllowedNSSAI(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandConfiguredNSSAITable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType},
}

func TestNasTypeAllowedNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeConfigurationUpdateCommandConfiguredNSSAITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeConfigurationUpdateCommandConfiguredNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAllowedNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeConfigurationUpdateCommandConfiguredNSSAILenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type SNSSAIValue struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeSNSSAIValueTable = []SNSSAIValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypeAllowedNSSAIGetSetSNSSAIValue(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeSNSSAIValueTable {
		a.SetLen(table.inLen)
		a.SetSNSSAIValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetSNSSAIValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetSNSSAIValue())
		}
	}
}

type testAllowedNSSAIDataTemplate struct {
	in  nasType.AllowedNSSAI
	out nasType.AllowedNSSAI
}

var AllowedNSSAITestData = []nasType.AllowedNSSAI{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, 2, []uint8{0x00, 0x01}},
}

var AllowedNSSAIExpectedTestData = []nasType.AllowedNSSAI{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, 2, []uint8{0x00, 0x01}},
}

var AllowedNSSAITable = []testAllowedNSSAIDataTemplate{
	{AllowedNSSAITestData[0], AllowedNSSAIExpectedTestData[0]},
}

func TestNasTypeAllowedNSSAI(t *testing.T) {
	for i, table := range AllowedNSSAITable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)

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
