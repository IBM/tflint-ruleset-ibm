# ibm_iam_effective_account_settings_data

Schema validation for the ibm_iam_effective_account_settings data source.

This rule validates `ibm_iam_effective_account_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/accounts/{account_id}/settings/identity` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_effective_account_settings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `account_id` | yes | — |
| `assigned_templates` | no | — |
| `effective` | no | — |
| `include_history` | no | — |
| `resolve_user_mfa` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_effective_account_settings_data_invalid_account_id` | `account_id` |
| `ibm_iam_effective_account_settings_data_invalid_include_history` | `include_history` |
| `ibm_iam_effective_account_settings_data_invalid_resolve_user_mfa` | `resolve_user_mfa` |

## Example

### Valid

```hcl
data "ibm_iam_effective_account_settings" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
