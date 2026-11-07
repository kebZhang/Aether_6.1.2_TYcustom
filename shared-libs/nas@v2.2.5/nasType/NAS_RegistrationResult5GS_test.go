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

func TestNasTypeNewRegistrationResult5GS(t *testing.T) {
	a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType, nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType},
}

func TestNasTypeRegistrationResult5GSGetSetIei(t *testing.T) {
	a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeRegistrationResult5GSGetSetLen(t *testing.T) {
	a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type NasTypeSMSAlloweduint8Data struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationResult5GSSMSAllowed = []NasTypeSMSAlloweduint8Data{
	{0x01, 0x01},
	// {0x0, 0x0},
}

func TestNasTypeRegistrationResult5GSGetSetSMSAllowed(t *testing.T) {
	a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
	for _, table := range nasTypeRegistrationResult5GSSMSAllowed {
		a.SetSMSAllowed(table.in)
		if !reflect.DeepEqual(table.out, a.GetSMSAllowed()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSMSAllowed())
		}
	}
}

type NasTypeRegistrationResultValue5GSuint8Data struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationResult5GSRegistrationResultValue5GS = []NasTypeRegistrationResultValue5GSuint8Data{
	{0x1, 0x1},
	{0x0, 0x0},
}

func TestNasTypeRegistrationResult5GSGetSetRegistrationResultValue5GS(t *testing.T) {
	a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
	for _, table := range nasTypeRegistrationResult5GSRegistrationResultValue5GS {
		a.SetRegistrationResultValue5GS(table.in)
		if !reflect.DeepEqual(table.out, a.GetRegistrationResultValue5GS()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetRegistrationResultValue5GS())
		}
	}
}

type testRegistrationResult5GSDataTemplate struct {
	inSMSAllowed                 uint8
	inRegistrationResultValue5GS uint8
	in                           nasType.RegistrationResult5GS
	out                          nasType.RegistrationResult5GS
}

var registrationResult5GSTestData = []nasType.RegistrationResult5GS{
	{nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType, 1, 0x05},
}

var registrationResult5GSExpectedData = []nasType.RegistrationResult5GS{
	{nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType, 1, 0x0f},
}

var registrationResult5GSDataTestTable = []testRegistrationResult5GSDataTemplate{
	{0x07, 0x1F, registrationResult5GSTestData[0], registrationResult5GSExpectedData[0]},
}

func TestNasTypeRegistrationResult5GS(t *testing.T) {
	for _, table := range registrationResult5GSDataTestTable {
		a := nasType.NewRegistrationResult5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetSMSAllowed(table.inSMSAllowed)
		a.SetRegistrationResultValue5GS(table.inRegistrationResultValue5GS)
		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("Not equal: expected %v, got %v", table.out.Octet, a.Octet)
		}
	}
}
