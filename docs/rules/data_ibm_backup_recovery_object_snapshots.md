# ibm_backup_recovery_object_snapshots_data

Schema validation for the ibm_backup_recovery_object_snapshots data source.

This rule validates `ibm_backup_recovery_object_snapshots` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_object_snapshots_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `from_time_usecs` | no | — |
| `object_action_keys` | no | — |
| `object_id` | yes | — |
| `protection_group_ids` | no | — |
| `region_ids` | no | — |
| `run_instance_ids` | no | — |
| `run_start_from_time_usecs` | no | — |
| `run_start_to_time_usecs` | no | — |
| `run_types` | no | — |
| `snapshot_actions` | no | — |
| `snapshots` | no | — |
| `to_time_usecs` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_object_snapshots_data_invalid_from_time_usecs` | `from_time_usecs` |
| `ibm_backup_recovery_object_snapshots_data_invalid_object_action_keys` | `object_action_keys` |
| `ibm_backup_recovery_object_snapshots_data_invalid_object_id` | `object_id` |
| `ibm_backup_recovery_object_snapshots_data_invalid_protection_group_ids` | `protection_group_ids` |
| `ibm_backup_recovery_object_snapshots_data_invalid_region_ids` | `region_ids` |
| `ibm_backup_recovery_object_snapshots_data_invalid_run_instance_ids` | `run_instance_ids` |
| `ibm_backup_recovery_object_snapshots_data_invalid_run_start_from_time_usecs` | `run_start_from_time_usecs` |
| `ibm_backup_recovery_object_snapshots_data_invalid_run_start_to_time_usecs` | `run_start_to_time_usecs` |
| `ibm_backup_recovery_object_snapshots_data_invalid_run_types` | `run_types` |
| `ibm_backup_recovery_object_snapshots_data_invalid_snapshot_actions` | `snapshot_actions` |
| `ibm_backup_recovery_object_snapshots_data_invalid_to_time_usecs` | `to_time_usecs` |
| `ibm_backup_recovery_object_snapshots_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_object_snapshots" "test" {
  object_id = 1
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
