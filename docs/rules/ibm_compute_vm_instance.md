# ibm_compute_vm_instance

Schema validation for the ibm_compute_vm_instance resource.

This rule validates `ibm_compute_vm_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_vm_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `block_storage_ids` | no | — |
| `cores` | no | — |
| `datacenter` | no | — |
| `datacenter_choice` | no | — |
| `dedicated_acct_host_only` | no | — |
| `dedicated_host_id` | no | — |
| `dedicated_host_name` | no | — |
| `disks` | no | — |
| `domain` | no | — |
| `evault` | no | — |
| `file_storage_ids` | no | — |
| `flavor_key_name` | no | — |
| `hostname` | no | — |
| `hourly_billing` | no | — |
| `image_id` | no | — |
| `ip_address_id` | no | — |
| `ip_address_id_private` | no | — |
| `ipv4_address` | no | — |
| `ipv4_address_private` | no | — |
| `ipv6_address` | no | — |
| `ipv6_address_id` | no | — |
| `ipv6_enabled` | no | — |
| `ipv6_static_enabled` | no | — |
| `local_disk` | no | — |
| `memory` | no | — |
| `network_speed` | no | — |
| `notes` | no | — |
| `os_reference_code` | no | — |
| `placement_group_id` | no | — |
| `placement_group_name` | no | — |
| `post_install_script_uri` | no | — |
| `private_interface_id` | no | — |
| `private_network_only` | no | — |
| `private_security_group_ids` | no | — |
| `private_subnet` | no | — |
| `private_subnet_id` | no | — |
| `private_vlan_id` | no | — |
| `public_bandwidth_limited` | no | — |
| `public_bandwidth_unlimited` | no | — |
| `public_interface_id` | no | — |
| `public_ipv6_subnet` | no | — |
| `public_ipv6_subnet_id` | no | — |
| `public_security_group_ids` | no | — |
| `public_subnet` | no | — |
| `public_subnet_id` | no | — |
| `public_vlan_id` | no | — |
| `quote_id` | no | — |
| `reserved_capacity_id` | no | — |
| `reserved_capacity_name` | no | — |
| `reserved_instance_primary_disk` | no | — |
| `resource_controller_url` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `secondary_ip_addresses` | no | — |
| `secondary_ip_count` | no | — |
| `ssh_key_ids` | no | — |
| `tags` | no | — |
| `transient` | no | — |
| `user_metadata` | no | — |
| `wait_time_minutes` | no | — |
| `bulk_vms` (block) | no | nested block |
| `bulk_vms.domain` | yes | — |
| `bulk_vms.hostname` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_vm_instance_invalid_block_storage_ids` | `block_storage_ids` |
| `ibm_compute_vm_instance_invalid_cores` | `cores` |
| `ibm_compute_vm_instance_invalid_datacenter` | `datacenter` |
| `ibm_compute_vm_instance_invalid_datacenter_choice` | `datacenter_choice` |
| `ibm_compute_vm_instance_invalid_dedicated_acct_host_only` | `dedicated_acct_host_only` |
| `ibm_compute_vm_instance_invalid_dedicated_host_id` | `dedicated_host_id` |
| `ibm_compute_vm_instance_invalid_dedicated_host_name` | `dedicated_host_name` |
| `ibm_compute_vm_instance_invalid_disks` | `disks` |
| `ibm_compute_vm_instance_invalid_domain` | `domain` |
| `ibm_compute_vm_instance_invalid_evault` | `evault` |
| `ibm_compute_vm_instance_invalid_file_storage_ids` | `file_storage_ids` |
| `ibm_compute_vm_instance_invalid_flavor_key_name` | `flavor_key_name` |
| `ibm_compute_vm_instance_invalid_hostname` | `hostname` |
| `ibm_compute_vm_instance_invalid_hourly_billing` | `hourly_billing` |
| `ibm_compute_vm_instance_invalid_image_id` | `image_id` |
| `ibm_compute_vm_instance_invalid_ipv6_enabled` | `ipv6_enabled` |
| `ibm_compute_vm_instance_invalid_ipv6_static_enabled` | `ipv6_static_enabled` |
| `ibm_compute_vm_instance_invalid_local_disk` | `local_disk` |
| `ibm_compute_vm_instance_invalid_memory` | `memory` |
| `ibm_compute_vm_instance_invalid_network_speed` | `network_speed` |
| `ibm_compute_vm_instance_invalid_notes` | `notes` |
| `ibm_compute_vm_instance_invalid_os_reference_code` | `os_reference_code` |
| `ibm_compute_vm_instance_invalid_placement_group_id` | `placement_group_id` |
| `ibm_compute_vm_instance_invalid_placement_group_name` | `placement_group_name` |
| `ibm_compute_vm_instance_invalid_post_install_script_uri` | `post_install_script_uri` |
| `ibm_compute_vm_instance_invalid_private_network_only` | `private_network_only` |
| `ibm_compute_vm_instance_invalid_private_security_group_ids` | `private_security_group_ids` |
| `ibm_compute_vm_instance_invalid_private_subnet` | `private_subnet` |
| `ibm_compute_vm_instance_invalid_private_vlan_id` | `private_vlan_id` |
| `ibm_compute_vm_instance_invalid_public_bandwidth_limited` | `public_bandwidth_limited` |
| `ibm_compute_vm_instance_invalid_public_bandwidth_unlimited` | `public_bandwidth_unlimited` |
| `ibm_compute_vm_instance_invalid_public_security_group_ids` | `public_security_group_ids` |
| `ibm_compute_vm_instance_invalid_public_subnet` | `public_subnet` |
| `ibm_compute_vm_instance_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_compute_vm_instance_invalid_quote_id` | `quote_id` |
| `ibm_compute_vm_instance_invalid_reserved_capacity_id` | `reserved_capacity_id` |
| `ibm_compute_vm_instance_invalid_reserved_capacity_name` | `reserved_capacity_name` |
| `ibm_compute_vm_instance_invalid_reserved_instance_primary_disk` | `reserved_instance_primary_disk` |
| `ibm_compute_vm_instance_invalid_secondary_ip_count` | `secondary_ip_count` |
| `ibm_compute_vm_instance_invalid_ssh_key_ids` | `ssh_key_ids` |
| `ibm_compute_vm_instance_invalid_tags` | `tags` |
| `ibm_compute_vm_instance_invalid_transient` | `transient` |
| `ibm_compute_vm_instance_invalid_user_metadata` | `user_metadata` |
| `ibm_compute_vm_instance_invalid_wait_time_minutes` | `wait_time_minutes` |
| `ibm_compute_vm_instance_invalid_bulk_vms__domain` | `bulk_vms.domain` |
| `ibm_compute_vm_instance_invalid_bulk_vms__hostname` | `bulk_vms.hostname` |

## Example

### Valid

```hcl
resource "ibm_compute_vm_instance" "test" {
}
```
