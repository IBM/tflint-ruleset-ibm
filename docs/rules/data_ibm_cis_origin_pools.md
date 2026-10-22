# ibm_cis_origin_pools_data

Schema validation for the ibm_cis_origin_pools data source.

This rule validates `ibm_cis_origin_pools` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_origin_pools_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `cis_origin_pools` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_pools_data_invalid_cis_id` | `cis_id` |

## Example

### Valid

```hcl
data "ibm_cis_origin_pools" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
