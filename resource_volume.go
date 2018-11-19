package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceVolume() *schema.Resource {
        return &schema.Resource{
                Create: resourceVolumeCreate,
                Read:   resourceVolumeRead,
                Update: resourceVolumeUpdate,
                Delete: resourceVolumeDelete,

                Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"size": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceVolumeCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
        return resourceVolumeRead(d, m)
}

func resourceVolumeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*MyClient)

	// Attempt to read from an upstream API
        obj, ok := client.Get(d.Id())

        // If the resource does not exist, inform Terraform. We want to immediately
        // return here to prevent further processing.
        if !ok {
          d.SetId("")
          return nil
        }

        d.Set("name", obj.Name)
        return nil
}

func resourceVolumeUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceVolumeRead(d, m)
}

func resourceVolumeDelete(d *schema.ResourceData, m interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors, but
        // it is added here for explicitness.
	d.SetId("")
        return nil
}
