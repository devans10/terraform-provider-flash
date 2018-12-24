package purestorage

import (
	"fmt"
	"log"

	"github.com/devans10/go-purestorage/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
)

// Config is the configuration for the Purestorage FlashArray Go Client.
// It holds the connection information to connect to the array API.
// Either the Username and Password or the API Token should be provided,
// but not both.
type Config struct {
	Username       string
	Password       string
	Target         string
	ApiToken       string
	Rest_version   string
	Verify_https   bool
	Ssl_cert       bool
	User_agent     string
	Request_kwargs map[string]string
}

// NewConfig returns a new Config from a supplied ResourceData.
func NewConfig(d *schema.ResourceData) (*Config, error) {

	// Handle the fact that (username and password) or api_token are
	// mutually exclusive, but one of the sets is required.
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	apitoken := d.Get("api_token").(string)

	if (username != "") && (password != "") && (apitoken != "") {
		return nil, fmt.Errorf("Username and Password or API Token must be provided, but not both")
	}

	if (username != "") && (password == "") {
		return nil, fmt.Errorf("Password must be provided with Username")
	}

	request_kwargs := make(map[string]string)

	for key, value := range d.Get("request_kwargs").(map[string]interface{}) {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		request_kwargs[strKey] = strValue
	}

	c := &Config{
		Username:       username,
		Password:       password,
		Target:         d.Get("target").(string),
		ApiToken:       apitoken,
		Rest_version:   d.Get("rest_version").(string),
		Verify_https:   d.Get("verify_https").(bool),
		Ssl_cert:       d.Get("ssl_cert").(bool),
		User_agent:     d.Get("user_agent").(string),
		Request_kwargs: request_kwargs,
	}

	return c, nil
}

// Client() returns a new client for accessing flasharray.
//
func (c *Config) Client() (*flasharray.Client, error) {

	client, err := flasharray.NewClient(c.Target, c.Username, c.Password, c.ApiToken, c.Rest_version, c.Verify_https, c.Ssl_cert, c.User_agent, c.Request_kwargs)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Pure Client configured for target: %s", c.Target)

	return client, err
}
