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

const testAccCheckPureHostgroupResourceName = "purestorage_hostgroup.tfhostgrouptest"

// Create a hostgroup
func TestAccResourcePureHostgroup_create(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_create_withHostlist(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfigWithHostlist(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					testAccCheckPureHostgroupHosts(testAccCheckPureHostgroupResourceName, fmt.Sprintf("tfhostgrouptesthost%d", rInt), true),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_create_withVolumes(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfigWithVolumes(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostgroupResourceName, "name", fmt.Sprintf("tfhostgrouptest%d", rInt)),
					testAccCheckPureHostgroupVolumes(testAccCheckPureHostgroupResourceName, fmt.Sprintf("tfhostgrouptest-volume-%d", rInt), true),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_update(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureHostgroupConfigRename(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostgroupResourceName, "name", "tfhostgrouptestrename"),
				),
			},
		},
	})
}

func TestAccResourcePureHostgroup_update_AddandRemoveVolume(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureHostgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureHostgroupConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureHostgroupConfigWithVolumes(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostgroupResourceName, "name", fmt.Sprintf("tfhostgrouptest%d", rInt)),
					testAccCheckPureHostgroupVolumes(testAccCheckPureHostgroupResourceName, fmt.Sprintf("tfhostgrouptest-volume-%d", rInt), true),
				),
			},
			{
				Config: testAccCheckPureHostgroupConfigWithoutVolumes(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureHostgroupExists(testAccCheckPureHostgroupResourceName, true),
					resource.TestCheckResourceAttr(testAccCheckPureHostgroupResourceName, "name", fmt.Sprintf("tfhostgrouptest%d", rInt)),
					testAccCheckPureHostgroupVolumes(testAccCheckPureHostgroupResourceName, fmt.Sprintf("tfhostgrouptest-volume-%d", rInt), false),
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
		}
		return fmt.Errorf("hostgroup '%s' stil exists", rs.Primary.ID)

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
			return fmt.Errorf("host %s still a member of hostgroup", host)
		}
		if exists {
			return fmt.Errorf("host %s not a member of hostgroup", host)
		}
		return nil
	}
}

func testAccCheckPureHostgroupVolumes(n string, volume string, exists bool) resource.TestCheckFunc {
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
		h, err := client.Hostgroups.ListHostgroupConnections(name)
		if err != nil {
			return fmt.Errorf("hostgroup does not exist: %s", n)
		}

		for _, elem := range h {
			if elem.Vol == volume {
				if exists {
					return nil
				}
				return fmt.Errorf("volume %s still connected to hostgroup", volume)
			}
		}
		if exists {
			return fmt.Errorf("volume %s not connected to hostgroup", volume)
		}
		return nil
	}
}

func testAccCheckPureHostgroupConfig(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest%d"
}`, rInt)
}

func testAccCheckPureHostgroupConfigWithHostlist(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfhostgrouptesthost" {
        name = "tfhostgrouptesthost%d"
}

resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest%d"
        hosts = ["${purestorage_host.tfhostgrouptesthost.name}"]
}`, rInt, rInt)
}

func testAccCheckPureHostgroupConfigRename() string {
	return fmt.Sprintf(`
resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptestrename"
}`)
}

func testAccCheckPureHostgroupConfigWithVolumes(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfhostgrouptest-volume" {
	name = "tfhostgrouptest-volume-%d"
	size = 1024000000
}

resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest%d"
	volume {
		vol = "${purestorage_volume.tfhostgrouptest-volume.name}"
		lun = 250
	}
}`, rInt, rInt)
}

func testAccCheckPureHostgroupConfigWithoutVolumes(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfhostgrouptest-volume" {
        name = "tfhostgrouptest-volume-%d"
        size = 1024000000
}

resource "purestorage_hostgroup" "tfhostgrouptest" {
        name = "tfhostgrouptest%d"
}`, rInt, rInt)
}
