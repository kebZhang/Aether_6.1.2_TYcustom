// SPDX-FileCopyrightText: 2026 Forsway Scandinavia AB
//
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"strings"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
)

// Every cause TS 24.501 V18.13.0 defines - all 45 of table 9.11.4.2.1, cross-checked against the
// annex B list, which agrees on every value - has to render, and the rendering has to carry the
// number: a log line naming the cause without its value cannot be matched against a capture.
func TestCause5GSMToStringRendersEveryDefinedCause(t *testing.T) {
	for cause, want := range map[uint8]string{
		nasMessage.Cause5GSMOperatorDeterminedBarring:                                   "Operator determined barring (8)",
		nasMessage.Cause5GSMInsufficientResources:                                       "Insufficient resources (26)",
		nasMessage.Cause5GSMMissingOrUnknownDNN:                                         "Missing or unknown DNN (27)",
		nasMessage.Cause5GSMUnknownPDUSessionType:                                       "Unknown PDU session type (28)",
		nasMessage.Cause5GSMUserAuthenticationOrAuthorizationFailed:                     "User authentication or authorization failed (29)",
		nasMessage.Cause5GSMRequestRejectedUnspecified:                                  "Request rejected, unspecified (31)",
		nasMessage.Cause5GSMServiceOptionNotSupported:                                   "Service option not supported (32)",
		nasMessage.Cause5GSMRequestedServiceOptionNotSubscribed:                         "Requested service option not subscribed (33)",
		nasMessage.Cause5GSMPTIAlreadyInUse:                                             "PTI already in use (35)",
		nasMessage.Cause5GSMRegularDeactivation:                                         "Regular deactivation (36)",
		nasMessage.Cause5GSM5GSQoSNotAccepted:                                           "5GS QoS not accepted (37)",
		nasMessage.Cause5GSMNetworkFailure:                                              "Network failure (38)",
		nasMessage.Cause5GSMReactivationRequested:                                       "Reactivation requested (39)",
		nasMessage.Cause5GSMSemanticErrorInTheTFTOperation:                              "Semantic error in the TFT operation (41)",
		nasMessage.Cause5GSMSyntacticalErrorInTheTFTOperation:                           "Syntactical error in the TFT operation (42)",
		nasMessage.Cause5GSMInvalidPDUSessionIdentity:                                   "Invalid PDU session identity (43)",
		nasMessage.Cause5GSMSemanticErrorsInPacketFilter:                                "Semantic errors in packet filter(s) (44)",
		nasMessage.Cause5GSMSyntacticalErrorInPacketFilter:                              "Syntactical error in packet filter(s) (45)",
		nasMessage.Cause5GSMOutOfLADNServiceArea:                                        "Out of LADN service area (46)",
		nasMessage.Cause5GSMPTIMismatch:                                                 "PTI mismatch (47)",
		nasMessage.Cause5GSMPDUSessionTypeIPv4OnlyAllowed:                               "PDU session type IPv4 only allowed (50)",
		nasMessage.Cause5GSMPDUSessionTypeIPv6OnlyAllowed:                               "PDU session type IPv6 only allowed (51)",
		nasMessage.Cause5GSMPDUSessionDoesNotExist:                                      "PDU session does not exist (54)",
		nasMessage.Cause5GSMPDUSessionTypeIPv4v6OnlyAllowed:                             "PDU session type IPv4v6 only allowed (57)",
		nasMessage.Cause5GSMPDUSessionTypeUnstructuredOnlyAllowed:                       "PDU session type Unstructured only allowed (58)",
		nasMessage.Cause5GSMUnsupported5QIValue:                                         "Unsupported 5QI value (59)",
		nasMessage.Cause5GSMPDUSessionTypeEthernetOnlyAllowed:                           "PDU session type Ethernet only allowed (61)",
		nasMessage.Cause5GSMInsufficientResourcesForSpecificSliceAndDNN:                 "Insufficient resources for specific slice and DNN (67)",
		nasMessage.Cause5GSMNotSupportedSSCMode:                                         "Not supported SSC mode (68)",
		nasMessage.Cause5GSMInsufficientResourcesForSpecificSlice:                       "Insufficient resources for specific slice (69)",
		nasMessage.Cause5GSMMissingOrUnknownDNNInASlice:                                 "Missing or unknown DNN in a slice (70)",
		nasMessage.Cause5GSMInvalidPTIValue:                                             "Invalid PTI value (81)",
		nasMessage.Cause5GSMMaximumDataRatePerUEForUserPlaneIntegrityProtectionIsTooLow: "Maximum data rate per UE for user-plane integrity protection is too low (82)",
		nasMessage.Cause5GSMSemanticErrorInTheQoSOperation:                              "Semantic error in the QoS operation (83)",
		nasMessage.Cause5GSMSyntacticalErrorInTheQoSOperation:                           "Syntactical error in the QoS operation (84)",
		nasMessage.Cause5GSMInvalidMappedEPSBearerIdentity:                              "Invalid mapped EPS bearer identity (85)",
		nasMessage.Cause5GSMUASServicesNotAllowed:                                       "UAS services not allowed (86)",
		nasMessage.Cause5GSMSemanticallyIncorrectMessage:                                "Semantically incorrect message (95)",
		nasMessage.Cause5GSMInvalidMandatoryInformation:                                 "Invalid mandatory information (96)",
		nasMessage.Cause5GSMMessageTypeNonExistentOrNotImplemented:                      "Message type non-existent or not implemented (97)",
		nasMessage.Cause5GSMMessageTypeNotCompatibleWithTheProtocolState:                "Message type not compatible with the protocol state (98)",
		nasMessage.Cause5GSMInformationElementNonExistentOrNotImplemented:               "Information element non-existent or not implemented (99)",
		nasMessage.Cause5GSMConditionalIEError:                                          "Conditional IE error (100)",
		nasMessage.Cause5GSMMessageNotCompatibleWithTheProtocolState:                    "Message not compatible with the protocol state (101)",
		nasMessage.Cause5GSMProtocolErrorUnspecified:                                    "Protocol error, unspecified (111)",
	} {
		if got := nasMessage.Cause5GSMToString(cause); got != want {
			t.Errorf("Cause5GSMToString(%d) = %q, want %q", cause, got, want)
		}
	}
}

