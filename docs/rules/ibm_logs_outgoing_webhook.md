# ibm_logs_outgoing_webhook

Manages ibm_logs_outgoing_webhook. The name of the Outbound Integration.

This rule validates `ibm_logs_outgoing_webhook` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `OutgoingWebhooksV1IbmEventNotificationsConfig` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_outgoing_webhook" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `external_id` | no | — |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `type` | yes | one of: `ibm_event_notifications` |
| `updated_at` | no | — |
| `url` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `webhook_id` | no | — |
| `ibm_event_notifications` (block) | no | nested block |
| `ibm_event_notifications.endpoint_type` | no | one of: `default_or_public`, `private` |
| `ibm_event_notifications.event_notifications_instance_id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `ibm_event_notifications.region_id` | yes | pattern `^[a-z]{2}-[a-z]+$` |
| `ibm_event_notifications.source_id` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `ibm_event_notifications.source_name` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_outgoing_webhook_invalid_name` | `name` |
| `ibm_logs_outgoing_webhook_invalid_type` | `type` |
| `ibm_logs_outgoing_webhook_invalid_url` | `url` |
| `ibm_logs_outgoing_webhook_invalid_ibm_event_notifications__endpoint_type` | `ibm_event_notifications.endpoint_type` |
| `ibm_logs_outgoing_webhook_invalid_ibm_event_notifications__event_notifications_instance_id` | `ibm_event_notifications.event_notifications_instance_id` |
| `ibm_logs_outgoing_webhook_invalid_ibm_event_notifications__region_id` | `ibm_event_notifications.region_id` |
| `ibm_logs_outgoing_webhook_invalid_ibm_event_notifications__source_id` | `ibm_event_notifications.source_id` |
| `ibm_logs_outgoing_webhook_invalid_ibm_event_notifications__source_name` | `ibm_event_notifications.source_name` |

## Example

### Valid

```hcl
resource "ibm_logs_outgoing_webhook" "test" {
  name = "example-name"
  type = "ibm_event_notifications"
}
```
