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

func TestNasTypeNewOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeOldPDUSessionIDULNASTransportOldPDUSessionIDTypeTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportOldPDUSessionIDType, nasMessage.ULNASTransportOldPDUSessionIDType},
}

func TestNasTypeOldPDUSessionIDGetSetIei(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDULNASTransportOldPDUSessionIDTypeTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

type nasTypeOldPDUSessionIDPduSessionIdentity2Value struct {
	in  uint8
	out uint8
}

var nasTypeOldPDUSessionIDPduSessionIdentity2ValueTable = []nasTypeOldPDUSessionIDPduSessionIdentity2Value{
	{0xff, 0xff},
}

func TestNasTypeOldPDUSessionIDGetSetOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDPduSessionIdentity2ValueTable {
		a.SetOldPDUSessionID(table.in)
		if !reflect.DeepEqual(table.out, a.GetOldPDUSessionID()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetOldPDUSessionID())
		}
	}
}

type nasTypeOldPDUSessionID struct {
	inIei                       uint8
	inPduSessionIdentity2Value  uint8
	outIei                      uint8
	outPduSessionIdentity2Value uint8
}

var nasTypeOldPDUSessionIDTable = []nasTypeOldPDUSessionID{
	{
		nasMessage.ULNASTransportOldPDUSessionIDType, 0xff,
		nasMessage.ULNASTransportOldPDUSessionIDType, 0xff,
	},
}

func TestNasTypeOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDTable {
		a.SetIei(table.inIei)
		a.SetOldPDUSessionID(table.inPduSessionIdentity2Value)
		if !reflect.DeepEqual(table.outIei, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.outIei, a.GetIei())
		}
		if !reflect.DeepEqual(table.outPduSessionIdentity2Value, a.GetOldPDUSessionID()) {
			t.Errorf("Not equal: expected %v, got %v", table.outPduSessionIdentity2Value, a.GetOldPDUSessionID())
		}
	}
}
