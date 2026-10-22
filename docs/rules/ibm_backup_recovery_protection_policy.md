# ibm_backup_recovery_protection_policy

Manages ibm_backup_recovery_protection_policy. Specifies the name of the Protection Policy.

This rule validates `ibm_backup_recovery_protection_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Protection Policy Request.` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_protection_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `data_lock` | no | — |
| `description` | no | API: Specifies the description of the Protection Policy. |
| `is_cbs_enabled` | no | — |
| `is_replicated` | no | — |
| `is_usable` | no | — |
| `last_modification_time_usecs` | no | — |
| `name` | yes | API: Specifies the name of the Protection Policy. |
| `num_protected_objects` | no | — |
| `num_protection_groups` | no | — |
| `policy_id` | no | — |
| `template_id` | no | — |
| `version` | no | API: Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases. |
| `x_ibm_tenant_id` | yes | — |
| `backup_policy` (block) | yes | nested block |
| `backup_policy.bmr` (block) | no | nested block |
| `backup_policy.bmr.retention` (block) | yes | nested block |
| `backup_policy.bmr.retention.duration` | yes | min 1 |
| `backup_policy.bmr.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.bmr.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.bmr.retention.data_lock_config.duration` | yes | min 1 |
| `backup_policy.bmr.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `backup_policy.bmr.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `backup_policy.bmr.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.bmr.schedule` (block) | yes | nested block |
| `backup_policy.bmr.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.bmr.schedule.day_schedule` (block) | no | nested block |
| `backup_policy.bmr.schedule.day_schedule.frequency` | yes | min 1 |
| `backup_policy.bmr.schedule.month_schedule` (block) | no | nested block |
| `backup_policy.bmr.schedule.month_schedule.day_of_month` | no | — |
| `backup_policy.bmr.schedule.month_schedule.day_of_week` | no | — |
| `backup_policy.bmr.schedule.month_schedule.week_of_month` | no | — |
| `backup_policy.bmr.schedule.week_schedule` (block) | no | nested block |
| `backup_policy.bmr.schedule.week_schedule.day_of_week` | yes | — |
| `backup_policy.bmr.schedule.year_schedule` (block) | no | nested block |
| `backup_policy.bmr.schedule.year_schedule.day_of_year` | yes | — |
| `backup_policy.cdp` (block) | no | nested block |
| `backup_policy.cdp.retention` (block) | yes | nested block |
| `backup_policy.cdp.retention.duration` | yes | min 1 |
| `backup_policy.cdp.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.cdp.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.cdp.retention.data_lock_config.duration` | yes | min 1 |
| `backup_policy.cdp.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `backup_policy.cdp.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `backup_policy.cdp.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.log` (block) | no | nested block |
| `backup_policy.log.retention` (block) | yes | nested block |
| `backup_policy.log.retention.duration` | yes | min 1 |
| `backup_policy.log.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.log.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.log.retention.data_lock_config.duration` | yes | min 1 |
| `backup_policy.log.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `backup_policy.log.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `backup_policy.log.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.log.schedule` (block) | yes | nested block |
| `backup_policy.log.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.log.schedule.hour_schedule` (block) | no | nested block |
| `backup_policy.log.schedule.hour_schedule.frequency` | yes | min 1 |
| `backup_policy.log.schedule.minute_schedule` (block) | no | nested block |
| `backup_policy.log.schedule.minute_schedule.frequency` | yes | min 1 |
| `backup_policy.regular` (block) | yes | nested block |
| `backup_policy.regular.full` (block) | no | nested block |
| `backup_policy.regular.full.schedule` (block) | no | nested block |
| `backup_policy.regular.full.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.regular.full.schedule.day_schedule` (block) | no | nested block |
| `backup_policy.regular.full.schedule.month_schedule` (block) | no | nested block |
| `backup_policy.regular.full.schedule.week_schedule` (block) | no | nested block |
| `backup_policy.regular.full.schedule.year_schedule` (block) | no | nested block |
| `backup_policy.regular.full_backups` (block) | no | nested block |
| `backup_policy.regular.full_backups.retention` (block) | yes | nested block |
| `backup_policy.regular.full_backups.retention.duration` | yes | min 1 |
| `backup_policy.regular.full_backups.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.regular.full_backups.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.regular.full_backups.schedule` (block) | yes | nested block |
| `backup_policy.regular.full_backups.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.regular.full_backups.schedule.day_schedule` (block) | no | nested block |
| `backup_policy.regular.full_backups.schedule.month_schedule` (block) | no | nested block |
| `backup_policy.regular.full_backups.schedule.week_schedule` (block) | no | nested block |
| `backup_policy.regular.full_backups.schedule.year_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule` (block) | yes | nested block |
| `backup_policy.regular.incremental.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.regular.incremental.schedule.day_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule.hour_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule.minute_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule.month_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule.week_schedule` (block) | no | nested block |
| `backup_policy.regular.incremental.schedule.year_schedule` (block) | no | nested block |
| `backup_policy.regular.primary_backup_target` (block) | no | nested block |
| `backup_policy.regular.primary_backup_target.target_type` | no | — |
| `backup_policy.regular.primary_backup_target.use_default_backup_target` | no | — |
| `backup_policy.regular.primary_backup_target.archival_target_settings` (block) | no | nested block |
| `backup_policy.regular.primary_backup_target.archival_target_settings.target_id` | yes | — |
| `backup_policy.regular.primary_backup_target.archival_target_settings.tier_settings` (block) | no | nested block |
| `backup_policy.regular.retention` (block) | no | nested block |
| `backup_policy.regular.retention.duration` | yes | min 1 |
| `backup_policy.regular.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.regular.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.regular.retention.data_lock_config.duration` | yes | min 1 |
| `backup_policy.regular.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `backup_policy.regular.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `backup_policy.regular.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.run_timeouts` (block) | no | nested block |
| `backup_policy.run_timeouts.backup_type` | no | — |
| `backup_policy.run_timeouts.timeout_mins` | no | — |
| `backup_policy.storage_array_snapshot` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.retention` (block) | yes | nested block |
| `backup_policy.storage_array_snapshot.retention.duration` | yes | min 1 |
| `backup_policy.storage_array_snapshot.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.storage_array_snapshot.retention.data_lock_config` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.retention.data_lock_config.duration` | yes | min 1 |
| `backup_policy.storage_array_snapshot.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `backup_policy.storage_array_snapshot.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `backup_policy.storage_array_snapshot.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.storage_array_snapshot.schedule` (block) | yes | nested block |
| `backup_policy.storage_array_snapshot.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `backup_policy.storage_array_snapshot.schedule.day_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.day_schedule.frequency` | yes | min 1 |
| `backup_policy.storage_array_snapshot.schedule.hour_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.hour_schedule.frequency` | yes | min 1 |
| `backup_policy.storage_array_snapshot.schedule.minute_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.minute_schedule.frequency` | yes | min 1 |
| `backup_policy.storage_array_snapshot.schedule.month_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.month_schedule.day_of_month` | no | — |
| `backup_policy.storage_array_snapshot.schedule.month_schedule.day_of_week` | no | — |
| `backup_policy.storage_array_snapshot.schedule.month_schedule.week_of_month` | no | — |
| `backup_policy.storage_array_snapshot.schedule.week_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.week_schedule.day_of_week` | yes | — |
| `backup_policy.storage_array_snapshot.schedule.year_schedule` (block) | no | nested block |
| `backup_policy.storage_array_snapshot.schedule.year_schedule.day_of_year` | yes | — |
| `blackout_window` (block) | no | nested block |
| `blackout_window.config_id` | no | — |
| `blackout_window.day` | yes | one of: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` |
| `blackout_window.end_time` (block) | yes | nested block |
| `blackout_window.end_time.hour` | yes | range 0–23 |
| `blackout_window.end_time.minute` | yes | range 0–59 |
| `blackout_window.end_time.time_zone` | no | zone format |
| `blackout_window.start_time` (block) | yes | nested block |
| `blackout_window.start_time.hour` | yes | range 0–23 |
| `blackout_window.start_time.minute` | yes | range 0–59 |
| `blackout_window.start_time.time_zone` | no | zone format |
| `cascaded_targets_config` (block) | no | nested block |
| `cascaded_targets_config.source_cluster_id` | yes | — |
| `cascaded_targets_config.remote_targets` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.archival_targets` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.backup_run_type` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.config_id` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.copy_on_run_success` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.target_id` | yes | — |
| `cascaded_targets_config.remote_targets.archival_targets.extended_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.extended_retention.config_id` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.extended_retention.run_type` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.extended_retention.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.extended_retention.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.log_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.log_retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.archival_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.archival_targets.log_retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.archival_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.archival_targets.retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.run_timeouts` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.run_timeouts.backup_type` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.run_timeouts.timeout_mins` | no | — |
| `cascaded_targets_config.remote_targets.archival_targets.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.schedule.frequency` | no | min 1 |
| `cascaded_targets_config.remote_targets.archival_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings.cloud_platform` | yes | — |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings.aws_tiering` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings.azure_tiering` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings.google_tiering` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.archival_targets.tier_settings.oracle_tiering` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.backup_run_type` | no | — |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.config_id` | no | — |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.copy_on_run_success` | no | — |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.run_timeouts` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.run_timeouts.backup_type` | no | — |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.run_timeouts.timeout_mins` | no | — |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.schedule.frequency` | no | min 1 |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.target` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.target.id` | no | API: Specifies the unique id of the cloud spin entity. |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.target.aws_params` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.cloud_spin_targets.target.azure_params` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.backup_run_type` | no | — |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.config_id` | no | — |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.copy_on_run_success` | no | — |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.params` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.params.id` | no | API: Specifies the unique id of the cloud spin entity. |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.run_timeouts` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.run_timeouts.backup_type` | no | — |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.run_timeouts.timeout_mins` | no | — |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.schedule.frequency` | no | min 1 |
| `cascaded_targets_config.remote_targets.onprem_deploy_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.replication_targets` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.backup_run_type` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.config_id` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.copy_on_run_success` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.target_type` | yes | — |
| `cascaded_targets_config.remote_targets.replication_targets.aws_target_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.aws_target_config.region` | yes | API: Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. |
| `cascaded_targets_config.remote_targets.replication_targets.aws_target_config.source_id` | yes | — |
| `cascaded_targets_config.remote_targets.replication_targets.azure_target_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.azure_target_config.resource_group` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.azure_target_config.source_id` | yes | — |
| `cascaded_targets_config.remote_targets.replication_targets.log_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.log_retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.replication_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.replication_targets.log_retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.remote_target_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.remote_target_config.cluster_id` | yes | — |
| `cascaded_targets_config.remote_targets.replication_targets.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.replication_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.replication_targets.retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.run_timeouts` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.run_timeouts.backup_type` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.run_timeouts.timeout_mins` | no | — |
| `cascaded_targets_config.remote_targets.replication_targets.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.replication_targets.schedule.frequency` | no | min 1 |
| `cascaded_targets_config.remote_targets.replication_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.rpaas_targets` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.backup_run_type` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.config_id` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.copy_on_run_success` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.target_id` | yes | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.target_type` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.log_retention` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.log_retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.rpaas_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.rpaas_targets.log_retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.retention` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.retention.duration` | yes | min 1 |
| `cascaded_targets_config.remote_targets.rpaas_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `cascaded_targets_config.remote_targets.rpaas_targets.retention.data_lock_config` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.run_timeouts` (block) | no | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.run_timeouts.backup_type` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.run_timeouts.timeout_mins` | no | — |
| `cascaded_targets_config.remote_targets.rpaas_targets.schedule` (block) | yes | nested block |
| `cascaded_targets_config.remote_targets.rpaas_targets.schedule.frequency` | no | min 1 |
| `cascaded_targets_config.remote_targets.rpaas_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `extended_retention` (block) | no | nested block |
| `extended_retention.config_id` | no | — |
| `extended_retention.run_type` | no | — |
| `extended_retention.retention` (block) | yes | nested block |
| `extended_retention.retention.duration` | yes | min 1 |
| `extended_retention.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `extended_retention.retention.data_lock_config` (block) | no | nested block |
| `extended_retention.retention.data_lock_config.duration` | yes | min 1 |
| `extended_retention.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `extended_retention.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `extended_retention.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `extended_retention.schedule` (block) | yes | nested block |
| `extended_retention.schedule.frequency` | no | min 1 |
| `extended_retention.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy` (block) | no | nested block |
| `remote_target_policy.archival_targets` (block) | no | nested block |
| `remote_target_policy.archival_targets.backup_run_type` | no | — |
| `remote_target_policy.archival_targets.config_id` | no | — |
| `remote_target_policy.archival_targets.copy_on_run_success` | no | — |
| `remote_target_policy.archival_targets.target_id` | yes | — |
| `remote_target_policy.archival_targets.extended_retention` (block) | no | nested block |
| `remote_target_policy.archival_targets.extended_retention.config_id` | no | — |
| `remote_target_policy.archival_targets.extended_retention.run_type` | no | — |
| `remote_target_policy.archival_targets.extended_retention.retention` (block) | yes | nested block |
| `remote_target_policy.archival_targets.extended_retention.retention.duration` | yes | min 1 |
| `remote_target_policy.archival_targets.extended_retention.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.extended_retention.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.archival_targets.extended_retention.schedule` (block) | yes | nested block |
| `remote_target_policy.archival_targets.extended_retention.schedule.frequency` | no | min 1 |
| `remote_target_policy.archival_targets.extended_retention.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.log_retention` (block) | no | nested block |
| `remote_target_policy.archival_targets.log_retention.duration` | yes | min 1 |
| `remote_target_policy.archival_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.log_retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.archival_targets.log_retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.archival_targets.log_retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.archival_targets.log_retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.archival_targets.log_retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.retention` (block) | yes | nested block |
| `remote_target_policy.archival_targets.retention.duration` | yes | min 1 |
| `remote_target_policy.archival_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.archival_targets.retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.archival_targets.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.archival_targets.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.archival_targets.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.run_timeouts` (block) | no | nested block |
| `remote_target_policy.archival_targets.run_timeouts.backup_type` | no | — |
| `remote_target_policy.archival_targets.run_timeouts.timeout_mins` | no | — |
| `remote_target_policy.archival_targets.schedule` (block) | yes | nested block |
| `remote_target_policy.archival_targets.schedule.frequency` | no | min 1 |
| `remote_target_policy.archival_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.archival_targets.tier_settings` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.cloud_platform` | yes | — |
| `remote_target_policy.archival_targets.tier_settings.aws_tiering` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.aws_tiering.tiers` (block) | yes | nested block |
| `remote_target_policy.archival_targets.tier_settings.azure_tiering` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.azure_tiering.tiers` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.google_tiering` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.google_tiering.tiers` (block) | yes | nested block |
| `remote_target_policy.archival_targets.tier_settings.oracle_tiering` (block) | no | nested block |
| `remote_target_policy.archival_targets.tier_settings.oracle_tiering.tiers` (block) | yes | nested block |
| `remote_target_policy.cloud_spin_targets` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.backup_run_type` | no | — |
| `remote_target_policy.cloud_spin_targets.config_id` | no | — |
| `remote_target_policy.cloud_spin_targets.copy_on_run_success` | no | — |
| `remote_target_policy.cloud_spin_targets.log_retention` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.log_retention.duration` | yes | min 1 |
| `remote_target_policy.cloud_spin_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.cloud_spin_targets.retention` (block) | yes | nested block |
| `remote_target_policy.cloud_spin_targets.retention.duration` | yes | min 1 |
| `remote_target_policy.cloud_spin_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.cloud_spin_targets.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.cloud_spin_targets.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.cloud_spin_targets.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.cloud_spin_targets.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.cloud_spin_targets.run_timeouts` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.run_timeouts.backup_type` | no | — |
| `remote_target_policy.cloud_spin_targets.run_timeouts.timeout_mins` | no | — |
| `remote_target_policy.cloud_spin_targets.schedule` (block) | yes | nested block |
| `remote_target_policy.cloud_spin_targets.schedule.frequency` | no | min 1 |
| `remote_target_policy.cloud_spin_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.cloud_spin_targets.target` (block) | yes | nested block |
| `remote_target_policy.cloud_spin_targets.target.id` | no | API: Specifies the unique id of the cloud spin entity. |
| `remote_target_policy.cloud_spin_targets.target.aws_params` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.target.aws_params.region` | yes | API: Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. |
| `remote_target_policy.cloud_spin_targets.target.aws_params.subnet_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.aws_params.vpc_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.aws_params.custom_tag_list` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.target.azure_params` (block) | no | nested block |
| `remote_target_policy.cloud_spin_targets.target.azure_params.availability_set_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.network_resource_group_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.resource_group_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.storage_account_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.storage_container_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.storage_resource_group_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_resource_group_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_storage_account_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_storage_container_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_subnet_id` | no | — |
| `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_virtual_network_id` | no | — |
| `remote_target_policy.onprem_deploy_targets` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.backup_run_type` | no | — |
| `remote_target_policy.onprem_deploy_targets.config_id` | no | — |
| `remote_target_policy.onprem_deploy_targets.copy_on_run_success` | no | — |
| `remote_target_policy.onprem_deploy_targets.log_retention` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.log_retention.duration` | yes | min 1 |
| `remote_target_policy.onprem_deploy_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.onprem_deploy_targets.params` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.params.id` | no | API: Specifies the unique id of the cloud spin entity. |
| `remote_target_policy.onprem_deploy_targets.retention` (block) | yes | nested block |
| `remote_target_policy.onprem_deploy_targets.retention.duration` | yes | min 1 |
| `remote_target_policy.onprem_deploy_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.onprem_deploy_targets.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.onprem_deploy_targets.run_timeouts` (block) | no | nested block |
| `remote_target_policy.onprem_deploy_targets.run_timeouts.backup_type` | no | — |
| `remote_target_policy.onprem_deploy_targets.run_timeouts.timeout_mins` | no | — |
| `remote_target_policy.onprem_deploy_targets.schedule` (block) | yes | nested block |
| `remote_target_policy.onprem_deploy_targets.schedule.frequency` | no | min 1 |
| `remote_target_policy.onprem_deploy_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.replication_targets` (block) | no | nested block |
| `remote_target_policy.replication_targets.backup_run_type` | no | — |
| `remote_target_policy.replication_targets.config_id` | no | — |
| `remote_target_policy.replication_targets.copy_on_run_success` | no | — |
| `remote_target_policy.replication_targets.target_type` | yes | — |
| `remote_target_policy.replication_targets.aws_target_config` (block) | no | nested block |
| `remote_target_policy.replication_targets.aws_target_config.region` | yes | API: Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. |
| `remote_target_policy.replication_targets.aws_target_config.source_id` | yes | — |
| `remote_target_policy.replication_targets.azure_target_config` (block) | no | nested block |
| `remote_target_policy.replication_targets.azure_target_config.resource_group` | no | — |
| `remote_target_policy.replication_targets.azure_target_config.source_id` | yes | — |
| `remote_target_policy.replication_targets.log_retention` (block) | no | nested block |
| `remote_target_policy.replication_targets.log_retention.duration` | yes | min 1 |
| `remote_target_policy.replication_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.replication_targets.log_retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.replication_targets.log_retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.replication_targets.log_retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.replication_targets.log_retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.replication_targets.log_retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.replication_targets.remote_target_config` (block) | no | nested block |
| `remote_target_policy.replication_targets.remote_target_config.cluster_id` | yes | — |
| `remote_target_policy.replication_targets.retention` (block) | yes | nested block |
| `remote_target_policy.replication_targets.retention.duration` | yes | min 1 |
| `remote_target_policy.replication_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.replication_targets.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.replication_targets.retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.replication_targets.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.replication_targets.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.replication_targets.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.replication_targets.run_timeouts` (block) | no | nested block |
| `remote_target_policy.replication_targets.run_timeouts.backup_type` | no | — |
| `remote_target_policy.replication_targets.run_timeouts.timeout_mins` | no | — |
| `remote_target_policy.replication_targets.schedule` (block) | yes | nested block |
| `remote_target_policy.replication_targets.schedule.frequency` | no | min 1 |
| `remote_target_policy.replication_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.rpaas_targets` (block) | no | nested block |
| `remote_target_policy.rpaas_targets.backup_run_type` | no | — |
| `remote_target_policy.rpaas_targets.config_id` | no | — |
| `remote_target_policy.rpaas_targets.copy_on_run_success` | no | — |
| `remote_target_policy.rpaas_targets.target_id` | yes | — |
| `remote_target_policy.rpaas_targets.target_type` | no | — |
| `remote_target_policy.rpaas_targets.log_retention` (block) | no | nested block |
| `remote_target_policy.rpaas_targets.log_retention.duration` | yes | min 1 |
| `remote_target_policy.rpaas_targets.log_retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.rpaas_targets.log_retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.rpaas_targets.log_retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.rpaas_targets.log_retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.rpaas_targets.log_retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.rpaas_targets.log_retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.rpaas_targets.retention` (block) | yes | nested block |
| `remote_target_policy.rpaas_targets.retention.duration` | yes | min 1 |
| `remote_target_policy.rpaas_targets.retention.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.rpaas_targets.retention.data_lock_config` (block) | no | nested block |
| `remote_target_policy.rpaas_targets.retention.data_lock_config.duration` | yes | min 1 |
| `remote_target_policy.rpaas_targets.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `remote_target_policy.rpaas_targets.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `remote_target_policy.rpaas_targets.retention.data_lock_config.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `remote_target_policy.rpaas_targets.run_timeouts` (block) | no | nested block |
| `remote_target_policy.rpaas_targets.run_timeouts.backup_type` | no | — |
| `remote_target_policy.rpaas_targets.run_timeouts.timeout_mins` | no | — |
| `remote_target_policy.rpaas_targets.schedule` (block) | yes | nested block |
| `remote_target_policy.rpaas_targets.schedule.frequency` | no | min 1 |
| `remote_target_policy.rpaas_targets.schedule.unit` | yes | one of: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years` |
| `retry_options` (block) | no | nested block |
| `retry_options.retries` | no | min 0 |
| `retry_options.retry_interval_mins` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_policy_invalid_data_lock` | `data_lock` |
| `ibm_backup_recovery_protection_policy_invalid_description` | `description` |
| `ibm_backup_recovery_protection_policy_invalid_is_cbs_enabled` | `is_cbs_enabled` |
| `ibm_backup_recovery_protection_policy_invalid_name` | `name` |
| `ibm_backup_recovery_protection_policy_invalid_template_id` | `template_id` |
| `ibm_backup_recovery_protection_policy_invalid_version` | `version` |
| `ibm_backup_recovery_protection_policy_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__duration` | `backup_policy.bmr.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__unit` | `backup_policy.bmr.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__data_lock_config__duration` | `backup_policy.bmr.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__data_lock_config__enable_worm_on_external_target` | `backup_policy.bmr.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__data_lock_config__mode` | `backup_policy.bmr.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__retention__data_lock_config__unit` | `backup_policy.bmr.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__unit` | `backup_policy.bmr.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__day_schedule__frequency` | `backup_policy.bmr.schedule.day_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__month_schedule__day_of_month` | `backup_policy.bmr.schedule.month_schedule.day_of_month` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__month_schedule__day_of_week` | `backup_policy.bmr.schedule.month_schedule.day_of_week` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__month_schedule__week_of_month` | `backup_policy.bmr.schedule.month_schedule.week_of_month` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__week_schedule__day_of_week` | `backup_policy.bmr.schedule.week_schedule.day_of_week` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__bmr__schedule__year_schedule__day_of_year` | `backup_policy.bmr.schedule.year_schedule.day_of_year` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__duration` | `backup_policy.cdp.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__unit` | `backup_policy.cdp.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__data_lock_config__duration` | `backup_policy.cdp.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__data_lock_config__enable_worm_on_external_target` | `backup_policy.cdp.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__data_lock_config__mode` | `backup_policy.cdp.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__cdp__retention__data_lock_config__unit` | `backup_policy.cdp.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__duration` | `backup_policy.log.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__unit` | `backup_policy.log.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__data_lock_config__duration` | `backup_policy.log.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__data_lock_config__enable_worm_on_external_target` | `backup_policy.log.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__data_lock_config__mode` | `backup_policy.log.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__retention__data_lock_config__unit` | `backup_policy.log.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__schedule__unit` | `backup_policy.log.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__schedule__hour_schedule__frequency` | `backup_policy.log.schedule.hour_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__log__schedule__minute_schedule__frequency` | `backup_policy.log.schedule.minute_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__full__schedule__unit` | `backup_policy.regular.full.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__full_backups__retention__duration` | `backup_policy.regular.full_backups.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__full_backups__retention__unit` | `backup_policy.regular.full_backups.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__full_backups__schedule__unit` | `backup_policy.regular.full_backups.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__incremental__schedule__unit` | `backup_policy.regular.incremental.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__primary_backup_target__target_type` | `backup_policy.regular.primary_backup_target.target_type` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__primary_backup_target__use_default_backup_target` | `backup_policy.regular.primary_backup_target.use_default_backup_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__primary_backup_target__archival_target_settings__target_id` | `backup_policy.regular.primary_backup_target.archival_target_settings.target_id` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__duration` | `backup_policy.regular.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__unit` | `backup_policy.regular.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__data_lock_config__duration` | `backup_policy.regular.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__data_lock_config__enable_worm_on_external_target` | `backup_policy.regular.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__data_lock_config__mode` | `backup_policy.regular.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__regular__retention__data_lock_config__unit` | `backup_policy.regular.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__run_timeouts__backup_type` | `backup_policy.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__run_timeouts__timeout_mins` | `backup_policy.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__duration` | `backup_policy.storage_array_snapshot.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__unit` | `backup_policy.storage_array_snapshot.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__data_lock_config__duration` | `backup_policy.storage_array_snapshot.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__data_lock_config__enable_worm_on_external_target` | `backup_policy.storage_array_snapshot.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__data_lock_config__mode` | `backup_policy.storage_array_snapshot.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__retention__data_lock_config__unit` | `backup_policy.storage_array_snapshot.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__unit` | `backup_policy.storage_array_snapshot.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__day_schedule__frequency` | `backup_policy.storage_array_snapshot.schedule.day_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__hour_schedule__frequency` | `backup_policy.storage_array_snapshot.schedule.hour_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__minute_schedule__frequency` | `backup_policy.storage_array_snapshot.schedule.minute_schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__month_schedule__day_of_month` | `backup_policy.storage_array_snapshot.schedule.month_schedule.day_of_month` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__month_schedule__day_of_week` | `backup_policy.storage_array_snapshot.schedule.month_schedule.day_of_week` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__month_schedule__week_of_month` | `backup_policy.storage_array_snapshot.schedule.month_schedule.week_of_month` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__week_schedule__day_of_week` | `backup_policy.storage_array_snapshot.schedule.week_schedule.day_of_week` |
| `ibm_backup_recovery_protection_policy_invalid_backup_policy__storage_array_snapshot__schedule__year_schedule__day_of_year` | `backup_policy.storage_array_snapshot.schedule.year_schedule.day_of_year` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__config_id` | `blackout_window.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__day` | `blackout_window.day` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__end_time__hour` | `blackout_window.end_time.hour` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__end_time__minute` | `blackout_window.end_time.minute` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__end_time__time_zone` | `blackout_window.end_time.time_zone` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__start_time__hour` | `blackout_window.start_time.hour` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__start_time__minute` | `blackout_window.start_time.minute` |
| `ibm_backup_recovery_protection_policy_invalid_blackout_window__start_time__time_zone` | `blackout_window.start_time.time_zone` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__source_cluster_id` | `cascaded_targets_config.source_cluster_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__backup_run_type` | `cascaded_targets_config.remote_targets.archival_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__config_id` | `cascaded_targets_config.remote_targets.archival_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__copy_on_run_success` | `cascaded_targets_config.remote_targets.archival_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__target_id` | `cascaded_targets_config.remote_targets.archival_targets.target_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__extended_retention__config_id` | `cascaded_targets_config.remote_targets.archival_targets.extended_retention.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__extended_retention__run_type` | `cascaded_targets_config.remote_targets.archival_targets.extended_retention.run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__log_retention__duration` | `cascaded_targets_config.remote_targets.archival_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__log_retention__unit` | `cascaded_targets_config.remote_targets.archival_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__retention__duration` | `cascaded_targets_config.remote_targets.archival_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__retention__unit` | `cascaded_targets_config.remote_targets.archival_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__run_timeouts__backup_type` | `cascaded_targets_config.remote_targets.archival_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__run_timeouts__timeout_mins` | `cascaded_targets_config.remote_targets.archival_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__schedule__frequency` | `cascaded_targets_config.remote_targets.archival_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__schedule__unit` | `cascaded_targets_config.remote_targets.archival_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__archival_targets__tier_settings__cloud_platform` | `cascaded_targets_config.remote_targets.archival_targets.tier_settings.cloud_platform` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__backup_run_type` | `cascaded_targets_config.remote_targets.cloud_spin_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__config_id` | `cascaded_targets_config.remote_targets.cloud_spin_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__copy_on_run_success` | `cascaded_targets_config.remote_targets.cloud_spin_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__log_retention__duration` | `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__log_retention__unit` | `cascaded_targets_config.remote_targets.cloud_spin_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__retention__duration` | `cascaded_targets_config.remote_targets.cloud_spin_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__retention__unit` | `cascaded_targets_config.remote_targets.cloud_spin_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__run_timeouts__backup_type` | `cascaded_targets_config.remote_targets.cloud_spin_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__run_timeouts__timeout_mins` | `cascaded_targets_config.remote_targets.cloud_spin_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__schedule__frequency` | `cascaded_targets_config.remote_targets.cloud_spin_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__schedule__unit` | `cascaded_targets_config.remote_targets.cloud_spin_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__cloud_spin_targets__target__id` | `cascaded_targets_config.remote_targets.cloud_spin_targets.target.id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__backup_run_type` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__config_id` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__copy_on_run_success` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__log_retention__duration` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__log_retention__unit` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__params__id` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.params.id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__retention__duration` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__retention__unit` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__run_timeouts__backup_type` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__run_timeouts__timeout_mins` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__schedule__frequency` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__onprem_deploy_targets__schedule__unit` | `cascaded_targets_config.remote_targets.onprem_deploy_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__backup_run_type` | `cascaded_targets_config.remote_targets.replication_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__config_id` | `cascaded_targets_config.remote_targets.replication_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__copy_on_run_success` | `cascaded_targets_config.remote_targets.replication_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__target_type` | `cascaded_targets_config.remote_targets.replication_targets.target_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__aws_target_config__region` | `cascaded_targets_config.remote_targets.replication_targets.aws_target_config.region` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__aws_target_config__source_id` | `cascaded_targets_config.remote_targets.replication_targets.aws_target_config.source_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__azure_target_config__resource_group` | `cascaded_targets_config.remote_targets.replication_targets.azure_target_config.resource_group` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__azure_target_config__source_id` | `cascaded_targets_config.remote_targets.replication_targets.azure_target_config.source_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__log_retention__duration` | `cascaded_targets_config.remote_targets.replication_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__log_retention__unit` | `cascaded_targets_config.remote_targets.replication_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__remote_target_config__cluster_id` | `cascaded_targets_config.remote_targets.replication_targets.remote_target_config.cluster_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__retention__duration` | `cascaded_targets_config.remote_targets.replication_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__retention__unit` | `cascaded_targets_config.remote_targets.replication_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__run_timeouts__backup_type` | `cascaded_targets_config.remote_targets.replication_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__run_timeouts__timeout_mins` | `cascaded_targets_config.remote_targets.replication_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__schedule__frequency` | `cascaded_targets_config.remote_targets.replication_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__replication_targets__schedule__unit` | `cascaded_targets_config.remote_targets.replication_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__backup_run_type` | `cascaded_targets_config.remote_targets.rpaas_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__config_id` | `cascaded_targets_config.remote_targets.rpaas_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__copy_on_run_success` | `cascaded_targets_config.remote_targets.rpaas_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__target_id` | `cascaded_targets_config.remote_targets.rpaas_targets.target_id` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__target_type` | `cascaded_targets_config.remote_targets.rpaas_targets.target_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__log_retention__duration` | `cascaded_targets_config.remote_targets.rpaas_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__log_retention__unit` | `cascaded_targets_config.remote_targets.rpaas_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__retention__duration` | `cascaded_targets_config.remote_targets.rpaas_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__retention__unit` | `cascaded_targets_config.remote_targets.rpaas_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__run_timeouts__backup_type` | `cascaded_targets_config.remote_targets.rpaas_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__run_timeouts__timeout_mins` | `cascaded_targets_config.remote_targets.rpaas_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__schedule__frequency` | `cascaded_targets_config.remote_targets.rpaas_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_cascaded_targets_config__remote_targets__rpaas_targets__schedule__unit` | `cascaded_targets_config.remote_targets.rpaas_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__config_id` | `extended_retention.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__run_type` | `extended_retention.run_type` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__duration` | `extended_retention.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__unit` | `extended_retention.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__data_lock_config__duration` | `extended_retention.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__data_lock_config__enable_worm_on_external_target` | `extended_retention.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__data_lock_config__mode` | `extended_retention.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__retention__data_lock_config__unit` | `extended_retention.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__schedule__frequency` | `extended_retention.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_extended_retention__schedule__unit` | `extended_retention.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__backup_run_type` | `remote_target_policy.archival_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__config_id` | `remote_target_policy.archival_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__copy_on_run_success` | `remote_target_policy.archival_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__target_id` | `remote_target_policy.archival_targets.target_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__config_id` | `remote_target_policy.archival_targets.extended_retention.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__run_type` | `remote_target_policy.archival_targets.extended_retention.run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__retention__duration` | `remote_target_policy.archival_targets.extended_retention.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__retention__unit` | `remote_target_policy.archival_targets.extended_retention.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__schedule__frequency` | `remote_target_policy.archival_targets.extended_retention.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__extended_retention__schedule__unit` | `remote_target_policy.archival_targets.extended_retention.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__duration` | `remote_target_policy.archival_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__unit` | `remote_target_policy.archival_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__data_lock_config__duration` | `remote_target_policy.archival_targets.log_retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.archival_targets.log_retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__data_lock_config__mode` | `remote_target_policy.archival_targets.log_retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__log_retention__data_lock_config__unit` | `remote_target_policy.archival_targets.log_retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__duration` | `remote_target_policy.archival_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__unit` | `remote_target_policy.archival_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__data_lock_config__duration` | `remote_target_policy.archival_targets.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.archival_targets.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__data_lock_config__mode` | `remote_target_policy.archival_targets.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__retention__data_lock_config__unit` | `remote_target_policy.archival_targets.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__run_timeouts__backup_type` | `remote_target_policy.archival_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__run_timeouts__timeout_mins` | `remote_target_policy.archival_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__schedule__frequency` | `remote_target_policy.archival_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__schedule__unit` | `remote_target_policy.archival_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__archival_targets__tier_settings__cloud_platform` | `remote_target_policy.archival_targets.tier_settings.cloud_platform` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__backup_run_type` | `remote_target_policy.cloud_spin_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__config_id` | `remote_target_policy.cloud_spin_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__copy_on_run_success` | `remote_target_policy.cloud_spin_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__duration` | `remote_target_policy.cloud_spin_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__unit` | `remote_target_policy.cloud_spin_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__data_lock_config__duration` | `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__data_lock_config__mode` | `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__log_retention__data_lock_config__unit` | `remote_target_policy.cloud_spin_targets.log_retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__duration` | `remote_target_policy.cloud_spin_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__unit` | `remote_target_policy.cloud_spin_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__data_lock_config__duration` | `remote_target_policy.cloud_spin_targets.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.cloud_spin_targets.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__data_lock_config__mode` | `remote_target_policy.cloud_spin_targets.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__retention__data_lock_config__unit` | `remote_target_policy.cloud_spin_targets.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__run_timeouts__backup_type` | `remote_target_policy.cloud_spin_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__run_timeouts__timeout_mins` | `remote_target_policy.cloud_spin_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__schedule__frequency` | `remote_target_policy.cloud_spin_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__schedule__unit` | `remote_target_policy.cloud_spin_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__id` | `remote_target_policy.cloud_spin_targets.target.id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__aws_params__region` | `remote_target_policy.cloud_spin_targets.target.aws_params.region` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__aws_params__subnet_id` | `remote_target_policy.cloud_spin_targets.target.aws_params.subnet_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__aws_params__vpc_id` | `remote_target_policy.cloud_spin_targets.target.aws_params.vpc_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__availability_set_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.availability_set_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__network_resource_group_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.network_resource_group_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__resource_group_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.resource_group_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__storage_account_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.storage_account_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__storage_container_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.storage_container_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__storage_resource_group_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.storage_resource_group_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__temp_vm_resource_group_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_resource_group_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__temp_vm_storage_account_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_storage_account_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__temp_vm_storage_container_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_storage_container_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__temp_vm_subnet_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_subnet_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__cloud_spin_targets__target__azure_params__temp_vm_virtual_network_id` | `remote_target_policy.cloud_spin_targets.target.azure_params.temp_vm_virtual_network_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__backup_run_type` | `remote_target_policy.onprem_deploy_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__config_id` | `remote_target_policy.onprem_deploy_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__copy_on_run_success` | `remote_target_policy.onprem_deploy_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__duration` | `remote_target_policy.onprem_deploy_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__unit` | `remote_target_policy.onprem_deploy_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__data_lock_config__duration` | `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__data_lock_config__mode` | `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__log_retention__data_lock_config__unit` | `remote_target_policy.onprem_deploy_targets.log_retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__params__id` | `remote_target_policy.onprem_deploy_targets.params.id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__duration` | `remote_target_policy.onprem_deploy_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__unit` | `remote_target_policy.onprem_deploy_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__data_lock_config__duration` | `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__data_lock_config__mode` | `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__retention__data_lock_config__unit` | `remote_target_policy.onprem_deploy_targets.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__run_timeouts__backup_type` | `remote_target_policy.onprem_deploy_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__run_timeouts__timeout_mins` | `remote_target_policy.onprem_deploy_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__schedule__frequency` | `remote_target_policy.onprem_deploy_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__onprem_deploy_targets__schedule__unit` | `remote_target_policy.onprem_deploy_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__backup_run_type` | `remote_target_policy.replication_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__config_id` | `remote_target_policy.replication_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__copy_on_run_success` | `remote_target_policy.replication_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__target_type` | `remote_target_policy.replication_targets.target_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__aws_target_config__region` | `remote_target_policy.replication_targets.aws_target_config.region` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__aws_target_config__source_id` | `remote_target_policy.replication_targets.aws_target_config.source_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__azure_target_config__resource_group` | `remote_target_policy.replication_targets.azure_target_config.resource_group` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__azure_target_config__source_id` | `remote_target_policy.replication_targets.azure_target_config.source_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__duration` | `remote_target_policy.replication_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__unit` | `remote_target_policy.replication_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__data_lock_config__duration` | `remote_target_policy.replication_targets.log_retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.replication_targets.log_retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__data_lock_config__mode` | `remote_target_policy.replication_targets.log_retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__log_retention__data_lock_config__unit` | `remote_target_policy.replication_targets.log_retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__remote_target_config__cluster_id` | `remote_target_policy.replication_targets.remote_target_config.cluster_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__duration` | `remote_target_policy.replication_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__unit` | `remote_target_policy.replication_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__data_lock_config__duration` | `remote_target_policy.replication_targets.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.replication_targets.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__data_lock_config__mode` | `remote_target_policy.replication_targets.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__retention__data_lock_config__unit` | `remote_target_policy.replication_targets.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__run_timeouts__backup_type` | `remote_target_policy.replication_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__run_timeouts__timeout_mins` | `remote_target_policy.replication_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__schedule__frequency` | `remote_target_policy.replication_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__replication_targets__schedule__unit` | `remote_target_policy.replication_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__backup_run_type` | `remote_target_policy.rpaas_targets.backup_run_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__config_id` | `remote_target_policy.rpaas_targets.config_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__copy_on_run_success` | `remote_target_policy.rpaas_targets.copy_on_run_success` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__target_id` | `remote_target_policy.rpaas_targets.target_id` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__target_type` | `remote_target_policy.rpaas_targets.target_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__duration` | `remote_target_policy.rpaas_targets.log_retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__unit` | `remote_target_policy.rpaas_targets.log_retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__data_lock_config__duration` | `remote_target_policy.rpaas_targets.log_retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.rpaas_targets.log_retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__data_lock_config__mode` | `remote_target_policy.rpaas_targets.log_retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__log_retention__data_lock_config__unit` | `remote_target_policy.rpaas_targets.log_retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__duration` | `remote_target_policy.rpaas_targets.retention.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__unit` | `remote_target_policy.rpaas_targets.retention.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__data_lock_config__duration` | `remote_target_policy.rpaas_targets.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__data_lock_config__enable_worm_on_external_target` | `remote_target_policy.rpaas_targets.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__data_lock_config__mode` | `remote_target_policy.rpaas_targets.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__retention__data_lock_config__unit` | `remote_target_policy.rpaas_targets.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__run_timeouts__backup_type` | `remote_target_policy.rpaas_targets.run_timeouts.backup_type` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__run_timeouts__timeout_mins` | `remote_target_policy.rpaas_targets.run_timeouts.timeout_mins` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__schedule__frequency` | `remote_target_policy.rpaas_targets.schedule.frequency` |
| `ibm_backup_recovery_protection_policy_invalid_remote_target_policy__rpaas_targets__schedule__unit` | `remote_target_policy.rpaas_targets.schedule.unit` |
| `ibm_backup_recovery_protection_policy_invalid_retry_options__retries` | `retry_options.retries` |
| `ibm_backup_recovery_protection_policy_invalid_retry_options__retry_interval_mins` | `retry_options.retry_interval_mins` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_protection_policy" "test" {
  name = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  backup_policy {
    regular {
    }
  }
}
```
