# ibm_pha_supported_locations_data

Schema validation for the ibm_pha_supported_locations data source.

This rule validates `ibm_pha_supported_locations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_supported_locations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `if_none_match` | no | — |
| `instance_id` | yes | — |
| `locations` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_supported_locations_data_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_supported_locations_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pha_supported_locations" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
