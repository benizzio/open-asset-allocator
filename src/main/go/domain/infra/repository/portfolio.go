package repository

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
)

const (
	portfolioSliceAtTimeTableName = "asset_value_fact"
)

type PortfolioRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolioSlices(limit int) ([]domain.PortfolioSliceAtTime, error) {

	var query = fmt.Sprintf(
		"SELECT * FROM %s ORDER BY create_timestamp DESC LIMIT {:limit}",
		portfolioSliceAtTimeTableName,
	)

	var result []domain.PortfolioSliceAtTime
	var err = repository.dbAdapter.BuildQuery(query).AddParam("limit", limit).FindInto(&result)
	return result, err
}

func BuildPortfolioRepository(dbAdapter *infra.RDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
