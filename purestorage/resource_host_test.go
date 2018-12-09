package purestorage

import (
	"fmt"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureHostResourceName = "purestorage_host.tfhosttest"

const testAccCheckPureHostConfig = `
resource "purestorage_host" "tfhosttest" {
	name = "tfhosttest"
}
`

// Create a host
func TestAccResourcePureHost_createHost(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostConfig,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPureHostExists(testAccCheckPureHostResourceName, true)),
			},
		},
	})
}

func testAccCheckPureHostDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*flasharray.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "purestorage_host" {
			continue
		}

		_, err := client.Hosts.GetHost(rs.Primary.ID, nil)
		if err != nil {
			return nil
		} else {
			return fmt.Errorf("host '%s' stil exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckPureHostExists(n string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		_, err := client.Hosts.GetHost(rs.Primary.ID, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("host does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}
