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

	if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), retention_data); err != nil {
		return err
	}
	d.SetPartial("all_for")
	d.SetPartial("days")
	d.SetPartial("per_day")
	d.SetPartial("target_all_for")
	d.SetPartial("target_days")
	d.SetPartial("target_per_day")

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

	if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), schedule_data); err != nil {
		return err
	}
	d.SetPartial("replicate_at")
	d.SetPartial("replicate_blackout")
	d.SetPartial("replicate_frequency")
	d.SetPartial("snap_at")
	d.SetPartial("snap_frequency")

	if replicate_enabled, ok := d.GetOk("replicate_enabled"); ok {
		if replicate_enabled.(bool) {
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

	if snap_enabled, ok := d.GetOk("snap_enabled"); ok {
		if snap_enabled.(bool) {
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

	retention_data := make(map[string]interface{})
	if d.HasChange("all_for") {
		retention_data["all_for"] = d.Get("all_for").(int)
	}

	if d.HasChange("days") {
		retention_data["days"] = d.Get("days").(int)
	}

	if d.HasChange("per_day") {
		retention_data["per_day"] = d.Get("per_day").(int)
	}

	if d.HasChange("target_all_for") {
		retention_data["target_all_for"] = d.Get("target_all_for").(int)
	}

	if d.HasChange("target_days") {
		retention_data["target_days"] = d.Get("target_days").(int)
	}

	if d.HasChange("target_per_day") {
		retention_data["target_per_day"] = d.Get("target_per_day").(int)
	}

	if len(retention_data) > 0 {
		if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), retention_data); err != nil {
			return err
		}
	}
	d.SetPartial("all_for")
	d.SetPartial("days")
	d.SetPartial("per_day")
	d.SetPartial("target_all_for")
	d.SetPartial("target_days")
	d.SetPartial("target_per_day")

	schedule_data := make(map[string]interface{})

	if d.HasChange("replicate_at") {
		schedule_data["replicate_at"] = d.Get("replicate_at").(int)
	}

	if d.HasChange("replicate_blackout") {
		schedule_data["replicate_blackout"] = d.Get("replicate_blackout").([]interface{})
	}

	if d.HasChange("replicate_frequency") {
		schedule_data["replicate_frequency"] = d.Get("replicate_frequency").(int)
	}

	if d.HasChange("snap_at") {
		schedule_data["snap_at"] = d.Get("snap_at").(int)
	}

	if d.HasChange("snap_frequency") {
		schedule_data["snap_frequency"] = d.Get("snap_frequency").(int)
	}

	if len(schedule_data) > 0 {
		if _, err = client.Protectiongroups.SetProtectiongroup(d.Id(), schedule_data); err != nil {
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
