package purestorage

import (
	"fmt"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureHostgroupResourceName = "purestorage_hostgroup.tfhostgrouptest"

// Create a hostgroup
func TestAccResourcePureHostgroup_create(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_create_withHostlist(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig_withHostlist(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					testAccCheckPureHostgroupHosts(testAccCheckPureHostgroupResourceName, "tfhostgrouptesthost", true),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureHostgroupConfig_rename(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostgroupResourceName, "name", "tfhostgrouptestrename"),
				),
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
		name := rs.Primary.Attributes["name"]
		_, err := client.Hostgroups.GetHostgroup(name, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("hostgroup does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}

func testAccCheckPureHostgroupHosts(n string, host string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		name := rs.Primary.Attributes["name"]
		h, err := client.Hostgroups.GetHostgroup(name, nil)
		if err != nil {
			return fmt.Errorf("hostgroup does not exist: %s", n)
		}

		if stringInSlice(host, h.Hosts) {
			if exists {
				return nil
			}
			return fmt.Errorf("Host %s still a member of hostgroup.", host)
		}
		if exists {
			return fmt.Errorf("Host %s not a member of hostgroup.", host)
		}
		return nil
	}
}

func testAccCheckPureHostgroupConfig() string {
	return fmt.Sprintf(`
resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest"
}`)
}

func testAccCheckPureHostgroupConfig_withHostlist() string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfhostgrouptesthost" {
        name = "tfhostgrouptesthost"
}

resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest"
        hosts = ["${purestorage_host.tfhostgrouptesthost.name}"]
}`)
}

func testAccCheckPureHostgroupConfig_rename() string {
	return fmt.Sprintf(`
resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptestrename"
}`)
}
