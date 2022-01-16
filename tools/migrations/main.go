package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-sdk/migration"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-template/resources"
)

func main() {
	if err := migration.Run(context.Background(), resources.Provider(), ""); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
