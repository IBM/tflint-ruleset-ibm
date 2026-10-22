# ibm_backup_recovery_manager_get_management_alerts_data

Schema validation for the ibm_backup_recovery_manager_get_management_alerts data source.

This rule validates `ibm_backup_recovery_manager_get_management_alerts` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_management_alerts_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_category_list` | no | — |
| `alert_id_list` | no | — |
| `alert_name` | no | — |
| `alert_property_key_list` | no | — |
| `alert_property_value_list` | no | — |
| `alert_severity_list` | no | — |
| `alert_state_list` | no | — |
| `alert_type_bucket_list` | no | — |
| `alert_type_list` | no | — |
| `alerts_list` | no | — |
| `cluster_identifiers` | no | — |
| `end_date_usecs` | no | — |
| `max_alerts` | no | — |
| `region_ids` | no | — |
| `service_instance_ids` | no | — |
| `start_date_usecs` | no | — |
| `tenant_ids` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_category_list` | `alert_category_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_id_list` | `alert_id_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_name` | `alert_name` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_property_key_list` | `alert_property_key_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_property_value_list` | `alert_property_value_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_severity_list` | `alert_severity_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_state_list` | `alert_state_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_type_bucket_list` | `alert_type_bucket_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_alert_type_list` | `alert_type_list` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_cluster_identifiers` | `cluster_identifiers` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_end_date_usecs` | `end_date_usecs` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_max_alerts` | `max_alerts` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_region_ids` | `region_ids` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_service_instance_ids` | `service_instance_ids` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_start_date_usecs` | `start_date_usecs` |
| `ibm_backup_recovery_manager_get_management_alerts_data_invalid_tenant_ids` | `tenant_ids` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_management_alerts" "test" {
}
```
