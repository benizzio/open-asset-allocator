package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	}
}

func (controller *AssetRESTController) getKnownAssets(context *gin.Context) {

	assets, err := controller.assetDomService.GetKnownAssets()
	if infra.HandleAPIError(context, "Error getting known assets", err) {
		return
	}

	var assetRESTModels = model.MapToAssetDTSs(assets)

	context.JSON(http.StatusOK, assetRESTModels)
}

func (controller *AssetRESTController) getAssetById(context *gin.Context) {

	var assetIdParamValue = context.Param(assetIdParam)
	assetId, err := strconv.Atoi(assetIdParamValue)
	if infra.HandleAPIError(context, "Error parsing asset ID", err) {
		return
	}

	asset, err := controller.assetDomService.FindAssetById(assetId)
	if infra.HandleAPIError(context, "Error getting asset by ID", err) {
		return
	}

	var assetDTS = model.MapToAssetDTS(asset)
	context.JSON(http.StatusOK, assetDTS)
}

func BuildAssetRESTController(assetDomService *service.AssetDomService) *AssetRESTController {
	return &AssetRESTController{
		assetDomService: assetDomService,
	}
}
