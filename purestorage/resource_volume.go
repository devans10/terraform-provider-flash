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
		Importer: &schema.ResourceImporter{
			State: resourcePureVolumeImport,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePureVolumeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	var v *flasharray.Volume
	var err error

	n, _ := d.GetOk("name")
	s, _ := d.GetOk("source")
	if s.(string) == "" {
		z, _ := d.GetOk("size")
		v, err = client.Volumes.CreateVolume(n.(string), z.(int), nil)
		if err != nil {
			return err
		}
	} else {
		v, err = client.Volumes.CopyVolume(n.(string), s.(string), false, nil)
		if err != nil {
			return err
		}
	}

	d.SetId(v.Name)
	return resourcePureVolumeRead(d, m)
}

func resourcePureVolumeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	vol, _ := client.Volumes.GetVolume(d.Id(), nil)

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
	var v *flasharray.Volume
	var err error

	oldVol, _ := client.Volumes.GetVolume(d.Id(), nil)

	if oldVol == nil {
		d.SetId("")
		return nil
	}

	n, _ := d.GetOk("name")
	if n.(string) != oldVol.Name {
		v, err = client.Volumes.RenameVolume(d.Id(), n.(string), nil)
		if err != nil {
			return err
		}
	}

	s, _ := d.GetOk("source")
	if s.(string) != oldVol.Source {
		v, err = client.Volumes.CopyVolume(d.Id(), s.(string), true, nil)
		if err != nil {
			return err
		}
	}

	z, _ := d.GetOk("size")
	if z.(int) > oldVol.Size {
		v, err = client.Volumes.ExtendVolume(d.Id(), z.(int), nil)
		if err != nil {
			return err
		}
	}

	d.SetId(v.Name)
	return resourcePureVolumeRead(d, m)
}

func resourcePureVolumeDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	_, err := client.Volumes.DeleteVolume(d.Id(), nil)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourcePureVolumeImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	vol, err := client.Volumes.GetVolume(d.Id(), nil)

	if err != nil {
		return nil, err
	}

	d.Set("name", vol.Name)
	d.Set("size", vol.Size)
	d.Set("serial", vol.Serial)
	d.Set("created", vol.Created)
	d.Set("source", vol.Source)
	return []*schema.ResourceData{d}, nil
}
