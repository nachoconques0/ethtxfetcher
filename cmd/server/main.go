// Package main defines the entrypoint for the service.
// It defines all the required options to run the service and
// the main function that will be executed when the service is started.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nachoconques0/ethtxfetcher/internal/app"
)

func main() {
	opts := []app.Option{
		app.WithHTTPPort(os.Getenv("HTTP_PORT")),
		app.WithNodeURL(os.Getenv("NODE_ETH_URL")),
	}

	application, err := app.New(opts...)
	if err != nil {
		log.Fatal(nil, fmt.Sprintf("error initianting application: %s", err.Error()))
	}

	application.Start()
}
