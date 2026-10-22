# ibm_is_vpn_server_route

Manages ibm_is_vpn_server_route. The user-defined name for this VPN route.

This rule validates `ibm_is_vpn_server_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPNServerRoutePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_server_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | one of: `deliver`, `drop`, `translate`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `created_at` | no | — |
| `destination` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_type` | no | — |
| `vpn_route` | no | — |
| `vpn_server` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_route_invalid_action` | `action` |
| `ibm_is_vpn_server_route_invalid_destination` | `destination` |
| `ibm_is_vpn_server_route_invalid_name` | `name` |
| `ibm_is_vpn_server_route_invalid_vpn_server` | `vpn_server` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_server_route" "test" {
  destination = "10.0.0.0/24"
  vpn_server = "example"
}
```
