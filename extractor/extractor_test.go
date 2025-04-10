// Package extractor tests the functionality of the extractor package.
package extractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

// RoundTripperFunc is a function type that implements the http.RoundTripper interface.
// This allows us to use functions as http.RoundTripper for testing.
type RoundTripperFunc func(*http.Request) (*http.Response, error)

// RoundTrip implements the http.RoundTripper interface.
func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

// testClientParams defines parameters used to configure the test HTTP client.
type testClientParams struct {
	// requestMethod is the HTTP method that should match the incoming request.
	requestMethod string
	// requestURL is the URL that should match the incoming request.
	requestURL string
	// responseBody is a function that returns the mock response body.
	responseBody func() ([]byte, error)
	// responseStatusCode is the HTTP status code to return in the response.
	responseStatusCode int
}

// testClient creates a mock HTTP client that returns predefined responses.
// It allows us to test the Extractor without making actual HTTP requests.
func testClient(p *testClientParams) *http.Client {
	httpClient := http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Default response for non-matching requests
			resp := &http.Response{
				Request:    r,
				StatusCode: p.responseStatusCode,
			}

			// If the request method or URL doesn't match expected values,
			// return the default response without a body
			if r.Method != p.requestMethod || r.URL.String() != p.requestURL {
				return resp, nil
			}

			// Get the mock response body
			data, err := p.responseBody()
			if err != nil {
				return resp, nil
			}

			// Create and return a response with the mock body
			body := bytes.NewReader(data)
			return &http.Response{
				Body:       io.NopCloser(body),
				Request:    r,
				StatusCode: p.responseStatusCode,
			}, nil
		}),
	}

	return &httpClient
}

// TestExtractorRun tests the Run method of the Extractor.
// It covers both successful and error scenarios.
func TestExtractorRun(t *testing.T) {
	tests := []struct {
		subTestName        string
		requestMethod      string
		requestURL         string
		responseBody       func() ([]byte, error)
		responseStatusCode int
		expectedError      error
	}{
		{
			subTestName:   "Success",
			requestMethod: "GET",
			requestURL:    "https://api.github.com/repos/graphql-go/graphql",
			responseBody: func() ([]byte, error) {
				// Mock a GitHub repository with 8 stars
				stargazersCount := int(8)
				repo := github.Repository{
					StargazersCount: &stargazersCount,
				}

				data, err := json.Marshal(repo)
				if err != nil {
					return nil, err
				}

				return data, nil
			},
			responseStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{
			subTestName:   "Handles client repositories get response error",
			requestMethod: "test invalid get method",
			requestURL:    "test invalid url",
			responseBody: func() ([]byte, error) {
				return nil, nil
			},
			responseStatusCode: http.StatusInternalServerError,
			expectedError:      errors.New("GET https://api.github.com/repos/graphql-go/graphql: 500  []"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			// Create an extractor and run it with test parameters
			ex := New()
			params := RunParams{
				HTTPClient: testClient(&testClientParams{
					requestMethod:      tt.requestMethod,
					requestURL:         tt.requestURL,
					responseBody:       tt.responseBody,
					responseStatusCode: tt.responseStatusCode,
				}),
				Organization:   "graphql-go",
				RepositoryName: "graphql",
			}

			// Execute the Run method and check results
			result, err := ex.Run(&params)
			if err != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				return
			}

			assert.NotNil(t, result)
		})
	}

}
