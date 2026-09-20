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

// BuildRoutes returns the HTTP routes handled by the asset REST controller.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
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
			Method:   http.MethodPost,
			Path:     "/api/asset",
			Handlers: gin.HandlersChain{controller.postAsset},
		},
		{
			Method:   http.MethodPut,
			Path:     "/api/asset",
			Handlers: gin.HandlersChain{controller.putAsset},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/external-asset",
			Handlers: gin.HandlersChain{controller.getExternalAssets},
		},
	}
}

func (controller *AssetRESTController) getKnownAssets(context *gin.Context) {

	var assetSearchQueryDTS model.AssetSearchQueryDTS
	valid, err := gininfra.BindAndValidateQueryWithInvalidResponse(context, &assetSearchQueryDTS)
	if err != nil {
		gininfra.HandleAPIError(context, "Error binding asset search query", err)
		return
	}
	if !valid {
		return
	}

	assets, err := controller.assetDomService.GetKnownAssets(assetSearchQueryDTS.TextSearch)
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

// postAsset handles POST requests that create an asset with optional external data. The database
// generates the asset ID; non-zero client-supplied IDs are rejected.
//
// Authored by: OpenCode
func (controller *AssetRESTController) postAsset(context *gin.Context) {

	var assetDTS model.AssetDTS
	valid, err := gininfra.BindAndValidateJSONWithInvalidResponse(context, &assetDTS)
	if err != nil {
		gininfra.HandleAPIError(context, bindAssetErrorMessage, err)
		return
	}
	if !valid {
		return
	}

	if assetDTS.Id != nil && !langext.IsZeroValue(*assetDTS.Id) {
		var validationErrors = validation.BuildCustomValidationErrorsBuilder().
			CustomValidationError(
				assetDTS,
				"Id",
				"custom",
				"Asset ID must be omitted or zero for creation",
				*assetDTS.Id,
			).
			Build()

		gininfra.RespondWithCustomValidationErrors(context, validationErrors, assetDTS)
		return
	}

	var asset = model.MapToAsset(&assetDTS)
	createdAsset, err := controller.assetDomService.CreateAsset(asset)
	if gininfra.HandleAPIError(context, "Error creating asset", err) {
		return
	}

	var responseBody = model.MapToAssetDTS(createdAsset)
	context.JSON(http.StatusCreated, responseBody)
}

// putAsset handles PUT requests to replace an existing asset's ticker, name, and external data
// fields. An omitted or null externalData value clears the persisted external data.
// Validates that the asset ID is present and non-zero before delegating to the domain service.
//
// Co-authored by: GitHub Copilot and OpenCode
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

// getExternalAssets handles GET requests that search external asset providers by query string.
//
// Co-authored by: OpenCode and benizzio
func (controller *AssetRESTController) getExternalAssets(context *gin.Context) {

	var externalAssetSearchQueryDTS model.ExternalAssetSearchQueryDTS
	valid, err := gininfra.BindAndValidateQueryWithInvalidResponse(context, &externalAssetSearchQueryDTS)
	if err != nil {
		gininfra.HandleAPIError(context, "Error binding external asset search query", err)
		return
	}
	if !valid {
		return
	}

	externalAssets, err := controller.assetDomService.SearchExternalAssets(
		context.Request.Context(),
		externalAssetSearchQueryDTS.Query,
	)
	if gininfra.HandleAPIError(context, "Error searching external assets", err) {
		return
	}

	var externalAssetDTSs = model.MapToExternalAssetDTSs(externalAssets)
	context.JSON(http.StatusOK, externalAssetDTSs)
}
