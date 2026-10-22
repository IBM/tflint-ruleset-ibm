# ibm_compute_reserved_capacity_data

Reads ibm_compute_reserved_capacity. Name of reserved instance

This rule validates `ibm_compute_reserved_capacity` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_reserved_capacity_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | no | — |
| `flavor` | no | — |
| `instances` | no | — |
| `most_recent` | no | — |
| `name` | yes | — |
| `pod` | no | — |
| `virtual_guests` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_reserved_capacity_data_invalid_most_recent` | `most_recent` |
| `ibm_compute_reserved_capacity_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_compute_reserved_capacity" "test" {
  name = "example"
}
```
