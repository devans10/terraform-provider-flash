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
			"all_for": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the length of time to keep the snapshots on the source array before they are eradicated.",
				Optional:    true,
				Default:     nil,
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
				Default:     nil,
			},
			"per_day": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of per_day snapshots to keep beyond the all_for period.",
				Optional:    true,
				Default:     nil,
			},
			"replicate_at": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Modifies the replication schedule of the protection group. Specifies the preferred time, on the hour, at which to replicate the snapshots.",
				Default:     nil,
			},
			"replicate_blackout": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "Modifies the replication schedule of the protection group. Specifies the range of time at which to suspend replication.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  nil,
						},
						"start": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  nil,
						},
					},
				},
				Optional: true,
				Default:  nil,
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
				Default:     nil,
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
				Default:     nil,
			},
			"target_all_for": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the length of time to keep the replicated snapshots on the targets.",
				Optional:    true,
				Default:     nil,
			},
			"target_days": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of days to keep the target_per_day replicated snapshots beyond the target_all_for period before they are eradicated.",
				Optional:    true,
				Default:     nil,
			},
			"target_per_day": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Modifies the retention policy of the protection group. Specifies the number of per_day replicated snapshots to keep beyond the target_all_for period.",
				Optional:    true,
				Default:     nil,
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
