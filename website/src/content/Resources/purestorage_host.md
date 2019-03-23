---
title: "purestorage_host"
date: 2019-03-22T21:19:24-04:00
lastmod: 2019-03-22T21:19:24-04:00
draft: false
description: ""
weight: 5
---

Provides a Pure Storage host resource.

## Example Usage

```
resource "purestorage_host" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

+ `name` - (Required) The name of the host
+ `iqn` - (Optional) List of iSCSI qualified names (IQNs) to the specified host.
+ `wwn` - (Optional) List of Fibre Channel worldwide names (WWNs) to the specified host.
+ `nqn` - (Optional) List of NVMeF qualified names (NQNs) to the specified host.
+ `host_password - (Optional) Host password for CHAP authentication.
+ `host_user` - (Optional) Host username for CHAP authentication.
+ `personality` - (Optional) Determines how the Purity system tunes the protocol used between the array and the initiator. One of "aix", "esxi", "hitachi-vsp", "hpux", "oracle-vm-server", "solaris", "vms", or null
+ `preferred_array - (Optional) List of preferred arrays.
+ `target_password` - (Optional) Target password for CHAP authentication.
+ `target_user` - (Optional) Target username for CHAP authentication.
+ `volume` - (Optional) Private volume connection
  + `vol` - Volume name to connect.
  + `lun` - LUN ID for the volume.

## Attribute Reference

The following attributes are exported:

+ `id` - The ID of the host
+ `name` - The name of the host
+ `iqn` - List of iSCSI qualified names (IQNs) to the specified host.
+ `wwn` - List of Fibre Channel worldwide names (WWNs) to the specified host.
+ `nqn` - List of NVMeF qualified names (NQNs) to the specified host.
+ `host_password - Host password for CHAP authentication.
+ `host_user` - Host username for CHAP authentication.
+ `personality` - Determines how the Purity system tunes the protocol used between the array and the initiator. One of "aix", "esxi", "hitachi-vsp", "hpux", "oracle-vm-server", "solaris", "vms", or null
+ `preferred_array - List of preferred arrays.
+ `hgroup` - hostgroup the host is a member of.
+ `target_password` - Target password for CHAP authentication.
+ `target_user` - Target username for CHAP authentication.
+ `volume` - Private volume connection
  + `vol` - Volume name to connect.
  + `lun` - LUN ID for the volume.

## Import

hosts can be imported using the host name

```
terraform import purestorage_host example
```
