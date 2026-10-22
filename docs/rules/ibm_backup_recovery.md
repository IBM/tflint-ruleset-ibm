# ibm_backup_recovery

Manages ibm_backup_recovery. Specifies the name of the Recovery.

This rule validates `ibm_backup_recovery` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery" {
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
| `messages` | no | — |
| `name` | yes | API: Specifies the name of the Recovery. |
| `parent_recovery_id` | no | — |
| `permissions` | no | — |
| `progress_task_id` | no | — |
| `recovery_action` | no | — |
| `recovery_id` | no | — |
| `request_initiator_type` | no | — |
| `retrieve_archive_tasks` | no | — |
| `snapshot_environment` | yes | — |
| `start_time_usecs` | no | — |
| `status` | no | — |
| `tear_down_message` | no | — |
| `tear_down_status` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `kubernetes_params` (block) | no | nested block |
| `kubernetes_params.recovery_action` | yes | — |
| `kubernetes_params.download_file_and_folder_params` (block) | no | nested block |
| `kubernetes_params.download_file_and_folder_params.download_file_path` | no | — |
| `kubernetes_params.download_file_and_folder_params.expiry_time_usecs` | no | — |
| `kubernetes_params.download_file_and_folder_params.files_and_folders` (block) | no | nested block |
| `kubernetes_params.download_file_and_folder_params.files_and_folders.absolute_path` | yes | — |
| `kubernetes_params.download_file_and_folder_params.files_and_folders.is_directory` | no | — |
| `kubernetes_params.download_file_and_folder_params.files_and_folders.is_view_file_recovery` | no | — |
| `kubernetes_params.objects` (block) | no | nested block |
| `kubernetes_params.objects.point_in_time_usecs` | no | — |
| `kubernetes_params.objects.protection_group_id` | no | — |
| `kubernetes_params.objects.protection_group_name` | no | — |
| `kubernetes_params.objects.recover_from_standby` | no | — |
| `kubernetes_params.objects.snapshot_id` | yes | — |
| `kubernetes_params.recover_file_and_folder_params` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.target_environment` | yes | — |
| `kubernetes_params.recover_file_and_folder_params.files_and_folders` (block) | yes | nested block |
| `kubernetes_params.recover_file_and_folder_params.files_and_folders.absolute_path` | yes | — |
| `kubernetes_params.recover_file_and_folder_params.files_and_folders.is_directory` | no | — |
| `kubernetes_params.recover_file_and_folder_params.files_and_folders.is_view_file_recovery` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.continue_on_error` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.overwrite_existing` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.preserve_attributes` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.recover_to_original_target` | yes | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config.absolute_path` | yes | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config.target_namespace` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config.target_pvc` (block) | yes | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config.target_source` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.original_target_config` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.original_target_config.alternate_path` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.original_target_config.recover_to_original_path` | yes | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.vlan_config` (block) | no | nested block |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.vlan_config.disable_vlan` | no | — |
| `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `kubernetes_params.recover_namespace_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.target_environment` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.preserve_mac_address` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_pvcs_only` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.skip_cluster_compatibility_check` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.label_combination_method` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.label_filter_entity_type` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.label_vector` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.excluded_pvcs` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.excluded_pvcs.metadata_only` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.label_combination_method` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.label_filter_entity_type` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.label_vector` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.point_in_time_usecs` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.protection_group_id` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.protection_group_name` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.recover_from_standby` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.recover_pvcs_only` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.snapshot_id` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.unbind_pvcs` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.use_instant_recovery` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.exclude_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.include_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.storage_class` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_cluster_scoped_resources` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_cluster_scoped_resources.snapshot_id` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.archival_target_id` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_id` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_instance_id` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_run_id` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_region_migration_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_region_migration_params.current_value` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_region_migration_params.new_value` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_target_config` (block) | yes | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_target_config.recover_to_new_source` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_target_config.new_source_config` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_zone_migration_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_zone_migration_params.current_value` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_zone_migration_params.new_value` | yes | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.rename_recovered_namespaces_params` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.rename_recovered_namespaces_params.prefix` | no | API: Specifies the prefix string to be added to recovered or cloned object names. |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.rename_recovered_namespaces_params.suffix` | no | API: Specifies the suffix string to be added to recovered or cloned object names. |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.storage_class` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.storage_class.use_storage_class_mapping` | no | — |
| `kubernetes_params.recover_namespace_params.kubernetes_target_params.storage_class.storage_class_mapping` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.vlan_config` (block) | no | nested block |
| `kubernetes_params.recover_namespace_params.vlan_config.disable_vlan` | no | — |
| `kubernetes_params.recover_namespace_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `mssql_params` (block) | no | nested block |
| `mssql_params.recovery_action` | yes | — |
| `mssql_params.recover_app_params` (block) | no | nested block |
| `mssql_params.recover_app_params.is_encrypted` | no | — |
| `mssql_params.recover_app_params.point_in_time_usecs` | no | — |
| `mssql_params.recover_app_params.protection_group_id` | no | — |
| `mssql_params.recover_app_params.protection_group_name` | no | — |
| `mssql_params.recover_app_params.recover_from_standby` | no | — |
| `mssql_params.recover_app_params.snapshot_id` | yes | — |
| `mssql_params.recover_app_params.target_environment` | yes | — |
| `mssql_params.recover_app_params.aag_info` (block) | no | nested block |
| `mssql_params.recover_app_params.aag_info.name` | no | API: Specifies the name of the Recovery. |
| `mssql_params.recover_app_params.aag_info.object_id` | no | — |
| `mssql_params.recover_app_params.host_info` (block) | no | nested block |
| `mssql_params.recover_app_params.host_info.environment` | no | one of: `kPhysical`, `kSQL` |
| `mssql_params.recover_app_params.host_info.id` | no | API: Specifies the id of the Recovery. |
| `mssql_params.recover_app_params.host_info.name` | no | API: Specifies the name of the Recovery. |
| `mssql_params.recover_app_params.sql_target_params` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.recover_to_new_source` | yes | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.data_file_directory_location` | yes | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.database_name` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.instance_name` | yes | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.keep_cdc` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.log_file_directory_location` | yes | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.native_log_recovery_with_clause` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.native_recovery_with_clause` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.overwriting_policy` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.replay_entire_last_log` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.restore_time_usecs` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.with_no_recovery` | no | — |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.host` (block) | yes | nested block |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.multi_stage_restore_options` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.new_source_config.secondary_data_files_dir_list` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.original_source_config` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.capture_tail_logs` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.data_file_directory_location` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.keep_cdc` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.log_file_directory_location` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.native_log_recovery_with_clause` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.native_recovery_with_clause` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.new_database_name` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.overwriting_policy` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.replay_entire_last_log` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.restore_time_usecs` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.with_no_recovery` | no | — |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.multi_stage_restore_options` (block) | no | nested block |
| `mssql_params.recover_app_params.sql_target_params.original_source_config.secondary_data_files_dir_list` (block) | no | nested block |
| `mssql_params.vlan_config` (block) | no | nested block |
| `mssql_params.vlan_config.disable_vlan` | no | — |
| `mssql_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `physical_params` (block) | no | nested block |
| `physical_params.recovery_action` | yes | — |
| `physical_params.download_file_and_folder_params` (block) | no | nested block |
| `physical_params.download_file_and_folder_params.download_file_path` | no | — |
| `physical_params.download_file_and_folder_params.expiry_time_usecs` | no | — |
| `physical_params.download_file_and_folder_params.files_and_folders` (block) | no | nested block |
| `physical_params.download_file_and_folder_params.files_and_folders.absolute_path` | yes | — |
| `physical_params.download_file_and_folder_params.files_and_folders.is_directory` | no | — |
| `physical_params.download_file_and_folder_params.files_and_folders.is_view_file_recovery` | no | — |
| `physical_params.mount_volume_params` (block) | no | nested block |
| `physical_params.mount_volume_params.target_environment` | yes | — |
| `physical_params.mount_volume_params.physical_target_params` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.mount_to_original_target` | yes | — |
| `physical_params.mount_volume_params.physical_target_params.read_only_mount` | no | — |
| `physical_params.mount_volume_params.physical_target_params.volume_names` | no | — |
| `physical_params.mount_volume_params.physical_target_params.new_target_config` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.new_target_config.mount_target` (block) | yes | nested block |
| `physical_params.mount_volume_params.physical_target_params.new_target_config.server_credentials` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.original_target_config` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.original_target_config.server_credentials` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.vlan_config` (block) | no | nested block |
| `physical_params.mount_volume_params.physical_target_params.vlan_config.disable_vlan` | no | — |
| `physical_params.mount_volume_params.physical_target_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `physical_params.objects` (block) | yes | nested block |
| `physical_params.objects.point_in_time_usecs` | no | — |
| `physical_params.objects.protection_group_id` | no | — |
| `physical_params.objects.protection_group_name` | no | — |
| `physical_params.objects.recover_from_standby` | no | — |
| `physical_params.objects.snapshot_id` | yes | — |
| `physical_params.recover_file_and_folder_params` (block) | no | nested block |
| `physical_params.recover_file_and_folder_params.target_environment` | yes | — |
| `physical_params.recover_file_and_folder_params.files_and_folders` (block) | yes | nested block |
| `physical_params.recover_file_and_folder_params.files_and_folders.absolute_path` | yes | — |
| `physical_params.recover_file_and_folder_params.files_and_folders.is_directory` | no | — |
| `physical_params.recover_file_and_folder_params.files_and_folders.is_view_file_recovery` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params` (block) | no | nested block |
| `physical_params.recover_file_and_folder_params.physical_target_params.alternate_restore_directory` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.continue_on_error` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.overwrite_existing` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.preserve_acls` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.preserve_attributes` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.preserve_timestamps` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.restore_entity_type` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.restore_to_original_paths` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.save_success_files` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.recover_target` (block) | yes | nested block |
| `physical_params.recover_file_and_folder_params.physical_target_params.recover_target.id` | yes | API: Specifies the id of the Recovery. |
| `physical_params.recover_file_and_folder_params.physical_target_params.vlan_config` (block) | no | nested block |
| `physical_params.recover_file_and_folder_params.physical_target_params.vlan_config.disable_vlan` | no | — |
| `physical_params.recover_file_and_folder_params.physical_target_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `physical_params.recover_volume_params` (block) | no | nested block |
| `physical_params.recover_volume_params.target_environment` | yes | — |
| `physical_params.recover_volume_params.physical_target_params` (block) | no | nested block |
| `physical_params.recover_volume_params.physical_target_params.force_unmount_volume` | no | — |
| `physical_params.recover_volume_params.physical_target_params.mount_target` (block) | yes | nested block |
| `physical_params.recover_volume_params.physical_target_params.mount_target.id` | yes | API: Specifies the id of the Recovery. |
| `physical_params.recover_volume_params.physical_target_params.vlan_config` (block) | no | nested block |
| `physical_params.recover_volume_params.physical_target_params.vlan_config.disable_vlan` | no | — |
| `physical_params.recover_volume_params.physical_target_params.vlan_config.id` | no | API: Specifies the id of the Recovery. |
| `physical_params.recover_volume_params.physical_target_params.volume_mapping` (block) | yes | nested block |
| `physical_params.recover_volume_params.physical_target_params.volume_mapping.destination_volume_guid` | yes | — |
| `physical_params.recover_volume_params.physical_target_params.volume_mapping.source_volume_guid` | yes | — |
| `physical_params.system_recovery_params` (block) | no | nested block |
| `physical_params.system_recovery_params.full_nas_path` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_invalid_name` | `name` |
| `ibm_backup_recovery_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_invalid_snapshot_environment` | `snapshot_environment` |
| `ibm_backup_recovery_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recovery_action` | `kubernetes_params.recovery_action` |
| `ibm_backup_recovery_invalid_kubernetes_params__download_file_and_folder_params__download_file_path` | `kubernetes_params.download_file_and_folder_params.download_file_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__download_file_and_folder_params__expiry_time_usecs` | `kubernetes_params.download_file_and_folder_params.expiry_time_usecs` |
| `ibm_backup_recovery_invalid_kubernetes_params__download_file_and_folder_params__files_and_folders__absolute_path` | `kubernetes_params.download_file_and_folder_params.files_and_folders.absolute_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__download_file_and_folder_params__files_and_folders__is_directory` | `kubernetes_params.download_file_and_folder_params.files_and_folders.is_directory` |
| `ibm_backup_recovery_invalid_kubernetes_params__download_file_and_folder_params__files_and_folders__is_view_file_recovery` | `kubernetes_params.download_file_and_folder_params.files_and_folders.is_view_file_recovery` |
| `ibm_backup_recovery_invalid_kubernetes_params__objects__point_in_time_usecs` | `kubernetes_params.objects.point_in_time_usecs` |
| `ibm_backup_recovery_invalid_kubernetes_params__objects__protection_group_id` | `kubernetes_params.objects.protection_group_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__objects__protection_group_name` | `kubernetes_params.objects.protection_group_name` |
| `ibm_backup_recovery_invalid_kubernetes_params__objects__recover_from_standby` | `kubernetes_params.objects.recover_from_standby` |
| `ibm_backup_recovery_invalid_kubernetes_params__objects__snapshot_id` | `kubernetes_params.objects.snapshot_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__target_environment` | `kubernetes_params.recover_file_and_folder_params.target_environment` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__files_and_folders__absolute_path` | `kubernetes_params.recover_file_and_folder_params.files_and_folders.absolute_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__files_and_folders__is_directory` | `kubernetes_params.recover_file_and_folder_params.files_and_folders.is_directory` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__files_and_folders__is_view_file_recovery` | `kubernetes_params.recover_file_and_folder_params.files_and_folders.is_view_file_recovery` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__continue_on_error` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.continue_on_error` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__overwrite_existing` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.overwrite_existing` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__preserve_attributes` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.preserve_attributes` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__recover_to_original_target` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.recover_to_original_target` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__new_target_config__absolute_path` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.new_target_config.absolute_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__original_target_config__alternate_path` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.original_target_config.alternate_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__original_target_config__recover_to_original_path` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.original_target_config.recover_to_original_path` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__vlan_config__disable_vlan` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_file_and_folder_params__kubernetes_target_params__vlan_config__id` | `kubernetes_params.recover_file_and_folder_params.kubernetes_target_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__target_environment` | `kubernetes_params.recover_namespace_params.target_environment` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__preserve_mac_address` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.preserve_mac_address` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_pvcs_only` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_pvcs_only` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__skip_cluster_compatibility_check` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.skip_cluster_compatibility_check` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__exclude_params__label_combination_method` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.label_combination_method` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__exclude_params__label_filter_entity_type` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.label_filter_entity_type` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__exclude_params__objects` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.exclude_params.objects` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__excluded_pvcs__metadata_only` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.excluded_pvcs.metadata_only` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__include_params__label_combination_method` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.label_combination_method` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__include_params__label_filter_entity_type` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.label_filter_entity_type` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__include_params__objects` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.include_params.objects` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__point_in_time_usecs` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.point_in_time_usecs` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__protection_group_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.protection_group_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__protection_group_name` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.protection_group_name` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__recover_from_standby` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.recover_from_standby` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__recover_pvcs_only` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.recover_pvcs_only` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__snapshot_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.snapshot_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__unbind_pvcs` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.unbind_pvcs` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__objects__use_instant_recovery` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.objects.use_instant_recovery` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_cluster_scoped_resources__snapshot_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_cluster_scoped_resources.snapshot_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_protection_group_runs_params__archival_target_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.archival_target_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_protection_group_runs_params__protection_group_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_protection_group_runs_params__protection_group_instance_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_instance_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recover_protection_group_runs_params__protection_group_run_id` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recover_protection_group_runs_params.protection_group_run_id` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recovery_region_migration_params__current_value` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_region_migration_params.current_value` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recovery_region_migration_params__new_value` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_region_migration_params.new_value` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recovery_target_config__recover_to_new_source` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_target_config.recover_to_new_source` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recovery_zone_migration_params__current_value` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_zone_migration_params.current_value` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__recovery_zone_migration_params__new_value` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.recovery_zone_migration_params.new_value` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__rename_recovered_namespaces_params__prefix` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.rename_recovered_namespaces_params.prefix` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__rename_recovered_namespaces_params__suffix` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.rename_recovered_namespaces_params.suffix` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__kubernetes_target_params__storage_class__use_storage_class_mapping` | `kubernetes_params.recover_namespace_params.kubernetes_target_params.storage_class.use_storage_class_mapping` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__vlan_config__disable_vlan` | `kubernetes_params.recover_namespace_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_kubernetes_params__recover_namespace_params__vlan_config__id` | `kubernetes_params.recover_namespace_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_mssql_params__recovery_action` | `mssql_params.recovery_action` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__is_encrypted` | `mssql_params.recover_app_params.is_encrypted` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__point_in_time_usecs` | `mssql_params.recover_app_params.point_in_time_usecs` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__protection_group_id` | `mssql_params.recover_app_params.protection_group_id` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__protection_group_name` | `mssql_params.recover_app_params.protection_group_name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__recover_from_standby` | `mssql_params.recover_app_params.recover_from_standby` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__snapshot_id` | `mssql_params.recover_app_params.snapshot_id` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__target_environment` | `mssql_params.recover_app_params.target_environment` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__aag_info__name` | `mssql_params.recover_app_params.aag_info.name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__aag_info__object_id` | `mssql_params.recover_app_params.aag_info.object_id` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__host_info__environment` | `mssql_params.recover_app_params.host_info.environment` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__host_info__id` | `mssql_params.recover_app_params.host_info.id` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__host_info__name` | `mssql_params.recover_app_params.host_info.name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__recover_to_new_source` | `mssql_params.recover_app_params.sql_target_params.recover_to_new_source` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__data_file_directory_location` | `mssql_params.recover_app_params.sql_target_params.new_source_config.data_file_directory_location` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__database_name` | `mssql_params.recover_app_params.sql_target_params.new_source_config.database_name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__instance_name` | `mssql_params.recover_app_params.sql_target_params.new_source_config.instance_name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__keep_cdc` | `mssql_params.recover_app_params.sql_target_params.new_source_config.keep_cdc` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__log_file_directory_location` | `mssql_params.recover_app_params.sql_target_params.new_source_config.log_file_directory_location` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__native_log_recovery_with_clause` | `mssql_params.recover_app_params.sql_target_params.new_source_config.native_log_recovery_with_clause` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__native_recovery_with_clause` | `mssql_params.recover_app_params.sql_target_params.new_source_config.native_recovery_with_clause` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__overwriting_policy` | `mssql_params.recover_app_params.sql_target_params.new_source_config.overwriting_policy` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__replay_entire_last_log` | `mssql_params.recover_app_params.sql_target_params.new_source_config.replay_entire_last_log` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__restore_time_usecs` | `mssql_params.recover_app_params.sql_target_params.new_source_config.restore_time_usecs` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__new_source_config__with_no_recovery` | `mssql_params.recover_app_params.sql_target_params.new_source_config.with_no_recovery` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__capture_tail_logs` | `mssql_params.recover_app_params.sql_target_params.original_source_config.capture_tail_logs` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__data_file_directory_location` | `mssql_params.recover_app_params.sql_target_params.original_source_config.data_file_directory_location` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__keep_cdc` | `mssql_params.recover_app_params.sql_target_params.original_source_config.keep_cdc` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__log_file_directory_location` | `mssql_params.recover_app_params.sql_target_params.original_source_config.log_file_directory_location` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__native_log_recovery_with_clause` | `mssql_params.recover_app_params.sql_target_params.original_source_config.native_log_recovery_with_clause` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__native_recovery_with_clause` | `mssql_params.recover_app_params.sql_target_params.original_source_config.native_recovery_with_clause` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__new_database_name` | `mssql_params.recover_app_params.sql_target_params.original_source_config.new_database_name` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__overwriting_policy` | `mssql_params.recover_app_params.sql_target_params.original_source_config.overwriting_policy` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__replay_entire_last_log` | `mssql_params.recover_app_params.sql_target_params.original_source_config.replay_entire_last_log` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__restore_time_usecs` | `mssql_params.recover_app_params.sql_target_params.original_source_config.restore_time_usecs` |
| `ibm_backup_recovery_invalid_mssql_params__recover_app_params__sql_target_params__original_source_config__with_no_recovery` | `mssql_params.recover_app_params.sql_target_params.original_source_config.with_no_recovery` |
| `ibm_backup_recovery_invalid_mssql_params__vlan_config__disable_vlan` | `mssql_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_mssql_params__vlan_config__id` | `mssql_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_physical_params__recovery_action` | `physical_params.recovery_action` |
| `ibm_backup_recovery_invalid_physical_params__download_file_and_folder_params__download_file_path` | `physical_params.download_file_and_folder_params.download_file_path` |
| `ibm_backup_recovery_invalid_physical_params__download_file_and_folder_params__expiry_time_usecs` | `physical_params.download_file_and_folder_params.expiry_time_usecs` |
| `ibm_backup_recovery_invalid_physical_params__download_file_and_folder_params__files_and_folders__absolute_path` | `physical_params.download_file_and_folder_params.files_and_folders.absolute_path` |
| `ibm_backup_recovery_invalid_physical_params__download_file_and_folder_params__files_and_folders__is_directory` | `physical_params.download_file_and_folder_params.files_and_folders.is_directory` |
| `ibm_backup_recovery_invalid_physical_params__download_file_and_folder_params__files_and_folders__is_view_file_recovery` | `physical_params.download_file_and_folder_params.files_and_folders.is_view_file_recovery` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__target_environment` | `physical_params.mount_volume_params.target_environment` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__physical_target_params__mount_to_original_target` | `physical_params.mount_volume_params.physical_target_params.mount_to_original_target` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__physical_target_params__read_only_mount` | `physical_params.mount_volume_params.physical_target_params.read_only_mount` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__physical_target_params__volume_names` | `physical_params.mount_volume_params.physical_target_params.volume_names` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__physical_target_params__vlan_config__disable_vlan` | `physical_params.mount_volume_params.physical_target_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_physical_params__mount_volume_params__physical_target_params__vlan_config__id` | `physical_params.mount_volume_params.physical_target_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_physical_params__objects__point_in_time_usecs` | `physical_params.objects.point_in_time_usecs` |
| `ibm_backup_recovery_invalid_physical_params__objects__protection_group_id` | `physical_params.objects.protection_group_id` |
| `ibm_backup_recovery_invalid_physical_params__objects__protection_group_name` | `physical_params.objects.protection_group_name` |
| `ibm_backup_recovery_invalid_physical_params__objects__recover_from_standby` | `physical_params.objects.recover_from_standby` |
| `ibm_backup_recovery_invalid_physical_params__objects__snapshot_id` | `physical_params.objects.snapshot_id` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__target_environment` | `physical_params.recover_file_and_folder_params.target_environment` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__files_and_folders__absolute_path` | `physical_params.recover_file_and_folder_params.files_and_folders.absolute_path` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__files_and_folders__is_directory` | `physical_params.recover_file_and_folder_params.files_and_folders.is_directory` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__files_and_folders__is_view_file_recovery` | `physical_params.recover_file_and_folder_params.files_and_folders.is_view_file_recovery` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__alternate_restore_directory` | `physical_params.recover_file_and_folder_params.physical_target_params.alternate_restore_directory` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__continue_on_error` | `physical_params.recover_file_and_folder_params.physical_target_params.continue_on_error` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__overwrite_existing` | `physical_params.recover_file_and_folder_params.physical_target_params.overwrite_existing` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__preserve_acls` | `physical_params.recover_file_and_folder_params.physical_target_params.preserve_acls` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__preserve_attributes` | `physical_params.recover_file_and_folder_params.physical_target_params.preserve_attributes` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__preserve_timestamps` | `physical_params.recover_file_and_folder_params.physical_target_params.preserve_timestamps` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__restore_entity_type` | `physical_params.recover_file_and_folder_params.physical_target_params.restore_entity_type` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__restore_to_original_paths` | `physical_params.recover_file_and_folder_params.physical_target_params.restore_to_original_paths` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__save_success_files` | `physical_params.recover_file_and_folder_params.physical_target_params.save_success_files` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__recover_target__id` | `physical_params.recover_file_and_folder_params.physical_target_params.recover_target.id` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__vlan_config__disable_vlan` | `physical_params.recover_file_and_folder_params.physical_target_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_physical_params__recover_file_and_folder_params__physical_target_params__vlan_config__id` | `physical_params.recover_file_and_folder_params.physical_target_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__target_environment` | `physical_params.recover_volume_params.target_environment` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__force_unmount_volume` | `physical_params.recover_volume_params.physical_target_params.force_unmount_volume` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__mount_target__id` | `physical_params.recover_volume_params.physical_target_params.mount_target.id` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__vlan_config__disable_vlan` | `physical_params.recover_volume_params.physical_target_params.vlan_config.disable_vlan` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__vlan_config__id` | `physical_params.recover_volume_params.physical_target_params.vlan_config.id` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__volume_mapping__destination_volume_guid` | `physical_params.recover_volume_params.physical_target_params.volume_mapping.destination_volume_guid` |
| `ibm_backup_recovery_invalid_physical_params__recover_volume_params__physical_target_params__volume_mapping__source_volume_guid` | `physical_params.recover_volume_params.physical_target_params.volume_mapping.source_volume_guid` |
| `ibm_backup_recovery_invalid_physical_params__system_recovery_params__full_nas_path` | `physical_params.system_recovery_params.full_nas_path` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery" "test" {
  name = "example"
  snapshot_environment = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
