# ibm_pi_network_data

Reads ibm_pi_network. The name of the network.

This rule validates `ibm_pi_network` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `advertise` | no | — |
| `arp_broadcast` | no | — |
| `available_ip_count` | no | — |
| `cidr` | no | — |
| `crn` | no | — |
| `dns` | no | — |
| `enable_dhcp` | no | — |
| `gateway` | no | — |
| `mtu` | no | — |
| `name` | no | — |
| `network_address_translation` | no | — |
| `peer_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_id` | no | — |
| `pi_network_name` | no | — |
| `type` | no | — |
| `used_ip_count` | no | — |
| `used_ip_percent` | no | — |
| `user_tags` | no | — |
| `vlan_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_data_invalid_pi_network_id` | `pi_network_id` |
| `ibm_pi_network_data_invalid_pi_network_name` | `pi_network_name` |

## Example

### Valid

```hcl
data "ibm_pi_network" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
