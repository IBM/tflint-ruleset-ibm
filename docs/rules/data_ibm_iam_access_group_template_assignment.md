# ibm_iam_access_group_template_assignment_data

Schema validation for the ibm_iam_access_group_template_assignment data source.

This rule validates `ibm_iam_access_group_template_assignment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TemplateAssignmentRequest` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_template_assignment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `assignments` | no | — |
| `first` | no | — |
| `last` | no | — |
| `limit` | no | — |
| `offset` | no | — |
| `status` | no | one of: `accepted`, `in_progress`, `succeeded`, `failed`, `superseded` |
| `target` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `template_id` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `template_version` | no | pattern `^[0-9]+$` |
| `total_count` | no | — |
| `transaction_id` | no | pattern `^[a-zA-Z0-9_-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_template_assignment_data_invalid_status` | `status` |
| `ibm_iam_access_group_template_assignment_data_invalid_target` | `target` |
| `ibm_iam_access_group_template_assignment_data_invalid_template_id` | `template_id` |
| `ibm_iam_access_group_template_assignment_data_invalid_template_version` | `template_version` |
| `ibm_iam_access_group_template_assignment_data_invalid_transaction_id` | `transaction_id` |

## Example

### Valid

```hcl
data "ibm_iam_access_group_template_assignment" "test" {
}
```
