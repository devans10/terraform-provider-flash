package purestorage

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

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

			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
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
