package gin

import (
	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/infra/validation"
)

// BindAndValidateQueryWithInvalidResponse binds the request query parameters to the provided
// struct and validates them.
//
// Authored by: GitHub Copilot
func BindAndValidateQueryWithInvalidResponse(context *gin.Context, bindingTarget interface{}) (bool, error) {

	var allErrorMessages []string

	var bindingErr = context.ShouldBindQuery(bindingTarget)
	if bindingErr != nil {

		var validationErrors = validation.MapValidationErrorsToMessages(bindingErr, bindingTarget)
		if validationErrors == nil {
			return false, bindingErr
		}

		allErrorMessages = append(allErrorMessages, validationErrors...)
	}

	var deepValidationErrors = validation.DeepValidate(bindingTarget)
	allErrorMessages = append(allErrorMessages, deepValidationErrors...)

	if len(allErrorMessages) > 0 {
		sendValidationErrorResponse(context, allErrorMessages)
		return false, nil
	}

	return true, nil
}
