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

func dataSourcePureFlashArray() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePureFlashArrayRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePureFlashArrayRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*flasharray.Client)

	flasharray, err := client.Array.Get(nil)
	if err != nil {
		return err
	}

	d.SetId(flasharray.Id)
	d.Set("name", flasharray.Array_name)
	d.Set("version", flasharray.Version)
	d.Set("revision", flasharray.Revision)
	return nil
}
