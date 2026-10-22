# ibm_pdr_dr_locations_data

Schema validation for the ibm_pdr_dr_locations data source.

This rule validates `ibm_pdr_dr_locations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pdr_dr_locations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `dr_locations` | no | — |
| `instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_dr_locations_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_dr_locations_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pdr_dr_locations" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
