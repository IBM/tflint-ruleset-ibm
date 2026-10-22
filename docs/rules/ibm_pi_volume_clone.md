# ibm_pi_volume_clone

Schema validation for the ibm_pi_volume_clone resource.

This rule validates `ibm_pi_volume_clone` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_clone" {
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
| `pi_replication_enabled` | no | — |
| `pi_target_storage_tier` | no | — |
| `pi_user_tags` | no | — |
| `pi_volume_clone_name` | yes | — |
| `pi_volume_ids` | yes | — |
| `status` | no | — |
| `task_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_clone_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_clone_invalid_pi_replication_enabled` | `pi_replication_enabled` |
| `ibm_pi_volume_clone_invalid_pi_target_storage_tier` | `pi_target_storage_tier` |
| `ibm_pi_volume_clone_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_volume_clone_invalid_pi_volume_clone_name` | `pi_volume_clone_name` |
| `ibm_pi_volume_clone_invalid_pi_volume_ids` | `pi_volume_ids` |

## Example

### Valid

```hcl
resource "ibm_pi_volume_clone" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_clone_name = "example"
  pi_volume_ids = "example"
}
```
