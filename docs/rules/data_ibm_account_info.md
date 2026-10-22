# ibm_account_info_data

Schema validation for the ibm_account_info data source.

This rule validates `ibm_account_info` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/v2/accounts/{account_id}/users` (`user-management`).

## Configuration

```hcl
rule "ibm_account_info_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `linked_softlayer_account` | no | — |
| `name` | no | — |
| `owner` | no | — |
| `owner_iamid` | no | — |
| `owner_userid` | no | — |
| `status` | no | — |
| `team_directory_enabled` | no | — |
| `traits` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_account_info_data_invalid_account_id` | `account_id` |

## Example

### Valid

```hcl
data "ibm_account_info" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
