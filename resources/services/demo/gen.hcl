service          = "demo"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = [" Delete me"]
}

resource "demo" "domain" "resource" {
  path = "github.com/cloudquery/cq-provider-template/resources/services/demo.Resource"

  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path   = "github.com/cloudquery/cq-provider-template/resources/services/demo.ResolverPath"
      params = ["AccountId"]
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "demoResolver" {
      path = "github.com/cloudquery/cq-provider-template/resources/services/demo.Resolver"
    }
  }

  column "metadata" {
    skip_prefix = true
  }
}