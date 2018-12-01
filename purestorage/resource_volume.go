package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
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
				Optional: true,
				Computed: true,
                        },
			"source": &schema.Schema{
				Type:	  schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:	  schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:	  schema.TypeString,
				Optional: true,
				Computed: true,
			},
                },
        }
}

func resourcePureVolumeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	if d.source == "" {
		v, err := client.Vols.CreateVolume(d.name, d.size, nil)
		if err != nil {
			return err
		}
	} else {
		v, err := client.Vols.CopyVolume(d.name, d.source, nil)
		if err != nil {
                        return err
                }
	}

	d.SetId(v.Name)
        return resourcePureVolumeRead(d, m)
}

func resourcePureVolumeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

        vol, ok := client.Vols.GetVolume(d.Id(), nil)

        if vol == nil {
          d.SetId("")
          return nil
        }

        d.Set("name", vol.Name)
	d.Set("size", vol.Size)
	d.Set("serial", vol.Serial)
	d.Set("created", vol.Created)
	d.Set("source", vol.Source)
        return nil
}

func resourcePureVolumeUpdate(d *schema.ResourceData, m interface{}) error {
        client := m.(*flasharray.Client)

	oldVol, ok := client.Vols.GetVolume(d.Id(), nil)

	if vol == nil {
          d.SetId("")
          return nil
        }

	if d.name != oldVol.Name {
		v, err := client.Vols.RenameVolume(d.Id(), d.name)
		if  err != nil {
			return err
		}
	}

        d.SetId(v.Name))
        return resourcePureVolumeRead(d, m)
}

func resourcePureVolumeDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(*flasharray.Client)
        _, err := client.Vols.DeleteVolume(d.Id())

        if err != nil {
          return err
        }

	d.SetId("")
        return nil
}
