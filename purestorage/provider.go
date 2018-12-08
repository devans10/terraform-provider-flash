package purestorage

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/mapstructure"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"api_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
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
		ResourcesMap: map[string]*schema.Resource{
			"purestorage_volume":          resourcePureVolume(),
			"purestorage_host":            resourcePureHost(),
			"purestorage_hostgroup":       resourcePureHostgroup(),
			"purestorage_protectiongroup": resourcePureProtectiongroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"purestorage_flasharray": dataSourcePureFlashArray(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing Pure client")
	return config.Client()
}
