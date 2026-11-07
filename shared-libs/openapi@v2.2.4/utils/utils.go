// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"net/http"

	openapi "github.com/omec-project/openapi/v2"
	"github.com/omec-project/openapi/v2/models"
)

// Problem detail cause constants
const (
	CauseAllocBdtPolicyIdFailed             = "ALLOC_BDT_POLICY_ID_FAILED"
	CauseAuthenticationRejected             = "AUTHENTICATION_REJECTED"
	CauseAmfSubscriptionNotFound            = "AMFSUBSCRIPTION_NOT_FOUND"
	CauseContextNotFound                    = "CONTEXT_NOT_FOUND"
	CauseCreateSubscriptionError            = "CREATE_SUBSCRIPTION_ERROR"
	CauseDataNotFound                       = "DATA_NOT_FOUND"
	CauseDeregistrationNotificationError    = "DEREGISTRATION_NOTIFICATION_ERROR"
	CauseEapPacketParseError                = "EAP_PACKET_PARSE_ERROR"
	CauseFetchError                         = "FETCH_ERROR"
	CauseHandoverFailure                    = "HANDOVER_FAILURE"
	CauseHigherPriorityRequestOngoing       = "HIGHER_PRIORITY_REQUEST_ONGOING"
	CauseIntegrityCheckFail                 = "INTEGRITY_CHECK_FAIL"
	CauseInvalidBodyFormat                  = "INVALID_BODY_FORMAT"
	CauseInvalidGuami                       = "INVALID_GUAMI"
	CauseInvalidMsgFormat                   = "INVALID_MSG_FORMAT"
	CauseInvalidRequest                     = "INVALID_REQUEST"
	CauseMandatoryIeIncorrect               = "MANDATORY_IE_INCORRECT"
	CauseMandatoryIeMissing                 = "MANDATORY_IE_MISSING"
	CauseMessageNotReceived                 = "MESSAGE_NOT_RECEIVED"
	CauseModifyNotAllowed                   = "MODIFY_NOT_ALLOWED"
	CauseNfDeleteError                      = "NF_DELETE_ERROR"
	CauseNotificationError                  = "NOTIFICATION_ERROR"
	CauseNotImplemented                     = "NOT_IMPLEMENTED"
	CauseRequestRejected                    = "REQUEST_REJECTED"
	CauseServerError                        = "SERVER_ERROR"
	CauseSnssaiNotSupported                 = "SNSSAI_NOT_SUPPORTED"
	CauseSubscriptionDeleteError            = "SUBSCRIPTION_DELETE_ERROR"
	CauseSubscriptionEmpty                  = "SUBSCRIPTION_EMPTY"
	CauseSubscriptionEventlistEmpty         = "SUBSCRIPTION_EVENTLIST_EMPTY"
	CauseSubscriptionNotFound               = "SUBSCRIPTION_NOT_FOUND"
	CauseSystemFailure                      = "SYSTEM_FAILURE"
	CauseTemporaryRejectHandoverOngoing     = "TEMPORARY_REJECT_HANDOVER_ONGOING"
	CauseTemporaryRejectRegistrationOngoing = "TEMPORARY_REJECT_REGISTRATION_ONGOING"
	CauseUdrNotFound                        = "UDR_NOT_FOUND"
	CauseUdrQueryFailed                     = "UDR_QUERY_FAILED"
	CauseUeInCmIdleState                    = "UE_IN_CM_IDLE_STATE"
	CauseUeNotReachable                     = "UE_NOT_REACHABLE"
	CauseUeNotServedByAmf                   = "UE_NOT_SERVED_BY_AMF"
	CauseUnexpectedResponseType             = "UNEXPECTED_RESPONSE_TYPE"
	CauseUnspecified                        = "UNSPECIFIED"
	CauseUnspecifiedNfFailure               = "UNSPECIFIED_NF_FAILURE"
	CauseUnsupportedResourceUri             = "UNSUPPORTED_RESOURCE_URI"
	CauseUserNotFound                       = "USER_NOT_FOUND"
)

func ProblemDetails(title string, status int, detail string) *models.ProblemDetails {
	problemDetails := models.NewProblemDetails()
	problemDetails.SetTitle(title)
	problemDetails.SetStatus(int32(status))
	if detail != "" {
		problemDetails.SetDetail(detail)
	}
	return problemDetails
}

func ProblemDetailsWithInvalidParams(title string, status int, detail string, invalidParams []models.InvalidParam) *models.ProblemDetails {
	problemDetails := ProblemDetails(title, status, detail)
	problemDetails.SetInvalidParams(invalidParams)
	return problemDetails
}

func ProblemDetailsSystemFailure(detail string) *models.ProblemDetails {
	problemDetails := ProblemDetails("System failure", http.StatusInternalServerError, detail)
	problemDetails.SetCause(CauseSystemFailure)
	return problemDetails
}

func ProblemDetailsMalformedRequestSyntax(detail string) *models.ProblemDetails {
	return ProblemDetails("Malformed request syntax", http.StatusBadRequest, detail)
}

func ProblemDetailsUnspecified() *models.ProblemDetails {
	problemDetails := ProblemDetails("Unspecified", http.StatusForbidden, "")
	problemDetails.SetCause(CauseUnspecified)
	return problemDetails
}

func ProblemDetailsDataNotFound() *models.ProblemDetails {
	problemDetails := ProblemDetails("Data not found", http.StatusNotFound, "")
	problemDetails.SetCause(CauseDataNotFound)
	return problemDetails
}

func ProblemDetailsUserNotFound() *models.ProblemDetails {
	problemDetails := ProblemDetails("User not found", http.StatusNotFound, "")
	problemDetails.SetCause(CauseUserNotFound)
	return problemDetails
}

func ProblemDetailsFromOpenAPIError(res *http.Response, err error) *models.ProblemDetails {
	if err == nil {
		return nil
	}

	problemDetails := ProblemDetailsSystemFailure(err.Error())
	if res != nil {
		problemDetails.SetStatus(int32(res.StatusCode))
	}

	if details, ok := openapi.ErrorModel[models.ProblemDetails](err); ok {
		if details.Title != nil {
			problemDetails.SetTitle(details.GetTitle())
		}
		if details.Detail != nil {
			problemDetails.SetDetail(details.GetDetail())
		}
		if details.Cause != nil {
			problemDetails.SetCause(details.GetCause())
		}
		if details.Status != nil {
			problemDetails.SetStatus(details.GetStatus())
		}
	}

	return problemDetails
}

func ProblemDetailsContextNotFound(detail string) *models.ProblemDetails {
	return ProblemDetailsWithCause("Context not found", http.StatusNotFound, detail, CauseContextNotFound)
}

func ProblemDetailsNotImplemented(detail string) *models.ProblemDetails {
	return ProblemDetailsWithCause("Not implemented", http.StatusNotImplemented, detail, CauseNotImplemented)
}

func ProblemDetailsMandatoryIeMissing(detail string) *models.ProblemDetails {
	return ProblemDetailsWithCause("Mandatory IE missing", http.StatusBadRequest, detail, CauseMandatoryIeMissing)
}

func ProblemDetailsMandatoryIeIncorrect(detail string) *models.ProblemDetails {
	return ProblemDetailsWithCause("Mandatory IE incorrect", http.StatusBadRequest, detail, CauseMandatoryIeIncorrect)
}

func ProblemDetailsWithCause(title string, status int, detail string, cause string) *models.ProblemDetails {
	problemDetails := ProblemDetails(title, status, detail)
	problemDetails.SetCause(cause)
	return problemDetails
}
