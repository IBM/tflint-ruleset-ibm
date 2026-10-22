# ibm_backup_recovery_manager_get_alerts_stats_data

Schema validation for the ibm_backup_recovery_manager_get_alerts_stats data source.

This rule validates `ibm_backup_recovery_manager_get_alerts_stats` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_alerts_stats_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `aggregated_alerts_stats` | no | — |
| `aggregated_cluster_stats` | no | — |
| `alert_source` | no | — |
| `cluster_ids` | no | — |
| `end_time_usecs` | yes | — |
| `exclude_stats_by_cluster` | no | — |
| `region_ids` | no | — |
| `service_instance_ids` | no | — |
| `start_time_usecs` | yes | — |
| `stats_by_cluster` | no | — |
| `tenant_ids` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_alert_source` | `alert_source` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_cluster_ids` | `cluster_ids` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_exclude_stats_by_cluster` | `exclude_stats_by_cluster` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_region_ids` | `region_ids` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_service_instance_ids` | `service_instance_ids` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_stats_data_invalid_tenant_ids` | `tenant_ids` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_alerts_stats" "test" {
  end_time_usecs = 1
  start_time_usecs = 1
}
```
