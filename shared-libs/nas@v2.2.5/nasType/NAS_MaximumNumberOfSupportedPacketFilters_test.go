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

func TestNasTypeNewMaximumNumberOfSupportedPacketFilters(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypePDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType, nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType},
}

func TestNasTypeMaximumNumberOfSupportedPacketFiltersGetSetIei(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	for _, table := range nasTypePDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeMaximumNumberOfSupportedPacketFilters struct {
	in  uint16
	out uint16
}

var nasTypeMaximumNumberOfSupportedPacketFiltersTable = []nasTypeMaximumNumberOfSupportedPacketFilters{
	{0x0100, 0x0100},
}

func TestNasTypeMaximumNumberOfSupportedPacketFiltersGetSetMaximumNumberOfSupportedPacketFilters(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	for _, table := range nasTypeMaximumNumberOfSupportedPacketFiltersTable {
		a.SetMaximumNumberOfSupportedPacketFilters(table.in)
		if !reflect.DeepEqual(table.out, a.GetMaximumNumberOfSupportedPacketFilters()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMaximumNumberOfSupportedPacketFilters())
		}
	}
}
