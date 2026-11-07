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

func TestNasTypeNewIntegrityProtectionMaximumDataRate(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionModificationRequestIntegrityProtectionMaximumDataRateTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetIei(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypePDUSessionModificationRequestIntegrityProtectionMaximumDataRateTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkData struct {
	in  uint8
	out uint8
}

var nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkTable = []nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkData{
	{0xff, 0xff},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkTable {
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(table.in)
		if !reflect.DeepEqual(table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink())
		}
	}
}

type nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkData struct {
	in  uint8
	out uint8
}

var nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkTable = []nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkData{
	{0xff, 0xff},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkTable {
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(table.in)
		if !reflect.DeepEqual(table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink())
		}
	}
}

type testIntegrityProtectionMaximumDataRateDataTemplate struct {
	inIei                                                             uint8
	inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink    uint8
	inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink  uint8
	outIei                                                            uint8
	outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink   uint8
	outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink uint8
}

var integrityProtectionMaximumDataRateTestTable = []testIntegrityProtectionMaximumDataRateDataTemplate{
	{
		nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, 0xff, 0x11,
		nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, 0xff, 0x11,
	},
}

func TestNasTypeIntegrityProtectionMaximumDataRate(t *testing.T) {
	for i, table := range integrityProtectionMaximumDataRateTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)

		a.SetIei(table.inIei)
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink)
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink())
		}
		if !reflect.DeepEqual(table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink()) {
			t.Errorf("in(%v): out %v, actual %x", table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink())
		}
	}
}
