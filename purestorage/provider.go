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
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider is the terraform resource provider called by main.go
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PURE_USERNAME", ""),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PURE_PASSWORD", ""),
			},

			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PURE_APITOKEN", ""),
			},

			"target": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PURE_TARGET", ""),
			},

			"rest_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"verify_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"ssl_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"request_kwargs": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Default:  nil,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"purestorage_flasharray": dataSourcePureFlashArray(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"purestorage_flasharray": schema.DataSourceResourceShim(
				"purestorage_flasharray",
				dataSourcePureFlashArray(),
			),
			"purestorage_volume":          resourcePureVolume(),
			"purestorage_host":            resourcePureHost(),
			"purestorage_hostgroup":       resourcePureHostgroup(),
			"purestorage_protectiongroup": resourcePureProtectiongroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c, err := NewConfig(d)
	if err != nil {
		return nil, err
	}

	return c.Client()
}
