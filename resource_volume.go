package main

import (
	"github`.com/devans10/go-pure-client/pureClient"
        "github.com/hashicorp/terraform/helper/schema"
)

func resourcePureVolume() *schema.Resource {
        return &schema.Resource{
                Create: resourcePureVolumeCreate,
                Read:   resourcePureVolumeRead,
                Update: resourcePureVolumeUpdate,
                Delete: resourcePureVolumeDelete,

                Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"size": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"source": &schema.Schema{
				Type:	  schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:	  schema.TypeString,
				Required: true,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:	  schema.TypeString,
				Required: true,
				Computed: true,
			},
                },
        }
}

func resourcePureVolumeCreate(d *schema.ResourceData, m interface{}) error {
	var v *string

	client := m.(*pureClient.Client)

	v, err = client.Vols.CreateVol(d)
	if err != nil {
		return err
	}

	d.SetId(*v)
        return resourceVolumeRead(d, m)
}

func resourcePureVolumeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*pureClient.Client)

        vol, ok := client.Vols.Read(d.Id())

        if vol == nil {
          d.SetId("")
          return nil
        }

        d.Set("name", vol.Name)
	d.Set("size", vol.Size)
	d.Set("serial", vol.Serial)
	d.Set("created", vol.Created)
        return nil
}

func resourceVolumeUpdate(d *schema.ResourceData, m interface{}) error {
        var v *string

        client := m.(*pureClient.Client)

        v, err = client.Vols.UpdateVol(d)
        if err != nil {
                return err
        }

        d.SetId(*v)
        return resourceVolumeRead(d, m)
}

func resourceVolumeDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(*pureClient.Client)
        err := client.Vols.DeleteVol(d.Id())

        if err != nil {
          return err
        }

	d.SetId("")
        return nil
}
