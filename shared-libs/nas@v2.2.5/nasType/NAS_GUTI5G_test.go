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

func TestNasTypeNewGUTI5G(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandGUTI5GTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandGUTI5GType, nasMessage.ConfigurationUpdateCommandGUTI5GType},
}

func TestNasTypeGUTI5GGetSetIei(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeConfigurationUpdateCommandGUTI5GTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeGUTI5GLenTable = []NasTypeLenUint16Data{
	{12, 12},
}

func TestNasTypeGUTI5GGetSetLen(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeGUTI5GSpare struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GSpareTable = []nasTypeGUTI5GSpare{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetSpare(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GSpareTable {
		a.SetLen(table.inLen)
		a.SetSpare(table.in)
		if !reflect.DeepEqual(table.out, a.GetSpare()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSpare())
		}
	}
}

type nasTypeGUTI5GTypeOfIdentity struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GTypeOfIdentityTable = []nasTypeGUTI5GTypeOfIdentity{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetTypeOfIdentity(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GTypeOfIdentityTable {
		a.SetLen(table.inLen)
		a.SetTypeOfIdentity(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfIdentity()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfIdentity())
		}
	}
}

type nasTypeGUTI5GMCCDigit2 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMCCDigit2Table = []nasTypeGUTI5GMCCDigit2{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMCCDigit2(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMCCDigit2Table {
		a.SetLen(table.inLen)
		a.SetMCCDigit2(table.in)
		if !reflect.DeepEqual(table.out, a.GetMCCDigit2()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMCCDigit2())
		}
	}
}

type nasTypeGUTI5GMCCDigit1 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMCCDigit1Table = []nasTypeGUTI5GMCCDigit1{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMCCDigit1(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMCCDigit1Table {
		a.SetLen(table.inLen)
		a.SetMCCDigit1(table.in)
		if !reflect.DeepEqual(table.out, a.GetMCCDigit1()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMCCDigit1())
		}
	}
}

type nasTypeGUTI5GMNCDigit3 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMNCDigit3Table = []nasTypeGUTI5GMNCDigit3{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMNCDigit3(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMNCDigit3Table {
		a.SetLen(table.inLen)
		a.SetMNCDigit3(table.in)
		if !reflect.DeepEqual(table.out, a.GetMNCDigit3()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMNCDigit3())
		}
	}
}

type nasTypeGUTI5GMCCDigit3 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMCCDigit3Table = []nasTypeGUTI5GMCCDigit3{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMCCDigit3(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMCCDigit3Table {
		a.SetLen(table.inLen)
		a.SetMCCDigit3(table.in)
		if !reflect.DeepEqual(table.out, a.GetMCCDigit3()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMCCDigit3())
		}
	}
}

type nasTypeGUTI5GMNCDigit2 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMNCDigit2Table = []nasTypeGUTI5GMNCDigit2{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMNCDigit2(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMNCDigit2Table {
		a.SetLen(table.inLen)
		a.SetMNCDigit2(table.in)
		if !reflect.DeepEqual(table.out, a.GetMNCDigit2()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMNCDigit2())
		}
	}
}

type nasTypeGUTI5GMNCDigit1 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GMNCDigit1Table = []nasTypeGUTI5GMNCDigit1{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetMNCDigit1(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GMNCDigit1Table {
		a.SetLen(table.inLen)
		a.SetMNCDigit1(table.in)
		if !reflect.DeepEqual(table.out, a.GetMNCDigit1()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMNCDigit1())
		}
	}
}

type nasTypeGUTI5GAMFRegionID struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GAMFRegionIDTable = []nasTypeGUTI5GAMFRegionID{
	{12, 0x01, 0x01},
}

func TestNasTypeGUTI5GGetSetAMFRegionID(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GAMFRegionIDTable {
		a.SetLen(table.inLen)
		a.SetAMFRegionID(table.in)
		if !reflect.DeepEqual(table.out, a.GetAMFRegionID()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAMFRegionID())
		}
	}
}

type nasTypeGUTI5GAMFSetID struct {
	inLen uint16
	in    uint16
	out   uint16
}

var nasTypeGUTI5GAMFSetIDTable = []nasTypeGUTI5GAMFSetID{
	{12, 0x101, 0x101},
}

func TestNasTypeGUTI5GGetSetAMFSetID(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GAMFSetIDTable {
		a.SetLen(table.inLen)
		a.SetAMFSetID(table.in)
		if !reflect.DeepEqual(table.out, a.GetAMFSetID()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAMFSetID())
		}
	}
}

type nasTypeGUTI5GAMFPointer struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeGUTI5GAMFPointerTable = []nasTypeGUTI5GAMFPointer{
	{12, 0x0f, 0x0f},
}

func TestNasTypeGUTI5GGetSetAMFPointer(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GAMFPointerTable {
		a.SetLen(table.inLen)
		a.SetAMFPointer(table.in)
		if !reflect.DeepEqual(table.out, a.GetAMFPointer()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAMFPointer())
		}
	}
}

type nasTypeGUTI5GTMSI5G struct {
	inLen uint16
	in    [4]uint8
	out   [4]uint8
}

var nasTypeGUTI5GTMSI5GTable = []nasTypeGUTI5GTMSI5G{
	{12, [4]uint8{0x01, 0x01, 0x01, 0x01}, [4]uint8{0x01, 0x01, 0x01, 0x01}},
}

