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

var RegistrationAcceptNSSAIInclusionModeTypeIeiInput uint8 = 0x0A

func TestNasTypeNewNSSAIInclusionMode(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeNSSAIInclusionModeRegistrationAcceptNSSAIInclusionModeTypeTable = []NasTypeIeiData{
	{RegistrationAcceptNSSAIInclusionModeTypeIeiInput, 0x0A},
}

func TestNasTypeNSSAIInclusionModeGetSetIei(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	for _, table := range nasTypeNSSAIInclusionModeRegistrationAcceptNSSAIInclusionModeTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeNSSAIInclusionMode struct {
	inIei                 uint8
	inNSSAIInclusionMode  uint8
	outIei                uint8
	outNSSAIInclusionMode uint8
}

var nasTypeNSSAIInclusionModeTable = []nasTypeNSSAIInclusionMode{
	{
		RegistrationAcceptNSSAIInclusionModeTypeIeiInput, 0x03,
		0x0A, 0x03,
	},
}

func TestNasTypeNSSAIInclusionMode(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	for _, table := range nasTypeNSSAIInclusionModeTable {
		a.SetNSSAIInclusionMode(table.inNSSAIInclusionMode)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outNSSAIInclusionMode, a.GetNSSAIInclusionMode()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNSSAIInclusionMode, a.GetNSSAIInclusionMode())
		}
	}
}
