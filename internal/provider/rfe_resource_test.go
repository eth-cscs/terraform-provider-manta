package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRfeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_rfe" "test" {
  id                 = "x1002c0s0b10"
  user               = "user"
  hostname           = "hostname"
  rediscoveronupdate = true
  enabled            = true
  password_wo        = "password"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_rfe.test", "id", "x1002c0s0b10"),
					resource.TestCheckResourceAttr("manta_rfe.test", "user", "user"),
					resource.TestCheckResourceAttr("manta_rfe.test", "hostname", "hostname"),
					resource.TestCheckResourceAttr("manta_rfe.test", "rediscoveronupdate", "true"),
					resource.TestCheckResourceAttr("manta_rfe.test", "enabled", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "manta_rfe.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "manta_rfe" "test" {
  id                 = "x1002c0s0b10"
  user               = "user"
  hostname           = "hostname"
  rediscoveronupdate = true
  enabled            = true
  password_wo        = "password"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_rfe.test", "id", "x1002c0s0b10"),
					resource.TestCheckResourceAttr("manta_rfe.test", "user", "user"),
					resource.TestCheckResourceAttr("manta_rfe.test", "hostname", "hostname"),
					resource.TestCheckResourceAttr("manta_rfe.test", "rediscoveronupdate", "true"),
					resource.TestCheckResourceAttr("manta_rfe.test", "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
