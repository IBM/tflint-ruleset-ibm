# ibm_is_vpc_dns_resolution_bindings_data

Schema validation for the ibm_is_vpc_dns_resolution_bindings data source.

This rule validates `ibm_is_vpc_dns_resolution_bindings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpc_dns_resolution_bindings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_resolution_bindings` | no | — |
| `vpc_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_dns_resolution_bindings_data_invalid_vpc_id` | `vpc_id` |

## Example

### Valid

```hcl
data "ibm_is_vpc_dns_resolution_bindings" "test" {
  vpc_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
