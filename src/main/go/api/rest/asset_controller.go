package rest

import (
	"net/http"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	"github.com/gin-gonic/gin"
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

	var assetIdOrTickerParamValue = context.Param(assetIdOrTickerParam)

	asset, err := controller.assetDomService.FindAssetByUniqueIdentifier(assetIdOrTickerParamValue)
	if infra.HandleAPIError(context, "Error getting asset by Id or Ticker", err) {
		return
	}

	if asset == nil {
		gininfra.SendDataNotFoundResponse(context, "Asset", assetIdOrTickerParamValue)
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
