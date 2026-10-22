# ibm_iam_role_template_version_data

Reads ibm_iam_role_template_version. Required field when creating a new template.

This rule validates `ibm_iam_role_template_version` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/role_templates/{role_template_id}/versions/{version}` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_template_version_data" {
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
| `description` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `role` | no | — |
| `role_template_id` | yes | pattern `^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `state` | no | — |
| `version` | yes | pattern `^[0-9]*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_template_version_data_invalid_role_template_id` | `role_template_id` |
| `ibm_iam_role_template_version_data_invalid_version` | `version` |

## Example

### Valid

```hcl
data "ibm_iam_role_template_version" "test" {
  role_template_id = "roleTemplate-01234567-89ab-cdef-0123-456789abcdef"
  version = "1"
}
```
