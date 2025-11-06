resource "manta_cloudinit_group" "compute" {
  name        = "compute"
  description = "The compute group"
  file = {
    content  = file(var.filename_cloud_init_group)
    encoding = "base64"
  }
}
