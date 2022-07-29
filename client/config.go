package client

// Provider Configuration
type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	// CHANGEME
	Debug  bool   `yaml:"debug,omitempty"`
	APIKey string `yaml:"api_key"`
}

func (Config) Example() string {
	return `
# CHANGEME:
# Here you define your default/example documentation.
# That is generated with cloudquery init YourProviderName
# Optional or required parameters
# debug: false
# api_key: ""	
`
}
