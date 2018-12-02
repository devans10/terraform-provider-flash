package purestorage

import (
	"testing"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/devans10/terraform-provider-purestorage/purestorage"
)

func TestProvider(t *testing.T) {
    if err := purestorage.Provider().(*schema.Provider).InternalValidate(); err != nil {
        t.Fatalf("err: %s", err)
    }
}
