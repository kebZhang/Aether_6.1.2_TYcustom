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

type nasTypeMessageAuthenticationCodeMACData struct {
	in  [4]uint8
	out [4]uint8
}

var nasTypeMessageAuthenticationCodeMACTable = []nasTypeMessageAuthenticationCodeMACData{
	{[4]uint8{0xff, 0xff, 0xff, 0xff}, [4]uint8{0xff, 0xff, 0xff, 0xff}},
}

func TestNasTypeNewMessageAuthenticationCode(t *testing.T) {
	a := nasType.NewMessageAuthenticationCode()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeMessageAuthenticationCode(t *testing.T) {
	a := nasType.NewMessageAuthenticationCode()
	for _, table := range nasTypeMessageAuthenticationCodeMACTable {
		a.SetMAC(table.in)
		if !reflect.DeepEqual(table.out, a.GetMAC()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMAC())
		}
	}
}
