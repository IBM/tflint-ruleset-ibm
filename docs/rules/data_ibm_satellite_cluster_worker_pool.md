# ibm_satellite_cluster_worker_pool_data

Reads ibm_satellite_cluster_worker_pool. worker pool name

This rule validates `ibm_satellite_cluster_worker_pool` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_cluster_worker_pool_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_scale_enabled` | no | — |
| `cluster` | yes | — |
| `flavor` | no | — |
| `host_labels` | no | — |
| `isolation` | no | — |
| `name` | yes | — |
| `openshift_license_source` | no | — |
| `operating_system` | no | — |
| `region` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `worker_count` | no | — |
| `worker_pool_labels` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_cluster_worker_pool_data_invalid_cluster` | `cluster` |
| `ibm_satellite_cluster_worker_pool_data_invalid_name` | `name` |
| `ibm_satellite_cluster_worker_pool_data_invalid_region` | `region` |
| `ibm_satellite_cluster_worker_pool_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_satellite_cluster_worker_pool" "test" {
  cluster = "example"
  name = "example"
}
```
