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

func TestNasTypeNewAdditionalInformation(t *testing.T) {
	a := nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeULNASTransportAdditionalInformationTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportAdditionalInformationType, nasMessage.ULNASTransportAdditionalInformationType},
}

func TestNasTypeAdditionalInformationGetSetIei(t *testing.T) {
	a := nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
	for _, table := range nasTypeULNASTransportAdditionalInformationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeULNASTransportAdditionalInformationTLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAdditionalInformationGetSetLen(t *testing.T) {
	a := nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
	for _, table := range nasTypeULNASTransportAdditionalInformationTLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type AdditionalInformationValue struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeAdditionalInformationValueTable = []AdditionalInformationValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypeAdditionalInformationGetSetAdditionalInformationValue(t *testing.T) {
	a := nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
	for _, table := range nasTypeAdditionalInformationValueTable {
		a.SetLen(table.inLen)
		a.SetAdditionalInformationValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetAdditionalInformationValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetAdditionalInformationValue())
		}
	}
}

type testAdditionalInformationDataTemplate struct {
	in  nasType.AdditionalInformation
	out nasType.AdditionalInformation
}

var additionalInformationTestData = []nasType.AdditionalInformation{
	{nasMessage.ULNASTransportAdditionalInformationType, 2, []uint8{0x00, 0x01}},
}

var additionalInformationExpectedTestData = []nasType.AdditionalInformation{
	{nasMessage.ULNASTransportAdditionalInformationType, 2, []uint8{0x00, 0x01}},
}

var additionalInformationTable = []testAdditionalInformationDataTemplate{
	{additionalInformationTestData[0], additionalInformationExpectedTestData[0]},
}

func TestNasTypeAdditionalInformation(t *testing.T) {
	for i, table := range additionalInformationTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetAdditionalInformationValue(table.in.Buffer)

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
