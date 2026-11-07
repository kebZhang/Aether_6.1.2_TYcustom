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

var RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput uint8 = 0x0C

func TestNasTypeNewNoncurrentNativeNASKeySetIdentifier(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandNoncurrentNativeNASKeySetIdentifierTable = []NasTypeIeiData{
	{RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput, 0x0C},
}

func TestNasTypeNoncurrentNativeNASKeySetIdentifierGetSetIei(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateCommandNoncurrentNativeNASKeySetIdentifierTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeNoncurrentNativeNASKeySetIdentifier struct {
	inIei                   uint8
	inTsc                   uint8
	inNasKeySetIdentifiler  uint8
	outIei                  uint8
	outTsc                  uint8
	outNasKeySetIdentifiler uint8
}

var nasTypeNoncurrentNativeNASKeySetIdentifierTable = []nasTypeNoncurrentNativeNASKeySetIdentifier{
	{
		RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput, 0x01, 0x01,
		0x0C, 0x01, 0x01,
	},
}

func TestNasTypeNoncurrentNativeNASKeySetIdentifier(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	for _, table := range nasTypeNoncurrentNativeNASKeySetIdentifierTable {
		a.SetTsc(table.inTsc)
		a.SetNasKeySetIdentifiler(table.inNasKeySetIdentifiler)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outTsc, a.GetTsc()) {
			t.Errorf("Not equal: expected %v, got %v", table.outTsc, a.GetTsc())
		}
		if !reflect.DeepEqual(table.outNasKeySetIdentifiler, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNasKeySetIdentifiler, a.GetNasKeySetIdentifiler())
		}
	}
}
