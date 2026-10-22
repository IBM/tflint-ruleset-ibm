# ibm_iam_user_mfa_enrollments_data

Schema validation for the ibm_iam_user_mfa_enrollments data source.

This rule validates `ibm_iam_user_mfa_enrollments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/preferences/accounts/{account_id}/identities/{iam_id}/{service}/{preference_id}` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_user_mfa_enrollments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_based_mfa` | no | — |
| `account_id` | yes | — |
| `effective_mfa_type` | no | — |
| `iam_id` | yes | — |
| `id_based_mfa` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_user_mfa_enrollments_data_invalid_account_id` | `account_id` |
| `ibm_iam_user_mfa_enrollments_data_invalid_iam_id` | `iam_id` |

## Example

### Valid

```hcl
data "ibm_iam_user_mfa_enrollments" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  iam_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
