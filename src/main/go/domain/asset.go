package domain

type Asset struct {
	Id     int
	Name   string
	Ticker string
}

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetById(id int) (*Asset, error)
}
