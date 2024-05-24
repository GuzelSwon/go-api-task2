terraform {
    required_providers {
        stackit = {
            source = "stackitcloud/stackit"
            version = "0.16.0"
        }
        helm = {
            source  = "hashicorp/helm"
            version = "2.12.1"
        }
        kubectl = {
            source = "gavinbunney/kubectl"
            version = "1.14.0"
        }
    }
}