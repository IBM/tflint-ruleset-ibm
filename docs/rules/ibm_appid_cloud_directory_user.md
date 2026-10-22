# ibm_appid_cloud_directory_user

Schema validation for the ibm_appid_cloud_directory_user resource.

This rule validates `ibm_appid_cloud_directory_user` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/cloud_directory/Users/{userId}:put` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_cloud_directory_user" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | no | API: The status of the user. This parameter determines whether the user account is active or not. If both parameters are passed (active: false and lockedUntil) - the user account is locked until the time that is configured. If only (active: false) is passed, the user account is locked forever (or until it changes again). |
| `create_profile` | no | — |
| `display_name` | no | — |
| `locked_until` | no | — |
| `meta` | no | — |
| `password` | yes | API: The password. |
| `status` | no | one of: `PENDING`, `CONFIRMED` |
| `subject` | no | — |
| `tenant_id` | yes | — |
| `user_id` | no | — |
| `user_name` | no | — |
| `email` (block) | yes | nested block |
| `email.primary` | no | API: The email is the user's primary email. |
| `email.value` | yes | API: The email. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_cloud_directory_user_invalid_active` | `active` |
| `ibm_appid_cloud_directory_user_invalid_create_profile` | `create_profile` |
| `ibm_appid_cloud_directory_user_invalid_display_name` | `display_name` |
| `ibm_appid_cloud_directory_user_invalid_locked_until` | `locked_until` |
| `ibm_appid_cloud_directory_user_invalid_password` | `password` |
| `ibm_appid_cloud_directory_user_invalid_status` | `status` |
| `ibm_appid_cloud_directory_user_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_cloud_directory_user_invalid_user_name` | `user_name` |
| `ibm_appid_cloud_directory_user_invalid_email__primary` | `email.primary` |
| `ibm_appid_cloud_directory_user_invalid_email__value` | `email.value` |

## Example

### Valid

```hcl
variable "password" {
  sensitive = true
}

resource "ibm_appid_cloud_directory_user" "test" {
  password = var.password
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  email {
    value = "example"
  }
}
```
