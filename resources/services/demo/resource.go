package demo

import (
	"context"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource resource --config gen.hcl --output .
func Resources() *schema.Table {
	return &schema.Table{
		Name:     "demo_domain_resource",
		Resolver: fetchDomainResources,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    ResolverPath("AccountId"),
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    Resolver,
			},
			{
				Name:        "name",
				Description: "The name of demo resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "Creation time of the resource",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Metadata.CreateDate"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDomainResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func Resolver(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return nil
}
func ResolverPath(_ string) schema.ColumnResolver {
	return nil
}

type Meta struct {
	// Creation time of the resource. Delete me
	CreateDate time.Time
}
type Resource struct {
	// The name of demo resource
	Name string
	// Metadata of demo resource
	Metadata Meta
}
