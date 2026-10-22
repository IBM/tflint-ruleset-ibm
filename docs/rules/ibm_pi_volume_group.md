# ibm_pi_volume_group

Schema validation for the ibm_pi_volume_group resource.

This rule validates `ibm_pi_volume_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `consistency_group_name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_consistency_group_name` | no | — |
| `pi_volume_group_name` | no | — |
| `pi_volume_ids` | yes | — |
| `replication_sites` | no | — |
| `replication_status` | no | — |
| `status_description_errors` | no | — |
| `volume_group_id` | no | — |
| `volume_group_status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_group_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_group_invalid_pi_consistency_group_name` | `pi_consistency_group_name` |
| `ibm_pi_volume_group_invalid_pi_volume_group_name` | `pi_volume_group_name` |
| `ibm_pi_volume_group_invalid_pi_volume_ids` | `pi_volume_ids` |

## Example

### Valid

```hcl
resource "ibm_pi_volume_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_ids = "example"
}
```
