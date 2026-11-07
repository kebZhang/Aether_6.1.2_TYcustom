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

func TestNasTypeNewAuthorizedQosRules(t *testing.T) {
	a := nasType.NewAuthorizedQosRules(nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

var nasTypeAuthenticationRequestAuthorizedQosRulesIeiTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType, nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType},
}

func TestNasTypeAuthorizedQosRulesGetSetIei(t *testing.T) {
	a := nasType.NewAuthorizedQosRules(nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType)
	for _, table := range nasTypeAuthenticationRequestAuthorizedQosRulesIeiTable {
		a.SetIei(table.in)
		if !reflect.DeepEqual(table.out, a.GetIei()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetIei())
		}
	}
}

var nasTypeAuthorizedQosRulesLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeAuthorizedQosRulesGetSetLen(t *testing.T) {
	a := nasType.NewAuthorizedQosRules(nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType)
	for _, table := range nasTypeAuthorizedQosRulesLenTable {
		a.SetLen(table.in)
		if !reflect.DeepEqual(table.out, a.GetLen()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetLen())
		}
	}
}

type nasTypetAuthorizedQosRulesQosRule struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeAuthorizedQosRulesTable = []nasTypetAuthorizedQosRulesQosRule{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeAuthorizedQosRulesGetSetAuthorizedQosRules(t *testing.T) {
	a := nasType.NewAuthorizedQosRules(0)
	for _, table := range nasTypeAuthorizedQosRulesTable {
		a.SetLen(table.inLen)
		a.SetQosRule(table.in)
		if !reflect.DeepEqual(table.out, a.GetQosRule()) {
			t.Errorf("in(%v): out %v, actual %x", table.in, table.out, a.GetQosRule())
		}
	}
}

type testAuthorizedQosRulesDataTemplate struct {
	in  nasType.AuthorizedQosRules
	out nasType.AuthorizedQosRules
}

var AuthorizedQosRulesTestData = []nasType.AuthorizedQosRules{
	{nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType, 2, []byte{0x00, 0x00}}, // AuthenticationResult
}

var AuthorizedQosRulesExpectedData = []nasType.AuthorizedQosRules{
	{nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType, 2, []byte{0x00, 0x00}}, // AuthenticationResult
}

var AuthorizedQosRulesTestTable = []testAuthorizedQosRulesDataTemplate{
	{AuthorizedQosRulesTestData[0], AuthorizedQosRulesExpectedData[0]},
}

func TestNasTypeAuthorizedQosRules(t *testing.T) {
	for i, table := range AuthorizedQosRulesTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthorizedQosRules(0) // AuthenticationResult

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetQosRule(table.in.Buffer)

		if !reflect.DeepEqual(table.out.Iei, a.Iei) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		}
		if !reflect.DeepEqual(table.out.Len, a.Len) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		}
		if !reflect.DeepEqual(table.out.Buffer, a.Buffer) {
			t.Errorf("in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
		}
	}
}
