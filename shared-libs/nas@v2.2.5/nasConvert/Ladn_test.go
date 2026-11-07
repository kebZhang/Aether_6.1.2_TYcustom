// Copyright (c) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import "testing"

const testDnnAbc = "abc"

func TestLadnToModels(t *testing.T) {
	tests := []struct {
		name     string
		buf      []uint8
		expected []string
	}{
		{
			name:     "single DNN",
			buf:      []uint8{0x03, 'a', 'b', 'c'},
			expected: []string{testDnnAbc},
		},
		{
			name:     "two DNNs",
			buf:      []uint8{0x03, 'a', 'b', 'c', 0x03, 'x', 'y', 'z'},
			expected: []string{testDnnAbc, "xyz"},
		},
		{
			// declared DNN length exceeds the remaining buffer
			name:     "overlength DNN",
			buf:      []uint8{0x05, 'a', 'b'},
			expected: nil,
		},
		{
			// a zero-length DNN entry must not spin forever
			name:     "zero length DNN",
			buf:      []uint8{0x00, 'a', 'b'},
			expected: nil,
		},
		{
			// TS 24.501 9.11.3.29: only the first 8 LADN DNN values are considered
			name:     "more than 8 DNN values",
			buf:      nineSingleByteDnns(),
			expected: []string{"0", "1", "2", "3", "4", "5", "6", "7"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LadnToModels(tc.buf)
			if len(got) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
				return
			}
			for i := range got {
				if got[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, got)
					break
				}
			}
		})
	}
}

// nineSingleByteDnns builds 9 concatenated [len][DNN] entries with single-character DNNs "0".."8"
func nineSingleByteDnns() []uint8 {
	var buf []uint8
	for i := 0; i < 9; i++ {
		buf = append(buf, 0x01, byte('0'+i))
	}
	return buf
}
