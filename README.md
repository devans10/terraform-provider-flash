Terraform Provider Purestorage
==================
<img src="https://travis-ci.com/devans10/terraform-provider-purestorage.svg?branch=master>

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by Dave Evans.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)
-	[Pure Storage Go Client](https://github.com/devans10/go-purestorage)

Usage
---------------------
Add the following to `$HOME/.terraformrc`

```
providers {
    purestorage = "$GOPATH/bin/terraform-provider-purestorage"
}
```

**Configure the Provider**

***Configure in TF configuration***

```
provider "purestorage" {
  api_token  = "${var.purestorage_apitoken}"
  target     = "${var.purestorage_target}"
}
```

or

```
provider "purestorage" {
  username   = "${var.purestorage_username}"
  password   = "${var.purestorage_password}"
  target     = "${var.purestorage_target}"
}
```


***Configure in environment***

Set username(`PURESTORAGE_USERNAME`) and password(`PURESTORAGE_PASSWORD`) or API Token(`PURESTORAGE_APITOKEN`) and endpoint(`PURESTORAGE_TARGET`) in environment.
```
provider "purestorage" {}
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/devans10/terraform-provider-purestorage`

```sh
$ mkdir -p $GOPATH/src/github.com/devans10; cd $GOPATH/src/github.com/devans10
$ git clone git@github.com:devans10/terraform-provider-purestorage
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/devans10/terraform-provider-purestorage
$ make build
```

Using the provider
----------------------

**Basic volume provision**

Create one volume
```
resource "purestorage_volume" "testvol_tf" {
	name = "testvol_tf"
	size = "1G"
}
```

Copy a volume
```
resource "purestorage_volume" "testvol_tf_copy" {
	name = "testvol_tf_copy"
	source = "testvol_tf"
}
```

Create a host
```
resource "purestorage_host" "testhosttf" {
	name = "testhosttf"
	connected_volumes = ["testvol_tf"]
}
```

Create a hostgroup
```
resource "purestorage_hostgroup" "testhgrouptf" {
	name = "testhgrouptf" 
	hosts = ["testhosttf"]
	connected_volumes = ["testvol_tf_copy"]
}
```

Create a Protection Group

Protection Group has a hosts, hgroups, and volumes parameters, but only 1 can be used.
```
resource "purestorage_protectiongroup" "testpgroup" {
	name = "testpgroup"
	volumes = ["testvol_tf", "testvol_tf_copy"]
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-purestorage
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

To run acceptance tests, run `make testacc`.
Volumes created during the acceptance tests are not eradicated. 

```sh
$ make testacc
```
