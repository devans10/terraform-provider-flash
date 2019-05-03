/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package purestorage

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/devans10/pugo/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureVolumeResourceName = "purestorage_volume.tfvolumetest"
const testAccCheckPureVolumeCloneResourceName = "purestorage_volume.tfclonevolumetest"

// The volumes created in theses tests will not be eradicated.
//
// Create a volume
func TestAccResourcePureVolume_create(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureVolumeConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "name", fmt.Sprintf("tfvolumetest-%d", rInt)),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "size", "1024000000"),
					resource.TestCheckResourceAttrSet(testAccCheckPureVolumeResourceName, "serial"),
				),
			},
		},
	})
}
func TestAccResourcePureVolume_clone(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureVolumeConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "name", fmt.Sprintf("tfvolumetest-%d", rInt)),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "size", "1024000000"),
					resource.TestCheckResourceAttrSet(testAccCheckPureVolumeResourceName, "serial"),
				),
			},
			{
				Config: testAccCheckPureVolumeConfigClone(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeCloneResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeCloneResourceName, "source", fmt.Sprintf("tfvolumetest-%d", rInt)),
				),
			},
		},
	})
}

func TestAccResourcePureVolume_update(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureVolumeConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "name", fmt.Sprintf("tfvolumetest-%d", rInt)),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "size", "1024000000"),
					resource.TestCheckResourceAttrSet(testAccCheckPureVolumeResourceName, "serial"),
				),
			},
			{
				Config: testAccCheckPureVolumeConfigResize(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "name", fmt.Sprintf("tfvolumetest-%d", rInt)),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "size", "2048000000"),
					resource.TestCheckResourceAttrSet(testAccCheckPureVolumeResourceName, "serial"),
				),
			},
			{
				Config: testAccCheckPureVolumeConfigRename(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureVolumeExists(testAccCheckPureVolumeResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "name", fmt.Sprintf("tfvolumetest-rename-%d", rInt)),
					resource.TestCheckResourceAttr(testAccCheckPureVolumeResourceName, "size", "2048000000"),
					resource.TestCheckResourceAttrSet(testAccCheckPureVolumeResourceName, "serial"),
				),
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
		}
		return fmt.Errorf("volume '%s' stil exists", rs.Primary.ID)
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

func testAccCheckPureVolumeConfig(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfvolumetest" {
        name = "tfvolumetest-%d"
        size = 1024000000
}`, rInt)
}

func testAccCheckPureVolumeConfigClone(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfvolumetest" {
        name = "tfvolumetest-%d"
        size = 1024000000
}

resource "purestorage_volume" "tfclonevolumetest" {
        name = "tfclonevolumetest-%d"
        source = "${purestorage_volume.tfvolumetest.name}"
}`, rInt, rInt)
}

func testAccCheckPureVolumeConfigResize(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfvolumetest" {
	name = "tfvolumetest-%d"
	size = 2048000000
}`, rInt)
}

func testAccCheckPureVolumeConfigRename(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfvolumetest" {
        name = "tfvolumetest-rename-%d"
        size = 2048000000
}`, rInt)
}
