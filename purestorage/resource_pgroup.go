package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePureProtectiongroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePureProtectiongroupCreate,
		Read:   resourcePureProtectiongroupRead,
		Update: resourcePureProtectiongroupUpdate,
		Delete: resourcePureProtectiongroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hosts": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ConflictsWith: []string{"volumes", "hgroups"},
				Optional:      true,
				Default:       nil,
			},
			"volumes": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ConflictsWith: []string{"hosts", "hgroups"},
				Optional:      true,
				Default:       nil,
			},
			"hgroups": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ConflictsWith: []string{"hosts", "volumes"},
				Optional:      true,
				Default:       nil,
			},
			"targets": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  nil,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
		},
	}
}

func resourcePureProtectiongroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	p, _ := d.GetOk("name")
	var hosts []string
	if h, ok := d.GetOk("hosts"); ok {
		for _, element := range h.([]interface{}) {
			hosts = append(hosts, element.(string))
		}
	}

	var volumes []string
	if v, ok := d.GetOk("volumes"); ok {
		for _, element := range v.([]interface{}) {
			volumes = append(volumes, element.(string))
		}
	}

	var hgroups []string
	if hg, ok := d.GetOk("hgroups"); ok {
		for _, element := range hg.([]interface{}) {
			hgroups = append(hgroups, element.(string))
		}
	}

	var targets []string
	if t, ok := d.GetOk("targets"); ok {
		for _, element := range t.([]interface{}) {
			targets = append(targets, element.(string))
		}
	}

	data := map[string][]string{"hostlist": hosts, "vollist": volumes, "hgrouplist": hgroups, "targets": targets}
	pgroup, err := client.Protectiongroups.CreateProtectiongroup(p.(string), data, nil)
	if err != nil {
		return err
	}

	d.SetId(pgroup.Name)
	return resourcePureProtectiongroupRead(d, m)
}

func resourcePureProtectiongroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	p, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil)

	if p == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", p.Name)
	d.Set("hosts", p.Hosts)
	d.Set("volumes", p.Volumes)
	d.Set("hgroups", p.Hgroups)
	d.Set("source", p.Source)
	d.Set("targets", p.Targets)
	return nil
}

func resourcePureProtectiongroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	v, _ := d.GetOk("name")
	h, err := client.Protectiongroups.RenameProtectiongroup(d.Id(), v.(string), nil)
	if err != nil {
		return err
	}

	var hosts []string
	if h, ok := d.GetOk("hosts"); ok {
		for _, element := range h.([]interface{}) {
			hosts = append(hosts, element.(string))
		}
	}
	data := map[string][]string{"hostlist": hosts}
	_, err = client.Protectiongroups.SetProtectiongroup(v.(string), nil, data)
	if err != nil {
		return err
	}
	d.SetId(h.Name)
	return resourcePureProtectiongroupRead(d, m)
}

func resourcePureProtectiongroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	_, err := client.Protectiongroups.DestroyProtectiongroup(d.Id(), nil)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
