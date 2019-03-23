---
title: "purestorage_flasharray"
date: 2019-03-22T21:09:33-04:00
lastmod: 2019-03-22T21:09:33-04:00
draft: false
description: ""
weight: 5
---

Get information on a flasharray.  This data source provides the name, version, revision of a flasharray.  This is useful if the flasharray is not managed by Terraform, or you need to utilize any of the flasharray data.

## Example Usage

Get the flasharray

```
data "purestorage_flasharray" "example" {
  name = "example"
}
```

## Argument Reference

The following arguements are supported:

+ name - (Required) The name of the flasharray

## Attribute Reference

The following attributes are exported:

+ `name`: Name of the flasharray
+ `revision`: Revision of the flasharray
+ `version`: The version of the flasharray
