resource "digitalocean_record" "sitename" {
  domain = "terraform-provider-flash.com"
  type = "A"
  name = var.sitename
  value = var.ipaddress
}
