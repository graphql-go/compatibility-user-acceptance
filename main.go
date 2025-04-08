package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/compatibility-base/config"
	"github.com/graphql-go/compatibility-user-acceptance/extractor"
)

func main() {
	cfg := config.Config{}

	ex := extractor.Extractor{}
	params := extractor.RunParams{
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}
	r, err := ex.Run(&params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.IsDebug)
	fmt.Printf("%+v\n", r)
}
