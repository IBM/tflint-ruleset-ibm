# ibm_pi_network

Schema validation for the ibm_pi_network resource.

This rule validates `ibm_pi_network` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Network` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_network" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `enable_dhcp` | no | — |
| `network_address_translation` | no | — |
| `network_id` | no | — |
| `peer_id` | no | — |
| `pi_advertise` | no | one of: `enable`, `disable` |
| `pi_arp_broadcast` | no | one of: `enable`, `disable` |
| `pi_cidr` | no | cidr format |
| `pi_cloud_instance_id` | yes | — |
| `pi_dns` | no | — |
| `pi_enable_dhcp` | no | — |
| `pi_gateway` | no | — |
| `pi_network_mtu` | no | — |
| `pi_network_name` | yes | — |
| `pi_network_type` | yes | one of: `pub-vlan`, `vlan` |
| `pi_user_tags` | no | — |
| `vlan_id` | no | — |
| `pi_ipaddress_range` (block) | no | nested block |
| `pi_ipaddress_range.pi_ending_ip_address` | yes | — |
| `pi_ipaddress_range.pi_starting_ip_address` | yes | — |
| `pi_network_peer` (block) | no | nested block |
| `pi_network_peer.id` | yes | API: Unique identifier for the trusted profile |
| `pi_network_peer.type` | no | one of: `L2`, `L3BGP`, `L3Static` |
| `pi_network_peer.network_address_translation` (block) | no | nested block |
| `pi_network_peer.network_address_translation.source_ip` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_invalid_pi_advertise` | `pi_advertise` |
| `ibm_pi_network_invalid_pi_arp_broadcast` | `pi_arp_broadcast` |
| `ibm_pi_network_invalid_pi_cidr` | `pi_cidr` |
| `ibm_pi_network_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_invalid_pi_dns` | `pi_dns` |
| `ibm_pi_network_invalid_pi_enable_dhcp` | `pi_enable_dhcp` |
| `ibm_pi_network_invalid_pi_gateway` | `pi_gateway` |
| `ibm_pi_network_invalid_pi_network_mtu` | `pi_network_mtu` |
| `ibm_pi_network_invalid_pi_network_name` | `pi_network_name` |
| `ibm_pi_network_invalid_pi_network_type` | `pi_network_type` |
| `ibm_pi_network_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_network_invalid_pi_ipaddress_range__pi_ending_ip_address` | `pi_ipaddress_range.pi_ending_ip_address` |
| `ibm_pi_network_invalid_pi_ipaddress_range__pi_starting_ip_address` | `pi_ipaddress_range.pi_starting_ip_address` |
| `ibm_pi_network_invalid_pi_network_peer__id` | `pi_network_peer.id` |
| `ibm_pi_network_invalid_pi_network_peer__type` | `pi_network_peer.type` |
| `ibm_pi_network_invalid_pi_network_peer__network_address_translation__source_ip` | `pi_network_peer.network_address_translation.source_ip` |

## Example

### Valid

```hcl
resource "ibm_pi_network" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_name = "example"
  pi_network_type = "pub-vlan"
}
```
