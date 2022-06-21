package client

import "github.com/cloudquery/cq-provider-sdk/cqproto"

// Provider Configuration

type ResourceConfig struct {
	Name  string
	Other map[string]interface{} `hcl:",inline"`
}

type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	// CHANGEME
	ExampleConfig string `hcl:"example_config"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []ResourceConfig

	requestedFormat cqproto.ConfigFormat
}

func NewConfig(f cqproto.ConfigFormat) *Config {
	return &Config{
		requestedFormat: f,
	}
}

func (c Config) Example() string {
	switch c.requestedFormat {
	case cqproto.ConfigHCL:
		return `configuration {
	// CHANGEME:
	//Here you define your default/example documentation.
	//That is generated with cloudquery init YourProviderName
	// Optional or required parameters
	// debug = false
	// api_key = ""	
}
`
	default:
		return `
CHANGEME:
Here you define your default/example documentation.
That is generated with cloudquery init YourProviderName
Optional or required parameters
debug: false
api_key: ""	
`
	}
}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
