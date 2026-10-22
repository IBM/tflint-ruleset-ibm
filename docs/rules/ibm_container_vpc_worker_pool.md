# ibm_container_vpc_worker_pool

Schema validation for the ibm_container_vpc_worker_pool resource.

This rule validates `ibm_container_vpc_worker_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Worker` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vpc_worker_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `autoscale_enabled` | no | — |
| `cluster` | yes | API: The cluster value. |
| `crk` | no | — |
| `entitlement` | no | API: The entitlement value. |
| `flavor` | yes | API: The flavor value. |
| `host_pool_id` | no | — |
| `import_on_create` | no | — |
| `kms_account_id` | no | — |
| `kms_instance_id` | no | — |
| `labels` | no | — |
| `operating_system` | no | — |
| `orphan_on_delete` | no | — |
| `resource_controller_url` | no | — |
| `resource_group_id` | no | — |
| `secondary_storage` | no | — |
| `security_groups` | no | — |
| `vpc_id` | yes | — |
| `worker_count` | yes | — |
| `worker_pool_id` | no | — |
| `worker_pool_name` | yes | — |
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
| `ibm_container_vpc_worker_pool_invalid_cluster` | `cluster` |
| `ibm_container_vpc_worker_pool_invalid_crk` | `crk` |
| `ibm_container_vpc_worker_pool_invalid_entitlement` | `entitlement` |
| `ibm_container_vpc_worker_pool_invalid_flavor` | `flavor` |
| `ibm_container_vpc_worker_pool_invalid_host_pool_id` | `host_pool_id` |
| `ibm_container_vpc_worker_pool_invalid_import_on_create` | `import_on_create` |
| `ibm_container_vpc_worker_pool_invalid_kms_account_id` | `kms_account_id` |
| `ibm_container_vpc_worker_pool_invalid_kms_instance_id` | `kms_instance_id` |
| `ibm_container_vpc_worker_pool_invalid_labels` | `labels` |
| `ibm_container_vpc_worker_pool_invalid_operating_system` | `operating_system` |
| `ibm_container_vpc_worker_pool_invalid_orphan_on_delete` | `orphan_on_delete` |
| `ibm_container_vpc_worker_pool_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_worker_pool_invalid_secondary_storage` | `secondary_storage` |
| `ibm_container_vpc_worker_pool_invalid_security_groups` | `security_groups` |
| `ibm_container_vpc_worker_pool_invalid_vpc_id` | `vpc_id` |
| `ibm_container_vpc_worker_pool_invalid_worker_count` | `worker_count` |
| `ibm_container_vpc_worker_pool_invalid_worker_pool_name` | `worker_pool_name` |
| `ibm_container_vpc_worker_pool_invalid_taints__effect` | `taints.effect` |
| `ibm_container_vpc_worker_pool_invalid_taints__key` | `taints.key` |
| `ibm_container_vpc_worker_pool_invalid_taints__value` | `taints.value` |
| `ibm_container_vpc_worker_pool_invalid_zones__name` | `zones.name` |
| `ibm_container_vpc_worker_pool_invalid_zones__subnet_id` | `zones.subnet_id` |

## Example

### Valid

```hcl
resource "ibm_container_vpc_worker_pool" "test" {
  cluster = "aaaaaaa"
  flavor = "aaaaaaa"
  vpc_id = "01234567-89ab-cdef-0123-456789abcdef"
  worker_count = 1
  worker_pool_name = "example"
  zones {
    name = "aaaaaaa"
    subnet_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
