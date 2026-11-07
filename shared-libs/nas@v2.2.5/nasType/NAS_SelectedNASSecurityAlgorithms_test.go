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

func TestNasTypeNewSelectedNASSecurityAlgorithms(t *testing.T) {
	a := nasType.NewSelectedNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionReleaseCompleteSelectedNASSecurityAlgorithmsTable = []NasTypeIeiData{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType},
}

func TestNasTypeSelectedNASSecurityAlgorithmsGetSetIei(t *testing.T) {
	a := nasType.NewSelectedNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypePDUSessionReleaseCompleteSelectedNASSecurityAlgorithmsTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeSelectedNASSecurityAlgorithmsTypeOfCipheringAlgorithmTable = []NasTypeLenuint8Data{
	{0x01, 0x01},
}

func TestNasTypeSelectedNASSecurityAlgorithmsGetSetTypeOfCipheringAlgorithm(t *testing.T) {
	a := nasType.NewSelectedNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypeSelectedNASSecurityAlgorithmsTypeOfCipheringAlgorithmTable {
		a.SetTypeOfCipheringAlgorithm(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfCipheringAlgorithm()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfCipheringAlgorithm())
		}
	}
}

type nasTypeSelectedNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmData struct {
	in  uint8
	out uint8
}

var nasTypeSelectedNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmTable = []nasTypeSelectedNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmData{
	{0x01, 0x01},
}

func TestNasTypeSelectedNASSecurityAlgorithmsGetSetTypeOfIntegrityProtectionAlgorithm(t *testing.T) {
	a := nasType.NewSelectedNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
	for _, table := range nasTypeSelectedNASSecurityAlgorithmsTypeOfIntegrityProtectionAlgorithmTable {
		a.SetTypeOfIntegrityProtectionAlgorithm(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfIntegrityProtectionAlgorithm()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfIntegrityProtectionAlgorithm())
		}
	}
}

type testSelectedNASSecurityAlgorithmsDataTemplate struct {
	inTypeOfCipheringAlgorithm           uint8
	inTypeOfIntegrityProtectionAlgorithm uint8
	in                                   nasType.SelectedNASSecurityAlgorithms
	out                                  nasType.SelectedNASSecurityAlgorithms
}

var SelectedNASSecurityAlgorithmsTestData = []nasType.SelectedNASSecurityAlgorithms{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, 0x01},
}

var SelectedNASSecurityAlgorithmsExpectedTestData = []nasType.SelectedNASSecurityAlgorithms{
	{nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType, 0x11},
}

var SelectedNASSecurityAlgorithmsTestTable = []testSelectedNASSecurityAlgorithmsDataTemplate{
	{0x01, 0x01, SelectedNASSecurityAlgorithmsTestData[0], SelectedNASSecurityAlgorithmsExpectedTestData[0]},
}

func TestNasTypeSelectedNASSecurityAlgorithms(t *testing.T) {
	for _, table := range SelectedNASSecurityAlgorithmsTestTable {
		a := nasType.NewSelectedNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)

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
