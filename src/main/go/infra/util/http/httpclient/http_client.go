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

// ExecuteGetJSON performs an HTTP GET request to the given URL, validates the response,
// decodes the JSON body into the target type T, and closes the response body.
// This is a convenience function that combines ExecuteGet, DecodeJSONResponse, and
// CloseResponseBody into a single call.
//
// Parameters:
//   - requestURL: the fully constructed URL to send the GET request to
//
// Returns:
//   - *T: a pointer to the decoded JSON response value
//   - error: if the request fails, returns a non-200 status, or JSON decoding fails
//
// Example:
//
//	type SearchResponse struct {
//	    Results []string `json:"results"`
//	}
//	response, err := httpclient.ExecuteGetJSON[SearchResponse]("https://api.example.com/search?q=test")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(response.Results)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func ExecuteGetJSON[T any](requestURL string) (*T, error) {

	var resp, err = ExecuteGet(requestURL)
	if err != nil {
		return nil, err
	}
	defer CloseResponseBody(resp)

	return DecodeJSONResponse[T](resp)
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
