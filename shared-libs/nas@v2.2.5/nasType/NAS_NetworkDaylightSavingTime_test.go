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

func TestNasTypeNewNetworkDaylightSavingTime(t *testing.T) {
	a := nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeNetworkDaylightSavingTimeConfigurationUpdateCommandNetworkDaylightSavingTimeable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType, nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType},
}

func TestNasTypeNetworkDaylightSavingTimeGetSetIei(t *testing.T) {
	a := nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
	for _, table := range nasTypeNetworkDaylightSavingTimeConfigurationUpdateCommandNetworkDaylightSavingTimeable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeNetworkDaylightSavingTimeLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeNetworkDaylightSavingTimeGetSetLen(t *testing.T) {
	a := nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
	for _, table := range nasTypeNetworkDaylightSavingTimeLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeNetworkDaylightSavingTimevalueData struct {
	in  uint8
	out uint8
}

var nasTypeNetworkDaylightSavingTimevalueTable = []nasTypeNetworkDaylightSavingTimevalueData{
	{0x03, 0x03},
}

func TestNasTypeNetworkDaylightSavingTimeGetSetvalue(t *testing.T) {
	a := nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
	for _, table := range nasTypeNetworkDaylightSavingTimevalueTable {
		a.Setvalue(table.in)
		if !reflect.DeepEqual(table.out, a.Getvalue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.Getvalue())
		}
	}
}

type testNetworkDaylightSavingTimeDataTemplate struct {
	inIei    uint8
	inLen    uint8
	invalue  uint8
	outIei   uint8
	outLen   uint8
	outvalue uint8
}

var testNetworkDaylightSavingTimeTestTable = []testNetworkDaylightSavingTimeDataTemplate{
	{
		nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType, 2, 0x03,
		nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType, 2, 0x03,
	},
}

func TestNasTypeNetworkDaylightSavingTime(t *testing.T) {
	for i, table := range testNetworkDaylightSavingTimeTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.Setvalue(table.invalue)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outvalue, a.Getvalue()) {
			t.Errorf("in(%v): out %v, actual %x", table.invalue, table.outvalue, a.Getvalue())
		}
	}
}
