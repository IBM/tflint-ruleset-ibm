# ibm_satellite_cluster_data

Reads ibm_satellite_cluster. Name or id of the cluster

This rule validates `ibm_satellite_cluster` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_cluster_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `health` | no | — |
| `infrastructure_topology` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `kube_version` | no | — |
| `location` | no | — |
| `name` | yes | — |
| `private_service_endpoint` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint` | no | — |
| `public_service_endpoint_url` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `server_url` | no | — |
| `state` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `worker_count` | no | — |
| `worker_pools` | no | — |
| `workers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_cluster_data_invalid_name` | `name` |
| `ibm_satellite_cluster_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_satellite_cluster" "test" {
  name = "example"
}
```
