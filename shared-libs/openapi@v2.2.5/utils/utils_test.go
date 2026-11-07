// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"errors"
	"net/http"
	"testing"

	openapi "github.com/omec-project/openapi/v2"
	"github.com/omec-project/openapi/v2/models"
)

func TestProblemDetailsFromOpenAPIErrorHandlesTransportError(t *testing.T) {
	problemDetails := ProblemDetailsFromOpenAPIError(nil, errors.New("EOF"))

	if problemDetails.GetStatus() != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, problemDetails.GetStatus())
	}
	if problemDetails.GetCause() != CauseSystemFailure {
		t.Fatalf("expected cause SYSTEM_FAILURE, got %q", problemDetails.GetCause())
	}
	if problemDetails.GetDetail() != "EOF" {
		t.Fatalf("expected detail EOF, got %q", problemDetails.GetDetail())
	}
}

func TestProblemDetailsFromOpenAPIErrorUsesResponseStatusForHTTPError(t *testing.T) {
	problemDetails := ProblemDetailsFromOpenAPIError(
		&http.Response{StatusCode: http.StatusBadGateway, Status: "502 Bad Gateway"},
		errors.New("EOF"),
	)

	if problemDetails.GetStatus() != http.StatusBadGateway {
		t.Fatalf("expected status %d, got %d", http.StatusBadGateway, problemDetails.GetStatus())
	}
	if problemDetails.GetCause() != CauseSystemFailure {
		t.Fatalf("expected cause SYSTEM_FAILURE, got %q", problemDetails.GetCause())
	}
	if problemDetails.GetDetail() != "EOF" {
		t.Fatalf("expected detail EOF, got %q", problemDetails.GetDetail())
	}
}

func TestProblemDetailsFromOpenAPIErrorPreservesStructuredProblemDetails(t *testing.T) {
	title := "Data not found"
	detail := "subscription was not found"
	cause := CauseDataNotFound
	status := int32(http.StatusNotFound)
	rawModel := models.NewProblemDetails()
	rawModel.SetTitle(title)
	rawModel.SetDetail(detail)
	rawModel.SetCause(cause)
	rawModel.SetStatus(status)

	problemDetails := ProblemDetailsFromOpenAPIError(
		&http.Response{StatusCode: http.StatusNotFound, Status: "404 Not Found"},
		&openapi.GenericOpenAPIError{
			RawError: openapi.FormatErrorMessage("404 Not Found", rawModel),
			RawModel: rawModel,
		},
	)

	if problemDetails.GetTitle() != title {
		t.Fatalf("expected title %q, got %q", title, problemDetails.GetTitle())
	}
	if problemDetails.GetDetail() != detail {
		t.Fatalf("expected detail %q, got %q", detail, problemDetails.GetDetail())
	}
	if problemDetails.GetCause() != cause {
		t.Fatalf("expected cause %q, got %q", cause, problemDetails.GetCause())
	}
	if problemDetails.GetStatus() != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, problemDetails.GetStatus())
	}
}

