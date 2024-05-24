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
variable "github_repo_url" {
  description = "Github repository URL of an application"
  default = "https://github.com/GuzelSwon/go-api-task2-k8s-deployment.git"
}
variable "auth_github_username" {
  description = "Github repository URL of an application"
  default = "GuzelSwon"
}
variable "auth_github_token" {
  description = "Github repository URL of an application"
  sensitive = true
}