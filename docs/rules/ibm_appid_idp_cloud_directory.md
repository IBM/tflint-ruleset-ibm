# ibm_appid_idp_cloud_directory

Schema validation for the ibm_appid_idp_cloud_directory resource.

This rule validates `ibm_appid_idp_cloud_directory` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_idp_cloud_directory" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `identity_confirm_access_mode` | no | one of: `FULL`, `RESTRICTIVE`, `OFF` |
| `identity_confirm_methods` | no | — |
| `identity_field` | no | — |
| `is_active` | yes | — |
| `reset_password_enabled` | no | — |
| `reset_password_notification_enabled` | no | — |
| `self_service_enabled` | no | — |
| `signup_enabled` | no | — |
| `tenant_id` | yes | — |
| `welcome_enabled` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_idp_cloud_directory_invalid_identity_confirm_access_mode` | `identity_confirm_access_mode` |
| `ibm_appid_idp_cloud_directory_invalid_identity_confirm_methods` | `identity_confirm_methods` |
| `ibm_appid_idp_cloud_directory_invalid_identity_field` | `identity_field` |
| `ibm_appid_idp_cloud_directory_invalid_is_active` | `is_active` |
| `ibm_appid_idp_cloud_directory_invalid_reset_password_enabled` | `reset_password_enabled` |
| `ibm_appid_idp_cloud_directory_invalid_reset_password_notification_enabled` | `reset_password_notification_enabled` |
| `ibm_appid_idp_cloud_directory_invalid_self_service_enabled` | `self_service_enabled` |
| `ibm_appid_idp_cloud_directory_invalid_signup_enabled` | `signup_enabled` |
| `ibm_appid_idp_cloud_directory_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_idp_cloud_directory_invalid_welcome_enabled` | `welcome_enabled` |

## Example

### Valid

```hcl
resource "ibm_appid_idp_cloud_directory" "test" {
  is_active = true
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
