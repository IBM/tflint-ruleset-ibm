# ibm_is_vpc_routing_table_route

Manages ibm_is_vpc_routing_table_route. The user-defined name for this route.

This rule validates `ibm_is_vpc_routing_table_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RoutePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_routing_table_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | one of: `delegate`, `delegate_vpc`, `deliver`, `drop`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `advertise` | no | API: Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property. All routes in a routing table with the same `destination` and `zone` must have the same `advertise` value. |
| `created_at` | no | — |
| `creator` | no | — |
| `destination` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `next_hop` | yes | API: If `action` is `deliver`, the next hop that packets will be delivered to (must not be `0.0.0.0`). For other `action` values, it must be omitted or specified as `0.0.0.0`. At most two routes per `zone` in a table can have the same `destination` and `priority`, and only when each route has an `action` of `deliver` and `next_hop` is an IP address. |
| `origin` | no | — |
| `priority` | no | range 0–4 |
| `route_id` | no | — |
| `routing_table` | yes | — |
| `vpc` | yes | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_routing_table_route_invalid_action` | `action` |
| `ibm_is_vpc_routing_table_route_invalid_advertise` | `advertise` |
| `ibm_is_vpc_routing_table_route_invalid_destination` | `destination` |
| `ibm_is_vpc_routing_table_route_invalid_name` | `name` |
| `ibm_is_vpc_routing_table_route_invalid_next_hop` | `next_hop` |
| `ibm_is_vpc_routing_table_route_invalid_priority` | `priority` |
| `ibm_is_vpc_routing_table_route_invalid_routing_table` | `routing_table` |
| `ibm_is_vpc_routing_table_route_invalid_vpc` | `vpc` |
| `ibm_is_vpc_routing_table_route_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_vpc_routing_table_route" "test" {
  destination = "10.0.0.0/24"
  next_hop = "example"
  routing_table = "example"
  vpc = "example"
  zone = "us-south-1"
}
```
