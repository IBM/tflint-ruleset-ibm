# ibm_is_vpc_routing_table_route_data

Reads ibm_is_vpc_routing_table_route. The user-defined name for this route.

This rule validates `ibm_is_vpc_routing_table_route` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCDNSResolutionBindingPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_routing_table_route_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `advertise` | no | — |
| `created_at` | no | — |
| `creator` | no | — |
| `destination` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `next_hop` | no | — |
| `origin` | no | — |
| `priority` | no | — |
| `route_id` | no | — |
| `routing_table` | yes | — |
| `vpc` | yes | API: The VPC to bind this VPC to for DNS resolution. The VPC must be different from the VPC specified in the URL, must have `dns.enable_hub` set to `true`, and may be in a different account (subject to IAM policies). Additionally, the VPC specified in the URL (this VPC) must have `dns.enable_hub` set to `false` and a `dns.resolution_binding_count` of zero. |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_routing_table_route_data_invalid_name` | `name` |
| `ibm_is_vpc_routing_table_route_data_invalid_route_id` | `route_id` |
| `ibm_is_vpc_routing_table_route_data_invalid_routing_table` | `routing_table` |
| `ibm_is_vpc_routing_table_route_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_vpc_routing_table_route" "test" {
  routing_table = "example"
  vpc = "example"
}
```
