# Terraform Provider Flash

[![Build Status](https://travis-ci.com/devans10/terraform-provider-flash.svg?branch=master)](https://travis-ci.com/devans10/terraform-provider-flash)

This is the repository for the Terraform Flash Provider.  The plugin provides resources for the Pure Storage FlashArray to be managed within Terraform.

For general information about Terraform, visit the [official website](https://terraform.io) and the [GitHub project page.](https://github.com/hashicorp/terraform)

This provider plugin is maintained by Dave Evans.

Please submit issues [here](https://github.com/devans10/terraform-provider-flash/issues).

## Requirements

------------

- [Terraform](https://www.terraform.io/downloads.html) 0.11.x (the provider was tested against 0.11.10)
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Usage

------------
Add the following to `$HOME/.terraformrc`

```sh

providers {
    purestorage = "$GOPATH/bin/terraform-provider-flash"
}

```

## Configure the Provider

***Configure in TF configuration***

```sh
provider "purestorage" {
  api_token  = "${var.purestorage_apitoken}"
  target     = "${var.purestorage_target}"
}
```

or

```sh
provider "purestorage" {
  username   = "${var.purestorage_username}"
  password   = "${var.purestorage_password}"
  target     = "${var.purestorage_target}"
}
```

***Configure in environment***

Set username(`PURE_USERNAME`) and password(`PURE_PASSWORD`) or API Token(`PURE_APITOKEN`) and endpoint(`PURE_TARGET`) in environment.

```sh
provider "purestorage" {}
```

## Building The Provider

------------

Clone repository to: `$GOPATH/src/github.com/devans10/terraform-provider-flash`

```sh
mkdir -p $GOPATH/src/github.com/devans10; cd $GOPATH/src/github.com/devans10
git clone git@github.com:devans10/terraform-provider-flash
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/devans10/terraform-provider-flash
make build
```

Using the provider

------------

## Basic volume provision

Create one volume

```sh
resource "purestorage_volume" "testvol_tf" {
  name = "testvol_tf"
  size = "1048000000"
}
```

Copy a volume

```sh
resource "purestorage_volume" "testvol_tf_copy" {
  name = "testvol_tf_copy"
  source = "testvol_tf"
}
```

Create a host

```sh
resource "purestorage_host" "testhosttf" {
  name = "testhosttf"
  volume {
    vol = "testvol_tf"
    lun = 1
  }
}
```

Create a hostgroup

```sh
resource "purestorage_hostgroup" "testhgrouptf" {
  name = "testhgrouptf"
  hosts = ["testhosttf"]
  volume {
    vol = "testvol_tf_copy"
    lun = 250
  }
}
```

Create a Protection Group

Protection Group has a hosts, hgroups, and volumes parameters, but only 1 can be used.

```sh
resource "purestorage_protectiongroup" "testpgroup" {
  name = "testpgroup"
  volumes = ["testvol_tf", "testvol_tf_copy"]
}
```

## Developing the Provider

------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-flash
...
```

In order to test the provider, you can simply run `make test`.

```sh
make test
```

To run acceptance tests, run `make testacc`.
Volumes and Protection Groups created during the acceptance tests are not eradicated.

```sh
make testacc
```
