# ibm_en_smtp_user_data

Schema validation for the ibm_en_smtp_user data source.

This rule validates `ibm_en_smtp_user` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/instances/{instance_id}/smtp/config/{id}/users/{user_id}:patch` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_user_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `description` | no | — |
| `domain` | no | — |
| `en_smtp_config_id` | yes | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `smtp_config_id` | no | — |
| `updated_at` | no | — |
| `user_id` | yes | API: The user identifier for the device registration |
| `username` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_user_data_invalid_en_smtp_config_id` | `en_smtp_config_id` |
| `ibm_en_smtp_user_data_invalid_instance_id` | `instance_id` |
| `ibm_en_smtp_user_data_invalid_user_id` | `user_id` |

## Example

### Valid

```hcl
data "ibm_en_smtp_user" "test" {
  en_smtp_config_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  user_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
