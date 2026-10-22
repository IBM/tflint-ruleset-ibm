# ibm_pi_volume_clone_data

Schema validation for the ibm_pi_volume_clone data source.

This rule validates `ibm_pi_volume_clone` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_clone_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `clone_volumes` | no | — |
| `failure_reason` | no | — |
| `percent_complete` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_clone_task_id` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_clone_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_clone_data_invalid_pi_volume_clone_task_id` | `pi_volume_clone_task_id` |

## Example

### Valid

```hcl
data "ibm_pi_volume_clone" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_clone_task_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
