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

type nasTypeNgksiAndServiceTypeData struct {
	inTsc                  uint8
	outTsc                 uint8
	inNASKeySetIdentifier  uint8
	outNASKeySetIdentifier uint8
	inServiceTypeValue     uint8
	outServiceTypeValue    uint8
}

var nasTypeNgksiAndServiceTypeTable = []nasTypeNgksiAndServiceTypeData{
	{0x01, 0x01, 0x07, 0x07, 0x7, 0x07},
}

func TestNasTypeNewServiceTypeAndNgksi(t *testing.T) {
	a := nasType.NewServiceTypeAndNgksi()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetNgksiAndServiceType(t *testing.T) {
	a := nasType.NewServiceTypeAndNgksi()
	for _, table := range nasTypeNgksiAndServiceTypeTable {
		a.SetTSC(table.inTsc)
		if !reflect.DeepEqual(table.outTsc, a.GetTSC()) {
			t.Errorf("Not equal: expected %v, got %v", table.outTsc, a.GetTSC())
		}
		// a.SetTSC(0)
		a.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		if !reflect.DeepEqual(table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler())
		}

		a.SetServiceTypeValue(table.inServiceTypeValue)
		if !reflect.DeepEqual(table.outServiceTypeValue, a.GetServiceTypeValue()) {
			t.Errorf("Not equal: expected %v, got %v", table.outServiceTypeValue, a.GetServiceTypeValue())
		}
	}
}
