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
			"connected_volumes": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  nil,
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

	var connected_volumes []string
	if cv, ok := d.GetOk("connected_volumes"); ok {
		for _, element := range cv.([]interface{}) {
			connected_volumes = append(connected_volumes, element.(string))
		}
	}

	if connected_volumes != nil {
		for _, volume := range connected_volumes {
			_, err = client.Hosts.ConnectHost(h.Name, volume, nil)
			if err != nil {
				return err
			}
		}
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

	var connected_volumes []string
	cv, _ := client.Hosts.ListHostConnections(host.Name, nil)
	for _, volume := range cv {
		connected_volumes = append(connected_volumes, volume.Vol)
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	d.Set("connected_volunmes", connected_volumes)
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

	var connected_volumes []string
	if cv, ok := d.GetOk("connected_volumes"); ok {
		for _, element := range cv.([]interface{}) {
			connected_volumes = append(connected_volumes, element.(string))
		}
	}
	if connected_volumes != nil {
		for _, volume := range connected_volumes {
			_, err := client.Hosts.DisconnectHost(d.Id(), volume, nil)
			if err != nil {
				return err
			}
		}
	}

	_, err := client.Hosts.DeleteHost(d.Id(), nil)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
