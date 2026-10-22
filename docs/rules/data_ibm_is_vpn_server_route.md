# ibm_is_vpn_server_route_data

Reads ibm_is_vpn_server_route. The unique user-defined name for this VPN server route

This rule validates `ibm_is_vpn_server_route` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_server_route_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `created_at` | no | — |
| `destination` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_type` | no | — |
| `vpn_server` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_route_data_invalid_identifier` | `identifier` |
| `ibm_is_vpn_server_route_data_invalid_name` | `name` |
| `ibm_is_vpn_server_route_data_invalid_vpn_server` | `vpn_server` |

## Example

### Valid

```hcl
data "ibm_is_vpn_server_route" "test" {
  vpn_server = "example"
}
```
