package resources_test

import (
	"testing"

	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-template/resources"

	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrations(t *testing.T) {
	migration.RunMigrationsTest(t, resources.Provider(), []string{"latest"})
}
