package domain

type Asset struct {
	Id     int
	Name   string
	Ticker string
}

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
}
