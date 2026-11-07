// SPDX-FileCopyrightText: 2025 Intel Corporation
// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/omec-project/openapi/v2"
	"github.com/omec-project/openapi/v2/models"
)

func TestRequestedNssaiToModels(t *testing.T) {
	testCases := []struct {
		name         string
		requestNssai nasType.RequestedNSSAI
		expected     []models.MappingOfSnssai
		expectError  bool
	}{
		{
			name: "Test correctness",
			requestNssai: nasType.RequestedNSSAI{
				Iei: nasMessage.RegistrationRequestRequestedNSSAIType,
				Len: 25,
				Buffer: []uint8{
					0x01, 0x01,
					0x02, 0x01, 0x02,
					0x04, 0x01, 0x01, 0x02, 0x03,
					0x05, 0x01, 0x01, 0x02, 0x03, 0x03,
					0x08, 0x01, 0x11, 0x22, 0x33, 0x04, 0x01, 0x02, 0x03,
				},
			},
			expected: []models.MappingOfSnssai{
				{
					ServingSnssai: models.Snssai{
						Sst: 1,
					},
				},
				{
					ServingSnssai: models.Snssai{
						Sst: 1,
					},
					HomeSnssai: models.Snssai{
						Sst: 2,
					},
				},
				{
					ServingSnssai: models.Snssai{
						Sst: 1,
						Sd:  openapi.PtrString("010203"),
					},
				},
				{
					ServingSnssai: models.Snssai{
						Sst: 1,
						Sd:  openapi.PtrString("010203"),
					},
					HomeSnssai: models.Snssai{
						Sst: 3,
					},
				},
				{
					ServingSnssai: models.Snssai{
						Sst: 1,
						Sd:  openapi.PtrString("112233"),
					},
					HomeSnssai: models.Snssai{
						Sst: 4,
						Sd:  openapi.PtrString("010203"),
					},
				},
			},
			expectError: false,
		},
		{
			name: "Test error handling",
			requestNssai: nasType.RequestedNSSAI{
				Iei: nasMessage.RegistrationRequestRequestedNSSAIType,
				Len: 2,
				Buffer: []uint8{
					0x09, 0x01,
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			// declared S-NSSAI contents length (0x08) exceeds the remaining buffer
			name: "Test overlength S-NSSAI contents",
			requestNssai: nasType.RequestedNSSAI{
				Iei: nasMessage.RegistrationRequestRequestedNSSAIType,
				Len: 3,
				Buffer: []uint8{
					0x08, 0x01, 0x02,
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			// advertised Len exceeds the actual buffer length
			name: "Test overlength requested NSSAI",
			requestNssai: nasType.RequestedNSSAI{
				Iei: nasMessage.RegistrationRequestRequestedNSSAIType,
				Len: 10,
				Buffer: []uint8{
					0x01, 0x01,
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			// buffer holds a full S-NSSAI, but the advertised Len only covers part of it
			name: "Test Len shorter than buffer must not expose trailing bytes",
			requestNssai: nasType.RequestedNSSAI{
				Iei: nasMessage.RegistrationRequestRequestedNSSAIType,
				Len: 2,
				Buffer: []uint8{
					0x04, 0x01, 0x01, 0x02, 0x03,
				},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			modelNssai, err := RequestedNssaiToModels(&tc.requestNssai)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(modelNssai, tc.expected) {
				t.Errorf("Expected %+v, got %+v", tc.expected, modelNssai)
			}
		})
	}
}
