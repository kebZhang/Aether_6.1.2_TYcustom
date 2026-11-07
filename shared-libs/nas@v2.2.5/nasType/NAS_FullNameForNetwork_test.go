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

func TestNasTypeNewFullNameForNetwork(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandFullNameForNetworkIeiTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandFullNameForNetworkType, nasMessage.ConfigurationUpdateCommandFullNameForNetworkType},
}

func TestNasTypeFullNameForNetworkGetSetIei(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeConfigurationUpdateCommandFullNameForNetworkIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeFullNameForNetworkLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeFullNameForNetworkGetSetLen(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypetFullNameForNetworkExt struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeFullNameForNetworkExtTable = []nasTypetFullNameForNetworkExt{
	{2, 0x01, 0x01},
}

func TestNasTypeFullNameForNetworkGetSetExt(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkExtTable {
		a.SetLen(table.inLen)
		a.SetExt(table.in)
		if !reflect.DeepEqual(table.out, a.GetExt()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetExt())
		}
	}
}

type nasTypetFullNameForNetworkCodingScheme struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeFullNameForNetworkCodingSchemeTable = []nasTypetFullNameForNetworkCodingScheme{
	{2, 0x07, 0x07},
}

func TestNasTypeFullNameForNetworkGetSetCodingScheme(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkCodingSchemeTable {
		a.SetLen(table.inLen)
		a.SetCodingScheme(table.in)
		if !reflect.DeepEqual(table.out, a.GetCodingScheme()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetCodingScheme())
		}
	}
}

type nasTypetFullNameForNetworkAddCI struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeFullNameForNetworkAddCITable = []nasTypetFullNameForNetworkAddCI{
	{2, 0x01, 0x01},
}

func TestNasTypeFullNameForNetworkGetSetAddCI(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkAddCITable {
		a.SetLen(table.inLen)
		a.SetAddCI(table.in)
		if !reflect.DeepEqual(table.out, a.GetAddCI()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetAddCI())
		}
	}
}

type nasTypetFullNameForNetworkNumberOfSpareBitsInLastOctet struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeFullNameForNetworkNumberOfSpareBitsInLastOctetTable = []nasTypetFullNameForNetworkNumberOfSpareBitsInLastOctet{
	{2, 0x07, 0x07},
}

func TestNasTypeFullNameForNetworkGetSetNumberOfSpareBitsInLastOctet(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkNumberOfSpareBitsInLastOctetTable {
		a.SetLen(table.inLen)
		a.SetNumberOfSpareBitsInLastOctet(table.in)
		if !reflect.DeepEqual(table.out, a.GetNumberOfSpareBitsInLastOctet()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetNumberOfSpareBitsInLastOctet())
		}
	}
}

type nasTypetFullNameForNetworkTextString struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeFullNameForNetworkTextStringTable = []nasTypetFullNameForNetworkTextString{
	{3, []uint8{0x07, 0x07}, []uint8{0x07, 0x07}},
}

func TestNasTypeFullNameForNetworkGetSetTextString(t *testing.T) {
	a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	for _, table := range nasTypeFullNameForNetworkTextStringTable {
		a.SetLen(table.inLen)
		a.SetTextString(table.in)
		if !reflect.DeepEqual(table.out, a.GetTextString()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetTextString())
		}
	}
}

type testFullNameForNetworkDataTemplate struct {
	inIei                           uint8
	inLen                           uint8
	inExt                           uint8
	inCodingScheme                  uint8
	inAddCI                         uint8
	inNumberOfSpareBitsInLastOctet  uint8
	inTextString                    []uint8
	outIei                          uint8
	outLen                          uint8
	outExt                          uint8
	outCodingScheme                 uint8
	outAddCI                        uint8
	outNumberOfSpareBitsInLastOctet uint8
	outTextString                   []uint8
}

var fullNameForNetworkestTable = []testFullNameForNetworkDataTemplate{
	{nasMessage.ConfigurationUpdateCommandFullNameForNetworkType, 3, 0x01, 0x01, 0x01, 0x01, []uint8{0x01, 0x01}, nasMessage.ConfigurationUpdateCommandFullNameForNetworkType, 3, 0x01, 0x01, 0x01, 0x01, []uint8{0x01, 0x01}},
}

func TestNasTypeFullNameForNetwork(t *testing.T) {
	for i, table := range fullNameForNetworkestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetExt(table.inExt)
		a.SetCodingScheme(table.inCodingScheme)
		a.SetAddCI(table.inAddCI)
		a.SetNumberOfSpareBitsInLastOctet(table.inNumberOfSpareBitsInLastOctet)
		a.SetTextString(table.inTextString)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outExt, a.GetExt()) {
			t.Errorf("in(%v): out %v, actual %x", table.inExt, table.outExt, a.GetExt())
		}
		if !reflect.DeepEqual(table.outCodingScheme, a.GetCodingScheme()) {
			t.Errorf("in(%v): out %v, actual %x", table.inCodingScheme, table.outCodingScheme, a.GetCodingScheme())
		}
		if !reflect.DeepEqual(table.outAddCI, a.GetAddCI()) {
			t.Errorf("in(%v): out %v, actual %x", table.inAddCI, table.outAddCI, a.GetAddCI())
		}
		if !reflect.DeepEqual(table.outNumberOfSpareBitsInLastOctet, a.GetNumberOfSpareBitsInLastOctet()) {
			t.Errorf("in(%v): out %v, actual %x", table.inNumberOfSpareBitsInLastOctet, table.outNumberOfSpareBitsInLastOctet, a.GetNumberOfSpareBitsInLastOctet())
		}
		if !reflect.DeepEqual(table.outTextString, a.GetTextString()) {
			t.Errorf("in(%v): out %v, actual %x", table.inTextString, table.outTextString, a.GetTextString())
		}
	}
}
