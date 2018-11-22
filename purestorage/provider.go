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
			"user": &schema.Schema{
				Type:		schema.TypeString,
				Optional:	true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"entrypoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
                ResourcesMap: map[string]*schema.Resource{
			"pure_volume": resourcePureVolume(),
			"pure_host": resourcePureHost(),
			"pure_hostgroup": resourcePureHostgroup(),
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
