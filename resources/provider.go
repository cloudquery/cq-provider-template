package resources

import (
	"embed"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-template/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	providerMigrations embed.FS
	Version            = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version: Version,
		// CHANGEME: Change to your provider name
		Name:      "YourProviderName",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			// CHANGEME: place here all supported resources
			"demo_resource": DemoResource(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
