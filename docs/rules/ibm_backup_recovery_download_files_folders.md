# ibm_backup_recovery_download_files_folders

Manages ibm_backup_recovery_download_files_folders. Specifies the name of the recovery task.

This rule validates `ibm_backup_recovery_download_files_folders` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_download_files_folders" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `glacier_retrieval_type` | no | — |
| `is_parent_recovery` | no | — |
| `name` | yes | API: Specifies the name of the Recovery. |
| `parent_recovery_id` | no | — |
| `recovery_action` | no | — |
| `recovery_can_tear_down` | no | — |
| `recovery_creation_info` | no | — |
| `recovery_end_time_usecs` | no | — |
| `recovery_is_multi_stage_restore` | no | — |
| `recovery_messages` | no | — |
| `recovery_mssql_params` | no | — |
| `recovery_name` | no | — |
| `recovery_permissions` | no | — |
| `recovery_physical_params` | no | — |
| `recovery_progress_task_id` | no | — |
| `recovery_request_initiator_type` | no | — |
| `recovery_retrieve_archive_tasks` | no | — |
| `recovery_snapshot_environment` | no | — |
| `recovery_start_time_usecs` | no | — |
| `recovery_status` | no | — |
| `recovery_tear_down_message` | no | — |
| `recovery_tear_down_status` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `documents` (block) | no | nested block |
| `documents.is_directory` | no | — |
| `documents.item_id` | yes | — |
| `files_and_folders` (block) | yes | nested block |
| `files_and_folders.absolute_path` | yes | — |
| `files_and_folders.is_directory` | no | — |
| `object` (block) | yes | nested block |
| `object.point_in_time_usecs` | no | — |
| `object.protection_group_id` | no | — |
| `object.protection_group_name` | no | — |
| `object.recover_from_standby` | no | — |
| `object.snapshot_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_download_files_folders_invalid_glacier_retrieval_type` | `glacier_retrieval_type` |
| `ibm_backup_recovery_download_files_folders_invalid_name` | `name` |
| `ibm_backup_recovery_download_files_folders_invalid_parent_recovery_id` | `parent_recovery_id` |
| `ibm_backup_recovery_download_files_folders_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_download_files_folders_invalid_documents__is_directory` | `documents.is_directory` |
| `ibm_backup_recovery_download_files_folders_invalid_documents__item_id` | `documents.item_id` |
| `ibm_backup_recovery_download_files_folders_invalid_files_and_folders__absolute_path` | `files_and_folders.absolute_path` |
| `ibm_backup_recovery_download_files_folders_invalid_files_and_folders__is_directory` | `files_and_folders.is_directory` |
| `ibm_backup_recovery_download_files_folders_invalid_object__point_in_time_usecs` | `object.point_in_time_usecs` |
| `ibm_backup_recovery_download_files_folders_invalid_object__protection_group_id` | `object.protection_group_id` |
| `ibm_backup_recovery_download_files_folders_invalid_object__protection_group_name` | `object.protection_group_name` |
| `ibm_backup_recovery_download_files_folders_invalid_object__recover_from_standby` | `object.recover_from_standby` |
| `ibm_backup_recovery_download_files_folders_invalid_object__snapshot_id` | `object.snapshot_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_download_files_folders" "test" {
  name = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  files_and_folders {
    absolute_path = "example"
  }
  object {
    snapshot_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
