# ibm_container_cluster

Manages ibm_container_cluster. The cluster name

This rule validates `ibm_container_cluster` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `albs` | no | — |
| `crn` | no | — |
| `datacenter` | yes | API: The datacenter value. |
| `default_pool_size` | no | — |
| `disk_encryption` | no | — |
| `entitlement` | no | API: The entitlement value. |
| `force_delete_storage` | no | — |
| `gateway_enabled` | no | — |
| `hardware` | yes | one of: `shared`, `dedicated` |
| `image_security_enforcement` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `kube_version` | no | — |
| `labels` | no | — |
| `machine_type` | no | — |
| `name` | yes | API: The name value. |
| `no_subnet` | no | — |
| `operating_system` | no | — |
| `patch_version` | no | — |
| `pod_subnet` | no | — |
| `private_service_endpoint` | no | — |
| `private_service_endpoint_url` | no | — |
| `private_vlan_id` | no | — |
| `public_service_endpoint` | no | — |
| `public_service_endpoint_url` | no | — |
| `public_vlan_id` | no | — |
| `region` | no | API: The region value. |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `retry_patch_version` | no | — |
| `server_url` | no | — |
| `service_subnet` | no | — |
| `subnet_id` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `update_all_workers` | no | — |
| `wait_for_worker_update` | no | — |
| `wait_till` | no | one of: `MasterNodeReady`, `OneWorkerNodeReady`, `IngressReady`, `normal` |
| `worker_pools` | no | — |
| `kms_config` (block) | no | nested block |
| `kms_config.account_id` | no | API: The account id value. |
| `kms_config.crk_id` | yes | API: The crk id value. |
| `kms_config.instance_id` | yes | API: The instance id value. |
| `kms_config.private_endpoint` | no | — |
| `taints` (block) | no | nested block |
| `taints.effect` | yes | — |
| `taints.key` | yes | API: The key to select. |
| `taints.value` | yes | API: The value value. |
| `webhook` (block) | no | nested block |
| `webhook.level` | yes | API: The notification level that you want to use for the webhook. |
| `webhook.type` | yes | one of: `slack` |
| `webhook.url` | yes | API: URL is a Key Protect endpoint. To get the endpoint, see http://ibm.biz/kp-endpoints |
| `workers_info` (block) | no | nested block |
| `workers_info.id` | no | API: The id value. |
| `workers_info.version` | no | API: The Kubernetes version to use for the update. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_invalid_datacenter` | `datacenter` |
| `ibm_container_cluster_invalid_default_pool_size` | `default_pool_size` |
| `ibm_container_cluster_invalid_disk_encryption` | `disk_encryption` |
| `ibm_container_cluster_invalid_entitlement` | `entitlement` |
| `ibm_container_cluster_invalid_force_delete_storage` | `force_delete_storage` |
| `ibm_container_cluster_invalid_gateway_enabled` | `gateway_enabled` |
| `ibm_container_cluster_invalid_hardware` | `hardware` |
| `ibm_container_cluster_invalid_image_security_enforcement` | `image_security_enforcement` |
| `ibm_container_cluster_invalid_kube_version` | `kube_version` |
| `ibm_container_cluster_invalid_labels` | `labels` |
| `ibm_container_cluster_invalid_machine_type` | `machine_type` |
| `ibm_container_cluster_invalid_name` | `name` |
| `ibm_container_cluster_invalid_no_subnet` | `no_subnet` |
| `ibm_container_cluster_invalid_operating_system` | `operating_system` |
| `ibm_container_cluster_invalid_patch_version` | `patch_version` |
| `ibm_container_cluster_invalid_pod_subnet` | `pod_subnet` |
| `ibm_container_cluster_invalid_private_service_endpoint` | `private_service_endpoint` |
| `ibm_container_cluster_invalid_private_vlan_id` | `private_vlan_id` |
| `ibm_container_cluster_invalid_public_service_endpoint` | `public_service_endpoint` |
| `ibm_container_cluster_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_container_cluster_invalid_region` | `region` |
| `ibm_container_cluster_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_cluster_invalid_retry_patch_version` | `retry_patch_version` |
| `ibm_container_cluster_invalid_service_subnet` | `service_subnet` |
| `ibm_container_cluster_invalid_subnet_id` | `subnet_id` |
| `ibm_container_cluster_invalid_tags` | `tags` |
| `ibm_container_cluster_invalid_update_all_workers` | `update_all_workers` |
| `ibm_container_cluster_invalid_wait_for_worker_update` | `wait_for_worker_update` |
| `ibm_container_cluster_invalid_wait_till` | `wait_till` |
| `ibm_container_cluster_invalid_kms_config__account_id` | `kms_config.account_id` |
| `ibm_container_cluster_invalid_kms_config__crk_id` | `kms_config.crk_id` |
| `ibm_container_cluster_invalid_kms_config__instance_id` | `kms_config.instance_id` |
| `ibm_container_cluster_invalid_kms_config__private_endpoint` | `kms_config.private_endpoint` |
| `ibm_container_cluster_invalid_taints__effect` | `taints.effect` |
| `ibm_container_cluster_invalid_taints__key` | `taints.key` |
| `ibm_container_cluster_invalid_taints__value` | `taints.value` |
| `ibm_container_cluster_invalid_webhook__level` | `webhook.level` |
| `ibm_container_cluster_invalid_webhook__type` | `webhook.type` |
| `ibm_container_cluster_invalid_webhook__url` | `webhook.url` |
| `ibm_container_cluster_invalid_workers_info__id` | `workers_info.id` |
| `ibm_container_cluster_invalid_workers_info__version` | `workers_info.version` |

## Example

### Valid

```hcl
resource "ibm_container_cluster" "test" {
  datacenter = "aaaaaaa"
  hardware = "shared"
  name = "aaaaaaa"
}
```
