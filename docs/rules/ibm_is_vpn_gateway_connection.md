# ibm_is_vpn_gateway_connection

Manages ibm_is_vpn_gateway_connection. VPN Gateway connection name

This rule validates `ibm_is_vpn_gateway_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPNGatewayConnectionPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_gateway_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | one of: `restart`, `clear`, `hold`, `none` |
| `admin_state_up` | no | API: If set to false, the VPN gateway connection is shut down |
| `authentication_mode` | no | — |
| `created_at` | no | — |
| `distribute_traffic` | no | API: Indicates whether the traffic is distributed between the `up` tunnels of the VPN gateway connection when the VPC route's next hop is a VPN connection. If `false`, the traffic is only routed through the `up` tunnel with the lower `public_ip` address. Before enabling it on VPN connections to on-prem private networks, review [distributing traffic restrictions](https://cloud.ibm.com/docs/vpc?topic=vpc |
| `establish_mode` | no | one of: `bidirectional`, `peer_only`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `gateway_connection` | no | — |
| `href` | no | — |
| `ike_policy` | no | API: The IKE policy to use. Specify `null` to remove any existing policy, [resulting in auto-negotiation](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#ike-auto-negotiation-phase-1) |
| `interval` | no | range 1–86399 |
| `ipsec_policy` | no | API: The IPsec policy to use. Specify `null` to remove any existing policy, [resulting in auto-negotiation](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#ipsec-auto-negotiation-phase-2) |
| `local_cidrs` | no | — |
| `mode` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `peer_address` | no | — |
| `peer_cidrs` | no | — |
| `preshared_key` | yes | — |
| `related_crn` | no | — |
| `resource_type` | no | — |
| `routing_protocol` | no | one of: `bgp`, `none`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `status` | no | — |
| `status_reasons` | no | — |
| `timeout` | no | range 2–86399 |
| `tunnels` | no | — |
| `vpn_gateway` | yes | — |
| `local` (block) | no | nested block |
| `local.cidrs` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `local.ike_identities` (block) | no | nested block |
| `local.ike_identities.type` | yes | — |
| `local.ike_identities.value` | no | — |
| `peer` (block) | no | nested block |
| `peer.address` | no | API: The IP address of the peer VPN gateway for this connection. |
| `peer.asn` | no | range 64512–65534 |
| `peer.cidrs` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `peer.fqdn` | no | API: The FQDN of the peer VPN gateway for this connection. |
| `peer.ike_identity` (block) | no | nested block |
| `peer.ike_identity.type` | yes | — |
| `peer.ike_identity.value` | no | — |
| `tunnel` (block) | no | nested block |
| `tunnel.neighbor_ip` | yes | API: The IP address of the neighbor on the virtual tunnel interface. This serves as the destination address for BGP peering sessions on the peer gateway within the tunnel. |
| `tunnel.tunnel_interface_ip` | yes | API: The IP address assigned to the VPN gateway's virtual tunnel interface. This serves as the source address for BGP peering sessions initiated from the VPN gateway towards the peer gateway within the tunnel. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_connection_invalid_action` | `action` |
| `ibm_is_vpn_gateway_connection_invalid_admin_state_up` | `admin_state_up` |
| `ibm_is_vpn_gateway_connection_invalid_distribute_traffic` | `distribute_traffic` |
| `ibm_is_vpn_gateway_connection_invalid_establish_mode` | `establish_mode` |
| `ibm_is_vpn_gateway_connection_invalid_ike_policy` | `ike_policy` |
| `ibm_is_vpn_gateway_connection_invalid_interval` | `interval` |
| `ibm_is_vpn_gateway_connection_invalid_ipsec_policy` | `ipsec_policy` |
| `ibm_is_vpn_gateway_connection_invalid_local_cidrs` | `local_cidrs` |
| `ibm_is_vpn_gateway_connection_invalid_name` | `name` |
| `ibm_is_vpn_gateway_connection_invalid_peer_address` | `peer_address` |
| `ibm_is_vpn_gateway_connection_invalid_peer_cidrs` | `peer_cidrs` |
| `ibm_is_vpn_gateway_connection_invalid_preshared_key` | `preshared_key` |
| `ibm_is_vpn_gateway_connection_invalid_routing_protocol` | `routing_protocol` |
| `ibm_is_vpn_gateway_connection_invalid_timeout` | `timeout` |
| `ibm_is_vpn_gateway_connection_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_connection_invalid_local__cidrs` | `local.cidrs` |
| `ibm_is_vpn_gateway_connection_invalid_local__ike_identities__type` | `local.ike_identities.type` |
| `ibm_is_vpn_gateway_connection_invalid_local__ike_identities__value` | `local.ike_identities.value` |
| `ibm_is_vpn_gateway_connection_invalid_peer__address` | `peer.address` |
| `ibm_is_vpn_gateway_connection_invalid_peer__asn` | `peer.asn` |
| `ibm_is_vpn_gateway_connection_invalid_peer__cidrs` | `peer.cidrs` |
| `ibm_is_vpn_gateway_connection_invalid_peer__fqdn` | `peer.fqdn` |
| `ibm_is_vpn_gateway_connection_invalid_peer__ike_identity__type` | `peer.ike_identity.type` |
| `ibm_is_vpn_gateway_connection_invalid_peer__ike_identity__value` | `peer.ike_identity.value` |
| `ibm_is_vpn_gateway_connection_invalid_tunnel__neighbor_ip` | `tunnel.neighbor_ip` |
| `ibm_is_vpn_gateway_connection_invalid_tunnel__tunnel_interface_ip` | `tunnel.tunnel_interface_ip` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_gateway_connection" "test" {
  name = "example-name"
  preshared_key = "example"
  vpn_gateway = "example"
}
```
