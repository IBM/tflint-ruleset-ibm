# ibm_iam_account_settings_external_interaction_data

Schema validation for the ibm_iam_account_settings_external_interaction data source.

This rule validates `ibm_iam_account_settings_external_interaction` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/v1/accounts/{account_id}/settings/access_management` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_account_settings_external_interaction_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `account_id` | yes | pattern `^[A-Za-z0-9-_]+$` |
| `external_account_identity_interaction` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_external_interaction_data_invalid_accept_language` | `accept_language` |
| `ibm_iam_account_settings_external_interaction_data_invalid_account_id` | `account_id` |

## Example

### Valid

```hcl
data "ibm_iam_account_settings_external_interaction" "test" {
  account_id = "example-name"
}
```
