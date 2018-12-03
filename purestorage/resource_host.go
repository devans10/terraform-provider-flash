package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePureHost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePureHostCreate,
		Read:   resourcePureHostRead,
		Update: resourcePureHostUpdate,
		Delete: resourcePureHostDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iqn": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"wwn": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"hgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePureHostCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	v, _ := d.GetOk("name")
	h, err := client.Hosts.CreateHost(v.(string), nil)
	if err != nil {
		return err
	}

	d.SetId(h.Name)
	return resourcePureHostRead(d, m)
}

func resourcePureHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	host, _ := client.Hosts.GetHost(d.Id(), nil)

	if host == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	return nil
}

func resourcePureHostUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	v, _ := d.GetOk("name")
	h, err := client.Hosts.RenameHost(d.Id(), v.(string), nil)
	if err != nil {
		return err
	}

	d.SetId(h.Name)
	return resourcePureHostRead(d, m)
}

func resourcePureHostDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	_, err := client.Hosts.DeleteHost(d.Id(), nil)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
