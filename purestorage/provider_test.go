package purestorage

import (
	"github.com/hashicorp/terraform/helper/schema"
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("PURESTORAGE_TARGET"); v == "" {
		t.Fatalf("PURESTORAGE_TARGET must be set for acceptance tests")
	}
	if v := os.Getenv("PURESTORAGE_USERNAME"); v == "" {
		t.Fatalf("PURESTORAGE_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("PURESTORAGE_PASSWORD"); v == "" {
		t.Fatalf("PURESTORAGE_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("PURESTORAGE_APITOKEN"); v == "" {
		t.Fatalf("PURESTORAGE_APITOKEN must be set for acceptance tests")
	}
}
