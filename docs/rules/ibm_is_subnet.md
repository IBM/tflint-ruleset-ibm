# ibm_is_subnet

Manages ibm_is_subnet. Subnet name

This rule validates `ibm_is_subnet` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Subnet` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `available_ipv4_address_count` | no | — |
| `crn` | no | — |
| `ip_version` | no | one of: `ipv4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `ipv4_cidr_block` | no | cidr format |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_acl` | no | API: The network ACL for this subnet |
| `public_gateway` | no | API: The public gateway to use for internet-bound traffic for this subnet. |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group for this subnet |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `routing_table` | no | API: The routing table for this subnet |
| `routing_table_crn` | no | crn format |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `total_ipv4_address_count` | no | API: The total number of IPv4 addresses in this subnet. Note: This is calculated as 2&lt;sup&gt;(32 - prefix length)&lt;/sup&gt;. For example, the prefix length `/24` gives:&lt;br&gt; 2&lt;sup&gt;(32 - 24)&lt;/sup&gt; = 2&lt;sup&gt;8&lt;/sup&gt; = 256 addresses. |
| `vpc` | yes | API: The VPC this subnet resides in |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_invalid_access_tags` | `access_tags` |
| `ibm_is_subnet_invalid_ip_version` | `ip_version` |
| `ibm_is_subnet_invalid_ipv4_cidr_block` | `ipv4_cidr_block` |
| `ibm_is_subnet_invalid_name` | `name` |
| `ibm_is_subnet_invalid_network_acl` | `network_acl` |
| `ibm_is_subnet_invalid_public_gateway` | `public_gateway` |
| `ibm_is_subnet_invalid_resource_group` | `resource_group` |
| `ibm_is_subnet_invalid_routing_table` | `routing_table` |
| `ibm_is_subnet_invalid_routing_table_crn` | `routing_table_crn` |
| `ibm_is_subnet_invalid_tags` | `tags` |
| `ibm_is_subnet_invalid_total_ipv4_address_count` | `total_ipv4_address_count` |
| `ibm_is_subnet_invalid_vpc` | `vpc` |
| `ibm_is_subnet_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_subnet" "test" {
  name = "example-name"
  vpc = "example"
  zone = "us-south-1"
}
```
