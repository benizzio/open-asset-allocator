package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"strings"
)

const (
	portfolioSliceAtTimeTableName = "asset_value_fact"
	getAllPortfolioSlicesSQL      = `WITH time_frame_tags 
			AS (SELECT DISTINCT time_frame_tag, create_timestamp FROM [table] ORDER BY create_timestamp DESC LIMIT {:timeFrameLimit})
		SELECT pst.*, ass.ticker as "asset.ticker", COALESCE(ass.name, '') as "asset.name" 
		FROM [table] pst
		JOIN asset ass ON ass.id = pst.asset_id
		WHERE pst.time_frame_tag IN (SELECT time_frame_tag FROM time_frame_tags)
		ORDER BY pst.time_frame_tag DESC, pst.class ASC, pst.cash_reserve DESC, ass.ticker ASC`
)

const (
	querySlicesError = "Error querying portfolio slices"
)

type PortfolioRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolioSlices(timeFrameLimit int) (
	[]domain.PortfolioSliceAtTime,
	error,
) {
	var query = strings.ReplaceAll(
		getAllPortfolioSlicesSQL,
		"[table]",
		portfolioSliceAtTimeTableName,
	)

	var result []domain.PortfolioSliceAtTime
	err := repository.dbAdapter.BuildQuery(query).AddParam("timeFrameLimit", timeFrameLimit).FindInto(&result)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, querySlicesError, repository)
}

func BuildPortfolioRepository(dbAdapter *infra.RDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
