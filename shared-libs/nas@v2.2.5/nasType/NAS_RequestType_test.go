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

func TestNasTypeNewRequestType(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRequestTypeIeiTable = []NasTypeIeiData{
	{0x08, 0x08},
}

func TestNasTypeRequestTypeGetSetIei(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	for _, table := range nasTypeRequestTypeIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeRequestRequestTypeValueType struct {
	in  uint8
	out uint8
}

var nasTypeRequestTypeRequestTypeValueTable = []nasTypeRequestRequestTypeValueType{
	{0x03, 0x03},
}

func TestNasTypeRequestTypeGetSetRequestTypeValue(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	for _, table := range nasTypeRequestTypeRequestTypeValueTable {
		a.SetRequestTypeValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetRequestTypeValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRequestTypeValue())
		}
	}
}

type RequestTypeTestDataTemplate struct {
	in  nasType.RequestType
	out nasType.RequestType
}

var RequestTypeTestData = []nasType.RequestType{
	{nasMessage.ULNASTransportRequestTypeType + 0x01},
}

var RequestTypeExpectedTestData = []nasType.RequestType{
	{0x81},
}

var RequestTypeTable = []RequestTypeTestDataTemplate{
	{RequestTypeTestData[0], RequestTypeExpectedTestData[0]},
}

func TestNasTypeRequestType(t *testing.T) {
	for _, table := range RequestTypeTable {
		a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
		a.SetIei(0x08)
		a.SetRequestTypeValue(0x01)

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
