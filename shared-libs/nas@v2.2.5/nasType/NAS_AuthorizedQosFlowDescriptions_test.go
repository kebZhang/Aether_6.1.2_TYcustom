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

func TestNasTypeNewAuthorizedQosFlowDescriptions(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetIei(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsLenTable = []NasTypeLenUint16Data{
	{12, 12},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetLen(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeQoSFlowDescription struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeQoSFlowDescriptionTable = []nasTypeQoSFlowDescription{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetQoSFlowDescription(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypeQoSFlowDescriptionTable {
		a.SetLen(table.inLen)
		a.SetQoSFlowDescriptions(table.in)
		if !reflect.DeepEqual(table.out, a.GetQoSFlowDescriptions()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetQoSFlowDescriptions())
		}
	}
}

type testAuthorizedQosFlowDescriptionsDataTemplate struct {
	in  nasType.AuthorizedQosFlowDescriptions
	out nasType.AuthorizedQosFlowDescriptions
}

var AuthorizedQosFlowDescriptionsTestData = []nasType.AuthorizedQosFlowDescriptions{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, 2, []uint8{0x00, 0x01}},
}

var AuthorizedQosFlowDescriptionsExpectedTestData = []nasType.AuthorizedQosFlowDescriptions{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, 2, []uint8{0x00, 0x01}},
}

var AuthorizedQosFlowDescriptionsTable = []testAuthorizedQosFlowDescriptionsDataTemplate{
	{AuthorizedQosFlowDescriptionsTestData[0], AuthorizedQosFlowDescriptionsExpectedTestData[0]},
}

func TestNasTypeAuthorizedQosFlowDescriptions(t *testing.T) {
	for i, table := range AuthorizedQosFlowDescriptionsTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetQoSFlowDescriptions(table.in.Buffer)

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
