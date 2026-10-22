# ibm_iam_identity_preferences_data

Schema validation for the ibm_iam_identity_preferences data source.

This rule validates `ibm_iam_identity_preferences` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/preferences/accounts/{account_id}/identities/{iam_id}/{service}/{preference_id}` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_identity_preferences_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `iam_id` | yes | — |
| `preferences` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_identity_preferences_data_invalid_account_id` | `account_id` |
| `ibm_iam_identity_preferences_data_invalid_iam_id` | `iam_id` |

## Example

### Valid

```hcl
data "ibm_iam_identity_preferences" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  iam_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
