package purestorage

import (
	"log"
	"os"

	"github.com/devans10/go-purestorage/flasharray"
)

type Config struct {
	Username       string            `mapstructure:"username"`
	Password       string            `mapstructure:"password"`
	Target         string            `mapstructure:"target"`
	ApiToken       string            `mapstructure:"api_token"`
	Rest_version   string            `mapstructure:"rest_version"`
	Verify_https   bool              `mapstructure:"verify_https"`
	Ssl_cert       bool              `mapstructure:"ssl_cert"`
	User_agent     string            `mapsturcture:"user_agent"`
	Request_kwargs map[string]string `mapstructure:"request_kwargs"`
}

// Client() returns a new client for accessing flasharray.
//
func (c *Config) Client() (*flasharray.Client, error) {

	if v := os.Getenv("PURE_USERNAME"); v != "" {
		c.Username = v
	}
	if v := os.Getenv("PURE_PASSWORD"); v != "" {
		c.Password = v
	}
	if v := os.Getenv("PURE_TARGET"); v != "" {
		c.Target = v
	}
	if v := os.Getenv("PURE_APITOKEN"); v != "" {
		c.ApiToken = v
	}

	client, err := flasharray.NewClient(c.Target, c.Username, c.Password, c.ApiToken, c.Rest_version, c.Verify_https, c.Ssl_cert, c.User_agent, c.Request_kwargs)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Pure Client configured for target: %s", c.Target)

	return client, err
}
