# ibm_container_vpc_worker

Schema validation for the ibm_container_vpc_worker resource.

This rule validates `ibm_container_vpc_worker` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_vpc_worker" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `check_ptx_status` | no | — |
| `cluster_name` | yes | — |
| `ip` | no | — |
| `kube_config_path` | no | — |
| `ptx_timeout` | no | — |
| `replace_worker` | yes | — |
| `resource_group_id` | no | — |
| `sds` | no | — |
| `sds_timeout` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_worker_invalid_check_ptx_status` | `check_ptx_status` |
| `ibm_container_vpc_worker_invalid_cluster_name` | `cluster_name` |
| `ibm_container_vpc_worker_invalid_kube_config_path` | `kube_config_path` |
| `ibm_container_vpc_worker_invalid_ptx_timeout` | `ptx_timeout` |
| `ibm_container_vpc_worker_invalid_replace_worker` | `replace_worker` |
| `ibm_container_vpc_worker_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_worker_invalid_sds` | `sds` |
| `ibm_container_vpc_worker_invalid_sds_timeout` | `sds_timeout` |

## Example

### Valid

```hcl
resource "ibm_container_vpc_worker" "test" {
  cluster_name = "example"
  replace_worker = "example"
}
```
