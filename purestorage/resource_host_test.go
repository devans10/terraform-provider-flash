package purestorage

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureHostResourceName = "purestorage_host.tfhosttest"

// Create a host
func TestAccResourcePureHost_create(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostExists(testAccCheckPureHostResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostResourceName, "name", fmt.Sprintf("tfhosttest%d", rInt)),
				),
			},
		},
	})
}

func TestAccResourcePureHost_createWithWWN(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostConfig_withWWN(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostExists(testAccCheckPureHostResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostResourceName, "name", fmt.Sprintf("tfhosttest%d", rInt)),
					testAccCheckPureHostWWN(testAccCheckPureHostResourceName, "0000999900009999", true),
				),
			},
		},
	})
}

func TestAccResourcePureHost_update(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostExists(testAccCheckPureHostResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostResourceName, "name", fmt.Sprintf("tfhosttest%d", rInt)),
				),
			},
			{
				Config: testAccCheckPureHostConfig_withWWN(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostWWN(testAccCheckPureHostResourceName, "0000999900009999", true),
				),
			},
			{
				Config: testAccCheckPureHostConfig_rename(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testAccCheckPureHostResourceName, "name", fmt.Sprintf("tfhosttestrename%d", rInt)),
				),
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
		name, ok := rs.Primary.Attributes["name"]
		_, err := client.Hosts.GetHost(name, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("host does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}

func testAccCheckPureHostWWN(n string, wwn string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		name, ok := rs.Primary.Attributes["name"]
		h, err := client.Hosts.GetHost(name, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("host does not exist: %s", n)
			}
			return nil
		}
		if stringInSlice(wwn, h.Wwn) {
			return nil
		}
		return fmt.Errorf("wwn %s not set for host.", wwn)
	}
}

func testAccCheckPureHostConfig_basic(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfhosttest" {
        name = "tfhosttest%d"
}`, rInt)
}

func testAccCheckPureHostConfig_rename(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfhosttest" {
        name = "tfhosttestrename%d"
	wwn = ["0000999900009999"]
}`, rInt)
}

func testAccCheckPureHostConfig_withWWN(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfhosttest" {
        name = "tfhosttest%d"
	wwn = ["0000999900009999"]
}`, rInt)
}
