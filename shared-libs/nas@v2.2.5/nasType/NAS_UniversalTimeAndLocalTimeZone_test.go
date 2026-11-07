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

func TestNasTypeNewUniversalTimeAndLocalTimeZone(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeServiceRequestUniversalTimeAndLocalTimeZoneTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType, nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetIei(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeServiceRequestUniversalTimeAndLocalTimeZoneTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneYear struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneYearTable = []nasTypeUniversalTimeAndLocalTimeZoneYear{
	{2, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetYear(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneYearTable {
		a.SetYear(table.in)

		if !reflect.DeepEqual(table.out, a.GetYear()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetYear())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneMonth struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneMonthTable = []nasTypeUniversalTimeAndLocalTimeZoneMonth{
	{3, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetMonth(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneMonthTable {
		a.SetMonth(table.in)

		if !reflect.DeepEqual(table.out, a.GetMonth()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMonth())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneDay struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneDayTable = []nasTypeUniversalTimeAndLocalTimeZoneDay{
	{2, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetDay(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneDayTable {
		a.SetDay(table.in)

		if !reflect.DeepEqual(table.out, a.GetDay()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetDay())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneHour struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneHourTable = []nasTypeUniversalTimeAndLocalTimeZoneHour{
	{3, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetHour(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneHourTable {
		a.SetHour(table.in)

		if !reflect.DeepEqual(table.out, a.GetHour()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetHour())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneMinute struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneMinuteTable = []nasTypeUniversalTimeAndLocalTimeZoneMinute{
	{2, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetMinute(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneMinuteTable {
		a.SetMinute(table.in)

		if !reflect.DeepEqual(table.out, a.GetMinute()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMinute())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneSecond struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneSecondTable = []nasTypeUniversalTimeAndLocalTimeZoneSecond{
	{2, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetSecond(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneSecondTable {
		a.SetSecond(table.in)

		if !reflect.DeepEqual(table.out, a.GetSecond()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSecond())
		}
	}
}

type nasTypeUniversalTimeAndLocalTimeZoneTimeZone struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeUniversalTimeAndLocalTimeZoneTimeZoneTable = []nasTypeUniversalTimeAndLocalTimeZoneTimeZone{
	{2, 0x01, 0x01},
}

func TestNasTypeUniversalTimeAndLocalTimeZoneGetSetTimeZone(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range nasTypeUniversalTimeAndLocalTimeZoneTimeZoneTable {
		a.SetTimeZone(table.in)

		if !reflect.DeepEqual(table.out, a.GetTimeZone()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTimeZone())
		}
	}
}

type testUniversalTimeAndLocalTimeZoneDataTemplate struct {
	in  nasType.UniversalTimeAndLocalTimeZone
	out nasType.UniversalTimeAndLocalTimeZone
}

var UniversalTimeAndLocalTimeZoneTestData = []nasType.UniversalTimeAndLocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType, [7]uint8{}},
}

var UniversalTimeAndLocalTimeZoneExpectedData = []nasType.UniversalTimeAndLocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType, [7]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}},
}

var UniversalTimeAndLocalTimeZoneTable = []testUniversalTimeAndLocalTimeZoneDataTemplate{
	{UniversalTimeAndLocalTimeZoneTestData[0], UniversalTimeAndLocalTimeZoneExpectedData[0]},
}

func TestNasTypeUniversalTimeAndLocalTimeZone(t *testing.T) {
	a := nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
	for _, table := range UniversalTimeAndLocalTimeZoneTable {
		a.SetYear(0x01)
		a.SetMonth(0x01)
		a.SetDay(0x01)
		a.SetHour(0x01)
		a.SetMinute(0x01)
		a.SetSecond(0x01)
		a.SetTimeZone(0x01)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
