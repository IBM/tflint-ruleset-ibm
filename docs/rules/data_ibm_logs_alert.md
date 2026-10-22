# ibm_logs_alert_data

Reads ibm_logs_alert. Alert name.

This rule validates `ibm_logs_alert` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_alert_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active_when` | no | — |
| `condition` | no | — |
| `description` | no | — |
| `expiration` | no | — |
| `filters` | no | — |
| `incident_settings` | no | — |
| `is_active` | no | — |
| `logs_alert_id` | yes | — |
| `meta_labels` | no | — |
| `meta_labels_strings` | no | — |
| `name` | no | — |
| `notification_groups` | no | — |
| `notification_payload_filters` | no | — |
| `severity` | no | — |
| `unique_identifier` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_alert_data_invalid_logs_alert_id` | `logs_alert_id` |

## Example

### Valid

```hcl
data "ibm_logs_alert" "test" {
  logs_alert_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
