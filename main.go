package main

import (
	"fmt"

	"github.com/graphql-go/compatibility-base/config"
)

func main() {
	cfg := config.Config{}

	fmt.Println(cfg.IsDebug)
}
