# ibm_is_volume_jobs_data

Schema validation for the ibm_is_volume_jobs data source.

This rule validates `ibm_is_volume_jobs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_volume_jobs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `jobs` | no | — |
| `volume_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volume_jobs_data_invalid_volume_id` | `volume_id` |

## Example

### Valid

```hcl
data "ibm_is_volume_jobs" "test" {
  volume_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
