package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"strings"
)

const (
	portfolioAllocationTableName  = "portfolio_allocation_fact"
	getAllPortfolioAllocationsSQL = `
		WITH time_frame_tags 
			AS (SELECT DISTINCT time_frame_tag, create_timestamp FROM [table] ORDER BY create_timestamp DESC LIMIT {:timeFrameLimit})
		SELECT pa.*, ass.ticker as "asset.ticker", COALESCE(ass.name, '') as "asset.name" 
		FROM [table] pa
		JOIN asset ass ON ass.id = pa.asset_id
		` + infra.WhereClausePlaceholder + `
		ORDER BY pa.time_frame_tag DESC, pa.class ASC, pa.cash_reserve DESC, ass.ticker ASC
	`
	getPortfoliosSQL = `
		SELECT p.id, p.name, p.allocation_structure
		FROM portfolio p
	`
)

const (
	queryAllocationsError = "Error querying portfolio allocations"
	queryPortfoliosError  = "Error querying portfolios"
	queryPortfolioError   = "Error querying single portfolio"
)

type PortfolioRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolios() ([]domain.Portfolio, error) {

	var result []domain.Portfolio
	err := repository.dbAdapter.BuildQuery(getPortfoliosSQL).Build().FindInto(&result)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryPortfoliosError, repository)
}

func (repository *PortfolioRDBMSRepository) GetPortfolio(id int) (domain.Portfolio, error) {

	var query = getPortfoliosSQL + `
		WHERE p.id = {:id}
	`

	var result domain.Portfolio
	err := repository.dbAdapter.BuildQuery(query).AddParam("id", id).Build().GetInto(&result)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryPortfolioError, repository)
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolioAllocations(id int, timeFrameLimit int) (
	[]domain.PortfolioAllocation,
	error,
) {
	var query = strings.ReplaceAll(
		getAllPortfolioAllocationsSQL,
		"[table]",
		portfolioAllocationTableName,
	)

	var result []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(query).
		AddParam("timeFrameLimit", timeFrameLimit).
		AddWhereClause("AND pa.time_frame_tag IN (SELECT time_frame_tag FROM time_frame_tags)").
		AddWhereClauseAndParam("AND pa.portfolio_id = {:portfolioId}", "portfolioId", id).
		Build().FindInto(&result)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

func BuildPortfolioRepository(dbAdapter *infra.RDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
