package anticorruption

import (
	"time"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/infra/integration"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

// serviceOrigin is a zero-value pointer used as the origin type reference
// for AppError construction in package-level mapping functions.
var serviceOrigin = (*YahooFinanceAssetIntegrationService)(nil)

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

// QuoteAssetLastClosePrice queries the Yahoo Finance chart API for the last close price of the given
// asset and returns it as a domain ExternalAssetQuote.
// Validates that the asset source matches domain.YahooFinanceSource before proceeding.
//
// Parameters:
//   - asset: the external asset to quote, must have Source set to domain.YahooFinanceSource
//
// Returns:
//   - *domain.ExternalAssetQuote: the asset quote with last close price, date, currency, and identifiers
//   - error: if the asset source does not match, or propagated from the integration client,
//     or if the response structure is invalid
//
// Example:
//
//	var service = &YahooFinanceAssetIntegrationService{
//	    Client: &integration.YahooFinanceAssetIntegrationClient{},
//	}
//	var asset = &domain.ExternalAsset{Source: domain.YahooFinanceSource, Ticker: "AAPL", ExchangeId: "NMS"}
//	quote, err := service.QuoteAssetLastClosePrice(asset)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(quote.LastCloseQuote, quote.Currency)
//
// Co-authored by: GitHub Copilot (claude-opus-4.6) and benizzio
func (service *YahooFinanceAssetIntegrationService) QuoteAssetLastClosePrice(
	asset *domain.ExternalAsset,
) (*domain.ExternalAssetQuote, error) {

	if asset.Source != domain.YahooFinanceSource {
		return nil, infra.BuildAppErrorFormatted(
			service,
			"unexpected asset source %s for Yahoo Finance anticorruption service",
			asset.Source,
		)
	}

	var chartResponse, err = service.Client.QuoteAssetLastClosePrice(asset.Ticker)
	if err != nil {
		return nil, err
	}

	var externalAssetQuote, mapErr = mapToExternalAssetQuote(chartResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return externalAssetQuote, nil
}

// mapToExternalAssetQuote converts a Yahoo Finance chart response DTS to a domain ExternalAssetQuote.
// Extracts the last close price from the indicators and the last timestamp for the close date.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func mapToExternalAssetQuote(
	chartResponse *integration.YahooFinanceChartResponseDTS,
) (*domain.ExternalAssetQuote, error) {

	var results = chartResponse.Chart.Result
	if len(results) == 0 {
		return nil, infra.BuildAppError("Yahoo Finance chart response contains no results", serviceOrigin)
	}

	var result = results[0]

	var currencyUnit, currencyErr = currency.ParseISO(result.Meta.Currency)
	if currencyErr != nil {
		return nil, infra.BuildAppErrorFormatted(serviceOrigin, "error parsing currency %s: %v", result.Meta.Currency, currencyErr)
	}

	lastCloseQuote, lastCloseDate, err := extractLastClose(&result)
	if err != nil {
		return nil, err
	}

	return &domain.ExternalAssetQuote{
		Ticker:         result.Meta.Symbol,
		ExchangeId:     result.Meta.ExchangeName,
		Currency:       currencyUnit,
		LastCloseQuote: lastCloseQuote,
		LastCloseDate:  lastCloseDate,
	}, nil
}

// extractLastClose retrieves the last close price and its corresponding timestamp
// from the chart result indicators and timestamps arrays.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func extractLastClose(
	result *integration.YahooFinanceChartResultDTS,
) (decimal.Decimal, time.Time, error) {

	if len(result.Indicators.Quote) == 0 || len(result.Indicators.Quote[0].Close) == 0 {
		return decimal.Decimal{}, time.Time{}, infra.BuildAppError("Yahoo Finance chart response contains no quote indicators", serviceOrigin)
	}

	if len(result.Timestamps) == 0 {
		return decimal.Decimal{}, time.Time{}, infra.BuildAppError("Yahoo Finance chart response contains no timestamps", serviceOrigin)
	}

	var closePrices = result.Indicators.Quote[0].Close
	var lastCloseQuote = decimal.NewFromFloat(closePrices[len(closePrices)-1])
	var lastCloseDate = time.Unix(result.Timestamps[len(result.Timestamps)-1], 0)

	return lastCloseQuote, lastCloseDate, nil
}
