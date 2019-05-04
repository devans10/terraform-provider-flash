resource "digitalocean_record" "sitename" {
  domain = "terraform-purestorage.com"
  type = "A"
  name = "${var.sitename}"
  value = "${var.ipaddress}"
}
