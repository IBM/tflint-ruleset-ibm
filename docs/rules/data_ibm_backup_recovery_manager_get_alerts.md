# ibm_backup_recovery_manager_get_alerts_data

Schema validation for the ibm_backup_recovery_manager_get_alerts data source.

This rule validates `ibm_backup_recovery_manager_get_alerts` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_alerts_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_categories` | no | — |
| `alert_ids` | no | — |
| `alert_name` | no | — |
| `alert_severities` | no | — |
| `alert_states` | no | — |
| `alert_type_buckets` | no | — |
| `alert_types` | no | — |
| `alerts` | no | — |
| `all_under_hierarchy` | no | — |
| `end_time_usecs` | no | — |
| `max_alerts` | no | — |
| `property_key` | no | — |
| `property_value` | no | — |
| `resolution_ids` | no | — |
| `start_time_usecs` | no | — |
| `tenant_ids` | no | — |
| `x_scope_identifier` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_categories` | `alert_categories` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_ids` | `alert_ids` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_name` | `alert_name` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_severities` | `alert_severities` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_states` | `alert_states` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_type_buckets` | `alert_type_buckets` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_alert_types` | `alert_types` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_all_under_hierarchy` | `all_under_hierarchy` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_max_alerts` | `max_alerts` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_property_key` | `property_key` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_property_value` | `property_value` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_resolution_ids` | `resolution_ids` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_tenant_ids` | `tenant_ids` |
| `ibm_backup_recovery_manager_get_alerts_data_invalid_x_scope_identifier` | `x_scope_identifier` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_alerts" "test" {
}
```
