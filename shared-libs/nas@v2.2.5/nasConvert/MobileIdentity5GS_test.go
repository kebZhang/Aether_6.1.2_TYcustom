// SPDX-FileCopyrightText: 2025 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import (
	"testing"
)

func TestGutiToString(t *testing.T) {
	tests := []struct {
		name    string
		buf     []byte
		wantGut string
	}{
		{
			name:    "buffer too short",
			buf:     []byte{0x01, 0x02, 0x03},
			wantGut: "",
		},
		{
			name:    "buffer too long",
			buf:     []byte{0xf1, 0x12, 0x93, 0x11, 0x22, 0x33, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			wantGut: "",
		},
		{
			name:    "valid GUTI buffer",
			buf:     []byte{0xf1, 0x12, 0x93, 0x11, 0x22, 0x33, 0x01, 0x02, 0x03, 0x04, 0x05},
			wantGut: "21311922330102030405",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, guti := GutiToString(tc.buf)
			if guti != tc.wantGut {
				t.Errorf("expected %q, got %q", tc.wantGut, guti)
			}
		})
	}
}

func TestNaiToString(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expectedNai string
		expectErr   bool
	}{
		{
			name:        "nil buffer",
			input:       nil,
			expectedNai: "",
			expectErr:   true,
		},
		{
			name:        "empty buffer",
			input:       []byte{},
			expectedNai: "",
			expectErr:   true,
		},
		{
			name:        "type-only buffer",
			input:       []byte{0x01},
			expectedNai: "",
			expectErr:   true,
		},
		{
			name:        "minimum valid buffer",
			input:       []byte{0x01, 0x23},
			expectedNai: "nai-1-23",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NaiToString(tt.input)
			if (err != nil) != tt.expectErr {
				t.Fatalf("NaiToString() error = %v, expectErr %v", err, tt.expectErr)
			}
			if result != tt.expectedNai {
				t.Errorf("NaiToString() = %q, expected %q", result, tt.expectedNai)
			}
		})
	}
}

func TestSuciToString(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		expectedSuci string
		expectedPlmn string
		expectErr    bool
	}{
		{
			name:         "nil buffer",
			input:        nil,
			expectedSuci: "",
			expectedPlmn: "",
			expectErr:    true,
		},
		{
			name:         "empty buffer",
			input:        []byte{},
			expectedSuci: "",
			expectedPlmn: "",
			expectErr:    true,
		},
		{
			name:         "undersized imsi buffer",
			input:        []byte{0x00, 0x21, 0x63, 0x54, 0x76, 0x98, 0x00},
			expectedSuci: "",
			expectedPlmn: "",
			expectErr:    true,
		},
		{
			name:         "imsi buffer missing scheme output",
			input:        []byte{0x00, 0x21, 0x63, 0x54, 0x76, 0xf8, 0x00, 0x01},
			expectedSuci: "",
			expectedPlmn: "",
			expectErr:    true,
		},
		{
			name:         "valid imsi buffer",
			input:        []byte{0x00, 0x21, 0x63, 0x54, 0x76, 0xf8, 0x00, 0x01, 0x32},
			expectedSuci: "suci-0-123-456-678-0-1-23",
			expectedPlmn: "123456",
			expectErr:    false,
		},
		{
			name:         "invalid nai buffer",
			input:        []byte{0x10},
			expectedSuci: "",
			expectedPlmn: "",
			expectErr:    true,
		},
		{
			name:         "valid nai buffer",
			input:        []byte{0x10, 0x23},
			expectedSuci: "nai-1-23",
			expectedPlmn: "",
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suci, plmn, err := SuciToString(tt.input)
			if (err != nil) != tt.expectErr {
				t.Fatalf("SuciToString() error = %v, expectErr %v", err, tt.expectErr)
			}
			if suci != tt.expectedSuci {
				t.Errorf("SuciToString() suci = %q, expected %q", suci, tt.expectedSuci)
			}
			if plmn != tt.expectedPlmn {
				t.Errorf("SuciToString() plmnId = %q, expected %q", plmn, tt.expectedPlmn)
			}
		})
	}
}

func TestPeiToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte // Hexadecimal representation of the input
		expected string // Expected output
	}{
		{
			name:     "Valid IMEI",
			input:    []byte{0x3b, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83, 0x85}, // IMEI: 356938035643859
			expected: "imei-356938035643858",
		},
		{
			name:     "Valid IMEISV",
			input:    []byte{0x35, 0x55, 0x12, 0x18, 0x01, 0x31, 0x66, 0x44, 0xF6}, // IMEISV: 3549774394518016
			expected: "imeisv-3552181101366446",
		},
		{
			name:     "Invalid IMEI length",
			input:    []byte{0x3b, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83}, // IMEI: 13 digits
			expected: "",
		},
		{
			name:     "Invalid IMEISV length",
			input:    []byte{0x35, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83, 0x95, 0x01, 0x01}, // IMEISV: 18 digits
			expected: "",
		},
		{
			name:     "Invalid character in IMEI",
			input:    []byte{0x9b, 0x09, 0x0a, 0x80, 0x26, 0x74, 0x81, 0x35}, // invalid character 'a'
			expected: "",
		},
		{
			name:     "Invalid character in IMEISV",
			input:    []byte{0x35, 0x45, 0x79, 0x47, 0xd3, 0x54, 0x81, 0x10, 0xF6}, // invalid character 'd'
			expected: "",
		},
		{
			name:     "Invalid TAC/SNR (Luhn check fails)",
			input:    []byte{0x3b, 0x65, 0x39, 0x58, 0x53, 0x46, 0x83, 0x15}, // IMEI: Fails Luhn check
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PeiToString(tt.input)
			if result != tt.expected {
				t.Errorf("PeiToString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
