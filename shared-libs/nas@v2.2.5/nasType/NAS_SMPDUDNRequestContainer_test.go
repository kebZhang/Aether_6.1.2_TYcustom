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

func TestNasTypeNewSMPDUDNRequestContainer(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSMPDUDNRequestContainerTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType},
}

func TestNasTypeSMPDUDNRequestContainerGetSetIei(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeSMPDUDNRequestContainerLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeSMPDUDNRequestContainerGetSetLen(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeSMPDUDNRequestContainerDNSpecificIdentityData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeSMPDUDNRequestContainerDNSpecificIdentityTable = []nasTypeSMPDUDNRequestContainerDNSpecificIdentityData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeSMPDUDNRequestContainerGetSetDNSpecificIdentity(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerDNSpecificIdentityTable {
		a.SetLen(table.inLen) // fix it, set input length
		a.SetDNSpecificIdentity(table.in)
		if !reflect.DeepEqual(table.out, a.GetDNSpecificIdentity()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetDNSpecificIdentity())
		}
	}
}

type testSMPDUDNRequestContainerDataTemplate struct {
	in  nasType.SMPDUDNRequestContainer
	out nasType.SMPDUDNRequestContainer
}

var SMPDUDNRequestContainerTestData = []nasType.SMPDUDNRequestContainer{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, 2, []uint8{}},
}

var SMPDUDNRequestContainerExpectedTestData = []nasType.SMPDUDNRequestContainer{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, 2, []uint8{0x01, 0x01}},
}

var SMPDUDNRequestContainerTestTable = []testSMPDUDNRequestContainerDataTemplate{
	{SMPDUDNRequestContainerTestData[0], SMPDUDNRequestContainerExpectedTestData[0]},
}

func TestNasTypeSMPDUDNRequestContainer(t *testing.T) {
	for i, table := range SMPDUDNRequestContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetDNSpecificIdentity([]uint8{0x01, 0x01})

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
