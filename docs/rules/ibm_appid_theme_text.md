# ibm_appid_theme_text

Schema validation for the ibm_appid_theme_text resource.

This rule validates `ibm_appid_theme_text` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/config/ui/theme_text:put` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_theme_text" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `footnote` | no | API: The footnote. |
| `tab_title` | no | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_theme_text_invalid_footnote` | `footnote` |
| `ibm_appid_theme_text_invalid_tab_title` | `tab_title` |
| `ibm_appid_theme_text_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_theme_text" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
