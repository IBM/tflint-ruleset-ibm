# ibm_pi_shared_processor_pool_data

Reads ibm_pi_shared_processor_pool. The name of the shared processor pool.

This rule validates `ibm_pi_shared_processor_pool` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SharedProcessorPoolServer` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_shared_processor_pool_data" {
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
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_shared_processor_pool_id` | yes | — |
| `reserved_cores` | no | — |
| `status` | no | — |
| `status_detail` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_shared_processor_pool_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_shared_processor_pool_data_invalid_pi_shared_processor_pool_id` | `pi_shared_processor_pool_id` |

## Example

### Valid

```hcl
data "ibm_pi_shared_processor_pool" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_shared_processor_pool_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
