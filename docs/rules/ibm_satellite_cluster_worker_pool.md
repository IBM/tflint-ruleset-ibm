# ibm_satellite_cluster_worker_pool

Manages ibm_satellite_cluster_worker_pool. The name for the worker pool

This rule validates `ibm_satellite_cluster_worker_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_cluster_worker_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | — |
| `disk_encryption` | no | — |
| `entitlement` | no | — |
| `flavor` | no | — |
| `host_labels` | no | — |
| `isolation` | no | — |
| `name` | yes | — |
| `operating_system` | no | — |
| `resource_group_id` | no | — |
| `worker_count` | no | — |
| `worker_pool_labels` | no | — |
| `zones` (block) | no | nested block |
| `zones.id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_cluster_worker_pool_invalid_cluster` | `cluster` |
| `ibm_satellite_cluster_worker_pool_invalid_disk_encryption` | `disk_encryption` |
| `ibm_satellite_cluster_worker_pool_invalid_entitlement` | `entitlement` |
| `ibm_satellite_cluster_worker_pool_invalid_flavor` | `flavor` |
| `ibm_satellite_cluster_worker_pool_invalid_host_labels` | `host_labels` |
| `ibm_satellite_cluster_worker_pool_invalid_isolation` | `isolation` |
| `ibm_satellite_cluster_worker_pool_invalid_name` | `name` |
| `ibm_satellite_cluster_worker_pool_invalid_operating_system` | `operating_system` |
| `ibm_satellite_cluster_worker_pool_invalid_resource_group_id` | `resource_group_id` |
| `ibm_satellite_cluster_worker_pool_invalid_worker_count` | `worker_count` |
| `ibm_satellite_cluster_worker_pool_invalid_worker_pool_labels` | `worker_pool_labels` |
| `ibm_satellite_cluster_worker_pool_invalid_zones__id` | `zones.id` |

## Example

### Valid

```hcl
resource "ibm_satellite_cluster_worker_pool" "test" {
  cluster = "example"
  name = "example"
}
```
