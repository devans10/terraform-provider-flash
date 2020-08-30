# Volume

Provides a Pure Storage volume resource

## Example Usage

```sh
resource "purestorage_volume" "vol" {
  provider = flash
  name     = "volume_name"
  size     = 1073741824
}
```

## Argument Reference

The following arguments are supported:

+ `name` - (Required) The name of the volume.
+ `size` - (Optional) The size of the volume in bytes. type: integer
+ `source` - (Optional) The source volume to copy.

*NOTE: `size` or `source` can be specified upon volume creation, but not both.*

## Attribute Reference

The following attributes are exported:

+ `id` - The ID of the volume.
+ `name` - The name of the volume.
+ `size` - The size of the volume in bytes. type: integer
+ `source` - The source of volume.
+ `serial` - The serial ID of the volume.
+ `created` - The date volume was created. 

## Import

volume can be imported using the volume name

```sh
terraform import purestorage_volume vol
```
