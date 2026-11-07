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

func TestNasTypeNewEPSNASMessageContainer(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationRequestEPSNASMessageContainerIeiTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, nasMessage.RegistrationRequestEPSNASMessageContainerType},
}

func TestNasTypeEPSNASMessageContainerGetSetIei(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeRegistrationRequestEPSNASMessageContainerIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeEPSNASMessageContainerLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeEPSNASMessageContainerGetSetLen(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeEPSNASMessageContainerLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeEPSNASMessageContainerEPANASMessageContainer struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeEPSNASMessageContainerEPANASMessageContainerTable = []nasTypeEPSNASMessageContainerEPANASMessageContainer{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeEPSNASMessageContainerGetSetEPANASMessageContainer(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeEPSNASMessageContainerEPANASMessageContainerTable {
		a.SetLen(table.inLen)
		a.SetEPANASMessageContainer(table.in)
		if !reflect.DeepEqual(table.out, a.GetEPANASMessageContainer()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetEPANASMessageContainer())
		}
	}
}

type testEPSNASMessageContainerDataTemplate struct {
	in  nasType.EPSNASMessageContainer
	out nasType.EPSNASMessageContainer
}

var ePSNASMessageContainerTestData = []nasType.EPSNASMessageContainer{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, 3, []byte{0x02, 0x1f, 0x22}},
}

var ePSNASMessageContainerExpectedData = []nasType.EPSNASMessageContainer{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, 3, []byte{0x02, 0x1f, 0x22}},
}

var ePSNASMessageContainerTestTable = []testEPSNASMessageContainerDataTemplate{
	{ePSNASMessageContainerTestData[0], ePSNASMessageContainerExpectedData[0]},
}

func TestNasTypeEPSNASMessageContainer(t *testing.T) {
	for i, table := range ePSNASMessageContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewEPSNASMessageContainer(0)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetEPANASMessageContainer(table.in.Buffer)

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
