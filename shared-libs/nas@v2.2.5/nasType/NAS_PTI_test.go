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

func TestNasTypeNewPTI(t *testing.T) {
	a := nasType.NewPTI()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePTI struct {
	in  uint8
	out uint8
}

var nasTypePTITable = []nasTypePTI{
	{0x03, 0x03},
}

func TestNasTypePTIGetSetPDUSessionIdentity(t *testing.T) {
	a := nasType.NewPTI()
	for _, table := range nasTypePTITable {
		a.SetPTI(table.in)
		if !reflect.DeepEqual(table.out, a.GetPTI()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetPTI())
		}
	}
}

type PTITestDataTemplate struct {
	in  nasType.PTI
	out nasType.PTI
}

var PTITestData = []nasType.PTI{
	{0x03},
}

var PTIExpectedTestData = []nasType.PTI{
	{0x03},
}

var PTITable = []PTITestDataTemplate{
	{PTITestData[0], PTIExpectedTestData[0]},
}

func TestNasTypePTI(t *testing.T) {
	for _, table := range PTITable {
		a := nasType.NewPTI()

		a.SetPTI(table.in.GetPTI())
		if !reflect.DeepEqual(table.out.GetPTI(), a.GetPTI()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetPTI(), a.GetPTI())
		}
	}
}
