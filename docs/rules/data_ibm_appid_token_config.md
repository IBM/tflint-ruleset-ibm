# ibm_appid_token_config_data

Schema validation for the ibm_appid_token_config data source.

This rule validates `ibm_appid_token_config` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_token_config_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_token_claim` | no | — |
| `access_token_expires_in` | no | — |
| `anonymous_access_enabled` | no | — |
| `anonymous_token_expires_in` | no | — |
| `id_token_claim` | no | — |
| `refresh_token_enabled` | no | — |
| `refresh_token_expires_in` | no | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_token_config_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_token_config" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
