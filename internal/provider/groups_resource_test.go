package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_group" "test" {
  label       = "test"
  description = "Test group"
  members     = ["x1000c0s0b0n0"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_group.test", "label", "test"),
					resource.TestCheckResourceAttr("manta_group.test", "description", "Test group"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_group" "test" {
  label       = "test"
  description = "Test group"
  members     = ["x1000c0s0b0n0"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_group.test", "label", "test"),
					resource.TestCheckResourceAttr("manta_group.test", "description", "Test group"),
				),
			},
		},
	})
}
