# ibm_iam_action_control_assignment_data

Schema validation for the ibm_iam_action_control_assignment data source.

This rule validates `ibm_iam_action_control_assignment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/v1/action_control_assignments/{assignment_id}` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_action_control_assignment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `assignment_id` | yes | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `operation` | no | — |
| `resources` | no | — |
| `status` | no | — |
| `target` | no | — |
| `template` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_action_control_assignment_data_invalid_assignment_id` | `assignment_id` |

## Example

### Valid

```hcl
data "ibm_iam_action_control_assignment" "test" {
  assignment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
