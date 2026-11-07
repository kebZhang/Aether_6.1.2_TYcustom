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

var AlwaysonPDUSessionRequestedIeiInput uint8 = 0x0B

func TestNasTypeNewAlwaysonPDUSessionRequested(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedTable = []NasTypeIeiData{
	{AlwaysonPDUSessionRequestedIeiInput, 0x0B},
}

func TestNasTypeAlwaysonPDUSessionRequestedGetSetIei(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeAlwaysonPDUSessionRequestedAPSI struct {
	in  uint8
	out uint8
}

var nasTypeAlwaysonPDUSessionRequestedAPSRTable = []nasTypeAlwaysonPDUSessionRequestedAPSI{
	{0x01, 0x01},
}

func TestNasTypeAlwaysonPDUSessionRequestedGetSetAPSR(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(nasMessage.PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType)
	for _, table := range nasTypeAlwaysonPDUSessionRequestedAPSRTable {
		a.SetAPSR(table.in)
		if !reflect.DeepEqual(table.out, a.GetAPSR()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetAPSR())
		}
	}
}

type testAlwaysonPDUSessionRequestedDataTemplate struct {
	in  nasType.AlwaysonPDUSessionRequested
	out nasType.AlwaysonPDUSessionRequested
}

var alwaysonPDUSessionRequestedTestData = []nasType.AlwaysonPDUSessionRequested{
	{(0xB0 + 0x01)},
}

var alwaysonPDUSessionRequestedExpectedTestData = []nasType.AlwaysonPDUSessionRequested{
	{(0xB0 + 0x01)},
}

var alwaysonPDUSessionRequestedTestTable = []testAlwaysonPDUSessionRequestedDataTemplate{
	{alwaysonPDUSessionRequestedTestData[0], alwaysonPDUSessionRequestedExpectedTestData[0]},
}

func TestNasTypeAlwaysonPDUSessionRequested(t *testing.T) {
	for _, table := range alwaysonPDUSessionRequestedTestTable {
		a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)

		a.SetIei(AlwaysonPDUSessionRequestedIeiInput)
		a.SetAPSR(table.in.GetAPSR())

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
