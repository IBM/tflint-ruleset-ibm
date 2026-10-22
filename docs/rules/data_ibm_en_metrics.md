# ibm_en_metrics_data

Schema validation for the ibm_en_metrics data source.

This rule validates `ibm_en_metrics` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Metrics` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_metrics_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `destination_id` | no | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `destination_type` | no | — |
| `email_to` | no | — |
| `gte` | yes | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `lte` | yes | — |
| `metrics` | no | — |
| `notification_id` | no | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `smtp_config_id` | no | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `source_id` | no | — |
| `subject` | no | API: The subject of the notification |
| `subscription_id` | no | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_metrics_data_invalid_destination_id` | `destination_id` |
| `ibm_en_metrics_data_invalid_destination_type` | `destination_type` |
| `ibm_en_metrics_data_invalid_email_to` | `email_to` |
| `ibm_en_metrics_data_invalid_gte` | `gte` |
| `ibm_en_metrics_data_invalid_instance_id` | `instance_id` |
| `ibm_en_metrics_data_invalid_lte` | `lte` |
| `ibm_en_metrics_data_invalid_notification_id` | `notification_id` |
| `ibm_en_metrics_data_invalid_smtp_config_id` | `smtp_config_id` |
| `ibm_en_metrics_data_invalid_source_id` | `source_id` |
| `ibm_en_metrics_data_invalid_subject` | `subject` |
| `ibm_en_metrics_data_invalid_subscription_id` | `subscription_id` |

## Example

### Valid

```hcl
data "ibm_en_metrics" "test" {
  gte = "example"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  lte = "example"
}
```
