# ibm_iam_account_settings_template_assignment_data

Schema validation for the ibm_iam_account_settings_template_assignment data source.

This rule validates `ibm_iam_account_settings_template_assignment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/account_settings_assignments/:post` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_account_settings_template_assignment_data" {
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
| `target` | no | API: Identifier of target to deploy to |
| `target_type` | no | one of: `Account`, `AccountGroup` |
| `template_id` | no | API: ID of the template to assign |
| `template_version` | no | min 1 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_template_assignment_data_invalid_assignment_id` | `assignment_id` |
| `ibm_iam_account_settings_template_assignment_data_invalid_include_history` | `include_history` |
| `ibm_iam_account_settings_template_assignment_data_invalid_target` | `target` |
| `ibm_iam_account_settings_template_assignment_data_invalid_target_type` | `target_type` |
| `ibm_iam_account_settings_template_assignment_data_invalid_template_id` | `template_id` |
| `ibm_iam_account_settings_template_assignment_data_invalid_template_version` | `template_version` |

## Example

### Valid

```hcl
data "ibm_iam_account_settings_template_assignment" "test" {
  assignment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
