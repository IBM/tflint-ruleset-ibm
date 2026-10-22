# ibm_iam_access_group_data

Schema validation for the ibm_iam_access_group data source.

This rule validates `ibm_iam_access_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/groups/{access_group_id}:patch` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_group_name` | no | — |
| `groups` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_data_invalid_access_group_name` | `access_group_name` |

## Example

### Valid

```hcl
data "ibm_iam_access_group" "test" {
}
```
