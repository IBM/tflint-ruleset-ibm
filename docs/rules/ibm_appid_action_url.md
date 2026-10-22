# ibm_appid_action_url

Schema validation for the ibm_appid_action_url resource.

This rule validates `ibm_appid_action_url` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_action_url" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `on_user_verified`, `on_reset_password` |
| `tenant_id` | yes | — |
| `url` | yes | API: The URL. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_action_url_invalid_action` | `action` |
| `ibm_appid_action_url_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_action_url_invalid_url` | `url` |

## Example

### Valid

```hcl
resource "ibm_appid_action_url" "test" {
  action = "on_user_verified"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  url = "aaaaaaa"
}
```
