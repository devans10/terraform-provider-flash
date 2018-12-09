package purestorage

import (
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func testAccClientPreCheck(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC to run purestorage acceptance tests (provider connection is required)")
	}
	testAccPreCheck(t)
}

func testAccClientGenerateConfig(t *testing.T) *Config {

	return &Config{
		Username:       os.Getenv("PURE_USERNAME"),
		Password:       os.Getenv("PURE_PASSWORD"),
		Target:         os.Getenv("PURE_TARGET"),
		ApiToken:       os.Getenv("PURE_APITOKEN"),
		Rest_version:   "",
		Verify_https:   false,
		Ssl_cert:       false,
		User_agent:     "",
		Request_kwargs: nil,
	}
}

func TestNewConfigWithApiToken(t *testing.T) {
	expected := &Config{
		Username:       "",
		Password:       "",
		Target:         "purestorage.flasharray",
		ApiToken:       "foobar",
		Rest_version:   "1.15",
		Verify_https:   false,
		Ssl_cert:       false,
		User_agent:     "useragent",
		Request_kwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.ApiToken)
	d.Set("rest_version", expected.Rest_version)
	d.Set("verify_https", expected.Verify_https)
	d.Set("ssl_cert", expected.Ssl_cert)
	d.Set("user_agent", expected.User_agent)

	actual, err := NewConfig(d)
	if err != nil {
		t.Fatalf("error creating new configuration: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func TestNewConfigWithUsernameAndPassword(t *testing.T) {
	expected := &Config{
		Username:       "foo",
		Password:       "bar",
		Target:         "purestorage.flasharray",
		ApiToken:       "",
		Rest_version:   "1.15",
		Verify_https:   false,
		Ssl_cert:       false,
		User_agent:     "useragent",
		Request_kwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.ApiToken)
	d.Set("rest_version", expected.Rest_version)
	d.Set("verify_https", expected.Verify_https)
	d.Set("ssl_cert", expected.Ssl_cert)
	d.Set("user_agent", expected.User_agent)

	actual, err := NewConfig(d)
	if err != nil {
		t.Fatalf("error creating new configuration: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func TestNewConfigWithAllAuth(t *testing.T) {
	expected := &Config{
		Username:       "foo",
		Password:       "bar",
		Target:         "purestorage.flasharray",
		ApiToken:       "foobar",
		Rest_version:   "1.15",
		Verify_https:   false,
		Ssl_cert:       false,
		User_agent:     "useragent",
		Request_kwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.ApiToken)
	d.Set("rest_version", expected.Rest_version)
	d.Set("verify_https", expected.Verify_https)
	d.Set("ssl_cert", expected.Ssl_cert)
	d.Set("user_agent", expected.User_agent)

	_, err := NewConfig(d)
	if err == nil {
		t.Fatalf("error NOT generated when username, password, and api_token provided.")
	}
}
