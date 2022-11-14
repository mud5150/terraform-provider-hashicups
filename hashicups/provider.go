package hashicups

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function returns a *schema.Provider. Here we are creating a schema.Provider in
// place. The schema provider will have references to all configuration, resources, and data sources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{ // Data source map will reference all data sources
			"hashicups_coffees": dataSourceCoffees(), // via their "main" function that returns a schema.Resource
			"hashicups_order":   dataSourceOrder(),
		},
		ConfigureContextFunc: providerConfigure, // This configures the client for the provider and allows retrieval via provider meta
	}
}

//
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := hashicups.NewClient(nil, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to auth user for authenticated HashiCups client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := hashicups.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
