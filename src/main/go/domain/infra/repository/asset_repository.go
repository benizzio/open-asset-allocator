package repository

import (
	"database/sql"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"strconv"
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
	return langext.ToPointerSlice(result), infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error getting known assets",
		repository,
	)
}

func (repository *AssetRDBMSRepository) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(knownAssetsSQL)

	var whereClause string
	if _, err := strconv.Atoi(uniqueIdentifier); err == nil {
		whereClause = "AND (id = {:uniqueIdentifier} OR ticker = {:uniqueIdentifier})"
	} else {
		whereClause = "AND ticker = {:uniqueIdentifier}"
	}

	queryBuilder.AddWhereClauseAndParam(
		whereClause,
		"uniqueIdentifier",
		uniqueIdentifier,
	)

	var result domain.Asset
	err := queryBuilder.Build().GetInto(&result)

	if err != nil && err.Error() == sql.ErrNoRows.Error() {
		return nil, nil
	}

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
