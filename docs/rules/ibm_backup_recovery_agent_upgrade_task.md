# ibm_backup_recovery_agent_upgrade_task

Manages ibm_backup_recovery_agent_upgrade_task. Specifies the name of the task.

This rule validates `ibm_backup_recovery_agent_upgrade_task` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Create agent upgrade task params` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_agent_upgrade_task" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `agent_ids` | no | — |
| `agents` | no | — |
| `cluster_version` | no | — |
| `description` | no | API: Specifies the description of the task. |
| `end_time_usecs` | no | — |
| `error` | no | — |
| `is_retryable` | no | — |
| `name` | no | API: Specifies the name of the task. |
| `retried_task_id` | no | — |
| `retry_task_id` | no | — |
| `schedule_end_time_usecs` | no | — |
| `schedule_time_usecs` | no | — |
| `start_time_usecs` | no | — |
| `status` | no | — |
| `type` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_agent_upgrade_task_invalid_agent_ids` | `agent_ids` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_description` | `description` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_name` | `name` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_retry_task_id` | `retry_task_id` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_schedule_end_time_usecs` | `schedule_end_time_usecs` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_schedule_time_usecs` | `schedule_time_usecs` |
| `ibm_backup_recovery_agent_upgrade_task_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_agent_upgrade_task" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
