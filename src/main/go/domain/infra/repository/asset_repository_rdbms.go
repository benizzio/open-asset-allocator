package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

const (
	assetsSQL = `
		SELECT id, ticker, name, external_data FROM asset
	` + rdbms.WhereClausePlaceholder
)

// assetRowScanner reads a persisted asset row, including its optional external data payload,
// into the domain model.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func assetRowScanner(rows *sql.Rows) (domain.Asset, error) {

	var asset domain.Asset
	var externalData domain.ExternalAssetData
	var externalDataValue interface{}

	scanErr := rows.Scan(
		&asset.Id,
		&asset.Ticker,
		&asset.Name,
		&externalDataValue,
	)
	if scanErr != nil {
		return asset, scanErr
	}

	if externalDataValue != nil {
		scanErr = externalData.Scan(externalDataValue)
		if scanErr != nil {
			return asset, scanErr
		}
		asset.ExternalData = &externalData
	}

	return asset, scanErr
}

type AssetRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

// GetKnownAssets retrieves all persisted assets, including optional external data.
//
// Example:
//
//	assets, err := assetRepository.GetKnownAssets()
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func (repository *AssetRDBMSRepository) GetKnownAssets() ([]*domain.Asset, error) {
	var queryExecutor = rdbms.BuildQuery[domain.Asset](repository.dbAdapter, assetsSQL).Build()
	result, err := queryExecutor.FindWithRowScanner(assetRowScanner)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error getting known assets",
			repository,
		)
	}

	return langext.ToPointerSlice(result), nil
}

// FindAssetByUniqueIdentifier retrieves a single asset by numeric id or ticker. Numeric input is
// matched against both columns to preserve the existing lookup behavior.
//
// Example:
//
//	asset, err := assetRepository.FindAssetByUniqueIdentifier("ARCA:BIL")
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func (repository *AssetRDBMSRepository) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {

	var queryBuilder = rdbms.BuildQuery[domain.Asset](repository.dbAdapter, assetsSQL)

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

	var queryExecutor = queryBuilder.Build()
	result, err := queryExecutor.GetWithRowScanner(assetRowScanner)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error getting asset by id",
			repository,
		)
	}

	return &result, nil
}

// UpdateAsset updates the ticker and name fields of an existing asset identified by its ID.
// Returns the freshly-read updated asset from the database.
//
// Example:
//
//	updatedAsset, err := assetRepository.UpdateAsset(asset)
//
// Co-authored by: OpenCode and GitHub Copilot
func (repository *AssetRDBMSRepository) UpdateAsset(asset *domain.Asset) (*domain.Asset, error) {

	err := repository.dbAdapter.UpdateListedFields(asset, "Ticker", "Name")
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error updating asset", repository)
	}

	var updatedAsset domain.Asset
	err = repository.dbAdapter.Read(&updatedAsset, asset.Id)
	return &updatedAsset, infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error retrieving updated asset",
		repository,
	)
}

// InsertAssetsInTransaction bulk-inserts new assets within an existing SQL transaction,
// including persisted external data when present, and returns the persisted records.
//
// Example:
//
//	err := adapter.RunInTransaction(func(transContext *rdbms.SQLTransactionalContext) error {
//		persistedAssets, err := assetRepository.InsertAssetsInTransaction(transContext, assets)
//		_ = persistedAssets
//		return err
//	})
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
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

	var columns = []string{"ticker", "name", "external_data"}

	values, tickers, err := buildAssetInsertValues(assets)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting assets",
			repository,
		)
	}

	err = repository.dbAdapter.InsertBulkInTransaction(transactionalContext, "asset", columns, values)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting assets",
			repository,
		)
	}

	return repository.FindAssetsByTickersInTransaction(transactionalContext, tickers)
}

// buildAssetInsertValues prepares the bulk insert values and ticker lookup list for a batch of
// assets.
//
// Authored by: OpenCode
func buildAssetInsertValues(assets []*domain.Asset) ([][]interface{}, []string, error) {
	var values = make([][]interface{}, len(assets))
	var tickers = make([]string, len(assets))

	for i, persistedAsset := range assets {
		assetInsertValue, err := buildAssetInsertValue(persistedAsset)
		if err != nil {
			return nil, nil, err
		}

		values[i] = assetInsertValue
		tickers[i] = persistedAsset.Ticker
	}

	return values, tickers, nil
}

// buildAssetInsertValue prepares a single asset bulk insert row, including the optional external
// data payload.
//
// Authored by: OpenCode
func buildAssetInsertValue(asset *domain.Asset) ([]interface{}, error) {
	var externalDataValue interface{}
	var err error
	if asset.ExternalData != nil {
		externalDataValue, err = asset.ExternalData.Value()
		if err != nil {
			return nil, err
		}
	}

	return []interface{}{
		asset.Ticker,
		asset.Name,
		externalDataValue,
	}, nil
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
