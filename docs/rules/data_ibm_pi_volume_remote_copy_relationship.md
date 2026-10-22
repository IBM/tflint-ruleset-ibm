# ibm_pi_volume_remote_copy_relationship_data

Reads ibm_pi_volume_remote_copy_relationship. The remote copy relationship name.

This rule validates `ibm_pi_volume_remote_copy_relationship` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_remote_copy_relationship_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auxiliary_changed_volume_name` | no | — |
| `auxiliary_volume_name` | no | — |
| `consistency_group_name` | no | — |
| `copy_type` | no | — |
| `cycle_period_seconds` | no | — |
| `cycling_mode` | no | — |
| `freeze_time` | no | — |
| `master_changed_volume_name` | no | — |
| `master_volume_name` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_id` | yes | — |
| `primary_role` | no | — |
| `progress` | no | — |
| `remote_copy_id` | no | — |
| `state` | no | — |
| `synchronized` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_remote_copy_relationship_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_remote_copy_relationship_data_invalid_pi_volume_id` | `pi_volume_id` |

## Example

### Valid

```hcl
data "ibm_pi_volume_remote_copy_relationship" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
