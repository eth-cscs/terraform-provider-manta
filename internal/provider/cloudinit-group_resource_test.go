package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCloudInitGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_cloudinit_group" "test" {
  description = "The compute group"
  file = {
    content  = "IyMgdGVtcGxhdGU6IGppbmphCiNjbG91ZC1jb25maWcKbWVyZ2VfaG93OgotIG5hbWU6IGxpc3QKICBzZXR0aW5nczogW2FwcGVuZF0KLSBuYW1lOiBkaWN0CiAgc2V0dGluZ3M6IFtub19yZXBsYWNlLCByZWN1cnNlX2xpc3RdCnVzZXJzOgogIC0gbmFtZTogcm9vdAogICAgc3NoX2F1dGhvcml6Z      WRfa2V5czoge3sgZHMubWV0YV9kYXRhLmluc3RhbmNlX2RhdGEudjEucHVibGljX2tleXMgfX0KZGlzYWJsZV9yb290OiBmYWxzZQo="
    encoding = "base64"
  }
  name = "compute"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_cloudinit_group.test", "name", "compute"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_cloudinit_group" "test" {
  description = "The compute group"
  file = {
    content  = "IyMgdGVtcGxhdGU6IGppbmphCiNjbG91ZC1jb25maWcKbWVyZ2VfaG93OgotIG5hbWU6IGxpc3QKICBzZXR0aW5nczogW2FwcGVuZF0KLSBuYW1lOiBkaWN0CiAgc2V0dGluZ3M6IFtub19yZXBsYWNlLCByZWN1cnNlX2xpc3RdCnVzZXJzOgogIC0gbmFtZTogcm9vdAogICAgc3NoX2F1dGhvcml6Z      WRfa2V5czoge3sgZHMubWV0YV9kYXRhLmluc3RhbmNlX2RhdGEudjEucHVibGljX2tleXMgfX0KZGlzYWJsZV9yb290OiBmYWxzZQo="
    encoding = "base64"
  }
  name = "compute"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_cloudinit_group.test", "name", "compute"),
				),
			},
		},
	})
}
