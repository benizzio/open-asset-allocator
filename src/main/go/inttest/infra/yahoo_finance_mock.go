// Package infra provides shared infrastructure for integration tests, including the Yahoo
// Finance mock server used by API tests to isolate external asset lookups from live services.
//
// Co-authored by: OpenCode and benizzio
package infra

import (
	"testing"

	"github.com/nhatthm/httpmock"
)

var yahooFinanceMockServer *httpmock.Server

// SetYahooFinanceMockServer stores the shared Yahoo Finance mock server instance for the
// integration test suite.
//
// Authored by: GitHub Copilot
func SetYahooFinanceMockServer(mockServer *httpmock.Server) {
	yahooFinanceMockServer = mockServer
}

// BuildAndStartYahooFinanceMockServer creates and starts the shared Yahoo Finance mock server
// used by integration tests. The server is started once for the suite and individual tests must
// reset its expectations to preserve isolation.
//
// Authored by: GitHub Copilot
func BuildAndStartYahooFinanceMockServer() *httpmock.Server {
	var mockServer = httpmock.NewServer()
	mockServer.WithDefaultResponseHeaders(map[string]string{"Content-Type": "application/json"})
	return mockServer
}

// GetYahooFinanceMockServer returns the shared Yahoo Finance mock server instance.
//
// Authored by: GitHub Copilot
func GetYahooFinanceMockServer() *httpmock.Server {
	return yahooFinanceMockServer
}

// SetupYahooFinanceMockTest resets the shared Yahoo Finance mock server for the given test and
// registers cleanup that verifies all expectations were met and clears state afterwards.
//
// Authored by: GitHub Copilot
func SetupYahooFinanceMockTest(t *testing.T) *httpmock.Server {
	t.Helper()

	var mockServer = GetYahooFinanceMockServer()
	if mockServer == nil {
		t.Fatalf("Yahoo Finance mock server is not initialized; configure it in test bootstrap")
	}

	mockServer.WithTest(t)
	resetYahooFinanceMockServer(mockServer)

	t.Cleanup(func() {
		if err := mockServer.ExpectationsWereMet(); err != nil {
			t.Errorf("Yahoo Finance mock expectations were not met: %v", err)
		}
		resetYahooFinanceMockServer(mockServer)
	})

	return mockServer
}

// resetYahooFinanceMockServer clears all expectations and request history from the shared mock
// server to preserve test isolation between integration tests.
//
// Authored by: GitHub Copilot
func resetYahooFinanceMockServer(mockServer *httpmock.Server) {
	if mockServer == nil {
		return
	}

	mockServer.ResetExpectations()
	mockServer.Requests = nil
}
