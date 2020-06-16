package aws

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwsEc2LocalGatewayVirtualInterface_Filter(t *testing.T) {
	// Hide Outposts testing behind consistent environment variable
	outpostArn := os.Getenv("AWS_OUTPOST_ARN")
	if outpostArn == "" {
		t.Skip(
			"Environment variable AWS_OUTPOST_ARN is not set. " +
				"This environment variable must be set to the ARN of " +
				"a deployed Outpost to enable this test.")
	}

	dataSourceName := "data.aws_ec2_local_gateway_virtual_interface.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsEc2LocalGatewayVirtualInterfaceConfigFilter(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceName, "id", regexp.MustCompile(`^lgw-vif-`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_address", regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+/\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_bgp_asn", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_gateway_id", regexp.MustCompile(`^lgw-`)),
					resource.TestMatchResourceAttr(dataSourceName, "peer_address", regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+/\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "peer_bgp_asn", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestMatchResourceAttr(dataSourceName, "vlan", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

func TestAccDataSourceAwsEc2LocalGatewayVirtualInterface_Id(t *testing.T) {
	// Hide Outposts testing behind consistent environment variable
	outpostArn := os.Getenv("AWS_OUTPOST_ARN")
	if outpostArn == "" {
		t.Skip(
			"Environment variable AWS_OUTPOST_ARN is not set. " +
				"This environment variable must be set to the ARN of " +
				"a deployed Outpost to enable this test.")
	}

	dataSourceName := "data.aws_ec2_local_gateway_virtual_interface.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsEc2LocalGatewayVirtualInterfaceConfigId(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceName, "id", regexp.MustCompile(`^lgw-vif-`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_address", regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+/\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_bgp_asn", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "local_gateway_id", regexp.MustCompile(`^lgw-`)),
					resource.TestMatchResourceAttr(dataSourceName, "peer_address", regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+/\d+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "peer_bgp_asn", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestMatchResourceAttr(dataSourceName, "vlan", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

func testAccDataSourceAwsEc2LocalGatewayVirtualInterfaceConfigFilter() string {
	return `
data "aws_ec2_local_gateways" "test" {}

data "aws_ec2_local_gateway_virtual_interface_group" "test" {
  local_gateway_id = tolist(data.aws_ec2_local_gateways.test.ids)[0]
}

data "aws_ec2_local_gateway_virtual_interface" "test" {
  filter {
    name   = "local-gateway-virtual-interface-id"
    values = [tolist(data.aws_ec2_local_gateway_virtual_interface_group.test.local_gateway_virtual_interface_ids)[0]]
  }
}
`
}

func testAccDataSourceAwsEc2LocalGatewayVirtualInterfaceConfigId() string {
	return `
data "aws_ec2_local_gateways" "test" {}

data "aws_ec2_local_gateway_virtual_interface_group" "test" {
  local_gateway_id = tolist(data.aws_ec2_local_gateways.test.ids)[0]
}

data "aws_ec2_local_gateway_virtual_interface" "test" {
  id = tolist(data.aws_ec2_local_gateway_virtual_interface_group.test.local_gateway_virtual_interface_ids)[0]
}
`
}
