# ibm_appid_theme_color

Schema validation for the ibm_appid_theme_color resource.

This rule validates `ibm_appid_theme_color` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_theme_color" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `header_color` | yes | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_theme_color_invalid_header_color` | `header_color` |
| `ibm_appid_theme_color_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_theme_color" "test" {
  header_color = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
