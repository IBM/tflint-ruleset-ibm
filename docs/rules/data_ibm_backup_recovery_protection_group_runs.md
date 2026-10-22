# ibm_backup_recovery_protection_group_runs_data

Schema validation for the ibm_backup_recovery_protection_group_runs data source.

This rule validates `ibm_backup_recovery_protection_group_runs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_group_runs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `archival_run_status` | no | — |
| `cloud_spin_run_status` | no | — |
| `end_time_usecs` | no | — |
| `exclude_non_restorable_runs` | no | — |
| `filter_by_copy_task_end_time` | no | — |
| `filter_by_end_time` | no | — |
| `include_object_details` | no | — |
| `local_backup_run_status` | no | — |
| `num_runs` | no | — |
| `only_return_successful_copy_run` | no | — |
| `protection_group_id` | yes | — |
| `replication_run_status` | no | — |
| `request_initiator_type` | no | — |
| `run_id` | no | — |
| `run_tags` | no | — |
| `run_types` | no | — |
| `runs` | no | — |
| `snapshot_target_types` | no | — |
| `start_time_usecs` | no | — |
| `total_runs` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_group_runs_data_invalid_archival_run_status` | `archival_run_status` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_cloud_spin_run_status` | `cloud_spin_run_status` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_exclude_non_restorable_runs` | `exclude_non_restorable_runs` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_filter_by_copy_task_end_time` | `filter_by_copy_task_end_time` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_filter_by_end_time` | `filter_by_end_time` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_include_object_details` | `include_object_details` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_local_backup_run_status` | `local_backup_run_status` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_num_runs` | `num_runs` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_only_return_successful_copy_run` | `only_return_successful_copy_run` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_protection_group_id` | `protection_group_id` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_replication_run_status` | `replication_run_status` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_run_id` | `run_id` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_run_tags` | `run_tags` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_run_types` | `run_types` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_snapshot_target_types` | `snapshot_target_types` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_protection_group_runs_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_group_runs" "test" {
  protection_group_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
