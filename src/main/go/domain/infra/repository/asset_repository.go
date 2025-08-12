package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

const (
	assetsSQL = `
		SELECT * FROM asset
	` + rdbms.WhereClausePlaceholder
)

func assetRowScanner(rows *sql.Rows) (domain.Asset, error) {
	var asset domain.Asset
	scanErr := rows.Scan(
		&asset.Id,
		&asset.Ticker,
		&asset.Name,
	)
	return asset, scanErr
}

type AssetRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
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

func (repository *AssetRDBMSRepository) InsertAssetsInTransaction(
	transContext context.Context,
	assets []*domain.Asset,
) ([]*domain.Asset, error) {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return nil, infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

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

	err := repository.dbAdapter.InsertBulkInTransaction(transactionalContext, "asset", columns, values)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting assets",
			repository,
		)
	}

	return repository.FindAssetsByTickersInTransaction(transactionalContext, tickers)
}

func (repository *AssetRDBMSRepository) FindAssetsByTickersInTransaction(
	transContext context.Context,
	tickers []string,
) ([]*domain.Asset, error) {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return nil, infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	var queryExecutor = rdbms.BuildQueryInTransaction[domain.Asset](transactionalContext, assetsSQL).
		AddWhereClauseAndParams("AND ticker = ANY($1)", tickers).
		Build()

	persistedAssets, err := queryExecutor.Find(assetRowScanner)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error retrieving persisted assets",
			repository,
		)
	}

	return langext.ToPointerSlice(persistedAssets), nil
}

func BuildAssetRDBMSRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *AssetRDBMSRepository {
	return &AssetRDBMSRepository{
		dbAdapter: dbAdapter,
	}
}
