# ibm_tg_connection

Manages ibm_tg_connection. The user-defined name for this transit gateway.

This rule validates `ibm_tg_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TransitGatewayConnectionTemplate` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `base_connection_id` | no | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `base_network_type` | no | one of: `classic`, `vpc` |
| `cidr` | no | cidr format |
| `connection_id` | no | — |
| `created_at` | no | — |
| `default_prefix_filter` | no | — |
| `gateway` | yes | — |
| `local_gateway_ip` | no | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `local_tunnel_ip` | no | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `name` | no | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `network_account_id` | no | API: The ID of the account which owns the network that is being connected. Generally only used if the network is in a different account than the gateway. This field is required for type `unbound_gre_tunnel` when the associated_network_type is `classic` or network_type is `redundant_gre` and the GRE tunnel is in a different account than the gateway. |
| `network_id` | no | API: The ID of the network being connected via this connection. For network types `vpc`, `vpn_gateway`, `power_virtual_server` and `directlink` this is the CRN of the VPC / VPN / PowerVS / Direct Link gateway respectively. This field is required for network type `vpc`, `power_virtual_server`, `vpn_gateway` and `directlink` connections. It is also required for `redundant_gre` connections when the base_n |
| `network_type` | yes | one of: `classic`, `directlink`, `vpc`, `gre_tunnel`, `unbound_gre_tunnel`, `power_virtual_server`, `redundant_gre`, `vpn_gateway` |
| `related_crn` | no | — |
| `remote_bgp_asn` | no | range 1–4294967295 |
| `remote_gateway_ip` | no | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `remote_tunnel_ip` | no | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `request_status` | no | — |
| `status` | no | — |
| `updated_at` | no | — |
| `zone` | no | zone format |
| `tunnels` (block) | no | nested block |
| `tunnels.local_bgp_asn` | no | range 1–4294967295 |
| `tunnels.local_gateway_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `tunnels.local_tunnel_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `tunnels.name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `tunnels.remote_bgp_asn` | no | range 1–4294967295 |
| `tunnels.remote_gateway_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `tunnels.remote_tunnel_ip` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})$` |
| `tunnels.zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_connection_invalid_base_connection_id` | `base_connection_id` |
| `ibm_tg_connection_invalid_base_network_type` | `base_network_type` |
| `ibm_tg_connection_invalid_cidr` | `cidr` |
| `ibm_tg_connection_invalid_default_prefix_filter` | `default_prefix_filter` |
| `ibm_tg_connection_invalid_gateway` | `gateway` |
| `ibm_tg_connection_invalid_local_gateway_ip` | `local_gateway_ip` |
| `ibm_tg_connection_invalid_local_tunnel_ip` | `local_tunnel_ip` |
| `ibm_tg_connection_invalid_name` | `name` |
| `ibm_tg_connection_invalid_network_account_id` | `network_account_id` |
| `ibm_tg_connection_invalid_network_id` | `network_id` |
| `ibm_tg_connection_invalid_network_type` | `network_type` |
| `ibm_tg_connection_invalid_remote_bgp_asn` | `remote_bgp_asn` |
| `ibm_tg_connection_invalid_remote_gateway_ip` | `remote_gateway_ip` |
| `ibm_tg_connection_invalid_remote_tunnel_ip` | `remote_tunnel_ip` |
| `ibm_tg_connection_invalid_zone` | `zone` |
| `ibm_tg_connection_invalid_tunnels__local_bgp_asn` | `tunnels.local_bgp_asn` |
| `ibm_tg_connection_invalid_tunnels__local_gateway_ip` | `tunnels.local_gateway_ip` |
| `ibm_tg_connection_invalid_tunnels__local_tunnel_ip` | `tunnels.local_tunnel_ip` |
| `ibm_tg_connection_invalid_tunnels__name` | `tunnels.name` |
| `ibm_tg_connection_invalid_tunnels__remote_bgp_asn` | `tunnels.remote_bgp_asn` |
| `ibm_tg_connection_invalid_tunnels__remote_gateway_ip` | `tunnels.remote_gateway_ip` |
| `ibm_tg_connection_invalid_tunnels__remote_tunnel_ip` | `tunnels.remote_tunnel_ip` |
| `ibm_tg_connection_invalid_tunnels__zone` | `tunnels.zone` |

## Example

### Valid

```hcl
resource "ibm_tg_connection" "test" {
  gateway = "example"
  network_type = "classic"
}
```
