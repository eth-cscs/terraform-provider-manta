resource "manta_group" "compute" {
  label   = "compute"
  members = local.list_nodes
}
