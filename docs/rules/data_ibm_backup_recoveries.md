# ibm_backup_recoveries_data

Schema validation for the ibm_backup_recoveries data source.

This rule validates `ibm_backup_recoveries` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recoveries_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `archival_target_type` | no | — |
| `end_time_usecs` | no | — |
| `ids` | no | — |
| `recoveries` | no | — |
| `recovery_actions` | no | — |
| `return_only_child_recoveries` | no | — |
| `snapshot_environments` | no | — |
| `snapshot_target_type` | no | — |
| `start_time_usecs` | no | — |
| `status` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recoveries_data_invalid_archival_target_type` | `archival_target_type` |
| `ibm_backup_recoveries_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recoveries_data_invalid_ids` | `ids` |
| `ibm_backup_recoveries_data_invalid_recovery_actions` | `recovery_actions` |
| `ibm_backup_recoveries_data_invalid_return_only_child_recoveries` | `return_only_child_recoveries` |
| `ibm_backup_recoveries_data_invalid_snapshot_environments` | `snapshot_environments` |
| `ibm_backup_recoveries_data_invalid_snapshot_target_type` | `snapshot_target_type` |
| `ibm_backup_recoveries_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recoveries_data_invalid_status` | `status` |
| `ibm_backup_recoveries_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recoveries" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
