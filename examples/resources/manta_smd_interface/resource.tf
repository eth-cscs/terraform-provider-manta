resource "manta_smd_interface" "eth" {
  component_id = node.xname
  mac_address  = node.mac
  ip_addresses = [node.ip]
}
