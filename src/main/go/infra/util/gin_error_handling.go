package util

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// sendValidationErrorResponse sends a standardized HTTP response for validation validationErrors.
//
// Authored by: GitHub Copilot
func sendValidationErrorResponse(context *gin.Context, errorMessages []string) {
	context.JSON(
		http.StatusBadRequest, model.ErrorResponse{
			ErrorMessage: "Validation failed",
			Details:      errorMessages,
		},
	)
}

func SendDataNotFoundResponse(context *gin.Context, dataType string, id string) {
	context.JSON(
		http.StatusNotFound,
		model.ErrorResponse{
			ErrorMessage: "Data not found",
			Details:      []string{dataType + " with ID " + id + " not found"},
		},
	)
}
