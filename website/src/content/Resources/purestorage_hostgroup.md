---
title: "purestorage_hostgroup"
date: 2019-03-22T21:19:24-04:00
lastmod: 2019-03-22T21:19:24-04:00
draft: false
description: ""
weight: 5
---

Provides a Pure Storage hostgroup resource.

## Example Usage

```
resource "purestorage_hostgroup" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

+ `name` - (Required) The name of the hostgroup
+ `hosts` - (Optional) List of member hosts
+ `volume` - (Optional) Shared volume connection
  + `vol` - Volume name to connect.
  + `lun` - LUN ID for the volume.

## Attribute Reference

The following attributes are exported:

+ `id` - The ID of the host
+ `name` - The name of the host
+ `hosts` - The list of member hosts
+ `volume` - Shared volume connection
  + `vol` - Volume name to connect.
  + `lun` - LUN ID for the volume.

## Import

hostgroups can be imported using the hostgroup name

```
terraform import purestorage_hostgroup example
```
