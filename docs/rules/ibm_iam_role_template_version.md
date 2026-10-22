# ibm_iam_role_template_version

Manages ibm_iam_role_template_version. Required field when creating a new template.

This rule validates `ibm_iam_role_template_version` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Role` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_template_version" {
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
| `description` | no | API: The description of the role. |
| `href` | no | — |
| `id` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `role_template_id` | yes | pattern `^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `state` | no | — |
| `version` | no | — |
| `role` (block) | no | nested block |
| `role.actions` | yes | API: The action ID. |
| `role.description` | no | API: The description of the role. |
| `role.display_name` | yes | API: The display the name of the role that is shown in the console. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_template_version_invalid_committed` | `committed` |
| `ibm_iam_role_template_version_invalid_description` | `description` |
| `ibm_iam_role_template_version_invalid_name` | `name` |
| `ibm_iam_role_template_version_invalid_role_template_id` | `role_template_id` |
| `ibm_iam_role_template_version_invalid_role__actions` | `role.actions` |
| `ibm_iam_role_template_version_invalid_role__description` | `role.description` |
| `ibm_iam_role_template_version_invalid_role__display_name` | `role.display_name` |

## Example

### Valid

```hcl
resource "ibm_iam_role_template_version" "test" {
  role_template_id = "roleTemplate-01234567-89ab-cdef-0123-456789abcdef"
}
```
