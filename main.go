package main

import (
        "github.com/hashicorp/terraform/plugin"
	"github.com/devans10/terraform-provider-purestorage/purestorage"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: purestorage.Provider,
        })
}
