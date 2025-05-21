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

	// Create a new extractor instance.
	ex := extractor.New(cfg)

	// Configure extraction parameters for the graphql-go/graphql repository.
	params := extractor.RunParams{
		HTTPClient:     &http.Client{},
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}

	// Run the extractor to fetch repository metrics.
	r, err := ex.Run(&params)
	if err != nil {
		log.Fatal(err)
	}

	// Display debug information if enabled.
	fmt.Println(cfg.IsDebug)

	// Display the extraction results.
	fmt.Printf("%+v\n", r)

	header := cfg.GraphqlJSImplementation.Repo.String(implementation.RefImplementationPrefix)
	headerWidth := uint(15)

	cliParams := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{
			Models: bubbletea.Models{
				bubbletea.NewChoicesModel(&bubbletea.ChoicesModelParams{
					Order:   1,
					Choices: cfg.AvailableImplementations,
					UI: bubbletea.ChoicesModelUIParams{
						Header: header,
					},
				}),
				bubbletea.NewTableModel(&bubbletea.TableModelParams{
					Order: 2,
					Headers: []bubbletea.TableHeader{
						{Title: "Metric", Width: 35},
						{Title: "Spec: https://github.com/graphql/graphql-js", Width: headerWidth},
						{Title: "Impl: https://github.com/graphql-go/graphql", Width: headerWidth},
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
				}),
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}

	cli := cmd.New(&cliParams)

	runResult, err := cli.Run(&cmd.RunParams{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(runResult)
}
