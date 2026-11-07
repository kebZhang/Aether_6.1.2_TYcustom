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

func TestNasTypeNewNon3GppDeregistrationTimerValue(t *testing.T) {
	a := nasType.NewNon3GppDeregistrationTimerValue(nasMessage.ServiceRejectT3346ValueType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeNon3GppDeregistrationTimerValueServiceRejectT3346ValueTypeTable = []NasTypeIeiData{
	{nasMessage.ServiceRejectT3346ValueType, nasMessage.ServiceRejectT3346ValueType},
}

func TestNasTypeNon3GppDeregistrationTimerValueGetSetIei(t *testing.T) {
	a := nasType.NewNon3GppDeregistrationTimerValue(nasMessage.ServiceRejectT3346ValueType)
	for _, table := range nasTypeNon3GppDeregistrationTimerValueServiceRejectT3346ValueTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeNon3GppDeregistrationTimerValueLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeNon3GppDeregistrationTimerValueGetSetLen(t *testing.T) {
	a := nasType.NewNon3GppDeregistrationTimerValue(nasMessage.ServiceRejectT3346ValueType)
	for _, table := range nasTypeNon3GppDeregistrationTimerValueLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypeNon3GppDeregistrationTimerValueGPRSTimer2ValueData struct {
	in  uint8
	out uint8
}

var nasTypeNon3GppDeregistrationTimerValueGPRSTimer2ValueTable = []nasTypeNon3GppDeregistrationTimerValueGPRSTimer2ValueData{
	{0x0f, 0x0f},
}

func TestNasTypeNon3GppDeregistrationTimerValueGetSetGPRSTimer2Value(t *testing.T) {
	a := nasType.NewNon3GppDeregistrationTimerValue(nasMessage.ServiceRejectT3346ValueType)
	for _, table := range nasTypeNon3GppDeregistrationTimerValueGPRSTimer2ValueTable {
		a.SetGPRSTimer2Value(table.in)
		if !reflect.DeepEqual(table.out, a.GetGPRSTimer2Value()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetGPRSTimer2Value())
		}
	}
}

type testNon3GppDeregistrationTimerValueDataTemplate struct {
	inIei              uint8
	inLen              uint8
	inGPRSTimer2Value  uint8
	outIei             uint8
	outLen             uint8
	outGPRSTimer2Value uint8
}

var testNon3GppDeregistrationTimerValueTestTable = []testNon3GppDeregistrationTimerValueDataTemplate{
	{
		nasMessage.ServiceRejectT3346ValueType, 2, 0x0f,
		nasMessage.ServiceRejectT3346ValueType, 2, 0x0f,
	},
}

func TestNasTypeNon3GppDeregistrationTimerValue(t *testing.T) {
	for i, table := range testNon3GppDeregistrationTimerValueTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewNon3GppDeregistrationTimerValue(nasMessage.ServiceRejectT3346ValueType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetGPRSTimer2Value(table.inGPRSTimer2Value)

		if !reflect.DeepEqual(table.outIei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		}
		if !reflect.DeepEqual(table.outLen, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		}
		if !reflect.DeepEqual(table.outGPRSTimer2Value, a.GetGPRSTimer2Value()) {
			t.Errorf("in(%v): out %v, actual %x", table.inGPRSTimer2Value, table.outGPRSTimer2Value, a.GetGPRSTimer2Value())
		}
	}
}
