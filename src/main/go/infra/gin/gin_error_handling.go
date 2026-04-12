// gin_error_handling.go provides unified HTTP error response handling for the Gin REST API layer.
//
// Co-authored by: GitHub Copilot
package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/infra"
)

// HandleAPIError handles an error from the API layer by logging it and sending an appropriate HTTP response.
// It first attempts to match the error to a known domain error type (e.g. DomainValidationError) and sends
// a specific response. If no domain match is found, it falls back to a generic 500 Internal Server Error.
// Returns true if an error was present and handled, false otherwise.
//
// Example:
//
//	assets, err := service.GetAssets()
//	if gininfra.HandleAPIError(context, "Error getting assets", err) {
//		return
//	}
//
// Co-authored by: GitHub Copilot
func HandleAPIError(context *gin.Context, message string, cause error) bool {

	var handle = cause != nil
	if handle {

		glog.Error(message, ": ", cause)

		var handled = handleDomainError(context, cause)
		if handled {
			return true
		}

		// Fallback for unhandled errors
		context.JSON(
			http.StatusInternalServerError, model.ErrorResponse{
				ErrorMessage: "Internal server error",
			},
		)
	}

	return handle
}

// handleDomainError checks if the error matches a known domain error type and sends the corresponding HTTP response.
//
// Co-authored by: GitHub Copilot
func handleDomainError(context *gin.Context, cause error) bool {

	if domValidationError, ok := errors.AsType[*infra.DomainValidationError](cause); ok {

		var validationMessages = make([]string, len(domValidationError.Causes))
		for i, validationError := range domValidationError.Causes {
			validationMessages[i] = validationError.Message
		}
		context.JSON(
			http.StatusBadRequest, model.ErrorResponse{
				ErrorMessage: domValidationError.Message,
				Details:      validationMessages,
			},
		)

		return true
	}

	return false
}

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

// SendDataNotFoundResponse sends a standardized HTTP 404 response for a missing resource.
func SendDataNotFoundResponse(context *gin.Context, dataType string, id string) {
	context.JSON(
		http.StatusNotFound,
		model.ErrorResponse{
			ErrorMessage: "Data not found",
			Details:      []string{dataType + " with identifier " + id + " not found"},
		},
	)
}
