# ibm_iam_trusted_profile_template_assignment_data

Schema validation for the ibm_iam_trusted_profile_template_assignment data source.

This rule validates `ibm_iam_trusted_profile_template_assignment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/v1/profile_assignments/{assignment_id}` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_template_assignment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `assignment_id` | yes | — |
| `context` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `entity_tag` | no | — |
| `history` | no | — |
| `href` | no | — |
| `id` | no | — |
| `include_history` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `resources` | no | — |
| `status` | no | — |
| `target` | no | — |
| `target_type` | no | — |
| `template_id` | no | — |
| `template_version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_template_assignment_data_invalid_assignment_id` | `assignment_id` |
| `ibm_iam_trusted_profile_template_assignment_data_invalid_include_history` | `include_history` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_template_assignment" "test" {
  assignment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
