package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBssResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_bss_boot_parameters" "test" {
  macs   = ["00:de:ad:be:ef:00"]
  params = "params"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_bss_boot_parameters.test", "params", "params"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_bss_boot_parameters" "test" {
  macs   = ["00:de:ad:be:ef:00"]
  params = "params"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_bss_boot_parameters.test", "params", "params"),
				),
			},
		},
	})
}
