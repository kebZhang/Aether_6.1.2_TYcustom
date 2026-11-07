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

func TestNasTypeNewMobileIdentity5GS(t *testing.T) {
	a := nasType.NewMobileIdentity5GS(nasMessage.RegistrationRequestAdditionalGUTIType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeMobileIdentity5GSRegistrationRequestAdditionalGUTITable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestAdditionalGUTIType, nasMessage.RegistrationRequestAdditionalGUTIType},
}

func TestNasTypeMobileIdentity5GSGetSetIei(t *testing.T) {
	a := nasType.NewMobileIdentity5GS(nasMessage.RegistrationRequestAdditionalGUTIType)
	for _, table := range nasTypeMobileIdentity5GSRegistrationRequestAdditionalGUTITable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeMobileIdentity5GSLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeMobileIdentity5GSGetSetLen(t *testing.T) {
	a := nasType.NewMobileIdentity5GS(nasMessage.RegistrationRequestAdditionalGUTIType)
	for _, table := range nasTypeMobileIdentity5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeMobileIdentity5GSMobileIdentity5GSContentsData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeMobileIdentity5GSMobileIdentity5GSContentsTable = []nasTypeMobileIdentity5GSMobileIdentity5GSContentsData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeMobileIdentity5GSGetSetMobileIdentity5GSContents(t *testing.T) {
	a := nasType.NewMobileIdentity5GS(nasMessage.RegistrationRequestAdditionalGUTIType)
	for _, table := range nasTypeMobileIdentity5GSMobileIdentity5GSContentsTable {
		a.SetLen(table.inLen)
		a.SetMobileIdentity5GSContents(table.in)
		if !reflect.DeepEqual(table.out, a.GetMobileIdentity5GSContents()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMobileIdentity5GSContents())
		}
	}
}

type testMobileIdentity5GSDataTemplate struct {
	inIei                        uint8
	inLen                        uint16
	inMobileIdentity5GSContents  []uint8
	outIei                       uint8
	outLen                       uint16
	outMobileIdentity5GSContents []uint8
}

var testMobileIdentity5GSTestTable = []testMobileIdentity5GSDataTemplate{
	{
		nasMessage.RegistrationRequestAdditionalGUTIType, 2,
		[]uint8{0xff, 0xff},
		nasMessage.RegistrationRequestAdditionalGUTIType, 2,
		[]uint8{0xff, 0xff},
	},
}

func TestNasTypeMobileIdentity5GS(t *testing.T) {
	for i, table := range testMobileIdentity5GSTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewMobileIdentity5GS(nasMessage.RegistrationRequestAdditionalGUTIType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetMobileIdentity5GSContents(table.inMobileIdentity5GSContents)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outMobileIdentity5GSContents, a.GetMobileIdentity5GSContents()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMobileIdentity5GSContents, table.outMobileIdentity5GSContents, a.GetMobileIdentity5GSContents())
		}
	}
}
