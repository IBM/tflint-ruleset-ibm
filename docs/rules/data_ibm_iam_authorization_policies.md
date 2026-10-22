# ibm_iam_authorization_policies_data

Schema validation for the ibm_iam_authorization_policies data source.

This rule validates `ibm_iam_authorization_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PolicyTemplateRequest` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_authorization_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | pattern `^[A-Za-z0-9-]*$` |
| `policies` | no | — |
| `sort` | no | — |
| `transaction_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_authorization_policies_data_invalid_account_id` | `account_id` |
| `ibm_iam_authorization_policies_data_invalid_sort` | `sort` |
| `ibm_iam_authorization_policies_data_invalid_transaction_id` | `transaction_id` |

## Example

### Valid

```hcl
data "ibm_iam_authorization_policies" "test" {
}
```
