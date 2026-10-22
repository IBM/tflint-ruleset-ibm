# ibm_appid_idp_cloud_directory_data

Schema validation for the ibm_appid_idp_cloud_directory data source.

This rule validates `ibm_appid_idp_cloud_directory` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_idp_cloud_directory_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `identity_confirm_access_mode` | no | — |
| `identity_confirm_methods` | no | — |
| `identity_field` | no | — |
| `is_active` | no | — |
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
| `ibm_appid_idp_cloud_directory_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_idp_cloud_directory" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
