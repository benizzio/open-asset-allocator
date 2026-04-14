package gin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gingonic "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

	require.NoError(t, err)
	assert.False(t, valid)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errorResponse model.ErrorResponse
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &errorResponse))
	assert.Equal(t, "Validation failed", errorResponse.ErrorMessage)
	require.Len(t, errorResponse.Details, 2)
	assert.Equal(t, "malformed or invalid query parameters", errorResponse.Details[0])
	assert.Equal(t, "Field 'Limit' failed validation: is required", errorResponse.Details[1])
}
