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

func TestNasTypeNewIMEISV(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeSecurityModeCompleteIMEISVTypeTable = []NasTypeIeiData{
	{nasMessage.SecurityModeCompleteIMEISVType, nasMessage.SecurityModeCompleteIMEISVType},
}

func TestNasTypeIMEISVGetSetIei(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeSecurityModeCompleteIMEISVTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeIMEISVLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeIMEISVGetSetLen(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeIMEISVIdentityDigit1 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeIMEISVIdentityDigit1Table = []nasTypeIMEISVIdentityDigit1{
	{2, 0x01, 0x01},
}

func TestNasTypeIMEISVGetSetIdentityDigit1(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVIdentityDigit1Table {
		a.SetLen(table.inLen)
		a.SetIdentityDigit1(table.in)
		if !reflect.DeepEqual(table.out, a.GetIdentityDigit1()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIdentityDigit1())
		}
	}
}

type nasTypeIMEISVOddEvenIdic struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeIMEISVOddEvenIdicTable = []nasTypeIMEISVOddEvenIdic{
	{2, 0x01, 0x01},
}

func TestNasTypeIMEISVGetSetOddEvenIdic(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVOddEvenIdicTable {
		a.SetLen(table.inLen)
		a.SetOddEvenIdic(table.in)
		if !reflect.DeepEqual(table.out, a.GetOddEvenIdic()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetOddEvenIdic())
		}
	}
}

type nasTypeIMEISVTypeOfIdentity struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeIMEISVTypeOfIdentityTable = []nasTypeIMEISVTypeOfIdentity{
	{2, 0x07, 0x07},
}

func TestNasTypeIMEISVGetSetTypeOfIdentity(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVTypeOfIdentityTable {
		a.SetLen(table.inLen)
		a.SetTypeOfIdentity(table.in)
		if !reflect.DeepEqual(table.out, a.GetTypeOfIdentity()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTypeOfIdentity())
		}
	}
}

type nasTypeIMEISVIdentityDigitP_1 struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeIMEISVIdentityDigitP_1Table = []nasTypeIMEISVIdentityDigitP_1{
	{2, 0x01, 0x01},
}

func TestNasTypeIMEISVGetSetIdentityDigitP_1(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVIdentityDigitP_1Table {
		a.SetLen(table.inLen)
		a.SetIdentityDigitP_1(table.in)
		if !reflect.DeepEqual(table.out, a.GetIdentityDigitP_1()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIdentityDigitP_1())
		}
	}
}

type nasTypeIMEISVGetIdentityDigitP struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeIMEISVGetIdentityDigitPTable = []nasTypeIMEISVGetIdentityDigitP{
	{2, 0x0f, 0x0f},
}

func TestNasTypeIMEISVGetSetGetIdentityDigitP(t *testing.T) {
	a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
	for _, table := range nasTypeIMEISVGetIdentityDigitPTable {
		a.SetLen(table.inLen)
		a.SetIdentityDigitP(table.in)
		if !reflect.DeepEqual(table.out, a.GetIdentityDigitP()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIdentityDigitP())
		}
	}
}

type testIMEISVDataTemplate struct {
	inIei              uint8
	inLen              uint16
	inIdentityDigit1   uint8
	inOddEvenIdic      uint8
	inTypeOfIdentity   uint8
	inIdentityDigitP_1 uint8
	inIdentityDigitP   uint8

	outIei              uint8
	outLen              uint16
	outIdentityDigit1   uint8
	outOddEvenIdic      uint8
	outTypeOfIdentity   uint8
	outIdentityDigitP_1 uint8
	outIdentityDigitP   uint8
}

var iMEISVTestTable = []testIMEISVDataTemplate{
	{
		nasMessage.SecurityModeCompleteIMEISVType, 2, 0x01, 0x01, 0x01, 0x01, 0x01,
		nasMessage.SecurityModeCompleteIMEISVType, 2, 0x01, 0x01, 0x01, 0x01, 0x01,
	},
}

func TestNasTypeIMEISV(t *testing.T) {
	for i, table := range iMEISVTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetIdentityDigit1(table.inIdentityDigit1)
		a.SetOddEvenIdic(table.inOddEvenIdic)
		a.SetTypeOfIdentity(table.inTypeOfIdentity)
		a.SetIdentityDigitP_1(table.inIdentityDigitP_1)
		a.SetIdentityDigitP(table.inIdentityDigitP)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outLen, a.GetLen()) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.GetLen())
		}
		if !reflect.DeepEqual(table.outIdentityDigit1, a.GetIdentityDigit1()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIdentityDigit1, table.outIdentityDigit1, a.GetIdentityDigit1())
		}
		if !reflect.DeepEqual(table.outOddEvenIdic, a.GetOddEvenIdic()) {
			t.Errorf("in(%v): out %v, actual %x", table.inOddEvenIdic, table.outOddEvenIdic, a.GetOddEvenIdic())
		}
		if !reflect.DeepEqual(table.outTypeOfIdentity, a.GetTypeOfIdentity()) {
			t.Errorf("in(%v): out %v, actual %x", table.inTypeOfIdentity, table.outTypeOfIdentity, a.GetTypeOfIdentity())
		}
		if !reflect.DeepEqual(table.outIdentityDigitP_1, a.GetIdentityDigitP_1()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIdentityDigitP_1, table.outIdentityDigitP_1, a.GetIdentityDigitP_1())
		}
		if !reflect.DeepEqual(table.outIdentityDigitP, a.GetIdentityDigitP()) {
			t.Errorf("in(%v): out %v, actual %x", table.inIdentityDigitP, table.outIdentityDigitP, a.GetIdentityDigitP())
		}
	}
}
