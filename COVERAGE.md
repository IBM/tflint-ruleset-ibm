# Coverage

Ground truth is the IBM Cloud Terraform provider `ResourcesMap` and `DataSourcesMap`.

| | Count |
| --- | ---: |
| Service packages | 58 |
| Provider resources discovered | 636 |
| Resources covered | 636 |
| Resources skipped | 0 |
| Provider data sources discovered | 865 |
| Data sources covered | 865 |
| Data sources skipped | 0 |
| Generated rules | 1501 |
| Attribute rules | 10533 |
| Registered checks | 12034 |
| Empty-schema stubs (fail fixture n/a) | 9 |
| Coverage gaps | 0 |
| Stale generated files | 0 |

No skipped constructors. Every provider map entry has extracted schema, rule, test, pass fixture, fail fixture (or empty-schema waiver), registration, and docs.

## accountmanagement

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `accountmanagement` | `ibm_account_info` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## appconfiguration

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `appconfiguration` | `ibm_app_config_collection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_environment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_feature` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_integration_en` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_integration_kms` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_property` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_segment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_collection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_collections` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_environment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_environments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_feature` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_features` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_integration_en` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_integration_kms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_integrations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_properties` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_property` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_segment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_segments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_snapshot` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appconfiguration` | `ibm_app_config_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## appid

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `appid` | `ibm_appid_action_url` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_apm` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application_roles` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application_scopes` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_audit_status` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_cloud_directory_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_cloud_directory_user` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_cloud_directory` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_custom` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_facebook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_google` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_saml` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_languages` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_mfa` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_mfa_channel` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_password_regex` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_redirect_urls` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_role` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_theme_color` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_theme_text` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_token_config` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_user_roles` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_action_url` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_apm` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application_roles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_application_scopes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_applications` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_audit_status` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_cloud_directory_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_cloud_directory_user` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_cloud_directory` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_custom` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_facebook` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_google` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_saml` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_idp_saml_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_languages` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_mfa` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_mfa_channel` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_password_regex` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_redirect_urls` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_role` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_roles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_theme_color` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_theme_text` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_token_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `appid` | `ibm_appid_user_roles` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## atracker

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `atracker` | `ibm_atracker_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `atracker` | `ibm_atracker_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `atracker` | `ibm_atracker_target` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `atracker` | `ibm_atracker_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `atracker` | `ibm_atracker_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## backuprecovery

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `backuprecovery` | `ibm_backup_recovery` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_agent_upgrade_task` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connection_registration_token` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_access_token` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_agent_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_update_user` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_data_source_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_data_source_connector_patch` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_download_files_folders` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_cancel_cluster_upgrades` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_create_cluster_upgrades` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_update_cluster_upgrades` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_perform_action_on_protection_group_run_request` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_group_run_request` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_source_refresh` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_restore_points` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_source_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_update_protection_group_run_request` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recoveries` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_agent_upgrade_tasks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_agent_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_agents` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_get_users` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_logs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connector_status` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_connectors_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_data_source_connections` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_data_source_connectors` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_download_agent` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_download_files` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_download_indexed_files` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_alerts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_alerts_resolution` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_alerts_stats` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_alerts_summary` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_cluster_info` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_compatible_clusters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_management_alerts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_management_alerts_summary` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_manager_get_upgrades_info` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_object_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_group_runs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_protection_sources` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_registration_info` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_search_indexed_object` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_search_objects` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_search_protected_objects` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_source_registration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `backuprecovery` | `ibm_backup_recovery_source_registrations` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## catalogmanagement

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `catalogmanagement` | `ibm_cm_account` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_catalog` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_object` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_offering` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_offering_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_validation` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_account` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_catalog` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_object` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_offering` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_offering_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_preset` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `catalogmanagement` | `ibm_cm_version` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cdtektonpipeline

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_definition` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_property` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_trigger` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_trigger_property` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_definition` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_property` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_trigger` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtektonpipeline` | `ibm_cd_tekton_pipeline_trigger_property` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cdtoolchain

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cdtoolchain` | `ibm_cd_toolchain` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_appconfig` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_artifactory` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_bitbucketgit` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_cos` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_custom` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_devopsinsights` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_eventnotifications` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_githubconsolidated` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_gitlab` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_hashicorpvault` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_hostedgit` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_jenkins` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_jira` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_keyprotect` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_nexus` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_pagerduty` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_pipeline` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_privateworker` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_saucelabs` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_secretsmanager` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_securitycompliance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_slack` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_sonarqube` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_appconfig` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_artifactory` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_bitbucketgit` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_cos` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_custom` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_devopsinsights` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_eventnotifications` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_githubconsolidated` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_gitlab` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_hashicorpvault` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_hostedgit` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_jenkins` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_jira` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_keyprotect` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_nexus` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_pagerduty` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_pipeline` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_privateworker` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_saucelabs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_secretsmanager` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_securitycompliance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_slack` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchain_tool_sonarqube` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cdtoolchain` | `ibm_cd_toolchains` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cis

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cis` | `ibm_cis` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_advanced_certificate_pack_order` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_alert` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_bot_management` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_cache_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_certificate_order` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_certificate_upload` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_list` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_list_items` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_page` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_dns_record` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_dns_records_batch` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_dns_records_import` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_domain` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_domain_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_edge_functions_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_edge_functions_trigger` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_filter` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_firewall` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_firewall_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_global_load_balancer` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_healthcheck` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_logpush_job` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_mtls` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_mtls_app` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_auth` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_certificate_order` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_page_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_range_app` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_rate_limit` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_routing` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_entrypoint_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_version_detach` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_tls_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_package` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_webhook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_alerts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_bot_analytics` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_bot_managements` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_cache_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_certificates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_certificates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_list_items` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_lists` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_custom_pages` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_dns_records` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_domain` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_edge_functions_actions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_edge_functions_triggers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_filters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_firewall` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_firewall_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_global_load_balancers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_healthchecks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ip_addresses` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_logpush_jobs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_managed_lists` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_mtls_apps` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_mtlss` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_auths` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_certificates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_origin_pools` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_page_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_range_apps` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_rate_limit` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_entrypoint_versions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_rules_by_tag` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_ruleset_versions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_rulesets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_packages` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_waf_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cis` | `ibm_cis_webhooks` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## classicinfrastructure

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `classicinfrastructure` | `ibm_cdn` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_autoscale_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_autoscale_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_bare_metal` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_dedicated_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_monitor` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_placement_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_provisioning_hook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_reserved_capacity` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_ssh_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_ssl_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_user` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_vm_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_domain` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_domain_registration_nameservers` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_record` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_reverse_record` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_secondary` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_firewall` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_firewall_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_hardware_firewall_shared` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_ipsec_vpn` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_service` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_service_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_vpx` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_vpx_ha` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_vpx_service` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lb_vpx_vip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lbaas` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lbaas_health_monitor` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lbaas_server_instance_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_multi_vlan_firewall` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_gateway_vlan_association` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_interface_sg_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_public_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_vlan` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_vlan_spanning` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_object_storage_account` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_security_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_security_group_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_ssl_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_storage_block` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_storage_evault` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_storage_file` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_subnet` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_bare_metal` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_image_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_placement_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_reserved_capacity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_ssh_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_compute_vm_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_domain` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_domain_registration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_dns_secondary` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_lbaas` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_network_vlan` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `classicinfrastructure` | `ibm_security_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cloudant

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cloudant` | `ibm_cloudant` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudant` | `ibm_cloudant_database` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudant` | `ibm_cloudant` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudant` | `ibm_cloudant_database` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cloudfoundry

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cloudfoundry` | `ibm_app` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_domain_private` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_domain_shared` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_org` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_service_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_service_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_space` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_account` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_domain_private` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_domain_shared` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_app_route` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_org` | data | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `cloudfoundry` | `ibm_org_quota` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_service_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_service_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_service_plan` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudfoundry` | `ibm_space` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cloudshell

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cloudshell` | `ibm_cloud_shell_account_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cloudshell` | `ibm_cloud_shell_account_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## codeengine

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `codeengine` | `ibm_code_engine_allowed_outbound_destination` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_app` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_binding` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_build` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_config_map` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_domain_mapping` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_function` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_job` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_persistent_data_store` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_project` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_allowed_outbound_destination` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_app` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_binding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_build` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_config_map` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_domain_mapping` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_function` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_job` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_persistent_data_store` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_project` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `codeengine` | `ibm_code_engine_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## configurationaggregator

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `configurationaggregator` | `ibm_config_aggregator_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `configurationaggregator` | `ibm_config_aggregator_configurations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `configurationaggregator` | `ibm_config_aggregator_resource_collection_status` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `configurationaggregator` | `ibm_config_aggregator_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## contextbasedrestrictions

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `contextbasedrestrictions` | `ibm_cbr_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `contextbasedrestrictions` | `ibm_cbr_zone` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `contextbasedrestrictions` | `ibm_cbr_zone_addresses` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `contextbasedrestrictions` | `ibm_cbr_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `contextbasedrestrictions` | `ibm_cbr_zone` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `contextbasedrestrictions` | `ibm_cbr_zone_addresses` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## cos

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `cos` | `ibm_cos_backup_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_backup_vault` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_lifecycle_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_object` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_object_lock_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_replication_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_website_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_backup_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_backup_vault` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `cos` | `ibm_cos_bucket_object` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## database

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `database` | `ibm_database` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_backup` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_backups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_connection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_point_in_time_recovery` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_remotes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_task` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `database` | `ibm_database_tasks` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## db2

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `db2` | `ibm_db2` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_allowlist_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_autoscale` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_backup` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_connection_info` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_tuneable_param` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_users` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `db2` | `ibm_db2_whitelist_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## directlink

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `directlink` | `ibm_dl_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_macsec_cak` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_macsec_config` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_provider_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_route_report` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_virtual_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_export_route_filter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_export_route_filters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_macsec_cak` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_macsec_caks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateway_macsec_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_import_route_filter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_import_route_filters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_offering_speeds` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_port` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_ports` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_provider_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_provider_ports` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_route_report` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_route_reports` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `directlink` | `ibm_dl_routers` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## dnsservices

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `dnsservices` | `ibm_dns_custom_resolver` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_custom_resolver_forwarding_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_custom_resolver_secondary_zone` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glb` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glb_monitor` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glb_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_linked_zone` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_permitted_network` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_resource_record` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_zone` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_custom_resolver_forwarding_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_custom_resolver_secondary_zones` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_custom_resolvers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glb_monitors` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glb_pools` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_glbs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_permitted_networks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_resource_records` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `dnsservices` | `ibm_dns_zones` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## drautomationservice

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `drautomationservice` | `ibm_pdr_managedr` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_validate_apikey` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_dr_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_dr_summary_response` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_event` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_events` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_dr_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_dr_summary_response` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_event` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_events` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_grs_location_pairs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_machine_types` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_managed_vm_list` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_get_powervs_workspace` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_grs_location_pairs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_last_operation` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_machine_types` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_managed_vm_list` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `drautomationservice` | `ibm_pdr_powervs_workspaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## enterprise

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `enterprise` | `ibm_enterprise` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `enterprise` | `ibm_enterprise_account` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `enterprise` | `ibm_enterprise_account_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `enterprise` | `ibm_enterprise_account_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `enterprise` | `ibm_enterprise_accounts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `enterprise` | `ibm_enterprises` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## eventnotification

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `eventnotification` | `ibm_en_app_configuration_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_code_engine_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_android` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_app_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_ce` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_chrome` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_cos` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_custom_email` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_custom_sms` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_event_streams` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_firefox` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_huawei` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_ios` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_msteams` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_pagerduty` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_safari` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_slack` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_sn` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_webhook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_email_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_event_streams_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_ibmsource` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_integration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_integration_cos` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_pagerduty_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_slack_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_setting` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_user` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_source` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_android` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_app_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_ce` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_chrome` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_cos` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_custom_email` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_custom_sms` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_email` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_event_streams` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_firefox` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_huawei` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_ios` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_msteams` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_pagerduty` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_safari` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_scheduler` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_slack` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_sms` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_sn` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_webhook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_topic` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_webhook_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_app_configuration_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_bounce_metrics` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_code_engine_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_android` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_app_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_ce` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_chrome` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_cos` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_custom_email` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_custom_sms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_event_streams` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_firefox` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_huawei` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_ios` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_msteams` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_pagerduty` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_safari` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_slack` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_sn` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destination_webhook` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_destinations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_email_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_email_templates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_event_streams_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_integration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_integration_cos` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_integrations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_metrics` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_pagerduty_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_pre_defined_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_pre_defined_templates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_slack_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_allowed_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_configurations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_user` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_smtp_users` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_source` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_sources` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_android` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_app_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_ce` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_chrome` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_cos` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_custom_email` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_custom_sms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_email` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_event_streams` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_firefox` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_huawei` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_ios` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_msteams` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_pagerduty` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_safari` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_scheduler` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_slack` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_sms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_sn` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscription_webhook` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_subscriptions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_topic` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_topics` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventnotification` | `ibm_en_webhook_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## eventstreams

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `eventstreams` | `ibm_event_streams_mirroring_config` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_quota` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_schema` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_schema_global_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_topic` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_mirroring_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_quota` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_schema` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_schema_global_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `eventstreams` | `ibm_event_streams_topic` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## functions

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `functions` | `ibm_function_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_namespace` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_package` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_trigger` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_action` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_namespace` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_package` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `functions` | `ibm_function_trigger` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## globaltagging

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `globaltagging` | `ibm_iam_access_tag` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `globaltagging` | `ibm_resource_tag` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `globaltagging` | `ibm_iam_access_tag` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `globaltagging` | `ibm_resource_tag` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## hpcs

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `hpcs` | `ibm_hpcs` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_key_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_keystore` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_managed_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_vault` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_key_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_keystore` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_managed_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `hpcs` | `ibm_hpcs_vault` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## iamaccessgroup

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `iamaccessgroup` | `ibm_iam_access_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_account_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_dynamic_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_members` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_template_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_template_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_template_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamaccessgroup` | `ibm_iam_access_group_template_versions` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## iamidentity

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `iamidentity` | `ibm_iam_account_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_account_settings_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_account_settings_template_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_api_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_identity_preference` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_idp` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_idp_account_setting` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_service_api_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_service_id` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_serviceid_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_claim_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_identities` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_identity` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_link` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_template_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_user_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_account_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_account_settings_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_account_settings_template_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_api_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_auth_token` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_effective_account_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_identity_preference` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_identity_preferences` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_idp` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_idp_account_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_idps` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_service_id` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_serviceid_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_claim_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_claim_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_identities` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_identity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_link` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_links` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profile_template_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_trusted_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_user_mfa_enrollments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_user_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iamidentity` | `ibm_iam_users` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## iampolicy

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `iampolicy` | `ibm_iam_access_group_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_account_settings_external_interaction` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_template_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_authorization_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_authorization_policy_detach` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_custom_role` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_template_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_template_version` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_service_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_trusted_profile_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_user_invite` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_user_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_access_group_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_account_settings_external_interaction` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_assignments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_action_control_template_version` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_authorization_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_assignments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_policy_template_version` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_actions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_assignments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_role_template_version` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_roles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_service_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_trusted_profile_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `iampolicy` | `ibm_iam_user_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## kms

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `kms` | `ibm_kms_cryptounits` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_instance_policies` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_alias` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_policies` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_rings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_with_policy_overrides` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_adapter` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_client_cert` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kp_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_instance_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_key_rings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_keys` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_adapter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_adapters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_client_cert` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_client_certs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_object` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kms_kmip_objects` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kms` | `ibm_kp_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## kubernetes

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `kubernetes` | `ibm_container_addons` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_alb` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_alb_cert` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_alb_create` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_api_key_reset` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_bind_service` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster_feature` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_secret_opaque` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_secret_tls` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_nlb_dns` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_storage_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vni_baremetal_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_alb` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_alb_create` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_cluster` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_worker` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_worker_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_worker_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_worker_pool_zone_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_addons` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_alb` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_alb_cert` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_bind_service` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster_versions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_cluster_worker` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host_flavor` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host_flavors` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_dedicated_host_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_secret_opaque` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_ingress_secret_tls` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_nlb_dns` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_storage_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vni_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vni_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_alb` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_cluster` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_cluster_alb` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_cluster_worker` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_cluster_worker_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_vpc_worker_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `kubernetes` | `ibm_container_worker_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## logs

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `logs` | `ibm_logs_alert` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_alert_definition` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_dashboard` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_dashboard_folder` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_data_access_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_data_usage_metrics` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_e2m` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_enrichment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_extension_deployment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_log_data_retention_tags` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_outgoing_webhook` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_rule_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_stream` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_view` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_view_folder` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_alert` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_alert_definition` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_alert_definitions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_alerts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_dashboard` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_dashboard_folders` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_data_access_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_data_usage_metrics` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_e2m` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_e2ms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_enrichments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_extension` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_extension_deployment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_extensions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_log_data_retention_tags` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_outgoing_webhook` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_outgoing_webhooks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_rule_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_rule_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_stream` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_streams` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_view` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_view_folder` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_view_folders` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logs` | `ibm_logs_views` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## logsrouter

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `logsrouter` | `ibm_logs_router_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logsrouter` | `ibm_logs_router_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logsrouter` | `ibm_logs_router_target` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logsrouter` | `ibm_logs_router_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `logsrouter` | `ibm_logs_router_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## logsrouting

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `logsrouting` | `ibm_logs_router_tenant` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `logsrouting` | `ibm_logs_router_tenants` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## metricsrouter

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `metricsrouter` | `ibm_metrics_router_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `metricsrouter` | `ibm_metrics_router_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `metricsrouter` | `ibm_metrics_router_target` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `metricsrouter` | `ibm_metrics_router_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `metricsrouter` | `ibm_metrics_router_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## mqcloud

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `mqcloud` | `ibm_mqcloud_application` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_keystore_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_queue_manager` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_truststore_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_user` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_virtual_private_endpoint_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_application` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_keystore_certificate` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_queue_manager` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_queue_manager_options` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_queue_manager_status` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_truststore_certificate` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_user` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_virtual_private_endpoint_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `mqcloud` | `ibm_mqcloud_virtual_private_endpoint_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## pag

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `pag` | `ibm_pag_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `pag` | `ibm_pag_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## partnercentersell

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `partnercentersell` | `ibm_onboarding_catalog_deployment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_catalog_plan` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_catalog_product` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_iam_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_product` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `partnercentersell` | `ibm_onboarding_resource_broker` | resource | yes | yes | yes | yes | yes | yes | yes | covered |

## platformnotifications

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `platformnotifications` | `ibm_notification_distribution_list_destination` | resource | yes | yes | yes | yes | yes | yes | yes | covered |

## power

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `power` | `ibm_pi_capture` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_cloud_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_cloud_connection_network_attach` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_console_language` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_dhcp` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_host_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_ike_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_image` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_image_export` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_vpmem_volumes` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_ipsec_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_address_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_address_group_member` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_interface` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peer` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peer_route_filter` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_port_attach` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_group_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_group_member` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_group_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_placement_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_shared_processor_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_spp_placement_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_virtual_serial_number` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_attach` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_clone` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_onboarding` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_vpn_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_workspace` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_available_hosts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_catalog_images` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_cloud_connection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_cloud_connections` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_cloud_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_console_languages` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_datacenter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_datacenters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_dhcp` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_dhcps` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_disaster_recovery_location` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_disaster_recovery_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_host` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_host_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_host_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_hosts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_image` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_images` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_network` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_networks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_snapshot` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_volumes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_vpmem_volume` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instance_vpmem_volumes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_instances` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_keys` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_address_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_address_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_interface` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peer` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peer_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peer_route_filter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_peers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_port` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_network_security_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_networks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_placement_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_placement_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_public_network` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_pvm_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_route` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_route_report` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_sap_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_sap_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_shared_processor_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_shared_processor_pools` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_software_tiers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_spp_placement_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_spp_placement_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_storage_pool_capacity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_storage_pools_capacity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_storage_tiers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_storage_type_capacity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_storage_types_capacity` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_system_pools` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_tenant` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_virtual_serial_number` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_virtual_serial_numbers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_clone` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_flash_copy_mappings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group_details` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group_remote_copy_relationships` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_group_storage_details` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_groups_details` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_onboarding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_onboardings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_remote_copy_relationship` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_snapshot` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volume_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_volumes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_workspace` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `power` | `ibm_pi_workspaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## powerhaautomationservice

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `powerhaautomationservice` | `ibm_pha_api_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_cluster_nodes` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_deployment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_cluster_nodes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_deployment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_last_operation` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_powervs_workspaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `powerhaautomationservice` | `ibm_pha_supported_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## project

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `project` | `ibm_project` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `project` | `ibm_project_config` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `project` | `ibm_project_environment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `project` | `ibm_project` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `project` | `ibm_project_config` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `project` | `ibm_project_environment` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## pushnotification

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `pushnotification` | `ibm_pn_application_chrome` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `pushnotification` | `ibm_pn_application_chrome` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## registry

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `registry` | `ibm_cr_namespace` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `registry` | `ibm_cr_retention_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `registry` | `ibm_cr_namespaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## resourcecontroller

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `resourcecontroller` | `ibm_resource_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_reclamation_delete` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_quota` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcecontroller` | `ibm_resource_reclamations` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## resourcemanager

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `resourcemanager` | `ibm_resource_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcemanager` | `ibm_resource_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `resourcemanager` | `ibm_resource_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## satellite

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `satellite` | `ibm_satellite_cluster` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_cluster_worker_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_cluster_worker_pool_zone_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_endpoint` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_link` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_location` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_location_nlb_dns` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_storage_assignment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_storage_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_attach_host_script` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_cluster` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_cluster_worker_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_cluster_worker_pool_zone_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_endpoint` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_link` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_location` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_location_nlb_dns` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_storage_assignment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `satellite` | `ibm_satellite_storage_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## scc

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `scc` | `ibm_scc_account_settings` | resource | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_control_library` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_instance_settings` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_profile` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_profile_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_provider_type_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_rule_attachment` | resource | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_scope` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_template` | resource | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_template_attachment` | resource | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_account_location` | data | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_account_location_settings` | data | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_account_locations` | data | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_account_notification_settings` | data | yes | yes | yes | yes | n/a | yes | yes | covered-empty-schema |
| `scc` | `ibm_scc_control_libraries` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_control_library` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_instance_settings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_latest_reports` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_profile_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_provider_type` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_provider_type_collection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_provider_type_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_provider_types` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_controls` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_evaluations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_resources` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_summary` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_tags` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_report_violation_drift` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_scope` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `scc` | `ibm_scc_scope_collection` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## schematics

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `schematics` | `ibm_schematics_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_deploy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_health` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_prs` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_inventory` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_job` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_resource_query` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_workspace` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_action` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_deploy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_health` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agent_prs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_agents` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_inventory` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_job` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_kms` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_output` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_resource_query` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_state` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `schematics` | `ibm_schematics_workspace` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## sdsaas

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `sdsaas` | `ibm_sds_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `sdsaas` | `ibm_sds_volume` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `sdsaas` | `ibm_sds_volume_mapping` | resource | yes | yes | yes | yes | yes | yes | yes | covered |

