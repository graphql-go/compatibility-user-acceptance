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
	r, err := ex.Run(&extractor.RunParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.IsDebug)
	fmt.Printf("%+v\n", r)
}
