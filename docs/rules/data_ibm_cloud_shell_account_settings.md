# ibm_cloud_shell_account_settings_data

Schema validation for the ibm_cloud_shell_account_settings data source.

This rule validates `ibm_cloud_shell_account_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/api/v1/user/accounts/{account_id}/settings:post` (`cloudshell`).

## Configuration

```hcl
rule "ibm_cloud_shell_account_settings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | pattern `^[a-zA-Z0-9-]*$` |
| `created_at` | no | — |
| `created_by` | no | — |
| `default_enable_new_features` | no | — |
| `default_enable_new_regions` | no | — |
| `enabled` | no | — |
| `features` | no | — |
| `regions` | no | — |
| `rev` | no | — |
| `type` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloud_shell_account_settings_data_invalid_account_id` | `account_id` |

## Example

### Valid

```hcl
data "ibm_cloud_shell_account_settings" "test" {
  account_id = "example-name"
}
```
