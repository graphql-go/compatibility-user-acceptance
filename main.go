// Package main provides the entry point for the compatibility-user-acceptance tool.
// This tool validates compatibility of GraphQL implementations against graphql-js
// by comparing various GitHub repository metrics.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/compatibility-base/config"
	"github.com/graphql-go/compatibility-user-acceptance/extractor"
)

// main is the entry point of the application. It initializes the extractor,
// fetches repository metrics from GitHub, and displays the results.
func main() {
	// Load configuration
	cfg := config.Config{}

	// Create a new extractor instance
	ex := extractor.New()
	
	// Configure extraction parameters for the graphql-go/graphql repository
	params := extractor.RunParams{
		HTTPClient:     &http.Client{},
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}
	
	// Run the extractor to fetch repository metrics
	r, err := ex.Run(&params)
	if err != nil {
		log.Fatal(err)
	}

	// Display debug information if enabled
	fmt.Println(cfg.IsDebug)
	
	// Display the extraction results
	fmt.Printf("%+v\n", r)
}
