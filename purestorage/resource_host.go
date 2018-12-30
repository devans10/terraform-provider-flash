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
		Importer: &schema.ResourceImporter{
			State: resourcePureHostImport,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iqn": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
			},
			"wwn": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	var wwnlist []string
	if wl, ok := d.GetOk("wwn"); ok {
		for _, element := range wl.([]interface{}) {
			wwnlist = append(wwnlist, element.(string))
		}
	}
	var iqnlist []string
	if il, ok := d.GetOk("iqn"); ok {
		for _, element := range il.([]interface{}) {
			iqnlist = append(iqnlist, element.(string))
		}
	}

	data := map[string][]string{"wwnlist": wwnlist, "iqnlist": iqnlist}
	h, err := client.Hosts.CreateHost(v.(string), data)
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
			_, err = client.Hosts.ConnectHost(h.Name, volume)
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

	host, _ := client.Hosts.GetHost(d.Id())

	if host == nil {
		d.SetId("")
		return nil
	}

	var connected_volumes []string
	cv, _ := client.Hosts.ListHostConnections(host.Name)
	for _, volume := range cv {
		connected_volumes = append(connected_volumes, volume.Vol)
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	d.Set("connected_volumes", connected_volumes)
	return nil
}

func resourcePureHostUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	var h *flasharray.Host
	var err error

	oldHost, _ := client.Hosts.GetHost(d.Id())

	n, _ := d.GetOk("name")
	if n.(string) != oldHost.Name {
		h, err = client.Hosts.RenameHost(d.Id(), n.(string))
		if err != nil {
			return err
		}
		d.SetId(h.Name)
	}

	var wwnlist []string
	wl, _ := d.GetOk("wwn")
	for _, element := range wl.([]interface{}) {
		wwnlist = append(wwnlist, element.(string))
	}
	if !sameStringSlice(wwnlist, oldHost.Wwn) {
		data := map[string]interface{}{"wwnlist": wwnlist}
		h, err = client.Hosts.SetHost(d.Id(), data)
		if err != nil {
			return err
		}
	}

	var iqnlist []string
	il, _ := d.GetOk("iqn")
	for _, element := range il.([]interface{}) {
		iqnlist = append(iqnlist, element.(string))
	}
	if !sameStringSlice(iqnlist, oldHost.Iqn) {
		data := map[string]interface{}{"iqnlist": iqnlist}
		h, err = client.Hosts.SetHost(d.Id(), data)
		if err != nil {
			return err
		}
	}

	var connected_volumes []string
	cv, _ := d.GetOk("connected_volumes")
	for _, element := range cv.([]interface{}) {
		connected_volumes = append(connected_volumes, element.(string))
	}
	var current_volumes []string
	curvols, _ := client.Hosts.ListHostConnections(d.Id())
	for _, volume := range curvols {
		current_volumes = append(current_volumes, volume.Vol)
	}

	if !sameStringSlice(connected_volumes, current_volumes) {
		connect_volumes := difference(connected_volumes, current_volumes)
		for _, volume := range connect_volumes {
			_, err = client.Hosts.ConnectHost(d.Id(), volume)
			if err != nil {
				return err
			}
		}

		disconnect_volumes := difference(current_volumes, connected_volumes)
		for _, volume := range disconnect_volumes {
			_, err = client.Hosts.DisconnectHost(d.Id(), volume)
			if err != nil {
				return err
			}
		}
	}
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
			_, err := client.Hosts.DisconnectHost(d.Id(), volume)
			if err != nil {
				return err
			}
		}
	}

	_, err := client.Hosts.DeleteHost(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourcePureHostImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	host, err := client.Hosts.GetHost(d.Id())

	if err != nil {
		return nil, err
	}

	var connected_volumes []string
	cv, _ := client.Hosts.ListHostConnections(host.Name)
	for _, volume := range cv {
		connected_volumes = append(connected_volumes, volume.Vol)
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	d.Set("connected_volumes", connected_volumes)
	return []*schema.ResourceData{d}, nil
}
