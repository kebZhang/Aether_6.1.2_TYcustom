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

func TestNasTypeNewTAIList(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeTAIListTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptTAIListType, nasMessage.RegistrationAcceptTAIListType},
}

func TestNasTypeTAIListGetSetIei(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeTAIListLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeTAIListGetSetLen(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeTAIListPartialTrackingAreaIdentityListData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeTAIListPartialTrackingAreaIdentityListTable = []nasTypeTAIListPartialTrackingAreaIdentityListData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeTAIListGetSetPartialTrackingAreaIdentityList(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListPartialTrackingAreaIdentityListTable {
		a.SetLen(table.inLen) // fix it, set input length
		a.SetPartialTrackingAreaIdentityList(table.in)
		if !reflect.DeepEqual(table.out, a.GetPartialTrackingAreaIdentityList()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetPartialTrackingAreaIdentityList())
		}
	}
}

type testTAIListDataTemplate struct {
	in  nasType.TAIList
	out nasType.TAIList
}

var TAIListTestData = []nasType.TAIList{
	{nasMessage.RegistrationAcceptTAIListType, 2, []uint8{}},
}

var TAIListExpectedTestData = []nasType.TAIList{
	{nasMessage.RegistrationAcceptTAIListType, 2, []uint8{0x01, 0x01}},
}

var TAIListTestTable = []testTAIListDataTemplate{
	{TAIListTestData[0], TAIListExpectedTestData[0]},
}

func TestNasTypeTAIList(t *testing.T) {
	for i, table := range TAIListTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetPartialTrackingAreaIdentityList([]uint8{0x01, 0x01})

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
