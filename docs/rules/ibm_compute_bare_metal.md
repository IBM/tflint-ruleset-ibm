# ibm_compute_bare_metal

Schema validation for the ibm_compute_bare_metal resource.

This rule validates `ibm_compute_bare_metal` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_bare_metal" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `block_storage_ids` | no | — |
| `datacenter` | no | — |
| `disk_key_names` | no | — |
| `domain` | yes | — |
| `extended_hardware_testing` | no | — |
| `file_storage_ids` | no | — |
| `fixed_config_preset` | no | — |
| `global_identifier` | no | — |
| `gpu_key_name` | no | — |
| `gpu_secondary_key_name` | no | — |
| `hostname` | no | — |
| `hourly_billing` | no | — |
| `image_template_id` | no | — |
| `ipv6_address` | no | — |
| `ipv6_address_id` | no | — |
| `ipv6_enabled` | no | — |
| `ipv6_static_enabled` | no | — |
| `memory` | no | — |
| `network_speed` | no | — |
| `notes` | no | — |
| `os_key_name` | no | — |
| `os_reference_code` | no | — |
| `package_key_name` | no | — |
| `post_install_script_uri` | no | — |
| `private_ipv4_address` | no | — |
| `private_ipv4_address_id` | no | — |
| `private_network_only` | no | — |
| `private_subnet` | no | — |
| `private_vlan_id` | no | — |
| `process_key_name` | no | — |
| `public_bandwidth` | no | — |
| `public_ipv4_address` | no | — |
| `public_ipv4_address_id` | no | — |
| `public_subnet` | no | — |
| `public_vlan_id` | no | — |
| `quote_id` | no | — |
| `redundant_network` | no | — |
| `redundant_power_supply` | no | — |
| `restricted_network` | no | — |
| `secondary_ip_addresses` | no | — |
| `secondary_ip_count` | no | — |
| `software_guard_extensions` | no | — |
| `ssh_key_ids` | no | — |
| `tags` | no | — |
| `tcp_monitoring` | no | — |
| `unbonded_network` | no | — |
| `user_metadata` | no | — |
| `storage_groups` (block) | no | nested block |
| `storage_groups.array_size` | no | — |
| `storage_groups.array_type_id` | yes | — |
| `storage_groups.hard_drives` | yes | — |
| `storage_groups.partition_template_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_bare_metal_invalid_block_storage_ids` | `block_storage_ids` |
| `ibm_compute_bare_metal_invalid_datacenter` | `datacenter` |
| `ibm_compute_bare_metal_invalid_disk_key_names` | `disk_key_names` |
| `ibm_compute_bare_metal_invalid_domain` | `domain` |
| `ibm_compute_bare_metal_invalid_extended_hardware_testing` | `extended_hardware_testing` |
| `ibm_compute_bare_metal_invalid_file_storage_ids` | `file_storage_ids` |
| `ibm_compute_bare_metal_invalid_fixed_config_preset` | `fixed_config_preset` |
| `ibm_compute_bare_metal_invalid_gpu_key_name` | `gpu_key_name` |
| `ibm_compute_bare_metal_invalid_gpu_secondary_key_name` | `gpu_secondary_key_name` |
| `ibm_compute_bare_metal_invalid_hostname` | `hostname` |
| `ibm_compute_bare_metal_invalid_hourly_billing` | `hourly_billing` |
| `ibm_compute_bare_metal_invalid_image_template_id` | `image_template_id` |
| `ibm_compute_bare_metal_invalid_ipv6_enabled` | `ipv6_enabled` |
| `ibm_compute_bare_metal_invalid_ipv6_static_enabled` | `ipv6_static_enabled` |
| `ibm_compute_bare_metal_invalid_memory` | `memory` |
| `ibm_compute_bare_metal_invalid_network_speed` | `network_speed` |
| `ibm_compute_bare_metal_invalid_notes` | `notes` |
| `ibm_compute_bare_metal_invalid_os_key_name` | `os_key_name` |
| `ibm_compute_bare_metal_invalid_os_reference_code` | `os_reference_code` |
| `ibm_compute_bare_metal_invalid_package_key_name` | `package_key_name` |
| `ibm_compute_bare_metal_invalid_post_install_script_uri` | `post_install_script_uri` |
| `ibm_compute_bare_metal_invalid_private_network_only` | `private_network_only` |
| `ibm_compute_bare_metal_invalid_private_subnet` | `private_subnet` |
| `ibm_compute_bare_metal_invalid_private_vlan_id` | `private_vlan_id` |
| `ibm_compute_bare_metal_invalid_process_key_name` | `process_key_name` |
| `ibm_compute_bare_metal_invalid_public_bandwidth` | `public_bandwidth` |
| `ibm_compute_bare_metal_invalid_public_subnet` | `public_subnet` |
| `ibm_compute_bare_metal_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_compute_bare_metal_invalid_quote_id` | `quote_id` |
| `ibm_compute_bare_metal_invalid_redundant_network` | `redundant_network` |
| `ibm_compute_bare_metal_invalid_redundant_power_supply` | `redundant_power_supply` |
| `ibm_compute_bare_metal_invalid_restricted_network` | `restricted_network` |
| `ibm_compute_bare_metal_invalid_secondary_ip_count` | `secondary_ip_count` |
| `ibm_compute_bare_metal_invalid_software_guard_extensions` | `software_guard_extensions` |
| `ibm_compute_bare_metal_invalid_ssh_key_ids` | `ssh_key_ids` |
| `ibm_compute_bare_metal_invalid_tags` | `tags` |
| `ibm_compute_bare_metal_invalid_tcp_monitoring` | `tcp_monitoring` |
| `ibm_compute_bare_metal_invalid_unbonded_network` | `unbonded_network` |
| `ibm_compute_bare_metal_invalid_user_metadata` | `user_metadata` |
| `ibm_compute_bare_metal_invalid_storage_groups__array_size` | `storage_groups.array_size` |
| `ibm_compute_bare_metal_invalid_storage_groups__array_type_id` | `storage_groups.array_type_id` |
| `ibm_compute_bare_metal_invalid_storage_groups__hard_drives` | `storage_groups.hard_drives` |
| `ibm_compute_bare_metal_invalid_storage_groups__partition_template_id` | `storage_groups.partition_template_id` |

## Example

### Valid

```hcl
resource "ibm_compute_bare_metal" "test" {
  domain = "example"
}
```
