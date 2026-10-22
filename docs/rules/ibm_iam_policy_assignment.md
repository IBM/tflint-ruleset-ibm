# ibm_iam_policy_assignment

Schema validation for the ibm_iam_policy_assignment resource.

This rule validates `ibm_iam_policy_assignment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ActionControlAssignmentTemplate` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_policy_assignment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `resources` | no | — |
| `status` | no | — |
| `target` | yes | — |
| `template` | no | — |
| `template_version` | no | pattern `^[0-9]*$` |
| `version` | yes | pattern `^[0-9]*$` |
| `templates` (block) | yes | nested block |
| `templates.id` | yes | pattern `^actionControlTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `templates.version` | yes | pattern `^[0-9]*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_policy_assignment_invalid_accept_language` | `accept_language` |
| `ibm_iam_policy_assignment_invalid_target` | `target` |
| `ibm_iam_policy_assignment_invalid_template_version` | `template_version` |
| `ibm_iam_policy_assignment_invalid_version` | `version` |
| `ibm_iam_policy_assignment_invalid_templates__id` | `templates.id` |
| `ibm_iam_policy_assignment_invalid_templates__version` | `templates.version` |

## Example

### Valid

```hcl
resource "ibm_iam_policy_assignment" "test" {
  target = "example"
  version = "1"
  templates {
    id = "actionControlTemplate-01234567-89ab-cdef-0123-456789abcdef"
    version = "1"
  }
}
```
