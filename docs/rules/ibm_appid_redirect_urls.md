# ibm_appid_redirect_urls

Schema validation for the ibm_appid_redirect_urls resource.

This rule validates `ibm_appid_redirect_urls` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_redirect_urls" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `tenant_id` | yes | — |
| `urls` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_redirect_urls_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_redirect_urls_invalid_urls` | `urls` |

## Example

### Valid

```hcl
resource "ibm_appid_redirect_urls" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  urls = "example"
}
```
