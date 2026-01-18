package model

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// TestPortfolioDTSMaxLengthValidation tests the max length validation for PortfolioDTS fields.
func TestPortfolioDTSMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		portfolioName  string
		shouldPass  bool
	}{
		{
			name:        "valid name - under max length",
			portfolioName:  "My Portfolio",
			shouldPass:  true,
		},
		{
			name:        "valid name - exactly at max length",
			portfolioName:  strings.Repeat("a", 100),
			shouldPass:  true,
		},
		{
			name:        "invalid name - exceeds max length",
			portfolioName:  strings.Repeat("a", 101),
			shouldPass:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			portfolio := PortfolioDTS{
				Name: tt.portfolioName,
			}
			err := validate.Struct(portfolio)

			if tt.shouldPass {
				// For required fields, we only check max when there's content
				if err != nil && !hasOnlyRequiredErrors(err) {
					t.Errorf("Expected validation to pass, but got error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected validation to fail, but it passed")
				}
			}
		})
	}
}

// TestAssetDTSMaxLengthValidation tests the max length validation for AssetDTS fields.
func TestAssetDTSMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name       string
		ticker     string
		assetName  string
		shouldPass bool
	}{
		{
			name:       "valid ticker and name",
			ticker:     "AAPL",
			assetName:  "Apple Inc.",
			shouldPass: true,
		},
		{
			name:       "ticker at max length",
			ticker:     strings.Repeat("A", 20),
			assetName:  "Test Asset",
			shouldPass: true,
		},
		{
			name:       "ticker exceeds max length",
			ticker:     strings.Repeat("A", 21),
			assetName:  "Test Asset",
			shouldPass: false,
		},
		{
			name:       "name at max length",
			ticker:     "TEST",
			assetName:  strings.Repeat("a", 100),
			shouldPass: true,
		},
		{
			name:       "name exceeds max length",
			ticker:     "TEST",
			assetName:  strings.Repeat("a", 101),
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := AssetDTS{
				Ticker: tt.ticker,
				Name:   tt.assetName,
			}
			err := validate.Struct(asset)

			if tt.shouldPass {
				if err != nil {
					t.Errorf("Expected validation to pass, but got error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected validation to fail, but it passed")
				}
			}
		})
	}
}

// TestAllocationPlanDTSMaxLengthValidation tests the max length validation for AllocationPlanDTS fields.
func TestAllocationPlanDTSMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name       string
		planName   string
		planType   string
		shouldPass bool
	}{
		{
			name:       "valid plan name and type",
			planName:   "My Allocation Plan",
			planType:   "ALLOCATION_PLAN",
			shouldPass: true,
		},
		{
			name:       "plan name at max length",
			planName:   strings.Repeat("a", 100),
			planType:   "ALLOCATION_PLAN",
			shouldPass: true,
		},
		{
			name:       "plan name exceeds max length",
			planName:   strings.Repeat("a", 101),
			planType:   "ALLOCATION_PLAN",
			shouldPass: false,
		},
		{
			name:       "plan type at max length",
			planName:   "Test Plan",
			planType:   strings.Repeat("T", 50),
			shouldPass: true,
		},
		{
			name:       "plan type exceeds max length",
			planName:   "Test Plan",
			planType:   strings.Repeat("T", 51),
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := AllocationPlanDTS{
				Name: tt.planName,
				Type: tt.planType,
			}
			err := validate.Struct(plan)

			if tt.shouldPass {
				// For required fields with min constraints, we may get validation errors
				// We only care about max length failures
				if err != nil && hasMaxLengthError(err) {
					t.Errorf("Expected max length validation to pass, but got error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected validation to fail, but it passed")
				}
				if !hasMaxLengthError(err) {
					t.Errorf("Expected max length error, but got different error: %v", err)
				}
			}
		})
	}
}

// TestPortfolioAllocationDTSMaxLengthValidation tests the max length validation for PortfolioAllocationDTS fields.
func TestPortfolioAllocationDTSMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		class       string
		assetTicker string
		assetName   string
		shouldPass  bool
	}{
		{
			name:        "valid class",
			class:       "STOCKS",
			assetTicker: "AAPL",
			assetName:   "Apple Inc.",
			shouldPass:  true,
		},
		{
			name:        "class at max length",
			class:       strings.Repeat("C", 100),
			assetTicker: "TEST",
			assetName:   "Test Asset",
			shouldPass:  true,
		},
		{
			name:        "class exceeds max length",
			class:       strings.Repeat("C", 101),
			assetTicker: "TEST",
			assetName:   "Test Asset",
			shouldPass:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := PortfolioAllocationDTS{
				Class:       tt.class,
				AssetTicker: tt.assetTicker,
				AssetName:   tt.assetName,
			}
			err := validate.Struct(allocation)

			if tt.shouldPass {
				// For required fields, we may get validation errors
				// We only care about max length failures
				if err != nil && hasMaxLengthError(err) {
					t.Errorf("Expected max length validation to pass, but got error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected validation to fail, but it passed")
				}
			}
		})
	}
}

// hasOnlyRequiredErrors checks if validation errors are only 'required' errors
func hasOnlyRequiredErrors(err error) bool {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}
	for _, fieldError := range validationErrors {
		if fieldError.Tag() != "required" && fieldError.Tag() != "min" {
			return false
		}
	}
	return true
}

// hasMaxLengthError checks if any of the validation errors is a max length error
func hasMaxLengthError(err error) bool {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}
	for _, fieldError := range validationErrors {
		if fieldError.Tag() == "max" {
			return true
		}
	}
	return false
}