func TestNasTypeGUTI5GGetSetTMSI5G(t *testing.T) {
	a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	for _, table := range nasTypeGUTI5GTMSI5GTable {
		a.SetLen(table.inLen)
		a.SetTMSI5G(table.in)
		if !reflect.DeepEqual(table.out, a.GetTMSI5G()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTMSI5G())
		}
	}
}

type testGUTI5GDataTemplate struct {
	inIei            uint8
	inLen            uint16
	inSpare          uint8
	inTypeOfIdentity uint8
	inMCCDigit2      uint8
	inMCCDigit1      uint8
	inMNCDigit3      uint8
	inMCCDigit3      uint8
	inMNCDigit2      uint8
	inMNCDigit1      uint8
	inAMFRegionID    uint8
	inAMFSetID       uint16
	inAMFPointer     uint8
	inTMSI5G         [4]uint8

	outIei            uint8
	outLen            uint16
	outSpare          uint8
	outTypeOfIdentity uint8
	outMCCDigit2      uint8
	outMCCDigit1      uint8
	outMNCDigit3      uint8
	outMCCDigit3      uint8
	outMNCDigit2      uint8
	outMNCDigit1      uint8
	outAMFRegionID    uint8
	outAMFSetID       uint16
	outAMFPointer     uint8
	outTMSI5G         [4]uint8
}

var gUTI5GTestTable = []testGUTI5GDataTemplate{
	{
		nasMessage.ConfigurationUpdateCommandGUTI5GType, 12, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x101, 0x01,
		[4]uint8{0x01, 0x01, 0x01, 0x01},
		nasMessage.ConfigurationUpdateCommandGUTI5GType, 12, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x101, 0x01,
		[4]uint8{0x01, 0x01, 0x01, 0x01},
	},
}

func TestNasTypeGUTI5G(t *testing.T) {
	for i, table := range gUTI5GTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetSpare(table.inSpare)
		a.SetTypeOfIdentity(table.inTypeOfIdentity)
		a.SetMCCDigit2(table.inMCCDigit2)
		a.SetMCCDigit1(table.inMCCDigit1)
		a.SetMNCDigit3(table.inMNCDigit3)
		a.SetMCCDigit3(table.inMCCDigit3)
		a.SetMNCDigit2(table.inMNCDigit2)
		a.SetMNCDigit1(table.inMNCDigit1)
		a.SetAMFRegionID(table.inAMFRegionID)
		a.SetAMFSetID(table.inAMFSetID)
		a.SetAMFPointer(table.inAMFPointer)
		a.SetTMSI5G(table.inTMSI5G)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outLen, a.GetLen()) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.GetLen())
		}
		if !reflect.DeepEqual(table.outSpare, a.GetSpare()) {
			t.Errorf("in(%v): out %v, actual %x", table.inSpare, table.outSpare, a.GetSpare())
		}
		if !reflect.DeepEqual(table.outTypeOfIdentity, a.GetTypeOfIdentity()) {
			t.Errorf("in(%v): out %v, actual %x", table.inTypeOfIdentity, table.outTypeOfIdentity, a.GetTypeOfIdentity())
		}
		if !reflect.DeepEqual(table.outMCCDigit2, a.GetMCCDigit2()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMCCDigit2, table.outMCCDigit2, a.GetMCCDigit2())
		}
		if !reflect.DeepEqual(table.outMCCDigit1, a.GetMCCDigit1()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMCCDigit1, table.outMCCDigit1, a.GetMCCDigit1())
		}
		if !reflect.DeepEqual(table.outMNCDigit3, a.GetMNCDigit3()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMNCDigit3, table.outMNCDigit3, a.GetMNCDigit3())
		}
		if !reflect.DeepEqual(table.outMCCDigit3, a.GetMCCDigit3()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMCCDigit3, table.outMCCDigit3, a.GetMCCDigit3())
		}
		if !reflect.DeepEqual(table.outMNCDigit2, a.GetMNCDigit2()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMNCDigit2, table.outMNCDigit2, a.GetMNCDigit2())
		}
		if !reflect.DeepEqual(table.outMNCDigit1, a.GetMNCDigit1()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMNCDigit1, table.outMNCDigit1, a.GetMNCDigit1())
		}
		if !reflect.DeepEqual(table.outAMFRegionID, a.GetAMFRegionID()) {
			t.Errorf("in(%v): out %v, actual %x", table.inAMFRegionID, table.outAMFRegionID, a.GetAMFRegionID())
		}
		if !reflect.DeepEqual(table.outAMFSetID, a.GetAMFSetID()) {
			t.Errorf("in(%v): out %v, actual %x", table.inAMFSetID, table.outAMFSetID, a.GetAMFSetID())
		}
		if !reflect.DeepEqual(table.outAMFPointer, a.GetAMFPointer()) {
			t.Errorf("in(%v): out %v, actual %x", table.inAMFPointer, table.outAMFPointer, a.GetAMFPointer())
		}
		if !reflect.DeepEqual(table.outTMSI5G, a.GetTMSI5G()) {
			t.Errorf("in(%v): out %v, actual %x", table.inTMSI5G, table.outTMSI5G, a.GetTMSI5G())
		}
	}
}
