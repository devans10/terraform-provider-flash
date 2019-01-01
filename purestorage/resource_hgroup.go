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
		Importer: &schema.ResourceImporter{
			State: resourcePureHostgroupImport,
		},

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
				Optional: true,
				Default:  nil,
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

func resourcePureHostgroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	var hgroup *flasharray.Hostgroup
	var err error

	var hosts []string
	if h, ok := d.GetOk("hosts"); ok {
		for _, element := range h.([]interface{}) {
			hosts = append(hosts, element.(string))
		}
	}
	data := map[string][]string{"hostlist": hosts}
	if hgroup, err = client.Hostgroups.CreateHostgroup(d.Get("name").(string), data); err != nil {
		return err
	}
	d.SetId(hgroup.Name)
	d.SetPartial("name")
	d.SetPartial("hosts")

	var connected_volumes []string
	if cv, ok := d.GetOk("connected_volumes"); ok {
		for _, element := range cv.([]interface{}) {
			connected_volumes = append(connected_volumes, element.(string))
		}
	}

	if connected_volumes != nil {
		for _, volume := range connected_volumes {
			_, err = client.Hostgroups.ConnectHostgroup(hgroup.Name, volume)
			if err != nil {
				return err
			}
		}
	}

	d.Partial(false)
	return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	h, _ := client.Hostgroups.GetHostgroup(d.Id(), nil)

	if h == nil {
		d.SetId("")
		return nil
	}

	var connected_volumes []string
	cv, _ := client.Hostgroups.ListHostgroupConnections(h.Name)
	for _, volume := range cv {
		connected_volumes = append(connected_volumes, volume.Vol)
	}

	d.Set("name", h.Name)
	d.Set("hosts", h.Hosts)
	d.Set("connected_volumes", connected_volumes)
	return nil
}

func resourcePureHostgroupUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*flasharray.Client)
	var h *flasharray.Hostgroup
	var err error

	if d.HasChange("name") {
		if h, err = client.Hostgroups.RenameHostgroup(d.Id(), d.Get("name").(string)); err != nil {
			return err
		}
		d.SetId(h.Name)
	}
	d.SetPartial("name")

	if d.HasChange("hosts") {
		var hosts []string
		for _, element := range d.Get("hosts").([]interface{}) {
			hosts = append(hosts, element.(string))
		}
		data := map[string][]string{"hostlist": hosts}
		if _, err = client.Hostgroups.SetHostgroup(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("hosts")

	if d.HasChange("connected_volumes") {
		var connected_volumes []string
		cv, _ := d.GetOk("connected_volumes")
		for _, element := range cv.([]interface{}) {
			connected_volumes = append(connected_volumes, element.(string))
		}
		var current_volumes []string
		curvols, _ := client.Hostgroups.ListHostgroupConnections(d.Id())
		for _, volume := range curvols {
			current_volumes = append(current_volumes, volume.Vol)
		}

		connect_volumes := difference(connected_volumes, current_volumes)
		for _, volume := range connect_volumes {
			if _, err = client.Hostgroups.ConnectHostgroup(d.Id(), volume); err != nil {
				return err
			}
		}

		disconnect_volumes := difference(current_volumes, connected_volumes)
		for _, volume := range disconnect_volumes {
			if _, err = client.Hostgroups.DisconnectHostgroup(d.Id(), volume); err != nil {
				return err
			}
		}
	}

	d.Partial(false)
	return resourcePureHostgroupRead(d, m)
}

func resourcePureHostgroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	var connected_volumes []string
	if cv, ok := d.GetOk("connected_volumes"); ok {
		for _, element := range cv.([]interface{}) {
			connected_volumes = append(connected_volumes, element.(string))
		}
	}
	if connected_volumes != nil {
		for _, volume := range connected_volumes {
			_, err := client.Hostgroups.DisconnectHostgroup(d.Id(), volume)
			if err != nil {
				return err
			}
		}
	}

	var hosts []string
	data := map[string][]string{"hostlist": hosts}
	_, err := client.Hostgroups.SetHostgroup(d.Id(), data)
	if err != nil {
		return err
	}

	_, err = client.Hostgroups.DeleteHostgroup(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourcePureHostgroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	h, err := client.Hostgroups.GetHostgroup(d.Id(), nil)

	if err != nil {
		return nil, err
	}

	var connected_volumes []string
	cv, _ := client.Hostgroups.ListHostgroupConnections(h.Name)
	for _, volume := range cv {
		connected_volumes = append(connected_volumes, volume.Vol)
	}

	d.Set("name", h.Name)
	d.Set("hosts", h.Hosts)
	d.Set("connected_volumes", connected_volumes)
	return []*schema.ResourceData{d}, nil
}
