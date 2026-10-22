# ibm_backup_recovery_manager_get_management_alerts_summary_data

Schema validation for the ibm_backup_recovery_manager_get_management_alerts_summary data source.

This rule validates `ibm_backup_recovery_manager_get_management_alerts_summary` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_management_alerts_summary_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alerts_summary` | no | — |
| `cluster_identifiers` | no | — |
| `end_time_usecs` | no | — |
| `start_time_usecs` | no | — |
| `states_list` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_management_alerts_summary_data_invalid_cluster_identifiers` | `cluster_identifiers` |
| `ibm_backup_recovery_manager_get_management_alerts_summary_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_manager_get_management_alerts_summary_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_manager_get_management_alerts_summary_data_invalid_states_list` | `states_list` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_management_alerts_summary" "test" {
}
```
