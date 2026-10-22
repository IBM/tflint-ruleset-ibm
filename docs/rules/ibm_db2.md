# ibm_db2

Manages ibm_db2. A name for the resource instance

This rule validates `ibm_db2` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/manage/update_db2_licenses:put` (`db2-on-cloud-db2-saas-perf-v4`).

## Configuration

```hcl
rule "ibm_db2" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allow_cleanup` | no | — |
| `backup_location` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `dashboard_url` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `disk_encryption_instance_crn` | no | crn format |
| `disk_encryption_key_crn` | no | crn format |
| `extensions` | no | — |
| `guid` | no | — |
| `high_availability` | no | — |
| `instance_type` | no | — |
| `last_operation` | no | — |
| `location` | yes | — |
| `locked` | no | — |
| `name` | yes | — |
| `onetime_credentials` | no | — |
| `oracle_compatibility` | no | — |
| `parameters` | no | — |
| `parameters_json` | no | — |
| `plan` | yes | — |
| `plan_history` | no | — |
| `resource_aliases_url` | no | — |
| `resource_bindings_url` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_id` | no | — |
| `resource_keys_url` | no | — |
| `resource_name` | no | — |
| `resource_plan_id` | no | — |
| `resource_status` | no | — |
| `restored_at` | no | — |
| `restored_by` | no | — |
| `scheduled_reclaim_at` | no | — |
| `scheduled_reclaim_by` | no | — |
| `service` | yes | — |
| `service_endpoints` | no | one of: `public`, `private`, `public-and-private` |
| `state` | no | — |
| `status` | no | — |
| `sub_type` | no | — |
| `subscription_id` | no | — |
| `tags` | no | — |
| `target_crn` | no | — |
| `type` | no | — |
| `update_at` | no | — |
| `update_by` | no | — |
| `allowlist_config` (block) | no | nested block |
| `allowlist_config.ip_addresses` (block) | no | nested block |
| `allowlist_config.ip_addresses.address` | no | — |
| `allowlist_config.ip_addresses.description` | no | — |
| `autoscale_config` (block) | no | nested block |
| `autoscale_config.auto_scaling_allow_plan_limit` | no | one of: `YES`, `NO` |
| `autoscale_config.auto_scaling_enabled` | no | one of: `YES`, `NO` |
| `autoscale_config.auto_scaling_over_time_period` | no | one of: `5`, `15`, `30`, `45`, `60` |
| `autoscale_config.auto_scaling_pause_limit` | no | — |
| `autoscale_config.auto_scaling_threshold` | no | — |
| `custom_setting_config` (block) | no | nested block |
| `custom_setting_config.db` (block) | no | nested block |
| `custom_setting_config.db.act_sortmem_limit` | no | — |
| `custom_setting_config.db.alt_collate` | no | — |
| `custom_setting_config.db.app_ctl_heap_sz` | no | — |
| `custom_setting_config.db.appgroup_mem_sz` | no | — |
| `custom_setting_config.db.appl_memory` | no | — |
| `custom_setting_config.db.applheapsz` | no | — |
| `custom_setting_config.db.archretrydelay` | no | — |
| `custom_setting_config.db.authn_cache_duration` | no | — |
| `custom_setting_config.db.auto_cg_stats` | no | — |
| `custom_setting_config.db.auto_maint` | no | — |
| `custom_setting_config.db.auto_reorg` | no | — |
| `custom_setting_config.db.auto_reval` | no | — |
| `custom_setting_config.db.auto_runstats` | no | — |
| `custom_setting_config.db.auto_sampling` | no | — |
| `custom_setting_config.db.auto_stats_views` | no | — |
| `custom_setting_config.db.auto_stmt_stats` | no | — |
| `custom_setting_config.db.auto_tbl_maint` | no | — |
| `custom_setting_config.db.autorestart` | no | — |
| `custom_setting_config.db.avg_appls` | no | — |
| `custom_setting_config.db.catalogcache_sz` | no | — |
| `custom_setting_config.db.chngpgs_thresh` | no | — |
| `custom_setting_config.db.cur_commit` | no | — |
| `custom_setting_config.db.database_memory` | no | — |
| `custom_setting_config.db.db_collname` | no | — |
| `custom_setting_config.db.db_mem_thresh` | no | — |
| `custom_setting_config.db.dbheap` | no | — |
| `custom_setting_config.db.ddl_compression_def` | no | — |
| `custom_setting_config.db.ddl_constraint_def` | no | — |
| `custom_setting_config.db.dec_arithmetic` | no | — |
| `custom_setting_config.db.dec_to_char_fmt` | no | — |
| `custom_setting_config.db.decflt_rounding` | no | — |
| `custom_setting_config.db.dft_degree` | no | — |
| `custom_setting_config.db.dft_extent_sz` | no | — |
| `custom_setting_config.db.dft_loadrec_ses` | no | — |
| `custom_setting_config.db.dft_mttb_types` | no | — |
| `custom_setting_config.db.dft_prefetch_sz` | no | — |
| `custom_setting_config.db.dft_queryopt` | no | — |
| `custom_setting_config.db.dft_refresh_age` | no | — |
| `custom_setting_config.db.dft_schemas_dcc` | no | — |
| `custom_setting_config.db.dft_sqlmathwarn` | no | — |
| `custom_setting_config.db.dft_table_org` | no | — |
| `custom_setting_config.db.dlchktime` | no | — |
| `custom_setting_config.db.enable_xmlchar` | no | — |
| `custom_setting_config.db.extended_row_sz` | no | — |
| `custom_setting_config.db.groupheap_ratio` | no | — |
| `custom_setting_config.db.indexrec` | no | — |
| `custom_setting_config.db.large_aggregation` | no | — |
| `custom_setting_config.db.locklist` | no | — |
| `custom_setting_config.db.locktimeout` | no | — |
| `custom_setting_config.db.log_appl_info` | no | — |
| `custom_setting_config.db.log_ddl_stmts` | no | — |
| `custom_setting_config.db.log_disk_cap` | no | — |
| `custom_setting_config.db.logindexbuild` | no | — |
| `custom_setting_config.db.maxappls` | no | — |
| `custom_setting_config.db.maxfilop` | no | — |
| `custom_setting_config.db.maxlocks` | no | — |
| `custom_setting_config.db.min_dec_div_3` | no | — |
| `custom_setting_config.db.mon_act_metrics` | no | — |
| `custom_setting_config.db.mon_deadlock` | no | — |
| `custom_setting_config.db.mon_lck_msg_lvl` | no | — |
| `custom_setting_config.db.mon_locktimeout` | no | — |
| `custom_setting_config.db.mon_lockwait` | no | — |
| `custom_setting_config.db.mon_lw_thresh` | no | — |
| `custom_setting_config.db.mon_obj_metrics` | no | — |
| `custom_setting_config.db.mon_pkglist_sz` | no | — |
| `custom_setting_config.db.mon_req_metrics` | no | — |
| `custom_setting_config.db.mon_rtn_data` | no | — |
| `custom_setting_config.db.mon_rtn_execlist` | no | — |
| `custom_setting_config.db.mon_uow_data` | no | — |
| `custom_setting_config.db.mon_uow_execlist` | no | — |
| `custom_setting_config.db.mon_uow_pkglist` | no | — |
| `custom_setting_config.db.nchar_mapping` | no | — |
| `custom_setting_config.db.num_freqvalues` | no | — |
| `custom_setting_config.db.num_iocleaners` | no | — |
| `custom_setting_config.db.num_ioservers` | no | — |
| `custom_setting_config.db.num_log_span` | no | — |
| `custom_setting_config.db.num_quantiles` | no | — |
| `custom_setting_config.db.opt_buffpage` | no | — |
| `custom_setting_config.db.opt_direct_wrkld` | no | — |
| `custom_setting_config.db.opt_locklist` | no | — |
| `custom_setting_config.db.opt_maxlocks` | no | — |
| `custom_setting_config.db.opt_sortheap` | no | — |
| `custom_setting_config.db.page_age_trgt_gcr` | no | — |
| `custom_setting_config.db.page_age_trgt_mcr` | no | — |
| `custom_setting_config.db.pckcachesz` | no | — |
| `custom_setting_config.db.pl_stack_trace` | no | — |
| `custom_setting_config.db.self_tuning_mem` | no | — |
| `custom_setting_config.db.seqdetect` | no | — |
| `custom_setting_config.db.sheapthres_shr` | no | — |
| `custom_setting_config.db.softmax` | no | — |
| `custom_setting_config.db.sortheap` | no | — |
| `custom_setting_config.db.sql_ccflags` | no | — |
| `custom_setting_config.db.stat_heap_sz` | no | — |
| `custom_setting_config.db.stmt_conc` | no | — |
| `custom_setting_config.db.stmtheap` | no | — |
| `custom_setting_config.db.string_units` | no | — |
| `custom_setting_config.db.systime_period_adj` | no | — |
| `custom_setting_config.db.trackmod` | no | — |
| `custom_setting_config.db.util_heap_sz` | no | — |
| `custom_setting_config.db.wlm_admission_ctrl` | no | — |
| `custom_setting_config.db.wlm_agent_load_trgt` | no | — |
| `custom_setting_config.db.wlm_cpu_limit` | no | — |
| `custom_setting_config.db.wlm_cpu_share_mode` | no | — |
| `custom_setting_config.db.wlm_cpu_shares` | no | — |
| `custom_setting_config.dbm` (block) | no | nested block |
| `custom_setting_config.dbm.comm_bandwidth` | no | — |
| `custom_setting_config.dbm.cpuspeed` | no | — |
| `custom_setting_config.dbm.dft_mon_bufpool` | no | — |
| `custom_setting_config.dbm.dft_mon_lock` | no | — |
| `custom_setting_config.dbm.dft_mon_sort` | no | — |
| `custom_setting_config.dbm.dft_mon_stmt` | no | — |
| `custom_setting_config.dbm.dft_mon_table` | no | — |
| `custom_setting_config.dbm.dft_mon_timestamp` | no | — |
| `custom_setting_config.dbm.dft_mon_uow` | no | — |
| `custom_setting_config.dbm.diaglevel` | no | — |
| `custom_setting_config.dbm.federated_async` | no | — |
| `custom_setting_config.dbm.indexrec` | no | — |
| `custom_setting_config.dbm.intra_parallel` | no | — |
| `custom_setting_config.dbm.keepfenced` | no | — |
| `custom_setting_config.dbm.max_connretries` | no | — |
| `custom_setting_config.dbm.max_querydegree` | no | — |
| `custom_setting_config.dbm.mon_heap_sz` | no | — |
| `custom_setting_config.dbm.multipartsizemb` | no | — |
| `custom_setting_config.dbm.notifylevel` | no | — |
| `custom_setting_config.dbm.num_initagents` | no | — |
| `custom_setting_config.dbm.num_initfenced` | no | — |
| `custom_setting_config.dbm.num_poolagents` | no | — |
| `custom_setting_config.dbm.resync_interval` | no | — |
| `custom_setting_config.dbm.rqrioblk` | no | — |
| `custom_setting_config.dbm.start_stop_time` | no | — |
| `custom_setting_config.dbm.util_impact_lim` | no | — |
| `custom_setting_config.dbm.wlm_disp_concur` | no | — |
| `custom_setting_config.dbm.wlm_disp_cpu_shares` | no | — |
| `custom_setting_config.dbm.wlm_disp_min_util` | no | — |
| `custom_setting_config.dbm.wlm_dispatcher` | no | — |
| `custom_setting_config.registry` (block) | no | nested block |
| `custom_setting_config.registry.db2_alternate_authz_behaviour` | no | — |
| `custom_setting_config.registry.db2_antijoin` | no | — |
| `custom_setting_config.registry.db2_ats_enable` | no | — |
| `custom_setting_config.registry.db2_bidi` | no | — |
| `custom_setting_config.registry.db2_compopt` | no | — |
| `custom_setting_config.registry.db2_deferred_prepare_semantics` | no | — |
| `custom_setting_config.registry.db2_evaluncommitted` | no | — |
| `custom_setting_config.registry.db2_extended_optimization` | no | — |
| `custom_setting_config.registry.db2_index_pctfree_default` | no | — |
| `custom_setting_config.registry.db2_inlist_to_nljn` | no | — |
| `custom_setting_config.registry.db2_lock_to_rb` | no | — |
| `custom_setting_config.registry.db2_minimize_listprefetch` | no | — |
| `custom_setting_config.registry.db2_object_table_entries` | no | — |
| `custom_setting_config.registry.db2_opt_max_temp_size` | no | — |
| `custom_setting_config.registry.db2_optprofile` | no | — |
| `custom_setting_config.registry.db2_optstats_log` | no | — |
| `custom_setting_config.registry.db2_parallel_io` | no | — |
| `custom_setting_config.registry.db2_reduced_optimization` | no | — |
| `custom_setting_config.registry.db2_selectivity` | no | — |
| `custom_setting_config.registry.db2_skipdeleted` | no | — |
| `custom_setting_config.registry.db2_skipinserted` | no | — |
| `custom_setting_config.registry.db2_stmm` | no | — |
| `custom_setting_config.registry.db2_sync_release_lock_attributes` | no | — |
| `custom_setting_config.registry.db2_truncate_reusestorage` | no | — |
| `custom_setting_config.registry.db2_use_alternate_page_cleaning` | no | — |
| `custom_setting_config.registry.db2_view_reopt_values` | no | — |
| `custom_setting_config.registry.db2_wlm_settings` | no | — |
| `custom_setting_config.registry.db2_workload` | no | — |
| `users_config` (block) | no | nested block |
| `users_config.email` | no | — |
| `users_config.iam` | no | — |
| `users_config.ibmid` | no | — |
| `users_config.id` | no | — |
| `users_config.locked` | no | one of: `yes`, `no` |
| `users_config.name` | no | — |
| `users_config.password` | no | — |
| `users_config.role` | no | — |
| `users_config.authentication` (block) | no | nested block |
| `users_config.authentication.method` | no | one of: `internal` |
| `users_config.authentication.policy_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_db2_invalid_backup_location` | `backup_location` |
| `ibm_db2_invalid_disk_encryption_instance_crn` | `disk_encryption_instance_crn` |
| `ibm_db2_invalid_disk_encryption_key_crn` | `disk_encryption_key_crn` |
| `ibm_db2_invalid_high_availability` | `high_availability` |
| `ibm_db2_invalid_instance_type` | `instance_type` |
| `ibm_db2_invalid_location` | `location` |
| `ibm_db2_invalid_name` | `name` |
| `ibm_db2_invalid_oracle_compatibility` | `oracle_compatibility` |
| `ibm_db2_invalid_parameters` | `parameters` |
| `ibm_db2_invalid_parameters_json` | `parameters_json` |
| `ibm_db2_invalid_plan` | `plan` |
| `ibm_db2_invalid_resource_group_id` | `resource_group_id` |
| `ibm_db2_invalid_service` | `service` |
| `ibm_db2_invalid_service_endpoints` | `service_endpoints` |
| `ibm_db2_invalid_subscription_id` | `subscription_id` |
| `ibm_db2_invalid_tags` | `tags` |
| `ibm_db2_invalid_allowlist_config__ip_addresses__address` | `allowlist_config.ip_addresses.address` |
| `ibm_db2_invalid_allowlist_config__ip_addresses__description` | `allowlist_config.ip_addresses.description` |
| `ibm_db2_invalid_autoscale_config__auto_scaling_allow_plan_limit` | `autoscale_config.auto_scaling_allow_plan_limit` |
| `ibm_db2_invalid_autoscale_config__auto_scaling_enabled` | `autoscale_config.auto_scaling_enabled` |
| `ibm_db2_invalid_autoscale_config__auto_scaling_over_time_period` | `autoscale_config.auto_scaling_over_time_period` |
| `ibm_db2_invalid_autoscale_config__auto_scaling_pause_limit` | `autoscale_config.auto_scaling_pause_limit` |
| `ibm_db2_invalid_autoscale_config__auto_scaling_threshold` | `autoscale_config.auto_scaling_threshold` |
| `ibm_db2_invalid_custom_setting_config__db__act_sortmem_limit` | `custom_setting_config.db.act_sortmem_limit` |
| `ibm_db2_invalid_custom_setting_config__db__alt_collate` | `custom_setting_config.db.alt_collate` |
| `ibm_db2_invalid_custom_setting_config__db__app_ctl_heap_sz` | `custom_setting_config.db.app_ctl_heap_sz` |
| `ibm_db2_invalid_custom_setting_config__db__appgroup_mem_sz` | `custom_setting_config.db.appgroup_mem_sz` |
| `ibm_db2_invalid_custom_setting_config__db__appl_memory` | `custom_setting_config.db.appl_memory` |
| `ibm_db2_invalid_custom_setting_config__db__applheapsz` | `custom_setting_config.db.applheapsz` |
| `ibm_db2_invalid_custom_setting_config__db__archretrydelay` | `custom_setting_config.db.archretrydelay` |
| `ibm_db2_invalid_custom_setting_config__db__authn_cache_duration` | `custom_setting_config.db.authn_cache_duration` |
| `ibm_db2_invalid_custom_setting_config__db__auto_cg_stats` | `custom_setting_config.db.auto_cg_stats` |
| `ibm_db2_invalid_custom_setting_config__db__auto_maint` | `custom_setting_config.db.auto_maint` |
| `ibm_db2_invalid_custom_setting_config__db__auto_reorg` | `custom_setting_config.db.auto_reorg` |
| `ibm_db2_invalid_custom_setting_config__db__auto_reval` | `custom_setting_config.db.auto_reval` |
| `ibm_db2_invalid_custom_setting_config__db__auto_runstats` | `custom_setting_config.db.auto_runstats` |
| `ibm_db2_invalid_custom_setting_config__db__auto_sampling` | `custom_setting_config.db.auto_sampling` |
| `ibm_db2_invalid_custom_setting_config__db__auto_stats_views` | `custom_setting_config.db.auto_stats_views` |
| `ibm_db2_invalid_custom_setting_config__db__auto_stmt_stats` | `custom_setting_config.db.auto_stmt_stats` |
| `ibm_db2_invalid_custom_setting_config__db__auto_tbl_maint` | `custom_setting_config.db.auto_tbl_maint` |
| `ibm_db2_invalid_custom_setting_config__db__autorestart` | `custom_setting_config.db.autorestart` |
| `ibm_db2_invalid_custom_setting_config__db__avg_appls` | `custom_setting_config.db.avg_appls` |
| `ibm_db2_invalid_custom_setting_config__db__catalogcache_sz` | `custom_setting_config.db.catalogcache_sz` |
| `ibm_db2_invalid_custom_setting_config__db__chngpgs_thresh` | `custom_setting_config.db.chngpgs_thresh` |
| `ibm_db2_invalid_custom_setting_config__db__cur_commit` | `custom_setting_config.db.cur_commit` |
| `ibm_db2_invalid_custom_setting_config__db__database_memory` | `custom_setting_config.db.database_memory` |
| `ibm_db2_invalid_custom_setting_config__db__db_collname` | `custom_setting_config.db.db_collname` |
| `ibm_db2_invalid_custom_setting_config__db__db_mem_thresh` | `custom_setting_config.db.db_mem_thresh` |
| `ibm_db2_invalid_custom_setting_config__db__dbheap` | `custom_setting_config.db.dbheap` |
| `ibm_db2_invalid_custom_setting_config__db__ddl_compression_def` | `custom_setting_config.db.ddl_compression_def` |
| `ibm_db2_invalid_custom_setting_config__db__ddl_constraint_def` | `custom_setting_config.db.ddl_constraint_def` |
| `ibm_db2_invalid_custom_setting_config__db__dec_arithmetic` | `custom_setting_config.db.dec_arithmetic` |
| `ibm_db2_invalid_custom_setting_config__db__dec_to_char_fmt` | `custom_setting_config.db.dec_to_char_fmt` |
| `ibm_db2_invalid_custom_setting_config__db__decflt_rounding` | `custom_setting_config.db.decflt_rounding` |
| `ibm_db2_invalid_custom_setting_config__db__dft_degree` | `custom_setting_config.db.dft_degree` |
| `ibm_db2_invalid_custom_setting_config__db__dft_extent_sz` | `custom_setting_config.db.dft_extent_sz` |
| `ibm_db2_invalid_custom_setting_config__db__dft_loadrec_ses` | `custom_setting_config.db.dft_loadrec_ses` |
| `ibm_db2_invalid_custom_setting_config__db__dft_mttb_types` | `custom_setting_config.db.dft_mttb_types` |
| `ibm_db2_invalid_custom_setting_config__db__dft_prefetch_sz` | `custom_setting_config.db.dft_prefetch_sz` |
| `ibm_db2_invalid_custom_setting_config__db__dft_queryopt` | `custom_setting_config.db.dft_queryopt` |
| `ibm_db2_invalid_custom_setting_config__db__dft_refresh_age` | `custom_setting_config.db.dft_refresh_age` |
| `ibm_db2_invalid_custom_setting_config__db__dft_schemas_dcc` | `custom_setting_config.db.dft_schemas_dcc` |
| `ibm_db2_invalid_custom_setting_config__db__dft_sqlmathwarn` | `custom_setting_config.db.dft_sqlmathwarn` |
| `ibm_db2_invalid_custom_setting_config__db__dft_table_org` | `custom_setting_config.db.dft_table_org` |
| `ibm_db2_invalid_custom_setting_config__db__dlchktime` | `custom_setting_config.db.dlchktime` |
| `ibm_db2_invalid_custom_setting_config__db__enable_xmlchar` | `custom_setting_config.db.enable_xmlchar` |
| `ibm_db2_invalid_custom_setting_config__db__extended_row_sz` | `custom_setting_config.db.extended_row_sz` |
| `ibm_db2_invalid_custom_setting_config__db__groupheap_ratio` | `custom_setting_config.db.groupheap_ratio` |
| `ibm_db2_invalid_custom_setting_config__db__indexrec` | `custom_setting_config.db.indexrec` |
| `ibm_db2_invalid_custom_setting_config__db__large_aggregation` | `custom_setting_config.db.large_aggregation` |
| `ibm_db2_invalid_custom_setting_config__db__locklist` | `custom_setting_config.db.locklist` |
| `ibm_db2_invalid_custom_setting_config__db__locktimeout` | `custom_setting_config.db.locktimeout` |
| `ibm_db2_invalid_custom_setting_config__db__log_appl_info` | `custom_setting_config.db.log_appl_info` |
| `ibm_db2_invalid_custom_setting_config__db__log_ddl_stmts` | `custom_setting_config.db.log_ddl_stmts` |
| `ibm_db2_invalid_custom_setting_config__db__log_disk_cap` | `custom_setting_config.db.log_disk_cap` |
| `ibm_db2_invalid_custom_setting_config__db__logindexbuild` | `custom_setting_config.db.logindexbuild` |
| `ibm_db2_invalid_custom_setting_config__db__maxappls` | `custom_setting_config.db.maxappls` |
| `ibm_db2_invalid_custom_setting_config__db__maxfilop` | `custom_setting_config.db.maxfilop` |
| `ibm_db2_invalid_custom_setting_config__db__maxlocks` | `custom_setting_config.db.maxlocks` |
| `ibm_db2_invalid_custom_setting_config__db__min_dec_div_3` | `custom_setting_config.db.min_dec_div_3` |
| `ibm_db2_invalid_custom_setting_config__db__mon_act_metrics` | `custom_setting_config.db.mon_act_metrics` |
| `ibm_db2_invalid_custom_setting_config__db__mon_deadlock` | `custom_setting_config.db.mon_deadlock` |
| `ibm_db2_invalid_custom_setting_config__db__mon_lck_msg_lvl` | `custom_setting_config.db.mon_lck_msg_lvl` |
| `ibm_db2_invalid_custom_setting_config__db__mon_locktimeout` | `custom_setting_config.db.mon_locktimeout` |
| `ibm_db2_invalid_custom_setting_config__db__mon_lockwait` | `custom_setting_config.db.mon_lockwait` |
| `ibm_db2_invalid_custom_setting_config__db__mon_lw_thresh` | `custom_setting_config.db.mon_lw_thresh` |
| `ibm_db2_invalid_custom_setting_config__db__mon_obj_metrics` | `custom_setting_config.db.mon_obj_metrics` |
| `ibm_db2_invalid_custom_setting_config__db__mon_pkglist_sz` | `custom_setting_config.db.mon_pkglist_sz` |
| `ibm_db2_invalid_custom_setting_config__db__mon_req_metrics` | `custom_setting_config.db.mon_req_metrics` |
| `ibm_db2_invalid_custom_setting_config__db__mon_rtn_data` | `custom_setting_config.db.mon_rtn_data` |
| `ibm_db2_invalid_custom_setting_config__db__mon_rtn_execlist` | `custom_setting_config.db.mon_rtn_execlist` |
| `ibm_db2_invalid_custom_setting_config__db__mon_uow_data` | `custom_setting_config.db.mon_uow_data` |
| `ibm_db2_invalid_custom_setting_config__db__mon_uow_execlist` | `custom_setting_config.db.mon_uow_execlist` |
| `ibm_db2_invalid_custom_setting_config__db__mon_uow_pkglist` | `custom_setting_config.db.mon_uow_pkglist` |
| `ibm_db2_invalid_custom_setting_config__db__nchar_mapping` | `custom_setting_config.db.nchar_mapping` |
| `ibm_db2_invalid_custom_setting_config__db__num_freqvalues` | `custom_setting_config.db.num_freqvalues` |
| `ibm_db2_invalid_custom_setting_config__db__num_iocleaners` | `custom_setting_config.db.num_iocleaners` |
| `ibm_db2_invalid_custom_setting_config__db__num_ioservers` | `custom_setting_config.db.num_ioservers` |
| `ibm_db2_invalid_custom_setting_config__db__num_log_span` | `custom_setting_config.db.num_log_span` |
| `ibm_db2_invalid_custom_setting_config__db__num_quantiles` | `custom_setting_config.db.num_quantiles` |
| `ibm_db2_invalid_custom_setting_config__db__opt_buffpage` | `custom_setting_config.db.opt_buffpage` |
| `ibm_db2_invalid_custom_setting_config__db__opt_direct_wrkld` | `custom_setting_config.db.opt_direct_wrkld` |
| `ibm_db2_invalid_custom_setting_config__db__opt_locklist` | `custom_setting_config.db.opt_locklist` |
| `ibm_db2_invalid_custom_setting_config__db__opt_maxlocks` | `custom_setting_config.db.opt_maxlocks` |
| `ibm_db2_invalid_custom_setting_config__db__opt_sortheap` | `custom_setting_config.db.opt_sortheap` |
| `ibm_db2_invalid_custom_setting_config__db__page_age_trgt_gcr` | `custom_setting_config.db.page_age_trgt_gcr` |
| `ibm_db2_invalid_custom_setting_config__db__page_age_trgt_mcr` | `custom_setting_config.db.page_age_trgt_mcr` |
| `ibm_db2_invalid_custom_setting_config__db__pckcachesz` | `custom_setting_config.db.pckcachesz` |
| `ibm_db2_invalid_custom_setting_config__db__pl_stack_trace` | `custom_setting_config.db.pl_stack_trace` |
| `ibm_db2_invalid_custom_setting_config__db__self_tuning_mem` | `custom_setting_config.db.self_tuning_mem` |
| `ibm_db2_invalid_custom_setting_config__db__seqdetect` | `custom_setting_config.db.seqdetect` |
| `ibm_db2_invalid_custom_setting_config__db__sheapthres_shr` | `custom_setting_config.db.sheapthres_shr` |
| `ibm_db2_invalid_custom_setting_config__db__softmax` | `custom_setting_config.db.softmax` |
| `ibm_db2_invalid_custom_setting_config__db__sortheap` | `custom_setting_config.db.sortheap` |
| `ibm_db2_invalid_custom_setting_config__db__sql_ccflags` | `custom_setting_config.db.sql_ccflags` |
| `ibm_db2_invalid_custom_setting_config__db__stat_heap_sz` | `custom_setting_config.db.stat_heap_sz` |
| `ibm_db2_invalid_custom_setting_config__db__stmt_conc` | `custom_setting_config.db.stmt_conc` |
| `ibm_db2_invalid_custom_setting_config__db__stmtheap` | `custom_setting_config.db.stmtheap` |
| `ibm_db2_invalid_custom_setting_config__db__string_units` | `custom_setting_config.db.string_units` |
| `ibm_db2_invalid_custom_setting_config__db__systime_period_adj` | `custom_setting_config.db.systime_period_adj` |
| `ibm_db2_invalid_custom_setting_config__db__trackmod` | `custom_setting_config.db.trackmod` |
| `ibm_db2_invalid_custom_setting_config__db__util_heap_sz` | `custom_setting_config.db.util_heap_sz` |
| `ibm_db2_invalid_custom_setting_config__db__wlm_admission_ctrl` | `custom_setting_config.db.wlm_admission_ctrl` |
| `ibm_db2_invalid_custom_setting_config__db__wlm_agent_load_trgt` | `custom_setting_config.db.wlm_agent_load_trgt` |
| `ibm_db2_invalid_custom_setting_config__db__wlm_cpu_limit` | `custom_setting_config.db.wlm_cpu_limit` |
| `ibm_db2_invalid_custom_setting_config__db__wlm_cpu_share_mode` | `custom_setting_config.db.wlm_cpu_share_mode` |
| `ibm_db2_invalid_custom_setting_config__db__wlm_cpu_shares` | `custom_setting_config.db.wlm_cpu_shares` |
| `ibm_db2_invalid_custom_setting_config__dbm__comm_bandwidth` | `custom_setting_config.dbm.comm_bandwidth` |
| `ibm_db2_invalid_custom_setting_config__dbm__cpuspeed` | `custom_setting_config.dbm.cpuspeed` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_bufpool` | `custom_setting_config.dbm.dft_mon_bufpool` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_lock` | `custom_setting_config.dbm.dft_mon_lock` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_sort` | `custom_setting_config.dbm.dft_mon_sort` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_stmt` | `custom_setting_config.dbm.dft_mon_stmt` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_table` | `custom_setting_config.dbm.dft_mon_table` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_timestamp` | `custom_setting_config.dbm.dft_mon_timestamp` |
| `ibm_db2_invalid_custom_setting_config__dbm__dft_mon_uow` | `custom_setting_config.dbm.dft_mon_uow` |
| `ibm_db2_invalid_custom_setting_config__dbm__diaglevel` | `custom_setting_config.dbm.diaglevel` |
| `ibm_db2_invalid_custom_setting_config__dbm__federated_async` | `custom_setting_config.dbm.federated_async` |
| `ibm_db2_invalid_custom_setting_config__dbm__indexrec` | `custom_setting_config.dbm.indexrec` |
| `ibm_db2_invalid_custom_setting_config__dbm__intra_parallel` | `custom_setting_config.dbm.intra_parallel` |
| `ibm_db2_invalid_custom_setting_config__dbm__keepfenced` | `custom_setting_config.dbm.keepfenced` |
| `ibm_db2_invalid_custom_setting_config__dbm__max_connretries` | `custom_setting_config.dbm.max_connretries` |
| `ibm_db2_invalid_custom_setting_config__dbm__max_querydegree` | `custom_setting_config.dbm.max_querydegree` |
| `ibm_db2_invalid_custom_setting_config__dbm__mon_heap_sz` | `custom_setting_config.dbm.mon_heap_sz` |
| `ibm_db2_invalid_custom_setting_config__dbm__multipartsizemb` | `custom_setting_config.dbm.multipartsizemb` |
| `ibm_db2_invalid_custom_setting_config__dbm__notifylevel` | `custom_setting_config.dbm.notifylevel` |
| `ibm_db2_invalid_custom_setting_config__dbm__num_initagents` | `custom_setting_config.dbm.num_initagents` |
| `ibm_db2_invalid_custom_setting_config__dbm__num_initfenced` | `custom_setting_config.dbm.num_initfenced` |
| `ibm_db2_invalid_custom_setting_config__dbm__num_poolagents` | `custom_setting_config.dbm.num_poolagents` |
| `ibm_db2_invalid_custom_setting_config__dbm__resync_interval` | `custom_setting_config.dbm.resync_interval` |
| `ibm_db2_invalid_custom_setting_config__dbm__rqrioblk` | `custom_setting_config.dbm.rqrioblk` |
| `ibm_db2_invalid_custom_setting_config__dbm__start_stop_time` | `custom_setting_config.dbm.start_stop_time` |
| `ibm_db2_invalid_custom_setting_config__dbm__util_impact_lim` | `custom_setting_config.dbm.util_impact_lim` |
| `ibm_db2_invalid_custom_setting_config__dbm__wlm_disp_concur` | `custom_setting_config.dbm.wlm_disp_concur` |
| `ibm_db2_invalid_custom_setting_config__dbm__wlm_disp_cpu_shares` | `custom_setting_config.dbm.wlm_disp_cpu_shares` |
| `ibm_db2_invalid_custom_setting_config__dbm__wlm_disp_min_util` | `custom_setting_config.dbm.wlm_disp_min_util` |
| `ibm_db2_invalid_custom_setting_config__dbm__wlm_dispatcher` | `custom_setting_config.dbm.wlm_dispatcher` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_alternate_authz_behaviour` | `custom_setting_config.registry.db2_alternate_authz_behaviour` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_antijoin` | `custom_setting_config.registry.db2_antijoin` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_ats_enable` | `custom_setting_config.registry.db2_ats_enable` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_bidi` | `custom_setting_config.registry.db2_bidi` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_compopt` | `custom_setting_config.registry.db2_compopt` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_deferred_prepare_semantics` | `custom_setting_config.registry.db2_deferred_prepare_semantics` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_evaluncommitted` | `custom_setting_config.registry.db2_evaluncommitted` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_extended_optimization` | `custom_setting_config.registry.db2_extended_optimization` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_index_pctfree_default` | `custom_setting_config.registry.db2_index_pctfree_default` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_inlist_to_nljn` | `custom_setting_config.registry.db2_inlist_to_nljn` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_lock_to_rb` | `custom_setting_config.registry.db2_lock_to_rb` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_minimize_listprefetch` | `custom_setting_config.registry.db2_minimize_listprefetch` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_object_table_entries` | `custom_setting_config.registry.db2_object_table_entries` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_opt_max_temp_size` | `custom_setting_config.registry.db2_opt_max_temp_size` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_optprofile` | `custom_setting_config.registry.db2_optprofile` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_optstats_log` | `custom_setting_config.registry.db2_optstats_log` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_parallel_io` | `custom_setting_config.registry.db2_parallel_io` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_reduced_optimization` | `custom_setting_config.registry.db2_reduced_optimization` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_selectivity` | `custom_setting_config.registry.db2_selectivity` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_skipdeleted` | `custom_setting_config.registry.db2_skipdeleted` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_skipinserted` | `custom_setting_config.registry.db2_skipinserted` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_stmm` | `custom_setting_config.registry.db2_stmm` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_sync_release_lock_attributes` | `custom_setting_config.registry.db2_sync_release_lock_attributes` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_truncate_reusestorage` | `custom_setting_config.registry.db2_truncate_reusestorage` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_use_alternate_page_cleaning` | `custom_setting_config.registry.db2_use_alternate_page_cleaning` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_view_reopt_values` | `custom_setting_config.registry.db2_view_reopt_values` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_wlm_settings` | `custom_setting_config.registry.db2_wlm_settings` |
| `ibm_db2_invalid_custom_setting_config__registry__db2_workload` | `custom_setting_config.registry.db2_workload` |
| `ibm_db2_invalid_users_config__email` | `users_config.email` |
| `ibm_db2_invalid_users_config__iam` | `users_config.iam` |
| `ibm_db2_invalid_users_config__ibmid` | `users_config.ibmid` |
| `ibm_db2_invalid_users_config__id` | `users_config.id` |
| `ibm_db2_invalid_users_config__locked` | `users_config.locked` |
| `ibm_db2_invalid_users_config__name` | `users_config.name` |
| `ibm_db2_invalid_users_config__password` | `users_config.password` |
| `ibm_db2_invalid_users_config__role` | `users_config.role` |
| `ibm_db2_invalid_users_config__authentication__method` | `users_config.authentication.method` |
| `ibm_db2_invalid_users_config__authentication__policy_id` | `users_config.authentication.policy_id` |

## Example

### Valid

```hcl
resource "ibm_db2" "test" {
  location = "example"
  name = "example"
  plan = "example"
  service = "example"
}
```
