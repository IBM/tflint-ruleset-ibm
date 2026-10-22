# ibm_pi_volume_snapshot_data

Reads ibm_pi_volume_snapshot. The volume snapshot name.

This rule validates `ibm_pi_volume_snapshot` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_snapshot_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `crn` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_snapshot_id` | yes | — |
| `size` | no | — |
| `status` | no | — |
| `updated_date` | no | — |
| `volume_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_snapshot_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_snapshot_data_invalid_pi_volume_snapshot_id` | `pi_volume_snapshot_id` |

## Example

### Valid

```hcl
data "ibm_pi_volume_snapshot" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_snapshot_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
