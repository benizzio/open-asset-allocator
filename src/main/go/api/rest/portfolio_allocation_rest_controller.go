package rest

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PortfolioAllocationRESTController struct {
	portfolioDomService *service.PortfolioDomService
}

func (controller *PortfolioAllocationRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     specificPortfolioPath + "/history",
			Handlers: gin.HandlersChain{controller.getPortfolioAllocationHistory},
		},
	}
}

func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var timeFrameTagParamValue = context.Query(timeFrameTagParam)

	portfolioHistory, err := controller.getPortfolioAllocationHistoryUpstack(timeFrameTagParamValue, portfolioId)
	if err != nil {
		var errorDetail string
		if !langext.IsZeroValue(timeFrameTagParamValue) {
			errorDetail = fmt.Sprintf(" for time frame tag %s", timeFrameTagParamValue)
		}
		infra.HandleAPIError(
			context,
			fmt.Sprintf("Error getting portfolio history %s", errorDetail),
			err,
		)
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapToPortfolioHistoryDTSs(portfolioHistory)

	context.JSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
}

// TODO: refactor to use domain.PortfolioObservationTimestamp instead of TimeFrameTag
func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistoryUpstack(
	timeFrameTagParamValue string,
	portfolioId int,
) ([]*domain.PortfolioAllocation, error) {

	var portfolioHistory []*domain.PortfolioAllocation
	var err error

	if !langext.IsZeroValue(timeFrameTagParamValue) {
		var timeFrameTag = domain.TimeFrameTag(timeFrameTagParamValue)
		portfolioHistory, err = controller.portfolioDomService.FindPortfolioAllocations(portfolioId, timeFrameTag)
	} else {
		portfolioHistory, err = controller.portfolioDomService.GetPortfolioAllocationHistory(portfolioId)
	}

	return portfolioHistory, err
}

func BuildPortfolioAllocationRESTController(portfolioDomService *service.PortfolioDomService) *PortfolioAllocationRESTController {
	return &PortfolioAllocationRESTController{
		portfolioDomService,
	}
}
