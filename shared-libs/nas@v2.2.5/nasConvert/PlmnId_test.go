// Copyright (c) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import "testing"

func TestPlmnIDToString(t *testing.T) {
	tests := []struct {
		name     string
		nasBuf   []byte
		expected string
	}{
		{
			name:     "buffer too short",
			nasBuf:   []byte{0x01, 0x02},
			expected: "",
		},
		{
			name:     "valid PLMN ID buffer",
			nasBuf:   []byte{0x12, 0x93, 0x11},
			expected: "213119",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PlmnIDToString(tc.nasBuf)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
