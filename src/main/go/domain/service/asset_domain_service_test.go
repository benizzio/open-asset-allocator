package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/infra"
)

// TestGetKnownAssetsRejectsExcessSearchTerms verifies that term-count validation occurs before
// repository access. A nil repository makes any accidental repository call fail the test.
//
// Authored by: OpenCode
func TestGetKnownAssetsRejectsExcessSearchTerms(t *testing.T) {
	var assetService = BuildAssetDomService(nil, nil)
	var textSearch = strings.TrimSpace(strings.Repeat("term ", maxAssetTextSearchTerms+1))

	var assets, err = assetService.GetKnownAssets(textSearch)

	var validationError *infra.DomainValidationError
	require.ErrorAs(t, err, &validationError)
	assert.Nil(t, assets)
	assert.Equal(t, "Asset text search validation failed", validationError.Message)
	assert.Equal(t, "Text search must not exceed 20 terms", validationError.Causes[0].Message)
}

// TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch verifies that syntactically nonblank
// input cannot fall through to the unrestricted repository query.
//
// Authored by: OpenCode
func TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(t *testing.T) {
	var assetService = BuildAssetDomService(nil, nil)

	var assets, err = assetService.GetKnownAssets(`""`)

	assert.NoError(t, err)
	assert.Empty(t, assets)
}
