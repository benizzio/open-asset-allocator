package repository

import (
	"database/sql"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"strconv"
)

const (
	assetsSQL = `
		SELECT * FROM asset
	` + infra.WhereClausePlaceholder
)

type AssetRDBMSRepository struct {
	dbAdapter infra.RepositoryRDBMSAdapter
}

func (repository *AssetRDBMSRepository) GetKnownAssets() ([]*domain.Asset, error) {
	var queryBuilder = repository.dbAdapter.BuildQuery(assetsSQL)
	var result []domain.Asset
	err := queryBuilder.Build().FindInto(&result)
	return langext.ToPointerSlice(result), infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error getting known assets",
		repository,
	)
}

func (repository *AssetRDBMSRepository) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(assetsSQL)

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

func (repository *AssetRDBMSRepository) InsertAssets(
	transContext *infra.TransactionalContext,
	assets []*domain.Asset,
) ([]*domain.Asset, error) {

	if len(assets) == 0 {
		return assets, nil
	}

	var columns = []string{"ticker", "name"}

	var values = make([][]interface{}, len(assets))
	var tickers = make([]string, len(assets))
	for i, asset := range assets {
		values[i] = []interface{}{
			asset.Ticker,
			asset.Name,
		}
		tickers[i] = asset.Ticker
	}

	err := repository.dbAdapter.InsertBulkInTransaction(transContext, "asset", columns, values)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting assets",
			repository,
		)
	}

	var queryBuilder = infra.BuildQueryWithinTransaction[domain.Asset](transContext, assetsSQL).AddWhereClauseAndParams(
		"AND ticker IN ?",
		tickers,
	).Build()

	persistedAssets, err := queryBuilder.Find(
		func(rows *sql.Rows) (domain.Asset, error) {
			var asset domain.Asset
			scanErr := rows.Scan(
				&asset.Id,
				&asset.Ticker,
				&asset.Name,
			)
			return asset, scanErr
		},
	)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error retrieving persisted assets",
			repository,
		)
	}

	return langext.ToPointerSlice(persistedAssets), nil
}

func BuildAssetRDBMSRepository(dbAdapter infra.RepositoryRDBMSAdapter) *AssetRDBMSRepository {
	return &AssetRDBMSRepository{
		dbAdapter: dbAdapter,
	}
}
