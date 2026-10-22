# ibm_container_dedicated_host_pool_data

Reads ibm_container_dedicated_host_pool. The name of the dedicated host pool

This rule validates `ibm_container_dedicated_host_pool` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_dedicated_host_pool_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `flavor_class` | no | — |
| `host_count` | no | — |
| `host_pool_id` | yes | — |
| `metro` | no | — |
| `name` | no | — |
| `state` | no | — |
| `worker_pools` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_pool_data_invalid_host_pool_id` | `host_pool_id` |

## Example

### Valid

```hcl
data "ibm_container_dedicated_host_pool" "test" {
  host_pool_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
