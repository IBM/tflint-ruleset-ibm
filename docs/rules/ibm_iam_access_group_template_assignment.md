# ibm_iam_access_group_template_assignment

Schema validation for the ibm_iam_access_group_template_assignment resource.

This rule validates `ibm_iam_access_group_template_assignment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TemplateAssignmentRequest` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_template_assignment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `operation` | no | — |
| `status` | no | — |
| `target` | yes | pattern `^[a-zA-Z0-9_-]+$` |
| `target_type` | yes | one of: `Account`, `AccountGroup`; pattern `^[a-zA-Z-]+$` |
| `template_id` | yes | pattern `^[a-zA-Z0-9_-]+$` |
| `template_version` | yes | pattern `^[0-9]+$` |
| `transaction_id` | no | pattern `^[a-zA-Z0-9_-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_template_assignment_invalid_target` | `target` |
| `ibm_iam_access_group_template_assignment_invalid_target_type` | `target_type` |
| `ibm_iam_access_group_template_assignment_invalid_template_id` | `template_id` |
| `ibm_iam_access_group_template_assignment_invalid_template_version` | `template_version` |
| `ibm_iam_access_group_template_assignment_invalid_transaction_id` | `transaction_id` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group_template_assignment" "test" {
  target = "example-name"
  target_type = "Account"
  template_id = "example-name"
  template_version = "1"
}
```
