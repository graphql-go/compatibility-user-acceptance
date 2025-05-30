// Package main provides the entry point for the compatibility-user-acceptance tool.
// This tool validates compatibility of GraphQL implementations against graphql-js
// by comparing various GitHub repository metrics.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/compatibility-base/bubbletea"
	"github.com/graphql-go/compatibility-base/cmd"
	"github.com/graphql-go/compatibility-base/config"
	"github.com/graphql-go/compatibility-base/implementation"
	"github.com/graphql-go/compatibility-user-acceptance/extractor"
)

// main is the entry point of the application. It initializes the extractor,
// fetches repository metrics from GitHub, and displays the results.
func main() {

	// Load configuration.
	cfg := config.New()

	defaultSpecTableHeader := fmt.Sprintf("Ref: %s", cfg.GraphqlJSImplementation.Repo.URL)
	defaultImplTableHeader := "Impl: https://github.com/graphql-go/graphql"
	choicesModelUIHeader := cfg.GraphqlJSImplementation.Repo.String(implementation.RefImplementationPrefix)

	// Display debug information if enabled.
	if cfg.IsDebug {
		log.Printf("DEBUG: %v", cfg.IsDebug)
	}

	// Create a new extractor instance.
	ex := extractor.New(cfg)

	// Configure extraction parameters for the graphql-go/graphql repository.
	params := extractor.RunParams{
		HTTPClient:     &http.Client{},
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}

	// Run the extractor to fetch repository metrics.
	// TODO(@chris-ramon): Wire extractor run result to the CLI.
	if _, err := ex.Run(&params); err != nil {
		log.Fatal(err)
	}

	cliParams := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{
			Models: bubbletea.Models{
				bubbletea.NewChoicesModel(&bubbletea.ChoicesModelParams{
					Order:   1,
					Choices: cfg.AvailableImplementations,
					UI: bubbletea.ChoicesModelUIParams{
						Header: choicesModelUIHeader,
					},
				}),
				newTableModel(defaultSpecTableHeader, defaultImplTableHeader),
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}

	cli := cmd.New(&cliParams)

	resultCallback := func(result *bubbletea.BubbleTeaResult) error {
		choicesModelUIHeader := result.ChoicesModelResult.Choice
		tableModel := newTableModel(defaultSpecTableHeader, choicesModelUIHeader)

		if err := cli.UpdateModel(tableModel); err != nil {
			log.Printf("failed to update table model: %v", err)
			return err
		}

		return nil
	}

	runParams := &cmd.RunParams{
		ResultCallback: resultCallback,
	}

	if _, err := cli.Run(runParams); err != nil {
		log.Fatal(err)
	}
}

// `newTableModel` creates and returns a pointer to `bubbletea.TableModel`.
func newTableModel(specificationHeader string, implementationHeader string) *bubbletea.TableModel {
	headerWidth := uint(16)

	return bubbletea.NewTableModel(&bubbletea.TableModelParams{
		Order: 2,
		Headers: []bubbletea.TableHeader{
			{Title: "Criteria", Width: 35},
			{Title: specificationHeader, Width: headerWidth},
			{Title: implementationHeader, Width: headerWidth},
			{Title: "Diff Ratio", Width: headerWidth},
			{Title: "Max Diff", Width: headerWidth},
			{Title: "Result", Width: headerWidth},
		},
		Rows: [][]string{
			[]string{"GitHub:", "", "", "", "", ""},
			[]string{"License", "MIT", "MIT", "0%", "0%", "✅"},
			[]string{"Number Of Stars", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Issues Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Issues Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Pull Requests Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Pull Requests Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Forks", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Last Commit Date", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Contributors", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"GraphQL Compatibility Keywords:", "", "", "", "", ""},
			[]string{"Number Of Comments Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"Number Of Comments Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			[]string{"GraphQL:", "", "", "", "", ""},
			[]string{"Specification Version", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		},
	})
}
