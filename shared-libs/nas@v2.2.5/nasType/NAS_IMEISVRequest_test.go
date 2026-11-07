// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

var SecurityModeCommandIMEISVRequestTypeIeiInput uint8 = 0x0E

func TestNasTypeNewIMEISVRequest(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentRequestIMEISVRequestTable = []NasTypeIeiData{
	{SecurityModeCommandIMEISVRequestTypeIeiInput, 0x0E},
}

func TestNasTypeIMEISVRequestGetSetIei(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestIMEISVRequestTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeIMEISVRequestIMEISVRequestValue struct {
	in  uint8
	out uint8
}

var nasTypeIMEISVRequestIMEISVRequestValueTable = []nasTypeIMEISVRequestIMEISVRequestValue{
	{0x07, 0x07},
}

func TestNasTypeIMEISVRequestGetSetIMEISVRequestValue(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	for _, table := range nasTypeIMEISVRequestIMEISVRequestValueTable {
		a.SetIMEISVRequestValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetIMEISVRequestValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIMEISVRequestValue())
		}
	}
}

type testIMEISVRequestDataTemplate struct {
	inIei                uint8
	inIMEISVRequestValue uint8

	outIei                uint8
	outIMEISVRequestValue uint8
}

var iMEISVRequestTestTable = []testIMEISVRequestDataTemplate{
	{
		SecurityModeCommandIMEISVRequestTypeIeiInput, 0x07,
		0x0E, 0x07,
	},
}

func TestNasTypeIMEISVRequest(t *testing.T) {
	for _, table := range iMEISVRequestTestTable {
		a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)

		a.SetIei(table.inIei)
		a.SetIMEISVRequestValue(table.inIMEISVRequestValue)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outIMEISVRequestValue, a.GetIMEISVRequestValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIMEISVRequestValue, table.outIMEISVRequestValue, a.GetIMEISVRequestValue())
		}
	}
}
