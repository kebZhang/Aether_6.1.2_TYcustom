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

func TestNasTypeNewPDUSessionReactivationResultErrorCause(t *testing.T) {
	a := nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationAcceptPDUSessionReactivationResultErrorCauseTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType, nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType},
}

func TestNasTypePDUSessionReactivationResultErrorCauseGetSetIei(t *testing.T) {
	a := nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
	for _, table := range nasTypeRegistrationAcceptPDUSessionReactivationResultErrorCauseTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeRegistrationAcceptPDUSessionReactivationResultErrorCauseLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypePDUSessionReactivationResultErrorCauseGetSetLen(t *testing.T) {
	a := nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
	for _, table := range nasTypeRegistrationAcceptPDUSessionReactivationResultErrorCauseLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type PDUSessionIDAndCauseValue struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypePDUSessionIDAndCauseValueTable = []PDUSessionIDAndCauseValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypePDUSessionReactivationResultErrorCauseGetSetPDUSessionIDAndCauseValue(t *testing.T) {
	a := nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
	for _, table := range nasTypePDUSessionIDAndCauseValueTable {
		a.SetLen(table.inLen)
		a.SetPDUSessionIDAndCauseValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetPDUSessionIDAndCauseValue()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetPDUSessionIDAndCauseValue())
		}
	}
}

type testPDUSessionReactivationResultErrorCauseDataTemplate struct {
	in  nasType.PDUSessionReactivationResultErrorCause
	out nasType.PDUSessionReactivationResultErrorCause
}

var pDUSessionReactivationResultErrorCauseTestData = []nasType.PDUSessionReactivationResultErrorCause{
	{nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType, 2, []uint8{0x00, 0x01}},
}

var pDUSessionReactivationResultErrorCauseExpectedTestData = []nasType.PDUSessionReactivationResultErrorCause{
	{nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType, 2, []uint8{0x00, 0x01}},
}

var pDUSessionReactivationResultErrorCauseInformationTable = []testPDUSessionReactivationResultErrorCauseDataTemplate{
	{pDUSessionReactivationResultErrorCauseTestData[0], pDUSessionReactivationResultErrorCauseExpectedTestData[0]},
}

func TestNasTypePDUSessionReactivationResultErrorCauseData(t *testing.T) {
	for i, table := range pDUSessionReactivationResultErrorCauseInformationTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetPDUSessionIDAndCauseValue(table.in.Buffer)

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
