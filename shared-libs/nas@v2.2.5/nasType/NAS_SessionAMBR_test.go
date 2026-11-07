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

func TestNasTypeNewSessionAMBR(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSessionAMBRPDUSessionEstablishmentAcceptSessionAMBRTypeTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationCommandSessionAMBRType, nasMessage.PDUSessionModificationCommandSessionAMBRType},
}

func TestNasTypeSessionAMBRGetSetIei(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRPDUSessionEstablishmentAcceptSessionAMBRTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeSessionAMBRLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeSessionAMBRGetSetLen(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeSessionAMBRUnitForSessionAMBRForDownlinkValueData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeSessionAMBRUnitForSessionAMBRForDownlinkValueTable = []nasTypeSessionAMBRUnitForSessionAMBRForDownlinkValueData{
	{2, 0x01, 0x01},
}

func TestNasTypeSessionAMBRGetSetUnitForSessionAMBRForDownlink(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRUnitForSessionAMBRForDownlinkValueTable {
		a.SetLen(table.inLen)
		a.SetUnitForSessionAMBRForDownlink(table.in)
		if !reflect.DeepEqual(table.out, a.GetUnitForSessionAMBRForDownlink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUnitForSessionAMBRForDownlink())
		}
	}
}

type nasTypeSessionAMBRSessionAMBRForDownlinkData struct {
	inLen uint8
	in    [2]uint8
	out   [2]uint8
}

var nasTypeSessionAMBRSessionAMBRForDownlinkTable = []nasTypeSessionAMBRSessionAMBRForDownlinkData{
	{2, [2]uint8{0x01, 0x01}, [2]uint8{0x01, 0x01}},
}

func TestNasTypeSessionAMBRGetSetSessionAMBRForDownlink(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRSessionAMBRForDownlinkTable {
		a.SetLen(table.inLen)
		a.SetSessionAMBRForDownlink(table.in)
		if !reflect.DeepEqual(table.out, a.GetSessionAMBRForDownlink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSessionAMBRForDownlink())
		}
	}
}

type nasTypeSessionAMBRUnitForSessionAMBRForUplinkValueData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeSessionAMBRUnitForSessionAMBRForUplinkValueTable = []nasTypeSessionAMBRUnitForSessionAMBRForUplinkValueData{
	{2, 0x01, 0x01},
}

func TestNasTypeSessionAMBRGetSetUnitForSessionAMBRForUplink(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRUnitForSessionAMBRForUplinkValueTable {
		a.SetLen(table.inLen)
		a.SetUnitForSessionAMBRForUplink(table.in)
		if !reflect.DeepEqual(table.out, a.GetUnitForSessionAMBRForUplink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetUnitForSessionAMBRForUplink())
		}
	}
}

type nasTypeSessionAMBRSessionAMBRForUplinkData struct {
	inLen uint8
	in    [2]uint8
	out   [2]uint8
}

var nasTypeSessionAMBRSessionAMBRForUplinkTable = []nasTypeSessionAMBRSessionAMBRForUplinkData{
	{2, [2]uint8{0x01, 0x01}, [2]uint8{0x01, 0x01}},
}

func TestNasTypeSessionAMBRGetSetSessionAMBRForUplink(t *testing.T) {
	a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
	for _, table := range nasTypeSessionAMBRSessionAMBRForUplinkTable {
		a.SetLen(table.inLen)
		a.SetSessionAMBRForUplink(table.in)
		if !reflect.DeepEqual(table.out, a.GetSessionAMBRForUplink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSessionAMBRForUplink())
		}
	}
}

type testSessionAMBRDataTemplate struct {
	in  nasType.SessionAMBR
	out nasType.SessionAMBR
}

var sessionAMBRTestData = []nasType.SessionAMBR{
	{nasMessage.PDUSessionModificationCommandSessionAMBRType, 6, [6]uint8{}},
}

var sessionAMBRExpectedTestData = []nasType.SessionAMBR{
	{nasMessage.PDUSessionModificationCommandSessionAMBRType, 6, [6]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01}},
}

var sessionAMBRTable = []testSessionAMBRDataTemplate{
	{sessionAMBRTestData[0], sessionAMBRExpectedTestData[0]},
}

func TestNasTypeSessionAMBR(t *testing.T) {
	for i, table := range sessionAMBRTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)

		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetUnitForSessionAMBRForDownlink(0x01)
		a.SetSessionAMBRForDownlink([2]uint8{0x01, 0x01})
		a.SetUnitForSessionAMBRForUplink(0x01)
		a.SetSessionAMBRForUplink([2]uint8{0x01, 0x01})

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("out %v, actual %x", table.out.Octet, a.Octet)
		}
	}
}
