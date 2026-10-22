# ibm_container_vpc_cluster_alb_data

Schema validation for the ibm_container_vpc_cluster_alb data source.

This rule validates `ibm_container_vpc_cluster_alb` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/ingress/v2/clusters/{id_or_name}/albs/{alb_id}/autoscale` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vpc_cluster_alb_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alb_id` | yes | — |
| `alb_type` | no | — |
| `cluster` | no | — |
| `disable_deployment` | no | — |
| `enable` | no | — |
| `load_balancer_hostname` | no | — |
| `name` | no | — |
| `resize` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `status` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_cluster_alb_data_invalid_alb_id` | `alb_id` |
| `ibm_container_vpc_cluster_alb_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_container_vpc_cluster_alb" "test" {
  alb_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
