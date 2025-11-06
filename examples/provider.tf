terraform {
  required_providers {
    manta = {
      source  = "registry.terraform.io/eth-cscs/manta"
      version = ">= 0.0.1"
    }
  }
  required_version = ">= 1.0.0"
}
