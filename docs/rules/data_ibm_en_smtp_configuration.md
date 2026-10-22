# ibm_en_smtp_configuration_data

Reads ibm_en_smtp_configuration. SMTP name.

This rule validates `ibm_en_smtp_configuration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IBMCloudObjectStorageDestinationConfig` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_configuration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config` | no | — |
| `description` | no | — |
| `domain` | no | — |
| `en_smtp_configuration_id` | yes | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `name` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_configuration_data_invalid_en_smtp_configuration_id` | `en_smtp_configuration_id` |
| `ibm_en_smtp_configuration_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_en_smtp_configuration" "test" {
  en_smtp_configuration_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
