package purestorage

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccCheckPureProtectiongroupResourceName = "purestorage_protectiongroup.tfprotectiongrouptest"

// Create a protectiongroup
func TestAccResourcePureProtectiongroup_create(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_create_withHosts(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_withHosts(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureHostExists("purestorage_host.tfpgrouptesthost", true),
					testAccCheckPureProtectiongroupHosts(testAccCheckPureProtectiongroupResourceName, "tfpgrouptesthost", true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_create_withHostgroups(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_withHostgroups(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureHostgroupExists("purestorage_hostgroup.tfpgrouptesthgroup", true),
					testAccCheckPureProtectiongroupHostgroups(testAccCheckPureProtectiongroupResourceName, "tfpgrouptesthgroup", true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_create_withVolumes(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_withVolumes(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureVolumeExists("purestorage_volume.tfpgrouptest-volume", true),
					testAccCheckPureProtectiongroupVolumes(testAccCheckPureProtectiongroupResourceName, fmt.Sprintf("tfpgrouptest-volume-%d", rInt), true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_create_withSchedule(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_withSchedule(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_create_withRetention(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_withRetention(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_update_withHosts(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureProtectiongroupConfig_withHosts(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureHostExists("purestorage_host.tfpgrouptesthost", true),
					testAccCheckPureProtectiongroupHosts(testAccCheckPureProtectiongroupResourceName, "tfpgrouptesthost", true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_update_withHostgroups(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureProtectiongroupConfig_withHostgroups(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureHostgroupExists("purestorage_hostgroup.tfpgrouptesthgroup", true),
					testAccCheckPureProtectiongroupHostgroups(testAccCheckPureProtectiongroupResourceName, "tfpgrouptesthgroup", true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_update_withVolumes(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureProtectiongroupConfig_withVolumes(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
					testAccCheckPureVolumeExists("purestorage_volume.tfpgrouptest-volume", true),
					testAccCheckPureProtectiongroupVolumes(testAccCheckPureProtectiongroupResourceName, fmt.Sprintf("tfpgrouptest-volume-%d", rInt), true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_update_withSchedule(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureProtectiongroupConfig_withSchedule(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
		},
	})
}

func TestAccResourcePureProtectiongroup_update_withRetention(t *testing.T) {
	rInt := rand.Int()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPureProtectiongroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPureProtectiongroupConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
			},
			{
				Config: testAccCheckPureProtectiongroupConfig_withRetention(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPureProtectiongroupExists(testAccCheckPureProtectiongroupResourceName, true),
				),
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
		name := rs.Primary.Attributes["name"]
		_, err := client.Protectiongroups.GetProtectiongroup(name, nil, nil)
		if err != nil {
			if exists {
				return fmt.Errorf("protectiongroup does not exist: %s", n)
			}
			return nil
		}
		return nil
	}
}

func testAccCheckPureProtectiongroupHosts(n string, host string, exists bool) resource.TestCheckFunc {
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
		p, err := client.Protectiongroups.GetProtectiongroup(name, nil, nil)
		if err != nil {
			return fmt.Errorf("protectiongroup does not exist: %s", n)
		}
		if stringInSlice(host, p.Hosts) {
			if exists {
				return nil
			}
			return fmt.Errorf("Host %s still connected to Protection Group %s.", host, name)
		}
		if exists {
			return fmt.Errorf("Host %s not connected to Protection Group %s.", host, name)
		}
		return nil
	}
}

func testAccCheckPureProtectiongroupHostgroups(n string, hostgroup string, exists bool) resource.TestCheckFunc {
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
		p, err := client.Protectiongroups.GetProtectiongroup(name, nil, nil)
		if err != nil {
			return fmt.Errorf("protectiongroup does not exist: %s", name)
		}
		if stringInSlice(hostgroup, p.Hgroups) {
			if exists {
				return nil
			}
			return fmt.Errorf("Hostgroup %s still connected to Protection Group %s.", hostgroup, name)
		}
		if exists {
			return fmt.Errorf("Hostgroup %s not connected to Protection Group %s.", hostgroup, name)
		}
		return nil
	}
}

func testAccCheckPureProtectiongroupVolumes(n string, volume string, exists bool) resource.TestCheckFunc {
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
		p, err := client.Protectiongroups.GetProtectiongroup(name, nil, nil)
		if err != nil {
			return fmt.Errorf("protectiongroup does not exist: %s", n)
		}
		if stringInSlice(volume, p.Volumes) {
			if exists {
				return nil
			}
			return fmt.Errorf("Volume %s still connected to Protection Group.", volume)
		}
		if exists {
			return fmt.Errorf("Volume %s not connected to Protection Group.", volume)
		}
		return nil
	}
}

func testAccCheckPureProtectiongroupConfig_basic(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
        name = "tfprotectiongrouptest-%d"
}`, rInt)
}

func testAccCheckPureProtectiongroupConfig_withHosts(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_host" "tfpgrouptesthost" {
        name = "tfpgrouptesthost"
}

resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
        name = "tfprotectiongrouptest-%d"
        hosts = ["${purestorage_host.tfpgrouptesthost.name}"]
}`, rInt)
}

func testAccCheckPureProtectiongroupConfig_withVolumes(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_volume" "tfpgrouptest-volume" {
	name = "tfpgrouptest-volume-%d"
	size = 1024000000
}

resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
	name = "tfprotectiongrouptest-%d"
	volumes = ["${purestorage_volume.tfpgrouptest-volume.name}"]
}`, rInt, rInt)
}

func testAccCheckPureProtectiongroupConfig_withHostgroups(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_hostgroup" "tfpgrouptesthgroup" {
	name = "tfpgrouptesthgroup"
}

resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
        name = "tfprotectiongrouptest-%d"
        hgroups = ["${purestorage_hostgroup.tfpgrouptesthgroup.name}"]
}`, rInt)
}

func testAccCheckPureProtectiongroupConfig_withSchedule(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
        name = "tfprotectiongrouptest-%d"
	replicate_enabled = "true"
	replicate_at = "3600"
	replicate_frequency = "86400"
	snap_enabled = "true"
	snap_at = "60"
	snap_frequency = "86400"
}`, rInt)
}

func testAccCheckPureProtectiongroupConfig_withRetention(rInt int) string {
	return fmt.Sprintf(`
resource "purestorage_protectiongroup" "tfprotectiongrouptest" {
	name = "tfprotectiongrouptest-%d"
	all_for = 86400
	days = 8
	per_day = 5
}`, rInt)
}
