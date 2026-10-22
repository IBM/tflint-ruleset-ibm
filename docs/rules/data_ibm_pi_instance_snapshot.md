# ibm_pi_instance_snapshot_data

Reads ibm_pi_instance_snapshot. The name of the Power Systems Virtual Machine instance snapshot.

This rule validates `ibm_pi_instance_snapshot` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_snapshot_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `creation_date` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `last_updated_date` | no | — |
| `name` | no | — |
| `percent_complete` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_snapshot_id` | yes | — |
| `status` | no | — |
| `status_detail` | no | — |
| `user_tags` | no | — |
| `volume_snapshots` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_snapshot_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_snapshot_data_invalid_pi_snapshot_id` | `pi_snapshot_id` |

## Example

### Valid

```hcl
data "ibm_pi_instance_snapshot" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_snapshot_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
