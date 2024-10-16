package api

import (
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PortfolioRESTController struct {
}

func (controller *PortfolioRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{Method: http.MethodGet, Path: "/api/portfolio", Handlers: gin.HandlersChain{controller.getPortfolioHistory}},
	}
}

func (controller *PortfolioRESTController) getPortfolioHistory(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testVar)
}

type test struct {
	TestField1 string `json:"testField1"`
	TestField2 string `json:"testField2"`
}

var testVar = []test{
	{TestField1: "test1", TestField2: "test2"},
	{TestField1: "test3", TestField2: "test4"},
}

func BuildPortfolioRESTController() *PortfolioRESTController {
	return &PortfolioRESTController{}
}
