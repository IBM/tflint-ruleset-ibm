# ibm_backup_recovery_restore_points

Schema validation for the ibm_backup_recovery_restore_points resource.

This rule validates `ibm_backup_recovery_restore_points` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_restore_points" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `end_time_usecs` | yes | — |
| `environment` | yes | one of: `kPhysical`, `kSQL` |
| `full_snapshot_info` | no | — |
| `protection_group_ids` | yes | — |
| `source_id` | no | — |
| `start_time_usecs` | yes | — |
| `time_range_info` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_restore_points_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_restore_points_invalid_environment` | `environment` |
| `ibm_backup_recovery_restore_points_invalid_protection_group_ids` | `protection_group_ids` |
| `ibm_backup_recovery_restore_points_invalid_source_id` | `source_id` |
| `ibm_backup_recovery_restore_points_invalid_start_time_usecs` | `start_time_usecs` |
| `ibm_backup_recovery_restore_points_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_restore_points" "test" {
  end_time_usecs = 1
  environment = "kPhysical"
  protection_group_ids = "example"
  start_time_usecs = 1
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
