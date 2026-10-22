# ibm_is_vpc

Manages ibm_is_vpc. VPC name

This rule validates `ibm_is_vpc` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPC` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `address_prefix_management` | no | one of: `auto`, `manual`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `classic_access` | no | API: Indicates whether this VPC is connected to Classic Infrastructure. If true, this VPC's resources have private network connectivity to the account's Classic Infrastructure resources. |
| `crn` | no | — |
| `cse_source_addresses` | no | — |
| `default_address_prefixes` | no | — |
| `default_network_acl` | no | — |
| `default_network_acl_crn` | no | — |
| `default_network_acl_name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `default_routing_table` | no | — |
| `default_routing_table_crn` | no | — |
| `default_routing_table_name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `default_security_group` | no | — |
| `default_security_group_crn` | no | — |
| `default_security_group_name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `no_sg_acl_rules` | no | — |
| `public_address_ranges` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group for this VPC |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `security_group` | no | — |
| `status` | no | — |
| `subnets` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `dns` (block) | no | nested block |
| `dns.enable_hub` | no | API: Indicates whether this VPC is enabled as a DNS name resolution hub. |
| `dns.resolver` (block) | no | nested block |
| `dns.resolver.dns_binding_name` | no | — |
| `dns.resolver.type` | no | one of: `delegated`, `manual`, `system`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `dns.resolver.vpc_crn` | no | crn format |
| `dns.resolver.vpc_id` | no | — |
| `dns.resolver.manual_servers` (block) | no | nested block |
| `dns.resolver.manual_servers.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `dns.resolver.manual_servers.zone_affinity` | no | API: If present, DHCP configuration for this zone will have this DNS server listed first. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_invalid_access_tags` | `access_tags` |
| `ibm_is_vpc_invalid_address_prefix_management` | `address_prefix_management` |
| `ibm_is_vpc_invalid_classic_access` | `classic_access` |
| `ibm_is_vpc_invalid_default_network_acl_name` | `default_network_acl_name` |
| `ibm_is_vpc_invalid_default_routing_table_name` | `default_routing_table_name` |
| `ibm_is_vpc_invalid_default_security_group_name` | `default_security_group_name` |
| `ibm_is_vpc_invalid_name` | `name` |
| `ibm_is_vpc_invalid_no_sg_acl_rules` | `no_sg_acl_rules` |
| `ibm_is_vpc_invalid_resource_group` | `resource_group` |
| `ibm_is_vpc_invalid_tags` | `tags` |
| `ibm_is_vpc_invalid_dns__enable_hub` | `dns.enable_hub` |
| `ibm_is_vpc_invalid_dns__resolver__dns_binding_name` | `dns.resolver.dns_binding_name` |
| `ibm_is_vpc_invalid_dns__resolver__type` | `dns.resolver.type` |
| `ibm_is_vpc_invalid_dns__resolver__vpc_crn` | `dns.resolver.vpc_crn` |
| `ibm_is_vpc_invalid_dns__resolver__vpc_id` | `dns.resolver.vpc_id` |
| `ibm_is_vpc_invalid_dns__resolver__manual_servers__address` | `dns.resolver.manual_servers.address` |
| `ibm_is_vpc_invalid_dns__resolver__manual_servers__zone_affinity` | `dns.resolver.manual_servers.zone_affinity` |

## Example

### Valid

```hcl
resource "ibm_is_vpc" "test" {
  name = "example-name"
}
```
