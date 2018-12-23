package purestorage

import (
	"fmt"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureVolumeResourceName = "purestorage_volume.tfvolumetest"

const testAccCheckPureVolumeConfig = `
resource "purestorage_volume" "tfvolumetest" {
	name = "tfvolumetest"
	size = 1024000000
}
`

// The volumes created in theses tests will not be eradicated.
// So to run these tests multiple times within 24 hours, the
// tests volumes will have to be eradicated manually.
//
// Create a volume
func TestAccResourcePureVolume_createVolume(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureVolumeConfig,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true)),
			},
		},
	})
}

func testAccCheckPureVolumeDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*flasharray.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "purestorage_volume" {
			continue
		}

		_, err := client.Volumes.GetVolume(rs.Primary.ID, nil)
		if err != nil {
			return nil
		} else {
			return fmt.Errorf("volume '%s' stil exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckPureVolumeExists(n string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*flasharray.Client)
		_, err := client.Volumes.GetVolume(rs.Primary.ID, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("volume does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}
