# ibm_en_destination_slack

Manages ibm_en_destination_slack. The Destintion name.

This rule validates `ibm_en_destination_slack` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_slack" {
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
| `type` | yes | one of: `webhook`, `smtp_ibm`, `sms_ibm`, `push_android`, `push_ios`, `slack`, `ibmce`, `pagerduty`, `push_safari`, `msteams`, `servicenow`, `ibmcos`, `push_huawei`, `smtp_custom`, `smtp_custom_sandbox`, `sms_custom`, `event_streams`, `app_configuration` |
| `updated_at` | no | — |
| `config` (block) | no | nested block |
| `config.params` (block) | no | nested block |
| `config.params.token` | no | API: Token of slack application |
| `config.params.type` | yes | one of: `features` |
| `config.params.url` | no | pattern `^(?:https?://)?(?:[^@/
]+@)?(?:www.)?([^:/
]+)` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_slack_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_slack_invalid_description` | `description` |
| `ibm_en_destination_slack_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_slack_invalid_name` | `name` |
| `ibm_en_destination_slack_invalid_type` | `type` |
| `ibm_en_destination_slack_invalid_config__params__token` | `config.params.token` |
| `ibm_en_destination_slack_invalid_config__params__type` | `config.params.type` |
| `ibm_en_destination_slack_invalid_config__params__url` | `config.params.url` |

## Example

### Valid

```hcl
resource "ibm_en_destination_slack" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
}
```
