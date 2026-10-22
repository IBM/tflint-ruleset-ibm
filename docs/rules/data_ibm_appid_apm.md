# ibm_appid_apm_data

Schema validation for the ibm_appid_apm data source.

This rule validates `ibm_appid_apm` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_apm_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | no | — |
| `lockout_policy` | no | — |
| `min_password_change_interval` | no | — |
| `password_expiration` | no | — |
| `password_reuse` | no | — |
| `prevent_password_with_username` | no | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_apm_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_apm" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
