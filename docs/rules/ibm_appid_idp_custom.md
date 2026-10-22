# ibm_appid_idp_custom

Schema validation for the ibm_appid_idp_custom resource.

This rule validates `ibm_appid_idp_custom` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_idp_custom" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `is_active` | yes | — |
| `public_key` | no | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_idp_custom_invalid_is_active` | `is_active` |
| `ibm_appid_idp_custom_invalid_public_key` | `public_key` |
| `ibm_appid_idp_custom_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_idp_custom" "test" {
  is_active = true
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
