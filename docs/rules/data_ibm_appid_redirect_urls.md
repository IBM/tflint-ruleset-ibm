# ibm_appid_redirect_urls_data

Schema validation for the ibm_appid_redirect_urls data source.

This rule validates `ibm_appid_redirect_urls` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_redirect_urls_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `tenant_id` | yes | — |
| `urls` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_redirect_urls_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_redirect_urls" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
