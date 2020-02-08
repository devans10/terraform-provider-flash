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
	"fmt"
	"log"

	"github.com/devans10/pugo/flasharray"
	"github.com/hashicorp/terraform/helper/schema"
)

// Config is the configuration for the Purestorage FlashArray Go Client.
// It holds the connection information to connect to the array API.
// Either the Username and Password or the API Token should be provided,
// but not both.
type Config struct {
	Username      string
	Password      string
	Target        string
	APIToken      string
	RestVersion   string
	VerifyHTTPS   bool
	SslCert       bool
	RequestKwargs map[string]string
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

	requestKwargs := make(map[string]string)

	for key, value := range d.Get("request_kwargs").(map[string]interface{}) {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		requestKwargs[strKey] = strValue
	}

	c := &Config{
		Username:      username,
		Password:      password,
		Target:        d.Get("target").(string),
		APIToken:      apitoken,
		RestVersion:   d.Get("rest_version").(string),
		VerifyHTTPS:   d.Get("verify_https").(bool),
		SslCert:       d.Get("ssl_cert").(bool),
		RequestKwargs: requestKwargs,
	}

	return c, nil
}

// Client returns a new client for accessing flasharray.
//
func (c *Config) Client() (*flasharray.Client, error) {

	client, err := flasharray.NewClient(c.Target, c.Username, c.Password, c.APIToken, c.RestVersion, c.VerifyHTTPS, c.SslCert, "Terraform", c.RequestKwargs)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Pure Client configured for target: %s", c.Target)

	return client, err
}
