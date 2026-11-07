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

func TestNasTypeNewSSCMode(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSSCModeIeiTable = []NasTypeIeiData{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetIei(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	for _, table := range nasTypeSSCModeIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeRequestSpareType struct {
	in  uint8
	out uint8
}

var nasTypeSSCModeSpareTable = []nasTypeRequestSpareType{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetSpare(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	for _, table := range nasTypeSSCModeSpareTable {
		a.SetSpare(table.in)
		if !reflect.DeepEqual(table.out, a.GetSpare()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSpare())
		}
	}
}

type nasTypeRequestSSCModeType struct {
	in  uint8
	out uint8
}

var nasTypeSSCModeSSCModeTable = []nasTypeRequestSSCModeType{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetSSCMode(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	for _, table := range nasTypeSSCModeSSCModeTable {
		a.SetSSCMode(table.in)
		if !reflect.DeepEqual(table.out, a.GetSSCMode()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSSCMode())
		}
	}
}

type SSCModeTestDataTemplate struct {
	in  nasType.SSCMode
	out nasType.SSCMode
}

var SSCModeTestData = []nasType.SSCMode{
	{nasMessage.PDUSessionEstablishmentRequestSSCModeType},
}

var SSCModeExpectedTestData = []nasType.SSCMode{
	{0x19},
}

var SSCModeTable = []SSCModeTestDataTemplate{
	{SSCModeTestData[0], SSCModeExpectedTestData[0]},
}

func TestNasTypeSSCMode(t *testing.T) {
	for _, table := range SSCModeTable {
		a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
		a.SetIei(0x01)
		a.SetSpare(0x01)
		a.SetSSCMode(0x01)

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
