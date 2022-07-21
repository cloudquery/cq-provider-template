package provider

import (
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-template/client"
	"github.com/cloudquery/cq-provider-template/resources/services/demo"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version: Version,
		// CHANGEME: Change to your provider name
		Name:      "YourProviderName",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			// CHANGEME: place here all supported resources
			"demo_resource": demo.Resources(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
