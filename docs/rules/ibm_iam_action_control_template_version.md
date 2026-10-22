# ibm_iam_action_control_template_version

Manages ibm_iam_action_control_template_version. Required field when creating a new template.

This rule validates `ibm_iam_action_control_template_version` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/action_control_templates/{action_control_template_id}/versions/{version}:put` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_action_control_template_version" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `action_control_template_id` | yes | pattern `^actionControlTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `committed` | no | API: Committed status of the template version. If committed is set to true, then the template version can no longer be updated. |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `description` | no | API: Description of the action control template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the action control for enterprise users managing IAM templates. |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | API: Required field when creating a new template. Otherwise, this field is optional. If the field is included, it will change the name value for all existing versions of the template. |
| `version` | no | — |
| `action_control` (block) | no | nested block |
| `action_control.actions` | yes | API: IAM action ID. |
| `action_control.description` | no | API: Description of the action control. |
| `action_control.service_name` | yes | API: The service name that the action control refers. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_action_control_template_version_invalid_action_control_template_id` | `action_control_template_id` |
| `ibm_iam_action_control_template_version_invalid_committed` | `committed` |
| `ibm_iam_action_control_template_version_invalid_description` | `description` |
| `ibm_iam_action_control_template_version_invalid_name` | `name` |
| `ibm_iam_action_control_template_version_invalid_action_control__actions` | `action_control.actions` |
| `ibm_iam_action_control_template_version_invalid_action_control__description` | `action_control.description` |
| `ibm_iam_action_control_template_version_invalid_action_control__service_name` | `action_control.service_name` |

## Example

### Valid

```hcl
resource "ibm_iam_action_control_template_version" "test" {
  action_control_template_id = "actionControlTemplate-01234567-89ab-cdef-0123-456789abcdef"
}
```
