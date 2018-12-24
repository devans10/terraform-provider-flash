package purestorage

import (
	"fmt"
	"log"

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

// resourcePureVolumeCreate creates a Pure Volume on a FlashArray according
// to the schema Resource Data provided.
// If the size parameter is provided, a new Volume of that size will be created.
// If the source parameter is provided, a new Volume that is a copy of the source
// volume will be created.
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

// resourcePureVolumeRead sets the values for the given volume ID
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

// resourcePureVolumeUpdate will update the attributes of the volume.
//
// If a new source is provided, a snapshot of the current volume will be
// taken before the source volume is copied over the current volume. This
// should help protect from any accidental overwrites.
//
// If a new size is provided, it must be larger than the current size.  Only
// extending volumes is supported at this time, since truncating volumes can
// lead to data loss.
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
		snapshot, err := client.Volumes.CreateSnapshot(d.Id(), "", nil)
		if err != nil {
			return err
		}
		log.Printf("[INFO] Created volume snapshot %s before overwriting volume %s.", snapshot.Name, d.Id())
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
	if z.(int) < oldVol.Size {
		return fmt.Errorf("Error: New size must be larger than current size. Truncating volumes not supported.")
	}

	d.SetId(v.Name)
	return resourcePureVolumeRead(d, m)
}

// resourcePureVolumeDelete will delete the volume specified.
// The volume will NOT be eradicated. This is to reduce the chance of
// data loss.  The volume's timer will start for 24 hours, at that time
// the volume will be eradicated.
func resourcePureVolumeDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	_, err := client.Volumes.DeleteVolume(d.Id(), nil)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// resourcePureVolumeImport imports a volume into Terraform.
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
