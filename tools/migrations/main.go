package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-template/resources/provider"
)

func main() {
	if err := migration.Run(context.Background(), provider.Provider(), ""); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
