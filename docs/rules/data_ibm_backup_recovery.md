# ibm_backup_recovery_data

Reads ibm_backup_recovery. Specifies the name of the Recovery.

This rule validates `ibm_backup_recovery` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `can_tear_down` | no | — |
| `creation_info` | no | — |
| `end_time_usecs` | no | — |
| `is_multi_stage_restore` | no | — |
| `is_parent_recovery` | no | — |
| `kubernetes_params` | no | — |
| `messages` | no | — |
| `mssql_params` | no | — |
| `name` | no | — |
| `parent_recovery_id` | no | — |
| `permissions` | no | — |
| `physical_params` | no | — |
| `progress_task_id` | no | — |
| `recovery_action` | no | — |
| `recovery_id` | yes | — |
| `retrieve_archive_tasks` | no | — |
| `snapshot_environment` | no | — |
| `start_time_usecs` | no | — |
| `status` | no | — |
| `tear_down_message` | no | — |
| `tear_down_status` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_data_invalid_recovery_id` | `recovery_id` |
| `ibm_backup_recovery_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery" "test" {
  recovery_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
