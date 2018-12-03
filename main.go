package main

import (
	"github.com/devans10/terraform-provider-purestorage/purestorage"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: purestorage.Provider,
	})
}
