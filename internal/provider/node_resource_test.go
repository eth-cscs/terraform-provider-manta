package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNodeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_node" "test" {
  id    = "x1000c0s0b1n1"
  state = "On"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_node.test", "id", "x1000c0s0b1n1"),
					resource.TestCheckResourceAttr("manta_node.test", "state", "On"),
					resource.TestCheckResourceAttr("manta_node.test", "arch", "X86"),
					resource.TestCheckResourceAttr("manta_node.test", "class", "River"),
					resource.TestCheckResourceAttr("manta_node.test", "enabled", "true"),
					resource.TestCheckResourceAttr("manta_node.test", "flag", "OK"),
					resource.TestCheckResourceAttr("manta_node.test", "nettype", "Sling"),
					resource.TestCheckResourceAttr("manta_node.test", "nid", "16400389"),
					resource.TestCheckResourceAttr("manta_node.test", "role", "Compute"),
					resource.TestCheckResourceAttr("manta_node.test", "state", "On"),
					resource.TestCheckResourceAttr("manta_node.test", "type", "Node"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "manta_node.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "manta_node" "test" {
  id    = "x1000c0s0b1n1"
  state = "On"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_node.test", "id", "x1000c0s0b1n1"),
					resource.TestCheckResourceAttr("manta_node.test", "state", "On"),
					resource.TestCheckResourceAttr("manta_node.test", "arch", "X86"),
					resource.TestCheckResourceAttr("manta_node.test", "class", "River"),
					resource.TestCheckResourceAttr("manta_node.test", "enabled", "true"),
					resource.TestCheckResourceAttr("manta_node.test", "flag", "OK"),
					resource.TestCheckResourceAttr("manta_node.test", "nettype", "Sling"),
					resource.TestCheckResourceAttr("manta_node.test", "nid", "16400389"),
					resource.TestCheckResourceAttr("manta_node.test", "role", "Compute"),
					resource.TestCheckResourceAttr("manta_node.test", "state", "On"),
					resource.TestCheckResourceAttr("manta_node.test", "type", "Node"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
