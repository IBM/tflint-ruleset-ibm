# ibm_is_vpn_gateway_connection_data

Reads ibm_is_vpn_gateway_connection. The user-defined name for this VPN gateway connection.

This rule validates `ibm_is_vpn_gateway_connection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPNGatewayConnectionDynamicRouteModePeerPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_gateway_connection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `admin_state_up` | no | — |
| `authentication_mode` | no | — |
| `created_at` | no | — |
| `dead_peer_detection` | no | — |
| `distribute_traffic` | no | — |
| `establish_mode` | no | — |
| `href` | no | — |
| `ike_policy` | no | — |
| `ipsec_policy` | no | — |
| `local` | no | — |
| `local_cidrs` | no | — |
| `mode` | no | — |
| `name` | no | — |
| `peer` | no | — |
| `peer_address` | no | — |
| `peer_cidrs` | no | — |
| `psk` | no | — |
| `resource_type` | no | — |
| `routing_protocol` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tunnels` | no | — |
| `vpn_gateway` | no | — |
| `vpn_gateway_connection` | no | — |
| `vpn_gateway_connection_name` | no | — |
| `vpn_gateway_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_connection_data_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_connection_data_invalid_vpn_gateway_connection` | `vpn_gateway_connection` |
| `ibm_is_vpn_gateway_connection_data_invalid_vpn_gateway_connection_name` | `vpn_gateway_connection_name` |
| `ibm_is_vpn_gateway_connection_data_invalid_vpn_gateway_name` | `vpn_gateway_name` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway_connection" "test" {
}
```
