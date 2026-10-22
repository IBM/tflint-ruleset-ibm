# ibm_cbr_zone_addresses_data

Schema validation for the ibm_cbr_zone_addresses data source.

This rule validates `ibm_cbr_zone_addresses` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cbr_zone_addresses_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `addresses` | no | — |
| `zone_addresses_id` | yes | — |
| `zone_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_zone_addresses_data_invalid_zone_addresses_id` | `zone_addresses_id` |

## Example

### Valid

```hcl
data "ibm_cbr_zone_addresses" "test" {
  zone_addresses_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
