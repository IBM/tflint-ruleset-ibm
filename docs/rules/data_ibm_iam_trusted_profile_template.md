# ibm_iam_trusted_profile_template_data

Reads ibm_iam_trusted_profile_template. The name of the trusted profile template.

This rule validates `ibm_iam_trusted_profile_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/profile_templates/{template_id}/versions/{version}` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `entity_tag` | no | — |
| `history` | no | — |
| `id` | no | — |
| `include_history` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `policy_template_references` | no | — |
| `profile` | no | — |
| `template_id` | yes | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_template_data_invalid_include_history` | `include_history` |
| `ibm_iam_trusted_profile_template_data_invalid_template_id` | `template_id` |
| `ibm_iam_trusted_profile_template_data_invalid_version` | `version` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_template" "test" {
  template_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
