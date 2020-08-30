# Flasharray

Get information on a FlashArray.  This data source provides the name, version, revision of a FlashArray.  This is useful if the FlashArray is not managed by Terraform, or you need to utilize any of the FlashArray data.

## Example Usage

Get the FlashArray

```sh
data "purestorage_flasharray" "example" {
  provider = flash
  name     = "example"
}
```

## Argument Reference

The following arguements are supported:

+ name - (Required) The name of the FlashArray

## Attribute Reference

The following attributes are exported:

+ `name`: Name of the FlashArray
+ `revision`: Revision of the FlashArray
+ `version`: The version of the FlashArray
