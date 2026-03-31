package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

// ExecuteGet performs an HTTP GET request to the given URL and validates the response status code.
// Returns the response if the status code is http.StatusOK. For non-200 responses, the response
// body is closed before returning the error.
//
// Parameters:
//   - requestURL: the fully constructed URL to send the GET request to
//
// Returns:
//   - *http.Response: the HTTP response with an open body (caller is responsible for closing)
//   - error: if the request fails or the response status is not 200
//
// Example:
//
//	resp, err := httpclient.ExecuteGet("https://api.example.com/data?q=test")
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(resp)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func ExecuteGet(requestURL string) (*http.Response, error) {

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		CloseResponseBody(resp)
		return nil, fmt.Errorf("HTTP GET request to %s returned status %d", requestURL, resp.StatusCode)
	}

	return resp, nil
}

// DecodeJSONResponse decodes the body of an HTTP response into the target type T.
// Uses json.NewDecoder for stream-based decoding.
//
// Parameters:
//   - resp: the HTTP response with a readable body
//
// Returns:
//   - *T: a pointer to the decoded value
//   - error: if JSON decoding fails
//
// Example:
//
//	type MyResponse struct {
//	    Name string `json:"name"`
//	}
//	resp, err := httpclient.ExecuteGet("https://api.example.com/data")
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(resp)
//	decoded, err := httpclient.DecodeJSONResponse[MyResponse](resp)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func DecodeJSONResponse[T any](resp *http.Response) (*T, error) {

	var result T
	var err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CloseResponseBody closes the HTTP response body and logs any error that occurs.
// Intended to be used in defer statements after obtaining an HTTP response.
//
// Parameters:
//   - resp: the HTTP response whose body should be closed
//
// Example:
//
//	resp, err := httpclient.ExecuteGet("https://api.example.com/data")
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(resp)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func CloseResponseBody(resp *http.Response) {
	var err = resp.Body.Close()
	if err != nil {
		glog.Errorf("Error closing HTTP response body: %v", err)
	}
}
