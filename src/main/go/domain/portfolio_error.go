// Package domain defines the portfolio and allocation concepts used by the application.
// Authored by: GPT-6.1 Terra
package domain

import "fmt"

// PortfolioNotFoundError identifies a portfolio that could not be found and retains the underlying cause.
// For example, callers can use errors.As(err, &notFound) to inspect the portfolio ID,
// and errors.Is to inspect its cause.
// Authored by: GPT-6.1 Terra
type PortfolioNotFoundError struct {
	PortfolioId int64
	Cause       error
}

// Error describes the missing portfolio, for example when logging a failed observation creation.
// Authored by: GPT-6.1 Terra
func (portfolioError *PortfolioNotFoundError) Error() string {
	return fmt.Sprintf("Portfolio %d does not exist", portfolioError.PortfolioId)
}

// Unwrap exposes the underlying cause, for example to errors.Is or errors.As.
// Authored by: GPT-6.1 Terra
func (portfolioError *PortfolioNotFoundError) Unwrap() error {
	return portfolioError.Cause
}
