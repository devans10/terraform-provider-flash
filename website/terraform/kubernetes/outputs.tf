output "cluster_endpoint" {
  value = "${data.digitalocean_kubernetes_cluster.production.endpoint}"
}

output "client_certificate" {
  value = "${base64decode(data.digitalocean_kubernetes_cluster.production.kube_config.0.client_certificate)}"
}

output "client_key" {
  value = "${base64decode(data.digitalocean_kubernetes_cluster.production.kube_config.0.client_key)}"
}

output "cluster_ca_certificate" {
  value = "${base64decode(data.digitalocean_kubernetes_cluster.production.kube_config.0.cluster_ca_certificate)}"
}

output "nginx_ingress_ip" {
  value = "${data.digitalocean_loadbalancer.nginx-ingress.ip}"
}
