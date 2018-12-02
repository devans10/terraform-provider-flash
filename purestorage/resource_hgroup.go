package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
        "github.com/hashicorp/terraform/helper/schema"
)

func resourcePureHostgroup() *schema.Resource {
        return &schema.Resource{
                Create: resourcePureHostgroupCreate,
                Read:   resourcePureHostgroupRead,
                Update: resourcePureHostgroupUpdate,
                Delete: resourcePureHostgroupDelete,

                Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Elem:	  schema.TypeString,
				Optional: true,
			},
                },
        }
}

func resourcePureHostgroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	v, _ := d.GetOk("name")
	host, err := client.Hostgroups.CreateHostgroup(v.(string), nil)
	if err != nil {
		return err
	}

	d.SetId(host.Name)
        return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

        h, _ := client.Hostgroups.GetHostgroup(d.Id(), nil)

        if h == nil {
          d.SetId("")
          return nil
        }

        d.Set("name", h.Name)
	d.Set("hosts", h.Hosts)
        return nil
}

func resourcePureHostgroupUpdate(d *schema.ResourceData, m interface{}) error {
        client := m.(*flasharray.Client)

	v, _ := d.GetOk("name")
	h, err := client.Hostgroups.RenameHostgroup(d.Id(), v.(string), nil)
        if err != nil {
                return err
        }

        d.SetId(h.Name)
        return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(*flasharray.Client)
        _, err := client.Hostgroups.DeleteHostgroup(d.Id(), nil)

        if err != nil {
          return err
        }

	d.SetId("")
        return nil
}
