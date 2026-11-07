// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nas

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

var hexString = "7e00560102000021e440b883d63a9f9c56b3703217152eba2010068f241c77748000b2180e54a9760068"

func TestNasGmmMessage(t *testing.T) {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		t.Fatal(err)
	}
	m := NewMessage()
	err = m.GmmMessageDecode(&data)
	if err != nil {
		t.Fatalf("Unexpected non-nil value: %v", err)
	}

	buff := new(bytes.Buffer)
	err = m.GmmMessageEncode(buff)
	if err != nil {
		t.Fatalf("Unexpected non-nil value: %v", err)
	}
}

func TestNasGsmMessage(t *testing.T) {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		t.Fatal(err)
	}
	m := NewMessage()
	err = m.GsmMessageDecode(&data)
	if err == nil {
		t.Fatal("Expected value not to be nil")
	}

	buff := new(bytes.Buffer)
	err = m.GsmMessageEncode(buff)
	if err == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestPlainNas(t *testing.T) {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		t.Fatal(err)
	}
	m := NewMessage()
	err = m.PlainNasDecode(&data)
	if err != nil {
		t.Fatalf("Unexpected non-nil value: %v", err)
	}
	buff, err1 := m.PlainNasEncode()
	if err1 != nil {
		t.Fatalf("Unexpected non-nil value: %v", err1)
	}
	if !reflect.DeepEqual(data, buff) {
		t.Errorf("Expect: 0x%0x\nOutput: 0x%0x", data, buff)
	}
}
