package resources

import (
	"embed"
	// TODO: change the following to your own package
	"github.com/cloudquery/cq-provider-template/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	providerMigrations embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "your_provider_name",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"demo_resource": DemoResource(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
