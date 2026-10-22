# ibm_is_vpc_routing_table_data

Reads ibm_is_vpc_routing_table. The user-defined name for this routing table.

This rule validates `ibm_is_vpc_routing_table` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCDNSResolutionBindingPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_routing_table_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_routes_from` | no | — |
| `access_tags` | no | — |
| `advertise_routes_to` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `is_default` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `route_direct_link_ingress` | no | — |
| `route_internet_ingress` | no | — |
| `route_transit_gateway_ingress` | no | — |
| `route_vpc_zone_ingress` | no | — |
| `routes` | no | — |
| `routing_table` | no | — |
| `subnets` | no | — |
| `tags` | no | — |
| `vpc` | yes | API: The VPC to bind this VPC to for DNS resolution. The VPC must be different from the VPC specified in the URL, must have `dns.enable_hub` set to `true`, and may be in a different account (subject to IAM policies). Additionally, the VPC specified in the URL (this VPC) must have `dns.enable_hub` set to `false` and a `dns.resolution_binding_count` of zero. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_routing_table_data_invalid_name` | `name` |
| `ibm_is_vpc_routing_table_data_invalid_routing_table` | `routing_table` |
| `ibm_is_vpc_routing_table_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_vpc_routing_table" "test" {
  vpc = "example"
}
```
