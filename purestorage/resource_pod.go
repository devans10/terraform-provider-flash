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

func resourcePurePod() *schema.Resource {
	return &schema.Resource{
		Create: resourcePurePodCreate,
		Read:   resourcePurePodRead,
		Update: resourcePurePodUpdate,
		Delete: resourcePurePodDelete,
		Importer: &schema.ResourceImporter{
			State: resourcePurePodImport,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the pod",
				Required:    true,
			},
			"failover_preference": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of Failover arrays",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Source for the Pod",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

// resourcePurePodCreate creates a Pure Pod on a FlashArray
func resourcePurePodCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	var p *flasharray.Pod
	var err error

	n, _ := d.GetOk("name")
	if p, err = client.Pods.CreatePod(n.(string), nil); err != nil {
		return err
	}

	d.SetId(p.Name)
	return resourcePurePodRead(d, m)
}

// resourcePurePodRead sets the values for the given name
func resourcePurePodRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	pod, _ := client.Pods.GetPod(d.Id(), nil)

	if pod == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", pod.Name)
	d.Set("source", pod.Source)
	d.Set("failover_preference", pod.FailoverPreference)
	return nil
}

// resourcePurePodUpdate will update the attributes of the pod.
func resourcePurePodUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	client := m.(*flasharray.Client)
	var v *flasharray.Pod
	var err error

	if d.HasChange("name") {
		if v, err = client.Pods.RenamePod(d.Id(), d.Get("name").(string)); err != nil {
			return err
		}
		d.SetId(v.Name)
	}
	d.SetPartial("name")

	// TODO: Rest of the Update Function

	d.Partial(false)

	return resourcePureVolumeRead(d, m)
}

// resourcePurePodDelete will delete the Pod specified.
func resourcePurePodDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)
	_, err := client.Pods.DeletePod(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// resourcePurePodImport imports a pod into Terraform.
func resourcePurePodImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*flasharray.Client)

	pod, err := client.Pods.GetPod(d.Id(), nil)

	if err != nil {
		return nil, err
	}

	d.Set("name", pod.Name)
	d.Set("source", pod.Source)
	d.Set("failover_preference", pod.FailoverPreference)
	return []*schema.ResourceData{d}, nil
}
