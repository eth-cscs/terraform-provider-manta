# File: variables.tf

variable "base_url" {
  description = "OpenCHAMI base URL"
  type        = string
  default     = "https://foobar.openchami.cluster:8443"
}

variable "access_token" {
  description = "OpenCHAMI access token"
  type        = string
  nullable    = false
}
