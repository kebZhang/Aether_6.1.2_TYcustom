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

func TestNasTypeNewLocalTimeZone(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandLocalTimeZoneTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, nasMessage.ConfigurationUpdateCommandLocalTimeZoneType},
}

func TestNasTypeLocalTimeZoneGetSetIei(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	for _, table := range nasTypeConfigurationUpdateCommandLocalTimeZoneTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeLocalTimeZoneTimeZoneData struct {
	in  uint8
	out uint8
}

var nasTypeLocalTimeZoneOctetTable = []nasTypeLocalTimeZoneTimeZoneData{
	{0xff, 0xff},
}

func TestNasTypeLocalTimeZoneGetSetTimeZone(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	for _, table := range nasTypeLocalTimeZoneOctetTable {
		a.SetTimeZone(table.in)
		if !reflect.DeepEqual(table.out, a.GetTimeZone()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTimeZone())
		}
	}
}

type testLocalTimeZoneDataTemplate struct {
	in  nasType.LocalTimeZone
	out nasType.LocalTimeZone
}

var LocalTimeZoneTestData = []nasType.LocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, 0xff},
}

var LocalTimeZoneExpectedTestData = []nasType.LocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, 0xff},
}

var LocalTimeZoneTestTable = []testLocalTimeZoneDataTemplate{
	{LocalTimeZoneTestData[0], LocalTimeZoneExpectedTestData[0]},
}

func TestNasTypeLocalTimeZone(t *testing.T) {
	for i, table := range LocalTimeZoneTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)

		a.SetIei(table.in.GetIei())
		a.SetTimeZone(table.in.Octet)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
