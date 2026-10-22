# ibm_backup_recovery_agent_upgrade_tasks_data

Schema validation for the ibm_backup_recovery_agent_upgrade_tasks data source.

This rule validates `ibm_backup_recovery_agent_upgrade_tasks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_agent_upgrade_tasks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ids` | no | — |
| `tasks` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_agent_upgrade_tasks_data_invalid_ids` | `ids` |
| `ibm_backup_recovery_agent_upgrade_tasks_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_agent_upgrade_tasks" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
