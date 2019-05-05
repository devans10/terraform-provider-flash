/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package purestorage

import (
	"github.com/devans10/pugo/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
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
				Type:        schema.TypeString,
				Description: "Name of the host",
				Required:    true,
			},
			"iqn": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of iSCSI qualified names (IQNs) to the specified host.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
			},
			"wwn": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of Fibre Channel worldwide names (WWNs) to the specified host.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
			},
			"nqn": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of NVMeF qualified names (NQNs) to the specified host.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
			},
			"host_password": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Host password for CHAP authentication.",
				Computed:    true,
				Optional:    true,
			},
			"host_user": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Host username for CHAP authentication.",
				Optional:    true,
				Default:     "",
			},
			"personality": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "Determines how the Purity system tunes the protocol used between the array and the initiator.",
				Optional:     true,
				Default:      "",
				ValidateFunc: validation.StringInSlice([]string{"", "aix", "esxi", "hitachi-vsp", "hpux", "oracle-vm-server", "solaris", "vms"}, false),
			},
			"preferred_array": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of preferred arrays.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  nil,
			},
			"hgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_password": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Target password for CHAP authentication.",
				Computed:    true,
				Optional:    true,
			},
			"target_user": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Target username for CHAP authentication.",
				Optional:    true,
				Default:     "",
			},
			"volume": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"lun": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePureHostCreate(d *schema.ResourceData, m interface{}) error {

	d.Partial(true)
	client := m.(*flasharray.Client)
	var h *flasharray.Host
	var err error

	v, _ := d.GetOk("name")

	data := make(map[string]interface{})

	if wl, ok := d.GetOk("wwn"); ok {
		var wwnlist []string
		for _, element := range wl.([]interface{}) {
			wwnlist = append(wwnlist, element.(string))
		}
		data["wwnlist"] = wwnlist
	}

	if il, ok := d.GetOk("iqn"); ok {
		var iqnlist []string
		for _, element := range il.([]interface{}) {
			iqnlist = append(iqnlist, element.(string))
		}
		data["iqnlist"] = iqnlist
	}

	if nl, ok := d.GetOk("nqn"); ok {
		var nqnlist []string
		for _, element := range nl.([]interface{}) {
			nqnlist = append(nqnlist, element.(string))
		}
		data["nqnlist"] = nqnlist
	}

	if pa, ok := d.GetOk("preferred_array"); ok {
		var preferredArray []string
		for _, element := range pa.([]interface{}) {
			preferredArray = append(preferredArray, element.(string))
		}
		data["preferred_array"] = preferredArray
	}

	if len(data) > 0 {
		h, err = client.Hosts.CreateHost(v.(string), data)
		if err != nil {
			return err
		}
	} else {
		h, err = client.Hosts.CreateHost(v.(string), nil)
		if err != nil {
			return err
		}
	}
	d.SetId(h.Name)
	d.SetPartial("name")
	d.SetPartial("wwn")
	d.SetPartial("iqn")
	d.SetPartial("nqn")
	d.SetPartial("preferred_array")

	chapDetails := make(map[string]interface{})
	if hostPassword, ok := d.GetOk("host_password"); ok {
		chapDetails["host_password"] = hostPassword.(string)
	}

	if hostUser, ok := d.GetOk("host_user"); ok {
		chapDetails["host_user"] = hostUser.(string)
	}

	if targetPassword, ok := d.GetOk("target_password"); ok {
		chapDetails["target_password"] = targetPassword.(string)
	}

	if targetUser, ok := d.GetOk("target_user"); ok {
		chapDetails["target_user"] = targetUser.(string)
	}

	if len(chapDetails) > 0 {
		h, err = client.Hosts.SetHost(h.Name, chapDetails)
		if err != nil {
			return err
		}
	}
	d.SetPartial("host_password")
	d.SetPartial("host_user")
	d.SetPartial("target_password")
	d.SetPartial("target_user")

	if personality, ok := d.GetOk("personality"); ok {
		h, err = client.Hosts.SetHost(h.Name, map[string]string{"personality": personality.(string)})
		if err != nil {
			return err
		}
	}
	d.SetPartial("personality")

	if cv := d.Get("volume").(*schema.Set).List(); len(cv) > 0 {
		for _, volume := range cv {
			vol, _ := volume.(map[string]interface{})
			data := make(map[string]interface{})
			if vol["lun"] != 0 {
				data["lun"] = vol["lun"].(int)
			}
			if _, err := client.Hosts.ConnectHost(h.Name, vol["vol"].(string), data); err != nil {
				return err
			}
		}
	}

	d.Partial(false)

	return resourcePureHostRead(d, m)
}

func resourcePureHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	host, _ := client.Hosts.GetHost(d.Id(), nil)

	if host == nil {
		d.SetId("")
		return nil
	}

	if volumes, _ := client.Hosts.ListHostConnections(host.Name, map[string]string{"private": "true"}); volumes != nil {
		if err := d.Set("volume", flattenVolume(volumes)); err != nil {
			return err
		}
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	d.Set("nqn", host.Nqn)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"preferred_array": "true"})
	d.Set("preferred_array", host.PreferredArray)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"personality": "true"})
	d.Set("personality", host.Personality)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"chap": "true"})
	d.Set("host_password", host.HostPassword)
	d.Set("host_user", host.HostUser)
	d.Set("target_password", host.TargetPassword)
	d.Set("target_user", host.TargetUser)

	return nil
}

func resourcePureHostUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*flasharray.Client)
	var h *flasharray.Host
	var err error

	if d.HasChange("name") {
		if h, err = client.Hosts.RenameHost(d.Id(), d.Get("name").(string)); err != nil {
			return err
		}
		d.SetId(h.Name)
	}
	d.SetPartial("name")

	if d.HasChange("wwn") {
		var wwnlist []string
		wl, _ := d.GetOk("wwn")
		for _, element := range wl.([]interface{}) {
			wwnlist = append(wwnlist, element.(string))
		}
		data := map[string]interface{}{"wwnlist": wwnlist}
		if _, err = client.Hosts.SetHost(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("wwn")

	if d.HasChange("iqn") {
		var iqnlist []string
		il, _ := d.GetOk("iqn")
		for _, element := range il.([]interface{}) {
			iqnlist = append(iqnlist, element.(string))
		}
		data := map[string]interface{}{"iqnlist": iqnlist}
		if _, err = client.Hosts.SetHost(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("iqn")

	if d.HasChange("nqn") {
		var nqnlist []string
		nl, _ := d.GetOk("nqn")
		for _, element := range nl.([]interface{}) {
			nqnlist = append(nqnlist, element.(string))
		}
		data := map[string]interface{}{"nqnlist": nqnlist}
		if _, err = client.Hosts.SetHost(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("nqn")

	if d.HasChange("preferred_array") {
		var preferredArray []string
		pa, _ := d.GetOk("preferred_array")
		for _, element := range pa.([]interface{}) {
			preferredArray = append(preferredArray, element.(string))
		}
		data := map[string]interface{}{"preferred_array": preferredArray}
		if _, err = client.Hosts.SetHost(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("preferred_array")

	chapDetails := make(map[string]interface{})

	if d.HasChange("host_password") {
		chapDetails["host_password"] = d.Get("host_password").(string)
	}

	if d.HasChange("host_user") {
		chapDetails["host_user"] = d.Get("host_user").(string)
	}

	if d.HasChange("target_password") {
		chapDetails["target_password"] = d.Get("target_password").(string)
	}

	if d.HasChange("target_user") {
		chapDetails["target_user"] = d.Get("target_user").(string)
	}

	if len(chapDetails) > 0 {
		if _, err = client.Hosts.SetHost(d.Id(), chapDetails); err != nil {
			return err
		}
	}
	d.SetPartial("host_password")
	d.SetPartial("host_user")
	d.SetPartial("target_password")
	d.SetPartial("target_user")

	if d.HasChange("personality") {
		if _, err = client.Hosts.SetHost(d.Id(), map[string]string{"personality": d.Get("personality").(string)}); err != nil {
			return err
		}
	}
	d.SetPartial("personality")

	if d.HasChange("volume") {
		o, n := d.GetChange("volume")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)
		disconnectVolumes := os.Difference(ns).List()
		connectVolumes := ns.Difference(os).List()

		if len(connectVolumes) > 0 {
			for _, volume := range connectVolumes {
				data := make(map[string]interface{})
				vol := volume.(map[string]interface{})
				if vol["lun"] != 0 {
					data["lun"] = vol["lun"].(int)
				}
				if _, err = client.Hosts.ConnectHost(d.Id(), vol["vol"].(string), data); err != nil {
					return err
				}
			}
		}

		if len(disconnectVolumes) > 0 {
			for _, volume := range disconnectVolumes {
				vol := volume.(map[string]interface{})
				if _, err = client.Hosts.DisconnectHost(d.Id(), vol["vol"].(string)); err != nil {
					return err
				}
			}
		}
	}
	d.Partial(false)

	return resourcePureHostRead(d, m)
}

func resourcePureHostDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	volumes := d.Get("volume").(*schema.Set).List()
	for _, volume := range volumes {
		vol := volume.(map[string]interface{})
		if _, err := client.Hosts.DisconnectHost(d.Id(), vol["vol"].(string)); err != nil {
			return err
		}
	}

	if _, err := client.Hosts.DeleteHost(d.Id()); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourcePureHostImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	host, err := client.Hosts.GetHost(d.Id(), nil)

	if err != nil {
		return nil, err
	}

	if volumes, _ := client.Hosts.ListHostConnections(host.Name, map[string]string{"private": "true"}); volumes != nil {
		if err := d.Set("volume", flattenVolume(volumes)); err != nil {
			return nil, err
		}
	}

	d.Set("name", host.Name)
	d.Set("iqn", host.Iqn)
	d.Set("wwn", host.Wwn)
	d.Set("nqn", host.Nqn)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"preferred_array": "true"})
	d.Set("preferred_array", host.PreferredArray)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"personality": "true"})
	d.Set("personality", host.Personality)

	host, _ = client.Hosts.GetHost(d.Id(), map[string]string{"chap": "true"})
	d.Set("host_password", host.HostPassword)
	d.Set("host_user", host.HostUser)
	d.Set("target_password", host.TargetPassword)
	d.Set("target_user", host.TargetUser)

	return []*schema.ResourceData{d}, nil
}
