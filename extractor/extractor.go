package extractor

import (
	"context"

	"github.com/google/go-github/github"
)

type Extractor struct {
}

type RunParams struct {
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
	client := github.NewClient(nil)
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
