# ibm_container_dedicated_host_data

Schema validation for the ibm_container_dedicated_host data source.

This rule validates `ibm_container_dedicated_host` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_dedicated_host_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `flavor` | no | — |
| `host_id` | yes | — |
| `host_pool_id` | yes | — |
| `life_cycle` | no | — |
| `placement_enabled` | no | — |
| `resources` | no | — |
| `workers` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_data_invalid_host_id` | `host_id` |
| `ibm_container_dedicated_host_data_invalid_host_pool_id` | `host_pool_id` |

## Example

### Valid

```hcl
data "ibm_container_dedicated_host" "test" {
  host_id = "01234567-89ab-cdef-0123-456789abcdef"
  host_pool_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
