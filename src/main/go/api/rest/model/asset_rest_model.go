package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
)

// ================================================
// TYPES
// ================================================

type AssetDTS struct {
	Id           *langext.ParseableInt64 `json:"id"`
	Name         string                  `json:"name" validate:"required,max=100"`
	Ticker       string                  `json:"ticker" validate:"required,max=40"`
	ExternalData *ExternalAssetDataDTS   `json:"externalData,omitempty"`
}

// ExternalAssetDTS is the REST data transfer structure for external asset search results.
// Maps all fields from the domain ExternalAsset, including Name and ExchangeName which are
// excluded from the domain type's JSON serialization (used for persistence) but required in
// API responses.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type ExternalAssetDTS struct {
	Source       string `json:"source" validate:"required"`
	Ticker       string `json:"ticker" validate:"required"`
	ExchangeId   string `json:"exchangeId" validate:"required"`
	Name         string `json:"name,omitempty"`
	ExchangeName string `json:"exchangeName,omitempty"`
}

type ExternalAssetDataDTS struct {
	Data []ExternalAssetDTS `json:"data"`
}

// ExternalAssetSearchQueryDTS is the request data transfer structure for external asset
// search query parameters.
//
// Authored by: GitHub Copilot
type ExternalAssetSearchQueryDTS struct {
	Query string `form:"query" json:"query" validate:"required,max=100"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

// MapToAssetDTS maps a domain Asset to its REST DTS representation, including its optional
// persisted external asset data.
//
// Example:
//
//	var assetDTS = model.MapToAssetDTS(asset)
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func MapToAssetDTS(asset *domain.Asset) *AssetDTS {

	if asset == nil {
		return nil
	}

	var assetId = langext.ParseableInt64(asset.Id)
	return &AssetDTS{
		Id:           &assetId,
		Name:         asset.Name,
		Ticker:       asset.Ticker,
		ExternalData: mapToExternalAssetDataDTS(asset.ExternalData),
	}
}

func MapToAssetDTSs(assets []*domain.Asset) []*AssetDTS {
	var assetsDTS = make([]*AssetDTS, len(assets))
	for index, asset := range assets {
		assetsDTS[index] = MapToAssetDTS(asset)
	}
	return assetsDTS
}

func MapToAsset(assetDTS *AssetDTS) *domain.Asset {

	if assetDTS == nil {
		return nil
	}

	var assetId int64
	if assetDTS.Id != nil {
		assetId = int64(*assetDTS.Id)
	}
	return &domain.Asset{
		Id:     assetId,
		Name:   assetDTS.Name,
		Ticker: assetDTS.Ticker,
	}
}

func MapToAssets(assetsDTS []*AssetDTS) []*domain.Asset {
	var assets = make([]*domain.Asset, len(assetsDTS))
	for index, assetDTS := range assetsDTS {
		assets[index] = MapToAsset(assetDTS)
	}
	return assets
}

// mapToExternalAssetDataDTS maps persisted external asset data to its REST representation.
//
// Authored by: OpenCode
func mapToExternalAssetDataDTS(externalData *domain.ExternalAssetData) *ExternalAssetDataDTS {

	if externalData == nil {
		return nil
	}

	var externalAssetDTSs = make([]ExternalAssetDTS, len(externalData.Data))
	for index := range externalData.Data {
		externalAssetDTSs[index] = *MapToExternalAssetDTS(&externalData.Data[index])
	}

	return &ExternalAssetDataDTS{
		Data: externalAssetDTSs,
	}
}

// MapToExternalAssetDTS maps a domain ExternalAsset to its REST DTS representation.
//
// Parameters:
//   - externalAsset: the domain ExternalAsset to map
//
// Returns:
//   - *ExternalAssetDTS: the mapped REST DTS, or nil if the input is nil
//
// Example:
//
//	var dts = model.MapToExternalAssetDTS(externalAsset)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func MapToExternalAssetDTS(externalAsset *domain.ExternalAsset) *ExternalAssetDTS {

	if externalAsset == nil {
		return nil
	}

	return &ExternalAssetDTS{
		Source:       string(externalAsset.Source),
		Ticker:       externalAsset.Ticker,
		ExchangeId:   externalAsset.ExchangeId,
		Name:         externalAsset.Name,
		ExchangeName: externalAsset.ExchangeName,
	}
}

// MapToExternalAssetDTSs maps a slice of domain ExternalAsset pointers to their REST DTS
// representations.
//
// Parameters:
//   - externalAssets: the slice of domain ExternalAsset pointers to map
//
// Returns:
//   - []*ExternalAssetDTS: the mapped slice of REST DTSs
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func MapToExternalAssetDTSs(externalAssets []*domain.ExternalAsset) []*ExternalAssetDTS {
	var externalAssetDTSs = make([]*ExternalAssetDTS, len(externalAssets))
	for index, externalAsset := range externalAssets {
		externalAssetDTSs[index] = MapToExternalAssetDTS(externalAsset)
	}
	return externalAssetDTSs
}
