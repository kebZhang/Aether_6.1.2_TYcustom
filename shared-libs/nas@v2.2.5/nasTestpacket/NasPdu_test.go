// SPDX-FileCopyrightText: 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0
//

package nasTestpacket

import (
	"bytes"
	"slices"
	"testing"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/nasMessage"
)

func TestGetSecurityModeComplete(t *testing.T) {
	tests := []struct {
		name                string
		nasMessageContainer []uint8
		expectedMinLength   int
		validateIMEISV      bool
		validateContainer   bool
	}{
		{
			name:                "SecurityModeComplete without NAS message container",
			nasMessageContainer: nil,
			expectedMinLength:   10, // Reduced minimum expected length
			validateIMEISV:      true,
			validateContainer:   false,
		},
		{
			name:                "SecurityModeComplete with empty NAS message container",
			nasMessageContainer: []uint8{},
			expectedMinLength:   12, // Reduced minimum expected length
			validateIMEISV:      true,
			validateContainer:   true,
		},
		{
			name:                "SecurityModeComplete with NAS message container",
			nasMessageContainer: []uint8{0x01, 0x02, 0x03, 0x04, 0x05},
			expectedMinLength:   17, // Reduced minimum expected length
			validateIMEISV:      true,
			validateContainer:   true,
		},
		{
			name:                "SecurityModeComplete with large NAS message container",
			nasMessageContainer: make([]uint8, 100), // 100 bytes
			expectedMinLength:   112,                // Reduced minimum expected length
			validateIMEISV:      true,
			validateContainer:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			result := GetSecurityModeComplete(tt.nasMessageContainer)

			// Basic validations
			if result == nil {
				t.Fatal("Result should not be nil")
			}

			if len(result) < tt.expectedMinLength {
				t.Errorf("Result length %d should be at least the expected minimum %d",
					len(result), tt.expectedMinLength)
			}

			// Validate the message is not empty
			if len(result) == 0 {
				t.Error("Result should not be empty")
			}

			// Parse the result to validate structure
			validateSecurityModeCompleteMessage(t, result, tt.nasMessageContainer,
				tt.validateIMEISV, tt.validateContainer)
		})
	}
}

func TestGetSecurityModeComplete_MessageStructure(t *testing.T) {
	t.Run("Validate message header fields", func(t *testing.T) {
		result := GetSecurityModeComplete(nil)

		// The result should contain properly encoded NAS message
		if len(result) == 0 {
			t.Fatal("Result should not be empty")
		}

		// First byte should contain Extended Protocol Discriminator
		// For 5GS MM messages, this should be 0x7E
		if result[0] != 0x7E {
			t.Errorf("Extended Protocol Discriminator should be 0x7E for 5GS MM, got 0x%02X", result[0])
		}

		// Second byte contains spare half octet and security header type
		// For plain NAS, security header type should be 0x00
		securityHeaderType := result[1] & 0x0F
		if securityHeaderType != 0x00 {
			t.Errorf("Security header type should be 0x00 for plain NAS, got 0x%02X", securityHeaderType)
		}

		// Third byte should be the message type for Security Mode Complete
		// This should be 0x5E according to TS 24.501
		if result[2] != 0x5E {
			t.Errorf("Message type should be 0x5E for Security Mode Complete, got 0x%02X", result[2])
		}
	})
}

