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

func TestNasTypeNewCause5GMM(t *testing.T) {
	a := nasType.NewCause5GMM(nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeDeregistrationRequestUETerminatedDeregistrationCause5GMMTable = []NasTypeIeiData{
	{nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType, nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType},
}

func TestNasTypeCause5GMMGetSetIei(t *testing.T) {
	a := nasType.NewCause5GMM(nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
	for _, table := range nasTypeDeregistrationRequestUETerminatedDeregistrationCause5GMMTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeCause5GMMCauseValueData struct {
	in  uint8
	out uint8
}

var nasTypeCause5GMMOctetTable = []nasTypeCause5GMMCauseValueData{
	{0xff, 0xff},
}

func TestNasTypeCause5GMMGetSetCauseValue(t *testing.T) {
	a := nasType.NewCause5GMM(nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
	for _, table := range nasTypeCause5GMMOctetTable {
		a.SetCauseValue(table.in)
		if !reflect.DeepEqual(table.out, a.GetCauseValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetCauseValue())
		}
	}
}

type testCause5GMMDataTemplate struct {
	in  nasType.Cause5GMM
	out nasType.Cause5GMM
}

var cause5GMMTestData = []nasType.Cause5GMM{
	{nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType, 0xff},
}

var cause5GMMExpectedTestData = []nasType.Cause5GMM{
	{nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType, 0xff},
}

var cause5GMMTestTable = []testCause5GMMDataTemplate{
	{cause5GMMTestData[0], cause5GMMExpectedTestData[0]},
}

func TestNasTypeCause5GMM(t *testing.T) {
	for i, table := range cause5GMMTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewCause5GMM(nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)

		a.SetIei(table.in.GetIei())
		a.SetCauseValue(table.in.Octet)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Octet, a.Octet) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)
		}
	}
}
