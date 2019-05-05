locals {
  dockercfg = {
    "${var.docker_server}" = {
      email    = "${var.docker_email}"
      username = "${var.docker_username}"
      password = "${var.docker_password}"
    }
  }
}
resource "kubernetes_secret" "tf-regsecret" {
  metadata {
    name = "tf-regsecret"
  }

  data {
    ".dockercfg" = "${ jsonencode(local.dockercfg) }"
  }

  type = "kubernetes.io/dockercfg"
}

resource "kubernetes_deployment" "terraform-purestorage" {
  metadata {
    name = "terraform-purestorage"
    labels {
      app = "terraform-purestorage"
    }
  }

  spec {
    replicas = 2

    selector {
      match_labels {
        app = "terraform-purestorage"
      }
    }

    template {
      metadata {
        labels {
          app = "terraform-purestorage"
        }
      }

      spec {
        container {
           name = "terraform-purestorage"
           image = "${var.image}"
           image_pull_policy = "Always"
           port {
             container_port = "80"
           }
         }
         image_pull_secrets {
           name = "${kubernetes_secret.tf-regsecret.metadata.0.name}"
         }
       }
     }
   }
}


resource "kubernetes_service" "terraform-purestorage" {
  metadata {
    name = "terraform-purestorage"
  }
  spec {
    selector {
      app = "${kubernetes_deployment.terraform-purestorage.metadata.0.labels.app}"
    }
    port {
      port = 80
      target_port = 80
    }
  }
}


