# ibm_iam_idps_data

Schema validation for the ibm_iam_idps data source.

This rule validates `ibm_iam_idps` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/v1/accounts/{account_id}/limits/identity` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_idps_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `idps` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_idps_data_invalid_account_id` | `account_id` |

## Example

### Valid

```hcl
data "ibm_iam_idps" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
