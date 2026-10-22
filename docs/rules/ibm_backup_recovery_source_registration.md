# ibm_backup_recovery_source_registration

Manages ibm_backup_recovery_source_registration. The user specified name for this source.

This rule validates `ibm_backup_recovery_source_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_source_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authentication_status` | no | — |
| `auto_proetction_group_id` | no | — |
| `connection_id` | no | — |
| `connector_group_id` | no | — |
| `data_source_connection_id` | no | — |
| `encryption_key` | no | — |
| `environment` | yes | one of: `kPhysical`, `kSQL` |
| `external_metadata` | no | — |
| `is_internal_encrypted` | no | — |
| `last_refreshed_time_msecs` | no | — |
| `name` | no | API: Specifies the name of the Recovery. |
| `registration_time_msecs` | no | — |
| `source_id` | no | — |
| `source_info` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `advanced_configs` (block) | no | nested block |
| `advanced_configs.key` | yes | API: Specifies the unique key for custom property. |
| `advanced_configs.value` | yes | API: Specifies the value for the above custom key. |
| `connections` (block) | no | nested block |
| `connections.connection_id` | no | — |
| `connections.connector_group_id` | no | — |
| `connections.data_source_connection_id` | no | — |
| `connections.entity_id` | no | — |
| `kubernetes_params` (block) | no | nested block |
| `kubernetes_params.client_private_key` | yes | — |
| `kubernetes_params.cohesity_dataprotect_plugin_image_location` | no | — |
| `kubernetes_params.data_mover_image_location` | yes | — |
| `kubernetes_params.datamover_service_type` | no | — |
| `kubernetes_params.endpoint` | yes | — |
| `kubernetes_params.init_container_image_location` | no | — |
| `kubernetes_params.kubernetes_distribution` | yes | — |
| `kubernetes_params.kubernetes_type` | no | — |
| `kubernetes_params.priority_class_name` | no | — |
| `kubernetes_params.san_fields` | no | — |
| `kubernetes_params.velero_aws_plugin_image_location` | no | — |
| `kubernetes_params.velero_image_location` | no | — |
| `kubernetes_params.velero_openshift_plugin_image_location` | no | — |
| `kubernetes_params.auto_protect_config` (block) | no | nested block |
| `kubernetes_params.auto_protect_config.error_message` | no | — |
| `kubernetes_params.auto_protect_config.is_default_auto_protected` | yes | — |
| `kubernetes_params.auto_protect_config.policy_id` | yes | — |
| `kubernetes_params.auto_protect_config.protection_group_id` | no | — |
| `kubernetes_params.auto_protect_config.storage_domain_id` | no | — |
| `kubernetes_params.default_vlan_params` (block) | no | nested block |
| `kubernetes_params.default_vlan_params.disable_vlan` | no | — |
| `kubernetes_params.default_vlan_params.interface_name` | no | — |
| `kubernetes_params.default_vlan_params.vlan_id` | no | — |
| `kubernetes_params.resource_annotations` (block) | no | nested block |
| `kubernetes_params.resource_annotations.key` | yes | API: Specifies the unique key for custom property. |
| `kubernetes_params.resource_annotations.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.resource_labels` (block) | no | nested block |
| `kubernetes_params.resource_labels.key` | yes | API: Specifies the unique key for custom property. |
| `kubernetes_params.resource_labels.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.service_annotations` (block) | no | nested block |
| `kubernetes_params.service_annotations.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.service_annotations.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.vlan_info_vec` (block) | no | nested block |
| `kubernetes_params.vlan_info_vec.service_annotations` (block) | no | nested block |
| `kubernetes_params.vlan_info_vec.service_annotations.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.vlan_info_vec.service_annotations.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.vlan_info_vec.vlan_params` (block) | no | nested block |
| `kubernetes_params.vlan_info_vec.vlan_params.disable_vlan` | no | — |
| `kubernetes_params.vlan_info_vec.vlan_params.interface_name` | no | — |
| `kubernetes_params.vlan_info_vec.vlan_params.vlan_id` | no | — |
| `physical_params` (block) | no | nested block |
| `physical_params.applications` | no | — |
| `physical_params.endpoint` | yes | — |
| `physical_params.force_register` | no | — |
| `physical_params.host_type` | no | — |
| `physical_params.physical_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_source_registration_invalid_connection_id` | `connection_id` |
| `ibm_backup_recovery_source_registration_invalid_connector_group_id` | `connector_group_id` |
| `ibm_backup_recovery_source_registration_invalid_data_source_connection_id` | `data_source_connection_id` |
| `ibm_backup_recovery_source_registration_invalid_encryption_key` | `encryption_key` |
| `ibm_backup_recovery_source_registration_invalid_environment` | `environment` |
| `ibm_backup_recovery_source_registration_invalid_is_internal_encrypted` | `is_internal_encrypted` |
| `ibm_backup_recovery_source_registration_invalid_name` | `name` |
| `ibm_backup_recovery_source_registration_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_source_registration_invalid_advanced_configs__key` | `advanced_configs.key` |
| `ibm_backup_recovery_source_registration_invalid_advanced_configs__value` | `advanced_configs.value` |
| `ibm_backup_recovery_source_registration_invalid_connections__connection_id` | `connections.connection_id` |
| `ibm_backup_recovery_source_registration_invalid_connections__connector_group_id` | `connections.connector_group_id` |
| `ibm_backup_recovery_source_registration_invalid_connections__data_source_connection_id` | `connections.data_source_connection_id` |
| `ibm_backup_recovery_source_registration_invalid_connections__entity_id` | `connections.entity_id` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__client_private_key` | `kubernetes_params.client_private_key` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__cohesity_dataprotect_plugin_image_location` | `kubernetes_params.cohesity_dataprotect_plugin_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__data_mover_image_location` | `kubernetes_params.data_mover_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__datamover_service_type` | `kubernetes_params.datamover_service_type` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__endpoint` | `kubernetes_params.endpoint` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__init_container_image_location` | `kubernetes_params.init_container_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__kubernetes_distribution` | `kubernetes_params.kubernetes_distribution` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__kubernetes_type` | `kubernetes_params.kubernetes_type` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__priority_class_name` | `kubernetes_params.priority_class_name` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__san_fields` | `kubernetes_params.san_fields` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__velero_aws_plugin_image_location` | `kubernetes_params.velero_aws_plugin_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__velero_image_location` | `kubernetes_params.velero_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__velero_openshift_plugin_image_location` | `kubernetes_params.velero_openshift_plugin_image_location` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__auto_protect_config__error_message` | `kubernetes_params.auto_protect_config.error_message` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__auto_protect_config__is_default_auto_protected` | `kubernetes_params.auto_protect_config.is_default_auto_protected` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__auto_protect_config__policy_id` | `kubernetes_params.auto_protect_config.policy_id` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__auto_protect_config__protection_group_id` | `kubernetes_params.auto_protect_config.protection_group_id` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__auto_protect_config__storage_domain_id` | `kubernetes_params.auto_protect_config.storage_domain_id` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__default_vlan_params__disable_vlan` | `kubernetes_params.default_vlan_params.disable_vlan` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__default_vlan_params__interface_name` | `kubernetes_params.default_vlan_params.interface_name` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__default_vlan_params__vlan_id` | `kubernetes_params.default_vlan_params.vlan_id` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__resource_annotations__key` | `kubernetes_params.resource_annotations.key` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__resource_annotations__value` | `kubernetes_params.resource_annotations.value` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__resource_labels__key` | `kubernetes_params.resource_labels.key` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__resource_labels__value` | `kubernetes_params.resource_labels.value` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__service_annotations__key` | `kubernetes_params.service_annotations.key` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__service_annotations__value` | `kubernetes_params.service_annotations.value` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__vlan_info_vec__service_annotations__key` | `kubernetes_params.vlan_info_vec.service_annotations.key` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__vlan_info_vec__service_annotations__value` | `kubernetes_params.vlan_info_vec.service_annotations.value` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__vlan_info_vec__vlan_params__disable_vlan` | `kubernetes_params.vlan_info_vec.vlan_params.disable_vlan` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__vlan_info_vec__vlan_params__interface_name` | `kubernetes_params.vlan_info_vec.vlan_params.interface_name` |
| `ibm_backup_recovery_source_registration_invalid_kubernetes_params__vlan_info_vec__vlan_params__vlan_id` | `kubernetes_params.vlan_info_vec.vlan_params.vlan_id` |
| `ibm_backup_recovery_source_registration_invalid_physical_params__applications` | `physical_params.applications` |
| `ibm_backup_recovery_source_registration_invalid_physical_params__endpoint` | `physical_params.endpoint` |
| `ibm_backup_recovery_source_registration_invalid_physical_params__force_register` | `physical_params.force_register` |
| `ibm_backup_recovery_source_registration_invalid_physical_params__host_type` | `physical_params.host_type` |
| `ibm_backup_recovery_source_registration_invalid_physical_params__physical_type` | `physical_params.physical_type` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_source_registration" "test" {
  environment = "kPhysical"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
