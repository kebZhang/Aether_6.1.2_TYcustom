// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewNgksiAndRegistrationType5GS(t *testing.T) {
	a := nasType.NewNgksiAndRegistrationType5GS()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var RegistrationType5GSAndNgksiFORTable = []NasTypeLenuint8Data{
	{0x1, 0x1},
}

func TestNasTypeRegistrationType5GSAndNgksiGetSetFOR(t *testing.T) {
	a := nasType.NewNgksiAndRegistrationType5GS()
	for _, table := range RegistrationType5GSAndNgksiFORTable {
		a.SetFOR(table.in)
		if !reflect.DeepEqual(table.out, a.GetFOR()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetFOR())
		}
	}
}

var RegistrationType5GSAndNgksiRegistrationType5GSTable = []NasTypeLenuint8Data{
	{0x07, 0x7},
}

func TestNasTypeRegistrationType5GSAndNgksiGetSetRegistrationType5GS(t *testing.T) {
	a := nasType.NewNgksiAndRegistrationType5GS()
	for _, table := range RegistrationType5GSAndNgksiRegistrationType5GSTable {
		a.SetRegistrationType5GS(table.in)
		if !reflect.DeepEqual(table.out, a.GetRegistrationType5GS()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRegistrationType5GS())
		}
	}
}

var RegistrationType5GSAndNgksiTSCTable = []NasTypeLenuint8Data{
	{0x1, 0x01},
}

func TestNasTypeRegistrationType5GSAndNgksiGetSetTSC(t *testing.T) {
	a := nasType.NewNgksiAndRegistrationType5GS()
	for _, table := range RegistrationType5GSAndNgksiTSCTable {
		a.SetTSC(table.in)
		if !reflect.DeepEqual(table.out, a.GetTSC()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetTSC())
		}
	}
}

var RegistrationType5GSAndNgksiNasKeySetIdentifilerTable = []NasTypeLenuint8Data{
	{0x04, 0x04},
}

func TestNasTypeRegistrationType5GSAndNgksiGetSetNasKeySetIdentifiler(t *testing.T) {
	a := nasType.NewNgksiAndRegistrationType5GS()
	for _, table := range RegistrationType5GSAndNgksiNasKeySetIdentifilerTable {
		a.SetNasKeySetIdentifiler(table.in)
		if !reflect.DeepEqual(table.out, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetNasKeySetIdentifiler())
		}
	}
}

type testRegistrationType5GSAndNgksiDataTemplate struct {
	inFOR                  uint8
	inRegistrationType5GS  uint8
	inTSC                  uint8
	inNasKeySetIdentifiler uint8
	in                     nasType.NgksiAndRegistrationType5GS
	out                    nasType.NgksiAndRegistrationType5GS
}

var registrationType5GSAndNgksiTestData = []nasType.NgksiAndRegistrationType5GS{
	{0x01},
}

var registrationType5GSAndNgksiExpectedTestData = []nasType.NgksiAndRegistrationType5GS{
	{0x99},
}

var registrationType5GSAndNgksiTestTable = []testRegistrationType5GSAndNgksiDataTemplate{
	{0x1, 0x1, 0x1, 0x1, registrationType5GSAndNgksiTestData[0], registrationType5GSAndNgksiExpectedTestData[0]},
}

func TestNasTypeRegistrationType5GSAndNgksi(t *testing.T) {
	for _, table := range registrationType5GSAndNgksiTestTable {
		a := nasType.NewNgksiAndRegistrationType5GS()

		a.SetFOR(table.inFOR)
		a.SetRegistrationType5GS(table.inRegistrationType5GS)
		a.SetTSC(table.inTSC)
		a.SetNasKeySetIdentifiler(table.inNasKeySetIdentifiler)

		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
