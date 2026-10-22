# ibm_appid_action_url_data

Schema validation for the ibm_appid_action_url data source.

This rule validates `ibm_appid_action_url` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_action_url_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `on_user_verified`, `on_reset_password` |
| `tenant_id` | yes | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_action_url_data_invalid_action` | `action` |
| `ibm_appid_action_url_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_action_url" "test" {
  action = "on_user_verified"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
