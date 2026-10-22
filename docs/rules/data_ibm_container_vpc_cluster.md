# ibm_container_vpc_cluster_data

Reads ibm_container_vpc_cluster. Name or id of the cluster

This rule validates `ibm_container_vpc_cluster` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vpc_cluster_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alb_type` | no | one of: `private`, `public`, `all` |
| `albs` | no | — |
| `api_key_id` | no | — |
| `api_key_owner_email` | no | — |
| `api_key_owner_name` | no | — |
| `cluster_name` | no | — |
| `cluster_name_id` | no | — |
| `crn` | no | — |
| `health` | no | — |
| `image_security_enforcement` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `kube_version` | no | — |
| `master_url` | no | — |
| `name` | no | API: The name value. |
| `network_plugin` | no | — |
| `offering` | no | — |
| `pod_subnet` | no | — |
| `private_service_endpoint` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint` | no | — |
| `public_service_endpoint_url` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service_subnet` | no | — |
| `state` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `vpe_service_endpoint_url` | no | — |
| `wait_till` | no | one of: `MasterNodeReady`, `OneWorkerNodeReady`, `IngressReady`, `normal` |
| `wait_till_timeout` | no | — |
| `worker_count` | no | — |
| `worker_pools` | no | — |
| `workers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_cluster_data_invalid_alb_type` | `alb_type` |
| `ibm_container_vpc_cluster_data_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_vpc_cluster_data_invalid_name` | `name` |
| `ibm_container_vpc_cluster_data_invalid_network_plugin` | `network_plugin` |
| `ibm_container_vpc_cluster_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_cluster_data_invalid_wait_till` | `wait_till` |
| `ibm_container_vpc_cluster_data_invalid_wait_till_timeout` | `wait_till_timeout` |

## Example

### Valid

```hcl
data "ibm_container_vpc_cluster" "test" {
}
```
