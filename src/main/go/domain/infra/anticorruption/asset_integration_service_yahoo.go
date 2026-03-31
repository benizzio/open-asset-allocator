package anticorruption

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/infra/integration"
)

// YahooFinanceAssetIntegrationService is the anticorruption layer service that translates
// Yahoo Finance integration DTSs into domain model types.
// Delegates HTTP communication to the YahooFinanceAssetIntegrationClient.
//
// Co-authored by: GitHub Copilot (claude-opus-4.6) and benizzio
type YahooFinanceAssetIntegrationService struct {
	Client *integration.YahooFinanceAssetIntegrationClient
}

// SearchAssets queries the Yahoo Finance API for assets matching the given query value and
// returns them as domain ExternalAsset instances.
//
// Parameters:
//   - queryValue: the search term to query
//
// Returns:
//   - []*domain.ExternalAsset: the list of matched external assets translated from Yahoo Finance data
//   - error: propagated from the integration client if the call fails
//
// Example:
//
//	var service = &YahooFinanceAssetIntegrationService{
//	    Client: &integration.YahooFinanceAssetIntegrationClient{},
//	}
//	assets, err := service.SearchAssets("AAPL")
//	if err != nil {
//	    // handle error
//	}
//	for _, asset := range assets {
//	    fmt.Println(asset.Ticker, asset.ExchangeId)
//	}
//
// Co-authored by: GitHub Copilot (claude-opus-4.6) and benizzio
func (service *YahooFinanceAssetIntegrationService) SearchAssets(
	queryValue string,
) ([]*domain.ExternalAsset, error) {

	var searchResponse, err = service.Client.SearchAssets(queryValue)
	if err != nil {
		return nil, err
	}

	var externalAssets = mapToExternalAssets(searchResponse.Quotes)

	return externalAssets, nil
}

// mapToExternalAssets converts a slice of Yahoo Finance quote DTSs to domain ExternalAsset pointers.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func mapToExternalAssets(quotes []integration.YahooFinanceSearchQuoteDTS) []*domain.ExternalAsset {

	var externalAssets = make([]*domain.ExternalAsset, len(quotes))
	for i, quote := range quotes {
		externalAssets[i] = mapToExternalAsset(&quote)
	}

	return externalAssets
}

// mapToExternalAsset converts a single Yahoo Finance quote DTS to a domain ExternalAsset.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func mapToExternalAsset(quote *integration.YahooFinanceSearchQuoteDTS) *domain.ExternalAsset {
	return &domain.ExternalAsset{
		Source:       domain.YahooFinanceSource,
		Ticker:       quote.Symbol,
		ExchangeId:   quote.Exchange,
		Name:         quote.LongName,
		ExchangeName: quote.ExchDisp,
	}
}
