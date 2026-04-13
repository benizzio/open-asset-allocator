package gin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gingonic "github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
)

// malformedQueryBindingTarget uses an integer field so query binding can fail before validation.
//
// Authored by: OpenCode
type malformedQueryBindingTarget struct {
	Limit int `form:"limit" validate:"required"`
}

// TestBindAndValidateQueryWithInvalidResponse_MalformedQuery verifies that malformed query values
// stay on the validation response path instead of being returned as raw binding errors.
//
// Authored by: OpenCode
func TestBindAndValidateQueryWithInvalidResponse_MalformedQuery(t *testing.T) {
	gingonic.SetMode(gingonic.TestMode)

	var recorder = httptest.NewRecorder()
	var context, _ = gingonic.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/test?limit=abc", nil)

	var bindingTarget malformedQueryBindingTarget
	var valid, err = BindAndValidateQueryWithInvalidResponse(context, &bindingTarget)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if valid {
		t.Fatal("Expected query binding to be invalid")
	}
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var errorResponse model.ErrorResponse
	if decodeErr := json.Unmarshal(recorder.Body.Bytes(), &errorResponse); decodeErr != nil {
		t.Fatalf("Expected valid JSON response, got %v", decodeErr)
	}
	if errorResponse.ErrorMessage != "Validation failed" {
		t.Fatalf("Expected validation error message, got %q", errorResponse.ErrorMessage)
	}
	if len(errorResponse.Details) != 2 {
		t.Fatalf("Expected two validation details, got %#v", errorResponse.Details)
	}
	if errorResponse.Details[0] != "malformed or invalid query parameters" {
		t.Fatalf("Expected malformed query fallback detail first, got %#v", errorResponse.Details)
	}
	if errorResponse.Details[1] != "Field 'Limit' failed validation: is required" {
		t.Fatalf("Expected required validation detail second, got %#v", errorResponse.Details)
	}
}
