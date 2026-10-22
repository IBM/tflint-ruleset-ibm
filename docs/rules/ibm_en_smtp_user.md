# ibm_en_smtp_user

Schema validation for the ibm_en_smtp_user resource.

This rule validates `ibm_en_smtp_user` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/instances/{instance_id}/smtp/config/{id}/users/{user_id}:patch` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_user" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `description` | no | API: SMTP user description |
| `domain` | no | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `password` | no | — |
| `smtp_config_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `updated_at` | no | — |
| `user_id` | no | — |
| `username` | no | — |
| `username_to_clone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_user_invalid_description` | `description` |
| `ibm_en_smtp_user_invalid_instance_id` | `instance_id` |
| `ibm_en_smtp_user_invalid_smtp_config_id` | `smtp_config_id` |
| `ibm_en_smtp_user_invalid_username_to_clone` | `username_to_clone` |

## Example

### Valid

```hcl
resource "ibm_en_smtp_user" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  smtp_config_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
