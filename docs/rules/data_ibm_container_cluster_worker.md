# ibm_container_cluster_worker_data

Schema validation for the ibm_container_cluster_worker data source.

This rule validates `ibm_container_cluster_worker` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster_worker_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_guid` | no | — |
| `org_guid` | no | — |
| `private_ip` | no | — |
| `private_vlan` | no | — |
| `public_ip` | no | — |
| `public_vlan` | no | — |
| `region` | no | API: The region value. |
| `resource_controller_url` | no | — |
| `resource_group_id` | no | — |
| `space_guid` | no | — |
| `state` | no | — |
| `status` | no | — |
| `worker_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_worker_data_invalid_account_guid` | `account_guid` |
| `ibm_container_cluster_worker_data_invalid_org_guid` | `org_guid` |
| `ibm_container_cluster_worker_data_invalid_region` | `region` |
| `ibm_container_cluster_worker_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_cluster_worker_data_invalid_space_guid` | `space_guid` |
| `ibm_container_cluster_worker_data_invalid_worker_id` | `worker_id` |

## Example

### Valid

```hcl
data "ibm_container_cluster_worker" "test" {
  worker_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
