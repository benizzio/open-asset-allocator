// Package infra_test verifies public infrastructure configuration behavior.
// Authored by: OpenCode
package infra_test

import (
	"testing"

	"github.com/benizzio/open-asset-allocator/infra"
)

// TestReadConfigAPIOnly verifies that API_ONLY enables API-only Gin routing only
// when it is set to true. For example, API_ONLY=true omits static-content routes.
// Authored by: OpenCode
func TestReadConfigAPIOnly(t *testing.T) {
	var testCases = []struct {
		name    string
		value   string
		expects bool
	}{
		{name: "empty", value: "", expects: false},
		{name: "false", value: "false", expects: false},
		{name: "true", value: "true", expects: true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Setenv("API_ONLY", testCase.value)

			if actual := infra.ReadConfig().GinServerConfig.ApiOnly; actual != testCase.expects {
				t.Errorf("ReadConfig().GinServerConfig.ApiOnly = %t, want %t", actual, testCase.expects)
			}
		})
	}
}
