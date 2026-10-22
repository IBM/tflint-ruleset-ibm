# ibm_container_vpc_cluster_worker_data

Schema validation for the ibm_container_vpc_cluster_worker data source.

This rule validates `ibm_container_vpc_cluster_worker` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_vpc_cluster_worker_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_name_id` | yes | — |
| `flavor` | no | — |
| `host_pool_id` | no | — |
| `kube_version` | no | — |
| `network_interfaces` | no | — |
| `pool_id` | no | — |
| `pool_name` | no | — |
| `resource_controller_url` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `worker_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_cluster_worker_data_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_vpc_cluster_worker_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_cluster_worker_data_invalid_worker_id` | `worker_id` |

## Example

### Valid

```hcl
data "ibm_container_vpc_cluster_worker" "test" {
  cluster_name_id = "01234567-89ab-cdef-0123-456789abcdef"
  worker_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
