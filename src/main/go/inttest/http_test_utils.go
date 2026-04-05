package inttest

import (
	"net/http"

	"github.com/golang/glog"
)

// deferCloseResponseBody closes an HTTP response body and logs any error.
// Intended to be used as a deferred call after an HTTP request in integration tests.
//
// Usage:
//
//	response, err := http.Get(url)
//	assert.NoError(t, err)
//	defer deferCloseResponseBody(response)
//
// Co-authored by: GitHub Copilot
func deferCloseResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		glog.Errorf("Error closing response body: %v", err)
	}
}
