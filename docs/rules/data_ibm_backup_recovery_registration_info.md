# ibm_backup_recovery_registration_info_data

Schema validation for the ibm_backup_recovery_registration_info data source.

This rule validates `ibm_backup_recovery_registration_info` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_registration_info_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `all_under_hierarchy` | no | — |
| `encryption_key` | no | — |
| `environments` | no | — |
| `ids` | no | — |
| `include_applications_tree_info` | no | — |
| `include_entity_permission_info` | no | — |
| `include_external_metadata` | no | — |
| `include_source_credentials` | no | — |
| `maintenance_status` | no | — |
| `prune_non_critical_info` | no | — |
| `request_initiator_type` | no | — |
| `root_nodes` | no | — |
| `sids` | no | — |
| `stats` | no | — |
| `stats_by_env` | no | — |
| `tenant_ids` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_registration_info_data_invalid_all_under_hierarchy` | `all_under_hierarchy` |
| `ibm_backup_recovery_registration_info_data_invalid_encryption_key` | `encryption_key` |
| `ibm_backup_recovery_registration_info_data_invalid_environments` | `environments` |
| `ibm_backup_recovery_registration_info_data_invalid_ids` | `ids` |
| `ibm_backup_recovery_registration_info_data_invalid_include_applications_tree_info` | `include_applications_tree_info` |
| `ibm_backup_recovery_registration_info_data_invalid_include_entity_permission_info` | `include_entity_permission_info` |
| `ibm_backup_recovery_registration_info_data_invalid_include_external_metadata` | `include_external_metadata` |
| `ibm_backup_recovery_registration_info_data_invalid_include_source_credentials` | `include_source_credentials` |
| `ibm_backup_recovery_registration_info_data_invalid_maintenance_status` | `maintenance_status` |
| `ibm_backup_recovery_registration_info_data_invalid_prune_non_critical_info` | `prune_non_critical_info` |
| `ibm_backup_recovery_registration_info_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_registration_info_data_invalid_sids` | `sids` |
| `ibm_backup_recovery_registration_info_data_invalid_tenant_ids` | `tenant_ids` |
| `ibm_backup_recovery_registration_info_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_registration_info_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_registration_info" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
