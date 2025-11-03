package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCloudInitResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_cloudinit_defaults" "test" {
  base_url    = "http://cloud-init/cloud-init"
  public_keys = ["ssh-ed25519 AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA user1@demo-head"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_cloudinit_defaults.test", "base_url", "http://cloud-init/cloud-init"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_cloudinit_defaults" "test" {
  base_url    = "http://cloud-init/cloud-init"
  public_keys = ["ssh-ed25519 AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA user1@demo-head"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_cloudinit_defaults.test", "base_url", "http://cloud-init/cloud-init"),
				),
			},
		},
	})
}
