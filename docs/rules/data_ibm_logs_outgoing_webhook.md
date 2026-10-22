# ibm_logs_outgoing_webhook_data

Reads ibm_logs_outgoing_webhook. The name of the Outbound Integration.

This rule validates `ibm_logs_outgoing_webhook` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_outgoing_webhook_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `external_id` | no | — |
| `ibm_event_notifications` | no | — |
| `logs_outgoing_webhook_id` | yes | — |
| `name` | no | — |
| `type` | no | — |
| `updated_at` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_outgoing_webhook_data_invalid_logs_outgoing_webhook_id` | `logs_outgoing_webhook_id` |

## Example

### Valid

```hcl
data "ibm_logs_outgoing_webhook" "test" {
  logs_outgoing_webhook_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