// The empty string is the contract for an undefined cause, matching Cause5GMMToString. It is what
// lets a caller use this to decide whether a value belongs in the 5GSM cause IE at all - a UE
// decodes an undefined value as 31 "request rejected, unspecified", so nothing downstream reports
// the mistake.
func TestCause5GSMToStringReportsUndefinedCausesAsEmpty(t *testing.T) {
	for _, cause := range []uint8{
		0x00,
		0x1e, // 30, between "user authentication or authorization failed" and "request rejected, unspecified"
		0x5b, // 91, a 5GMM cause (DNN not supported or not subscribed in the slice) with no 5GSM counterpart
		0xff,
	} {
		if got := nasMessage.Cause5GSMToString(cause); got != "" {
			t.Errorf("Cause5GSMToString(%d) = %q, want an empty string for an undefined cause", cause, got)
		}
	}
}

// 34 is the awkward one, and it is asserted rather than left to be rediscovered.
//
// Cause5GSMServiceOptionTemporarilyOutOfOrder has been in this package for as long as the rest,
// but TS 24.501 V18.13.0 does not define 34 for 5GSM - not in table 9.11.4.2.1 and not in annex B,
// which agree with each other on all 45 values it does define. The constant stays, because
// removing it would break callers, and it does not render: reporting it would say a value that
// must not be put in the 5GSM cause IE is fit to send.
func TestCause5GSMToStringDoesNotRenderTheUndefined34(t *testing.T) {
	if got := nasMessage.Cause5GSMToString(nasMessage.Cause5GSMServiceOptionTemporarilyOutOfOrder); got != "" {
		t.Errorf("Cause5GSMToString(34) = %q; TS 24.501 does not define 34 for 5GSM", got)
	}
}

// 5GSM and 5GMM are separate registers that share numeric values, which is the mistake this
// function exists to catch: a 5GMM cause written into the 5GSM cause IE reads as a valid, wrong
// cause or as none at all. Both directions are asserted, so neither function can quietly grow to
// cover the other's register.
func TestCause5GSMAndCause5GMMStayInTheirOwnRegisters(t *testing.T) {
	const dnnNotSubscribedInSlice = nasMessage.Cause5GMMDNNNotSupportedOrNotSubscribedInTheSlice // 91

	if got := nasMessage.Cause5GSMToString(dnnNotSubscribedInSlice); got != "" {
		t.Errorf("Cause5GSMToString(%d) = %q; 91 is a 5GMM cause and must not resolve as 5GSM",
			dnnNotSubscribedInSlice, got)
	}
	if got := nasMessage.Cause5GMMToString(nasMessage.Cause5GSMMissingOrUnknownDNNInASlice); got != "" {
		t.Errorf("Cause5GMMToString(70) = %q; 70 is a 5GSM cause and must not resolve as 5GMM", got)
	}

	// Where the registers do overlap they have to agree only on the number, not on being the same
	// cause: 69 is "insufficient resources for specific slice" in both, and each function must
	// answer from its own table.
	if got := nasMessage.Cause5GSMToString(nasMessage.Cause5GSMInsufficientResourcesForSpecificSlice); !strings.HasSuffix(got, "(69)") {
		t.Errorf("Cause5GSMToString(69) = %q, want it to render as 69", got)
	}
	if got := nasMessage.Cause5GMMToString(nasMessage.Cause5GMMInsufficientResourcesForSpecificSlice); !strings.HasSuffix(got, "(69)") {
		t.Errorf("Cause5GMMToString(69) = %q, want it to render as 69", got)
	}
}
