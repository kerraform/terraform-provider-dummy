package dummy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:               map[string]*schema.Schema{},
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureContextFunc,
	}
}

func providerConfigureContextFunc(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, nil
}
