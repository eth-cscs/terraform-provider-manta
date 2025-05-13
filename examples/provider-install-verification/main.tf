terraform {
  required_providers {
    manta = {
      source  = "registry.terraform.io/scitas/manta" # to be confirmed
      version = ">= 0.0.1"
    }
  }
  required_version = ">= 1.0.0"
}

provider "manta" {
  base_url     = var.base_url
  access_token = var.access_token
}

resource "manta_rfe" "rfe" {
  id   = "x1002c0s0b6"
  user = "root"
}

resource "manta_rfe" "rfe_all" {
  id                 = "x1002c0s0b10"
  user               = "user"
  hostname           = "hostname"
  rediscoveronupdate = true
  enabled            = true
  password           = "password"
}

resource "manta_node" "node" {
  id    = "x1000c0s0b1n1"
  state = "On"
}

data "manta_version" "version" {}

output "manta_version" {
  value = data.manta_version.version
}

output "rfe" {
  value = manta_rfe.rfe
}

output "rfe_all" {
  value = manta_rfe.rfe_all
}

output "node" {
  value = manta_node.node
}
