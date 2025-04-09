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

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

type testClientParams struct {
	requestMethod      string
	requestURL         string
	responseBody       func() ([]byte, error)
	responseStatusCode int
}

func testClient(p *testClientParams) *http.Client {
	httpClient := http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Request:    r,
				StatusCode: p.responseStatusCode,
			}

			if r.Method != p.requestMethod || r.URL.String() != p.requestURL {
				return resp, nil
			}

			data, err := p.responseBody()
			if err != nil {
				return resp, nil
			}

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
			ex := &Extractor{}
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

			result, err := ex.Run(&params)
			if err != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				return
			}

			assert.NotNil(t, result)
		})
	}

}
