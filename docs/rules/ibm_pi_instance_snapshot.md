# ibm_pi_instance_snapshot

Schema validation for the ibm_pi_instance_snapshot resource.

This rule validates `ibm_pi_instance_snapshot` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_snapshot" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `crn` | no | — |
| `last_update_date` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_description` | no | — |
| `pi_instance_name` | yes | — |
| `pi_snapshot_name` | yes | — |
| `pi_user_tags` | no | — |
| `pi_volume_ids` | no | — |
| `snapshot_id` | no | — |
| `status` | no | — |
| `status_detail` | no | — |
| `volume_snapshots` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_snapshot_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_snapshot_invalid_pi_description` | `pi_description` |
| `ibm_pi_instance_snapshot_invalid_pi_instance_name` | `pi_instance_name` |
| `ibm_pi_instance_snapshot_invalid_pi_snapshot_name` | `pi_snapshot_name` |
| `ibm_pi_instance_snapshot_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_instance_snapshot_invalid_pi_volume_ids` | `pi_volume_ids` |

## Example

### Valid

```hcl
resource "ibm_pi_instance_snapshot" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_name = "example"
  pi_snapshot_name = "example"
}
```
