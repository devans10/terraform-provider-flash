package purestorage

import (
	"reflect"

	"github.com/devans10/go-purestorage/flasharray"
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
					Type: schema.TypeString,
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
			"allowed": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Allows (true) or disallows (false) a protection group from being replicated.",
				Optional:    true,
				Default:     false,
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
	client := m.(*flasharray.Client)
	data := make(map[string]interface{})
	p, _ := d.GetOk("name")

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

	pgroup, err := client.Protectiongroups.CreateProtectiongroup(p.(string), data, nil)
	if err != nil {
		return err
	}

	retention_data := make(map[string]interface{})
	if all_for, ok := d.GetOk("all_for"); ok {
		retention_data["all_for"] = all_for
	}

	if days, ok := d.GetOk("days"); ok {
		retention_data["days"] = days
	}

	if per_day, ok := d.GetOk("per_day"); ok {
		retention_data["per_day"] = per_day
	}

	if target_all_for, ok := d.GetOk("target_all_for"); ok {
		retention_data["target_all_for"] = target_all_for
	}

	if target_days, ok := d.GetOk("target_days"); ok {
		retention_data["target_days"] = target_days
	}

	if target_per_day, ok := d.GetOk("target_per_day"); ok {
		retention_data["target_per_day"] = target_per_day
	}

	pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, retention_data)
	if err != nil {
		return err
	}

	schedule_data := make(map[string]interface{})

	if replicate_at, ok := d.GetOk("replicate_at"); ok {
		schedule_data["replicate_at"] = replicate_at
	}

	if replicate_blackout, ok := d.GetOk("replicate_blackout"); ok {
		schedule_data["replicate_blackout"] = replicate_blackout
	}

	if replicate_frequency, ok := d.GetOk("replicate_frequency"); ok {
		schedule_data["replicate_frequency"] = replicate_frequency
	}

	if snap_at, ok := d.GetOk("snap_at"); ok {
		schedule_data["snap_at"] = snap_at
	}

	if snap_frequency, ok := d.GetOk("snap_frequency"); ok {
		schedule_data["snap_frequency"] = snap_frequency
	}

	pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, schedule_data)
	if err != nil {
		return err
	}

	allowed_data := make(map[string]interface{})
	if allowed, ok := d.GetOk("allowed"); ok {
		allowed_data["allowed"] = allowed

		pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, allowed_data)
		if err != nil {
			return err
		}
	}

	if replicate_enabled, ok := d.GetOk("replicate_enabled"); ok {
		if replicate_enabled.(bool) {
			_, err = client.Protectiongroups.EnablePgroupReplication(p.(string), nil)
			if err != nil {
				return err
			}
		} else {
			_, err = client.Protectiongroups.DisablePgroupReplication(p.(string), nil)
			if err != nil {
				return err
			}
		}
	}

	if snap_enabled, ok := d.GetOk("snap_enabled"); ok {
		if snap_enabled.(bool) {
			_, err = client.Protectiongroups.EnablePgroupSnapshots(p.(string), nil)
			if err != nil {
				return err
			}
		} else {
			_, err = client.Protectiongroups.DisablePgroupSnapshots(p.(string), nil)
			if err != nil {
				return err
			}
		}
	}

	d.SetId(pgroup.Name)
	return resourcePureProtectiongroupRead(d, m)
}

func resourcePureProtectiongroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	p, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, nil)

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

	data := map[string]bool{"schedule": true}
	s, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, data)
	if s != nil {
		d.Set("replicate_at", s.ReplicateAt)
		d.Set("replicate_blackout", s.ReplicateBlackout)
		d.Set("replicate_frequency", s.ReplicateFrequency)
		d.Set("replicate_enabled", s.ReplicateEnabled)
		d.Set("snap_at", s.SnapAt)
		d.Set("snap_enabled", s.SnapEnabled)
		d.Set("snap_frequency", s.SnapFrequency)
	}

	data = map[string]bool{"retention": true}
	r, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, data)
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
	client := m.(*flasharray.Client)

	pgroup, err := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, nil)
	if err != nil {
		return err
	}

	p, _ := d.GetOk("name")
	if pgroup.Name != p.(string) {
		pgroup, err = client.Protectiongroups.RenameProtectiongroup(pgroup.Name, p.(string), nil)
		if err != nil {
			return err
		}
		d.SetId(pgroup.Name)
	}

	pgroup, err = client.Protectiongroups.GetProtectiongroup(d.Id(), nil, nil)

	data := make(map[string]interface{})
	if h, ok := d.GetOk("hosts"); ok {
		var hosts []string
		for _, element := range h.([]interface{}) {
			hosts = append(hosts, element.(string))
		}
		if !sameStringSlice(pgroup.Hosts, hosts) {
			data["hostlist"] = hosts
		}
	}

	if v, ok := d.GetOk("volumes"); ok {
		var volumes []string
		for _, element := range v.([]interface{}) {
			volumes = append(volumes, element.(string))
		}
		if !sameStringSlice(pgroup.Volumes, volumes) {
			data["vollist"] = volumes
		}
	}

	if hg, ok := d.GetOk("hgroups"); ok {
		var hgroups []string
		for _, element := range hg.([]interface{}) {
			hgroups = append(hgroups, element.(string))
		}
		if !sameStringSlice(pgroup.Hgroups, hgroups) {
			data["hgrouplist"] = hgroups
		}
	}

	if t, ok := d.GetOk("targets"); ok {
		var targets []string
		for _, element := range t.([]interface{}) {
			targets = append(targets, element.(string))
		}
		if !sameStringSlice(pgroup.Targets, targets) {
			data["targetlist"] = targets
		}
	}

	if data != nil {
		pgroup, err = client.Protectiongroups.SetProtectiongroup(d.Id(), nil, data)
		if err != nil {
			return err
		}
	}

	pgroup, err = client.Protectiongroups.GetProtectiongroup(d.Id(), nil, map[string]bool{"retention": true})
	if err != nil {
		return err
	}

	retention_data := make(map[string]interface{})
	if all_for, ok := d.GetOk("all_for"); ok {
		if pgroup.Allfor != all_for {
			retention_data["all_for"] = all_for
		}
	}

	if days, ok := d.GetOk("days"); ok {
		if pgroup.Days != days {
			retention_data["days"] = days
		}
	}

	if per_day, ok := d.GetOk("per_day"); ok {
		if pgroup.Perday != per_day {
			retention_data["per_day"] = per_day
		}
	}

	if target_all_for, ok := d.GetOk("target_all_for"); ok {
		if pgroup.TargetAllfor != target_all_for {
			retention_data["target_all_for"] = target_all_for
		}
	}

	if target_days, ok := d.GetOk("target_days"); ok {
		if pgroup.TargetDays != target_days {
			retention_data["target_days"] = target_days
		}
	}

	if target_per_day, ok := d.GetOk("target_per_day"); ok {
		if pgroup.TargetPerDay != target_per_day {
			retention_data["target_per_day"] = target_per_day
		}
	}

	if retention_data != nil {
		pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, retention_data)
		if err != nil {
			return err
		}
	}

	pgroup, err = client.Protectiongroups.GetProtectiongroup(d.Id(), nil, map[string]bool{"schedule": true})
	if err != nil {
		return err
	}

	schedule_data := make(map[string]interface{})

	if replicate_at, ok := d.GetOk("replicate_at"); ok {
		if pgroup.ReplicateAt != replicate_at {
			schedule_data["replicate_at"] = replicate_at
		}
	}

	if replicate_blackout, ok := d.GetOk("replicate_blackout"); ok {
		if !reflect.DeepEqual(pgroup.ReplicateBlackout, replicate_blackout) {
			schedule_data["replicate_blackout"] = replicate_blackout
		}
	}

	if replicate_frequency, ok := d.GetOk("replicate_frequency"); ok {
		if pgroup.ReplicateFrequency != replicate_frequency {
			schedule_data["replicate_frequency"] = replicate_frequency
		}
	}

	if snap_at, ok := d.GetOk("snap_at"); ok {
		if pgroup.SnapAt != snap_at {
			schedule_data["snap_at"] = snap_at
		}
	}

	if snap_frequency, ok := d.GetOk("snap_frequency"); ok {
		if pgroup.SnapFrequency != snap_frequency {
			schedule_data["snap_frequency"] = snap_frequency
		}
	}

	if schedule_data != nil {
		pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, schedule_data)
		if err != nil {
			return err
		}
	}
	/*
		allowed_data := make(map[string]interface{})
		if allowed, ok := d.GetOk("allowed"); ok {
			allowed_data["allowed"] = allowed

			pgroup, err = client.Protectiongroups.SetProtectiongroup(p.(string), nil, allowed_data)
			if err != nil {
				return err
			}
		}
	*/
	if replicate_enabled, ok := d.GetOk("replicate_enabled"); ok {
		if replicate_enabled.(bool) {
			_, err = client.Protectiongroups.EnablePgroupReplication(p.(string), nil)
			if err != nil {
				return err
			}
		} else {
			_, err = client.Protectiongroups.DisablePgroupReplication(p.(string), nil)
			if err != nil {
				return err
			}
		}
	}

	if snap_enabled, ok := d.GetOk("snap_enabled"); ok {
		if snap_enabled.(bool) {
			_, err = client.Protectiongroups.EnablePgroupSnapshots(p.(string), nil)
			if err != nil {
				return err
			}
		} else {
			_, err = client.Protectiongroups.DisablePgroupSnapshots(p.(string), nil)
			if err != nil {
				return err
			}
		}
	}

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

func resourcePureProtectiongroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	p, err := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, nil)

	if err != nil {
		return nil, err
	}

	d.Set("name", p.Name)
	d.Set("hosts", p.Hosts)
	d.Set("volumes", p.Volumes)
	d.Set("hgroups", p.Hgroups)
	d.Set("source", p.Source)
	d.Set("targets", p.Targets)

	data := map[string]bool{"schedule": true}
	s, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, data)
	if s != nil {
		d.Set("replicate_at", s.ReplicateAt)
		d.Set("replicate_blackout", s.ReplicateBlackout)
		d.Set("replicate_frequency", s.ReplicateFrequency)
		d.Set("replicate_enabled", s.ReplicateEnabled)
		d.Set("snap_at", s.SnapAt)
		d.Set("snap_enabled", s.SnapEnabled)
		d.Set("snap_frequency", s.SnapFrequency)
	}

	data = map[string]bool{"retention": true}
	r, _ := client.Protectiongroups.GetProtectiongroup(d.Id(), nil, data)
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
