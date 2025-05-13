<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_manta"></a> [manta](#requirement\_manta) | >= 0.0.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_manta"></a> [manta](#provider\_manta) | >= 0.0.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| manta_node.node | resource |
| manta_rfe.rfe | resource |
| manta_rfe.rfe_all | resource |
| manta_version.version | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_access_token"></a> [access\_token](#input\_access\_token) | OpenCHAMI access token | `string` | n/a | yes |
| <a name="input_base_url"></a> [base\_url](#input\_base\_url) | OpenCHAMI base URL | `string` | `"https://scitas.openchami.cluster:8443"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_manta_version"></a> [manta\_version](#output\_manta\_version) | n/a |
| <a name="output_node"></a> [node](#output\_node) | n/a |
| <a name="output_rfe"></a> [rfe](#output\_rfe) | n/a |
| <a name="output_rfe_all"></a> [rfe\_all](#output\_rfe\_all) | n/a |
<!-- END_TF_DOCS -->