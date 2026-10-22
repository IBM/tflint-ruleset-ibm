# ibm_backup_recovery_update_protection_group_run_request

Schema validation for the ibm_backup_recovery_update_protection_group_run_request resource.

This rule validates `ibm_backup_recovery_update_protection_group_run_request` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Group` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_update_protection_group_run_request" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `failed_runs` | no | — |
| `group_id` | yes | — |
| `run_id` | no | — |
| `successful_run_ids` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `update_protection_group_run_params` (block) | yes | nested block |
| `update_protection_group_run_params.run_id` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.archival_target_type` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.copy_only_fully_successful` | no | — |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.id` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention` (block) | no | nested block |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention.duration` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention.unit` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention.data_lock_config` (block) | no | nested block |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.archival_target_type` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.data_lock` | no | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.days_to_keep` | no | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.delete_snapshot` | no | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.enable_legal_hold` | no | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.id` | yes | — |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.name` | no | API: Specifies the name of the user group. |
| `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.resync` | no | — |
| `update_protection_group_run_params.local_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.local_snapshot_config.data_lock` | no | — |
| `update_protection_group_run_params.local_snapshot_config.days_to_keep` | no | — |
| `update_protection_group_run_params.local_snapshot_config.delete_snapshot` | no | — |
| `update_protection_group_run_params.local_snapshot_config.enable_legal_hold` | no | — |
| `update_protection_group_run_params.replication_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.id` | yes | — |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention` (block) | no | nested block |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention.duration` | yes | — |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention.unit` | yes | — |
| `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention.data_lock_config` (block) | no | nested block |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config` (block) | no | nested block |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.data_lock` | no | — |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.days_to_keep` | no | — |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.delete_snapshot` | no | — |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.enable_legal_hold` | no | — |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.id` | yes | — |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.name` | no | API: Specifies the name of the user group. |
| `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.resync` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_group_id` | `group_id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__run_id` | `update_protection_group_run_params.run_id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__new_snapshot_config__archival_target_type` | `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.archival_target_type` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__new_snapshot_config__copy_only_fully_successful` | `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.copy_only_fully_successful` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__new_snapshot_config__id` | `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__new_snapshot_config__retention__duration` | `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention.duration` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__new_snapshot_config__retention__unit` | `update_protection_group_run_params.archival_snapshot_config.new_snapshot_config.retention.unit` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__archival_target_type` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.archival_target_type` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__data_lock` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.data_lock` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__days_to_keep` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.days_to_keep` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__delete_snapshot` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.delete_snapshot` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__enable_legal_hold` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.enable_legal_hold` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__id` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__name` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.name` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__archival_snapshot_config__update_existing_snapshot_config__resync` | `update_protection_group_run_params.archival_snapshot_config.update_existing_snapshot_config.resync` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__local_snapshot_config__data_lock` | `update_protection_group_run_params.local_snapshot_config.data_lock` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__local_snapshot_config__days_to_keep` | `update_protection_group_run_params.local_snapshot_config.days_to_keep` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__local_snapshot_config__delete_snapshot` | `update_protection_group_run_params.local_snapshot_config.delete_snapshot` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__local_snapshot_config__enable_legal_hold` | `update_protection_group_run_params.local_snapshot_config.enable_legal_hold` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__new_snapshot_config__id` | `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__new_snapshot_config__retention__duration` | `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention.duration` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__new_snapshot_config__retention__unit` | `update_protection_group_run_params.replication_snapshot_config.new_snapshot_config.retention.unit` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__data_lock` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.data_lock` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__days_to_keep` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.days_to_keep` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__delete_snapshot` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.delete_snapshot` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__enable_legal_hold` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.enable_legal_hold` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__id` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.id` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__name` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.name` |
| `ibm_backup_recovery_update_protection_group_run_request_invalid_update_protection_group_run_params__replication_snapshot_config__update_existing_snapshot_config__resync` | `update_protection_group_run_params.replication_snapshot_config.update_existing_snapshot_config.resync` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_update_protection_group_run_request" "test" {
  group_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  update_protection_group_run_params {
    run_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
