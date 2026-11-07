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

var ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput uint8 = 0x09

func TestNasTypeNewNetworkSlicingIndication(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeConfigurationUpdateCommandNetworkSlicingIndicationTable = []NasTypeIeiData{
	{ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput, 0x09},
}

func TestNasTypeNetworkSlicingIndicationGetSetIei(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateCommandNetworkSlicingIndicationTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeNetworkSlicingIndication struct {
	inDCNI   uint8
	outDCNI  uint8
	inNSSCI  uint8
	outNSSCI uint8
	outIei   uint8
}

var nasTypeNetworkSlicingIndicationTable = []nasTypeNetworkSlicingIndication{
	{0x01, 0x01, 0x01, 0x01, 0x09},
}

func TestNasTypeNetworkSlicingIndication(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	for _, table := range nasTypeNetworkSlicingIndicationTable {
		a.SetDCNI(table.inDCNI)
		a.SetNSSCI(table.inNSSCI)

		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outDCNI, a.GetDCNI()) {
			t.Errorf("Not equal: expected %v, got %v", table.outDCNI, a.GetDCNI())
		}
		if !reflect.DeepEqual(table.outNSSCI, a.GetNSSCI()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNSSCI, a.GetNSSCI())
		}
	}
}
