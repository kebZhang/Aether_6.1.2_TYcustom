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

func TestNasTypeNewSORTransparentContainer(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSORTransparentContainerTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, nasMessage.RegistrationAcceptSORTransparentContainerType},
}

func TestNasTypeSORTransparentContainerGetSetIei(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeSORTransparentContainerLenData struct {
	in  uint16
	out uint16
}

var nasTypeSORTransparentContainerLenTable = []nasTypeSORTransparentContainerLenData{
	{2, 2},
}

func TestNasTypeSORTransparentContainerGetSetLen(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeSORTransparentContainerSORContentData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeSORTransparentContainerSORContentTable = []nasTypeSORTransparentContainerSORContentData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeSORTransparentContainerGetSetSORContent(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerSORContentTable {
		a.SetLen(table.inLen)
		a.SetSORContent(table.in)
		if !reflect.DeepEqual(table.out, a.GetSORContent()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetSORContent())
		}
	}
}

type testSORTransparentContainerDataTemplate struct {
	in  nasType.SORTransparentContainer
	out nasType.SORTransparentContainer
}

var SORTransparentContainerTestData = []nasType.SORTransparentContainer{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, 2, []uint8{}},
}

var SORTransparentContainerExpectedTestData = []nasType.SORTransparentContainer{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, 2, []uint8{0x01, 0x01}},
}

var SORTransparentContainerTestTable = []testSORTransparentContainerDataTemplate{
	{SORTransparentContainerTestData[0], SORTransparentContainerExpectedTestData[0]},
}

func TestNasTypeSORTransparentContainer(t *testing.T) {
	for i, table := range SORTransparentContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetSORContent([]uint8{0x01, 0x01})

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
