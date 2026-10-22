# ibm_container_cluster_data

Reads ibm_container_cluster. Name or id of the cluster

This rule validates `ibm_container_cluster` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_guid` | no | — |
| `alb_type` | no | one of: `private`, `public`, `all` |
| `albs` | no | — |
| `api_key_id` | no | — |
| `api_key_owner_email` | no | — |
| `api_key_owner_name` | no | — |
| `bounded_services` | no | — |
| `cluster_name` | no | — |
| `cluster_name_id` | no | — |
| `crn` | no | — |
| `image_security_enforcement` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `is_trusted` | no | — |
| `list_bounded_services` | no | — |
| `name` | no | API: The name value. |
| `org_guid` | no | — |
| `private_service_endpoint` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint` | no | — |
| `public_service_endpoint_url` | no | — |
| `region` | no | API: The region value. |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `server_url` | no | — |
| `space_guid` | no | — |
| `state` | no | — |
| `vlans` | no | — |
| `wait_till` | no | one of: `MasterNodeReady`, `OneWorkerNodeReady`, `normal` |
| `wait_till_timeout` | no | — |
| `worker_count` | no | — |
| `worker_pools` | no | — |
| `workers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_data_invalid_account_guid` | `account_guid` |
| `ibm_container_cluster_data_invalid_alb_type` | `alb_type` |
| `ibm_container_cluster_data_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_cluster_data_invalid_list_bounded_services` | `list_bounded_services` |
| `ibm_container_cluster_data_invalid_name` | `name` |
| `ibm_container_cluster_data_invalid_org_guid` | `org_guid` |
| `ibm_container_cluster_data_invalid_region` | `region` |
| `ibm_container_cluster_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_cluster_data_invalid_space_guid` | `space_guid` |
| `ibm_container_cluster_data_invalid_wait_till` | `wait_till` |
| `ibm_container_cluster_data_invalid_wait_till_timeout` | `wait_till_timeout` |

## Example

### Valid

```hcl
data "ibm_container_cluster" "test" {
}
```
