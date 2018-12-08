package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePureFlashArray() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePureFlashArrayRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the flasharray",
				Optional:    true,
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "ID of the flasharray",
				Computed:    true,
				Optional:    true,
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Version of the flasharray",
				Computed:    true,
				Optional:    true,
			},
			"revision": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Revision of the flasharray",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func dataSourcePureFlashArrayRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	flasharray, err := client.Array.Get(nil)
	if err != nil {
		return err
	}

	d.SetId(flasharray.Id)
	d.Set("name", flasharray.Array_name)
	d.Set("version", flasharray.Version)
	d.Set("revision", flasharray.Revision)
	return nil
}
