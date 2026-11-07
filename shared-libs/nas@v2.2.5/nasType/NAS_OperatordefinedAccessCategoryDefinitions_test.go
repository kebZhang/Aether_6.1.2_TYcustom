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

func TestNasTypeNewOperatordefinedAccessCategoryDefinitions(t *testing.T) {
	a := nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeOperatordefinedAccessCategoryDefinitionsConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsTypeTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType, nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType},
}

func TestNasTypeOperatordefinedAccessCategoryDefinitionsGetSetIei(t *testing.T) {
	a := nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)
	for _, table := range nasTypeOperatordefinedAccessCategoryDefinitionsConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeOperatordefinedAccessCategoryDefinitionsLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeOperatordefinedAccessCategoryDefinitionsGetSetLen(t *testing.T) {
	a := nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)
	for _, table := range nasTypeOperatordefinedAccessCategoryDefinitionsLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeOperatordefinedAccessCategoryDefinitionsOperatorDefinedAccessCategoryDefintiionData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeOperatordefinedAccessCategoryDefinitionsOperatorDefinedAccessCategoryDefintiionTable = []nasTypeOperatordefinedAccessCategoryDefinitionsOperatorDefinedAccessCategoryDefintiionData{
	{2, []uint8{0x0f, 0x0f}, []uint8{0x0f, 0x0f}},
}

func TestNasTypeOperatordefinedAccessCategoryDefinitionsGetSetOperatorDefinedAccessCategoryDefintiion(t *testing.T) {
	a := nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)
	for _, table := range nasTypeOperatordefinedAccessCategoryDefinitionsOperatorDefinedAccessCategoryDefintiionTable {
		a.SetLen(table.inLen)
		a.SetOperatorDefinedAccessCategoryDefintiion(table.in)
		if !reflect.DeepEqual(table.out, a.GetOperatorDefinedAccessCategoryDefintiion()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetOperatorDefinedAccessCategoryDefintiion())
		}
	}
}

type testOperatordefinedAccessCategoryDefinitionsDataTemplate struct {
	inIei                                      uint8
	inLen                                      uint16
	inOperatorDefinedAccessCategoryDefintiion  []uint8
	outIei                                     uint8
	outLen                                     uint16
	outOperatorDefinedAccessCategoryDefintiion []uint8
}

var testOperatordefinedAccessCategoryDefinitionsTestTable = []testOperatordefinedAccessCategoryDefinitionsDataTemplate{
	{
		nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType, 2,
		[]uint8{0x0f, 0x0f},
		nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType, 2,
		[]uint8{0x0f, 0x0f},
	},
}

func TestNasTypeOperatordefinedAccessCategoryDefinitions(t *testing.T) {
	for i, table := range testOperatordefinedAccessCategoryDefinitionsTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetOperatorDefinedAccessCategoryDefintiion(table.inOperatorDefinedAccessCategoryDefintiion)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outOperatorDefinedAccessCategoryDefintiion, a.GetOperatorDefinedAccessCategoryDefintiion()) {
			t.Errorf("in(%v): out %v, actual %x", table.inOperatorDefinedAccessCategoryDefintiion, table.outOperatorDefinedAccessCategoryDefintiion, a.GetOperatorDefinedAccessCategoryDefintiion())
		}
	}
}
