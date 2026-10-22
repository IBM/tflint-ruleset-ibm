# ibm_compute_placement_group_data

Schema validation for the ibm_compute_placement_group data source.

This rule validates `ibm_compute_placement_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_placement_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | no | — |
| `name` | yes | — |
| `pod` | no | — |
| `rule` | no | — |
| `virtual_guests` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_placement_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_compute_placement_group" "test" {
  name = "example"
}
```
