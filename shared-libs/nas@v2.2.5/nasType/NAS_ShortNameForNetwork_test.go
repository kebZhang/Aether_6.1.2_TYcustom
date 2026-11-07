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

func TestNasTypeNewShortNameForNetwork(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeShortNameForNetworkTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandShortNameForNetworkType, nasMessage.ConfigurationUpdateCommandShortNameForNetworkType},
}

func TestNasTypeShortNameForNetworkGetSetIei(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeShortNameForNetworkLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeShortNameForNetworkGetSetLen(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeShortNameForNetworkExtData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeShortNameForNetworkExtTable = []nasTypeShortNameForNetworkExtData{
	{2, 0x01, 0x01},
}

func TestNasTypeShortNameForNetworkGetSetExt(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkExtTable {
		a.SetLen(table.inLen) // fix it, set input length
		a.SetExt(table.in)
		if !reflect.DeepEqual(table.out, a.GetExt()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetExt())
		}
	}
}

type nasTypeShortNameForNetworkCodingSchemeData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeShortNameForNetworkCodingSchemeTable = []nasTypeShortNameForNetworkCodingSchemeData{
	{2, 0x01, 0x01},
}

func TestNasTypeShortNameForNetworkGetSetCodingScheme(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkCodingSchemeTable {
		a.SetLen(table.inLen)
		a.SetCodingScheme(table.in)
		if !reflect.DeepEqual(table.out, a.GetCodingScheme()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetCodingScheme())
		}
	}
}

type nasTypeShortNameForNetworkAddCIData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeShortNameForNetworkAddCITable = []nasTypeShortNameForNetworkAddCIData{
	{2, 0x01, 0x01},
}

func TestNasTypeShortNameForNetworkGetSetAddCI(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkAddCITable {
		a.SetLen(table.inLen)
		a.SetAddCI(table.in)
		if !reflect.DeepEqual(table.out, a.GetAddCI()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetAddCI())
		}
	}
}

type nasTypeShortNameForNetworkNumberOfSpareBitsInLastOctettData struct {
	inLen uint8
	in    uint8
	out   uint8
}

var nasTypeShortNameForNetworkNumberOfSpareBitsInLastOctetTable = []nasTypeShortNameForNetworkNumberOfSpareBitsInLastOctettData{
	{2, 0x01, 0x01},
}

func TestNasTypeShortNameForNetworkGetSetNumberOfSpareBitsInLastOctet(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkNumberOfSpareBitsInLastOctetTable {
		a.SetLen(table.inLen)
		a.SetNumberOfSpareBitsInLastOctet(table.in)
		if !reflect.DeepEqual(table.out, a.GetNumberOfSpareBitsInLastOctet()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetNumberOfSpareBitsInLastOctet())
		}
	}
}

type nasTypeShortNameForNetworkTextStringData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeShortNameForNetworkTextStringTable = []nasTypeShortNameForNetworkTextStringData{
	{3, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeShortNameForNetworkGetSetTextString(t *testing.T) {
	a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	for _, table := range nasTypeShortNameForNetworkTextStringTable {
		a.SetLen(table.inLen)
		a.SetTextString(table.in)
		if !reflect.DeepEqual(table.out, a.GetTextString()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetTextString())
		}
	}
}

type testShortNameForNetworkDataTemplate struct {
	in  nasType.ShortNameForNetwork
	out nasType.ShortNameForNetwork
}

var ShortNameForNetworkTestData = []nasType.ShortNameForNetwork{
	{nasMessage.ConfigurationUpdateCommandShortNameForNetworkType, 3, []uint8{}},
}

var ShortNameForNetworkExpectedTestData = []nasType.ShortNameForNetwork{
	{nasMessage.ConfigurationUpdateCommandShortNameForNetworkType, 3, []uint8{0x99, 0x01, 0x01}},
}

var ShortNameForNetworkTestTable = []testShortNameForNetworkDataTemplate{
	{ShortNameForNetworkTestData[0], ShortNameForNetworkExpectedTestData[0]},
}

func TestNasTypeShortNameForNetwork(t *testing.T) {
	for i, table := range ShortNameForNetworkTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetExt(0x01)
		a.SetCodingScheme(0x01)
		a.SetAddCI(0x01)
		a.SetNumberOfSpareBitsInLastOctet(0x01)
		a.SetTextString([]uint8{0x01, 0x01})

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
