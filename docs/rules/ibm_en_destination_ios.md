# ibm_en_destination_ios

Manages ibm_en_destination_ios. The Destintion name.

This rule validates `ibm_en_destination_ios` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_ios" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | yes | API: Certificate for APNS |
| `certificate_content_type` | yes | — |
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
| `config.params.bundle_id` | no | API: Bundle ID for token (Required when cert_type is p8) |
| `config.params.cert_type` | yes | API: Authentication type p12 |
| `config.params.is_sandbox` | yes | API: Sandbox mode for iOS destinations |
| `config.params.key_id` | no | API: Key ID for token (Required when cert_type is p8) |
| `config.params.password` | no | API: Password for ServiceNow account REST API |
| `config.params.pre_prod` | no | API: If pre prod enabled |
| `config.params.team_id` | no | API: Team ID for token (Required when cert_type is p8) |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_ios_invalid_certificate` | `certificate` |
| `ibm_en_destination_ios_invalid_certificate_content_type` | `certificate_content_type` |
| `ibm_en_destination_ios_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_ios_invalid_description` | `description` |
| `ibm_en_destination_ios_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_ios_invalid_name` | `name` |
| `ibm_en_destination_ios_invalid_type` | `type` |
| `ibm_en_destination_ios_invalid_config__params__bundle_id` | `config.params.bundle_id` |
| `ibm_en_destination_ios_invalid_config__params__cert_type` | `config.params.cert_type` |
| `ibm_en_destination_ios_invalid_config__params__is_sandbox` | `config.params.is_sandbox` |
| `ibm_en_destination_ios_invalid_config__params__key_id` | `config.params.key_id` |
| `ibm_en_destination_ios_invalid_config__params__password` | `config.params.password` |
| `ibm_en_destination_ios_invalid_config__params__pre_prod` | `config.params.pre_prod` |
| `ibm_en_destination_ios_invalid_config__params__team_id` | `config.params.team_id` |

## Example

### Valid

```hcl
resource "ibm_en_destination_ios" "test" {
  certificate = "aaaaaaa"
  certificate_content_type = "example"
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
}
```
