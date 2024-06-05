variable "service_account_token" {
  description = "A service account token"
  sensitive = true
}
variable "project_id" {
  description = "The project ID to host the cluster in"
  default     = "c17114ca-1a52-411d-a019-9c3e33f218e7"
}
variable "cluster_name" {
  description = "The name for the GKE cluster"
  default     = "cl-ythb0xi4"
}
variable "object_storage_access_key" {
  description = "An access key to object storage"
  sensitive = true
}
variable "object_storage_secret_key" {
  description = "A secret key to object storage"
  sensitive = true
}
variable "s3_bucket_name" {
  description = "The name of S3 bucket"
  default     = "go-api-app"
}
variable "go_api_app_github_repo_url" {
  description = "Github repository URL of an application"
  default = "https://github.com/GuzelSwon/go-api-task2-k8s-deployment.git"
}
variable "bitnami_github_repo_url" {
  description = "Github repository URL of bitnami"
  default = "https://github.com/bitnami/charts.git"
}
variable "fluentbit_github_repo_url" {
  description = "Github repository URL of fluentbit"
  default = "https://github.com/fluent/helm-charts.git"
}
variable "otel_github_repo_url" {
  description = "Github repository URL of Open Telemetry"
  default = "https://github.com/open-telemetry/opentelemetry-helm-charts.git"
}
variable "prometheus_github_repo_url" {
  description = "Github repository URL of prometheus"
  default = "https://github.com/prometheus-community/helm-charts.git"
}
variable "jaeger_github_repo_url" {
  description = "Github repository URL of jaeger"
  default = "https://github.com/jaegertracing/helm-charts.git"
}
variable "auth_github_username" {
  description = "Github repository URL of an application"
  default = "GuzelSwon"
}
variable "auth_github_token" {
  description = "Github repository URL of an application"
  sensitive = true
}
variable "registry_url" {
  default = "docker.io"
}
variable "registry_username" {
  default = "guzelkhuziakhmetova"
}
variable "registry_password" {
  sensitive = true
}
variable "registry_email" {
  default = "guzelkhuz@gmail.com"
}
variable "postgresql_secretname" {
  description = "Secret name of mysql"
  default = "postgresql-production"
}
variable "environment" {
  description = "Github environment value"
}