func TestProblemDetailsWithDetail(t *testing.T) {
	pd := ProblemDetails("Test Title", http.StatusBadRequest, "test detail")

	if pd.GetTitle() != "Test Title" {
		t.Fatalf("expected title %q, got %q", "Test Title", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
	if pd.GetDetail() != "test detail" {
		t.Fatalf("expected detail %q, got %q", "test detail", pd.GetDetail())
	}
}

func TestProblemDetailsWithEmptyDetail(t *testing.T) {
	pd := ProblemDetails("Test Title", http.StatusBadRequest, "")

	if pd.GetTitle() != "Test Title" {
		t.Fatalf("expected title %q, got %q", "Test Title", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
	// Detail should be unset (nil) when empty string is passed
	if pd.Detail != nil {
		t.Fatalf("expected Detail to be nil, got %q", pd.GetDetail())
	}
}

func TestProblemDetailsWithInvalidParams(t *testing.T) {
	reason1 := "missing"
	reason2 := "invalid format"
	invalidParams := []models.InvalidParam{
		{Param: "field1", Reason: &reason1},
		{Param: "field2", Reason: &reason2},
	}
	pd := ProblemDetailsWithInvalidParams("Validation Error", http.StatusBadRequest, "request validation failed", invalidParams)

	if pd.GetTitle() != "Validation Error" {
		t.Fatalf("expected title %q, got %q", "Validation Error", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
	if pd.GetDetail() != "request validation failed" {
		t.Fatalf("expected detail %q, got %q", "request validation failed", pd.GetDetail())
	}
	if len(pd.GetInvalidParams()) != 2 {
		t.Fatalf("expected 2 invalid params, got %d", len(pd.GetInvalidParams()))
	}
}

func TestProblemDetailsSystemFailureAlwaysSetsDetail(t *testing.T) {
	// In practice, always called with non-empty detail (err.Error())
	pd := ProblemDetailsSystemFailure("error occurred")

	if pd.GetDetail() != "error occurred" {
		t.Fatalf("expected detail %q, got %q", "error occurred", pd.GetDetail())
	}
	if pd.GetCause() != CauseSystemFailure {
		t.Fatalf("expected cause SYSTEM_FAILURE, got %q", pd.GetCause())
	}
	if pd.GetStatus() != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, pd.GetStatus())
	}
}

func TestProblemDetailsMalformedRequestSyntaxAlwaysSetsDetail(t *testing.T) {
	// In practice, always called with non-empty detail (error messages)
	pd := ProblemDetailsMalformedRequestSyntax("invalid JSON")

	if pd.GetDetail() != "invalid JSON" {
		t.Fatalf("expected detail %q, got %q", "invalid JSON", pd.GetDetail())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
}

func TestProblemDetailsUnspecifiedOmitsDetail(t *testing.T) {
	pd := ProblemDetailsUnspecified()

	if pd.GetTitle() != "Unspecified" {
		t.Fatalf("expected title %q, got %q", "Unspecified", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusForbidden {
		t.Fatalf("expected status %d, got %d", http.StatusForbidden, pd.GetStatus())
	}
	if pd.GetCause() != CauseUnspecified {
		t.Fatalf("expected cause UNSPECIFIED, got %q", pd.GetCause())
	}
	// Detail should be unset (nil)
	if pd.Detail != nil {
		t.Fatalf("expected Detail to be nil, got %q", pd.GetDetail())
	}
}

func TestProblemDetailsDataNotFoundOmitsDetail(t *testing.T) {
	pd := ProblemDetailsDataNotFound()

	if pd.GetTitle() != "Data not found" {
		t.Fatalf("expected title %q, got %q", "Data not found", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, pd.GetStatus())
	}
	if pd.GetCause() != CauseDataNotFound {
		t.Fatalf("expected cause DATA_NOT_FOUND, got %q", pd.GetCause())
	}
	// Detail should be unset (nil)
	if pd.Detail != nil {
		t.Fatalf("expected Detail to be nil, got %q", pd.GetDetail())
	}
}

func TestProblemDetailsUserNotFoundOmitsDetail(t *testing.T) {
	pd := ProblemDetailsUserNotFound()

	if pd.GetTitle() != "User not found" {
		t.Fatalf("expected title %q, got %q", "User not found", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, pd.GetStatus())
	}
	if pd.GetCause() != CauseUserNotFound {
		t.Fatalf("expected cause USER_NOT_FOUND, got %q", pd.GetCause())
	}
	// Detail should be unset (nil)
	if pd.Detail != nil {
		t.Fatalf("expected Detail to be nil, got %q", pd.GetDetail())
	}
}

func TestProblemDetailsContextNotFound(t *testing.T) {
	pd := ProblemDetailsContextNotFound("Guti[12345] Not Found")

	if pd.GetTitle() != "Context not found" {
		t.Fatalf("expected title %q, got %q", "Context not found", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, pd.GetStatus())
	}
	if pd.GetCause() != CauseContextNotFound {
		t.Fatalf("expected cause CONTEXT_NOT_FOUND, got %q", pd.GetCause())
	}
	if pd.GetDetail() != "Guti[12345] Not Found" {
		t.Fatalf("expected detail %q, got %q", "Guti[12345] Not Found", pd.GetDetail())
	}
}

func TestProblemDetailsNotImplemented(t *testing.T) {
	pd := ProblemDetailsNotImplemented("feature not available")

	if pd.GetTitle() != "Not implemented" {
		t.Fatalf("expected title %q, got %q", "Not implemented", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusNotImplemented {
		t.Fatalf("expected status %d, got %d", http.StatusNotImplemented, pd.GetStatus())
	}
	if pd.GetCause() != CauseNotImplemented {
		t.Fatalf("expected cause NOT_IMPLEMENTED, got %q", pd.GetCause())
	}
	if pd.GetDetail() != "feature not available" {
		t.Fatalf("expected detail %q, got %q", "feature not available", pd.GetDetail())
	}
}

func TestProblemDetailsMandatoryIeMissing(t *testing.T) {
	pd := ProblemDetailsMandatoryIeMissing("required field missing")

	if pd.GetTitle() != "Mandatory IE missing" {
		t.Fatalf("expected title %q, got %q", "Mandatory IE missing", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
	if pd.GetCause() != CauseMandatoryIeMissing {
		t.Fatalf("expected cause MANDATORY_IE_MISSING, got %q", pd.GetCause())
	}
	if pd.GetDetail() != "required field missing" {
		t.Fatalf("expected detail %q, got %q", "required field missing", pd.GetDetail())
	}
}

func TestProblemDetailsMandatoryIeIncorrect(t *testing.T) {
	pd := ProblemDetailsMandatoryIeIncorrect("field format incorrect")

	if pd.GetTitle() != "Mandatory IE incorrect" {
		t.Fatalf("expected title %q, got %q", "Mandatory IE incorrect", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, pd.GetStatus())
	}
	if pd.GetCause() != CauseMandatoryIeIncorrect {
		t.Fatalf("expected cause MANDATORY_IE_INCORRECT, got %q", pd.GetCause())
	}
	if pd.GetDetail() != "field format incorrect" {
		t.Fatalf("expected detail %q, got %q", "field format incorrect", pd.GetDetail())
	}
}

func TestProblemDetailsWithCause(t *testing.T) {
	pd := ProblemDetailsWithCause("Custom Error", http.StatusServiceUnavailable, "service unavailable", "CUSTOM_CAUSE")

	if pd.GetTitle() != "Custom Error" {
		t.Fatalf("expected title %q, got %q", "Custom Error", pd.GetTitle())
	}
	if pd.GetStatus() != http.StatusServiceUnavailable {
		t.Fatalf("expected status %d, got %d", http.StatusServiceUnavailable, pd.GetStatus())
	}
	if pd.GetCause() != "CUSTOM_CAUSE" {
		t.Fatalf("expected cause CUSTOM_CAUSE, got %q", pd.GetCause())
	}
	if pd.GetDetail() != "service unavailable" {
		t.Fatalf("expected detail %q, got %q", "service unavailable", pd.GetDetail())
	}
}
