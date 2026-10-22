# ibm_iam_role_template

Manages ibm_iam_role_template. Required field when creating a new template.

This rule validates `ibm_iam_role_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RolePostRequest` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_template" {
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
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | yes | pattern `^[A-Z]{1}[A-Za-z0-9]{0,29}$` |
| `role_template_id` | no | — |
| `state` | no | — |
| `version` | no | — |
| `role` (block) | no | nested block |
| `role.actions` | yes | API: The action ID. |
| `role.description` | no | API: The description of the role. |
| `role.display_name` | yes | API: The display the name of the role that is shown in the console. |
| `role.name` | yes | pattern `^[A-Z]{1}[A-Za-z0-9]{0,29}$` |
| `role.service_name` | yes | API: The service name. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_template_invalid_committed` | `committed` |
| `ibm_iam_role_template_invalid_description` | `description` |
| `ibm_iam_role_template_invalid_name` | `name` |
| `ibm_iam_role_template_invalid_role__actions` | `role.actions` |
| `ibm_iam_role_template_invalid_role__description` | `role.description` |
| `ibm_iam_role_template_invalid_role__display_name` | `role.display_name` |
| `ibm_iam_role_template_invalid_role__name` | `role.name` |
| `ibm_iam_role_template_invalid_role__service_name` | `role.service_name` |

## Example

### Valid

```hcl
resource "ibm_iam_role_template" "test" {
  name = "PREFIX"
}
```
