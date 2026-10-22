# ibm_iam_user_policy_data

Schema validation for the ibm_iam_user_policy data source.

This rule validates `ibm_iam_user_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `V2PolicySubjectAttribute` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_user_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ibm_id` | yes | — |
| `policies` | no | — |
| `sort` | no | — |
| `transaction_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_user_policy_data_invalid_ibm_id` | `ibm_id` |
| `ibm_iam_user_policy_data_invalid_sort` | `sort` |
| `ibm_iam_user_policy_data_invalid_transaction_id` | `transaction_id` |

## Example

### Valid

```hcl
data "ibm_iam_user_policy" "test" {
  ibm_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
