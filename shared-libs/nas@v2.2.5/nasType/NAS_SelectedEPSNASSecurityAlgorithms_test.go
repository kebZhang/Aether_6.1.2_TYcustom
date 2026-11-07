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

func TestNasTypeNewSelectedEPSNASSecurityAlgorithms(t *testing.T) {
	a := nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionReleaseCompleteSelectedEPSNASSecurityAlgorithmsTable = []NasTypeIeiData{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType},
}

func TestNasTypeSelectedEPSNASSecurityAlgorithmsGetSetIei(t *testing.T) {
	a := nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypePDUSessionReleaseCompleteSelectedEPSNASSecurityAlgorithmsTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfCipheringAlgorithmTable = []NasTypeLenuint8Data{
	{0x01, 0x01},
}

func TestNasTypeSelectedEPSNASSecurityAlgorithmsGetSetTypeOfCipheringAlgorithm(t *testing.T) {
	a := nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfCipheringAlgorithmTable {
		a.SetTypeOfCipheringAlgorithm(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfCipheringAlgorithm()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfCipheringAlgorithm())
		}
	}
}

type nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmData struct {
	in  uint8
	out uint8
}

var nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmTable = []nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmData{
	{0x01, 0x01},
}

func TestNasTypeSelectedEPSNASSecurityAlgorithmsGetSetTypeOfIntegrityProtectionAlgorithm(t *testing.T) {
	a := nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypeSelectedEPSNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmTable {
		a.SetTypeOfIntegrityProtectionAlgorithm(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfIntegrityProtectionAlgorithm()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfIntegrityProtectionAlgorithm())
		}
	}
}

type testSelectedEPSNASSecurityAlgorithmsDataTemplate struct {
	inTypeOfCipheringAlgorithm           uint8
	inTypeOfIntegrityProtectionAlgorithm uint8
	in                                   nasType.SelectedEPSNASSecurityAlgorithms
	out                                  nasType.SelectedEPSNASSecurityAlgorithms
}

var SelectedEPSNASSecurityAlgorithmsTestData = []nasType.SelectedEPSNASSecurityAlgorithms{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, 0x01},
}

var SelectedEPSNASSecurityAlgorithmsExpectedTestData = []nasType.SelectedEPSNASSecurityAlgorithms{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, 0x11},
}

var SelectedEPSNASSecurityAlgorithmsTestTable = []testSelectedEPSNASSecurityAlgorithmsDataTemplate{
	{0x01, 0x01, SelectedEPSNASSecurityAlgorithmsTestData[0], SelectedEPSNASSecurityAlgorithmsExpectedTestData[0]},
}

func TestNasTypeSelectedEPSNASSecurityAlgorithms(t *testing.T) {
	for _, table := range SelectedEPSNASSecurityAlgorithmsTestTable {
		a := nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)

		a.SetIei(table.in.GetIei())
		a.SetTypeOfCipheringAlgorithm(table.inTypeOfCipheringAlgorithm)
		a.SetTypeOfIntegrityProtectionAlgorithm(table.inTypeOfIntegrityProtectionAlgorithm)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
