/*
Copyright 2021 Upbound Inc.
*/

package ec2

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ec2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_instance", func(r *config.Resource) {
		r.UseAsync = true
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "VPCSecurityGroupIDRefs",
			SelectorFieldName: "VPCSecurityGroupIDSelector",
		}
		r.References["security_groups"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["root_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["network_interface.network_interface_id"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.References["ebs_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/upbound/upjet/issues/107
			IgnoredFields: []string{
				"subnet_id",
				"network_interface",
				"private_ip",
				"source_dest_check",
				"vpc_security_group_ids",
				"associate_public_ip_address",
			},
		}
		config.MoveToStatus(r.TerraformResource, "security_groups")
	})

	p.AddResourceConfigurator("aws_ec2_tag", func(r *config.Resource) {
		delete(r.References, "resource_id")
	})

	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["gateway_id"] = config.Reference{
			Type: "InternetGateway",
		}
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_interface_id"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
		r.References["vpc_peering_connection_id"] = config.Reference{
			Type: "VPCPeeringConnection",
		}
		r.References["vpc_endpoint_id"] = config.Reference{
			Type: "VPCEndpoint",
		}
		r.References["nat_gateway_id"] = config.Reference{
			Type: "NATGateway",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_security_group", func(r *config.Resource) {
		// Mutually exclusive with aws_security_group_rule
		config.MoveToStatus(r.TerraformResource, "ingress", "egress")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"name", "name_prefix",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_security_group_rule", func(r *config.Resource) {
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["source_security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"cidr_blocks",
				"ipv6_cidr_blocks",
				"self",
				"source_security_group_id",
			},
		}
	})

	p.AddResourceConfigurator("aws_nat_gateway", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("aws_vpc", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"ipv6_cidr_block",
			},
		}
		r.UseAsync = true
	})

}
