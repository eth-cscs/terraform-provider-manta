resource "manta_rfe" "rfe" {
  id                 = rfe.xname
  user               = "root"
  hostname           = rfe.hostname
  rediscoveronupdate = true
  enabled            = true
  password_wo        = rfe.password
}
