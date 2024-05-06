resource "stackit_ske_cluster" "<name>" {
  project_id         = var.project_id
  name               = var.cluster_name
  kubernetes_version = "1.27"
  node_pools = [
    {
      name               = var.cluster_name
      machine_type       = "c1.3"
      os_version         = "3760.2.0"
      minimum            = "1"
      maximum            = "1"
      availability_zones = ["eu01-1", "eu01-2", "eu01-3"]
    }
  ]
  maintenance = {
    enable_kubernetes_version_updates    = true
    enable_machine_image_version_updates = true
    start                                = "01:00:00Z"
    end                                  = "02:00:00Z"
  }
}