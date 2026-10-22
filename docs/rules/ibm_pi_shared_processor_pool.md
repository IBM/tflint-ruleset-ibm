# ibm_pi_shared_processor_pool

Schema validation for the ibm_pi_shared_processor_pool resource.

This rule validates `ibm_pi_shared_processor_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_shared_processor_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allocated_cores` | no | — |
| `available_cores` | no | — |
| `creation_date` | no | — |
| `crn` | no | — |
| `dedicated_host_id` | no | — |
| `host_id` | no | — |
| `instances` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_host_id` | no | — |
| `pi_shared_processor_pool_host_group` | yes | — |
| `pi_shared_processor_pool_name` | yes | — |
| `pi_shared_processor_pool_placement_group_id` | no | — |
| `pi_shared_processor_pool_placement_groups` | no | — |
| `pi_shared_processor_pool_reserved_cores` | yes | — |
| `pi_user_tags` | no | — |
| `shared_processor_pool_id` | no | — |
| `spp_placement_groups` | no | — |
| `status` | no | — |
| `status_detail` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_shared_processor_pool_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_shared_processor_pool_invalid_pi_host_id` | `pi_host_id` |
| `ibm_pi_shared_processor_pool_invalid_pi_shared_processor_pool_host_group` | `pi_shared_processor_pool_host_group` |
| `ibm_pi_shared_processor_pool_invalid_pi_shared_processor_pool_name` | `pi_shared_processor_pool_name` |
| `ibm_pi_shared_processor_pool_invalid_pi_shared_processor_pool_placement_group_id` | `pi_shared_processor_pool_placement_group_id` |
| `ibm_pi_shared_processor_pool_invalid_pi_shared_processor_pool_placement_groups` | `pi_shared_processor_pool_placement_groups` |
| `ibm_pi_shared_processor_pool_invalid_pi_shared_processor_pool_reserved_cores` | `pi_shared_processor_pool_reserved_cores` |
| `ibm_pi_shared_processor_pool_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_shared_processor_pool_invalid_spp_placement_groups` | `spp_placement_groups` |

## Example

### Valid

```hcl
resource "ibm_pi_shared_processor_pool" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_shared_processor_pool_host_group = "example"
  pi_shared_processor_pool_name = "example"
  pi_shared_processor_pool_reserved_cores = 1
}
```
