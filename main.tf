provider "stackit" {
  region = "eu01"
  service_account_token = var.service_account_token
}

terraform {
  backend "s3" {
    bucket = "go-api-task"
    key    = "go-api-task/tfstate"
    endpoints = {
      s3 = "https://object.storage.eu01.onstackit.cloud"
    }
    region                      = "eu01"
    skip_credentials_validation = true
    skip_region_validation      = true
    skip_s3_checksum            = true
    skip_requesting_account_id  = true
    secret_key                  = "${var.object_storage_secret_key}"
    access_key                  = "${var.object_storage_access_key}"
  }
}

resource "stackit_ske_kubeconfig" "kubeconfig" {
  project_id   = var.project_id
  cluster_name = stackit_ske_cluster.ske.name
  refresh      = true
}

resource "kubernetes_secret" "repo_access" {
  depends_on = [helm_release.argocd]
  metadata {
    name      = "repo-access"
    namespace = "argocd"
    labels = {
      "argocd.argoproj.io/secret-type" = "repository"
    }
  }
 
  data = {
    "type"          = "git"
    "name"          = "deployment"
    "url"           = var.go_api_app_github_repo_url
    "project"       = "default"
    "username"      = var.auth_github_username
    "password"      = var.auth_github_token
  }
 
  type = "Opaque"
}

provider "helm" {
  kubernetes {
    config_path = null
    config_context = null
    host = yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster.server

    client_certificate     = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-certificate-data"])
    client_key             = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-key-data"])
    cluster_ca_certificate = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster["certificate-authority-data"])
  }
}

provider "kubectl" {
  host = yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster.server

  client_certificate     = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-certificate-data"])
  client_key             = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-key-data"])
  cluster_ca_certificate = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster["certificate-authority-data"])
  load_config_file       = false
}

provider "kubernetes" {
  config_path = null
  host = yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster.server

  client_certificate     = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-certificate-data"])
  client_key             = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).users[0].user["client-key-data"])
  cluster_ca_certificate = base64decode(yamldecode(stackit_ske_kubeconfig.kubeconfig.kube_config).clusters[0].cluster["certificate-authority-data"])
}

resource "helm_release" "argocd" {
  name = "argocd"
  repository = "https://argoproj.github.io/argo-helm"
  chart = "argo-cd"
  namespace = "argocd"
  create_namespace = true
  timeout = 600
  skip_crds = true
}

resource "stackit_secretsmanager_instance" "secretsmanager_instance" {
  project_id = var.project_id
  name       = "instance1"
}

resource "stackit_secretsmanager_user" "secretsmanager_user" {
  project_id    = var.project_id
  instance_id   = stackit_secretsmanager_instance.secretsmanager_instance.instance_id
  description   = "User1"
  write_enabled = false
}

resource "kubernetes_secret" "dockerhub-secret" {
  depends_on = [stackit_ske_cluster.ske]
  metadata {
    name      = "dockerhub-secret"
    namespace = "default"
  }
 
  type = "kubernetes.io/dockerconfigjson"
 
  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        "${var.registry_url}" = {
          "username" = var.registry_username,
          "password" = var.registry_password,
          "email"    = var.registry_email
          "auth"     = base64encode("${var.registry_username}:${var.registry_password}")
        }
      }
    })
  }
}

resource "kubectl_manifest" "argocd_go_api_app" {
  depends_on = [kubernetes_secret.dockerhub-secret, kubernetes_secret.repo_access, helm_release.argocd, stackit_ske_cluster.ske]
  yaml_body = templatefile("${path.module}/argocd_go-api-app.yaml", {
    github_repo_url = var.go_api_app_github_repo_url
    helm_chart_path = "helm-chart"
    environment = var.environment
    resource_name = "go-api-app"

    secretsmanager_instance_id = stackit_secretsmanager_user.secretsmanager_user.instance_id
    secretsmanager_username = stackit_secretsmanager_user.secretsmanager_user.username
    postgresql_secretname = var.postgresql_secretname
  })
}

resource "kubectl_manifest" "argocd_postgresql" {
  depends_on = [kubernetes_secret.repo_access, helm_release.argocd, stackit_ske_cluster.ske]
  yaml_body = templatefile("${path.module}/argocd_template.yaml", {
    github_repo_url = var.bitnami_github_repo_url
    helm_chart_path = "bitnami/postgresql"
    environment = var.environment
    resource_name = "postgresql"
  })
}

resource "kubectl_manifest" "argocd_fluentbit" {
  depends_on = [helm_release.argocd, stackit_ske_cluster.ske]
  yaml_body = templatefile("${path.module}/argocd_template.yaml", {
    github_repo_url = var.fluentbit_github_repo_url
    helm_chart_path = "charts/fluent-bit"
    environment = var.environment
    resource_name = "fluentbit"
  })
}

resource "kubectl_manifest" "argocd_otel" {
  depends_on = [helm_release.argocd, stackit_ske_cluster.ske]
  yaml_body = templatefile("${path.module}/argocd_otel.yaml", {
    github_repo_url = var.otel_github_repo_url
    helm_chart_path = "charts/opentelemetry-collector"
    environment = var.environment
    resource_name = "otel"
    mode = "daemonset"
    logs_collector_enabled = "true"
    logs_collector_include_collector_logs = "true"
    image_repository = "otel/opentelemetry-collector-k8s"
  })
}

resource "kubectl_manifest" "argocd_prometheus" {
  depends_on = [helm_release.argocd, stackit_ske_cluster.ske]
  yaml_body = templatefile("${path.module}/argocd_template.yaml", {
    github_repo_url = var.prometheus_github_repo_url
    helm_chart_path = "charts/prometheus"
    environment = var.environment
    resource_name = "prometheus"
  })
}