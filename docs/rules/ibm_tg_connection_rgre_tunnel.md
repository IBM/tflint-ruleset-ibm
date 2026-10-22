# ibm_tg_connection_rgre_tunnel

Manages ibm_tg_connection_rgre_tunnel. The user-defined name for this tunnel connection.

This rule validates `ibm_tg_connection_rgre_tunnel` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TransitGatewayTunnelTemplate` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_connection_rgre_tunnel" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `base_network_type` | no | one of: `classic`, `vpc` |
| `connection_id` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `created_at` | no | — |
| `gateway` | yes | — |
| `local_bgp_asn` | no | range 1–4294967295 |
| `local_gateway_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `local_tunnel_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `mtu` | no | — |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `network_account_id` | no | API: The ID of the account which owns the network that is being connected. Generally only used if the network is in a different account than the gateway. This field is required for type `unbound_gre_tunnel` when the associated_network_type is `classic` or network_type is `redundant_gre` and the GRE tunnel is in a different account than the gateway. |
| `network_id` | no | API: The ID of the network being connected via this connection. For network types `vpc`, `vpn_gateway`, `power_virtual_server` and `directlink` this is the CRN of the VPC / VPN / PowerVS / Direct Link gateway respectively. This field is required for network type `vpc`, `power_virtual_server`, `vpn_gateway` and `directlink` connections. It is also required for `redundant_gre` connections when the base_n |
| `remote_bgp_asn` | no | range 1–4294967295 |
| `remote_gateway_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `remote_tunnel_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `status` | no | — |
| `tunnel_id` | no | — |
| `updated_at` | no | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_connection_rgre_tunnel_invalid_base_network_type` | `base_network_type` |
| `ibm_tg_connection_rgre_tunnel_invalid_connection_id` | `connection_id` |
| `ibm_tg_connection_rgre_tunnel_invalid_gateway` | `gateway` |
| `ibm_tg_connection_rgre_tunnel_invalid_local_bgp_asn` | `local_bgp_asn` |
| `ibm_tg_connection_rgre_tunnel_invalid_local_gateway_ip` | `local_gateway_ip` |
| `ibm_tg_connection_rgre_tunnel_invalid_local_tunnel_ip` | `local_tunnel_ip` |
| `ibm_tg_connection_rgre_tunnel_invalid_name` | `name` |
| `ibm_tg_connection_rgre_tunnel_invalid_network_account_id` | `network_account_id` |
| `ibm_tg_connection_rgre_tunnel_invalid_network_id` | `network_id` |
| `ibm_tg_connection_rgre_tunnel_invalid_remote_bgp_asn` | `remote_bgp_asn` |
| `ibm_tg_connection_rgre_tunnel_invalid_remote_gateway_ip` | `remote_gateway_ip` |
| `ibm_tg_connection_rgre_tunnel_invalid_remote_tunnel_ip` | `remote_tunnel_ip` |
| `ibm_tg_connection_rgre_tunnel_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_tg_connection_rgre_tunnel" "test" {
  connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
  local_gateway_ip = "10.0.0.1"
  local_tunnel_ip = "10.0.0.1"
  name = "example-name"
  remote_gateway_ip = "10.0.0.1"
  remote_tunnel_ip = "10.0.0.1"
  zone = "us-south-1"
}
```
