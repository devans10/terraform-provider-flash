# Terraform Provider Flash

[![Build Status](https://travis-ci.com/devans10/terraform-provider-flash.svg?branch=master)](https://travis-ci.com/devans10/terraform-provider-flash)

## Note: Breaking change with Terraform 0.13. Check the Usage updates below.

This is the repository for the Terraform Provider Flash.  The plugin provides resources for the Pure Storage FlashArray to be managed within Terraform.

For general information about Terraform, visit the [official website](https://terraform.io) and the [GitHub project page.](https://github.com/hashicorp/terraform)

As of version 1.1.2, the provider is available in the [Terraform Registry](https://registry.terraform.io/providers/devans10/flash/latest?pollNotifications=true), and can be downloaded automatically when running `terraform init`.  

The documentation for the provider can be found on the [Provider's website](https://www.terraform-provider-flash.com)

This provider plugin is maintained by Dave Evans.

Please submit issues [here](https://github.com/devans10/terraform-provider-flash/issues).

## Requirements

------------

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x (the provider was tested against 0.13.1)
- [Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

## Usage

------------
Add the `required_providers` block to your terraform configuration.

```sh
terraform {
  required_providers {
    flash = {
      source  = "devans10/flash"
      version = "~> 1.1.2"
    }
  }
}
```

You will also need to list the provider in all of the resources.

```sh
resource "purestorage_volume" "vol1" {
  provider = flash
  name     = "volume_name"
  size     = 1073741824
}
```


*(Deprecated)*
Add the following to `$HOME/.terraformrc`

```sh

providers {
    purestorage = "$GOPATH/bin/terraform-provider-flash"
}

```

## Configure the Provider

***Configure in TF configuration***

```sh
provider "flash" {
  api_token  = "${var.purestorage_apitoken}"
  target     = "${var.purestorage_target}"
}
```

or

```sh
provider "flash" {
  username   = "${var.purestorage_username}"
  password   = "${var.purestorage_password}"
  target     = "${var.purestorage_target}"
}
```

***Configure in environment***

Set username(`PURE_USERNAME`) and password(`PURE_PASSWORD`) or API Token(`PURE_APITOKEN`) and endpoint(`PURE_TARGET`) in environment.

```sh
provider "flash" {}
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
  provider = flash
  name     = "testvol_tf"
  size     = "1048000000"
}
```

Copy a volume

```sh
resource "purestorage_volume" "testvol_tf_copy" {
  provider = flash
  name     = "testvol_tf_copy"
  source   = "testvol_tf"
}
```

Create a host

```sh
resource "purestorage_host" "testhosttf" {
  provider = flash
  name     = "testhosttf"
  volume {
    vol = "testvol_tf"
    lun = 1
  }
}
```

Create a hostgroup

```sh
resource "purestorage_hostgroup" "testhgrouptf" {
  provider = flash
  name     = "testhgrouptf"
  hosts    = ["testhosttf"]
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
  provider = flash
  name     = "testpgroup"
  volumes  = ["testvol_tf", "testvol_tf_copy"]
}
```

## Developing the Provider

------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

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

## Disclaimer

terraform-provider-flash and its developer(s) are not affiliated with or sponsored by Pure Storage.  The statements and opinions on this site are those of the developer(s) and do not necessarily represent those of Pure Storage. Pure Storage and the Pure Storage trademarks listed at [https://www.purestorage.com/pure-folio/showcase.html?type=pdf&path=/content/dam/pdf/en/legal/external-trademark-list.pdf](https://www.purestorage.com/pure-folio/showcase.html?type=pdf&path=/content/dam/pdf/en/legal/external-trademark-list.pdf) are trademarks of Pure Storage, Inc.
