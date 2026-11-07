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

func TestNasTypeNewMappedEPSBearerContexts(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationRequestMappedEPSBearerContextsTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType},
}

func TestNasTypeMappedEPSBearerContextsGetSetIei(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeRegistrationRequestMappedEPSBearerContextsTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeMappedEPSBearerContextsLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeMappedEPSBearerContextsGetSetLen(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeMappedEPSBearerContextsLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeMappedEPSBearerContextsMappedEPSBearerContextData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeMappedEPSBearerContextsMappedEPSBearerContextTable = []nasTypeMappedEPSBearerContextsMappedEPSBearerContextData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeMappedEPSBearerContextsGetSetMappedEPSBearerContext(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeMappedEPSBearerContextsMappedEPSBearerContextTable {
		a.SetLen(table.inLen)
		a.SetMappedEPSBearerContext(table.in)
		if !reflect.DeepEqual(table.out, a.GetMappedEPSBearerContext()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMappedEPSBearerContext())
		}
	}
}

type testMappedEPSBearerContextsDataTemplate struct {
	inIei                     uint8
	inLen                     uint16
	inMappedEPSBearerContext  []uint8
	outIei                    uint8
	outLen                    uint16
	outMappedEPSBearerContext []uint8
}

var testMappedEPSBearerContextsTestTable = []testMappedEPSBearerContextsDataTemplate{
	{
		nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, 2,
		[]uint8{0xff, 0xff},
		nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, 2,
		[]uint8{0xff, 0xff},
	},
}

func TestNasTypeMappedEPSBearerContexts(t *testing.T) {
	for i, table := range testMappedEPSBearerContextsTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetMappedEPSBearerContext(table.inMappedEPSBearerContext)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outMappedEPSBearerContext, a.GetMappedEPSBearerContext()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMappedEPSBearerContext, table.outMappedEPSBearerContext, a.GetMappedEPSBearerContext())
		}
	}
}
