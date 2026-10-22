# ibm_backup_recovery_protection_group

Manages ibm_backup_recovery_protection_group. Specifies the name of the Protection Group.

This rule validates `ibm_backup_recovery_protection_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_protection_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `abort_in_blackouts` | no | — |
| `cluster_id` | no | — |
| `delete_snapshots` | no | — |
| `description` | no | API: Description about the tenant. |
| `end_time_usecs` | no | — |
| `environment` | yes | one of: `kKubernetes`, `kPhysical`, `kSQL` |
| `group_id` | no | — |
| `invalid_entities` | no | — |
| `is_active` | no | — |
| `is_deleted` | no | — |
| `is_paused` | no | — |
| `is_protect_once` | no | — |
| `last_modified_timestamp_usecs` | no | — |
| `last_run` | no | — |
| `missing_entities` | no | — |
| `name` | yes | API: Specifies the name of the Recovery. |
| `num_protected_objects` | no | — |
| `pause_in_blackouts` | no | — |
| `permissions` | no | — |
| `policy_id` | yes | — |
| `priority` | no | one of: `kHigh`, `kLow`, `kMedium` |
| `qos_policy` | no | one of: `kBackupAll`, `kBackupHDD`, `kBackupSSD`, `kTestAndDevHigh` |
| `region_id` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `advanced_configs` (block) | no | nested block |
| `advanced_configs.key` | yes | API: Specifies the unique key for custom property. |
| `advanced_configs.value` | yes | API: Specifies the value for the above custom key. |
| `alert_policy` (block) | no | nested block |
| `alert_policy.backup_run_status` | yes | — |
| `alert_policy.raise_object_level_failure_alert` | no | — |
| `alert_policy.raise_object_level_failure_alert_after_each_attempt` | no | — |
| `alert_policy.raise_object_level_failure_alert_after_last_attempt` | no | — |
| `alert_policy.alert_targets` (block) | no | nested block |
| `alert_policy.alert_targets.email_address` | yes | — |
| `alert_policy.alert_targets.language` | no | one of: `en-us`, `ja-jp`, `zh-cn` |
| `alert_policy.alert_targets.recipient_type` | no | — |
| `kubernetes_params` (block) | no | nested block |
| `kubernetes_params.enable_indexing` | no | — |
| `kubernetes_params.exclude_label_ids` | no | — |
| `kubernetes_params.exclude_object_ids` | no | — |
| `kubernetes_params.label_ids` | no | — |
| `kubernetes_params.leverage_csi_snapshot` | no | — |
| `kubernetes_params.non_snapshot_backup` | no | — |
| `kubernetes_params.perform_source_side_deduplication` | no | — |
| `kubernetes_params.volume_backup_failure` | no | — |
| `kubernetes_params.exclude_params` (block) | no | nested block |
| `kubernetes_params.exclude_params.label_combination_method` | no | — |
| `kubernetes_params.exclude_params.label_filter_entity_type` | no | — |
| `kubernetes_params.exclude_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.exclude_params.label_vector` (block) | no | nested block |
| `kubernetes_params.exclude_params.label_vector.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.exclude_params.label_vector.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.exclude_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.exclude_params.selected_resources.api_group` | no | — |
| `kubernetes_params.exclude_params.selected_resources.is_cluster_scoped` | no | — |
| `kubernetes_params.exclude_params.selected_resources.kind` | no | API: The kind of the resource type. (Eg. VirtualMachine) |
| `kubernetes_params.exclude_params.selected_resources.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.exclude_params.selected_resources.version` | no | API: The version under the API group for the resource. This field is ignored for resource selection during recovery. |
| `kubernetes_params.exclude_params.selected_resources.resource_list` (block) | no | nested block |
| `kubernetes_params.exclude_params.selected_resources.resource_list.entity_id` | no | — |
| `kubernetes_params.exclude_params.selected_resources.resource_list.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.include_params` (block) | no | nested block |
| `kubernetes_params.include_params.label_combination_method` | no | — |
| `kubernetes_params.include_params.label_filter_entity_type` | no | — |
| `kubernetes_params.include_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.include_params.label_vector` (block) | no | nested block |
| `kubernetes_params.include_params.label_vector.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.include_params.label_vector.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.include_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.include_params.selected_resources.api_group` | no | — |
| `kubernetes_params.include_params.selected_resources.is_cluster_scoped` | no | — |
| `kubernetes_params.include_params.selected_resources.kind` | no | API: The kind of the resource type. (Eg. VirtualMachine) |
| `kubernetes_params.include_params.selected_resources.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.include_params.selected_resources.version` | no | API: The version under the API group for the resource. This field is ignored for resource selection during recovery. |
| `kubernetes_params.include_params.selected_resources.resource_list` (block) | no | nested block |
| `kubernetes_params.include_params.selected_resources.resource_list.entity_id` | no | — |
| `kubernetes_params.include_params.selected_resources.resource_list.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.objects` (block) | no | nested block |
| `kubernetes_params.objects.backup_only_pvc` | no | — |
| `kubernetes_params.objects.excluded_resources` | no | — |
| `kubernetes_params.objects.fail_backup_on_hook_failure` | no | — |
| `kubernetes_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `kubernetes_params.objects.included_resources` | no | — |
| `kubernetes_params.objects.exclude_params` (block) | no | nested block |
| `kubernetes_params.objects.exclude_params.label_combination_method` | no | — |
| `kubernetes_params.objects.exclude_params.label_filter_entity_type` | no | — |
| `kubernetes_params.objects.exclude_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.objects.exclude_params.label_vector` (block) | no | nested block |
| `kubernetes_params.objects.exclude_params.label_vector.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.objects.exclude_params.label_vector.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.objects.exclude_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.objects.exclude_params.selected_resources.api_group` | no | — |
| `kubernetes_params.objects.exclude_params.selected_resources.is_cluster_scoped` | no | — |
| `kubernetes_params.objects.exclude_params.selected_resources.kind` | no | API: The kind of the resource type. (Eg. VirtualMachine) |
| `kubernetes_params.objects.exclude_params.selected_resources.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.objects.exclude_params.selected_resources.version` | no | API: The version under the API group for the resource. This field is ignored for resource selection during recovery. |
| `kubernetes_params.objects.exclude_params.selected_resources.resource_list` (block) | no | nested block |
| `kubernetes_params.objects.exclude_pvcs` (block) | no | nested block |
| `kubernetes_params.objects.exclude_pvcs.metadata_only` | no | — |
| `kubernetes_params.objects.include_params` (block) | no | nested block |
| `kubernetes_params.objects.include_params.label_combination_method` | no | — |
| `kubernetes_params.objects.include_params.label_filter_entity_type` | no | — |
| `kubernetes_params.objects.include_params.objects` | no | API: Specifies the common snapshot parameters for a protected object. |
| `kubernetes_params.objects.include_params.label_vector` (block) | no | nested block |
| `kubernetes_params.objects.include_params.label_vector.key` | no | API: Specifies the unique key for custom property. |
| `kubernetes_params.objects.include_params.label_vector.value` | no | API: Specifies the value for the above custom key. |
| `kubernetes_params.objects.include_params.selected_resources` (block) | no | nested block |
| `kubernetes_params.objects.include_params.selected_resources.api_group` | no | — |
| `kubernetes_params.objects.include_params.selected_resources.is_cluster_scoped` | no | — |
| `kubernetes_params.objects.include_params.selected_resources.kind` | no | API: The kind of the resource type. (Eg. VirtualMachine) |
| `kubernetes_params.objects.include_params.selected_resources.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.objects.include_params.selected_resources.version` | no | API: The version under the API group for the resource. This field is ignored for resource selection during recovery. |
| `kubernetes_params.objects.include_params.selected_resources.resource_list` (block) | no | nested block |
| `kubernetes_params.objects.include_pvcs` (block) | no | nested block |
| `kubernetes_params.objects.include_pvcs.id` | no | API: Specifies the id of the Recovery. |
| `kubernetes_params.objects.include_pvcs.metadata_only` | no | — |
| `kubernetes_params.objects.include_pvcs.name` | no | API: Specifies the name of the Recovery. |
| `kubernetes_params.objects.quiesce_groups` (block) | no | nested block |
| `kubernetes_params.objects.quiesce_groups.quiesce_mode` | yes | — |
| `kubernetes_params.objects.quiesce_groups.quiesce_rules` (block) | yes | nested block |
| `kubernetes_params.objects.quiesce_groups.quiesce_rules.pod_selector_labels` (block) | no | nested block |
| `kubernetes_params.objects.quiesce_groups.quiesce_rules.post_snapshot_hooks` (block) | yes | nested block |
| `kubernetes_params.objects.quiesce_groups.quiesce_rules.pre_snapshot_hooks` (block) | yes | nested block |
| `kubernetes_params.vlan_params` (block) | no | nested block |
| `kubernetes_params.vlan_params.disable_vlan` | no | — |
| `kubernetes_params.vlan_params.interface_name` | no | — |
| `kubernetes_params.vlan_params.vlan_id` | no | — |
| `mssql_params` (block) | no | nested block |
| `mssql_params.protection_type` | yes | — |
| `mssql_params.file_protection_type_params` (block) | no | nested block |
| `mssql_params.file_protection_type_params.aag_backup_preference_type` | no | — |
| `mssql_params.file_protection_type_params.backup_system_dbs` | no | — |
| `mssql_params.file_protection_type_params.full_backups_copy_only` | no | — |
| `mssql_params.file_protection_type_params.log_backup_num_streams` | no | — |
| `mssql_params.file_protection_type_params.log_backup_with_clause` | no | — |
| `mssql_params.file_protection_type_params.perform_source_side_deduplication` | no | — |
| `mssql_params.file_protection_type_params.use_aag_preferences_from_server` | no | — |
| `mssql_params.file_protection_type_params.user_db_backup_preference_type` | no | — |
| `mssql_params.file_protection_type_params.additional_host_params` (block) | no | nested block |
| `mssql_params.file_protection_type_params.additional_host_params.disable_source_side_deduplication` | no | — |
| `mssql_params.file_protection_type_params.additional_host_params.host_id` | yes | — |
| `mssql_params.file_protection_type_params.advanced_settings` (block) | no | nested block |
| `mssql_params.file_protection_type_params.advanced_settings.cloned_db_backup_status` | no | — |
| `mssql_params.file_protection_type_params.advanced_settings.db_backup_if_not_online_status` | no | — |
| `mssql_params.file_protection_type_params.advanced_settings.missing_db_backup_status` | no | — |
| `mssql_params.file_protection_type_params.advanced_settings.offline_restoring_db_backup_status` | no | — |
| `mssql_params.file_protection_type_params.advanced_settings.read_only_db_backup_status` | no | — |
| `mssql_params.file_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` | no | — |
| `mssql_params.file_protection_type_params.exclude_filters` (block) | no | nested block |
| `mssql_params.file_protection_type_params.exclude_filters.filter_string` | no | — |
| `mssql_params.file_protection_type_params.exclude_filters.is_regular_expression` | no | — |
| `mssql_params.file_protection_type_params.objects` (block) | yes | nested block |
| `mssql_params.file_protection_type_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `mssql_params.file_protection_type_params.pre_post_script` (block) | no | nested block |
| `mssql_params.file_protection_type_params.pre_post_script.post_script` (block) | no | nested block |
| `mssql_params.file_protection_type_params.pre_post_script.post_script.is_active` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.post_script.params` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.post_script.path` | yes | — |
| `mssql_params.file_protection_type_params.pre_post_script.post_script.timeout_secs` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script` (block) | no | nested block |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script.continue_on_error` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script.is_active` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script.params` | no | — |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script.path` | yes | — |
| `mssql_params.file_protection_type_params.pre_post_script.pre_script.timeout_secs` | no | — |
| `mssql_params.native_protection_type_params` (block) | no | nested block |
| `mssql_params.native_protection_type_params.aag_backup_preference_type` | no | — |
| `mssql_params.native_protection_type_params.backup_system_dbs` | no | — |
| `mssql_params.native_protection_type_params.full_backups_copy_only` | no | — |
| `mssql_params.native_protection_type_params.log_backup_num_streams` | no | — |
| `mssql_params.native_protection_type_params.log_backup_with_clause` | no | — |
| `mssql_params.native_protection_type_params.num_streams` | no | — |
| `mssql_params.native_protection_type_params.use_aag_preferences_from_server` | no | — |
| `mssql_params.native_protection_type_params.user_db_backup_preference_type` | no | — |
| `mssql_params.native_protection_type_params.with_clause` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings` (block) | no | nested block |
| `mssql_params.native_protection_type_params.advanced_settings.cloned_db_backup_status` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings.db_backup_if_not_online_status` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings.missing_db_backup_status` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings.offline_restoring_db_backup_status` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings.read_only_db_backup_status` | no | — |
| `mssql_params.native_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` | no | — |
| `mssql_params.native_protection_type_params.exclude_filters` (block) | no | nested block |
| `mssql_params.native_protection_type_params.exclude_filters.filter_string` | no | — |
| `mssql_params.native_protection_type_params.exclude_filters.is_regular_expression` | no | — |
| `mssql_params.native_protection_type_params.objects` (block) | yes | nested block |
| `mssql_params.native_protection_type_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `mssql_params.native_protection_type_params.pre_post_script` (block) | no | nested block |
| `mssql_params.native_protection_type_params.pre_post_script.post_script` (block) | no | nested block |
| `mssql_params.native_protection_type_params.pre_post_script.post_script.is_active` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.post_script.params` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.post_script.path` | yes | — |
| `mssql_params.native_protection_type_params.pre_post_script.post_script.timeout_secs` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script` (block) | no | nested block |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script.continue_on_error` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script.is_active` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script.params` | no | — |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script.path` | yes | — |
| `mssql_params.native_protection_type_params.pre_post_script.pre_script.timeout_secs` | no | — |
| `mssql_params.volume_protection_type_params` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.aag_backup_preference_type` | no | — |
| `mssql_params.volume_protection_type_params.backup_db_volumes_only` | no | — |
| `mssql_params.volume_protection_type_params.backup_system_dbs` | no | — |
| `mssql_params.volume_protection_type_params.full_backups_copy_only` | no | — |
| `mssql_params.volume_protection_type_params.incremental_backup_after_restart` | no | — |
| `mssql_params.volume_protection_type_params.log_backup_num_streams` | no | — |
| `mssql_params.volume_protection_type_params.log_backup_with_clause` | no | — |
| `mssql_params.volume_protection_type_params.use_aag_preferences_from_server` | no | — |
| `mssql_params.volume_protection_type_params.user_db_backup_preference_type` | no | — |
| `mssql_params.volume_protection_type_params.additional_host_params` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.additional_host_params.enable_system_backup` | no | — |
| `mssql_params.volume_protection_type_params.additional_host_params.host_id` | yes | — |
| `mssql_params.volume_protection_type_params.additional_host_params.volume_guids` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.advanced_settings.cloned_db_backup_status` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings.db_backup_if_not_online_status` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings.missing_db_backup_status` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings.offline_restoring_db_backup_status` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings.read_only_db_backup_status` | no | — |
| `mssql_params.volume_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` | no | — |
| `mssql_params.volume_protection_type_params.exclude_filters` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.exclude_filters.filter_string` | no | — |
| `mssql_params.volume_protection_type_params.exclude_filters.is_regular_expression` | no | — |
| `mssql_params.volume_protection_type_params.indexing_policy` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.indexing_policy.enable_indexing` | yes | — |
| `mssql_params.volume_protection_type_params.indexing_policy.exclude_paths` | no | — |
| `mssql_params.volume_protection_type_params.indexing_policy.include_paths` | no | — |
| `mssql_params.volume_protection_type_params.objects` (block) | yes | nested block |
| `mssql_params.volume_protection_type_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `mssql_params.volume_protection_type_params.pre_post_script` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.pre_post_script.post_script` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.pre_post_script.post_script.is_active` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.post_script.params` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.post_script.path` | yes | — |
| `mssql_params.volume_protection_type_params.pre_post_script.post_script.timeout_secs` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script` (block) | no | nested block |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script.continue_on_error` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script.is_active` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script.params` | no | — |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script.path` | yes | — |
| `mssql_params.volume_protection_type_params.pre_post_script.pre_script.timeout_secs` | no | — |
| `physical_params` (block) | no | nested block |
| `physical_params.protection_type` | yes | — |
| `physical_params.file_protection_type_params` (block) | no | nested block |
| `physical_params.file_protection_type_params.allow_parallel_runs` | no | — |
| `physical_params.file_protection_type_params.cobmr_backup` | no | — |
| `physical_params.file_protection_type_params.continue_on_quiesce_failure` | no | — |
| `physical_params.file_protection_type_params.dedup_exclusion_source_ids` | no | — |
| `physical_params.file_protection_type_params.excluded_vss_writers` | no | — |
| `physical_params.file_protection_type_params.global_exclude_fs` | no | — |
| `physical_params.file_protection_type_params.global_exclude_paths` | no | — |
| `physical_params.file_protection_type_params.ignorable_errors` | no | — |
| `physical_params.file_protection_type_params.perform_brick_based_deduplication` | no | — |
| `physical_params.file_protection_type_params.perform_source_side_deduplication` | no | — |
| `physical_params.file_protection_type_params.quiesce` | no | — |
| `physical_params.file_protection_type_params.indexing_policy` (block) | no | nested block |
| `physical_params.file_protection_type_params.indexing_policy.enable_indexing` | yes | — |
| `physical_params.file_protection_type_params.indexing_policy.exclude_paths` | no | — |
| `physical_params.file_protection_type_params.indexing_policy.include_paths` | no | — |
| `physical_params.file_protection_type_params.objects` (block) | yes | nested block |
| `physical_params.file_protection_type_params.objects.excluded_vss_writers` | no | — |
| `physical_params.file_protection_type_params.objects.follow_nas_symlink_target` | no | — |
| `physical_params.file_protection_type_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `physical_params.file_protection_type_params.objects.metadata_file_path` | no | — |
| `physical_params.file_protection_type_params.objects.nested_volume_types_to_skip` | no | — |
| `physical_params.file_protection_type_params.objects.uses_path_level_skip_nested_volume_setting` | no | — |
| `physical_params.file_protection_type_params.objects.file_paths` (block) | no | nested block |
| `physical_params.file_protection_type_params.objects.file_paths.excluded_paths` | no | — |
| `physical_params.file_protection_type_params.objects.file_paths.included_path` | yes | — |
| `physical_params.file_protection_type_params.objects.file_paths.skip_nested_volumes` | no | — |
| `physical_params.file_protection_type_params.pre_post_script` (block) | no | nested block |
| `physical_params.file_protection_type_params.pre_post_script.post_script` (block) | no | nested block |
| `physical_params.file_protection_type_params.pre_post_script.post_script.is_active` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.post_script.params` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.post_script.path` | yes | — |
| `physical_params.file_protection_type_params.pre_post_script.post_script.timeout_secs` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.pre_script` (block) | no | nested block |
| `physical_params.file_protection_type_params.pre_post_script.pre_script.continue_on_error` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.pre_script.is_active` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.pre_script.params` | no | — |
| `physical_params.file_protection_type_params.pre_post_script.pre_script.path` | yes | — |
| `physical_params.file_protection_type_params.pre_post_script.pre_script.timeout_secs` | no | — |
| `physical_params.file_protection_type_params.task_timeouts` (block) | no | nested block |
| `physical_params.file_protection_type_params.task_timeouts.backup_type` | no | — |
| `physical_params.file_protection_type_params.task_timeouts.timeout_mins` | no | — |
| `physical_params.volume_protection_type_params` (block) | no | nested block |
| `physical_params.volume_protection_type_params.cobmr_backup` | no | — |
| `physical_params.volume_protection_type_params.continue_on_quiesce_failure` | no | — |
| `physical_params.volume_protection_type_params.dedup_exclusion_source_ids` | no | — |
| `physical_params.volume_protection_type_params.excluded_vss_writers` | no | — |
| `physical_params.volume_protection_type_params.incremental_backup_after_restart` | no | — |
| `physical_params.volume_protection_type_params.perform_source_side_deduplication` | no | — |
| `physical_params.volume_protection_type_params.quiesce` | no | — |
| `physical_params.volume_protection_type_params.indexing_policy` (block) | no | nested block |
| `physical_params.volume_protection_type_params.indexing_policy.enable_indexing` | yes | — |
| `physical_params.volume_protection_type_params.indexing_policy.exclude_paths` | no | — |
| `physical_params.volume_protection_type_params.indexing_policy.include_paths` | no | — |
| `physical_params.volume_protection_type_params.objects` (block) | yes | nested block |
| `physical_params.volume_protection_type_params.objects.enable_system_backup` | no | — |
| `physical_params.volume_protection_type_params.objects.excluded_vss_writers` | no | — |
| `physical_params.volume_protection_type_params.objects.id` | yes | API: Specifies the id of the Recovery. |
| `physical_params.volume_protection_type_params.objects.name` | no | API: Specifies the name of the Recovery. |
| `physical_params.volume_protection_type_params.objects.volume_guids` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script` (block) | no | nested block |
| `physical_params.volume_protection_type_params.pre_post_script.post_script` (block) | no | nested block |
| `physical_params.volume_protection_type_params.pre_post_script.post_script.is_active` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.post_script.params` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.post_script.path` | yes | — |
| `physical_params.volume_protection_type_params.pre_post_script.post_script.timeout_secs` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script` (block) | no | nested block |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script.continue_on_error` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script.is_active` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script.params` | no | — |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script.path` | yes | — |
| `physical_params.volume_protection_type_params.pre_post_script.pre_script.timeout_secs` | no | — |
| `sla` (block) | no | nested block |
| `sla.backup_run_type` | no | — |
| `sla.sla_minutes` | no | — |
| `start_time` (block) | no | nested block |
| `start_time.hour` | yes | — |
| `start_time.minute` | yes | — |
| `start_time.time_zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_group_invalid_abort_in_blackouts` | `abort_in_blackouts` |
| `ibm_backup_recovery_protection_group_invalid_delete_snapshots` | `delete_snapshots` |
| `ibm_backup_recovery_protection_group_invalid_description` | `description` |
| `ibm_backup_recovery_protection_group_invalid_end_time_usecs` | `end_time_usecs` |
| `ibm_backup_recovery_protection_group_invalid_environment` | `environment` |
| `ibm_backup_recovery_protection_group_invalid_is_paused` | `is_paused` |
| `ibm_backup_recovery_protection_group_invalid_name` | `name` |
| `ibm_backup_recovery_protection_group_invalid_pause_in_blackouts` | `pause_in_blackouts` |
| `ibm_backup_recovery_protection_group_invalid_policy_id` | `policy_id` |
| `ibm_backup_recovery_protection_group_invalid_priority` | `priority` |
| `ibm_backup_recovery_protection_group_invalid_qos_policy` | `qos_policy` |
| `ibm_backup_recovery_protection_group_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_protection_group_invalid_advanced_configs__key` | `advanced_configs.key` |
| `ibm_backup_recovery_protection_group_invalid_advanced_configs__value` | `advanced_configs.value` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__backup_run_status` | `alert_policy.backup_run_status` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__raise_object_level_failure_alert` | `alert_policy.raise_object_level_failure_alert` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__raise_object_level_failure_alert_after_each_attempt` | `alert_policy.raise_object_level_failure_alert_after_each_attempt` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__raise_object_level_failure_alert_after_last_attempt` | `alert_policy.raise_object_level_failure_alert_after_last_attempt` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__alert_targets__email_address` | `alert_policy.alert_targets.email_address` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__alert_targets__language` | `alert_policy.alert_targets.language` |
| `ibm_backup_recovery_protection_group_invalid_alert_policy__alert_targets__recipient_type` | `alert_policy.alert_targets.recipient_type` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__enable_indexing` | `kubernetes_params.enable_indexing` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_label_ids` | `kubernetes_params.exclude_label_ids` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_object_ids` | `kubernetes_params.exclude_object_ids` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__label_ids` | `kubernetes_params.label_ids` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__leverage_csi_snapshot` | `kubernetes_params.leverage_csi_snapshot` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__non_snapshot_backup` | `kubernetes_params.non_snapshot_backup` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__perform_source_side_deduplication` | `kubernetes_params.perform_source_side_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__volume_backup_failure` | `kubernetes_params.volume_backup_failure` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__label_combination_method` | `kubernetes_params.exclude_params.label_combination_method` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__label_filter_entity_type` | `kubernetes_params.exclude_params.label_filter_entity_type` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__objects` | `kubernetes_params.exclude_params.objects` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__label_vector__key` | `kubernetes_params.exclude_params.label_vector.key` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__label_vector__value` | `kubernetes_params.exclude_params.label_vector.value` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__api_group` | `kubernetes_params.exclude_params.selected_resources.api_group` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__is_cluster_scoped` | `kubernetes_params.exclude_params.selected_resources.is_cluster_scoped` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__kind` | `kubernetes_params.exclude_params.selected_resources.kind` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__name` | `kubernetes_params.exclude_params.selected_resources.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__version` | `kubernetes_params.exclude_params.selected_resources.version` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__resource_list__entity_id` | `kubernetes_params.exclude_params.selected_resources.resource_list.entity_id` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__exclude_params__selected_resources__resource_list__name` | `kubernetes_params.exclude_params.selected_resources.resource_list.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__label_combination_method` | `kubernetes_params.include_params.label_combination_method` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__label_filter_entity_type` | `kubernetes_params.include_params.label_filter_entity_type` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__objects` | `kubernetes_params.include_params.objects` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__label_vector__key` | `kubernetes_params.include_params.label_vector.key` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__label_vector__value` | `kubernetes_params.include_params.label_vector.value` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__api_group` | `kubernetes_params.include_params.selected_resources.api_group` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__is_cluster_scoped` | `kubernetes_params.include_params.selected_resources.is_cluster_scoped` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__kind` | `kubernetes_params.include_params.selected_resources.kind` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__name` | `kubernetes_params.include_params.selected_resources.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__version` | `kubernetes_params.include_params.selected_resources.version` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__resource_list__entity_id` | `kubernetes_params.include_params.selected_resources.resource_list.entity_id` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__include_params__selected_resources__resource_list__name` | `kubernetes_params.include_params.selected_resources.resource_list.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__backup_only_pvc` | `kubernetes_params.objects.backup_only_pvc` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__excluded_resources` | `kubernetes_params.objects.excluded_resources` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__fail_backup_on_hook_failure` | `kubernetes_params.objects.fail_backup_on_hook_failure` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__id` | `kubernetes_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__included_resources` | `kubernetes_params.objects.included_resources` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__label_combination_method` | `kubernetes_params.objects.exclude_params.label_combination_method` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__label_filter_entity_type` | `kubernetes_params.objects.exclude_params.label_filter_entity_type` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__objects` | `kubernetes_params.objects.exclude_params.objects` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__label_vector__key` | `kubernetes_params.objects.exclude_params.label_vector.key` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__label_vector__value` | `kubernetes_params.objects.exclude_params.label_vector.value` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__selected_resources__api_group` | `kubernetes_params.objects.exclude_params.selected_resources.api_group` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__selected_resources__is_cluster_scoped` | `kubernetes_params.objects.exclude_params.selected_resources.is_cluster_scoped` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__selected_resources__kind` | `kubernetes_params.objects.exclude_params.selected_resources.kind` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__selected_resources__name` | `kubernetes_params.objects.exclude_params.selected_resources.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_params__selected_resources__version` | `kubernetes_params.objects.exclude_params.selected_resources.version` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__exclude_pvcs__metadata_only` | `kubernetes_params.objects.exclude_pvcs.metadata_only` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__label_combination_method` | `kubernetes_params.objects.include_params.label_combination_method` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__label_filter_entity_type` | `kubernetes_params.objects.include_params.label_filter_entity_type` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__objects` | `kubernetes_params.objects.include_params.objects` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__label_vector__key` | `kubernetes_params.objects.include_params.label_vector.key` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__label_vector__value` | `kubernetes_params.objects.include_params.label_vector.value` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__selected_resources__api_group` | `kubernetes_params.objects.include_params.selected_resources.api_group` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__selected_resources__is_cluster_scoped` | `kubernetes_params.objects.include_params.selected_resources.is_cluster_scoped` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__selected_resources__kind` | `kubernetes_params.objects.include_params.selected_resources.kind` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__selected_resources__name` | `kubernetes_params.objects.include_params.selected_resources.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_params__selected_resources__version` | `kubernetes_params.objects.include_params.selected_resources.version` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_pvcs__id` | `kubernetes_params.objects.include_pvcs.id` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_pvcs__metadata_only` | `kubernetes_params.objects.include_pvcs.metadata_only` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__include_pvcs__name` | `kubernetes_params.objects.include_pvcs.name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__objects__quiesce_groups__quiesce_mode` | `kubernetes_params.objects.quiesce_groups.quiesce_mode` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__vlan_params__disable_vlan` | `kubernetes_params.vlan_params.disable_vlan` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__vlan_params__interface_name` | `kubernetes_params.vlan_params.interface_name` |
| `ibm_backup_recovery_protection_group_invalid_kubernetes_params__vlan_params__vlan_id` | `kubernetes_params.vlan_params.vlan_id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__protection_type` | `mssql_params.protection_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__aag_backup_preference_type` | `mssql_params.file_protection_type_params.aag_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__backup_system_dbs` | `mssql_params.file_protection_type_params.backup_system_dbs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__full_backups_copy_only` | `mssql_params.file_protection_type_params.full_backups_copy_only` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__log_backup_num_streams` | `mssql_params.file_protection_type_params.log_backup_num_streams` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__log_backup_with_clause` | `mssql_params.file_protection_type_params.log_backup_with_clause` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__perform_source_side_deduplication` | `mssql_params.file_protection_type_params.perform_source_side_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__use_aag_preferences_from_server` | `mssql_params.file_protection_type_params.use_aag_preferences_from_server` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__user_db_backup_preference_type` | `mssql_params.file_protection_type_params.user_db_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__additional_host_params__disable_source_side_deduplication` | `mssql_params.file_protection_type_params.additional_host_params.disable_source_side_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__additional_host_params__host_id` | `mssql_params.file_protection_type_params.additional_host_params.host_id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__cloned_db_backup_status` | `mssql_params.file_protection_type_params.advanced_settings.cloned_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__db_backup_if_not_online_status` | `mssql_params.file_protection_type_params.advanced_settings.db_backup_if_not_online_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__missing_db_backup_status` | `mssql_params.file_protection_type_params.advanced_settings.missing_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__offline_restoring_db_backup_status` | `mssql_params.file_protection_type_params.advanced_settings.offline_restoring_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__read_only_db_backup_status` | `mssql_params.file_protection_type_params.advanced_settings.read_only_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__advanced_settings__report_all_non_autoprotect_db_errors` | `mssql_params.file_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__exclude_filters__filter_string` | `mssql_params.file_protection_type_params.exclude_filters.filter_string` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__exclude_filters__is_regular_expression` | `mssql_params.file_protection_type_params.exclude_filters.is_regular_expression` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__objects__id` | `mssql_params.file_protection_type_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__post_script__is_active` | `mssql_params.file_protection_type_params.pre_post_script.post_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__post_script__params` | `mssql_params.file_protection_type_params.pre_post_script.post_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__post_script__path` | `mssql_params.file_protection_type_params.pre_post_script.post_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__post_script__timeout_secs` | `mssql_params.file_protection_type_params.pre_post_script.post_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__pre_script__continue_on_error` | `mssql_params.file_protection_type_params.pre_post_script.pre_script.continue_on_error` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__pre_script__is_active` | `mssql_params.file_protection_type_params.pre_post_script.pre_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__pre_script__params` | `mssql_params.file_protection_type_params.pre_post_script.pre_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__pre_script__path` | `mssql_params.file_protection_type_params.pre_post_script.pre_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__file_protection_type_params__pre_post_script__pre_script__timeout_secs` | `mssql_params.file_protection_type_params.pre_post_script.pre_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__aag_backup_preference_type` | `mssql_params.native_protection_type_params.aag_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__backup_system_dbs` | `mssql_params.native_protection_type_params.backup_system_dbs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__full_backups_copy_only` | `mssql_params.native_protection_type_params.full_backups_copy_only` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__log_backup_num_streams` | `mssql_params.native_protection_type_params.log_backup_num_streams` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__log_backup_with_clause` | `mssql_params.native_protection_type_params.log_backup_with_clause` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__num_streams` | `mssql_params.native_protection_type_params.num_streams` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__use_aag_preferences_from_server` | `mssql_params.native_protection_type_params.use_aag_preferences_from_server` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__user_db_backup_preference_type` | `mssql_params.native_protection_type_params.user_db_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__with_clause` | `mssql_params.native_protection_type_params.with_clause` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__cloned_db_backup_status` | `mssql_params.native_protection_type_params.advanced_settings.cloned_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__db_backup_if_not_online_status` | `mssql_params.native_protection_type_params.advanced_settings.db_backup_if_not_online_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__missing_db_backup_status` | `mssql_params.native_protection_type_params.advanced_settings.missing_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__offline_restoring_db_backup_status` | `mssql_params.native_protection_type_params.advanced_settings.offline_restoring_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__read_only_db_backup_status` | `mssql_params.native_protection_type_params.advanced_settings.read_only_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__advanced_settings__report_all_non_autoprotect_db_errors` | `mssql_params.native_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__exclude_filters__filter_string` | `mssql_params.native_protection_type_params.exclude_filters.filter_string` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__exclude_filters__is_regular_expression` | `mssql_params.native_protection_type_params.exclude_filters.is_regular_expression` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__objects__id` | `mssql_params.native_protection_type_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__post_script__is_active` | `mssql_params.native_protection_type_params.pre_post_script.post_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__post_script__params` | `mssql_params.native_protection_type_params.pre_post_script.post_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__post_script__path` | `mssql_params.native_protection_type_params.pre_post_script.post_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__post_script__timeout_secs` | `mssql_params.native_protection_type_params.pre_post_script.post_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__pre_script__continue_on_error` | `mssql_params.native_protection_type_params.pre_post_script.pre_script.continue_on_error` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__pre_script__is_active` | `mssql_params.native_protection_type_params.pre_post_script.pre_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__pre_script__params` | `mssql_params.native_protection_type_params.pre_post_script.pre_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__pre_script__path` | `mssql_params.native_protection_type_params.pre_post_script.pre_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__native_protection_type_params__pre_post_script__pre_script__timeout_secs` | `mssql_params.native_protection_type_params.pre_post_script.pre_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__aag_backup_preference_type` | `mssql_params.volume_protection_type_params.aag_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__backup_db_volumes_only` | `mssql_params.volume_protection_type_params.backup_db_volumes_only` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__backup_system_dbs` | `mssql_params.volume_protection_type_params.backup_system_dbs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__full_backups_copy_only` | `mssql_params.volume_protection_type_params.full_backups_copy_only` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__incremental_backup_after_restart` | `mssql_params.volume_protection_type_params.incremental_backup_after_restart` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__log_backup_num_streams` | `mssql_params.volume_protection_type_params.log_backup_num_streams` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__log_backup_with_clause` | `mssql_params.volume_protection_type_params.log_backup_with_clause` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__use_aag_preferences_from_server` | `mssql_params.volume_protection_type_params.use_aag_preferences_from_server` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__user_db_backup_preference_type` | `mssql_params.volume_protection_type_params.user_db_backup_preference_type` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__additional_host_params__enable_system_backup` | `mssql_params.volume_protection_type_params.additional_host_params.enable_system_backup` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__additional_host_params__host_id` | `mssql_params.volume_protection_type_params.additional_host_params.host_id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__additional_host_params__volume_guids` | `mssql_params.volume_protection_type_params.additional_host_params.volume_guids` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__cloned_db_backup_status` | `mssql_params.volume_protection_type_params.advanced_settings.cloned_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__db_backup_if_not_online_status` | `mssql_params.volume_protection_type_params.advanced_settings.db_backup_if_not_online_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__missing_db_backup_status` | `mssql_params.volume_protection_type_params.advanced_settings.missing_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__offline_restoring_db_backup_status` | `mssql_params.volume_protection_type_params.advanced_settings.offline_restoring_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__read_only_db_backup_status` | `mssql_params.volume_protection_type_params.advanced_settings.read_only_db_backup_status` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__advanced_settings__report_all_non_autoprotect_db_errors` | `mssql_params.volume_protection_type_params.advanced_settings.report_all_non_autoprotect_db_errors` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__exclude_filters__filter_string` | `mssql_params.volume_protection_type_params.exclude_filters.filter_string` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__exclude_filters__is_regular_expression` | `mssql_params.volume_protection_type_params.exclude_filters.is_regular_expression` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__indexing_policy__enable_indexing` | `mssql_params.volume_protection_type_params.indexing_policy.enable_indexing` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__indexing_policy__exclude_paths` | `mssql_params.volume_protection_type_params.indexing_policy.exclude_paths` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__indexing_policy__include_paths` | `mssql_params.volume_protection_type_params.indexing_policy.include_paths` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__objects__id` | `mssql_params.volume_protection_type_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__post_script__is_active` | `mssql_params.volume_protection_type_params.pre_post_script.post_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__post_script__params` | `mssql_params.volume_protection_type_params.pre_post_script.post_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__post_script__path` | `mssql_params.volume_protection_type_params.pre_post_script.post_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__post_script__timeout_secs` | `mssql_params.volume_protection_type_params.pre_post_script.post_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__pre_script__continue_on_error` | `mssql_params.volume_protection_type_params.pre_post_script.pre_script.continue_on_error` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__pre_script__is_active` | `mssql_params.volume_protection_type_params.pre_post_script.pre_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__pre_script__params` | `mssql_params.volume_protection_type_params.pre_post_script.pre_script.params` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__pre_script__path` | `mssql_params.volume_protection_type_params.pre_post_script.pre_script.path` |
| `ibm_backup_recovery_protection_group_invalid_mssql_params__volume_protection_type_params__pre_post_script__pre_script__timeout_secs` | `mssql_params.volume_protection_type_params.pre_post_script.pre_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__protection_type` | `physical_params.protection_type` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__allow_parallel_runs` | `physical_params.file_protection_type_params.allow_parallel_runs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__cobmr_backup` | `physical_params.file_protection_type_params.cobmr_backup` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__continue_on_quiesce_failure` | `physical_params.file_protection_type_params.continue_on_quiesce_failure` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__dedup_exclusion_source_ids` | `physical_params.file_protection_type_params.dedup_exclusion_source_ids` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__excluded_vss_writers` | `physical_params.file_protection_type_params.excluded_vss_writers` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__global_exclude_fs` | `physical_params.file_protection_type_params.global_exclude_fs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__global_exclude_paths` | `physical_params.file_protection_type_params.global_exclude_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__ignorable_errors` | `physical_params.file_protection_type_params.ignorable_errors` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__perform_brick_based_deduplication` | `physical_params.file_protection_type_params.perform_brick_based_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__perform_source_side_deduplication` | `physical_params.file_protection_type_params.perform_source_side_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__quiesce` | `physical_params.file_protection_type_params.quiesce` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__indexing_policy__enable_indexing` | `physical_params.file_protection_type_params.indexing_policy.enable_indexing` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__indexing_policy__exclude_paths` | `physical_params.file_protection_type_params.indexing_policy.exclude_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__indexing_policy__include_paths` | `physical_params.file_protection_type_params.indexing_policy.include_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__excluded_vss_writers` | `physical_params.file_protection_type_params.objects.excluded_vss_writers` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__follow_nas_symlink_target` | `physical_params.file_protection_type_params.objects.follow_nas_symlink_target` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__id` | `physical_params.file_protection_type_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__metadata_file_path` | `physical_params.file_protection_type_params.objects.metadata_file_path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__nested_volume_types_to_skip` | `physical_params.file_protection_type_params.objects.nested_volume_types_to_skip` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__uses_path_level_skip_nested_volume_setting` | `physical_params.file_protection_type_params.objects.uses_path_level_skip_nested_volume_setting` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__file_paths__excluded_paths` | `physical_params.file_protection_type_params.objects.file_paths.excluded_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__file_paths__included_path` | `physical_params.file_protection_type_params.objects.file_paths.included_path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__objects__file_paths__skip_nested_volumes` | `physical_params.file_protection_type_params.objects.file_paths.skip_nested_volumes` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__post_script__is_active` | `physical_params.file_protection_type_params.pre_post_script.post_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__post_script__params` | `physical_params.file_protection_type_params.pre_post_script.post_script.params` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__post_script__path` | `physical_params.file_protection_type_params.pre_post_script.post_script.path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__post_script__timeout_secs` | `physical_params.file_protection_type_params.pre_post_script.post_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__pre_script__continue_on_error` | `physical_params.file_protection_type_params.pre_post_script.pre_script.continue_on_error` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__pre_script__is_active` | `physical_params.file_protection_type_params.pre_post_script.pre_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__pre_script__params` | `physical_params.file_protection_type_params.pre_post_script.pre_script.params` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__pre_script__path` | `physical_params.file_protection_type_params.pre_post_script.pre_script.path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__pre_post_script__pre_script__timeout_secs` | `physical_params.file_protection_type_params.pre_post_script.pre_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__task_timeouts__backup_type` | `physical_params.file_protection_type_params.task_timeouts.backup_type` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__file_protection_type_params__task_timeouts__timeout_mins` | `physical_params.file_protection_type_params.task_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__cobmr_backup` | `physical_params.volume_protection_type_params.cobmr_backup` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__continue_on_quiesce_failure` | `physical_params.volume_protection_type_params.continue_on_quiesce_failure` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__dedup_exclusion_source_ids` | `physical_params.volume_protection_type_params.dedup_exclusion_source_ids` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__excluded_vss_writers` | `physical_params.volume_protection_type_params.excluded_vss_writers` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__incremental_backup_after_restart` | `physical_params.volume_protection_type_params.incremental_backup_after_restart` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__perform_source_side_deduplication` | `physical_params.volume_protection_type_params.perform_source_side_deduplication` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__quiesce` | `physical_params.volume_protection_type_params.quiesce` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__indexing_policy__enable_indexing` | `physical_params.volume_protection_type_params.indexing_policy.enable_indexing` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__indexing_policy__exclude_paths` | `physical_params.volume_protection_type_params.indexing_policy.exclude_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__indexing_policy__include_paths` | `physical_params.volume_protection_type_params.indexing_policy.include_paths` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__objects__enable_system_backup` | `physical_params.volume_protection_type_params.objects.enable_system_backup` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__objects__excluded_vss_writers` | `physical_params.volume_protection_type_params.objects.excluded_vss_writers` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__objects__id` | `physical_params.volume_protection_type_params.objects.id` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__objects__name` | `physical_params.volume_protection_type_params.objects.name` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__objects__volume_guids` | `physical_params.volume_protection_type_params.objects.volume_guids` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__post_script__is_active` | `physical_params.volume_protection_type_params.pre_post_script.post_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__post_script__params` | `physical_params.volume_protection_type_params.pre_post_script.post_script.params` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__post_script__path` | `physical_params.volume_protection_type_params.pre_post_script.post_script.path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__post_script__timeout_secs` | `physical_params.volume_protection_type_params.pre_post_script.post_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__pre_script__continue_on_error` | `physical_params.volume_protection_type_params.pre_post_script.pre_script.continue_on_error` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__pre_script__is_active` | `physical_params.volume_protection_type_params.pre_post_script.pre_script.is_active` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__pre_script__params` | `physical_params.volume_protection_type_params.pre_post_script.pre_script.params` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__pre_script__path` | `physical_params.volume_protection_type_params.pre_post_script.pre_script.path` |
| `ibm_backup_recovery_protection_group_invalid_physical_params__volume_protection_type_params__pre_post_script__pre_script__timeout_secs` | `physical_params.volume_protection_type_params.pre_post_script.pre_script.timeout_secs` |
| `ibm_backup_recovery_protection_group_invalid_sla__backup_run_type` | `sla.backup_run_type` |
| `ibm_backup_recovery_protection_group_invalid_sla__sla_minutes` | `sla.sla_minutes` |
| `ibm_backup_recovery_protection_group_invalid_start_time__hour` | `start_time.hour` |
| `ibm_backup_recovery_protection_group_invalid_start_time__minute` | `start_time.minute` |
| `ibm_backup_recovery_protection_group_invalid_start_time__time_zone` | `start_time.time_zone` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_protection_group" "test" {
  environment = "kKubernetes"
  name = "example"
  policy_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
