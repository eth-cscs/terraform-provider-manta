resource "manta_cloudinit_defaults" "defaults" {
  base_url    = var.cloud_init
  public_keys = var.public_keys
}
