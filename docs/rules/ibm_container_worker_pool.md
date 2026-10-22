# ibm_container_worker_pool

Schema validation for the ibm_container_worker_pool resource.

This rule validates `ibm_container_worker_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Worker` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_worker_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `autoscale_enabled` | no | — |
| `cluster` | yes | API: The cluster value. |
| `disk_encryption` | no | — |
| `entitlement` | no | API: The entitlement value. |
| `hardware` | no | one of: `shared`, `dedicated` |
| `import_on_create` | no | — |
| `labels` | no | — |
| `machine_type` | yes | — |
| `operating_system` | no | — |
| `orphan_on_delete` | no | — |
| `region` | no | API: The region value. |
| `resource_controller_url` | no | — |
| `resource_group_id` | no | — |
| `size_per_zone` | yes | zone format |
| `state` | no | — |
| `worker_pool_id` | no | — |
| `worker_pool_name` | yes | — |
| `zones` | no | — |
| `taints` (block) | no | nested block |
| `taints.effect` | yes | — |
| `taints.key` | yes | API: The key to select. |
| `taints.value` | yes | API: The value value. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_worker_pool_invalid_cluster` | `cluster` |
| `ibm_container_worker_pool_invalid_disk_encryption` | `disk_encryption` |
| `ibm_container_worker_pool_invalid_entitlement` | `entitlement` |
| `ibm_container_worker_pool_invalid_hardware` | `hardware` |
| `ibm_container_worker_pool_invalid_import_on_create` | `import_on_create` |
| `ibm_container_worker_pool_invalid_labels` | `labels` |
| `ibm_container_worker_pool_invalid_machine_type` | `machine_type` |
| `ibm_container_worker_pool_invalid_operating_system` | `operating_system` |
| `ibm_container_worker_pool_invalid_orphan_on_delete` | `orphan_on_delete` |
| `ibm_container_worker_pool_invalid_region` | `region` |
| `ibm_container_worker_pool_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_worker_pool_invalid_size_per_zone` | `size_per_zone` |
| `ibm_container_worker_pool_invalid_worker_pool_name` | `worker_pool_name` |
| `ibm_container_worker_pool_invalid_taints__effect` | `taints.effect` |
| `ibm_container_worker_pool_invalid_taints__key` | `taints.key` |
| `ibm_container_worker_pool_invalid_taints__value` | `taints.value` |

## Example

### Valid

```hcl
resource "ibm_container_worker_pool" "test" {
  cluster = "aaaaaaa"
  machine_type = "example"
  size_per_zone = "us-south-1"
  worker_pool_name = "example"
}
```
