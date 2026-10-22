# ibm_backup_recovery_manager_get_alerts_summary_data

Schema validation for the ibm_backup_recovery_manager_get_alerts_summary data source.

This rule validates `ibm_backup_recovery_manager_get_alerts_summary` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_alerts_summary_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alerts_summary` | no | — |
| `end_time_usecs` | no | — |
| `include_tenants` | no | — |
| `start_time_usecs` | no | — |
| `states_list` | no | — |
| `tenant_ids` | no | — |
| `x_scope_identifier` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_include_tenants` | `include_tenants` |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_states_list` | `states_list` |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_tenant_ids` | `tenant_ids` |
| `ibm_backup_recovery_manager_get_alerts_summary_data_invalid_x_scope_identifier` | `x_scope_identifier` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_alerts_summary" "test" {
}
```
