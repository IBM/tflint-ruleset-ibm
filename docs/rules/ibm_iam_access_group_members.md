# ibm_iam_access_group_members

Schema validation for the ibm_iam_access_group_members resource.

This rule validates `ibm_iam_access_group_members` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_access_group_members" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_group_id` | yes | — |
| `iam_profile_ids` | no | — |
| `iam_service_ids` | no | — |
| `ibm_ids` | no | — |
| `members` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_members_invalid_access_group_id` | `access_group_id` |
| `ibm_iam_access_group_members_invalid_iam_profile_ids` | `iam_profile_ids` |
| `ibm_iam_access_group_members_invalid_iam_service_ids` | `iam_service_ids` |
| `ibm_iam_access_group_members_invalid_ibm_ids` | `ibm_ids` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group_members" "test" {
  access_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
