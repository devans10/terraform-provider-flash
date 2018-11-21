package purestorage

import (
	"log"
	"os"

	"github.com/devans10/go-pure-client/pureClient"
)

type Config struct {
	User		string	`mapstructure:"user"`
	Password	string	`mapstructure:"password"`
	Entrypoint	string	`mapstructure:"entrypoint"`
}

// Client() returns a new client for accessing pingdom.
//
func (c *Config) Client() (*pureClient.Client, error) {

	if v := os.Getenv("PURE_USERNAME"); v != "" {
		c.User = v
	}
	if v := os.Getenv("PURE_PASSWORD"); v != "" {
		c.Password = v
	}
	if v := os.Getenv("PURE_ENTRYPOINT"); v != "" {
		c.Entrypoint = v
	}

	client := pureClinet.NewClient(c.User, c.Password, c.Entrypoint)

	log.Printf("[INFO] Pure Client configured for user: %s", c.User)

	return client, nil
}
