# ibm_backup_recovery_perform_action_on_protection_group_run_request

Schema validation for the ibm_backup_recovery_perform_action_on_protection_group_run_request resource.

This rule validates `ibm_backup_recovery_perform_action_on_protection_group_run_request` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Update Protection Group Run Request Body.` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_perform_action_on_protection_group_run_request" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `Pause`, `Resume`, `Cancel` |
| `group_id` | yes | — |
| `x_ibm_tenant_id` | yes | — |
| `cancel_params` (block) | no | nested block |
| `cancel_params.archival_task_id` | no | — |
| `cancel_params.cloud_spin_task_id` | no | — |
| `cancel_params.local_task_id` | no | — |
| `cancel_params.object_ids` | no | — |
| `cancel_params.replication_task_id` | no | — |
| `cancel_params.run_id` | yes | — |
| `pause_params` (block) | no | nested block |
| `pause_params.run_id` | yes | — |
| `resume_params` (block) | no | nested block |
| `resume_params.run_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_action` | `action` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_group_id` | `group_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__archival_task_id` | `cancel_params.archival_task_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__cloud_spin_task_id` | `cancel_params.cloud_spin_task_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__local_task_id` | `cancel_params.local_task_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__object_ids` | `cancel_params.object_ids` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__replication_task_id` | `cancel_params.replication_task_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_cancel_params__run_id` | `cancel_params.run_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_pause_params__run_id` | `pause_params.run_id` |
| `ibm_backup_recovery_perform_action_on_protection_group_run_request_invalid_resume_params__run_id` | `resume_params.run_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "test" {
  action = "Pause"
  group_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
