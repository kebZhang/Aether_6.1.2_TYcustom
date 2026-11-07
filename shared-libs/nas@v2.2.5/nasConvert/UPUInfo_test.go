// Copyright (c) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import "testing"

func TestUpuAckToModels(t *testing.T) {
	tests := []struct {
		name      string
		buf       []uint8
		expectErr bool
	}{
		{
			name:      "empty UPU ack buffer",
			buf:       []uint8{},
			expectErr: true,
		},
		{
			name:      "wrong length",
			buf:       []uint8{0x01, 0x02},
			expectErr: true,
		},
		{
			name:      "valid UPU ack buffer",
			buf:       append([]uint8{0x01}, make([]uint8, 16)...),
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := UpuAckToModels(tc.buf)
			if (err != nil) != tc.expectErr {
				t.Errorf("expected error: %v, got: %v", tc.expectErr, err)
			}
		})
	}
}
