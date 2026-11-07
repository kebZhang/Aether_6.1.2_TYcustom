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

func TestNasTypeNewUesUsageSetting(t *testing.T) {
	a := nasType.NewUesUsageSetting(nasMessage.RegistrationRequestUesUsageSettingType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeUesUsageSettingGetSetIei(t *testing.T) {
	a := nasType.NewUesUsageSetting(nasMessage.RegistrationRequestUesUsageSettingType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

func TestNasTypeUesUsageSettingGetSetLen(t *testing.T) {
	a := nasType.NewUesUsageSetting(nasMessage.RegistrationRequestUesUsageSettingType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeUesUsageSettingUesUsageSettingData struct {
	in  uint8
	out uint8
}

var nasTypeUesUsageSettingUesUsageSettingTable = []NasTypeUesUsageSettingUesUsageSettingData{
	{0x1, 0x1},
}

func TestNasTypeUesUsageSettingGetSetUesUsageSetting(t *testing.T) {
	a := nasType.NewUesUsageSetting(nasMessage.RegistrationRequestUesUsageSettingType)
	for _, table := range nasTypeUesUsageSettingUesUsageSettingTable {
		a.SetUesUsageSetting(table.in)
		if !reflect.DeepEqual(table.out, a.GetUesUsageSetting()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUesUsageSetting())
		}
	}
}

type testUesUsageSettingDataTemplate struct {
	in  nasType.UesUsageSetting
	out nasType.UesUsageSetting
}

var UesUsageSettingTestData = []nasType.UesUsageSetting{
	{nasMessage.RegistrationRequestUesUsageSettingType, 1, 0x01},
}

var UesUsageSettingExpectedData = []nasType.UesUsageSetting{
	{nasMessage.RegistrationRequestUesUsageSettingType, 1, 0x01},
}

var UesUsageSettingDataTestTable = []testUesUsageSettingDataTemplate{
	{UesUsageSettingTestData[0], UesUsageSettingExpectedData[0]},
}

func TestNasTypeUesUsageSetting(t *testing.T) {
	for _, table := range UesUsageSettingDataTestTable {
		a := nasType.NewUesUsageSetting(nasMessage.RegistrationRequestUesUsageSettingType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetUesUsageSetting(0x05)
		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
