package salesforce

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mocofound/terraform-provider-salesforce/salesforce/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.AzureProvider()
}
