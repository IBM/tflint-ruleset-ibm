# ibm_en_destination_safari

Manages ibm_en_destination_safari. The Destintion name.

This rule validates `ibm_en_destination_safari` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DestinationCreate` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_safari" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | yes | API: Certificate for APNS |
| `collect_failed_events` | no | API: Whether to collect the failed event in Cloud Object Storage bucket |
| `description` | no | API: The Destination description |
| `destination_id` | no | — |
| `icon_128x128` | no | API: Safari icon 128x128 |
| `icon_128x128_2x` | no | API: Safari icon 128x128@2x |
| `icon_128x128_2x_content_type` | no | — |
| `icon_128x128_content_type` | no | — |
| `icon_16x16` | no | API: Safari icon 16x16 |
| `icon_16x16_2x` | no | API: Safari icon 16x16@2x |
| `icon_16x16_2x_content_type` | no | — |
| `icon_16x16_content_type` | no | — |
| `icon_32x32` | no | API: Safari icon 32x32 |
| `icon_32x32_2x` | no | API: Safari icon 32x32@2x |
| `icon_32x32_2x_content_type` | no | — |
| `icon_32x32_content_type` | no | — |
| `instance_guid` | yes | — |
| `name` | yes | API: The Destination name |
| `subscription_count` | no | — |
| `subscription_names` | no | — |
| `type` | yes | one of: `webhook`, `push_android`, `push_ios`, `push_chrome`, `push_firefox`, `slack`, `ibmce`, `pagerduty`, `push_safari`, `msteams`, `servicenow`, `ibmcos`, `push_huawei`, `smtp_custom`, `smtp_custom_sandbox`, `sms_custom`, `event_streams`, `app_configuration` |
| `updated_at` | no | — |
| `config` (block) | no | nested block |
| `config.params` (block) | no | nested block |
| `config.params.cert_type` | yes | API: Authentication type p12 |
| `config.params.password` | yes | API: Password for ServiceNow account REST API |
| `config.params.pre_prod` | no | API: If pre prod enabled |
| `config.params.url_format_string` | no | API: Website url |
| `config.params.website_name` | no | pattern `^(?:https?://)?(?:[^@/
]+@)?(?:www.)?([^:/
]+)` |
| `config.params.website_push_id` | no | API: Website url |
| `config.params.website_url` | no | pattern `^(?:https?://)?(?:[^@/
]+@)?(?:www.)?([^:/
]+)` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_safari_invalid_certificate` | `certificate` |
| `ibm_en_destination_safari_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_safari_invalid_description` | `description` |
| `ibm_en_destination_safari_invalid_icon_128x128` | `icon_128x128` |
| `ibm_en_destination_safari_invalid_icon_128x128_2x` | `icon_128x128_2x` |
| `ibm_en_destination_safari_invalid_icon_128x128_2x_content_type` | `icon_128x128_2x_content_type` |
| `ibm_en_destination_safari_invalid_icon_128x128_content_type` | `icon_128x128_content_type` |
| `ibm_en_destination_safari_invalid_icon_16x16` | `icon_16x16` |
| `ibm_en_destination_safari_invalid_icon_16x16_2x` | `icon_16x16_2x` |
| `ibm_en_destination_safari_invalid_icon_16x16_2x_content_type` | `icon_16x16_2x_content_type` |
| `ibm_en_destination_safari_invalid_icon_16x16_content_type` | `icon_16x16_content_type` |
| `ibm_en_destination_safari_invalid_icon_32x32` | `icon_32x32` |
| `ibm_en_destination_safari_invalid_icon_32x32_2x` | `icon_32x32_2x` |
| `ibm_en_destination_safari_invalid_icon_32x32_2x_content_type` | `icon_32x32_2x_content_type` |
| `ibm_en_destination_safari_invalid_icon_32x32_content_type` | `icon_32x32_content_type` |
| `ibm_en_destination_safari_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_safari_invalid_name` | `name` |
| `ibm_en_destination_safari_invalid_type` | `type` |
| `ibm_en_destination_safari_invalid_config__params__cert_type` | `config.params.cert_type` |
| `ibm_en_destination_safari_invalid_config__params__password` | `config.params.password` |
| `ibm_en_destination_safari_invalid_config__params__pre_prod` | `config.params.pre_prod` |
| `ibm_en_destination_safari_invalid_config__params__url_format_string` | `config.params.url_format_string` |
| `ibm_en_destination_safari_invalid_config__params__website_name` | `config.params.website_name` |
| `ibm_en_destination_safari_invalid_config__params__website_push_id` | `config.params.website_push_id` |
| `ibm_en_destination_safari_invalid_config__params__website_url` | `config.params.website_url` |

## Example

### Valid

```hcl
resource "ibm_en_destination_safari" "test" {
  certificate = "aaaaaaa"
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "webhook"
}
```
