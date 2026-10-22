# ibm_appid_idp_google

Schema validation for the ibm_appid_idp_google resource.

This rule validates `ibm_appid_idp_google` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_idp_google" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `is_active` | yes | — |
| `redirect_url` | no | — |
| `tenant_id` | yes | — |
| `config` (block) | no | nested block |
| `config.application_id` | yes | — |
| `config.application_secret` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_idp_google_invalid_is_active` | `is_active` |
| `ibm_appid_idp_google_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_idp_google_invalid_config__application_id` | `config.application_id` |
| `ibm_appid_idp_google_invalid_config__application_secret` | `config.application_secret` |

## Example

### Valid

```hcl
resource "ibm_appid_idp_google" "test" {
  is_active = true
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
