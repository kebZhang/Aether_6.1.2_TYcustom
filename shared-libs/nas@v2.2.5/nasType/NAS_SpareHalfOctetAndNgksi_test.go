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

type nasTypeNewNgksiAndSpareHalfOctetData struct {
	inTsc                  uint8
	outTsc                 uint8
	inNASKeySetIdentifier  uint8
	outNASKeySetIdentifier uint8
	inSpareHalfOctet       uint8
	outSpareHalfOctet      uint8
}

var nasTypeNewNgksiAndSpareHalfOctetTable = []nasTypeNewNgksiAndSpareHalfOctetData{
	{0x1, 0x1, 0x7, 0x7, 0x7, 0x7},
}

func TestNasTypeNewSpareHalfOctetAndNgksi(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndNgksi()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetSpareHalfOctetAndNgksi(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndNgksi()
	for _, table := range nasTypeNewNgksiAndSpareHalfOctetTable {
		a.SetTSC(table.inTsc)
		if !reflect.DeepEqual(table.outTsc, a.GetTSC()) {
			t.Errorf("Not equal: expected %v, got %v", table.outTsc, a.GetTSC())
		}
		a.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		if !reflect.DeepEqual(table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler()) {
			t.Errorf("Not equal: expected %v, got %v", table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler())
		}

		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		if !reflect.DeepEqual(table.outSpareHalfOctet, a.GetSpareHalfOctet()) {
			t.Errorf("Not equal: expected %v, got %v", table.outSpareHalfOctet, a.GetSpareHalfOctet())
		}
	}
}
