package purestorage

import (
	"github.com/devans10/go-pure-client/pureClient"
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
                },
        }
}

func resourcePureHostgroupCreate(d *schema.ResourceData, m interface{}) error {
	var v *string

	client := m.(*pureClient.Client)

	v, err = client.Hostgroups.CreateHostgroup(d)
	if err != nil {
		return err
	}

	d.SetId(*v)
        return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*pureClient.Client)

        hgroup, ok := client.Hostgroups.Read(d.Id())

        if hgroup == nil {
          d.SetId("")
          return nil
        }

        d.Set("name", hgroup.Name)
        return nil
}

func resourcePureHostgroupUpdate(d *schema.ResourceData, m interface{}) error {
        var v *string

        client := m.(*pureClient.Client)

        v, err = client.Hostgroups.UpdateHostgroup(d)
        if err != nil {
                return err
        }

        d.SetId(*v)
        return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(*pureClient.Client)
        err := client.Hostgroups.DeleteHostgroup(d.Id())

        if err != nil {
          return err
        }

	d.SetId("")
        return nil
}
