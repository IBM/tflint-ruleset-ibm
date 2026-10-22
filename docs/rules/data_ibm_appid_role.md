# ibm_appid_role_data

Reads ibm_appid_role. Unique role name

This rule validates `ibm_appid_role` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_role_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access` | no | — |
| `description` | no | — |
| `name` | no | — |
| `role_id` | yes | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_role_data_invalid_role_id` | `role_id` |
| `ibm_appid_role_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_role" "test" {
  role_id = "01234567-89ab-cdef-0123-456789abcdef"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
