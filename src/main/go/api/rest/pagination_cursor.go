package rest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/domain"
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

// paginationCursorPayload binds a cursor to its query scope and one descending-order boundary.
// Authored by: OpenCode
type paginationCursorPayload struct {
	Version     int    `json:"version"`
	Portfolio   int64  `json:"portfolioId"`
	Collection  string `json:"collection"`
	Timestamp   string `json:"timestamp,omitempty"`
	PlannedDate string `json:"plannedDate,omitempty"`
	Id          int64  `json:"id"`
	PlanType    string `json:"planType,omitempty"`
}

// paginationCursorValidationDTS gives cursor errors the same field-oriented response shape as other API validation.
// Authored by: OpenCode
type paginationCursorValidationDTS struct {
	PageCursor string `json:"pageCursor"`
}

// parsedPaginationCursor contains the validated boundary shared by observation and allocation-plan handlers.
// Authored by: OpenCode
type parsedPaginationCursor struct {
	Timestamp   time.Time
	PlannedDate time.Time
	Id          int64
}

// readPaginationCursor detects parameter presence, decodes a continuation token, and responds with HTTP 400 on error.
// A present empty value represents the first page; an absent parameter leaves the legacy endpoint path untouched.
// For example, `?pageCursor=` returns (nil, true), while a missing parameter returns (nil, false).
// Co-authored by: GPT-6.1 Sol and OpenCode
func readPaginationCursor(
	context *gin.Context,
	portfolioId int64,
	collection string,
	planType string,
) (*parsedPaginationCursor, bool) {
	valuesByKey, queryError := url.ParseQuery(context.Request.URL.RawQuery)
	values, present := valuesByKey[paginationCursorQueryKey]
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

	cursor, err := decodePaginationCursor(values[0], portfolioId, collection, planType)
	if err != nil {
		respondWithPaginationCursorValidationError(context, err.Error())
		return nil, true
	}
	return cursor, true
}

// decodePaginationCursor validates and decodes one opaque cursor for a specific portfolio and collection.
// For example, decodePaginationCursor(token, 42, observationCursorCollection, "") rejects tokens from portfolio 5.
// Authored by: OpenCode
func decodePaginationCursor(
	token string,
	portfolioId int64,
	collection string,
	planType string,
) (*parsedPaginationCursor, error) {
	payload, err := decodePaginationCursorPayload(token)
	if err != nil {
		return nil, err
	}
	if err = validatePaginationCursorScope(payload, portfolioId, collection); err != nil {
		return nil, err
	}
	return decodePaginationCursorBoundary(payload, collection, planType)
}

// decodePaginationCursorPayload decodes one bounded, strict JSON cursor payload from URL-safe Base64.
// Authored by: OpenCode
func decodePaginationCursorPayload(token string) (paginationCursorPayload, error) {
	var payload paginationCursorPayload
	if len(token) > paginationCursorMaxLength {
		return payload, errors.New("must not exceed 1024 characters")
	}
	payloadBytes, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil || base64.RawURLEncoding.EncodeToString(payloadBytes) != token {
		return payload, errors.New("must be a valid URL-safe cursor")
	}
	decoder := json.NewDecoder(bytes.NewReader(payloadBytes))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&payload); err != nil {
		return payload, errors.New("contains an invalid cursor payload")
	}
	if err = decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return payload, errors.New("must contain exactly one cursor payload")
	}
	return payload, nil
}

// validatePaginationCursorScope verifies the cursor version, portfolio, collection, and record ID.
// Authored by: OpenCode
func validatePaginationCursorScope(
	payload paginationCursorPayload,
	portfolioId int64,
	collection string,
) error {
	if payload.Version != paginationCursorVersion {
		return errors.New("uses an unsupported cursor version")
	}
	if payload.Portfolio <= 0 || payload.Portfolio != portfolioId {
		return errors.New("belongs to a different portfolio")
	}
	if payload.Collection != collection {
		return errors.New("belongs to a different collection")
	}
	if payload.Id <= 0 {
		return errors.New("contains an invalid record ID")
	}
	return nil
}

// decodePaginationCursorBoundary validates the collection-specific values and builds a typed REST boundary.
// Authored by: OpenCode
func decodePaginationCursorBoundary(
	payload paginationCursorPayload,
	collection string,
	planType string,
) (*parsedPaginationCursor, error) {
	cursor := &parsedPaginationCursor{Id: payload.Id}
	switch collection {
	case observationCursorCollection:
		if payload.Timestamp == "" || payload.PlannedDate != "" || payload.PlanType != "" {
			return nil, errors.New("must contain an observation timestamp and ID only")
		}
		parsedTimestamp, err := time.Parse(time.RFC3339Nano, payload.Timestamp)
		if err != nil || parsedTimestamp.IsZero() {
			return nil, errors.New("contains an invalid observation timestamp")
		}
		cursor.Timestamp = parsedTimestamp
	case allocationPlanCollection:
		if payload.PlannedDate == "" || payload.Timestamp != "" {
			return nil, errors.New("must contain a planned date and ID only")
		}
		if payload.PlanType != planType {
			return nil, errors.New("belongs to a different allocation-plan type")
		}
		parsedDate, err := time.Parse(time.DateOnly, payload.PlannedDate)
		if err != nil || parsedDate.Format(time.DateOnly) != payload.PlannedDate {
			return nil, errors.New("contains an invalid planned date")
		}
		cursor.PlannedDate = parsedDate
	default:
		return nil, errors.New("uses an unsupported collection")
	}
	return cursor, nil
}

// encodeObservationPaginationCursor serializes an observation boundary into an opaque URL-safe cursor.
// For example, its return value can be sent unchanged as the next pageCursor query value.
// Authored by: OpenCode
func encodeObservationPaginationCursor(
	portfolioId int64,
	cursor *domain.ObservationPageCursor,
) (*string, error) {
	if cursor == nil {
		return nil, nil
	}
	return encodePaginationCursor(paginationCursorPayload{
		Version:    paginationCursorVersion,
		Portfolio:  portfolioId,
		Collection: observationCursorCollection,
		Timestamp:  cursor.Timestamp.UTC().Format(time.RFC3339Nano),
		Id:         cursor.Id,
	})
}

// encodeAllocationPlanPaginationCursor serializes a plan boundary for the plan type used by the endpoint.
// For example, its return value can be sent unchanged as the next pageCursor query value.
// Authored by: OpenCode
func encodeAllocationPlanPaginationCursor(
	portfolioId int64,
	planType string,
	cursor *domain.AllocationPlanPageCursor,
) (*string, error) {
	if cursor == nil {
		return nil, nil
	}
	return encodePaginationCursor(paginationCursorPayload{
		Version:     paginationCursorVersion,
		Portfolio:   portfolioId,
		Collection:  allocationPlanCollection,
		PlannedDate: cursor.PlannedDate.Format(time.DateOnly),
		Id:          cursor.Id,
		PlanType:    planType,
	})
}

// encodePaginationCursor encodes a typed cursor payload as unpadded URL-safe Base64 JSON.
// Authored by: OpenCode
func encodePaginationCursor(payload paginationCursorPayload) (*string, error) {
	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encoding pagination cursor: %w", err)
	}
	encodedCursor := base64.RawURLEncoding.EncodeToString(encodedPayload)
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
