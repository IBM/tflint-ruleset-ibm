# ibm_backup_recovery_search_protected_objects_data

Schema validation for the ibm_backup_recovery_search_protected_objects data source.

This rule validates `ibm_backup_recovery_search_protected_objects` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_search_protected_objects_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cdp_protected_only` | no | — |
| `environments` | no | — |
| `filter_snapshot_from_usecs` | no | — |
| `filter_snapshot_to_usecs` | no | — |
| `metadata` | no | — |
| `num_results` | no | — |
| `object_action_key` | no | — |
| `object_ids` | no | — |
| `objects` | no | — |
| `os_types` | no | — |
| `protection_group_ids` | no | — |
| `request_initiator_type` | no | — |
| `run_instance_ids` | no | — |
| `search_string` | no | — |
| `snapshot_actions` | no | — |
| `source_ids` | no | — |
| `sub_result_size` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_search_protected_objects_data_invalid_cdp_protected_only` | `cdp_protected_only` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_environments` | `environments` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_filter_snapshot_from_usecs` | `filter_snapshot_from_usecs` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_filter_snapshot_to_usecs` | `filter_snapshot_to_usecs` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_object_action_key` | `object_action_key` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_object_ids` | `object_ids` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_os_types` | `os_types` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_protection_group_ids` | `protection_group_ids` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_run_instance_ids` | `run_instance_ids` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_search_string` | `search_string` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_snapshot_actions` | `snapshot_actions` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_source_ids` | `source_ids` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_sub_result_size` | `sub_result_size` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_search_protected_objects_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_search_protected_objects" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
