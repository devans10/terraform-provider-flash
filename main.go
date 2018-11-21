package main

import (
        "github.com/hashicorp/terraform/plugin"
	"github.com/devan10/terraform-provider-purestorage/purestorage"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: purestorage.Provider
        })
}
