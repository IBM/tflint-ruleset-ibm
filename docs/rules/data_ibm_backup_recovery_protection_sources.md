# ibm_backup_recovery_protection_sources_data

Schema validation for the ibm_backup_recovery_protection_sources data source.

This rule validates `ibm_backup_recovery_protection_sources` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_protection_sources_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `after_cursor_entity_id` | no | — |
| `all_under_hierarchy` | no | — |
| `backup_recovery_protection_source_nodes_id` | no | — |
| `before_cursor_entity_id` | no | — |
| `encryption_key` | no | — |
| `environment` | no | one of: `kPhysical`, `kSQL` |
| `environments` | no | — |
| `exclude_aws_types` | no | — |
| `exclude_kubernetes_types` | no | — |
| `exclude_office365_types` | no | — |
| `exclude_types` | no | — |
| `get_teams_channels` | no | — |
| `has_valid_mailbox` | no | — |
| `has_valid_onedrive` | no | — |
| `include_datastores` | no | — |
| `include_entity_permission_info` | no | — |
| `include_networks` | no | — |
| `include_object_protection_info` | no | — |
| `include_sfdc_fields` | no | — |
| `include_source_credentials` | no | — |
| `include_system_v_apps` | no | — |
| `include_vm_folders` | no | — |
| `is_security_group` | no | — |
| `node_id` | no | — |
| `num_levels` | no | — |
| `page_size` | no | — |
| `protection_sources` | no | — |
| `prune_aggregation_info` | no | — |
| `prune_non_critical_info` | no | — |
| `request_initiator_type` | no | — |
| `sids` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_sources_data_invalid_after_cursor_entity_id` | `after_cursor_entity_id` |
| `ibm_backup_recovery_protection_sources_data_invalid_all_under_hierarchy` | `all_under_hierarchy` |
| `ibm_backup_recovery_protection_sources_data_invalid_backup_recovery_protection_source_nodes_id` | `backup_recovery_protection_source_nodes_id` |
| `ibm_backup_recovery_protection_sources_data_invalid_before_cursor_entity_id` | `before_cursor_entity_id` |
| `ibm_backup_recovery_protection_sources_data_invalid_encryption_key` | `encryption_key` |
| `ibm_backup_recovery_protection_sources_data_invalid_environment` | `environment` |
| `ibm_backup_recovery_protection_sources_data_invalid_environments` | `environments` |
| `ibm_backup_recovery_protection_sources_data_invalid_exclude_aws_types` | `exclude_aws_types` |
| `ibm_backup_recovery_protection_sources_data_invalid_exclude_kubernetes_types` | `exclude_kubernetes_types` |
| `ibm_backup_recovery_protection_sources_data_invalid_exclude_office365_types` | `exclude_office365_types` |
| `ibm_backup_recovery_protection_sources_data_invalid_exclude_types` | `exclude_types` |
| `ibm_backup_recovery_protection_sources_data_invalid_get_teams_channels` | `get_teams_channels` |
| `ibm_backup_recovery_protection_sources_data_invalid_has_valid_mailbox` | `has_valid_mailbox` |
| `ibm_backup_recovery_protection_sources_data_invalid_has_valid_onedrive` | `has_valid_onedrive` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_datastores` | `include_datastores` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_entity_permission_info` | `include_entity_permission_info` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_networks` | `include_networks` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_object_protection_info` | `include_object_protection_info` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_sfdc_fields` | `include_sfdc_fields` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_source_credentials` | `include_source_credentials` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_system_v_apps` | `include_system_v_apps` |
| `ibm_backup_recovery_protection_sources_data_invalid_include_vm_folders` | `include_vm_folders` |
| `ibm_backup_recovery_protection_sources_data_invalid_is_security_group` | `is_security_group` |
| `ibm_backup_recovery_protection_sources_data_invalid_node_id` | `node_id` |
| `ibm_backup_recovery_protection_sources_data_invalid_num_levels` | `num_levels` |
| `ibm_backup_recovery_protection_sources_data_invalid_page_size` | `page_size` |
| `ibm_backup_recovery_protection_sources_data_invalid_prune_aggregation_info` | `prune_aggregation_info` |
| `ibm_backup_recovery_protection_sources_data_invalid_prune_non_critical_info` | `prune_non_critical_info` |
| `ibm_backup_recovery_protection_sources_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_sources_data_invalid_sids` | `sids` |
| `ibm_backup_recovery_protection_sources_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_protection_sources_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_sources" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
