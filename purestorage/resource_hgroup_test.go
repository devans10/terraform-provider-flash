package purestorage

import (
	"fmt"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureHostgroupResourceName = "purestorage_hostgroup.tfhostgrouptest"

const testAccCheckPureHostgroupConfig = `
resource "purestorage_hostgroup" "tfhostgrouptest" {
	name = "tfhostgrouptest"
}
`

const testAccCheckPureHostgroupConfigWithHostlist = `
resource "purestorage_host" "tfhosttest" {
	name = "tfhosttest"
}

resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest"
	hostlist = ["${purestorage_host.tfhostest.name}"]
}
`

// Create a hostgroup
func TestAccResourcePureHostgroup_createHostgroup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true)),
			},
		},
	})
}

func testAccCheckPureHostgroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*flasharray.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "purestorage_hostgroup" {
			continue
		}

		_, err := client.Hostgroups.GetHostgroup(rs.Primary.ID, nil)
		if err != nil {
			return nil
		} else {
			return fmt.Errorf("hostgroup '%s' stil exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckPureHostgroupExists(n string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		_, err := client.Hostgroups.GetHostgroup(rs.Primary.ID, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("hostgroup does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}
