# ibm_container_worker_pool_data

Schema validation for the ibm_container_worker_pool data source.

This rule validates `ibm_container_worker_pool` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DisablePodSecurityPolicyRequest` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_worker_pool_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `autoscale_enabled` | no | — |
| `cluster` | yes | API: The cluster value. |
| `disk_encryption` | no | — |
| `hardware` | no | — |
| `labels` | no | — |
| `machine_type` | no | — |
| `operating_system` | no | — |
| `resource_group_id` | no | — |
| `size_per_zone` | no | — |
| `state` | no | — |
| `worker_pool_name` | yes | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_worker_pool_data_invalid_cluster` | `cluster` |
| `ibm_container_worker_pool_data_invalid_worker_pool_name` | `worker_pool_name` |

## Example

### Valid

```hcl
data "ibm_container_worker_pool" "test" {
  cluster = "aaaaaaa"
  worker_pool_name = "example"
}
```
