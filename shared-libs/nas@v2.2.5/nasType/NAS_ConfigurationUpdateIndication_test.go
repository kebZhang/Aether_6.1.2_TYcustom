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

var ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput uint8 = 0x0D

func TestNasTypeNewConfigurationUpdateIndication(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentRequestConfigurationUpdateIndicationTable = []NasTypeIeiData{
	{ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput, 0x0D},
}

func TestNasTypeConfigurationUpdateIndicationGetSetIei(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestConfigurationUpdateIndicationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeConfigurationUpdateIndicationRED struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateIndicationREDTable = []nasTypeConfigurationUpdateIndicationRED{
	{0x01, 0x01},
}

func TestNasTypeConfigurationUpdateIndicationGetSetRED(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateIndicationREDTable {
		a.SetRED(table.in)
		if !reflect.DeepEqual(table.out, a.GetRED()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRED())
		}
	}
}

type nasTypeConfigurationUpdateIndicationACK struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateIndicationACKTable = []nasTypeConfigurationUpdateIndicationACK{
	{0x01, 0x01},
}

func TestNasTypeConfigurationUpdateIndicationGetSetACK(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateIndicationACKTable {
		a.SetACK(table.in)
		if !reflect.DeepEqual(table.out, a.GetACK()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetACK())
		}
	}
}

type testConfigurationUpdateIndicationDataTemplate struct {
	inRED uint8
	inACK uint8
	in    nasType.ConfigurationUpdateIndication
	out   nasType.ConfigurationUpdateIndication
}

var configurationUpdateIndicationTestData = []nasType.ConfigurationUpdateIndication{
	{(0xD0 + 0x03)},
}

var configurationUpdateIndicationExpectedData = []nasType.ConfigurationUpdateIndication{
	{(0xD0 + 0x03)},
}

var configurationUpdateIndicationTestTable = []testConfigurationUpdateIndicationDataTemplate{
	{0x01, 0x01, configurationUpdateIndicationTestData[0], configurationUpdateIndicationExpectedData[0]},
}

func TestNasTypeConfigurationUpdateIndication(t *testing.T) {
	for _, table := range configurationUpdateIndicationTestTable {
		a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)

		a.SetIei(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
		a.SetRED(table.inRED)
		a.SetACK(table.inACK)

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
