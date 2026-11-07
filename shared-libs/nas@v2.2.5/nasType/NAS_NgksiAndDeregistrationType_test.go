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

func TestNasTypeNewNgksiAndDeregistrationType(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeNgksiAndDeregistrationTypeTSC struct {
	in  uint8
	out uint8
}

var nasTypeNgksiAndDeregistrationTypeTSCTable = []nasTypeNgksiAndDeregistrationTypeTSC{
	{0x01, 0x01},
}

func TestNasTypeNgksiAndDeregistrationTypeGetSetTSC(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	for _, table := range nasTypeNgksiAndDeregistrationTypeTSCTable {
		a.SetTSC(table.in)
		if !reflect.DeepEqual(table.out, a.GetTSC()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTSC())
		}
	}
}

type nasTypeNgksiAndDeregistrationTypeNasKeySetIdentifiler struct {
	in  uint8
	out uint8
}

var nasTypeNgksiAndDeregistrationTypeNasKeySetIdentifilerTable = []nasTypeNgksiAndDeregistrationTypeNasKeySetIdentifiler{
	{0x07, 0x07},
}

func TestNasTypeNgksiAndDeregistrationTypeGetSetNasKeySetIdentifiler(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	for _, table := range nasTypeNgksiAndDeregistrationTypeNasKeySetIdentifilerTable {
		a.SetNasKeySetIdentifiler(table.in)
		if !reflect.DeepEqual(table.out, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetNasKeySetIdentifiler())
		}
	}
}

type nasTypeNgksiAndDeregistrationTypeSwitchOff struct {
	in  uint8
	out uint8
}

var nasTypeNgksiAndDeregistrationTypeSwitchOffTable = []nasTypeNgksiAndDeregistrationTypeSwitchOff{
	{0x01, 0x01},
}

func TestNasTypeNgksiAndDeregistrationTypeGetSetSwitchOff(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	for _, table := range nasTypeNgksiAndDeregistrationTypeSwitchOffTable {
		a.SetSwitchOff(table.in)
		if !reflect.DeepEqual(table.out, a.GetSwitchOff()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSwitchOff())
		}
	}
}

type nasTypeNgksiAndDeregistrationTypeReRegistrationRequired struct {
	in  uint8
	out uint8
}

var nasTypeNgksiAndDeregistrationTypeReRegistrationRequiredTable = []nasTypeNgksiAndDeregistrationTypeReRegistrationRequired{
	{0x01, 0x01},
}

func TestNasTypeNgksiAndDeregistrationTypeGetSetReRegistrationRequired(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	for _, table := range nasTypeNgksiAndDeregistrationTypeReRegistrationRequiredTable {
		a.SetReRegistrationRequired(table.in)
		if !reflect.DeepEqual(table.out, a.GetReRegistrationRequired()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetReRegistrationRequired())
		}
	}
}

type nasTypeNgksiAndDeregistrationTypeAccessType struct {
	in  uint8
	out uint8
}

var nasTypeNgksiAndDeregistrationTypeAccessTypeTable = []nasTypeNgksiAndDeregistrationTypeAccessType{
	{0x03, 0x03},
}

func TestNasTypeNgksiAndDeregistrationTypeGetSetAccessType(t *testing.T) {
	a := nasType.NewNgksiAndDeregistrationType()
	for _, table := range nasTypeNgksiAndDeregistrationTypeAccessTypeTable {
		a.SetAccessType(table.in)
		if !reflect.DeepEqual(table.out, a.GetAccessType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAccessType())
		}
	}
}

type testNgksiAndDeregistrationTypeDataTemplate struct {
	inTSC                     uint8
	inNasKeySetIdentifiler    uint8
	inSwitchOff               uint8
	inReRegistrationRequired  uint8
	inAccessType              uint8
	outTSC                    uint8
	outNasKeySetIdentifiler   uint8
	outSwitchOff              uint8
	outReRegistrationRequired uint8
	outAccessType             uint8
}

var NgksiAndDeregistrationTypeTestTable = []testNgksiAndDeregistrationTypeDataTemplate{
	{
		0x01, 0x07, 0x01, 0x01, 0x03,
		0x01, 0x07, 0x01, 0x01, 0x03,
	},
}

func TestNasTypeNgksiAndDeregistrationType(t *testing.T) {
	for _, table := range NgksiAndDeregistrationTypeTestTable {
		a := nasType.NewNgksiAndDeregistrationType()

		a.SetTSC(table.inTSC)
		a.SetNasKeySetIdentifiler(table.inNasKeySetIdentifiler)
		a.SetSwitchOff(table.inSwitchOff)
		a.SetReRegistrationRequired(table.inReRegistrationRequired)
		a.SetAccessType(table.inAccessType)

		if !reflect.DeepEqual(table.outTSC, a.GetTSC()) {
			t.Errorf("Not equal: expected %v, got %v", table.outTSC, a.GetTSC())
		}
		if !reflect.DeepEqual(table.outNasKeySetIdentifiler, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNasKeySetIdentifiler, a.GetNasKeySetIdentifiler())
		}
		if !reflect.DeepEqual(table.outSwitchOff, a.GetSwitchOff()) {
			t.Errorf("Not equal: expected %v, got %v", table.outSwitchOff, a.GetSwitchOff())
		}
		if !reflect.DeepEqual(table.outReRegistrationRequired, a.GetReRegistrationRequired()) {
			t.Errorf("Not equal: expected %v, got %v", table.outReRegistrationRequired, a.GetReRegistrationRequired())
		}
		if !reflect.DeepEqual(table.outAccessType, a.GetAccessType()) {
			t.Errorf("Not equal: expected %v, got %v", table.outAccessType, a.GetAccessType())
		}
	}
}
