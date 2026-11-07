// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package security

import (
	"testing"
)

func TestSetterGetter(t *testing.T) {
	testCases := []struct {
		overflow uint16
		sqn      uint8
	}{
		{1, 2},
		{0, 0},
		{170, 35},
		{65535, 255},
	}

	count := Count{}

	for _, testCase := range testCases {
		count.Set(testCase.overflow, testCase.sqn)
		expected := (uint32(testCase.overflow) << 8) + uint32(testCase.sqn)
		if count.Get() != expected {
			t.Errorf("Get() Failed: expected %v, got %v", expected, count.Get())
		}
		if count.Overflow() != testCase.overflow {
			t.Errorf("Overflow() Failed: expected %v, got %v", testCase.overflow, count.Overflow())
		}
		if count.SQN() != testCase.sqn {
			t.Errorf("SQN() Failed: expected %v, got %v", testCase.sqn, count.SQN())
		}
	}
}

func TestAddOne(t *testing.T) {
	count := Count{}

	count.Set(0, 0)

	for i := range 4567 {
		count.AddOne()
		if count.Get() != uint32(i+1) {
			t.Errorf("AddOne() Test Failed: expected %v, got %v", uint32(i+1), count.Get())
		}
	}
}
