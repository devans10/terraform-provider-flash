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
			"iqn":  &schema.Schema{
                                Type:     schema.TypeString,
                                Required: false,
				Optional: true,
                        },
			"wwn": &schema.Schema{
				Type:	  schema.TypeString,
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
	var v *string

	client := m.(*flasharray.Client)

	v, err = client.Hosts.CreateHost(d)
	if err != nil {
		return err
	}

	d.SetId(*v)
        return resourcePureHostRead(d, m)
}

func resourcePureHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

        host, ok := client.Hosts.Read(d.Id())

        if host == nil {
          d.SetId("")
          return nil
        }

        d.Set("name", host.Name)
	d.Set("iqn", host.Size)
	d.Set("wwn", host.Serial)
        return nil
}

func resourcePureHostUpdate(d *schema.ResourceData, m interface{}) error {
        var v *string

        client := m.(*flasharray.Client)

        v, err = client.Hosts.UpdateHost(d)
        if err != nil {
                return err
        }

        d.SetId(*v)
        return resourcePureHostRead(d, m)
}

func resourcePureHostDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(*flasharray.Client)
        err := client.Hosts.DeleteHost(d.Id())

        if err != nil {
          return err
        }

	d.SetId("")
        return nil
}
