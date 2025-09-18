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
