# ibm_en_smtp_setting

Schema validation for the ibm_en_smtp_setting resource.

This rule validates `ibm_en_smtp_setting` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SMTPAllowedIPs` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_setting" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `smtp_config_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `settings` (block) | yes | nested block |
| `settings.subnets` | yes | API: The SMTP allowed Ips |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_setting_invalid_instance_guid` | `instance_guid` |
| `ibm_en_smtp_setting_invalid_smtp_config_id` | `smtp_config_id` |
| `ibm_en_smtp_setting_invalid_settings__subnets` | `settings.subnets` |

## Example

### Valid

```hcl
resource "ibm_en_smtp_setting" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  smtp_config_id = "01234567-89ab-cdef-0123-456789abcdef"
  settings {
    subnets = "aaaaaaa"
  }
}
```
