# ibm_pi_volume_group_action

Schema validation for the ibm_pi_volume_group_action resource.

This rule validates `ibm_pi_volume_group_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_group_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_group_id` | yes | — |
| `replication_status` | no | — |
| `volume_group_name` | no | — |
| `volume_group_status` | no | — |
| `pi_volume_group_action` (block) | yes | nested block |
| `pi_volume_group_action.reset` (block) | no | nested block |
| `pi_volume_group_action.reset.status` | yes | one of: `available` |
| `pi_volume_group_action.start` (block) | no | nested block |
| `pi_volume_group_action.start.source` | yes | one of: `master`, `aux` |
| `pi_volume_group_action.stop` (block) | no | nested block |
| `pi_volume_group_action.stop.access` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_group_action_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_group_action_invalid_pi_volume_group_id` | `pi_volume_group_id` |
| `ibm_pi_volume_group_action_invalid_pi_volume_group_action__reset__status` | `pi_volume_group_action.reset.status` |
| `ibm_pi_volume_group_action_invalid_pi_volume_group_action__start__source` | `pi_volume_group_action.start.source` |
| `ibm_pi_volume_group_action_invalid_pi_volume_group_action__stop__access` | `pi_volume_group_action.stop.access` |

## Example

### Valid

```hcl
resource "ibm_pi_volume_group_action" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_group_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_group_action {
  }
}
```
