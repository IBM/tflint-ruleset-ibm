# ibm_pi_cloud_instance_data

Schema validation for the ibm_pi_cloud_instance data source.

This rule validates `ibm_pi_cloud_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_cloud_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capabilities` | no | — |
| `enabled` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pvm_instances` | no | — |
| `region` | no | — |
| `tenant_id` | no | — |
| `total_instances` | no | — |
| `total_memory_consumed` | no | — |
| `total_processors_consumed` | no | — |
| `total_ssd_storage_consumed` | no | — |
| `total_standard_storage_consumed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_cloud_instance_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_cloud_instance" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
