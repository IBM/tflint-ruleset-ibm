# ibm_appid_cloud_directory_user_data

Schema validation for the ibm_appid_cloud_directory_user data source.

This rule validates `ibm_appid_cloud_directory_user` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_cloud_directory_user_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | no | — |
| `display_name` | no | — |
| `email` | no | — |
| `locked_until` | no | — |
| `meta` | no | — |
| `status` | no | — |
| `subject` | no | — |
| `tenant_id` | yes | — |
| `user_id` | yes | — |
| `user_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_cloud_directory_user_data_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_cloud_directory_user_data_invalid_user_id` | `user_id` |

## Example

### Valid

```hcl
data "ibm_appid_cloud_directory_user" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  user_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
