package extractor

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

type Extractor struct {
}

type RunParams struct {
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

func (e *Extractor) Run() (*RunResult, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	repo, _, err := client.Repositories.Get(ctx, "graphql-go", "graphql")
	if err != nil {
		log.Fatal(err)
	}

	r := &Repository{
		StarsCount: *repo.StargazersCount,
	}

	result := &RunResult{
		Repository: r,
	}

	return result, nil
}
