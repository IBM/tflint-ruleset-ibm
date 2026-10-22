# ibm_iam_trusted_profile_template_assignment

Schema validation for the ibm_iam_trusted_profile_template_assignment resource.

This rule validates `ibm_iam_trusted_profile_template_assignment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateTemplateAssignmentRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_template_assignment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `entity_tag` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `resources` | no | — |
| `status` | no | — |
| `target` | yes | API: Identifier of target to deploy to |
| `target_type` | yes | one of: `Account`, `AccountGroup` |
| `template_id` | yes | API: ID of the template to assign |
| `template_version` | yes | min 1 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_template_assignment_invalid_target` | `target` |
| `ibm_iam_trusted_profile_template_assignment_invalid_target_type` | `target_type` |
| `ibm_iam_trusted_profile_template_assignment_invalid_template_id` | `template_id` |
| `ibm_iam_trusted_profile_template_assignment_invalid_template_version` | `template_version` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_template_assignment" "test" {
  target = "example"
  target_type = "Account"
  template_id = "01234567-89ab-cdef-0123-456789abcdef"
  template_version = 1
}
```
