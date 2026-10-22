# ibm_pi_volume_data

Reads ibm_pi_volume. The name of the volume.

This rule validates `ibm_pi_volume` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auxiliary` | no | — |
| `auxiliary_volume_name` | no | — |
| `bootable` | no | — |
| `consistency_group_name` | no | — |
| `creation_date` | no | — |
| `crn` | no | — |
| `disk_type` | no | — |
| `freeze_time` | no | — |
| `group_id` | no | — |
| `io_throttle_rate` | no | — |
| `last_update_date` | no | — |
| `master_volume_name` | no | — |
| `mirroring_state` | no | — |
| `name` | no | — |
| `out_of_band_deleted` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_id` | no | — |
| `pi_volume_name` | no | — |
| `primary_role` | no | — |
| `replication_enabled` | no | — |
| `replication_sites` | no | — |
| `replication_status` | no | — |
| `replication_type` | no | — |
| `shreable` | no | — |
| `size` | no | — |
| `state` | no | — |
| `user_tags` | no | — |
| `volume_pool` | no | — |
| `volume_type` | no | — |
| `wwn` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_data_invalid_pi_volume_id` | `pi_volume_id` |
| `ibm_pi_volume_data_invalid_pi_volume_name` | `pi_volume_name` |

## Example

### Valid

```hcl
data "ibm_pi_volume" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
