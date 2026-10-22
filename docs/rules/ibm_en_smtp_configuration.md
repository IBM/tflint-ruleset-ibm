# ibm_en_smtp_configuration

Manages ibm_en_smtp_configuration. SMTP name.

This rule validates `ibm_en_smtp_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SMTPCreate` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config` | no | — |
| `description` | no | API: The description of SMTP configuration |
| `domain` | yes | API: Domain Name |
| `en_smtp_configuration_id` | no | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `name` | yes | API: The name of SMTP configuration |
| `updated_at` | no | — |
| `verification_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_configuration_invalid_description` | `description` |
| `ibm_en_smtp_configuration_invalid_domain` | `domain` |
| `ibm_en_smtp_configuration_invalid_instance_id` | `instance_id` |
| `ibm_en_smtp_configuration_invalid_name` | `name` |
| `ibm_en_smtp_configuration_invalid_verification_type` | `verification_type` |

## Example

### Valid

```hcl
resource "ibm_en_smtp_configuration" "test" {
  domain = "aaaaaaa"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "aaaaaaa"
}
```
