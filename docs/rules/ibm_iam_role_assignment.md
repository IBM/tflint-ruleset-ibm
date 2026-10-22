# ibm_iam_role_assignment

Schema validation for the ibm_iam_role_assignment resource.

This rule validates `ibm_iam_role_assignment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/role_assignments/{assignment_id}:patch` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_assignment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `operation` | no | — |
| `resources` | no | — |
| `status` | no | — |
| `template_version` | no | pattern `^[0-9]*$` |
| `target` (block) | yes | nested block |
| `target.id` | yes | — |
| `target.type` | yes | — |
| `templates` (block) | yes | nested block |
| `templates.id` | yes | — |
| `templates.version` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_assignment_invalid_template_version` | `template_version` |
| `ibm_iam_role_assignment_invalid_target__id` | `target.id` |
| `ibm_iam_role_assignment_invalid_target__type` | `target.type` |
| `ibm_iam_role_assignment_invalid_templates__id` | `templates.id` |
| `ibm_iam_role_assignment_invalid_templates__version` | `templates.version` |

## Example

### Valid

```hcl
resource "ibm_iam_role_assignment" "test" {
  target {
    id = "example"
    type = "example"
  }
  templates {
    id = "example"
    version = "example"
  }
}
```
