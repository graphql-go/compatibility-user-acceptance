package extractor

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
)

type Extractor struct {
}

type RunParams struct {
	HTTPClient     *http.Client
	Organization   string
	RepositoryName string
}

type Repository struct {
	StarsCount int
}

type RunResult struct {
	Repository *Repository
}

func New() *Extractor {
	return &Extractor{}
}

func (e *Extractor) Run(p *RunParams) (*RunResult, error) {
	client := github.NewClient(p.HTTPClient)
	ctx := context.Background()
	repo, _, err := client.Repositories.Get(ctx, p.Organization, p.RepositoryName)
	if err != nil {
		return nil, err
	}

	r := &Repository{
		StarsCount: *repo.StargazersCount,
	}

	result := &RunResult{
		Repository: r,
	}

	return result, nil
}