## secretsmanager

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `secretsmanager` | `ibm_sm_arbitrary_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_custom_credentials_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_custom_credentials_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_en_registration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_iam_credentials_configuration` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_iam_credentials_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_imported_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_kv_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_action_set_signed` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_action_sign_csr` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_intermediate_ca` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_root_ca` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_action_validate_manual_dns` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_ca_lets_encrypt` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_dns_cis` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_dns_classic_infrastructure` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_secret_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_service_credentials_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_username_password_secret` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_arbitrary_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_arbitrary_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_configurations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_custom_credentials_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_custom_credentials_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_custom_credentials_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_en_registration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_iam_credentials_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_iam_credentials_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_iam_credentials_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_imported_certificate` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_imported_certificate_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_kv_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_kv_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_intermediate_ca` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_root_ca` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_configuration_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_private_certificate_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_ca_lets_encrypt` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_dns_cis` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_configuration_dns_classic_infrastructure` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_public_certificate_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_secret_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_secret_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_secrets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_service_credentials_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_service_credentials_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_username_password_secret` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanager` | `ibm_sm_username_password_secret_metadata` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## secretsmanagerinstancemanagement

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `secretsmanagerinstancemanagement` | `ibm_sm_admin_token` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `secretsmanagerinstancemanagement` | `ibm_sm_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## transitgateway

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `transitgateway` | `ibm_tg_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_connection_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_connection_prefix_filter` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_connection_rgre_tunnel` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_route_report` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_connection_prefix_filter` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_connection_prefix_filters` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_location` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_locations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_route_report` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `transitgateway` | `ibm_tg_route_reports` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## usagereports

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `usagereports` | `ibm_billing_report_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `usagereports` | `ibm_billing_snapshot_list` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## vmware

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `vmware` | `ibm_vmaas_transit_gateway_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vmware` | `ibm_vmaas_vdc` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vmware` | `ibm_vmaas_transit_gateway_connection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vmware` | `ibm_vmaas_vdc` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## vpc

| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `vpc` | `ibm_is_backup_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy_plan` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_disk` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_initialization` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_allow_float` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_floating_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_interface` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnet` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnet_reserved_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_disk_management` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_floating_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_flow_log` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ike_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_deprecate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_export_job` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_obsolete` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_cluster_network_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_disk_management` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_action` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_membership` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interface` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interface_floating_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_reinitialize` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_software_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_template` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_volume_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ipsec_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policy_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pool` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pool_member` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acl` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acl_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_placement_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_account_policy` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_operations` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_revoke_account` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_address_range` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_reservation` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_reservation_activate` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_rule` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_target` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_delete_accessor_binding` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_mount_target` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_replica_operations` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_consistency_group` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ssh_key` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_network_acl_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_public_gateway_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_reserved_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_reserved_ip_patch` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_routing_table_attachment` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway_resource_binding` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_floating_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_ip` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_job` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_job_cancel` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_address_prefix` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_dns_resolution_binding` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_table` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_table_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_advertised_cidr` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_connection` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_client` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_route` | resource | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy_job` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy_jobs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy_plan` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_backup_policy_plans` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_capacities` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_disk` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_disks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_initialization` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_floating_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_floating_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_reserved_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interface_reserved_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_network_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_server_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_bare_metal_servers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_interface` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnet` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnet_reserved_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnet_reserved_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_network_subnets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_cluster_networks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_disk` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_disks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_host_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_dedicated_hosts` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_endpoint_gateway_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_floating_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_floating_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_flow_log` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_flow_logs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ike_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ike_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_bare_metal_server_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_export_job` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_export_jobs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_image_instance_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_images` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_cluster_network_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_cluster_network_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_disk` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_disks` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_action` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_actions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_manager_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_managers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_membership` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_group_memberships` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interface` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interface_reserved_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interface_reserved_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_network_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_software_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_software_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_template` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_templates` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_volume_attachment` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instance_volume_attachments` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_instances` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ipsec_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ipsec_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policy_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listener_policy_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_listeners` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pool` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pool_member` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pool_members` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_pools` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lb_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_lbs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acl` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acl_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acl_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_network_acls` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_operating_system` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_operating_systems` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_placement_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_placement_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_account_policies` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_account_policy` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_endpoint_gateway_binding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateway_endpoint_gateway_bindings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_private_path_service_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_address_range` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_address_ranges` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_public_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_region` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_regions` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_reservation` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_reservations` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_rule` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_rules` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_target` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_group_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_security_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_accessor_binding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_accessor_bindings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_mount_target` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_mount_targets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_snapshot` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_share_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_shares` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_clone` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_clones` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_consistency_group` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_consistency_groups` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshot_instance_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_snapshots` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_source_share` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ssh_key` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_ssh_keys` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_reserved_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnet_reserved_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_subnets` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway_resource_binding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateway_resource_bindings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_endpoint_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_floating_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_floating_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_ip` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interface_ips` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_virtual_network_interfaces` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_instance_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_job` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_jobs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_profile` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volume_profiles` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_volumes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_address_prefix` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_address_prefixes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_default_routing_table` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_dns_resolution_binding` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_dns_resolution_bindings` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_table` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_table_route` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_table_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpc_routing_tables` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpcs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_advertised_cidrs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_connection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_connection_local_cidrs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_connection_peer_cidrs` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_connections` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_service_connection` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateway_service_connections` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_gateways` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_client` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_client_configuration` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_clients` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_route` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_server_routes` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_vpn_servers` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_zone` | data | yes | yes | yes | yes | yes | yes | yes | covered |
| `vpc` | `ibm_is_zones` | data | yes | yes | yes | yes | yes | yes | yes | covered |

## Unregistered source files

These `resource_ibm_*` / `data_source_ibm_*` files exist under `ibm/service` but are not themselves provider map keys. They are helpers, gen2 variants, or share a constructor registered under another name. They are listed so they cannot be silently ignored.

- resource files: 629
- data source files: 851
