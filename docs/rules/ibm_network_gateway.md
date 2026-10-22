# ibm_network_gateway

Manages ibm_network_gateway. The name of the gateway

This rule validates `ibm_network_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `associated_vlans` | no | — |
| `name` | yes | — |
| `post_install_script_uri` | no | — |
| `private_ip_address_id` | no | — |
| `private_ipv4_address` | no | — |
| `private_vlan_id` | no | — |
| `public_ip_address_id` | no | — |
| `public_ipv4_address` | no | — |
| `public_ipv6_address_id` | no | — |
| `public_vlan_id` | no | — |
| `ssh_key_ids` | no | — |
| `status` | no | — |
| `members` (block) | yes | nested block |
| `members.datacenter` | yes | — |
| `members.disk_key_names` | no | — |
| `members.domain` | yes | — |
| `members.hostname` | no | — |
| `members.ipv6_enabled` | no | — |
| `members.memory` | yes | — |
| `members.network_speed` | no | — |
| `members.notes` | no | — |
| `members.os_key_name` | no | — |
| `members.package_key_name` | no | — |
| `members.post_install_script_uri` | no | — |
| `members.private_network_only` | no | — |
| `members.private_vlan_id` | no | — |
| `members.process_key_name` | no | — |
| `members.public_bandwidth` | no | — |
| `members.public_vlan_id` | no | — |
| `members.redundant_network` | no | — |
| `members.redundant_power_supply` | no | — |
| `members.ssh_key_ids` | no | — |
| `members.tags` | no | — |
| `members.tcp_monitoring` | no | — |
| `members.unbonded_network` | no | — |
| `members.user_metadata` | no | — |
| `members.storage_groups` (block) | no | nested block |
| `members.storage_groups.array_size` | no | — |
| `members.storage_groups.array_type_id` | yes | — |
| `members.storage_groups.hard_drives` | yes | — |
| `members.storage_groups.partition_template_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_gateway_invalid_name` | `name` |
| `ibm_network_gateway_invalid_post_install_script_uri` | `post_install_script_uri` |
| `ibm_network_gateway_invalid_ssh_key_ids` | `ssh_key_ids` |
| `ibm_network_gateway_invalid_members__datacenter` | `members.datacenter` |
| `ibm_network_gateway_invalid_members__disk_key_names` | `members.disk_key_names` |
| `ibm_network_gateway_invalid_members__domain` | `members.domain` |
| `ibm_network_gateway_invalid_members__hostname` | `members.hostname` |
| `ibm_network_gateway_invalid_members__ipv6_enabled` | `members.ipv6_enabled` |
| `ibm_network_gateway_invalid_members__memory` | `members.memory` |
| `ibm_network_gateway_invalid_members__network_speed` | `members.network_speed` |
| `ibm_network_gateway_invalid_members__notes` | `members.notes` |
| `ibm_network_gateway_invalid_members__os_key_name` | `members.os_key_name` |
| `ibm_network_gateway_invalid_members__package_key_name` | `members.package_key_name` |
| `ibm_network_gateway_invalid_members__post_install_script_uri` | `members.post_install_script_uri` |
| `ibm_network_gateway_invalid_members__private_network_only` | `members.private_network_only` |
| `ibm_network_gateway_invalid_members__private_vlan_id` | `members.private_vlan_id` |
| `ibm_network_gateway_invalid_members__process_key_name` | `members.process_key_name` |
| `ibm_network_gateway_invalid_members__public_bandwidth` | `members.public_bandwidth` |
| `ibm_network_gateway_invalid_members__public_vlan_id` | `members.public_vlan_id` |
| `ibm_network_gateway_invalid_members__redundant_network` | `members.redundant_network` |
| `ibm_network_gateway_invalid_members__redundant_power_supply` | `members.redundant_power_supply` |
| `ibm_network_gateway_invalid_members__ssh_key_ids` | `members.ssh_key_ids` |
| `ibm_network_gateway_invalid_members__tags` | `members.tags` |
| `ibm_network_gateway_invalid_members__tcp_monitoring` | `members.tcp_monitoring` |
| `ibm_network_gateway_invalid_members__unbonded_network` | `members.unbonded_network` |
| `ibm_network_gateway_invalid_members__user_metadata` | `members.user_metadata` |
| `ibm_network_gateway_invalid_members__storage_groups__array_size` | `members.storage_groups.array_size` |
| `ibm_network_gateway_invalid_members__storage_groups__array_type_id` | `members.storage_groups.array_type_id` |
| `ibm_network_gateway_invalid_members__storage_groups__hard_drives` | `members.storage_groups.hard_drives` |
| `ibm_network_gateway_invalid_members__storage_groups__partition_template_id` | `members.storage_groups.partition_template_id` |

## Example

### Valid

```hcl
resource "ibm_network_gateway" "test" {
  name = "example"
  members {
    datacenter = "example"
    domain = "example"
    memory = 1
  }
}
```
