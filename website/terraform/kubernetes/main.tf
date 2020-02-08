data "digitalocean_kubernetes_cluster" "production" {
  name = "${var.cluster_name}"
}

data "digitalocean_loadbalancer" "nginx-ingress" {
  name = "a733893ed9a4f47ecb52df353f062287"
}
