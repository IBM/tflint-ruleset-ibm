# ibm_is_instance

Manages ibm_is_instance. Instance name

This rule validates `ibm_is_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Instance` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `action` | no | one of: `stop`, `start`, `reboot` |
| `auto_delete_volume` | no | — |
| `availability_policy_host_failure` | no | one of: `restart`, `stop` |
| `bandwidth` | no | — |
| `cluster_network` | no | — |
| `confidential_compute_mode` | no | one of: `disabled`, `sgx`, `tdx`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `crn` | no | — |
| `dedicated_host` | no | API: If present, the dedicated host this virtual server instance has been placed on. |
| `dedicated_host_group` | no | — |
| `default_trusted_profile_auto_link` | no | — |
| `default_trusted_profile_target` | no | — |
| `disks` | no | — |
| `enable_secure_boot` | no | API: Indicates whether secure boot is enabled for this virtual server instance. |
| `force_action` | no | — |
| `force_recovery_time` | no | — |
| `gpu` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `image` | no | API: The image the virtual server instance was provisioned from |
| `instance_template` | no | — |
| `keys` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `memory` | no | — |
| `metadata_service_enabled` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `numa_count` | no | — |
| `placement_group` | no | — |
| `placement_target` | no | — |
| `profile` | no | API: The [profile](https://cloud.ibm.com/docs/vpc?topic=vpc-profiles) for this virtual server instance |
| `reservation` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group for this instance |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `software_attachments` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `threads_per_core` | no | range 1–2 |
| `total_network_bandwidth` | no | — |
| `total_volume_bandwidth` | no | min 500 |
| `user_data` | no | API: The [user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data) to make available when setting up the bare metal server. |
| `volume_attachments` | no | — |
| `volume_bandwidth_qos_mode` | no | one of: `pooled`, `weighted`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `volumes` | no | — |
| `vpc` | no | API: The VPC this virtual server instance resides in |
| `wait_before_delete` | no | — |
| `zone` | no | zone format |
| `availability` (block) | no | nested block |
| `availability.class` | no | — |
| `availability_policy` (block) | no | nested block |
| `availability_policy.host_failure` | no | one of: `restart`, `stop`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `availability_policy.preemption` | no | one of: `delete`, `stop`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `boot_volume` (block) | no | nested block |
| `boot_volume.auto_delete_volume` | no | — |
| `boot_volume.bandwidth` | no | API: The total bandwidth (in megabits per second) shared across the instance network attachments or instance network interfaces and storage volumes of the virtual server instance |
| `boot_volume.encryption` | no | — |
| `boot_volume.iops` | no | range 100–9007199254740991 |
| `boot_volume.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `boot_volume.profile` | no | API: The [profile](https://cloud.ibm.com/docs/vpc?topic=vpc-profiles) for this virtual server instance |
| `boot_volume.size` | no | range 1–100000 |
| `boot_volume.snapshot` | no | — |
| `boot_volume.snapshot_crn` | no | crn format |
| `boot_volume.tags` | no | — |
| `boot_volume.volume_id` | no | — |
| `boot_volume.allowed_use` (block) | no | nested block |
| `boot_volume.allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `boot_volume.allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `boot_volume.allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `catalog_offering` (block) | no | nested block |
| `catalog_offering.offering_crn` | no | crn format |
| `catalog_offering.plan_crn` | no | crn format |
| `catalog_offering.version_crn` | no | crn format |
| `cluster_network_attachments` (block) | no | nested block |
| `cluster_network_attachments.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `cluster_network_attachments.cluster_network_interface` (block) | yes | nested block |
| `cluster_network_attachments.cluster_network_interface.auto_delete` | no | — |
| `cluster_network_attachments.cluster_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `cluster_network_attachments.cluster_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `cluster_network_attachments.cluster_network_interface.primary_ip` (block) | no | nested block |
| `cluster_network_attachments.cluster_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `cluster_network_attachments.cluster_network_interface.primary_ip.auto_delete` | no | — |
| `cluster_network_attachments.cluster_network_interface.primary_ip.id` | no | pattern `^[-0-9a-z_]+$` |
| `cluster_network_attachments.cluster_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `cluster_network_attachments.cluster_network_interface.subnet` (block) | no | nested block |
| `cluster_network_attachments.cluster_network_interface.subnet.href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `cluster_network_attachments.cluster_network_interface.subnet.id` | no | pattern `^[-0-9a-z_]+$` |
| `metadata_service` (block) | no | nested block |
| `metadata_service.enabled` | no | API: Indicates whether the metadata service endpoint is available to the virtual server instance |
| `metadata_service.protocol` | no | one of: `http`, `https`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `metadata_service.response_hop_limit` | no | range 1–64 |
| `network_attachments` (block) | no | nested block |
| `network_attachments.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachments.virtual_network_interface` (block) | no | nested block |
| `network_attachments.virtual_network_interface.allow_ip_spoofing` | no | — |
| `network_attachments.virtual_network_interface.auto_delete` | no | — |
| `network_attachments.virtual_network_interface.enable_infrastructure_nat` | no | — |
| `network_attachments.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `network_attachments.virtual_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachments.virtual_network_interface.protocol_state_filtering_mode` | no | — |
| `network_attachments.virtual_network_interface.resource_group` | no | API: The resource group for this instance |
| `network_attachments.virtual_network_interface.security_groups` | no | — |
| `network_attachments.virtual_network_interface.subnet` | no | API: The subnet of the virtual network interface for the instance network attachment |
| `network_attachments.virtual_network_interface.ips` (block) | no | nested block |
| `network_attachments.virtual_network_interface.ips.reserved_ip` | yes | — |
| `network_attachments.virtual_network_interface.primary_ip` (block) | no | nested block |
| `network_attachments.virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `network_attachments.virtual_network_interface.primary_ip.auto_delete` | no | — |
| `network_attachments.virtual_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachments.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `network_interfaces` (block) | no | nested block |
| `network_interfaces.allow_ip_spoofing` | no | — |
| `network_interfaces.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_interfaces.primary_ipv4_address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `network_interfaces.security_groups` | no | — |
| `network_interfaces.subnet` | yes | API: The associated subnet |
| `network_interfaces.primary_ip` (block) | no | nested block |
| `network_interfaces.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `network_interfaces.primary_ip.auto_delete` | no | — |
| `network_interfaces.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_interfaces.primary_ip.reserved_ip` | no | — |
| `primary_network_attachment` (block) | no | nested block |
| `primary_network_attachment.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `primary_network_attachment.virtual_network_interface` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.allow_ip_spoofing` | no | — |
| `primary_network_attachment.virtual_network_interface.auto_delete` | no | — |
| `primary_network_attachment.virtual_network_interface.enable_infrastructure_nat` | no | — |
| `primary_network_attachment.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `primary_network_attachment.virtual_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `primary_network_attachment.virtual_network_interface.protocol_state_filtering_mode` | no | — |
| `primary_network_attachment.virtual_network_interface.resource_group` | no | API: The resource group for this instance |
| `primary_network_attachment.virtual_network_interface.security_groups` | no | — |
| `primary_network_attachment.virtual_network_interface.subnet` | no | API: The subnet of the virtual network interface for the instance network attachment |
| `primary_network_attachment.virtual_network_interface.ips` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.ips.reserved_ip` | yes | — |
| `primary_network_attachment.virtual_network_interface.primary_ip` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_network_attachment.virtual_network_interface.primary_ip.auto_delete` | no | — |
| `primary_network_attachment.virtual_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `primary_network_attachment.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `primary_network_interface` (block) | no | nested block |
| `primary_network_interface.allow_ip_spoofing` | no | — |
| `primary_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `primary_network_interface.port_speed` | no | — |
| `primary_network_interface.primary_ipv4_address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_network_interface.security_groups` | no | — |
| `primary_network_interface.subnet` | yes | API: The associated subnet |
| `primary_network_interface.primary_ip` (block) | no | nested block |
| `primary_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_network_interface.primary_ip.auto_delete` | no | — |
| `primary_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `primary_network_interface.primary_ip.reserved_ip` | no | — |
| `reservation_affinity` (block) | no | nested block |
| `reservation_affinity.policy` | no | one of: `automatic`, `disabled`, `manual`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `reservation_affinity.pool` (block) | no | nested block |
| `reservation_affinity.pool.id` | no | pattern `^[-0-9a-z_]+$` |
| `vcpu` (block) | no | nested block |
| `vcpu.percentage` | no | range 1–100 |
| `volume_prototypes` (block) | no | nested block |
| `volume_prototypes.delete_volume_on_instance_delete` | no | — |
| `volume_prototypes.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `volume_prototypes.volume_bandwidth` | no | — |
| `volume_prototypes.volume_capacity` | no | — |
| `volume_prototypes.volume_encryption_key` | no | — |
| `volume_prototypes.volume_iops` | no | — |
| `volume_prototypes.volume_name` | no | — |
| `volume_prototypes.volume_profile` | no | — |
| `volume_prototypes.volume_source_snapshot` | no | — |
| `volume_prototypes.volume_tags` | no | — |
| `volume_prototypes.allowed_use` (block) | no | nested block |
| `volume_prototypes.allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `volume_prototypes.allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `volume_prototypes.allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_invalid_access_tags` | `access_tags` |
| `ibm_is_instance_invalid_action` | `action` |
| `ibm_is_instance_invalid_auto_delete_volume` | `auto_delete_volume` |
| `ibm_is_instance_invalid_availability_policy_host_failure` | `availability_policy_host_failure` |
| `ibm_is_instance_invalid_confidential_compute_mode` | `confidential_compute_mode` |
| `ibm_is_instance_invalid_dedicated_host` | `dedicated_host` |
| `ibm_is_instance_invalid_dedicated_host_group` | `dedicated_host_group` |
| `ibm_is_instance_invalid_default_trusted_profile_auto_link` | `default_trusted_profile_auto_link` |
| `ibm_is_instance_invalid_default_trusted_profile_target` | `default_trusted_profile_target` |
| `ibm_is_instance_invalid_enable_secure_boot` | `enable_secure_boot` |
| `ibm_is_instance_invalid_force_action` | `force_action` |
| `ibm_is_instance_invalid_force_recovery_time` | `force_recovery_time` |
| `ibm_is_instance_invalid_image` | `image` |
| `ibm_is_instance_invalid_instance_template` | `instance_template` |
| `ibm_is_instance_invalid_keys` | `keys` |
| `ibm_is_instance_invalid_metadata_service_enabled` | `metadata_service_enabled` |
| `ibm_is_instance_invalid_name` | `name` |
| `ibm_is_instance_invalid_placement_group` | `placement_group` |
| `ibm_is_instance_invalid_profile` | `profile` |
| `ibm_is_instance_invalid_resource_group` | `resource_group` |
| `ibm_is_instance_invalid_tags` | `tags` |
| `ibm_is_instance_invalid_threads_per_core` | `threads_per_core` |
| `ibm_is_instance_invalid_total_volume_bandwidth` | `total_volume_bandwidth` |
| `ibm_is_instance_invalid_user_data` | `user_data` |
| `ibm_is_instance_invalid_volume_bandwidth_qos_mode` | `volume_bandwidth_qos_mode` |
| `ibm_is_instance_invalid_volumes` | `volumes` |
| `ibm_is_instance_invalid_vpc` | `vpc` |
| `ibm_is_instance_invalid_wait_before_delete` | `wait_before_delete` |
| `ibm_is_instance_invalid_zone` | `zone` |
| `ibm_is_instance_invalid_availability__class` | `availability.class` |
| `ibm_is_instance_invalid_availability_policy__host_failure` | `availability_policy.host_failure` |
| `ibm_is_instance_invalid_availability_policy__preemption` | `availability_policy.preemption` |
| `ibm_is_instance_invalid_boot_volume__auto_delete_volume` | `boot_volume.auto_delete_volume` |
| `ibm_is_instance_invalid_boot_volume__bandwidth` | `boot_volume.bandwidth` |
| `ibm_is_instance_invalid_boot_volume__encryption` | `boot_volume.encryption` |
| `ibm_is_instance_invalid_boot_volume__iops` | `boot_volume.iops` |
| `ibm_is_instance_invalid_boot_volume__name` | `boot_volume.name` |
| `ibm_is_instance_invalid_boot_volume__profile` | `boot_volume.profile` |
| `ibm_is_instance_invalid_boot_volume__size` | `boot_volume.size` |
| `ibm_is_instance_invalid_boot_volume__snapshot` | `boot_volume.snapshot` |
| `ibm_is_instance_invalid_boot_volume__snapshot_crn` | `boot_volume.snapshot_crn` |
| `ibm_is_instance_invalid_boot_volume__tags` | `boot_volume.tags` |
| `ibm_is_instance_invalid_boot_volume__volume_id` | `boot_volume.volume_id` |
| `ibm_is_instance_invalid_boot_volume__allowed_use__api_version` | `boot_volume.allowed_use.api_version` |
| `ibm_is_instance_invalid_boot_volume__allowed_use__bare_metal_server` | `boot_volume.allowed_use.bare_metal_server` |
| `ibm_is_instance_invalid_boot_volume__allowed_use__instance` | `boot_volume.allowed_use.instance` |
| `ibm_is_instance_invalid_catalog_offering__offering_crn` | `catalog_offering.offering_crn` |
| `ibm_is_instance_invalid_catalog_offering__plan_crn` | `catalog_offering.plan_crn` |
| `ibm_is_instance_invalid_catalog_offering__version_crn` | `catalog_offering.version_crn` |
| `ibm_is_instance_invalid_cluster_network_attachments__name` | `cluster_network_attachments.name` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__auto_delete` | `cluster_network_attachments.cluster_network_interface.auto_delete` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__id` | `cluster_network_attachments.cluster_network_interface.id` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__name` | `cluster_network_attachments.cluster_network_interface.name` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__primary_ip__address` | `cluster_network_attachments.cluster_network_interface.primary_ip.address` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__primary_ip__auto_delete` | `cluster_network_attachments.cluster_network_interface.primary_ip.auto_delete` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__primary_ip__id` | `cluster_network_attachments.cluster_network_interface.primary_ip.id` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__primary_ip__name` | `cluster_network_attachments.cluster_network_interface.primary_ip.name` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__subnet__href` | `cluster_network_attachments.cluster_network_interface.subnet.href` |
| `ibm_is_instance_invalid_cluster_network_attachments__cluster_network_interface__subnet__id` | `cluster_network_attachments.cluster_network_interface.subnet.id` |
| `ibm_is_instance_invalid_metadata_service__enabled` | `metadata_service.enabled` |
| `ibm_is_instance_invalid_metadata_service__protocol` | `metadata_service.protocol` |
| `ibm_is_instance_invalid_metadata_service__response_hop_limit` | `metadata_service.response_hop_limit` |
| `ibm_is_instance_invalid_network_attachments__name` | `network_attachments.name` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__allow_ip_spoofing` | `network_attachments.virtual_network_interface.allow_ip_spoofing` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__auto_delete` | `network_attachments.virtual_network_interface.auto_delete` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__enable_infrastructure_nat` | `network_attachments.virtual_network_interface.enable_infrastructure_nat` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__id` | `network_attachments.virtual_network_interface.id` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__name` | `network_attachments.virtual_network_interface.name` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__protocol_state_filtering_mode` | `network_attachments.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__resource_group` | `network_attachments.virtual_network_interface.resource_group` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__security_groups` | `network_attachments.virtual_network_interface.security_groups` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__subnet` | `network_attachments.virtual_network_interface.subnet` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__ips__reserved_ip` | `network_attachments.virtual_network_interface.ips.reserved_ip` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__primary_ip__address` | `network_attachments.virtual_network_interface.primary_ip.address` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__primary_ip__auto_delete` | `network_attachments.virtual_network_interface.primary_ip.auto_delete` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__primary_ip__name` | `network_attachments.virtual_network_interface.primary_ip.name` |
| `ibm_is_instance_invalid_network_attachments__virtual_network_interface__primary_ip__reserved_ip` | `network_attachments.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_instance_invalid_network_interfaces__allow_ip_spoofing` | `network_interfaces.allow_ip_spoofing` |
| `ibm_is_instance_invalid_network_interfaces__name` | `network_interfaces.name` |
| `ibm_is_instance_invalid_network_interfaces__primary_ipv4_address` | `network_interfaces.primary_ipv4_address` |
| `ibm_is_instance_invalid_network_interfaces__security_groups` | `network_interfaces.security_groups` |
| `ibm_is_instance_invalid_network_interfaces__subnet` | `network_interfaces.subnet` |
| `ibm_is_instance_invalid_network_interfaces__primary_ip__address` | `network_interfaces.primary_ip.address` |
| `ibm_is_instance_invalid_network_interfaces__primary_ip__auto_delete` | `network_interfaces.primary_ip.auto_delete` |
| `ibm_is_instance_invalid_network_interfaces__primary_ip__name` | `network_interfaces.primary_ip.name` |
| `ibm_is_instance_invalid_network_interfaces__primary_ip__reserved_ip` | `network_interfaces.primary_ip.reserved_ip` |
| `ibm_is_instance_invalid_primary_network_attachment__name` | `primary_network_attachment.name` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__allow_ip_spoofing` | `primary_network_attachment.virtual_network_interface.allow_ip_spoofing` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__auto_delete` | `primary_network_attachment.virtual_network_interface.auto_delete` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__enable_infrastructure_nat` | `primary_network_attachment.virtual_network_interface.enable_infrastructure_nat` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__id` | `primary_network_attachment.virtual_network_interface.id` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__name` | `primary_network_attachment.virtual_network_interface.name` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__protocol_state_filtering_mode` | `primary_network_attachment.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__resource_group` | `primary_network_attachment.virtual_network_interface.resource_group` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__security_groups` | `primary_network_attachment.virtual_network_interface.security_groups` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__subnet` | `primary_network_attachment.virtual_network_interface.subnet` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__ips__reserved_ip` | `primary_network_attachment.virtual_network_interface.ips.reserved_ip` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__primary_ip__address` | `primary_network_attachment.virtual_network_interface.primary_ip.address` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__primary_ip__auto_delete` | `primary_network_attachment.virtual_network_interface.primary_ip.auto_delete` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__primary_ip__name` | `primary_network_attachment.virtual_network_interface.primary_ip.name` |
| `ibm_is_instance_invalid_primary_network_attachment__virtual_network_interface__primary_ip__reserved_ip` | `primary_network_attachment.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_instance_invalid_primary_network_interface__allow_ip_spoofing` | `primary_network_interface.allow_ip_spoofing` |
| `ibm_is_instance_invalid_primary_network_interface__name` | `primary_network_interface.name` |
| `ibm_is_instance_invalid_primary_network_interface__port_speed` | `primary_network_interface.port_speed` |
| `ibm_is_instance_invalid_primary_network_interface__primary_ipv4_address` | `primary_network_interface.primary_ipv4_address` |
| `ibm_is_instance_invalid_primary_network_interface__security_groups` | `primary_network_interface.security_groups` |
| `ibm_is_instance_invalid_primary_network_interface__subnet` | `primary_network_interface.subnet` |
| `ibm_is_instance_invalid_primary_network_interface__primary_ip__address` | `primary_network_interface.primary_ip.address` |
| `ibm_is_instance_invalid_primary_network_interface__primary_ip__auto_delete` | `primary_network_interface.primary_ip.auto_delete` |
| `ibm_is_instance_invalid_primary_network_interface__primary_ip__name` | `primary_network_interface.primary_ip.name` |
| `ibm_is_instance_invalid_primary_network_interface__primary_ip__reserved_ip` | `primary_network_interface.primary_ip.reserved_ip` |
| `ibm_is_instance_invalid_reservation_affinity__policy` | `reservation_affinity.policy` |
| `ibm_is_instance_invalid_reservation_affinity__pool__id` | `reservation_affinity.pool.id` |
| `ibm_is_instance_invalid_vcpu__percentage` | `vcpu.percentage` |
| `ibm_is_instance_invalid_volume_prototypes__delete_volume_on_instance_delete` | `volume_prototypes.delete_volume_on_instance_delete` |
| `ibm_is_instance_invalid_volume_prototypes__name` | `volume_prototypes.name` |
| `ibm_is_instance_invalid_volume_prototypes__volume_bandwidth` | `volume_prototypes.volume_bandwidth` |
| `ibm_is_instance_invalid_volume_prototypes__volume_capacity` | `volume_prototypes.volume_capacity` |
| `ibm_is_instance_invalid_volume_prototypes__volume_encryption_key` | `volume_prototypes.volume_encryption_key` |
| `ibm_is_instance_invalid_volume_prototypes__volume_iops` | `volume_prototypes.volume_iops` |
| `ibm_is_instance_invalid_volume_prototypes__volume_name` | `volume_prototypes.volume_name` |
| `ibm_is_instance_invalid_volume_prototypes__volume_profile` | `volume_prototypes.volume_profile` |
| `ibm_is_instance_invalid_volume_prototypes__volume_source_snapshot` | `volume_prototypes.volume_source_snapshot` |
| `ibm_is_instance_invalid_volume_prototypes__volume_tags` | `volume_prototypes.volume_tags` |
| `ibm_is_instance_invalid_volume_prototypes__allowed_use__api_version` | `volume_prototypes.allowed_use.api_version` |
| `ibm_is_instance_invalid_volume_prototypes__allowed_use__bare_metal_server` | `volume_prototypes.allowed_use.bare_metal_server` |
| `ibm_is_instance_invalid_volume_prototypes__allowed_use__instance` | `volume_prototypes.allowed_use.instance` |

## Example

### Valid

```hcl
resource "ibm_is_instance" "test" {
  name = "example-name"
}
```
