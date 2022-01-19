package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/cloudquery/cq-provider-template/resources/provider"
)

func main() {
	p := provider.Provider()
	serve.Serve(&serve.Options{
		Name:                p.Name,
		Provider:            p,
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
