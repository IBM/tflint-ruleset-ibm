# ibm_iam_service_policy_data

Schema validation for the ibm_iam_service_policy data source.

This rule validates `ibm_iam_service_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_service_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `iam_id` | no | — |
| `iam_service_id` | no | — |
| `policies` | no | — |
| `sort` | no | — |
| `transaction_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_service_policy_data_invalid_iam_id` | `iam_id` |
| `ibm_iam_service_policy_data_invalid_iam_service_id` | `iam_service_id` |
| `ibm_iam_service_policy_data_invalid_sort` | `sort` |
| `ibm_iam_service_policy_data_invalid_transaction_id` | `transaction_id` |

## Example

### Valid

```hcl
data "ibm_iam_service_policy" "test" {
}
```
