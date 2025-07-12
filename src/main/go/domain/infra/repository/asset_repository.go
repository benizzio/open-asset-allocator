package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
)

const (
	knownAssetsSQL = `
		SELECT * FROM asset
	` + infra.WhereClausePlaceholder
)

type AssetRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AssetRDBMSRepository) GetKnownAssets() ([]*domain.Asset, error) {
	var queryBuilder = repository.dbAdapter.BuildQuery(knownAssetsSQL)
	var result []domain.Asset
	err := queryBuilder.Build().FindInto(&result)
	return nil, infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error getting known assets",
		repository,
	)
}

func (repository *AssetRDBMSRepository) FindAssetById(id int) (*domain.Asset, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(knownAssetsSQL)
	queryBuilder.AddWhereClauseAndParam("AND id = {:id}", "id", id)

	var result domain.Asset
	err := queryBuilder.Build().GetInto(&result)

	return &result, infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error getting asset by id",
		repository,
	)
}

func BuildAssetRDBMSRepository(dbAdapter *infra.RDBMSAdapter) *AssetRDBMSRepository {
	return &AssetRDBMSRepository{
		dbAdapter: dbAdapter,
	}
}
