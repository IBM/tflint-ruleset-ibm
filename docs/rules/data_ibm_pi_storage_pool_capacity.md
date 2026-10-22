# ibm_pi_storage_pool_capacity_data

Schema validation for the ibm_pi_storage_pool_capacity data source.

This rule validates `ibm_pi_storage_pool_capacity` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_storage_pool_capacity_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `max_allocation_size` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_storage_pool` | yes | — |
| `replication_enabled` | no | — |
| `storage_type` | no | — |
| `total_capacity` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_storage_pool_capacity_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_storage_pool_capacity_data_invalid_pi_storage_pool` | `pi_storage_pool` |

## Example

### Valid

```hcl
data "ibm_pi_storage_pool_capacity" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_storage_pool = "example"
}
```
