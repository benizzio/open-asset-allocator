package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
)

// ================================================
// TYPES
// ================================================

type AssetDTS struct {
	Id     *langext.ParseableInt64 `json:"id"`
	Name   string                  `json:"name" validate:"max=100"`
	Ticker string                  `json:"ticker" validate:"max=40"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapToAssetDTS(asset *domain.Asset) *AssetDTS {

	if asset == nil {
		return nil
	}

	var assetId = langext.ParseableInt64(asset.Id)
	return &AssetDTS{
		Id:     &assetId,
		Name:   asset.Name,
		Ticker: asset.Ticker,
	}
}

func MapToAssetDTSs(assets []*domain.Asset) []*AssetDTS {
	var assetsDTS = make([]*AssetDTS, len(assets))
	for i, asset := range assets {
		assetsDTS[i] = MapToAssetDTS(asset)
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
	for i, assetDTS := range assetsDTS {
		assets[i] = MapToAsset(assetDTS)
	}
	return assets
}
