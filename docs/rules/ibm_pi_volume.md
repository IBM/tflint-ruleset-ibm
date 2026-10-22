# ibm_pi_volume

Schema validation for the ibm_pi_volume resource.

This rule validates `ibm_pi_volume` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auxiliary` | no | — |
| `auxiliary_volume_name` | no | — |
| `consistency_group_name` | no | — |
| `crn` | no | — |
| `delete_on_termination` | no | — |
| `group_id` | no | — |
| `io_throttle_rate` | no | — |
| `master_volume_name` | no | — |
| `mirroring_state` | no | — |
| `out_of_band_deleted` | no | — |
| `pi_affinity_instance` | no | — |
| `pi_affinity_policy` | no | — |
| `pi_affinity_volume` | no | — |
| `pi_anti_affinity_instances` | no | — |
| `pi_anti_affinity_volumes` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_replication_enabled` | no | — |
| `pi_replication_sites` | no | — |
| `pi_user_tags` | no | — |
| `pi_volume_name` | yes | — |
| `pi_volume_pool` | no | — |
| `pi_volume_shareable` | no | — |
| `pi_volume_size` | yes | — |
| `pi_volume_type` | no | one of: `tier0`, `tier1`, `tier3`, `tier5k` |
| `primary_role` | no | — |
| `replication_sites` | no | — |
| `replication_status` | no | — |
| `replication_type` | no | — |
| `volume_id` | no | — |
| `volume_status` | no | — |
| `wwn` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_invalid_pi_affinity_instance` | `pi_affinity_instance` |
| `ibm_pi_volume_invalid_pi_affinity_policy` | `pi_affinity_policy` |
| `ibm_pi_volume_invalid_pi_affinity_volume` | `pi_affinity_volume` |
| `ibm_pi_volume_invalid_pi_anti_affinity_instances` | `pi_anti_affinity_instances` |
| `ibm_pi_volume_invalid_pi_anti_affinity_volumes` | `pi_anti_affinity_volumes` |
| `ibm_pi_volume_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_invalid_pi_replication_enabled` | `pi_replication_enabled` |
| `ibm_pi_volume_invalid_pi_replication_sites` | `pi_replication_sites` |
| `ibm_pi_volume_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_volume_invalid_pi_volume_name` | `pi_volume_name` |
| `ibm_pi_volume_invalid_pi_volume_pool` | `pi_volume_pool` |
| `ibm_pi_volume_invalid_pi_volume_shareable` | `pi_volume_shareable` |
| `ibm_pi_volume_invalid_pi_volume_size` | `pi_volume_size` |
| `ibm_pi_volume_invalid_pi_volume_type` | `pi_volume_type` |

## Example

### Valid

```hcl
resource "ibm_pi_volume" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_name = "example"
  pi_volume_size = 1
}
```
