resource "manta_bss_boot_parameters" "default" {
  macs   = [node.mac]
  kernel = var.bss_kernel
  initrd = var.bss_initrd
  params = var.bss_params
}
