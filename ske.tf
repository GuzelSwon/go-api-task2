resource "stackit_ske_cluster" "ske" {
  project_id         = var.project_id
  name               = var.cluster_name
  kubernetes_version = "1.28"
  node_pools = [
    {
      name               = var.cluster_name
      machine_type       = "c1.4"
      os_version         = "3815.2.2"
      minimum            = "1"
      maximum            = "1"
      availability_zones = ["eu01-1"]
    }
  ]
  maintenance = {
    enable_kubernetes_version_updates    = true
    enable_machine_image_version_updates = true
    start                                = "01:00:00Z"
    end                                  = "02:00:00Z"
  }
}

