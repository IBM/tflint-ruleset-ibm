# ibm_is_vpc_routing_table

Manages ibm_is_vpc_routing_table. The user-defined name for this routing table.

This rule validates `ibm_is_vpc_routing_table` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RoutingTablePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_routing_table" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_routes_from_resource_type` | no | — |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `advertise_routes_to` | no | one of: `direct_link`, `transit_gateway`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `is_default` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `route_direct_link_ingress` | no | API: Indicates whether this routing table is used to route traffic that originates from [Direct Link](https://cloud.ibm.com/docs/dl/) to this VPC. Updating to `true` selects this routing table, provided no other routing table in the VPC already has this property set to `true`, and no subnets are attached to this routing table. Updating to `false` deselects this routing table, provided `direct_link` is  |
| `route_internet_ingress` | no | API: Indicates whether this routing table is used to route traffic that originates from the internet. Updating to `true` selects this routing table, provided no other routing table in the VPC already has this property set to `true`. Updating to `false` deselects this routing table. Incoming traffic will be routed according to the routing table with two exceptions: - Traffic destined for IP addresses as |
| `route_transit_gateway_ingress` | no | API: Indicates whether this routing table is used to route traffic that originates from [Transit Gateway](https://cloud.ibm.com/docs/transit-gateway) to this VPC. Updating to `true` selects this routing table, provided no other routing table in the VPC already has this property set to `true`, and no subnets are attached to this routing table. Updating to `false` deselects this routing table, provided ` |
| `route_vpc_zone_ingress` | no | API: Indicates whether this routing table is used to route traffic that originates from subnets in other zones in this VPC. Updating to `true` selects this routing table, provided no other routing table in the VPC already has this property set to `true`, and no subnets are attached to this routing table. Updating to `false` deselects this routing table. Incoming traffic will be routed according to the  |
| `routing_table` | no | — |
| `subnets` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_routing_table_invalid_accept_routes_from_resource_type` | `accept_routes_from_resource_type` |
| `ibm_is_vpc_routing_table_invalid_access_tags` | `access_tags` |
| `ibm_is_vpc_routing_table_invalid_advertise_routes_to` | `advertise_routes_to` |
| `ibm_is_vpc_routing_table_invalid_name` | `name` |
| `ibm_is_vpc_routing_table_invalid_route_direct_link_ingress` | `route_direct_link_ingress` |
| `ibm_is_vpc_routing_table_invalid_route_internet_ingress` | `route_internet_ingress` |
| `ibm_is_vpc_routing_table_invalid_route_transit_gateway_ingress` | `route_transit_gateway_ingress` |
| `ibm_is_vpc_routing_table_invalid_route_vpc_zone_ingress` | `route_vpc_zone_ingress` |
| `ibm_is_vpc_routing_table_invalid_tags` | `tags` |
| `ibm_is_vpc_routing_table_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
resource "ibm_is_vpc_routing_table" "test" {
  vpc = "example"
}
```
