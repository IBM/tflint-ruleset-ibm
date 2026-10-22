# ibm_backup_recovery_manager_get_alerts_resolution_data

Schema validation for the ibm_backup_recovery_manager_get_alerts_resolution data source.

This rule validates `ibm_backup_recovery_manager_get_alerts_resolution` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_alerts_resolution_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_resolutions_list` | no | — |
| `max_resolutions` | yes | — |
| `resolution_id` | no | — |
| `resolution_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_alerts_resolution_data_invalid_max_resolutions` | `max_resolutions` |
| `ibm_backup_recovery_manager_get_alerts_resolution_data_invalid_resolution_id` | `resolution_id` |
| `ibm_backup_recovery_manager_get_alerts_resolution_data_invalid_resolution_name` | `resolution_name` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_alerts_resolution" "test" {
  max_resolutions = 1
}
```
