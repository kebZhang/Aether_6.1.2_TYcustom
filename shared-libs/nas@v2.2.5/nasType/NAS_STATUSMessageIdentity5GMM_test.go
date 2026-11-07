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

func TestNasTypeNewSTATUSMessageIdentity5GMM(t *testing.T) {
	a := nasType.NewSTATUSMessageIdentity5GMM()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeSTATUSMessageIdentity5GMM struct {
	in  uint8
	out uint8
}

var nasTypeSTATUSMessageIdentity5GMMTable = []nasTypeSTATUSMessageIdentity5GMM{
	{0x03, 0x03},
}

func TestNasTypeSTATUSMessageIdentity5GMMGetSetMessageType(t *testing.T) {
	a := nasType.NewSTATUSMessageIdentity5GMM()
	for _, table := range nasTypeSTATUSMessageIdentity5GMMTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}

type STATUSMessageIdentity5GMMTestDataTemplate struct {
	in  nasType.STATUSMessageIdentity5GMM
	out nasType.STATUSMessageIdentity5GMM
}

var STATUSMessageIdentity5GMMTestData = []nasType.STATUSMessageIdentity5GMM{
	{0x03},
}

var STATUSMessageIdentity5GMMExpectedTestData = []nasType.STATUSMessageIdentity5GMM{
	{0x03},
}

var STATUSMessageIdentity5GMMTable = []STATUSMessageIdentity5GMMTestDataTemplate{
	{STATUSMessageIdentity5GMMTestData[0], STATUSMessageIdentity5GMMExpectedTestData[0]},
}

func TestNasTypeSTATUSMessageIdentity5GMM(t *testing.T) {
	for _, table := range STATUSMessageIdentity5GMMTable {
		a := nasType.NewSTATUSMessageIdentity5GMM()

		a.SetMessageType(table.in.GetMessageType())
		if !reflect.DeepEqual(table.out.GetMessageType(), a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetMessageType(), a.GetMessageType())
		}
	}
}
