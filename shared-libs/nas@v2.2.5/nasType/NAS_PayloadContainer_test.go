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

func TestNasTypeNewPayloadContainer(t *testing.T) {
	a := nasType.NewPayloadContainer(nasMessage.RegistrationRequestPayloadContainerType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePayloadContainerRegistrationRequestPayloadContainerTypeTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestPayloadContainerType, nasMessage.RegistrationRequestPayloadContainerType},
}

func TestNasTypePayloadContainerGetSetIei(t *testing.T) {
	a := nasType.NewPayloadContainer(nasMessage.RegistrationRequestPayloadContainerType)
	for _, table := range nasTypePayloadContainerRegistrationRequestPayloadContainerTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypePayloadContainerLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypePayloadContainerGetSetLen(t *testing.T) {
	a := nasType.NewPayloadContainer(nasMessage.RegistrationRequestPayloadContainerType)
	for _, table := range nasTypePayloadContainerLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypePayloadContainerPayloadContainerContentsData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypePayloadContainerPayloadContainerContentsTable = []nasTypePayloadContainerPayloadContainerContentsData{
	{2, []uint8{0x0f, 0x0f}, []uint8{0x0f, 0x0f}},
}

func TestNasTypePayloadContainerGetSetPayloadContainerContents(t *testing.T) {
	a := nasType.NewPayloadContainer(nasMessage.RegistrationRequestPayloadContainerType)
	for _, table := range nasTypePayloadContainerPayloadContainerContentsTable {
		a.SetLen(table.inLen)
		a.SetPayloadContainerContents(table.in)
		if !reflect.DeepEqual(table.out, a.GetPayloadContainerContents()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPayloadContainerContents())
		}
	}
}

type testPayloadContainerDataTemplate struct {
	inIei                       uint8
	inLen                       uint16
	inPayloadContainerContents  []uint8
	outIei                      uint8
	outLen                      uint16
	outPayloadContainerContents []uint8
}

var testPayloadContainerTestTable = []testPayloadContainerDataTemplate{
	{
		nasMessage.RegistrationRequestPayloadContainerType, 2,
		[]uint8{0x0f, 0x0f},
		nasMessage.RegistrationRequestPayloadContainerType, 2,
		[]uint8{0x0f, 0x0f},
	},
}

func TestNasTypePayloadContainer(t *testing.T) {
	for i, table := range testPayloadContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPayloadContainer(nasMessage.RegistrationRequestPayloadContainerType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetPayloadContainerContents(table.inPayloadContainerContents)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outPayloadContainerContents, a.GetPayloadContainerContents()) {
			t.Errorf("in(%v): out %v, actual %x", table.inPayloadContainerContents, table.outPayloadContainerContents, a.GetPayloadContainerContents())
		}
	}
}
