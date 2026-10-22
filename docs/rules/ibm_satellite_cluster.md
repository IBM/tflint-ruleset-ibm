# ibm_satellite_cluster

Manages ibm_satellite_cluster. The unique name for the new IBM Cloud Satellite cluster

This rule validates `ibm_satellite_cluster` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_cluster" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `calico_ip_autodetection` | no | — |
| `crn` | no | — |
| `crn_token` | no | — |
| `default_worker_pool_labels` | no | — |
| `disable_public_service_endpoint` | no | — |
| `enable_config_admin` | no | — |
| `entitlement` | no | — |
| `host_labels` | no | — |
| `infrastructure_topology` | no | one of: `single-replica`, `highly-available` |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `kube_version` | no | — |
| `location` | yes | — |
| `master_status` | no | — |
| `master_url` | no | — |
| `name` | yes | — |
| `operating_system` | no | — |
| `patch_version` | no | — |
| `pod_subnet` | no | — |
| `private_service_endpoint_enabled` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint_enabled` | no | — |
| `public_service_endpoint_url` | no | — |
| `pull_secret` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `retry_patch_version` | no | — |
| `service_subnet` | no | — |
| `state` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `wait_for_worker_update` | no | — |
| `worker_count` | no | — |
| `zones` (block) | no | nested block |
| `zones.id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_cluster_invalid_calico_ip_autodetection` | `calico_ip_autodetection` |
| `ibm_satellite_cluster_invalid_crn_token` | `crn_token` |
| `ibm_satellite_cluster_invalid_default_worker_pool_labels` | `default_worker_pool_labels` |
| `ibm_satellite_cluster_invalid_disable_public_service_endpoint` | `disable_public_service_endpoint` |
| `ibm_satellite_cluster_invalid_enable_config_admin` | `enable_config_admin` |
| `ibm_satellite_cluster_invalid_entitlement` | `entitlement` |
| `ibm_satellite_cluster_invalid_host_labels` | `host_labels` |
| `ibm_satellite_cluster_invalid_infrastructure_topology` | `infrastructure_topology` |
| `ibm_satellite_cluster_invalid_kube_version` | `kube_version` |
| `ibm_satellite_cluster_invalid_location` | `location` |
| `ibm_satellite_cluster_invalid_name` | `name` |
| `ibm_satellite_cluster_invalid_operating_system` | `operating_system` |
| `ibm_satellite_cluster_invalid_patch_version` | `patch_version` |
| `ibm_satellite_cluster_invalid_pod_subnet` | `pod_subnet` |
| `ibm_satellite_cluster_invalid_pull_secret` | `pull_secret` |
| `ibm_satellite_cluster_invalid_resource_group_id` | `resource_group_id` |
| `ibm_satellite_cluster_invalid_retry_patch_version` | `retry_patch_version` |
| `ibm_satellite_cluster_invalid_service_subnet` | `service_subnet` |
| `ibm_satellite_cluster_invalid_tags` | `tags` |
| `ibm_satellite_cluster_invalid_wait_for_worker_update` | `wait_for_worker_update` |
| `ibm_satellite_cluster_invalid_worker_count` | `worker_count` |
| `ibm_satellite_cluster_invalid_zones__id` | `zones.id` |

## Example

### Valid

```hcl
resource "ibm_satellite_cluster" "test" {
  location = "example"
  name = "example"
}
```
