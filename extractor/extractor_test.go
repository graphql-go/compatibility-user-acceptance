package extractor

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func TestExtractorRun(t *testing.T) {
	httpClient := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Request: r,
			}
			if r.Method != "GET" || r.URL.String() != "https://api.github.com/repos/graphql-go/graphql" {
				resp.StatusCode = http.StatusNotFound
				return resp, nil
			}

			stargazersCount := int(10)
			repo := github.Repository{
				StargazersCount: &stargazersCount,
			}
			d, err := json.Marshal(repo)
			if err != nil {
				log.Printf("failed to do json marshal: %v", err)
				resp.StatusCode = http.StatusInternalServerError
				return resp, nil
			}

			body := bytes.NewReader(d)
			return &http.Response{
				Body:       io.NopCloser(body),
				Request:    r,
				StatusCode: 200,
			}, nil
		}),
	}

	ex := &Extractor{}
	params := RunParams{
		HTTPClient:     httpClient,
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}

	result, err := ex.Run(&params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
