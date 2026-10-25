package rest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	urlinfra "github.com/benizzio/open-asset-allocator/infra/util/url"
	"github.com/benizzio/open-asset-allocator/infra/validation"
)

const (
	paginationCursorQueryKey    = "pageCursor"
	paginationCursorVersion     = 1
	paginationCursorMaxLength   = 1024
	observationCursorCollection = "observations"
	allocationPlanCollection    = "allocation-plans"
)

// paginationCursorScope binds a continuation boundary to its version, portfolio, collection, and record ID.
// Authored by: GPT-6.1 Sol
type paginationCursorScope struct {
	Version    int    `json:"version"`
	Portfolio  int64  `json:"portfolioId"`
	Collection string `json:"collection"`
	Id         int64  `json:"id"`
}

// observationPaginationCursorPayload composes shared scope with an observation timestamp.
// Embedding preserves the flat version-1 JSON format without accepting allocation-plan fields.
// Authored by: GPT-6.1 Sol
type observationPaginationCursorPayload struct {
	paginationCursorScope
	Timestamp string `json:"timestamp"`
}

// allocationPlanPaginationCursorPayload composes shared scope with a planned date and plan-type filter.
// Embedding preserves the flat version-1 JSON format without accepting observation fields.
// Authored by: GPT-6.1 Sol
type allocationPlanPaginationCursorPayload struct {
	paginationCursorScope
	PlannedDate string `json:"plannedDate"`
	PlanType    string `json:"planType"`
}

// paginationCursorPayload limits shared token mechanics to the two concrete wire payloads.
// Authored by: GPT-6.1 Sol
type paginationCursorPayload interface {
	observationPaginationCursorPayload | allocationPlanPaginationCursorPayload
	cursorScope() paginationCursorScope
}

// cursorScope exposes the embedded scope for shared validation without erasing the payload type.
// Authored by: GPT-6.1 Sol
func (scope paginationCursorScope) cursorScope() paginationCursorScope {
	return scope
}

// paginationCursorValidationDTS gives cursor errors the same field-oriented response shape as other API validation.
// Authored by: OpenCode
type paginationCursorValidationDTS struct {
	PageCursor string `json:"pageCursor"`
}

// readPaginationCursor detects parameter presence and strictly decodes a collection-specific payload.
// A present empty value represents the first page; an absent parameter leaves the legacy endpoint path untouched.
// For example, `?pageCursor=` returns (nil, true), while a missing parameter returns (nil, false).
// Co-authored by: GPT-6.1 Sol and OpenCode
func readPaginationCursor[payloadType paginationCursorPayload](
	context *gin.Context,
	portfolioId int64,
	collection string,
) (*payloadType, bool) {
	var valuesByKey, queryError = url.ParseQuery(context.Request.URL.RawQuery)
	var values, present = valuesByKey[paginationCursorQueryKey]
	if queryError != nil && urlinfra.RawQueryHasParameter(context.Request.URL.RawQuery, paginationCursorQueryKey) {
		respondWithPaginationCursorValidationError(context, "must be URL-encoded")
		return nil, true
	}
	if !present {
		return nil, false
	}
	if len(values) != 1 {
		respondWithPaginationCursorValidationError(context, "must be supplied exactly once")
		return nil, true
	}
	if values[0] == "" {
		return nil, true
	}

	var payload, err = decodePaginationCursorPayload[payloadType](values[0])
	if err == nil {
		err = validatePaginationCursorScope(payload.cursorScope(), portfolioId, collection)
	}
	if err != nil {
		respondWithPaginationCursorValidationError(context, err.Error())
		return nil, true
	}
	return &payload, true
}

// readObservationPaginationCursor validates an observation token and returns its domain boundary directly.
// Authored by: GPT-6.1 Sol
func readObservationPaginationCursor(
	context *gin.Context,
	portfolioId int64,
) (*domain.ObservationPageCursor, bool) {
	var payload, present = readPaginationCursor[observationPaginationCursorPayload](
		context, portfolioId, observationCursorCollection,
	)
	if payload == nil {
		return nil, present
	}
	var timestamp, err = time.Parse(time.RFC3339Nano, payload.Timestamp)
	if err != nil || timestamp.IsZero() {
		respondWithPaginationCursorValidationError(context, "contains an invalid observation timestamp")
		return nil, true
	}
	return &domain.ObservationPageCursor{Timestamp: timestamp, Id: payload.Id}, true
}

// readAllocationPlanPaginationCursor validates plan-type scope and returns a planned-date domain boundary directly.
// Authored by: GPT-6.1 Sol
func readAllocationPlanPaginationCursor(
	context *gin.Context,
	portfolioId int64,
	planType string,
) (*domain.AllocationPlanPageCursor, bool) {
	var payload, present = readPaginationCursor[allocationPlanPaginationCursorPayload](
		context, portfolioId, allocationPlanCollection,
	)
	if payload == nil {
		return nil, present
	}
	if payload.PlanType != planType {
		respondWithPaginationCursorValidationError(context, "belongs to a different allocation-plan type")
		return nil, true
	}
	var plannedDate, err = time.Parse(time.DateOnly, payload.PlannedDate)
	if err != nil || plannedDate.Format(time.DateOnly) != payload.PlannedDate {
		respondWithPaginationCursorValidationError(context, "contains an invalid planned date")
		return nil, true
	}
	return &domain.AllocationPlanPageCursor{PlannedDate: plannedDate, Id: payload.Id}, true
}

