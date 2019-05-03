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
)

func resourcePureProtectiongroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePureProtectiongroupCreate,
		Read:   resourcePureProtectiongroupRead,
		Update: resourcePureProtectiongroupUpdate,
		Delete: resourcePureProtectiongroupDelete,
		Importer: &schema.ResourceImporter{
			State: resourcePureProtectiongroupImport,
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
					Type: schema.TypeMap,
				},
				Optional: true,
				Default:  nil,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"all_for": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the length of time to keep the snapshots on the source array before they are eradicated.",
				Optional:    true,
				Default:     86400,
			},
			"days": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of days to keep the per_day snapshots beyond the all_for period before they are eradicated.",
				Optional:    true,
				Default:     7,
			},
			"per_day": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of per_day snapshots to keep beyond the all_for period.",
				Optional:    true,
				Default:     4,
			},
			"replicate_at": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Modifies the replication schedule of the protection group. Specifies the preferred time, on the hour, at which to replicate the snapshots.",
				Default:     nil,
			},
			"replicate_blackout": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Modifies the replication schedule of the protection group. Specifies the range of time at which to suspend replication.",
				Optional:    true,
				Default:     nil,
			},
			"replicate_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Used to enable (true) or disable (false) the protection group replication schedule.",
				Optional:    true,
				Default:     false,
			},
			"replicate_frequency": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the replication schedule of the protection group. Specifies the replication frequency.",
				Optional:    true,
				Default:     14400,
			},
			"snap_at": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the snapshot schedule of the protection group. Specifies the preferred time, on the hour, at which to generate the snapshot.",
				Optional:    true,
				Default:     nil,
			},
			"snap_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Used to enable (true) or disable (false) the protection group snapshot schedule.",
				Optional:    true,
				Default:     false,
			},
			"snap_frequency": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the snapshot schedule of the protection group. Specifies the snapshot frequency.",
				Optional:    true,
				Default:     3600,
			},
			"target_all_for": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the length of time to keep the replicated snapshots on the targets.",
				Optional:    true,
				Default:     86400,
			},
			"target_days": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of days to keep the target_per_day replicated snapshots beyond the target_all_for period before they are eradicated.",
				Optional:    true,
				Default:     7,
			},
			"target_per_day": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of per_day replicated snapshots to keep beyond the target_all_for period.",
				Optional:    true,
				Default:     4,
			},
		},
	}
}

func resourcePureProtectiongroupCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	client := m.(*flasharray.Client)
	var pgroup *flasharray.Protectiongroup
	var err error

	data := make(map[string]interface{})

	if h, ok := d.GetOk("hosts"); ok {
		var hosts []string
		for _, element := range h.([]interface{}) {
			hosts = append(hosts, element.(string))
		}
		data["hostlist"] = hosts
	}

	if v, ok := d.GetOk("volumes"); ok {
		var volumes []string
		for _, element := range v.([]interface{}) {
			volumes = append(volumes, element.(string))
		}
		data["vollist"] = volumes
	}

	if hg, ok := d.GetOk("hgroups"); ok {
		var hgroups []string
		for _, element := range hg.([]interface{}) {
			hgroups = append(hgroups, element.(string))
		}
		data["hgrouplist"] = hgroups
	}

	if t, ok := d.GetOk("targets"); ok {
		var targets []string
		for _, element := range t.([]interface{}) {
			targets = append(targets, element.(string))
		}
		data["targetlist"] = targets
	}

	if pgroup, err = client.Protectiongroups.CreateProtectiongroup(d.Get("name").(string), data); err != nil {
		return err
	}
	d.SetId(pgroup.Name)
	d.SetPartial("name")
	d.SetPartial("hosts")
	d.SetPartial("volumes")
	d.SetPartial("hgroups")
	d.SetPartial("targets")

	retentionData := make(map[string]interface{})
	if allFor, ok := d.GetOk("all_for"); ok {
		retentionData["all_for"] = allFor
	}

	if days, ok := d.GetOk("days"); ok {
		retentionData["days"] = days
	}

	if perDay, ok := d.GetOk("per_day"); ok {
		retentionData["per_day"] = perDay
	}

	if targetAllFor, ok := d.GetOk("target_all_for"); ok {
		retentionData["target_all_for"] = targetAllFor
	}

	if targetDays, ok := d.GetOk("target_days"); ok {
		retentionData["target_days"] = targetDays
	}

	if targetPerDay, ok := d.GetOk("target_per_day"); ok {
		retentionData["target_per_day"] = targetPerDay
	}

	if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), retentionData); err != nil {
		return err
	}
	d.SetPartial("all_for")
	d.SetPartial("days")
	d.SetPartial("per_day")
	d.SetPartial("target_all_for")
	d.SetPartial("target_days")
	d.SetPartial("target_per_day")

	scheduleData := make(map[string]interface{})

	if replicateAt, ok := d.GetOk("replicate_at"); ok {
		scheduleData["replicate_at"] = replicateAt
	}

	if replicateBlackout, ok := d.GetOk("replicate_blackout"); ok {
		scheduleData["replicate_blackout"] = replicateBlackout
	}

	if replicateFrequency, ok := d.GetOk("replicate_frequency"); ok {
		scheduleData["replicate_frequency"] = replicateFrequency
	}

	if snapAt, ok := d.GetOk("snap_at"); ok {
		scheduleData["snap_at"] = snapAt
	}

	if snapFrequency, ok := d.GetOk("snap_frequency"); ok {
		scheduleData["snap_frequency"] = snapFrequency
	}

	if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), scheduleData); err != nil {
		return err
	}
	d.SetPartial("replicate_at")
	d.SetPartial("replicate_blackout")
	d.SetPartial("replicate_frequency")
	d.SetPartial("snap_at")
	d.SetPartial("snap_frequency")

	if replicateEnabled, ok := d.GetOk("replicate_enabled"); ok {
		if replicateEnabled.(bool) {
			if _, err = client.Protectiongroups.EnablePgroupReplication(d.Id()); err != nil {
				return err
			}
		} else {
			if _, err = client.Protectiongroups.DisablePgroupReplication(d.Id()); err != nil {
				return err
			}
		}
	}
	d.SetPartial("replicate_enabled")

	if snapEnabled, ok := d.GetOk("snap_enabled"); ok {
		if snapEnabled.(bool) {
			if _, err = client.Protectiongroups.EnablePgroupSnapshots(d.Id()); err != nil {
				return err
			}
		} else {
			if _, err = client.Protectiongroups.DisablePgroupSnapshots(d.Id()); err != nil {
				return err
			}
		}
	}
	d.Partial(false)

	return resourcePureProtectiongroupRead(d, m)
}

func resourcePureProtectiongroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	var p *flasharray.Protectiongroup

	if p, _ = client.Protectiongroups.GetProtectiongroup(d.Id(), nil); p == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", p.Name)
	d.Set("hosts", p.Hosts)
	d.Set("volumes", p.Volumes)
	d.Set("hgroups", p.Hgroups)
	d.Set("source", p.Source)
	d.Set("targets", p.Targets)

	params := map[string]string{"schedule": "true"}
	s, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), params)
	if s != nil {
		d.Set("replicate_at", s.ReplicateAt)
		d.Set("replicate_blackout", s.ReplicateBlackout)
		d.Set("replicate_frequency", s.ReplicateFrequency)
		d.Set("replicate_enabled", s.ReplicateEnabled)
		d.Set("snap_at", s.SnapAt)
		d.Set("snap_enabled", s.SnapEnabled)
		d.Set("snap_frequency", s.SnapFrequency)
	}

	params = map[string]string{"retention": "true"}
	r, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), params)
	if r != nil {
		d.Set("all_for", r.Allfor)
		d.Set("days", r.Days)
		d.Set("per_day", r.Perday)
		d.Set("target_all_for", r.TargetAllfor)
		d.Set("target_days", r.TargetDays)
		d.Set("target_per_day", r.TargetPerDay)
	}
	return nil
}

func resourcePureProtectiongroupUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	var pgroup *flasharray.Protectiongroup
	var err error
	client := m.(*flasharray.Client)

	if d.HasChange("name") {
		if pgroup, err = client.Protectiongroups.RenameProtectiongroup(pgroup.Name, d.Get("name").(string)); err != nil {
			return err
		}
		d.SetId(pgroup.Name)
	}
	d.SetPartial("name")

	data := make(map[string]interface{})
	if d.HasChange("hosts") {
		var hosts []string
		for _, element := range d.Get("hosts").([]interface{}) {
			hosts = append(hosts, element.(string))
		}
		data["hostlist"] = hosts
	}

	if d.HasChange("volumes") {
		var volumes []string
		for _, element := range d.Get("volumes").([]interface{}) {
			volumes = append(volumes, element.(string))
		}
		data["vollist"] = volumes
	}

	if d.HasChange("hgroups") {
		var hgroups []string
		for _, element := range d.Get("hgroups").([]interface{}) {
			hgroups = append(hgroups, element.(string))
		}
		data["hgrouplist"] = hgroups
	}

	if d.HasChange("targets") {
		var targets []string
		for _, element := range d.Get("targets").([]interface{}) {
			targets = append(targets, element.(string))
		}
		data["targetlist"] = targets
	}

	if len(data) > 0 {
		if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), data); err != nil {
			return err
		}
	}
	d.SetPartial("hosts")
	d.SetPartial("volumes")
	d.SetPartial("hgroups")
	d.SetPartial("targets")

	retentionData := make(map[string]interface{})
	if d.HasChange("all_for") {
		retentionData["all_for"] = d.Get("all_for").(int)
	}

	if d.HasChange("days") {
		retentionData["days"] = d.Get("days").(int)
	}

	if d.HasChange("per_day") {
		retentionData["per_day"] = d.Get("per_day").(int)
	}

	if d.HasChange("target_all_for") {
		retentionData["target_all_for"] = d.Get("target_all_for").(int)
	}

	if d.HasChange("target_days") {
		retentionData["target_days"] = d.Get("target_days").(int)
	}

	if d.HasChange("target_per_day") {
		retentionData["target_per_day"] = d.Get("target_per_day").(int)
	}

	if len(retentionData) > 0 {
		if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), retentionData); err != nil {
			return err
		}
	}
	d.SetPartial("all_for")
	d.SetPartial("days")
	d.SetPartial("per_day")
	d.SetPartial("target_all_for")
	d.SetPartial("target_days")
	d.SetPartial("target_per_day")

	scheduleData := make(map[string]interface{})

	if d.HasChange("replicate_at") {
		scheduleData["replicate_at"] = d.Get("replicate_at").(int)
	}

	if d.HasChange("replicate_blackout") {
		scheduleData["replicate_blackout"] = d.Get("replicate_blackout").([]interface{})
	}

	if d.HasChange("replicate_frequency") {
		scheduleData["replicate_frequency"] = d.Get("replicate_frequency").(int)
	}

	if d.HasChange("snap_at") {
		scheduleData["snap_at"] = d.Get("snap_at").(int)
	}

	if d.HasChange("snap_frequency") {
		scheduleData["snap_frequency"] = d.Get("snap_frequency").(int)
	}

	if len(scheduleData) > 0 {
		if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), scheduleData); err != nil {
			return err
		}
	}
	d.SetPartial("replicate_at")
	d.SetPartial("replicate_blackout")
	d.SetPartial("replicate_frequency")
	d.SetPartial("snap_at")
	d.SetPartial("snap_frequency")

	if d.HasChange("replicate_enabled") {
		if d.Get("replicate_enabled").(bool) {
			if _, err = client.Protectiongroups.EnablePgroupReplication(d.Id()); err != nil {
				return err
			}
		} else {
			if _, err = client.Protectiongroups.DisablePgroupReplication(d.Id()); err != nil {
				return err
			}
		}
	}
	d.SetPartial("replicate_enabled")

	if d.HasChange("snap_enabled") {
		if d.Get("snap_enabled").(bool) {
			if _, err = client.Protectiongroups.EnablePgroupSnapshots(d.Id()); err != nil {
				return err
			}
		} else {
			if _, err = client.Protectiongroups.DisablePgroupSnapshots(d.Id()); err != nil {
				return err
			}
		}
	}

	d.Partial(false)
	return resourcePureProtectiongroupRead(d, m)
}

func resourcePureProtectiongroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	_, err := client.Protectiongroups.DestroyProtectiongroup(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourcePureProtectiongroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	p, err := client.Protectiongroups.GetProtectiongroup(d.Id(), nil)

	if err != nil {
		return nil, err
	}

	d.Set("name", p.Name)
	d.Set("hosts", p.Hosts)
	d.Set("volumes", p.Volumes)
	d.Set("hgroups", p.Hgroups)
	d.Set("source", p.Source)
	d.Set("targets", p.Targets)

	params := map[string]string{"schedule": "true"}
	s, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), params)
	if s != nil {
		d.Set("replicate_at", s.ReplicateAt)
		d.Set("replicate_blackout", s.ReplicateBlackout)
		d.Set("replicate_frequency", s.ReplicateFrequency)
		d.Set("replicate_enabled", s.ReplicateEnabled)
		d.Set("snap_at", s.SnapAt)
		d.Set("snap_enabled", s.SnapEnabled)
		d.Set("snap_frequency", s.SnapFrequency)
	}

	params = map[string]string{"retention": "true"}
	r, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), params)
	if r != nil {
		d.Set("all_for", r.Allfor)
		d.Set("days", r.Days)
		d.Set("per_day", r.Perday)
		d.Set("target_all_for", r.TargetAllfor)
		d.Set("target_days", r.TargetDays)
		d.Set("target_per_day", r.TargetPerDay)
	}
	return []*schema.ResourceData{d}, nil
}
