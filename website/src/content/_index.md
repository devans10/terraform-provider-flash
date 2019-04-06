---
title: "Pure Storage Provider"
date: 2019-03-22T20:49:09-04:00
lastmod: 2019-03-22T20:49:09-04:00
draft: false
description: ""
weight: 5
---

# Pure Storage Provider

The Pure Storage provider is used to interact with the resources supported by the Pure Storage FlashArray.  The provider needs to be configured with the proper credentials before it can be used.

This is an open source community project and is not affiliated with the Pure Storage or HashiCorp companies.

## Installation

Download the provider from the [downloads](/downloads) page.
Copy the binary file to the user plugin directory, located at `%APPDATA%\terraform.d\plugins` on Windows and `~/.terraform.d/plugins` on Linux and MacOS.


## Building the provider from source

Clone repository to: `$GOPATH/src/github.com/devans10/terraform-provider-purestorage`

```
$ mkdir -p $GOPATH/src/github.com/devans10; cd $GOPATH/src/github.com/devans10
$ git clone git@github.com:devans10/terraform-provider-purestorage
```

Enter the provider directory and build the provider

```
$ cd $GOPATH/src/github.com/devans10/terraform-provider-purestorage
$ make build
```
Copy the binary file to the user plugin directory, located at `%APPDATA%\terraform.d\plugins` on Windows and `~/.terraform.d/plugins` on Linux an
d MacOS.
