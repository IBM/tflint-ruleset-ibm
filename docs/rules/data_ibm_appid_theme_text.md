# ibm_appid_theme_text_data

Schema validation for the ibm_appid_theme_text data source.

This rule validates `ibm_appid_theme_text` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_theme_text_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `footnote` | no | — |
| `tab_title` | no | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_theme_text_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_theme_text" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
