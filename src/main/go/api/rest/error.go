package rest

import (
	"errors"
	"net/http"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

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

func handleDomainError(context *gin.Context, cause error) bool {

	var domValidationError *infra.DomainValidationError
	if errors.As(cause, &domValidationError) {

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
