# ibm_is_subnet_routing_table_attachment

Manages ibm_is_subnet_routing_table_attachment. The name of the routing table

This rule validates `ibm_is_subnet_routing_table_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Subnet` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_routing_table_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `is_default` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `route_direct_link_ingress` | no | — |
| `route_transit_gateway_ingress` | no | — |
| `route_vpc_zone_ingress` | no | — |
| `routes` | no | — |
| `routing_table` | no | API: The routing table for this subnet |
| `routing_table_crn` | no | crn format |
| `subnet` | yes | — |
| `subnets` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_routing_table_attachment_invalid_routing_table` | `routing_table` |
| `ibm_is_subnet_routing_table_attachment_invalid_routing_table_crn` | `routing_table_crn` |
| `ibm_is_subnet_routing_table_attachment_invalid_subnet` | `subnet` |

## Example

### Valid

```hcl
resource "ibm_is_subnet_routing_table_attachment" "test" {
  subnet = "example"
}
```
