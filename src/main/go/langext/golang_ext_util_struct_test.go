package langext

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindStructFieldByNameOrJSONName verifies that struct fields can be resolved both by their
// Go field name and by their exported JSON tag name.
//
// Authored by: OpenCode
func TestFindStructFieldByNameOrJSONName(t *testing.T) {
	type nestedDTS struct {
		ExchangeID string `json:"exchangeId"`
	}

	type sampleDTS struct {
		AssetID *int64     `json:"assetId"`
		Nested  *nestedDTS `json:"nested,omitempty"`
	}

	t.Run(
		"FindByGoFieldName",
		func(t *testing.T) {
			structField, found := FindStructFieldByNameOrJSONName(reflect.TypeOf(sampleDTS{}), "AssetID")

			assert.True(t, found)
			assert.Equal(t, "AssetID", structField.Name)
		},
	)

	t.Run(
		"FindByJSONFieldName",
		func(t *testing.T) {
			structField, found := FindStructFieldByNameOrJSONName(reflect.TypeOf(sampleDTS{}), "assetId")

			assert.True(t, found)
			assert.Equal(t, "AssetID", structField.Name)
		},
	)

	t.Run(
		"FindWithPointerType",
		func(t *testing.T) {
			structField, found := FindStructFieldByNameOrJSONName(reflect.TypeOf(&sampleDTS{}), "nested")

			assert.True(t, found)
			assert.Equal(t, "Nested", structField.Name)
		},
	)

	t.Run(
		"MissingField",
		func(t *testing.T) {
			_, found := FindStructFieldByNameOrJSONName(reflect.TypeOf(sampleDTS{}), "missing")

			assert.False(t, found)
		},
	)
}
