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

func TestNasTypeNewNASMessageContainer(t *testing.T) {
	a := nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeNASMessageContainerRegistrationRequestAdditionalGUTITable = []NasTypeIeiData{
	{nasMessage.SecurityModeCompleteNASMessageContainerType, nasMessage.SecurityModeCompleteNASMessageContainerType},
}

func TestNasTypeNASMessageContainerGetSetIei(t *testing.T) {
	a := nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
	for _, table := range nasTypeNASMessageContainerRegistrationRequestAdditionalGUTITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeNASMessageContainerLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeNASMessageContainerGetSetLen(t *testing.T) {
	a := nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
	for _, table := range nasTypeNASMessageContainerLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeNASMessageContainerNASMessageContainerContentsData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeNASMessageContainerNASMessageContainerContentsTable = []nasTypeNASMessageContainerNASMessageContainerContentsData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeNASMessageContainerGetSetNASMessageContainerContents(t *testing.T) {
	a := nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
	for _, table := range nasTypeNASMessageContainerNASMessageContainerContentsTable {
		a.SetLen(table.inLen)
		a.SetNASMessageContainerContents(table.in)
		if !reflect.DeepEqual(table.out, a.GetNASMessageContainerContents()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetNASMessageContainerContents())
		}
	}
}

type testNASMessageContainerDataTemplate struct {
	inIei                          uint8
	inLen                          uint16
	inNASMessageContainerContents  []uint8
	outIei                         uint8
	outLen                         uint16
	outNASMessageContainerContents []uint8
}

var testNASMessageContainerTestTable = []testNASMessageContainerDataTemplate{
	{
		nasMessage.SecurityModeCompleteNASMessageContainerType, 2,
		[]uint8{0xff, 0xff},
		nasMessage.SecurityModeCompleteNASMessageContainerType, 2,
		[]uint8{0xff, 0xff},
	},
}

func TestNasTypeNASMessageContainer(t *testing.T) {
	for i, table := range testNASMessageContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetNASMessageContainerContents(table.inNASMessageContainerContents)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outNASMessageContainerContents, a.GetNASMessageContainerContents()) {
			t.Errorf("in(%v): out %v, actual %x", table.inNASMessageContainerContents, table.outNASMessageContainerContents, a.GetNASMessageContainerContents())
		}
	}
}
