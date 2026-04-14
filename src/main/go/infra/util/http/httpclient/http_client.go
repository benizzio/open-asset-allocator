package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
)

const defaultTimeout = 10 * time.Second

// RequestOption is a functional option that configures an HTTP request before execution.
// Used with ExecuteGet and ExecuteGetJSON to customize request headers, parameters, or
// other properties without modifying the function signatures for each new requirement.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type RequestOption func(*http.Request)

// WithHeader returns a RequestOption that sets a single header key-value pair on the request.
// If the header already exists, it is replaced.
//
// Parameters:
//   - key: the header name (e.g., "User-Agent", "Accept")
//   - value: the header value
//
// Returns:
//   - RequestOption: a function that applies the header to a request
//
// Example:
//
//	response, err := httpclient.ExecuteGet(context.Background(), url,
//	    httpclient.WithHeader("User-Agent", "Mozilla/5.0"),
//	    httpclient.WithHeader("Accept", "application/json"),
//	)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func WithHeader(key string, value string) RequestOption {
	return func(request *http.Request) {
		request.Header.Set(key, value)
	}
}

// ExecuteGet performs an HTTP GET request to the given URL and validates the response status code.
// Returns the response if the status code is http.StatusOK. For non-200 responses, the response
// body is closed before returning the error. Accepts variadic RequestOption functions to customize
// the request before execution.
//
// Parameters:
//   - requestURL: the fully constructed URL to send the GET request to
//   - options: variadic functional options applied to the request before execution
//
// Returns:
//   - *http.Response: the HTTP response with an open body (caller is responsible for closing)
//   - error: if the request fails or the response status is not 200
//
// Example:
//
//	response, err := httpclient.ExecuteGet(context.Background(), "https://api.example.com/data?q=test",
//	    httpclient.WithHeader("User-Agent", "MyApp/1.0"),
//	)
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(response)
//
// Co-authored by: OpenCode and benizzio
func ExecuteGet(requestContext context.Context, requestURL string, options ...RequestOption) (*http.Response, error) {

	var request, err = http.NewRequestWithContext(requestContext, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP GET request for %s: %w", requestURL, err)
	}

	for _, option := range options {
		option(request)
	}

	var client = &http.Client{Timeout: defaultTimeout}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		CloseResponseBody(response)
		return nil, fmt.Errorf("HTTP GET request to %s returned status %d", requestURL, response.StatusCode)
	}

	return response, nil
}

// ExecuteGetJSON performs an HTTP GET request to the given URL, validates the response,
// decodes the JSON body into the target type T, and closes the response body.
// This is a convenience function that combines ExecuteGet, DecodeJSONResponse, and
// CloseResponseBody into a single call. Accepts variadic RequestOption functions to customize
// the request before execution.
//
// Parameters:
//   - requestURL: the fully constructed URL to send the GET request to
//   - options: variadic functional options applied to the request before execution
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
//	response, err := httpclient.ExecuteGetJSON[SearchResponse](
//	    context.Background(),
//	    "https://api.example.com/search?q=test",
//	    httpclient.WithHeader("User-Agent", "MyApp/1.0"),
//	)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(response.Results)
//
// Co-authored by: OpenCode and benizzio
func ExecuteGetJSON[T any](requestContext context.Context, requestURL string, options ...RequestOption) (*T, error) {

	var response, err = ExecuteGet(requestContext, requestURL, options...)
	if err != nil {
		return nil, err
	}
	defer CloseResponseBody(response)

	return DecodeJSONResponse[T](response)
}

// DecodeJSONResponse decodes the body of an HTTP response into the target type T.
// Uses json.NewDecoder for stream-based decoding.
//
// Parameters:
//   - response: the HTTP response with a readable body
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
//	response, err := httpclient.ExecuteGet(context.Background(), "https://api.example.com/data")
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(response)
//	decoded, err := httpclient.DecodeJSONResponse[MyResponse](response)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func DecodeJSONResponse[T any](response *http.Response) (*T, error) {

	var result T
	var err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CloseResponseBody closes the HTTP response body and logs any error that occurs.
// Intended to be used in defer statements after obtaining an HTTP response.
//
// Parameters:
//   - response: the HTTP response whose body should be closed
//
// Example:
//
//	response, err := httpclient.ExecuteGet(context.Background(), "https://api.example.com/data")
//	if err != nil {
//	    // handle error
//	}
//	defer httpclient.CloseResponseBody(response)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func CloseResponseBody(response *http.Response) {
	var err = response.Body.Close()
	if err != nil {
		glog.Errorf("Error closing HTTP response body: %v", err)
	}
}
