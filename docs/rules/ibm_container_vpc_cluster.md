# ibm_container_vpc_cluster

Manages ibm_container_vpc_cluster. The cluster name

This rule validates `ibm_container_vpc_cluster` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vpc_cluster" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `albs` | no | — |
| `cos_instance_crn` | no | crn format |
| `crk` | no | — |
| `crn` | no | — |
| `disable_outbound_traffic_protection` | no | — |
| `disable_public_service_endpoint` | no | — |
| `enable_secure_by_default` | no | — |
| `entitlement` | no | API: The entitlement value. |
| `flavor` | yes | API: The flavor value. |
| `force_delete_storage` | no | — |
| `host_pool_id` | no | — |
| `image_security_enforcement` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `kms_account_id` | no | — |
| `kms_instance_id` | no | — |
| `kube_version` | no | — |
| `master_status` | no | — |
| `master_url` | no | — |
| `name` | yes | API: The name value. |
| `network_plugin` | no | — |
| `offering` | no | one of: `kubernetes`, `openshift`, `openshift-vs` |
| `operating_system` | no | — |
| `patch_version` | no | — |
| `pod_subnet` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint_url` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `retry_patch_version` | no | — |
| `secondary_storage` | no | — |
| `security_groups` | no | — |
| `service_subnet` | no | — |
| `state` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `update_all_workers` | no | — |
| `vpc_id` | yes | — |
| `vpe_service_endpoint_url` | no | — |
| `wait_for_worker_update` | no | — |
| `wait_till` | no | one of: `MasterNodeReady`, `OneWorkerNodeReady`, `IngressReady`, `normal` |
| `worker_count` | no | — |
| `worker_labels` | no | — |
| `kms_config` (block) | no | nested block |
| `kms_config.account_id` | no | API: The account id value. |
| `kms_config.crk_id` | yes | API: The crk id value. |
| `kms_config.instance_id` | yes | API: The instance id value. |
| `kms_config.private_endpoint` | no | — |
| `kms_config.wait_for_apply` | no | — |
| `taints` (block) | no | nested block |
| `taints.effect` | yes | — |
| `taints.key` | yes | API: The key to select. |
| `taints.value` | yes | API: The value value. |
| `zones` (block) | yes | nested block |
| `zones.name` | yes | API: The name value. |
| `zones.subnet_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_cluster_invalid_cos_instance_crn` | `cos_instance_crn` |
| `ibm_container_vpc_cluster_invalid_crk` | `crk` |
| `ibm_container_vpc_cluster_invalid_disable_outbound_traffic_protection` | `disable_outbound_traffic_protection` |
| `ibm_container_vpc_cluster_invalid_disable_public_service_endpoint` | `disable_public_service_endpoint` |
| `ibm_container_vpc_cluster_invalid_enable_secure_by_default` | `enable_secure_by_default` |
| `ibm_container_vpc_cluster_invalid_entitlement` | `entitlement` |
| `ibm_container_vpc_cluster_invalid_flavor` | `flavor` |
| `ibm_container_vpc_cluster_invalid_force_delete_storage` | `force_delete_storage` |
| `ibm_container_vpc_cluster_invalid_host_pool_id` | `host_pool_id` |
| `ibm_container_vpc_cluster_invalid_image_security_enforcement` | `image_security_enforcement` |
| `ibm_container_vpc_cluster_invalid_kms_account_id` | `kms_account_id` |
| `ibm_container_vpc_cluster_invalid_kms_instance_id` | `kms_instance_id` |
| `ibm_container_vpc_cluster_invalid_kube_version` | `kube_version` |
| `ibm_container_vpc_cluster_invalid_name` | `name` |
| `ibm_container_vpc_cluster_invalid_network_plugin` | `network_plugin` |
| `ibm_container_vpc_cluster_invalid_offering` | `offering` |
| `ibm_container_vpc_cluster_invalid_operating_system` | `operating_system` |
| `ibm_container_vpc_cluster_invalid_patch_version` | `patch_version` |
| `ibm_container_vpc_cluster_invalid_pod_subnet` | `pod_subnet` |
| `ibm_container_vpc_cluster_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_cluster_invalid_retry_patch_version` | `retry_patch_version` |
| `ibm_container_vpc_cluster_invalid_secondary_storage` | `secondary_storage` |
| `ibm_container_vpc_cluster_invalid_security_groups` | `security_groups` |
| `ibm_container_vpc_cluster_invalid_service_subnet` | `service_subnet` |
| `ibm_container_vpc_cluster_invalid_tags` | `tags` |
| `ibm_container_vpc_cluster_invalid_update_all_workers` | `update_all_workers` |
| `ibm_container_vpc_cluster_invalid_vpc_id` | `vpc_id` |
| `ibm_container_vpc_cluster_invalid_wait_for_worker_update` | `wait_for_worker_update` |
| `ibm_container_vpc_cluster_invalid_wait_till` | `wait_till` |
| `ibm_container_vpc_cluster_invalid_worker_count` | `worker_count` |
| `ibm_container_vpc_cluster_invalid_worker_labels` | `worker_labels` |
| `ibm_container_vpc_cluster_invalid_kms_config__account_id` | `kms_config.account_id` |
| `ibm_container_vpc_cluster_invalid_kms_config__crk_id` | `kms_config.crk_id` |
| `ibm_container_vpc_cluster_invalid_kms_config__instance_id` | `kms_config.instance_id` |
| `ibm_container_vpc_cluster_invalid_kms_config__private_endpoint` | `kms_config.private_endpoint` |
| `ibm_container_vpc_cluster_invalid_kms_config__wait_for_apply` | `kms_config.wait_for_apply` |
| `ibm_container_vpc_cluster_invalid_taints__effect` | `taints.effect` |
| `ibm_container_vpc_cluster_invalid_taints__key` | `taints.key` |
| `ibm_container_vpc_cluster_invalid_taints__value` | `taints.value` |
| `ibm_container_vpc_cluster_invalid_zones__name` | `zones.name` |
| `ibm_container_vpc_cluster_invalid_zones__subnet_id` | `zones.subnet_id` |

## Example

### Valid

```hcl
resource "ibm_container_vpc_cluster" "test" {
  flavor = "aaaaaaa"
  name = "aaaaaaa"
  vpc_id = "01234567-89ab-cdef-0123-456789abcdef"
  zones {
    name = "aaaaaaa"
    subnet_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
