# ibm_kp_key_data

Schema validation for the ibm_kp_key data source.

This rule validates `ibm_kp_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_kp_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `key_name` | no | — |
| `key_protect_id` | yes | — |
| `keys` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kp_key_data_invalid_key_name` | `key_name` |
| `ibm_kp_key_data_invalid_key_protect_id` | `key_protect_id` |

## Example

### Valid

```hcl
data "ibm_kp_key" "test" {
  key_protect_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
