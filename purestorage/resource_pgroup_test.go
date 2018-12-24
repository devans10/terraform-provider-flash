package purestorage

import (
	"fmt"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureProtectiongroupResourceName = "purestorage_protectiongroup.tfprotectiongrouptest"

const testAccCheckPureProtectiongroupConfig = `
resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
	name = "tfprotectiongrouptest"
}
`

const testAccCheckPureProtectiongroupConfigWithHostlist = `
resource "purestorage_host" "tfhosttest" {
	name = "tfhosttest"
}

resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
        name = "tfprotectiongrouptest"
	hostlist = ["${purestorage_host.tfhostest.name}"]
}
`

// Create a protectiongroup
func TestAccResourcePureProtectiongroup_createProtectiongroup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true)),
			},
		},
	})
}

func testAccCheckPureProtectiongroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*flasharray.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "purestorage_protectiongroup" {
			continue
		}

		_, err := client.Protectiongroups.GetProtectiongroup(rs.Primary.ID, nil, nil)
		if err != nil {
			return nil
		} else {
			return fmt.Errorf("protectiongroup '%s' stil exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckPureProtectiongroupExists(n string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		_, err := client.Protectiongroups.GetProtectiongroup(rs.Primary.ID, nil, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("protectiongroup does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}
