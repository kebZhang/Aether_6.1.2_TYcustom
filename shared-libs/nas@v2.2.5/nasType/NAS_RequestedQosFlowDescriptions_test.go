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

func TestNasTypeNewRequestedQosFlowDescriptions(t *testing.T) {
	a := nasType.NewRequestedQosFlowDescriptions(nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultRequestedQosFlowDescriptionsTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType, nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType},
}

func TestNasTypeRequestedQosFlowDescriptionsGetSetIei(t *testing.T) {
	a := nasType.NewRequestedQosFlowDescriptions(nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType)
	for _, table := range nasTypeAuthenticationResultRequestedQosFlowDescriptionsTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthenticationResultRequestedQosFlowDescriptionsLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeRequestedQosFlowDescriptionsGetSetLen(t *testing.T) {
	a := nasType.NewRequestedQosFlowDescriptions(nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType)
	for _, table := range nasTypeAuthenticationResultRequestedQosFlowDescriptionsLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeRequestedQosFlowDescriptionsData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeRequestedQosFlowDescriptionsTable = []nasTypeRequestedQosFlowDescriptionsData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeRequestedQosFlowDescriptionsGetSetContent(t *testing.T) {
	a := nasType.NewRequestedQosFlowDescriptions(nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType)
	for _, table := range nasTypeRequestedQosFlowDescriptionsTable {
		a.SetLen(table.inLen)
		a.SetQoSFlowDescriptions(table.in)
		if !reflect.DeepEqual(table.out, a.GetQoSFlowDescriptions()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetQoSFlowDescriptions())
		}
	}
}

type testRequestedQosFlowDescriptionsDataTemplate struct {
	in  nasType.RequestedQosFlowDescriptions
	out nasType.RequestedQosFlowDescriptions
}

var RequestedQosFlowDescriptionsTestData = []nasType.RequestedQosFlowDescriptions{
	{nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType, 2, []byte{0x01, 0x02}},
}

var RequestedQosFlowDescriptionsExpectedTestData = []nasType.RequestedQosFlowDescriptions{
	{nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType, 2, []byte{0x01, 0x02}},
}

var RequestedQosFlowDescriptionsTestTable = []testRequestedQosFlowDescriptionsDataTemplate{
	{RequestedQosFlowDescriptionsTestData[0], RequestedQosFlowDescriptionsExpectedTestData[0]},
}

func TestNasTypeRequestedQosFlowDescriptions(t *testing.T) {
	for i, table := range RequestedQosFlowDescriptionsTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRequestedQosFlowDescriptions(nasMessage.PDUSessionModificationRequestRequestedQosFlowDescriptionsType)

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