func TestGetSecurityModeComplete_EdgeCases(t *testing.T) {
	t.Run("Very large NAS message container", func(t *testing.T) {
		// Test with maximum reasonable size
		largeContainer := make([]uint8, 1000)
		for i := range largeContainer {
			largeContainer[i] = uint8(i % 256)
		}

		result := GetSecurityModeComplete(largeContainer)
		if result == nil {
			t.Fatal("Result should not be nil")
		}

		if len(result) <= 1000 {
			t.Errorf("Result length %d should be greater than 1000 to contain the large container", len(result))
		}
	})

	t.Run("NAS message container with specific pattern", func(t *testing.T) {
		// Test with a specific pattern to ensure data integrity
		pattern := []uint8{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
		result := GetSecurityModeComplete(pattern)

		if result == nil {
			t.Fatal("Result should not be nil")
		}

		if !slices.Contains(result, pattern[0]) {
			t.Error("Result should contain parts of the input pattern")
		}
	})
}

func TestGetSecurityModeComplete_Consistency(t *testing.T) {
	t.Run("Multiple calls with same input produce same output", func(t *testing.T) {
		container := []uint8{0x01, 0x02, 0x03}

		result1 := GetSecurityModeComplete(container)
		result2 := GetSecurityModeComplete(container)

		if !bytes.Equal(result1, result2) {
			t.Error("Multiple calls with same input should produce identical output")
		}
	})

	t.Run("Different inputs produce different outputs", func(t *testing.T) {
		container1 := []uint8{0x01, 0x02, 0x03}
		container2 := []uint8{0x04, 0x05, 0x06}

		result1 := GetSecurityModeComplete(container1)
		result2 := GetSecurityModeComplete(container2)

		if bytes.Equal(result1, result2) {
			t.Error("Different inputs should produce different outputs")
		}
	})
}

func TestGetSecurityModeComplete_NilInput(t *testing.T) {
	result := GetSecurityModeComplete(nil)

	if result == nil {
		t.Fatal("Result should not be nil even with nil input")
	}

	if len(result) == 0 {
		t.Error("Result should not be empty even with nil input")
	}
}

func TestGetSecurityModeComplete_EmptyInput(t *testing.T) {
	result := GetSecurityModeComplete([]uint8{})

	if result == nil {
		t.Fatal("Result should not be nil with empty input")
	}

	if len(result) == 0 {
		t.Error("Result should not be empty with empty input")
	}
}

// Helper function to validate the structure of the Security Mode Complete message
func validateSecurityModeCompleteMessage(t *testing.T, result []byte, expectedContainer []uint8,
	validateIMEISV, validateContainer bool,
) {
	t.Helper()

	if len(result) <= 3 {
		t.Fatalf("Message should be longer than header, got length %d", len(result))
	}

	message := nas.NewMessage()
	if err := message.PlainNasDecode(&result); err != nil {
		t.Fatalf("Failed to decode Security Mode Complete: %v", err)
	}

	if message.GmmMessage == nil || message.SecurityModeComplete == nil {
		t.Fatal("Decoded message does not contain Security Mode Complete")
	}

	securityModeComplete := message.SecurityModeComplete

	// Validate Extended Protocol Discriminator (should be 0x7E for 5GS MM)
	if securityModeComplete.GetExtendedProtocolDiscriminator() !=
		nasMessage.Epd5GSMobilityManagementMessage {
		t.Errorf("Invalid Extended Protocol Discriminator: expected 0x%02X, got 0x%02X",
			nasMessage.Epd5GSMobilityManagementMessage,
			securityModeComplete.GetExtendedProtocolDiscriminator())
	}

	// Validate Security Header Type (should be 0x00 for plain NAS)
	securityHeaderType := securityModeComplete.GetSecurityHeaderType()
	if securityHeaderType != 0x00 {
		t.Errorf("Invalid Security Header Type: expected 0x00, got 0x%02X", securityHeaderType)
	}

	// Validate Message Type (should be 0x5E for Security Mode Complete)
	if securityModeComplete.GetMessageType() != nas.MsgTypeSecurityModeComplete {
		t.Errorf("Invalid Message Type: expected 0x%02X, got 0x%02X",
			nas.MsgTypeSecurityModeComplete,
			securityModeComplete.GetMessageType())
	}

	if validateIMEISV {
		if securityModeComplete.IMEISV == nil {
			t.Error("IMEISV should be present in the message")
		} else if securityModeComplete.IMEISV.GetLen() != 9 {
			t.Errorf("IMEISV length should be 9, got %d", securityModeComplete.IMEISV.GetLen())
		}
	}

	if validateContainer && expectedContainer != nil {
		if securityModeComplete.NASMessageContainer == nil {
			t.Fatal("NAS message container should be present in the message")
		}

		if securityModeComplete.NASMessageContainer.GetLen() != uint16(len(expectedContainer)) {
			t.Errorf("NAS message container length should be %d, got %d",
				len(expectedContainer), securityModeComplete.NASMessageContainer.GetLen())
		}

		if !bytes.Equal(securityModeComplete.GetNASMessageContainerContents(), expectedContainer) {
			t.Errorf("NAS message container contents mismatch: expected %x, got %x",
				expectedContainer, securityModeComplete.GetNASMessageContainerContents())
		}
	}
}

// Modernized benchmark tests using b.Loop() (Go 1.22+)
func BenchmarkGetSecurityModeComplete(b *testing.B) {
	container := []uint8{0x01, 0x02, 0x03, 0x04, 0x05}

	for b.Loop() {
		GetSecurityModeComplete(container)
	}
}

func BenchmarkGetSecurityModeCompleteNoContainer(b *testing.B) {
	for b.Loop() {
		GetSecurityModeComplete(nil)
	}
}

func BenchmarkGetSecurityModeCompleteLargeContainer(b *testing.B) {
	largeContainer := make([]uint8, 1000)
	for i := range largeContainer {
		largeContainer[i] = uint8(i % 256)
	}

	for b.Loop() {
		GetSecurityModeComplete(largeContainer)
	}
}
