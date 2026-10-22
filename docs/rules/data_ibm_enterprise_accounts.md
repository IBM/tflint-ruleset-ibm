# ibm_enterprise_accounts_data

Reads ibm_enterprise_accounts. The name of the account.

This rule validates `ibm_enterprise_accounts` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/account-groups/{account_group_id}:patch` (`enterprise-apis-enterprise`).

## Configuration

```hcl
rule "ibm_enterprise_accounts_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accounts` | no | — |
| `name` | no | API: The new name of the account group. This field must have 3 - 60 characters. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_enterprise_accounts_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_enterprise_accounts" "test" {
}
```
