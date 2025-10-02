package model

import "github.com/benizzio/open-asset-allocator/domain"

// ================================================
// TYPES
// ================================================

type AssetDTS struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapToAssetDTS(asset *domain.Asset) *AssetDTS {

	if asset == nil {
		return nil
	}

	return &AssetDTS{
		Id:     asset.Id,
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

	return &domain.Asset{
		Id:     assetDTS.Id,
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
