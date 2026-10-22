# ibm_en_destination_app_configuration

Manages ibm_en_destination_app_configuration. The Destintion name.

This rule validates `ibm_en_destination_app_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_app_configuration" {
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
| `config` (block) | yes | nested block |
| `config.params` (block) | yes | nested block |
| `config.params.crn` | yes | crn format |
| `config.params.environment_id` | yes | API: Environment ID of App Configuration |
| `config.params.feature_id` | yes | API: Feature ID of App Configuration |
| `config.params.type` | yes | one of: `features` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_app_configuration_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_app_configuration_invalid_description` | `description` |
| `ibm_en_destination_app_configuration_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_app_configuration_invalid_name` | `name` |
| `ibm_en_destination_app_configuration_invalid_type` | `type` |
| `ibm_en_destination_app_configuration_invalid_config__params__crn` | `config.params.crn` |
| `ibm_en_destination_app_configuration_invalid_config__params__environment_id` | `config.params.environment_id` |
| `ibm_en_destination_app_configuration_invalid_config__params__feature_id` | `config.params.feature_id` |
| `ibm_en_destination_app_configuration_invalid_config__params__type` | `config.params.type` |

## Example

### Valid

```hcl
resource "ibm_en_destination_app_configuration" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
  config {
    params {
      crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
      environment_id = "01234567-89ab-cdef-0123-456789abcdef"
      feature_id = "01234567-89ab-cdef-0123-456789abcdef"
      type = "features"
    }
  }
}
```
