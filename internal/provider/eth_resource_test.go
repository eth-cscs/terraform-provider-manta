package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEthResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_eth" "test" {
  mac_address  = "00:de:ad:be:ef:00"
  ip_addresses = ["10.0.0.0"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_eth.test", "id", "00deadbeef00"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_eth" "test" {
  mac_address  = "00:de:ad:be:ef:00"
  ip_addresses = ["10.0.0.0"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_eth.test", "id", "00deadbeef00"),
				),
			},
		},
	})
}

func TestAccEthResourceComponentid(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "manta_eth" "test" {
  mac_address  = "00:de:ad:be:ef:00"
  ip_addresses = ["10.0.0.0"]
  component_id = "x0c0s0b0n0"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_eth.test", "id", "00deadbeef00"),
					resource.TestCheckResourceAttr("manta_eth.test", "mac_address", "00:de:ad:be:ef:00"),
					resource.TestCheckResourceAttr("manta_eth.test", "component_id", "x0c0s0b0n0"),
					resource.TestCheckResourceAttr("manta_eth.test", "type", "Node"),
				),
			},
			{
				Config: providerConfig + `
resource "manta_eth" "test" {
  mac_address  = "00:de:ad:be:ef:00"
  ip_addresses = ["10.0.0.0"]
  component_id = "x0c0s0b0n0"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("manta_eth.test", "id", "00deadbeef00"),
					resource.TestCheckResourceAttr("manta_eth.test", "mac_address", "00:de:ad:be:ef:00"),
					resource.TestCheckResourceAttr("manta_eth.test", "component_id", "x0c0s0b0n0"),
					resource.TestCheckResourceAttr("manta_eth.test", "type", "Node"),
				),
			},
		},
	})
}
