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

func TestNasTypeNewSequenceNumber(t *testing.T) {
	a := nasType.NewSequenceNumber()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypeSequenceNumber struct {
	in  uint8
	out uint8
}

var nasTypeSequenceNumberTable = []nasTypeSequenceNumber{
	{0x03, 0x03},
}

func TestNasTypeSequenceNumberGetSetSQN(t *testing.T) {
	a := nasType.NewSequenceNumber()
	for _, table := range nasTypeSequenceNumberTable {
		a.SetSQN(table.in)
		if !reflect.DeepEqual(table.out, a.GetSQN()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetSQN())
		}
	}
}

type SequenceNumberTestDataTemplate struct {
	in  nasType.SequenceNumber
	out nasType.SequenceNumber
}

var SequenceNumberTestData = []nasType.SequenceNumber{
	{0x03},
}

var SequenceNumberExpectedTestData = []nasType.SequenceNumber{
	{0x03},
}

var SequenceNumberTable = []SequenceNumberTestDataTemplate{
	{SequenceNumberTestData[0], SequenceNumberExpectedTestData[0]},
}

func TestNasTypeSequenceNumber(t *testing.T) {
	for _, table := range SequenceNumberTable {
		a := nasType.NewSequenceNumber()

		a.SetSQN(table.in.GetSQN())
		if !reflect.DeepEqual(table.out.GetSQN(), a.GetSQN()) {
			t.Errorf("Not equal: expected %v, got %v", table.out.GetSQN(), a.GetSQN())
		}
	}
}
