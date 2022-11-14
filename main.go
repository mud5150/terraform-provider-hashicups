package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-hashicups/hashicups"
)

// plugin.Serve is the main entrypoint into the provider. It takes a ServeOpts pointer
// which will have a reference to the ProviderFunc. Here this is using an anonymous function
// that returns a  *schema.Provider, which in this case is the hashicups(package) Provider function.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return hashicups.Provider()
		},
	})
}
