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

func TestNasTypeNewCause5GSM(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionReleaseCompleteCause5GSMTable = []NasTypeIeiData{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, nasMessage.PDUSessionReleaseCompleteCause5GSMType},
}

func TestNasTypeCause5GSMGetSetIei(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	for _, table := range nasTypePDUSessionReleaseCompleteCause5GSMTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeCause5GSMCauseValueData struct {
	in  uint8
	out uint8
}

var nasTypeCause5GSMOctetTable = []nasTypeCause5GSMCauseValueData{
	{0xff, 0xff},
}

func TestNasTypeCause5GSMGetSetCauseValue(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	for _, table := range nasTypeCause5GSMOctetTable {
		a.SetCauseValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetCauseValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetCauseValue())
		}
	}
}

type testCause5GSMDataTemplate struct {
	in  nasType.Cause5GSM
	out nasType.Cause5GSM
}

var cause5GSMTestData = []nasType.Cause5GSM{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, 0xff},
}

var cause5GSMExpectedTestData = []nasType.Cause5GSM{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, 0xff},
}

var cause5GSMTestTable = []testCause5GSMDataTemplate{
	{cause5GSMTestData[0], cause5GSMExpectedTestData[0]},
}

func TestNasTypeCause5GSM(t *testing.T) {
	for i, table := range cause5GSMTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)

		a.SetIei(table.in.GetIei())
		a.SetCauseValue(table.in.Octet)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
