package domain_infra

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
)

type PortfolioRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *PortfolioRepository) GetAllPortfolioSlices(limit int) []domain.PortfolioSliceAtTime {
	// TODO continue with data retrieval using dbx (boilerplate methods in adapter?)
	return nil
}

func BuildPortfolioRepository(dbAdapter *infra.RDBMSAdapter) *PortfolioRepository {
	return &PortfolioRepository{dbAdapter: dbAdapter}
}
