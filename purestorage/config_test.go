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
		Username:      os.Getenv("PURE_USERNAME"),
		Password:      os.Getenv("PURE_PASSWORD"),
		Target:        os.Getenv("PURE_TARGET"),
		APIToken:      os.Getenv("PURE_APITOKEN"),
		RestVersion:   "",
		VerifyHTTPS:   false,
		SslCert:       false,
		RequestKwargs: nil,
	}
}

func TestAccClient(t *testing.T) {
	testAccClientPreCheck(t)

	c := testAccClientGenerateConfig(t)

	_, err := c.Client()
	if err != nil {
		t.Fatalf("error stting up client: %s", err)
	}
}

func TestNewConfigWithApiToken(t *testing.T) {
	expected := &Config{
		Username:      "",
		Password:      "",
		Target:        "purestorage.flasharray",
		APIToken:      "foobar",
		RestVersion:   "1.15",
		VerifyHTTPS:   false,
		SslCert:       false,
		RequestKwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.APIToken)
	d.Set("rest_version", expected.RestVersion)
	d.Set("verify_https", expected.VerifyHTTPS)
	d.Set("ssl_cert", expected.SslCert)

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
		Username:      "foo",
		Password:      "bar",
		Target:        "purestorage.flasharray",
		APIToken:      "",
		RestVersion:   "1.15",
		VerifyHTTPS:   false,
		SslCert:       false,
		RequestKwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.APIToken)
	d.Set("rest_version", expected.RestVersion)
	d.Set("verify_https", expected.VerifyHTTPS)
	d.Set("ssl_cert", expected.SslCert)

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
		Username:      "foo",
		Password:      "bar",
		Target:        "purestorage.flasharray",
		APIToken:      "foobar",
		RestVersion:   "1.15",
		VerifyHTTPS:   false,
		SslCert:       false,
		RequestKwargs: map[string]string{},
	}

	r := &schema.Resource{Schema: Provider().(*schema.Provider).Schema}
	d := r.Data(nil)
	d.Set("username", expected.Username)
	d.Set("password", expected.Password)
	d.Set("target", expected.Target)
	d.Set("api_token", expected.APIToken)
	d.Set("rest_version", expected.RestVersion)
	d.Set("verify_https", expected.VerifyHTTPS)
	d.Set("ssl_cert", expected.SslCert)

	_, err := NewConfig(d)
	if err == nil {
		t.Fatalf("error NOT generated when username, password, and api_token provided.")
	}
}
