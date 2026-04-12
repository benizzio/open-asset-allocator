package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	"github.com/benizzio/open-asset-allocator/infra/validation"
	"github.com/benizzio/open-asset-allocator/langext"
)

type AssetRESTController struct {
	assetDomService *service.AssetDomService
}

func (controller *AssetRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/asset",
			Handlers: gin.HandlersChain{controller.getKnownAssets},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/asset/:" + assetIdOrTickerParam,
			Handlers: gin.HandlersChain{controller.getAssetById},
		},
		{
			Method:   http.MethodPut,
			Path:     "/api/asset",
			Handlers: gin.HandlersChain{controller.putAsset},
		},
	}
}

func (controller *AssetRESTController) getKnownAssets(context *gin.Context) {

	assets, err := controller.assetDomService.GetKnownAssets()
	if gininfra.HandleAPIError(context, "Error getting known assets", err) {
		return
	}

	var assetRESTModels = model.MapToAssetDTSs(assets)

	context.JSON(http.StatusOK, assetRESTModels)
}

func (controller *AssetRESTController) getAssetById(context *gin.Context) {

	var assetIdOrTickerParamValue = context.Param(assetIdOrTickerParam)

	asset, err := controller.assetDomService.FindAssetByUniqueIdentifier(assetIdOrTickerParamValue)
	if gininfra.HandleAPIError(context, "Error getting asset by Id or Ticker", err) {
		return
	}

	if asset == nil {
		gininfra.SendDataNotFoundResponse(context, "Asset", assetIdOrTickerParamValue)
		return
	}

	var assetDTS = model.MapToAssetDTS(asset)
	context.JSON(http.StatusOK, assetDTS)
}

// putAsset handles PUT requests to update an existing asset's ticker and name fields.
// Validates that the asset ID is present and non-zero before delegating to the domain service.
//
// Authored by: GitHub Copilot
func (controller *AssetRESTController) putAsset(context *gin.Context) {

	var assetDTS model.AssetDTS
	valid, err := gininfra.BindAndValidateJSONWithInvalidResponse(context, &assetDTS)
	if err != nil {
		gininfra.HandleAPIError(context, bindAssetErrorMessage, err)
		return
	}
	if !valid {
		return
	}

	if assetDTS.Id == nil || langext.IsZeroValue(*assetDTS.Id) {

		var validationErrors = validation.BuildCustomValidationErrorsBuilder().
			CustomValidationError(
				assetDTS,
				"Id",
				"required",
				"Asset ID is required for update",
				nil,
			).
			Build()

		gininfra.RespondWithCustomValidationErrors(context, validationErrors, assetDTS)

		return
	}

	var asset = model.MapToAsset(&assetDTS)
	updatedAsset, err := controller.assetDomService.UpdateAsset(asset)
	if gininfra.HandleAPIError(context, "Error updating asset", err) {
		return
	}

	var responseBody = model.MapToAssetDTS(updatedAsset)
	context.JSON(http.StatusOK, responseBody)
}

func BuildAssetRESTController(assetDomService *service.AssetDomService) *AssetRESTController {
	return &AssetRESTController{
		assetDomService: assetDomService,
	}
}
