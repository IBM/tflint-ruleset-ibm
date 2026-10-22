# ibm_logs_alert_definition_data

Reads ibm_logs_alert_definition. The name of the alert definition.

This rule validates `ibm_logs_alert_definition` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_alert_definition_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active_on` | no | — |
| `alert_version_id` | no | — |
| `created_time` | no | — |
| `deleted` | no | — |
| `description` | no | — |
| `enabled` | no | — |
| `entity_labels` | no | — |
| `flow` | no | — |
| `group_by_keys` | no | — |
| `incidents_settings` | no | — |
| `logs_alert_definition_id` | yes | — |
| `logs_anomaly` | no | — |
| `logs_immediate` | no | — |
| `logs_new_value` | no | — |
| `logs_ratio_threshold` | no | — |
| `logs_threshold` | no | — |
| `logs_time_relative_threshold` | no | — |
| `logs_unique_count` | no | — |
| `metric_anomaly` | no | — |
| `metric_threshold` | no | — |
| `name` | no | — |
| `notification_group` | no | — |
| `phantom_mode` | no | — |
| `priority` | no | — |
| `type` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_alert_definition_data_invalid_logs_alert_definition_id` | `logs_alert_definition_id` |

## Example

### Valid

```hcl
data "ibm_logs_alert_definition" "test" {
  logs_alert_definition_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
