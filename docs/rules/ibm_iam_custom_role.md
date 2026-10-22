# ibm_iam_custom_role

Manages ibm_iam_custom_role. The name of the custom Role

This rule validates `ibm_iam_custom_role` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Role` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_custom_role" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `actions` | yes | API: The action ID. |
| `crn` | no | — |
| `description` | no | API: The description of the role. |
| `display_name` | yes | API: The display the name of the role that is shown in the console. |
| `name` | yes | pattern `^[A-Z]{1}[A-Za-z0-9]{0,29}$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_name` | no | — |
| `service` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_custom_role_invalid_actions` | `actions` |
| `ibm_iam_custom_role_invalid_description` | `description` |
| `ibm_iam_custom_role_invalid_display_name` | `display_name` |
| `ibm_iam_custom_role_invalid_name` | `name` |
| `ibm_iam_custom_role_invalid_service` | `service` |

## Example

### Valid

```hcl
resource "ibm_iam_custom_role" "test" {
  actions = "aaaaaaa"
  display_name = "aaaaaaa"
  name = "PREFIX"
  service = "example"
}
```
