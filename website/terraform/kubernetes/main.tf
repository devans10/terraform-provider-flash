data "digitalocean_kubernetes_cluster" "production" {
  name = "${var.cluster_name}"
}

data "digitalocean_loadbalancer" "nginx-ingress" {
  name = "nginx-ingress"
}
