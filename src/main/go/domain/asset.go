package domain

type Asset struct {
	Id           int64
	Name         string
	Ticker       string
	ExternalData *ExternalAssetData
}

type AssetsPerTicker map[string]*Asset
