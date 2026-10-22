# ibm_en_destination_cos

Manages ibm_en_destination_cos. The Destintion name.

This rule validates `ibm_en_destination_cos` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_cos" {
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
| `config.params.bucket_name` | yes | API: Bucket Name for Cloud Object Storage |
| `config.params.endpoint` | no | API: Endpoint of Event Streams |
| `config.params.instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_cos_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_cos_invalid_description` | `description` |
| `ibm_en_destination_cos_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_cos_invalid_name` | `name` |
| `ibm_en_destination_cos_invalid_type` | `type` |
| `ibm_en_destination_cos_invalid_config__params__bucket_name` | `config.params.bucket_name` |
| `ibm_en_destination_cos_invalid_config__params__endpoint` | `config.params.endpoint` |
| `ibm_en_destination_cos_invalid_config__params__instance_id` | `config.params.instance_id` |

## Example

### Valid

```hcl
resource "ibm_en_destination_cos" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
}
```
