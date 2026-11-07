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

func TestNasTypeNewRequestedQosRules(t *testing.T) {
	a := nasType.NewRequestedQosRules(nasMessage.PDUSessionModificationRequestRequestedQosRulesType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationResultRequestedQosRulesTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestRequestedQosRulesType, nasMessage.PDUSessionModificationRequestRequestedQosRulesType},
}

func TestNasTypeRequestedQosRulesGetSetIei(t *testing.T) {
	a := nasType.NewRequestedQosRules(nasMessage.PDUSessionModificationRequestRequestedQosRulesType)
	for _, table := range nasTypeAuthenticationResultRequestedQosRulesTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthenticationResultRequestedQosRulesLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeRequestedQosRulesGetSetLen(t *testing.T) {
	a := nasType.NewRequestedQosRules(nasMessage.PDUSessionModificationRequestRequestedQosRulesType)
	for _, table := range nasTypeAuthenticationResultRequestedQosRulesLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeRequestedQosRulesData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeRequestedQosRulesTable = []nasTypeRequestedQosRulesData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeRequestedQosRulesGetSetContent(t *testing.T) {
	a := nasType.NewRequestedQosRules(nasMessage.PDUSessionModificationRequestRequestedQosRulesType)
	for _, table := range nasTypeRequestedQosRulesTable {
		a.SetLen(table.inLen)
		a.SetQoSRules(table.in)
		if !reflect.DeepEqual(table.out, a.GetQoSRules()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetQoSRules())
		}
	}
}

type testRequestedQosRulesDataTemplate struct {
	in  nasType.RequestedQosRules
	out nasType.RequestedQosRules
}

var RequestedQosRulesTestData = []nasType.RequestedQosRules{
	{nasMessage.PDUSessionModificationRequestRequestedQosRulesType, 2, []byte{0x01, 0x02}},
}

var RequestedQosRulesExpectedTestData = []nasType.RequestedQosRules{
	{nasMessage.PDUSessionModificationRequestRequestedQosRulesType, 2, []byte{0x01, 0x02}},
}

var RequestedQosRulesTestTable = []testRequestedQosRulesDataTemplate{
	{RequestedQosRulesTestData[0], RequestedQosRulesExpectedTestData[0]},
}

func TestNasTypeRequestedQosRules(t *testing.T) {
	for i, table := range RequestedQosRulesTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRequestedQosRules(nasMessage.PDUSessionModificationRequestRequestedQosRulesType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetQoSRules(table.in.Buffer)

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
