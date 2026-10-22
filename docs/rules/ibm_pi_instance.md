# ibm_pi_instance

Schema validation for the ibm_pi_instance resource.

This rule validates `ibm_pi_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `dedicated_host_id` | no | — |
| `effective_processor_compatibility_mode` | no | — |
| `fault` | no | — |
| `health_status` | no | — |
| `ibmi_rds` | no | — |
| `instance_id` | no | — |
| `max_memory` | no | — |
| `max_processors` | no | — |
| `max_virtual_cores` | no | — |
| `min_memory` | no | — |
| `min_processors` | no | — |
| `min_virtual_cores` | no | — |
| `operating_system` | no | — |
| `os_type` | no | — |
| `pi_affinity_instance` | no | — |
| `pi_affinity_policy` | no | one of: `affinity`, `anti-affinity` |
| `pi_affinity_volume` | no | — |
| `pi_allow_remote_restart` | no | — |
| `pi_anti_affinity_instances` | no | — |
| `pi_anti_affinity_volumes` | no | — |
| `pi_boot_volume_replication_enabled` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_deployment_type` | no | one of: `EPIC`, `VMNoStorage` |
| `pi_health_status` | no | one of: `OK`, `WARNING` |
| `pi_ibmi_css` | no | — |
| `pi_ibmi_pha` | no | — |
| `pi_ibmi_rds_users` | no | — |
| `pi_image_id` | yes | — |
| `pi_instance_name` | yes | — |
| `pi_key_pair_name` | no | — |
| `pi_license_repository_capacity` | no | — |
| `pi_memory` | no | — |
| `pi_pin_policy` | no | one of: `none`, `soft`, `hard` |
| `pi_placement_group_id` | no | — |
| `pi_preferred_processor_compatibility_mode` | no | — |
| `pi_proc_type` | no | one of: `dedicated`, `shared`, `capped` |
| `pi_processors` | no | — |
| `pi_replicants` | no | — |
| `pi_replication_policy` | no | one of: `affinity`, `anti-affinity`, `none` |
| `pi_replication_scheme` | no | one of: `prefix`, `suffix` |
| `pi_replication_sites` | no | — |
| `pi_retain_virtual_serial_number` | no | — |
| `pi_sap_deployment_type` | no | — |
| `pi_sap_profile_id` | no | — |
| `pi_shared_processor_pool` | no | — |
| `pi_storage_connection` | no | — |
| `pi_storage_pool` | no | — |
| `pi_storage_pool_affinity` | no | — |
| `pi_storage_type` | no | — |
| `pi_sys_type` | no | — |
| `pi_user_data` | no | — |
| `pi_user_tags` | no | — |
| `pi_virtual_cores_assigned` | no | — |
| `pi_virtual_optical_device` | no | one of: `attach` |
| `pi_volume_ids` | no | — |
| `pin_policy` | no | — |
| `progress` | no | — |
| `shared_processor_pool_id` | no | — |
| `status` | no | — |
| `pi_default_trusted_profile` (block) | no | nested block |
| `pi_default_trusted_profile.autolink` | no | — |
| `pi_default_trusted_profile.target` (block) | yes | nested block |
| `pi_default_trusted_profile.target.crn` | no | pattern `^(crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\.\/]*)?$` |
| `pi_default_trusted_profile.target.id` | no | API: Unique identifier for the trusted profile |
| `pi_default_trusted_profile.target.name` | no | — |
| `pi_deployment_target` (block) | no | nested block |
| `pi_deployment_target.id` | yes | API: Unique identifier for the trusted profile |
| `pi_deployment_target.type` | yes | one of: `host`, `hostGroup` |
| `pi_metadata_service` (block) | no | nested block |
| `pi_metadata_service.enabled` | yes | — |
| `pi_metadata_service.force_disable` | no | — |
| `pi_network` (block) | yes | nested block |
| `pi_network.ip_address` | no | ipv4 format |
| `pi_network.network_id` | yes | — |
| `pi_network.network_security_group_ids` | no | — |
| `pi_virtual_serial_number` (block) | no | nested block |
| `pi_virtual_serial_number.description` | no | API: Description of the PVM instance snapshot |
| `pi_virtual_serial_number.serial` | yes | — |
| `pi_virtual_serial_number.software_tier` | no | one of: `P05`, `P10`, `P20`, `P30` |
| `pi_vpmem_volumes` (block) | no | nested block |
| `pi_vpmem_volumes.name` | yes | — |
| `pi_vpmem_volumes.size` | yes | min 1 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_invalid_pi_affinity_instance` | `pi_affinity_instance` |
| `ibm_pi_instance_invalid_pi_affinity_policy` | `pi_affinity_policy` |
| `ibm_pi_instance_invalid_pi_affinity_volume` | `pi_affinity_volume` |
| `ibm_pi_instance_invalid_pi_allow_remote_restart` | `pi_allow_remote_restart` |
| `ibm_pi_instance_invalid_pi_anti_affinity_instances` | `pi_anti_affinity_instances` |
| `ibm_pi_instance_invalid_pi_anti_affinity_volumes` | `pi_anti_affinity_volumes` |
| `ibm_pi_instance_invalid_pi_boot_volume_replication_enabled` | `pi_boot_volume_replication_enabled` |
| `ibm_pi_instance_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_invalid_pi_deployment_type` | `pi_deployment_type` |
| `ibm_pi_instance_invalid_pi_health_status` | `pi_health_status` |
| `ibm_pi_instance_invalid_pi_ibmi_css` | `pi_ibmi_css` |
| `ibm_pi_instance_invalid_pi_ibmi_pha` | `pi_ibmi_pha` |
| `ibm_pi_instance_invalid_pi_ibmi_rds_users` | `pi_ibmi_rds_users` |
| `ibm_pi_instance_invalid_pi_image_id` | `pi_image_id` |
| `ibm_pi_instance_invalid_pi_instance_name` | `pi_instance_name` |
| `ibm_pi_instance_invalid_pi_key_pair_name` | `pi_key_pair_name` |
| `ibm_pi_instance_invalid_pi_license_repository_capacity` | `pi_license_repository_capacity` |
| `ibm_pi_instance_invalid_pi_memory` | `pi_memory` |
| `ibm_pi_instance_invalid_pi_pin_policy` | `pi_pin_policy` |
| `ibm_pi_instance_invalid_pi_placement_group_id` | `pi_placement_group_id` |
| `ibm_pi_instance_invalid_pi_preferred_processor_compatibility_mode` | `pi_preferred_processor_compatibility_mode` |
| `ibm_pi_instance_invalid_pi_proc_type` | `pi_proc_type` |
| `ibm_pi_instance_invalid_pi_processors` | `pi_processors` |
| `ibm_pi_instance_invalid_pi_replicants` | `pi_replicants` |
| `ibm_pi_instance_invalid_pi_replication_policy` | `pi_replication_policy` |
| `ibm_pi_instance_invalid_pi_replication_scheme` | `pi_replication_scheme` |
| `ibm_pi_instance_invalid_pi_replication_sites` | `pi_replication_sites` |
| `ibm_pi_instance_invalid_pi_retain_virtual_serial_number` | `pi_retain_virtual_serial_number` |
| `ibm_pi_instance_invalid_pi_sap_deployment_type` | `pi_sap_deployment_type` |
| `ibm_pi_instance_invalid_pi_sap_profile_id` | `pi_sap_profile_id` |
| `ibm_pi_instance_invalid_pi_shared_processor_pool` | `pi_shared_processor_pool` |
| `ibm_pi_instance_invalid_pi_storage_connection` | `pi_storage_connection` |
| `ibm_pi_instance_invalid_pi_storage_pool` | `pi_storage_pool` |
| `ibm_pi_instance_invalid_pi_storage_pool_affinity` | `pi_storage_pool_affinity` |
| `ibm_pi_instance_invalid_pi_storage_type` | `pi_storage_type` |
| `ibm_pi_instance_invalid_pi_sys_type` | `pi_sys_type` |
| `ibm_pi_instance_invalid_pi_user_data` | `pi_user_data` |
| `ibm_pi_instance_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_instance_invalid_pi_virtual_cores_assigned` | `pi_virtual_cores_assigned` |
| `ibm_pi_instance_invalid_pi_virtual_optical_device` | `pi_virtual_optical_device` |
| `ibm_pi_instance_invalid_pi_volume_ids` | `pi_volume_ids` |
| `ibm_pi_instance_invalid_pi_default_trusted_profile__autolink` | `pi_default_trusted_profile.autolink` |
| `ibm_pi_instance_invalid_pi_default_trusted_profile__target__crn` | `pi_default_trusted_profile.target.crn` |
| `ibm_pi_instance_invalid_pi_default_trusted_profile__target__id` | `pi_default_trusted_profile.target.id` |
| `ibm_pi_instance_invalid_pi_default_trusted_profile__target__name` | `pi_default_trusted_profile.target.name` |
| `ibm_pi_instance_invalid_pi_deployment_target__id` | `pi_deployment_target.id` |
| `ibm_pi_instance_invalid_pi_deployment_target__type` | `pi_deployment_target.type` |
| `ibm_pi_instance_invalid_pi_metadata_service__enabled` | `pi_metadata_service.enabled` |
| `ibm_pi_instance_invalid_pi_metadata_service__force_disable` | `pi_metadata_service.force_disable` |
| `ibm_pi_instance_invalid_pi_network__ip_address` | `pi_network.ip_address` |
| `ibm_pi_instance_invalid_pi_network__network_id` | `pi_network.network_id` |
| `ibm_pi_instance_invalid_pi_network__network_security_group_ids` | `pi_network.network_security_group_ids` |
| `ibm_pi_instance_invalid_pi_virtual_serial_number__description` | `pi_virtual_serial_number.description` |
| `ibm_pi_instance_invalid_pi_virtual_serial_number__serial` | `pi_virtual_serial_number.serial` |
| `ibm_pi_instance_invalid_pi_virtual_serial_number__software_tier` | `pi_virtual_serial_number.software_tier` |
| `ibm_pi_instance_invalid_pi_vpmem_volumes__name` | `pi_vpmem_volumes.name` |
| `ibm_pi_instance_invalid_pi_vpmem_volumes__size` | `pi_vpmem_volumes.size` |

## Example

### Valid

```hcl
resource "ibm_pi_instance" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_image_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_name = "example"
  pi_network {
    network_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
