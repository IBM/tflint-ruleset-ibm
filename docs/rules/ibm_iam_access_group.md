# ibm_iam_access_group

Manages ibm_iam_access_group. Name of the access group

This rule validates `ibm_iam_access_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/groups/{access_group_id}:patch` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `description` | no | API: Assign an optional description for the access group. This field has a limit of 250 characters. |
| `name` | yes | API: Give the access group a unique name that doesn't conflict with an existing access group in the account. This field is case-insensitive and has a limit of 100 characters. |
| `tags` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_invalid_description` | `description` |
| `ibm_iam_access_group_invalid_name` | `name` |
| `ibm_iam_access_group_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group" "test" {
  name = "example"
}
```
