# ibm_en_destination_webhook

Manages ibm_en_destination_webhook. The Destintion name.

This rule validates `ibm_en_destination_webhook` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_webhook" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collect_failed_events` | no | API: Whether to collect the failed event in Cloud Object Storage bucket |
| `description` | no | API: Destination description |
| `destination_id` | no | — |
| `instance_guid` | yes | — |
| `name` | yes | API: Destination name |
| `subscription_count` | no | — |
| `subscription_names` | no | — |
| `test_destination` | no | — |
| `test_status` | no | — |
| `type` | yes | one of: `webhook`, `smtp_ibm`, `sms_ibm`, `push_android`, `push_ios`, `slack`, `ibmce`, `pagerduty`, `push_safari`, `msteams`, `servicenow`, `ibmcos`, `push_huawei`, `smtp_custom`, `smtp_custom_sandbox`, `sms_custom`, `event_streams`, `app_configuration` |
| `updated_at` | no | — |
| `config` (block) | no | nested block |
| `config.params` (block) | no | nested block |
| `config.params.custom_headers` | no | API: Custom headers (Key-Value pair) for webhook call |
| `config.params.sensitive_headers` | no | API: List of sensitive headers from custom headers |
| `config.params.url` | yes | pattern `^(?:https?://)?(?:[^@/
]+@)?(?:www.)?([^:/
]+)` |
| `config.params.verb` | yes | one of: `get`, `post` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_webhook_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_webhook_invalid_description` | `description` |
| `ibm_en_destination_webhook_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_webhook_invalid_name` | `name` |
| `ibm_en_destination_webhook_invalid_test_destination` | `test_destination` |
| `ibm_en_destination_webhook_invalid_type` | `type` |
| `ibm_en_destination_webhook_invalid_config__params__custom_headers` | `config.params.custom_headers` |
| `ibm_en_destination_webhook_invalid_config__params__sensitive_headers` | `config.params.sensitive_headers` |
| `ibm_en_destination_webhook_invalid_config__params__url` | `config.params.url` |
| `ibm_en_destination_webhook_invalid_config__params__verb` | `config.params.verb` |

## Example

### Valid

```hcl
resource "ibm_en_destination_webhook" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
}
```
