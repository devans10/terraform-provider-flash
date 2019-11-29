data "digitalocean_kubernetes_cluster" "production" {
  name = "${var.cluster_name}"
}

data "digitalocean_loadbalancer" "nginx-ingress" {
  name = "ab125616b9d3311e9bbae1a4e7fc0bf1"
}
