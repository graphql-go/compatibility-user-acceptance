// Package extractor provides functionality to extract GitHub repository metrics
// for validating compatibility of GraphQL implementations against graphql-js.
package extractor

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
)

// Extractor is responsible for fetching and processing GitHub repository data.
type Extractor struct {
}

// RunParams contains the parameters required to run the extractor.
type RunParams struct {
	// HTTPClient is the HTTP client used for GitHub API requests.
	HTTPClient *http.Client

	// Organization is the GitHub organization name (e.g., "graphql-go").
	Organization string

	// RepositoryName is the GitHub repository name (e.g., "graphql").
	RepositoryName string
}

// Repository contains extracted repository metrics.
type Repository struct {
	// StarsCount is the number of stars on the GitHub repository.
	StarsCount int
}

// RunResult contains the results of running the extractor.
type RunResult struct {
	// Repository contains the extracted repository metrics.
	Repository *Repository
}

// New creates and returns a new Extractor instance.
func New() *Extractor {
	return &Extractor{}
}

// Run executes the extraction process using the provided parameters.
// It fetches repository data from GitHub and returns the extracted metrics.
// Returns an error if the GitHub API request fails.
func (e *Extractor) Run(p *RunParams) (*RunResult, error) {
	// Create a GitHub client using the provided HTTP client.
	client := github.NewClient(p.HTTPClient)
	ctx := context.Background()

	// Fetch repository information from GitHub.
	repo, _, err := client.Repositories.Get(ctx, p.Organization, p.RepositoryName)
	if err != nil {
		return nil, err
	}

	// Extract repository metrics.
	r := &Repository{
		StarsCount: *repo.StargazersCount,
	}

	// Create the result.
	result := &RunResult{
		Repository: r,
	}

	// Return the result.
	return result, nil
}