// decodePaginationCursorPayload decodes one bounded, strict JSON cursor payload from URL-safe Base64.
// Co-authored by: GPT-6.1 Sol and OpenCode
func decodePaginationCursorPayload[payloadType paginationCursorPayload](token string) (payloadType, error) {
	var payload payloadType
	if len(token) > paginationCursorMaxLength {
		return payload, infra.BuildAppError("must not exceed 1024 characters", payload)
	}
	var payloadBytes, err = base64.RawURLEncoding.DecodeString(token)
	if err != nil || base64.RawURLEncoding.EncodeToString(payloadBytes) != token {
		return payload, infra.BuildAppError("must be a valid URL-safe cursor", payload)
	}
	var decoder = json.NewDecoder(bytes.NewReader(payloadBytes))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&payload); err != nil {
		return payload, infra.PropagateAsAppErrorWithNewMessage(err, "contains an invalid cursor payload", payload)
	}
	if err = decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return payload, infra.BuildAppError("must contain exactly one cursor payload", payload)
	}
	return payload, nil
}

// validatePaginationCursorScope verifies the cursor version, portfolio, collection, and record ID.
// Co-authored by: GPT-6.1 Sol and OpenCode
func validatePaginationCursorScope(
	payload paginationCursorScope,
	portfolioId int64,
	collection string,
) error {
	if payload.Version != paginationCursorVersion {
		return infra.BuildAppError("uses an unsupported cursor version", payload)
	}
	if payload.Portfolio <= 0 || payload.Portfolio != portfolioId {
		return infra.BuildAppError("belongs to a different portfolio", payload)
	}
	if payload.Collection != collection {
		return infra.BuildAppError("belongs to a different collection", payload)
	}
	if payload.Id <= 0 {
		return infra.BuildAppError("contains an invalid record ID", payload)
	}
	return nil
}

// encodeObservationPaginationCursor serializes an observation boundary into an opaque URL-safe cursor.
// For example, its return value can be sent unchanged as the next pageCursor query value.
// Co-authored by: GPT-6.1 Sol and OpenCode
func encodeObservationPaginationCursor(
	portfolioId int64,
	cursor *domain.ObservationPageCursor,
) (*string, error) {
	if cursor == nil {
		return nil, nil
	}
	return encodePaginationCursor(observationPaginationCursorPayload{
		paginationCursorScope: paginationCursorScope{
			Version: paginationCursorVersion, Portfolio: portfolioId,
			Collection: observationCursorCollection, Id: cursor.Id,
		},
		Timestamp: cursor.Timestamp.UTC().Format(time.RFC3339Nano),
	})
}

// encodeAllocationPlanPaginationCursor serializes a plan boundary for the plan type used by the endpoint.
// For example, its return value can be sent unchanged as the next pageCursor query value.
// Co-authored by: GPT-6.1 Sol and OpenCode
func encodeAllocationPlanPaginationCursor(
	portfolioId int64,
	planType string,
	cursor *domain.AllocationPlanPageCursor,
) (*string, error) {
	if cursor == nil {
		return nil, nil
	}
	return encodePaginationCursor(allocationPlanPaginationCursorPayload{
		paginationCursorScope: paginationCursorScope{
			Version: paginationCursorVersion, Portfolio: portfolioId,
			Collection: allocationPlanCollection, Id: cursor.Id,
		},
		PlannedDate: cursor.PlannedDate.Format(time.DateOnly),
		PlanType:    planType,
	})
}

// encodePaginationCursor encodes a typed cursor payload as unpadded URL-safe Base64 JSON.
// Co-authored by: GPT-6.1 Sol and OpenCode
func encodePaginationCursor[payloadType paginationCursorPayload](payload payloadType) (*string, error) {
	var encodedPayload, err = json.Marshal(payload)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "encoding pagination cursor", payload)
	}
	var encodedCursor = base64.RawURLEncoding.EncodeToString(encodedPayload)
	return &encodedCursor, nil
}

// respondWithPaginationCursorValidationError sends a standard validation response without echoing the token.
// Authored by: OpenCode
func respondWithPaginationCursorValidationError(context *gin.Context, message string) {
	var target paginationCursorValidationDTS
	var validationErrors = validation.BuildCustomValidationErrorsBuilder().
		CustomValidationError(&target, "PageCursor", "custom", message, nil).
		Build()
	gininfra.RespondWithCustomValidationErrors(context, validationErrors, &target)
	context.Abort()
}
