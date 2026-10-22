# Rules

Schema-validation rules generated from the IBM Cloud Terraform provider and IBM Cloud OpenAPI documents. Each type has a schema rule plus one rule per writable argument.

These rules check schema validity only. They do not assess security posture (for example open security groups, broad IAM policies or missing encryption), so run a policy scanner alongside them.

## accountmanagement

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_account_info_data](data_ibm_account_info.md) | data | Schema validation for the ibm_account_info data source. |

## appconfiguration

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_app_config_collection_data](data_ibm_app_config_collection.md) | data | Reads ibm_app_config_collection. |
| [ibm_app_config_collection](ibm_app_config_collection.md) | resource | Manages ibm_app_config_collection. |
| [ibm_app_config_collections_data](data_ibm_app_config_collections.md) | data | Schema validation for the ibm_app_config_collections data source. |
| [ibm_app_config_environment_data](data_ibm_app_config_environment.md) | data | Reads ibm_app_config_environment. |
| [ibm_app_config_environment](ibm_app_config_environment.md) | resource | Manages ibm_app_config_environment. |
| [ibm_app_config_environments_data](data_ibm_app_config_environments.md) | data | Schema validation for the ibm_app_config_environments data source. |
| [ibm_app_config_feature_data](data_ibm_app_config_feature.md) | data | Reads ibm_app_config_feature. |
| [ibm_app_config_feature](ibm_app_config_feature.md) | resource | Manages ibm_app_config_feature. |
| [ibm_app_config_features_data](data_ibm_app_config_features.md) | data | Schema validation for the ibm_app_config_features data source. |
| [ibm_app_config_integration_en_data](data_ibm_app_config_integration_en.md) | data | Schema validation for the ibm_app_config_integration_en data source. |
| [ibm_app_config_integration_en](ibm_app_config_integration_en.md) | resource | Schema validation for the ibm_app_config_integration_en resource. |
| [ibm_app_config_integration_kms_data](data_ibm_app_config_integration_kms.md) | data | Schema validation for the ibm_app_config_integration_kms data source. |
| [ibm_app_config_integration_kms](ibm_app_config_integration_kms.md) | resource | Schema validation for the ibm_app_config_integration_kms resource. |
| [ibm_app_config_integrations_data](data_ibm_app_config_integrations.md) | data | Schema validation for the ibm_app_config_integrations data source. |
| [ibm_app_config_properties_data](data_ibm_app_config_properties.md) | data | Schema validation for the ibm_app_config_properties data source. |
| [ibm_app_config_property_data](data_ibm_app_config_property.md) | data | Reads ibm_app_config_property. |
| [ibm_app_config_property](ibm_app_config_property.md) | resource | Manages ibm_app_config_property. |
| [ibm_app_config_segment_data](data_ibm_app_config_segment.md) | data | Reads ibm_app_config_segment. |
| [ibm_app_config_segment](ibm_app_config_segment.md) | resource | Manages ibm_app_config_segment. |
| [ibm_app_config_segments_data](data_ibm_app_config_segments.md) | data | Schema validation for the ibm_app_config_segments data source. |
| [ibm_app_config_snapshot_data](data_ibm_app_config_snapshot.md) | data | Schema validation for the ibm_app_config_snapshot data source. |
| [ibm_app_config_snapshot](ibm_app_config_snapshot.md) | resource | Schema validation for the ibm_app_config_snapshot resource. |
| [ibm_app_config_snapshots_data](data_ibm_app_config_snapshots.md) | data | Schema validation for the ibm_app_config_snapshots data source. |

## appid

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_appid_action_url_data](data_ibm_appid_action_url.md) | data | Schema validation for the ibm_appid_action_url data source. |
| [ibm_appid_action_url](ibm_appid_action_url.md) | resource | Schema validation for the ibm_appid_action_url resource. |
| [ibm_appid_apm_data](data_ibm_appid_apm.md) | data | Schema validation for the ibm_appid_apm data source. |
| [ibm_appid_apm](ibm_appid_apm.md) | resource | Schema validation for the ibm_appid_apm resource. |
| [ibm_appid_application_data](data_ibm_appid_application.md) | data | Reads ibm_appid_application. |
| [ibm_appid_application](ibm_appid_application.md) | resource | Manages ibm_appid_application. |
| [ibm_appid_application_roles_data](data_ibm_appid_application_roles.md) | data | Schema validation for the ibm_appid_application_roles data source. |
| [ibm_appid_application_roles](ibm_appid_application_roles.md) | resource | Schema validation for the ibm_appid_application_roles resource. |
| [ibm_appid_application_scopes_data](data_ibm_appid_application_scopes.md) | data | Schema validation for the ibm_appid_application_scopes data source. |
| [ibm_appid_application_scopes](ibm_appid_application_scopes.md) | resource | Schema validation for the ibm_appid_application_scopes resource. |
| [ibm_appid_applications_data](data_ibm_appid_applications.md) | data | Schema validation for the ibm_appid_applications data source. |
| [ibm_appid_audit_status_data](data_ibm_appid_audit_status.md) | data | Schema validation for the ibm_appid_audit_status data source. |
| [ibm_appid_audit_status](ibm_appid_audit_status.md) | resource | Schema validation for the ibm_appid_audit_status resource. |
| [ibm_appid_cloud_directory_template_data](data_ibm_appid_cloud_directory_template.md) | data | Schema validation for the ibm_appid_cloud_directory_template data source. |
| [ibm_appid_cloud_directory_template](ibm_appid_cloud_directory_template.md) | resource | Schema validation for the ibm_appid_cloud_directory_template resource. |
| [ibm_appid_cloud_directory_user_data](data_ibm_appid_cloud_directory_user.md) | data | Schema validation for the ibm_appid_cloud_directory_user data source. |
| [ibm_appid_cloud_directory_user](ibm_appid_cloud_directory_user.md) | resource | Schema validation for the ibm_appid_cloud_directory_user resource. |
| [ibm_appid_idp_cloud_directory_data](data_ibm_appid_idp_cloud_directory.md) | data | Schema validation for the ibm_appid_idp_cloud_directory data source. |
| [ibm_appid_idp_cloud_directory](ibm_appid_idp_cloud_directory.md) | resource | Schema validation for the ibm_appid_idp_cloud_directory resource. |
| [ibm_appid_idp_custom_data](data_ibm_appid_idp_custom.md) | data | Schema validation for the ibm_appid_idp_custom data source. |
| [ibm_appid_idp_custom](ibm_appid_idp_custom.md) | resource | Schema validation for the ibm_appid_idp_custom resource. |
| [ibm_appid_idp_facebook_data](data_ibm_appid_idp_facebook.md) | data | Schema validation for the ibm_appid_idp_facebook data source. |
| [ibm_appid_idp_facebook](ibm_appid_idp_facebook.md) | resource | Schema validation for the ibm_appid_idp_facebook resource. |
| [ibm_appid_idp_google_data](data_ibm_appid_idp_google.md) | data | Schema validation for the ibm_appid_idp_google data source. |
| [ibm_appid_idp_google](ibm_appid_idp_google.md) | resource | Schema validation for the ibm_appid_idp_google resource. |
| [ibm_appid_idp_saml_data](data_ibm_appid_idp_saml.md) | data | Schema validation for the ibm_appid_idp_saml data source. |
| [ibm_appid_idp_saml](ibm_appid_idp_saml.md) | resource | Schema validation for the ibm_appid_idp_saml resource. |
| [ibm_appid_idp_saml_metadata_data](data_ibm_appid_idp_saml_metadata.md) | data | Schema validation for the ibm_appid_idp_saml_metadata data source. |
| [ibm_appid_languages_data](data_ibm_appid_languages.md) | data | Schema validation for the ibm_appid_languages data source. |
| [ibm_appid_languages](ibm_appid_languages.md) | resource | Schema validation for the ibm_appid_languages resource. |
| [ibm_appid_mfa_data](data_ibm_appid_mfa.md) | data | Schema validation for the ibm_appid_mfa data source. |
| [ibm_appid_mfa](ibm_appid_mfa.md) | resource | Schema validation for the ibm_appid_mfa resource. |
| [ibm_appid_mfa_channel_data](data_ibm_appid_mfa_channel.md) | data | Schema validation for the ibm_appid_mfa_channel data source. |
| [ibm_appid_mfa_channel](ibm_appid_mfa_channel.md) | resource | Schema validation for the ibm_appid_mfa_channel resource. |
| [ibm_appid_password_regex_data](data_ibm_appid_password_regex.md) | data | Schema validation for the ibm_appid_password_regex data source. |
| [ibm_appid_password_regex](ibm_appid_password_regex.md) | resource | Schema validation for the ibm_appid_password_regex resource. |
| [ibm_appid_redirect_urls_data](data_ibm_appid_redirect_urls.md) | data | Schema validation for the ibm_appid_redirect_urls data source. |
| [ibm_appid_redirect_urls](ibm_appid_redirect_urls.md) | resource | Schema validation for the ibm_appid_redirect_urls resource. |
| [ibm_appid_role_data](data_ibm_appid_role.md) | data | Reads ibm_appid_role. |
| [ibm_appid_role](ibm_appid_role.md) | resource | Manages ibm_appid_role. |
| [ibm_appid_roles_data](data_ibm_appid_roles.md) | data | Schema validation for the ibm_appid_roles data source. |
| [ibm_appid_theme_color_data](data_ibm_appid_theme_color.md) | data | Schema validation for the ibm_appid_theme_color data source. |
| [ibm_appid_theme_color](ibm_appid_theme_color.md) | resource | Schema validation for the ibm_appid_theme_color resource. |
| [ibm_appid_theme_text_data](data_ibm_appid_theme_text.md) | data | Schema validation for the ibm_appid_theme_text data source. |
| [ibm_appid_theme_text](ibm_appid_theme_text.md) | resource | Schema validation for the ibm_appid_theme_text resource. |
| [ibm_appid_token_config_data](data_ibm_appid_token_config.md) | data | Schema validation for the ibm_appid_token_config data source. |
| [ibm_appid_token_config](ibm_appid_token_config.md) | resource | Schema validation for the ibm_appid_token_config resource. |
| [ibm_appid_user_roles_data](data_ibm_appid_user_roles.md) | data | Schema validation for the ibm_appid_user_roles data source. |
| [ibm_appid_user_roles](ibm_appid_user_roles.md) | resource | Schema validation for the ibm_appid_user_roles resource. |

## atracker

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_atracker_route](ibm_atracker_route.md) | resource | Manages ibm_atracker_route. |
| [ibm_atracker_routes_data](data_ibm_atracker_routes.md) | data | Reads ibm_atracker_routes. |
| [ibm_atracker_settings](ibm_atracker_settings.md) | resource | Schema validation for the ibm_atracker_settings resource. |
| [ibm_atracker_target](ibm_atracker_target.md) | resource | Manages ibm_atracker_target. |
| [ibm_atracker_targets_data](data_ibm_atracker_targets.md) | data | Reads ibm_atracker_targets. |

## backuprecovery

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_backup_recoveries_data](data_ibm_backup_recoveries.md) | data | Schema validation for the ibm_backup_recoveries data source. |
| [ibm_backup_recovery_data](data_ibm_backup_recovery.md) | data | Reads ibm_backup_recovery. |
| [ibm_backup_recovery](ibm_backup_recovery.md) | resource | Manages ibm_backup_recovery. |
| [ibm_backup_recovery_agent_upgrade_task](ibm_backup_recovery_agent_upgrade_task.md) | resource | Manages ibm_backup_recovery_agent_upgrade_task. |
| [ibm_backup_recovery_agent_upgrade_tasks_data](data_ibm_backup_recovery_agent_upgrade_tasks.md) | data | Schema validation for the ibm_backup_recovery_agent_upgrade_tasks data source. |
| [ibm_backup_recovery_connection_registration_token](ibm_backup_recovery_connection_registration_token.md) | resource | Schema validation for the ibm_backup_recovery_connection_registration_token resource. |
| [ibm_backup_recovery_connector_access_token](ibm_backup_recovery_connector_access_token.md) | resource | Schema validation for the ibm_backup_recovery_connector_access_token resource. |
| [ibm_backup_recovery_connector_agent_config_data](data_ibm_backup_recovery_connector_agent_config.md) | data | Schema validation for the ibm_backup_recovery_connector_agent_config data source. |
| [ibm_backup_recovery_connector_agent_registration](ibm_backup_recovery_connector_agent_registration.md) | resource | Schema validation for the ibm_backup_recovery_connector_agent_registration resource. |
| [ibm_backup_recovery_connector_agents_data](data_ibm_backup_recovery_connector_agents.md) | data | Schema validation for the ibm_backup_recovery_connector_agents data source. |
| [ibm_backup_recovery_connector_get_users_data](data_ibm_backup_recovery_connector_get_users.md) | data | Schema validation for the ibm_backup_recovery_connector_get_users data source. |
| [ibm_backup_recovery_connector_logs_data](data_ibm_backup_recovery_connector_logs.md) | data | Schema validation for the ibm_backup_recovery_connector_logs data source. |
| [ibm_backup_recovery_connector_registration](ibm_backup_recovery_connector_registration.md) | resource | Schema validation for the ibm_backup_recovery_connector_registration resource. |
| [ibm_backup_recovery_connector_status_data](data_ibm_backup_recovery_connector_status.md) | data | Schema validation for the ibm_backup_recovery_connector_status data source. |
| [ibm_backup_recovery_connector_update_user](ibm_backup_recovery_connector_update_user.md) | resource | Schema validation for the ibm_backup_recovery_connector_update_user resource. |
| [ibm_backup_recovery_connectors_metadata_data](data_ibm_backup_recovery_connectors_metadata.md) | data | Schema validation for the ibm_backup_recovery_connectors_metadata data source. |
| [ibm_backup_recovery_data_source_connection](ibm_backup_recovery_data_source_connection.md) | resource | Schema validation for the ibm_backup_recovery_data_source_connection resource. |
| [ibm_backup_recovery_data_source_connections_data](data_ibm_backup_recovery_data_source_connections.md) | data | Schema validation for the ibm_backup_recovery_data_source_connections data source. |
| [ibm_backup_recovery_data_source_connector_patch](ibm_backup_recovery_data_source_connector_patch.md) | resource | Schema validation for the ibm_backup_recovery_data_source_connector_patch resource. |
| [ibm_backup_recovery_data_source_connectors_data](data_ibm_backup_recovery_data_source_connectors.md) | data | Schema validation for the ibm_backup_recovery_data_source_connectors data source. |
| [ibm_backup_recovery_download_agent_data](data_ibm_backup_recovery_download_agent.md) | data | Schema validation for the ibm_backup_recovery_download_agent data source. |
| [ibm_backup_recovery_download_files_data](data_ibm_backup_recovery_download_files.md) | data | Schema validation for the ibm_backup_recovery_download_files data source. |
| [ibm_backup_recovery_download_files_folders](ibm_backup_recovery_download_files_folders.md) | resource | Manages ibm_backup_recovery_download_files_folders. |
| [ibm_backup_recovery_download_indexed_files_data](data_ibm_backup_recovery_download_indexed_files.md) | data | Schema validation for the ibm_backup_recovery_download_indexed_files data source. |
| [ibm_backup_recovery_manager_cancel_cluster_upgrades](ibm_backup_recovery_manager_cancel_cluster_upgrades.md) | resource | Schema validation for the ibm_backup_recovery_manager_cancel_cluster_upgrades resource. |
| [ibm_backup_recovery_manager_create_cluster_upgrades](ibm_backup_recovery_manager_create_cluster_upgrades.md) | resource | Schema validation for the ibm_backup_recovery_manager_create_cluster_upgrades resource. |
| [ibm_backup_recovery_manager_get_alerts_data](data_ibm_backup_recovery_manager_get_alerts.md) | data | Schema validation for the ibm_backup_recovery_manager_get_alerts data source. |
| [ibm_backup_recovery_manager_get_alerts_resolution_data](data_ibm_backup_recovery_manager_get_alerts_resolution.md) | data | Schema validation for the ibm_backup_recovery_manager_get_alerts_resolution data source. |
| [ibm_backup_recovery_manager_get_alerts_stats_data](data_ibm_backup_recovery_manager_get_alerts_stats.md) | data | Schema validation for the ibm_backup_recovery_manager_get_alerts_stats data source. |
| [ibm_backup_recovery_manager_get_alerts_summary_data](data_ibm_backup_recovery_manager_get_alerts_summary.md) | data | Schema validation for the ibm_backup_recovery_manager_get_alerts_summary data source. |
| [ibm_backup_recovery_manager_get_cluster_info_data](data_ibm_backup_recovery_manager_get_cluster_info.md) | data | Schema validation for the ibm_backup_recovery_manager_get_cluster_info data source. |
| [ibm_backup_recovery_manager_get_compatible_clusters_data](data_ibm_backup_recovery_manager_get_compatible_clusters.md) | data | Schema validation for the ibm_backup_recovery_manager_get_compatible_clusters data source. |
| [ibm_backup_recovery_manager_get_management_alerts_data](data_ibm_backup_recovery_manager_get_management_alerts.md) | data | Schema validation for the ibm_backup_recovery_manager_get_management_alerts data source. |
| [ibm_backup_recovery_manager_get_management_alerts_summary_data](data_ibm_backup_recovery_manager_get_management_alerts_summary.md) | data | Schema validation for the ibm_backup_recovery_manager_get_management_alerts_summary data source. |
| [ibm_backup_recovery_manager_get_upgrades_info_data](data_ibm_backup_recovery_manager_get_upgrades_info.md) | data | Schema validation for the ibm_backup_recovery_manager_get_upgrades_info data source. |
| [ibm_backup_recovery_manager_update_cluster_upgrades](ibm_backup_recovery_manager_update_cluster_upgrades.md) | resource | Schema validation for the ibm_backup_recovery_manager_update_cluster_upgrades resource. |
| [ibm_backup_recovery_object_snapshots_data](data_ibm_backup_recovery_object_snapshots.md) | data | Schema validation for the ibm_backup_recovery_object_snapshots data source. |
| [ibm_backup_recovery_perform_action_on_protection_group_run_request](ibm_backup_recovery_perform_action_on_protection_group_run_request.md) | resource | Schema validation for the ibm_backup_recovery_perform_action_on_protection_group_run_request resource. |
| [ibm_backup_recovery_protection_group_data](data_ibm_backup_recovery_protection_group.md) | data | Reads ibm_backup_recovery_protection_group. |
| [ibm_backup_recovery_protection_group](ibm_backup_recovery_protection_group.md) | resource | Manages ibm_backup_recovery_protection_group. |
| [ibm_backup_recovery_protection_group_run_request](ibm_backup_recovery_protection_group_run_request.md) | resource | Schema validation for the ibm_backup_recovery_protection_group_run_request resource. |
| [ibm_backup_recovery_protection_group_runs_data](data_ibm_backup_recovery_protection_group_runs.md) | data | Schema validation for the ibm_backup_recovery_protection_group_runs data source. |
| [ibm_backup_recovery_protection_groups_data](data_ibm_backup_recovery_protection_groups.md) | data | Schema validation for the ibm_backup_recovery_protection_groups data source. |
| [ibm_backup_recovery_protection_policies_data](data_ibm_backup_recovery_protection_policies.md) | data | Schema validation for the ibm_backup_recovery_protection_policies data source. |
| [ibm_backup_recovery_protection_policy_data](data_ibm_backup_recovery_protection_policy.md) | data | Reads ibm_backup_recovery_protection_policy. |
| [ibm_backup_recovery_protection_policy](ibm_backup_recovery_protection_policy.md) | resource | Manages ibm_backup_recovery_protection_policy. |
| [ibm_backup_recovery_protection_source_refresh](ibm_backup_recovery_protection_source_refresh.md) | resource | Schema validation for the ibm_backup_recovery_protection_source_refresh resource. |
| [ibm_backup_recovery_protection_sources_data](data_ibm_backup_recovery_protection_sources.md) | data | Schema validation for the ibm_backup_recovery_protection_sources data source. |
| [ibm_backup_recovery_registration_info_data](data_ibm_backup_recovery_registration_info.md) | data | Schema validation for the ibm_backup_recovery_registration_info data source. |
| [ibm_backup_recovery_restore_points](ibm_backup_recovery_restore_points.md) | resource | Schema validation for the ibm_backup_recovery_restore_points resource. |
| [ibm_backup_recovery_search_indexed_object_data](data_ibm_backup_recovery_search_indexed_object.md) | data | Schema validation for the ibm_backup_recovery_search_indexed_object data source. |
| [ibm_backup_recovery_search_objects_data](data_ibm_backup_recovery_search_objects.md) | data | Schema validation for the ibm_backup_recovery_search_objects data source. |
| [ibm_backup_recovery_search_protected_objects_data](data_ibm_backup_recovery_search_protected_objects.md) | data | Schema validation for the ibm_backup_recovery_search_protected_objects data source. |
| [ibm_backup_recovery_source_registration_data](data_ibm_backup_recovery_source_registration.md) | data | Reads ibm_backup_recovery_source_registration. |
| [ibm_backup_recovery_source_registration](ibm_backup_recovery_source_registration.md) | resource | Manages ibm_backup_recovery_source_registration. |
| [ibm_backup_recovery_source_registrations_data](data_ibm_backup_recovery_source_registrations.md) | data | Schema validation for the ibm_backup_recovery_source_registrations data source. |
| [ibm_backup_recovery_update_protection_group_run_request](ibm_backup_recovery_update_protection_group_run_request.md) | resource | Schema validation for the ibm_backup_recovery_update_protection_group_run_request resource. |

## catalogmanagement

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cm_account_data](data_ibm_cm_account.md) | data | Schema validation for the ibm_cm_account data source. |
| [ibm_cm_account](ibm_cm_account.md) | resource | Schema validation for the ibm_cm_account resource. |
| [ibm_cm_catalog_data](data_ibm_cm_catalog.md) | data | Schema validation for the ibm_cm_catalog data source. |
| [ibm_cm_catalog](ibm_cm_catalog.md) | resource | Schema validation for the ibm_cm_catalog resource. |
| [ibm_cm_object_data](data_ibm_cm_object.md) | data | Reads ibm_cm_object. |
| [ibm_cm_object](ibm_cm_object.md) | resource | Manages ibm_cm_object. |
| [ibm_cm_offering_data](data_ibm_cm_offering.md) | data | Reads ibm_cm_offering. |
| [ibm_cm_offering](ibm_cm_offering.md) | resource | Manages ibm_cm_offering. |
| [ibm_cm_offering_instance_data](data_ibm_cm_offering_instance.md) | data | Schema validation for the ibm_cm_offering_instance data source. |
| [ibm_cm_offering_instance](ibm_cm_offering_instance.md) | resource | Schema validation for the ibm_cm_offering_instance resource. |
| [ibm_cm_preset_data](data_ibm_cm_preset.md) | data | Schema validation for the ibm_cm_preset data source. |
| [ibm_cm_validation](ibm_cm_validation.md) | resource | Schema validation for the ibm_cm_validation resource. |
| [ibm_cm_version_data](data_ibm_cm_version.md) | data | Schema validation for the ibm_cm_version data source. |
| [ibm_cm_version](ibm_cm_version.md) | resource | Manages ibm_cm_version. |

## cdtektonpipeline

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cd_tekton_pipeline_data](data_ibm_cd_tekton_pipeline.md) | data | Reads ibm_cd_tekton_pipeline. |
| [ibm_cd_tekton_pipeline](ibm_cd_tekton_pipeline.md) | resource | Manages ibm_cd_tekton_pipeline. |
| [ibm_cd_tekton_pipeline_definition_data](data_ibm_cd_tekton_pipeline_definition.md) | data | Schema validation for the ibm_cd_tekton_pipeline_definition data source. |
| [ibm_cd_tekton_pipeline_definition](ibm_cd_tekton_pipeline_definition.md) | resource | Schema validation for the ibm_cd_tekton_pipeline_definition resource. |
| [ibm_cd_tekton_pipeline_property_data](data_ibm_cd_tekton_pipeline_property.md) | data | Reads ibm_cd_tekton_pipeline_property. |
| [ibm_cd_tekton_pipeline_property](ibm_cd_tekton_pipeline_property.md) | resource | Manages ibm_cd_tekton_pipeline_property. |
| [ibm_cd_tekton_pipeline_trigger_data](data_ibm_cd_tekton_pipeline_trigger.md) | data | Reads ibm_cd_tekton_pipeline_trigger. |
| [ibm_cd_tekton_pipeline_trigger](ibm_cd_tekton_pipeline_trigger.md) | resource | Manages ibm_cd_tekton_pipeline_trigger. |
| [ibm_cd_tekton_pipeline_trigger_property_data](data_ibm_cd_tekton_pipeline_trigger_property.md) | data | Reads ibm_cd_tekton_pipeline_trigger_property. |
| [ibm_cd_tekton_pipeline_trigger_property](ibm_cd_tekton_pipeline_trigger_property.md) | resource | Manages ibm_cd_tekton_pipeline_trigger_property. |

## cdtoolchain

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cd_toolchain_data](data_ibm_cd_toolchain.md) | data | Reads ibm_cd_toolchain. |
| [ibm_cd_toolchain](ibm_cd_toolchain.md) | resource | Manages ibm_cd_toolchain. |
| [ibm_cd_toolchain_tool_appconfig_data](data_ibm_cd_toolchain_tool_appconfig.md) | data | Reads ibm_cd_toolchain_tool_appconfig. |
| [ibm_cd_toolchain_tool_appconfig](ibm_cd_toolchain_tool_appconfig.md) | resource | Manages ibm_cd_toolchain_tool_appconfig. |
| [ibm_cd_toolchain_tool_artifactory_data](data_ibm_cd_toolchain_tool_artifactory.md) | data | Reads ibm_cd_toolchain_tool_artifactory. |
| [ibm_cd_toolchain_tool_artifactory](ibm_cd_toolchain_tool_artifactory.md) | resource | Manages ibm_cd_toolchain_tool_artifactory. |
| [ibm_cd_toolchain_tool_bitbucketgit_data](data_ibm_cd_toolchain_tool_bitbucketgit.md) | data | Reads ibm_cd_toolchain_tool_bitbucketgit. |
| [ibm_cd_toolchain_tool_bitbucketgit](ibm_cd_toolchain_tool_bitbucketgit.md) | resource | Manages ibm_cd_toolchain_tool_bitbucketgit. |
| [ibm_cd_toolchain_tool_cos_data](data_ibm_cd_toolchain_tool_cos.md) | data | Reads ibm_cd_toolchain_tool_cos. |
| [ibm_cd_toolchain_tool_cos](ibm_cd_toolchain_tool_cos.md) | resource | Manages ibm_cd_toolchain_tool_cos. |
| [ibm_cd_toolchain_tool_custom_data](data_ibm_cd_toolchain_tool_custom.md) | data | Reads ibm_cd_toolchain_tool_custom. |
| [ibm_cd_toolchain_tool_custom](ibm_cd_toolchain_tool_custom.md) | resource | Manages ibm_cd_toolchain_tool_custom. |
| [ibm_cd_toolchain_tool_devopsinsights_data](data_ibm_cd_toolchain_tool_devopsinsights.md) | data | Reads ibm_cd_toolchain_tool_devopsinsights. |
| [ibm_cd_toolchain_tool_devopsinsights](ibm_cd_toolchain_tool_devopsinsights.md) | resource | Manages ibm_cd_toolchain_tool_devopsinsights. |
| [ibm_cd_toolchain_tool_eventnotifications_data](data_ibm_cd_toolchain_tool_eventnotifications.md) | data | Reads ibm_cd_toolchain_tool_eventnotifications. |
| [ibm_cd_toolchain_tool_eventnotifications](ibm_cd_toolchain_tool_eventnotifications.md) | resource | Manages ibm_cd_toolchain_tool_eventnotifications. |
| [ibm_cd_toolchain_tool_githubconsolidated_data](data_ibm_cd_toolchain_tool_githubconsolidated.md) | data | Reads ibm_cd_toolchain_tool_githubconsolidated. |
| [ibm_cd_toolchain_tool_githubconsolidated](ibm_cd_toolchain_tool_githubconsolidated.md) | resource | Manages ibm_cd_toolchain_tool_githubconsolidated. |
| [ibm_cd_toolchain_tool_gitlab_data](data_ibm_cd_toolchain_tool_gitlab.md) | data | Reads ibm_cd_toolchain_tool_gitlab. |
| [ibm_cd_toolchain_tool_gitlab](ibm_cd_toolchain_tool_gitlab.md) | resource | Manages ibm_cd_toolchain_tool_gitlab. |
| [ibm_cd_toolchain_tool_hashicorpvault_data](data_ibm_cd_toolchain_tool_hashicorpvault.md) | data | Reads ibm_cd_toolchain_tool_hashicorpvault. |
| [ibm_cd_toolchain_tool_hashicorpvault](ibm_cd_toolchain_tool_hashicorpvault.md) | resource | Manages ibm_cd_toolchain_tool_hashicorpvault. |
| [ibm_cd_toolchain_tool_hostedgit_data](data_ibm_cd_toolchain_tool_hostedgit.md) | data | Reads ibm_cd_toolchain_tool_hostedgit. |
| [ibm_cd_toolchain_tool_hostedgit](ibm_cd_toolchain_tool_hostedgit.md) | resource | Manages ibm_cd_toolchain_tool_hostedgit. |
| [ibm_cd_toolchain_tool_jenkins_data](data_ibm_cd_toolchain_tool_jenkins.md) | data | Reads ibm_cd_toolchain_tool_jenkins. |
| [ibm_cd_toolchain_tool_jenkins](ibm_cd_toolchain_tool_jenkins.md) | resource | Manages ibm_cd_toolchain_tool_jenkins. |
| [ibm_cd_toolchain_tool_jira_data](data_ibm_cd_toolchain_tool_jira.md) | data | Reads ibm_cd_toolchain_tool_jira. |
| [ibm_cd_toolchain_tool_jira](ibm_cd_toolchain_tool_jira.md) | resource | Manages ibm_cd_toolchain_tool_jira. |
| [ibm_cd_toolchain_tool_keyprotect_data](data_ibm_cd_toolchain_tool_keyprotect.md) | data | Reads ibm_cd_toolchain_tool_keyprotect. |
| [ibm_cd_toolchain_tool_keyprotect](ibm_cd_toolchain_tool_keyprotect.md) | resource | Manages ibm_cd_toolchain_tool_keyprotect. |
| [ibm_cd_toolchain_tool_nexus_data](data_ibm_cd_toolchain_tool_nexus.md) | data | Reads ibm_cd_toolchain_tool_nexus. |
| [ibm_cd_toolchain_tool_nexus](ibm_cd_toolchain_tool_nexus.md) | resource | Manages ibm_cd_toolchain_tool_nexus. |
| [ibm_cd_toolchain_tool_pagerduty_data](data_ibm_cd_toolchain_tool_pagerduty.md) | data | Reads ibm_cd_toolchain_tool_pagerduty. |
| [ibm_cd_toolchain_tool_pagerduty](ibm_cd_toolchain_tool_pagerduty.md) | resource | Manages ibm_cd_toolchain_tool_pagerduty. |
| [ibm_cd_toolchain_tool_pipeline_data](data_ibm_cd_toolchain_tool_pipeline.md) | data | Reads ibm_cd_toolchain_tool_pipeline. |
| [ibm_cd_toolchain_tool_pipeline](ibm_cd_toolchain_tool_pipeline.md) | resource | Manages ibm_cd_toolchain_tool_pipeline. |
| [ibm_cd_toolchain_tool_privateworker_data](data_ibm_cd_toolchain_tool_privateworker.md) | data | Reads ibm_cd_toolchain_tool_privateworker. |
| [ibm_cd_toolchain_tool_privateworker](ibm_cd_toolchain_tool_privateworker.md) | resource | Manages ibm_cd_toolchain_tool_privateworker. |
| [ibm_cd_toolchain_tool_saucelabs_data](data_ibm_cd_toolchain_tool_saucelabs.md) | data | Reads ibm_cd_toolchain_tool_saucelabs. |
| [ibm_cd_toolchain_tool_saucelabs](ibm_cd_toolchain_tool_saucelabs.md) | resource | Manages ibm_cd_toolchain_tool_saucelabs. |
| [ibm_cd_toolchain_tool_secretsmanager_data](data_ibm_cd_toolchain_tool_secretsmanager.md) | data | Reads ibm_cd_toolchain_tool_secretsmanager. |
| [ibm_cd_toolchain_tool_secretsmanager](ibm_cd_toolchain_tool_secretsmanager.md) | resource | Manages ibm_cd_toolchain_tool_secretsmanager. |
| [ibm_cd_toolchain_tool_securitycompliance_data](data_ibm_cd_toolchain_tool_securitycompliance.md) | data | Reads ibm_cd_toolchain_tool_securitycompliance. |
| [ibm_cd_toolchain_tool_securitycompliance](ibm_cd_toolchain_tool_securitycompliance.md) | resource | Manages ibm_cd_toolchain_tool_securitycompliance. |
| [ibm_cd_toolchain_tool_slack_data](data_ibm_cd_toolchain_tool_slack.md) | data | Reads ibm_cd_toolchain_tool_slack. |
| [ibm_cd_toolchain_tool_slack](ibm_cd_toolchain_tool_slack.md) | resource | Manages ibm_cd_toolchain_tool_slack. |
| [ibm_cd_toolchain_tool_sonarqube_data](data_ibm_cd_toolchain_tool_sonarqube.md) | data | Reads ibm_cd_toolchain_tool_sonarqube. |
| [ibm_cd_toolchain_tool_sonarqube](ibm_cd_toolchain_tool_sonarqube.md) | resource | Manages ibm_cd_toolchain_tool_sonarqube. |
| [ibm_cd_toolchains_data](data_ibm_cd_toolchains.md) | data | Reads ibm_cd_toolchains. |

## cis

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cis_data](data_ibm_cis.md) | data | Reads ibm_cis. |
| [ibm_cis](ibm_cis.md) | resource | Manages ibm_cis. |
| [ibm_cis_advanced_certificate_pack_order](ibm_cis_advanced_certificate_pack_order.md) | resource | Schema validation for the ibm_cis_advanced_certificate_pack_order resource. |
| [ibm_cis_alert](ibm_cis_alert.md) | resource | Manages ibm_cis_alert. |
| [ibm_cis_alerts_data](data_ibm_cis_alerts.md) | data | Schema validation for the ibm_cis_alerts data source. |
| [ibm_cis_bot_analytics_data](data_ibm_cis_bot_analytics.md) | data | Schema validation for the ibm_cis_bot_analytics data source. |
| [ibm_cis_bot_management](ibm_cis_bot_management.md) | resource | Schema validation for the ibm_cis_bot_management resource. |
| [ibm_cis_bot_managements_data](data_ibm_cis_bot_managements.md) | data | Schema validation for the ibm_cis_bot_managements data source. |
| [ibm_cis_cache_settings_data](data_ibm_cis_cache_settings.md) | data | Schema validation for the ibm_cis_cache_settings data source. |
| [ibm_cis_cache_settings](ibm_cis_cache_settings.md) | resource | Schema validation for the ibm_cis_cache_settings resource. |
| [ibm_cis_certificate_order](ibm_cis_certificate_order.md) | resource | Schema validation for the ibm_cis_certificate_order resource. |
| [ibm_cis_certificate_upload](ibm_cis_certificate_upload.md) | resource | Schema validation for the ibm_cis_certificate_upload resource. |
| [ibm_cis_certificates_data](data_ibm_cis_certificates.md) | data | Schema validation for the ibm_cis_certificates data source. |
| [ibm_cis_custom_certificates_data](data_ibm_cis_custom_certificates.md) | data | Schema validation for the ibm_cis_custom_certificates data source. |
| [ibm_cis_custom_list](ibm_cis_custom_list.md) | resource | Manages ibm_cis_custom_list. |
| [ibm_cis_custom_list_items_data](data_ibm_cis_custom_list_items.md) | data | Schema validation for the ibm_cis_custom_list_items data source. |
| [ibm_cis_custom_list_items](ibm_cis_custom_list_items.md) | resource | Schema validation for the ibm_cis_custom_list_items resource. |
| [ibm_cis_custom_lists_data](data_ibm_cis_custom_lists.md) | data | Schema validation for the ibm_cis_custom_lists data source. |
| [ibm_cis_custom_page](ibm_cis_custom_page.md) | resource | Schema validation for the ibm_cis_custom_page resource. |
| [ibm_cis_custom_pages_data](data_ibm_cis_custom_pages.md) | data | Schema validation for the ibm_cis_custom_pages data source. |
| [ibm_cis_dns_record](ibm_cis_dns_record.md) | resource | Manages ibm_cis_dns_record. |
| [ibm_cis_dns_records_data](data_ibm_cis_dns_records.md) | data | Schema validation for the ibm_cis_dns_records data source. |
| [ibm_cis_dns_records_batch](ibm_cis_dns_records_batch.md) | resource | Schema validation for the ibm_cis_dns_records_batch resource. |
| [ibm_cis_dns_records_import](ibm_cis_dns_records_import.md) | resource | Schema validation for the ibm_cis_dns_records_import resource. |
| [ibm_cis_domain_data](data_ibm_cis_domain.md) | data | Schema validation for the ibm_cis_domain data source. |
| [ibm_cis_domain](ibm_cis_domain.md) | resource | Schema validation for the ibm_cis_domain resource. |
| [ibm_cis_domain_settings](ibm_cis_domain_settings.md) | resource | Schema validation for the ibm_cis_domain_settings resource. |
| [ibm_cis_edge_functions_action](ibm_cis_edge_functions_action.md) | resource | Schema validation for the ibm_cis_edge_functions_action resource. |
| [ibm_cis_edge_functions_actions_data](data_ibm_cis_edge_functions_actions.md) | data | Schema validation for the ibm_cis_edge_functions_actions data source. |
| [ibm_cis_edge_functions_trigger](ibm_cis_edge_functions_trigger.md) | resource | Schema validation for the ibm_cis_edge_functions_trigger resource. |
| [ibm_cis_edge_functions_triggers_data](data_ibm_cis_edge_functions_triggers.md) | data | Schema validation for the ibm_cis_edge_functions_triggers data source. |
| [ibm_cis_filter](ibm_cis_filter.md) | resource | Schema validation for the ibm_cis_filter resource. |
| [ibm_cis_filters_data](data_ibm_cis_filters.md) | data | Schema validation for the ibm_cis_filters data source. |
| [ibm_cis_firewall_data](data_ibm_cis_firewall.md) | data | Schema validation for the ibm_cis_firewall data source. |
| [ibm_cis_firewall](ibm_cis_firewall.md) | resource | Schema validation for the ibm_cis_firewall resource. |
| [ibm_cis_firewall_rule](ibm_cis_firewall_rule.md) | resource | Schema validation for the ibm_cis_firewall_rule resource. |
| [ibm_cis_firewall_rules_data](data_ibm_cis_firewall_rules.md) | data | Schema validation for the ibm_cis_firewall_rules data source. |
| [ibm_cis_global_load_balancer](ibm_cis_global_load_balancer.md) | resource | Manages ibm_cis_global_load_balancer. |
| [ibm_cis_global_load_balancers_data](data_ibm_cis_global_load_balancers.md) | data | Schema validation for the ibm_cis_global_load_balancers data source. |
| [ibm_cis_healthcheck](ibm_cis_healthcheck.md) | resource | Schema validation for the ibm_cis_healthcheck resource. |
| [ibm_cis_healthchecks_data](data_ibm_cis_healthchecks.md) | data | Schema validation for the ibm_cis_healthchecks data source. |
| [ibm_cis_ip_addresses_data](data_ibm_cis_ip_addresses.md) | data | Schema validation for the ibm_cis_ip_addresses data source. |
| [ibm_cis_logpush_job](ibm_cis_logpush_job.md) | resource | Manages ibm_cis_logpush_job. |
| [ibm_cis_logpush_jobs_data](data_ibm_cis_logpush_jobs.md) | data | Schema validation for the ibm_cis_logpush_jobs data source. |
| [ibm_cis_managed_lists_data](data_ibm_cis_managed_lists.md) | data | Schema validation for the ibm_cis_managed_lists data source. |
| [ibm_cis_mtls](ibm_cis_mtls.md) | resource | Manages ibm_cis_mtls. |
| [ibm_cis_mtls_app](ibm_cis_mtls_app.md) | resource | Manages ibm_cis_mtls_app. |
| [ibm_cis_mtls_apps_data](data_ibm_cis_mtls_apps.md) | data | Schema validation for the ibm_cis_mtls_apps data source. |
| [ibm_cis_mtlss_data](data_ibm_cis_mtlss.md) | data | Schema validation for the ibm_cis_mtlss data source. |
| [ibm_cis_origin_auth](ibm_cis_origin_auth.md) | resource | Schema validation for the ibm_cis_origin_auth resource. |
| [ibm_cis_origin_auths_data](data_ibm_cis_origin_auths.md) | data | Schema validation for the ibm_cis_origin_auths data source. |
| [ibm_cis_origin_certificate_order](ibm_cis_origin_certificate_order.md) | resource | Schema validation for the ibm_cis_origin_certificate_order resource. |
| [ibm_cis_origin_certificates_data](data_ibm_cis_origin_certificates.md) | data | Schema validation for the ibm_cis_origin_certificates data source. |
| [ibm_cis_origin_pool](ibm_cis_origin_pool.md) | resource | Manages ibm_cis_origin_pool. |
| [ibm_cis_origin_pools_data](data_ibm_cis_origin_pools.md) | data | Schema validation for the ibm_cis_origin_pools data source. |
| [ibm_cis_page_rule](ibm_cis_page_rule.md) | resource | Schema validation for the ibm_cis_page_rule resource. |
| [ibm_cis_page_rules_data](data_ibm_cis_page_rules.md) | data | Schema validation for the ibm_cis_page_rules data source. |
| [ibm_cis_range_app](ibm_cis_range_app.md) | resource | Schema validation for the ibm_cis_range_app resource. |
| [ibm_cis_range_apps_data](data_ibm_cis_range_apps.md) | data | Schema validation for the ibm_cis_range_apps data source. |
| [ibm_cis_rate_limit_data](data_ibm_cis_rate_limit.md) | data | Schema validation for the ibm_cis_rate_limit data source. |
| [ibm_cis_rate_limit](ibm_cis_rate_limit.md) | resource | Schema validation for the ibm_cis_rate_limit resource. |
| [ibm_cis_routing](ibm_cis_routing.md) | resource | Schema validation for the ibm_cis_routing resource. |
| [ibm_cis_ruleset](ibm_cis_ruleset.md) | resource | Schema validation for the ibm_cis_ruleset resource. |
| [ibm_cis_ruleset_entrypoint_version](ibm_cis_ruleset_entrypoint_version.md) | resource | Schema validation for the ibm_cis_ruleset_entrypoint_version resource. |
| [ibm_cis_ruleset_entrypoint_versions_data](data_ibm_cis_ruleset_entrypoint_versions.md) | data | Schema validation for the ibm_cis_ruleset_entrypoint_versions data source. |
| [ibm_cis_ruleset_rule](ibm_cis_ruleset_rule.md) | resource | Schema validation for the ibm_cis_ruleset_rule resource. |
| [ibm_cis_ruleset_rules_by_tag_data](data_ibm_cis_ruleset_rules_by_tag.md) | data | Schema validation for the ibm_cis_ruleset_rules_by_tag data source. |
| [ibm_cis_ruleset_version_detach](ibm_cis_ruleset_version_detach.md) | resource | Schema validation for the ibm_cis_ruleset_version_detach resource. |
| [ibm_cis_ruleset_versions_data](data_ibm_cis_ruleset_versions.md) | data | Schema validation for the ibm_cis_ruleset_versions data source. |
| [ibm_cis_rulesets_data](data_ibm_cis_rulesets.md) | data | Schema validation for the ibm_cis_rulesets data source. |
| [ibm_cis_tls_settings](ibm_cis_tls_settings.md) | resource | Schema validation for the ibm_cis_tls_settings resource. |
| [ibm_cis_waf_group](ibm_cis_waf_group.md) | resource | Manages ibm_cis_waf_group. |
| [ibm_cis_waf_groups_data](data_ibm_cis_waf_groups.md) | data | Schema validation for the ibm_cis_waf_groups data source. |
| [ibm_cis_waf_package](ibm_cis_waf_package.md) | resource | Manages ibm_cis_waf_package. |
| [ibm_cis_waf_packages_data](data_ibm_cis_waf_packages.md) | data | Schema validation for the ibm_cis_waf_packages data source. |
| [ibm_cis_waf_rule](ibm_cis_waf_rule.md) | resource | Schema validation for the ibm_cis_waf_rule resource. |
| [ibm_cis_waf_rules_data](data_ibm_cis_waf_rules.md) | data | Schema validation for the ibm_cis_waf_rules data source. |
| [ibm_cis_webhook](ibm_cis_webhook.md) | resource | Manages ibm_cis_webhook. |
| [ibm_cis_webhooks_data](data_ibm_cis_webhooks.md) | data | Schema validation for the ibm_cis_webhooks data source. |

## classicinfrastructure

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cdn](ibm_cdn.md) | resource | Schema validation for the ibm_cdn resource. |
| [ibm_compute_autoscale_group](ibm_compute_autoscale_group.md) | resource | Manages ibm_compute_autoscale_group. |
| [ibm_compute_autoscale_policy](ibm_compute_autoscale_policy.md) | resource | Manages ibm_compute_autoscale_policy. |
| [ibm_compute_bare_metal_data](data_ibm_compute_bare_metal.md) | data | Schema validation for the ibm_compute_bare_metal data source. |
| [ibm_compute_bare_metal](ibm_compute_bare_metal.md) | resource | Schema validation for the ibm_compute_bare_metal resource. |
| [ibm_compute_dedicated_host](ibm_compute_dedicated_host.md) | resource | Schema validation for the ibm_compute_dedicated_host resource. |
| [ibm_compute_image_template_data](data_ibm_compute_image_template.md) | data | Reads ibm_compute_image_template. |
| [ibm_compute_monitor](ibm_compute_monitor.md) | resource | Schema validation for the ibm_compute_monitor resource. |
| [ibm_compute_placement_group_data](data_ibm_compute_placement_group.md) | data | Schema validation for the ibm_compute_placement_group data source. |
| [ibm_compute_placement_group](ibm_compute_placement_group.md) | resource | Manages ibm_compute_placement_group. |
| [ibm_compute_provisioning_hook](ibm_compute_provisioning_hook.md) | resource | Manages ibm_compute_provisioning_hook. |
| [ibm_compute_reserved_capacity_data](data_ibm_compute_reserved_capacity.md) | data | Reads ibm_compute_reserved_capacity. |
| [ibm_compute_reserved_capacity](ibm_compute_reserved_capacity.md) | resource | Manages ibm_compute_reserved_capacity. |
| [ibm_compute_ssh_key_data](data_ibm_compute_ssh_key.md) | data | Schema validation for the ibm_compute_ssh_key data source. |
| [ibm_compute_ssh_key](ibm_compute_ssh_key.md) | resource | Schema validation for the ibm_compute_ssh_key resource. |
| [ibm_compute_ssl_certificate](ibm_compute_ssl_certificate.md) | resource | Schema validation for the ibm_compute_ssl_certificate resource. |
| [ibm_compute_user](ibm_compute_user.md) | resource | Schema validation for the ibm_compute_user resource. |
| [ibm_compute_vm_instance_data](data_ibm_compute_vm_instance.md) | data | Schema validation for the ibm_compute_vm_instance data source. |
| [ibm_compute_vm_instance](ibm_compute_vm_instance.md) | resource | Schema validation for the ibm_compute_vm_instance resource. |
| [ibm_dns_domain_data](data_ibm_dns_domain.md) | data | Reads ibm_dns_domain. |
| [ibm_dns_domain](ibm_dns_domain.md) | resource | Manages ibm_dns_domain. |
| [ibm_dns_domain_registration_data](data_ibm_dns_domain_registration.md) | data | Reads ibm_dns_domain_registration. |
| [ibm_dns_domain_registration_nameservers](ibm_dns_domain_registration_nameservers.md) | resource | Schema validation for the ibm_dns_domain_registration_nameservers resource. |
| [ibm_dns_record](ibm_dns_record.md) | resource | Schema validation for the ibm_dns_record resource. |
| [ibm_dns_reverse_record](ibm_dns_reverse_record.md) | resource | Schema validation for the ibm_dns_reverse_record resource. |
| [ibm_dns_secondary_data](data_ibm_dns_secondary.md) | data | Schema validation for the ibm_dns_secondary data source. |
| [ibm_dns_secondary](ibm_dns_secondary.md) | resource | Schema validation for the ibm_dns_secondary resource. |
| [ibm_firewall](ibm_firewall.md) | resource | Schema validation for the ibm_firewall resource. |
| [ibm_firewall_policy](ibm_firewall_policy.md) | resource | Schema validation for the ibm_firewall_policy resource. |
| [ibm_hardware_firewall_shared](ibm_hardware_firewall_shared.md) | resource | Schema validation for the ibm_hardware_firewall_shared resource. |
| [ibm_ipsec_vpn](ibm_ipsec_vpn.md) | resource | Schema validation for the ibm_ipsec_vpn resource. |
| [ibm_lb](ibm_lb.md) | resource | Schema validation for the ibm_lb resource. |
| [ibm_lb_service](ibm_lb_service.md) | resource | Schema validation for the ibm_lb_service resource. |
| [ibm_lb_service_group](ibm_lb_service_group.md) | resource | Schema validation for the ibm_lb_service_group resource. |
| [ibm_lb_vpx](ibm_lb_vpx.md) | resource | Manages ibm_lb_vpx. |
| [ibm_lb_vpx_ha](ibm_lb_vpx_ha.md) | resource | Schema validation for the ibm_lb_vpx_ha resource. |
| [ibm_lb_vpx_service](ibm_lb_vpx_service.md) | resource | Manages ibm_lb_vpx_service. |
| [ibm_lb_vpx_vip](ibm_lb_vpx_vip.md) | resource | Manages ibm_lb_vpx_vip. |
| [ibm_lbaas_data](data_ibm_lbaas.md) | data | Schema validation for the ibm_lbaas data source. |
| [ibm_lbaas](ibm_lbaas.md) | resource | Manages ibm_lbaas. |
| [ibm_lbaas_health_monitor](ibm_lbaas_health_monitor.md) | resource | Schema validation for the ibm_lbaas_health_monitor resource. |
| [ibm_lbaas_server_instance_attachment](ibm_lbaas_server_instance_attachment.md) | resource | Schema validation for the ibm_lbaas_server_instance_attachment resource. |
| [ibm_multi_vlan_firewall](ibm_multi_vlan_firewall.md) | resource | Manages ibm_multi_vlan_firewall. |
| [ibm_network_gateway](ibm_network_gateway.md) | resource | Manages ibm_network_gateway. |
| [ibm_network_gateway_vlan_association](ibm_network_gateway_vlan_association.md) | resource | Schema validation for the ibm_network_gateway_vlan_association resource. |
| [ibm_network_interface_sg_attachment](ibm_network_interface_sg_attachment.md) | resource | Schema validation for the ibm_network_interface_sg_attachment resource. |
| [ibm_network_public_ip](ibm_network_public_ip.md) | resource | Schema validation for the ibm_network_public_ip resource. |
| [ibm_network_vlan_data](data_ibm_network_vlan.md) | data | Schema validation for the ibm_network_vlan data source. |
| [ibm_network_vlan](ibm_network_vlan.md) | resource | Manages ibm_network_vlan. |
| [ibm_network_vlan_spanning](ibm_network_vlan_spanning.md) | resource | Schema validation for the ibm_network_vlan_spanning resource. |
| [ibm_object_storage_account](ibm_object_storage_account.md) | resource | Schema validation for the ibm_object_storage_account resource. |
| [ibm_security_group_data](data_ibm_security_group.md) | data | Reads ibm_security_group. |
| [ibm_security_group](ibm_security_group.md) | resource | Manages ibm_security_group. |
| [ibm_security_group_rule](ibm_security_group_rule.md) | resource | Schema validation for the ibm_security_group_rule resource. |
| [ibm_ssl_certificate](ibm_ssl_certificate.md) | resource | Schema validation for the ibm_ssl_certificate resource. |
| [ibm_storage_block](ibm_storage_block.md) | resource | Schema validation for the ibm_storage_block resource. |
| [ibm_storage_evault](ibm_storage_evault.md) | resource | Schema validation for the ibm_storage_evault resource. |
| [ibm_storage_file](ibm_storage_file.md) | resource | Schema validation for the ibm_storage_file resource. |
| [ibm_subnet](ibm_subnet.md) | resource | Schema validation for the ibm_subnet resource. |

## cloudant

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cloudant_data](data_ibm_cloudant.md) | data | Reads ibm_cloudant. |
| [ibm_cloudant](ibm_cloudant.md) | resource | Manages ibm_cloudant. |
| [ibm_cloudant_database_data](data_ibm_cloudant_database.md) | data | Schema validation for the ibm_cloudant_database data source. |
| [ibm_cloudant_database](ibm_cloudant_database.md) | resource | Schema validation for the ibm_cloudant_database resource. |

## cloudfoundry

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_account_data](data_ibm_account.md) | data | Schema validation for the ibm_account data source. |
| [ibm_app_data](data_ibm_app.md) | data | Reads ibm_app. |
| [ibm_app](ibm_app.md) | resource | Manages ibm_app. |
| [ibm_app_domain_private_data](data_ibm_app_domain_private.md) | data | Reads ibm_app_domain_private. |
| [ibm_app_domain_private](ibm_app_domain_private.md) | resource | Manages ibm_app_domain_private. |
| [ibm_app_domain_shared_data](data_ibm_app_domain_shared.md) | data | Reads ibm_app_domain_shared. |
| [ibm_app_domain_shared](ibm_app_domain_shared.md) | resource | Manages ibm_app_domain_shared. |
| [ibm_app_route_data](data_ibm_app_route.md) | data | Schema validation for the ibm_app_route data source. |
| [ibm_app_route](ibm_app_route.md) | resource | Schema validation for the ibm_app_route resource. |
| [ibm_org_data](data_ibm_org.md) | data | Reads ibm_org. |
| [ibm_org](ibm_org.md) | resource | Manages ibm_org. |
| [ibm_org_quota_data](data_ibm_org_quota.md) | data | Reads ibm_org_quota. |
| [ibm_service_instance_data](data_ibm_service_instance.md) | data | Reads ibm_service_instance. |
| [ibm_service_instance](ibm_service_instance.md) | resource | Manages ibm_service_instance. |
| [ibm_service_key_data](data_ibm_service_key.md) | data | Reads ibm_service_key. |
| [ibm_service_key](ibm_service_key.md) | resource | Manages ibm_service_key. |
| [ibm_service_plan_data](data_ibm_service_plan.md) | data | Schema validation for the ibm_service_plan data source. |
| [ibm_space_data](data_ibm_space.md) | data | Reads ibm_space. |
| [ibm_space](ibm_space.md) | resource | Manages ibm_space. |

## cloudshell

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cloud_shell_account_settings_data](data_ibm_cloud_shell_account_settings.md) | data | Schema validation for the ibm_cloud_shell_account_settings data source. |
| [ibm_cloud_shell_account_settings](ibm_cloud_shell_account_settings.md) | resource | Schema validation for the ibm_cloud_shell_account_settings resource. |

## codeengine

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_code_engine_allowed_outbound_destination_data](data_ibm_code_engine_allowed_outbound_destination.md) | data | Reads ibm_code_engine_allowed_outbound_destination. |
| [ibm_code_engine_allowed_outbound_destination](ibm_code_engine_allowed_outbound_destination.md) | resource | Manages ibm_code_engine_allowed_outbound_destination. |
| [ibm_code_engine_app_data](data_ibm_code_engine_app.md) | data | Reads ibm_code_engine_app. |
| [ibm_code_engine_app](ibm_code_engine_app.md) | resource | Manages ibm_code_engine_app. |
| [ibm_code_engine_binding_data](data_ibm_code_engine_binding.md) | data | Schema validation for the ibm_code_engine_binding data source. |
| [ibm_code_engine_binding](ibm_code_engine_binding.md) | resource | Schema validation for the ibm_code_engine_binding resource. |
| [ibm_code_engine_build_data](data_ibm_code_engine_build.md) | data | Reads ibm_code_engine_build. |
| [ibm_code_engine_build](ibm_code_engine_build.md) | resource | Manages ibm_code_engine_build. |
| [ibm_code_engine_config_map_data](data_ibm_code_engine_config_map.md) | data | Reads ibm_code_engine_config_map. |
| [ibm_code_engine_config_map](ibm_code_engine_config_map.md) | resource | Manages ibm_code_engine_config_map. |
| [ibm_code_engine_domain_mapping_data](data_ibm_code_engine_domain_mapping.md) | data | Reads ibm_code_engine_domain_mapping. |
| [ibm_code_engine_domain_mapping](ibm_code_engine_domain_mapping.md) | resource | Manages ibm_code_engine_domain_mapping. |
| [ibm_code_engine_function_data](data_ibm_code_engine_function.md) | data | Reads ibm_code_engine_function. |
| [ibm_code_engine_function](ibm_code_engine_function.md) | resource | Manages ibm_code_engine_function. |
| [ibm_code_engine_job_data](data_ibm_code_engine_job.md) | data | Reads ibm_code_engine_job. |
| [ibm_code_engine_job](ibm_code_engine_job.md) | resource | Manages ibm_code_engine_job. |
| [ibm_code_engine_persistent_data_store_data](data_ibm_code_engine_persistent_data_store.md) | data | Reads ibm_code_engine_persistent_data_store. |
| [ibm_code_engine_persistent_data_store](ibm_code_engine_persistent_data_store.md) | resource | Manages ibm_code_engine_persistent_data_store. |
| [ibm_code_engine_project_data](data_ibm_code_engine_project.md) | data | Reads ibm_code_engine_project. |
| [ibm_code_engine_project](ibm_code_engine_project.md) | resource | Manages ibm_code_engine_project. |
| [ibm_code_engine_secret_data](data_ibm_code_engine_secret.md) | data | Reads ibm_code_engine_secret. |
| [ibm_code_engine_secret](ibm_code_engine_secret.md) | resource | Manages ibm_code_engine_secret. |

## configurationaggregator

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_config_aggregator_configurations_data](data_ibm_config_aggregator_configurations.md) | data | Schema validation for the ibm_config_aggregator_configurations data source. |
| [ibm_config_aggregator_resource_collection_status_data](data_ibm_config_aggregator_resource_collection_status.md) | data | Schema validation for the ibm_config_aggregator_resource_collection_status data source. |
| [ibm_config_aggregator_settings_data](data_ibm_config_aggregator_settings.md) | data | Schema validation for the ibm_config_aggregator_settings data source. |
| [ibm_config_aggregator_settings](ibm_config_aggregator_settings.md) | resource | Schema validation for the ibm_config_aggregator_settings resource. |

## contextbasedrestrictions

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cbr_rule_data](data_ibm_cbr_rule.md) | data | Schema validation for the ibm_cbr_rule data source. |
| [ibm_cbr_rule](ibm_cbr_rule.md) | resource | Schema validation for the ibm_cbr_rule resource. |
| [ibm_cbr_zone_data](data_ibm_cbr_zone.md) | data | Reads ibm_cbr_zone. |
| [ibm_cbr_zone](ibm_cbr_zone.md) | resource | Manages ibm_cbr_zone. |
| [ibm_cbr_zone_addresses_data](data_ibm_cbr_zone_addresses.md) | data | Schema validation for the ibm_cbr_zone_addresses data source. |
| [ibm_cbr_zone_addresses](ibm_cbr_zone_addresses.md) | resource | Schema validation for the ibm_cbr_zone_addresses resource. |

## cos

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cos_backup_policy_data](data_ibm_cos_backup_policy.md) | data | Schema validation for the ibm_cos_backup_policy data source. |
| [ibm_cos_backup_policy](ibm_cos_backup_policy.md) | resource | Schema validation for the ibm_cos_backup_policy resource. |
| [ibm_cos_backup_vault_data](data_ibm_cos_backup_vault.md) | data | Schema validation for the ibm_cos_backup_vault data source. |
| [ibm_cos_backup_vault](ibm_cos_backup_vault.md) | resource | Schema validation for the ibm_cos_backup_vault resource. |
| [ibm_cos_bucket_data](data_ibm_cos_bucket.md) | data | Schema validation for the ibm_cos_bucket data source. |
| [ibm_cos_bucket](ibm_cos_bucket.md) | resource | Schema validation for the ibm_cos_bucket resource. |
| [ibm_cos_bucket_lifecycle_configuration](ibm_cos_bucket_lifecycle_configuration.md) | resource | Schema validation for the ibm_cos_bucket_lifecycle_configuration resource. |
| [ibm_cos_bucket_object_data](data_ibm_cos_bucket_object.md) | data | Schema validation for the ibm_cos_bucket_object data source. |
| [ibm_cos_bucket_object](ibm_cos_bucket_object.md) | resource | Schema validation for the ibm_cos_bucket_object resource. |
| [ibm_cos_bucket_object_lock_configuration](ibm_cos_bucket_object_lock_configuration.md) | resource | Schema validation for the ibm_cos_bucket_object_lock_configuration resource. |
| [ibm_cos_bucket_replication_rule](ibm_cos_bucket_replication_rule.md) | resource | Schema validation for the ibm_cos_bucket_replication_rule resource. |
| [ibm_cos_bucket_website_configuration](ibm_cos_bucket_website_configuration.md) | resource | Schema validation for the ibm_cos_bucket_website_configuration resource. |

## database

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_database_data](data_ibm_database.md) | data | Reads ibm_database. |
| [ibm_database](ibm_database.md) | resource | Manages ibm_database. |
| [ibm_database_backup_data](data_ibm_database_backup.md) | data | Schema validation for the ibm_database_backup data source. |
| [ibm_database_backups_data](data_ibm_database_backups.md) | data | Schema validation for the ibm_database_backups data source. |
| [ibm_database_connection_data](data_ibm_database_connection.md) | data | Schema validation for the ibm_database_connection data source. |
| [ibm_database_point_in_time_recovery_data](data_ibm_database_point_in_time_recovery.md) | data | Schema validation for the ibm_database_point_in_time_recovery data source. |
| [ibm_database_remotes_data](data_ibm_database_remotes.md) | data | Schema validation for the ibm_database_remotes data source. |
| [ibm_database_task_data](data_ibm_database_task.md) | data | Schema validation for the ibm_database_task data source. |
| [ibm_database_tasks_data](data_ibm_database_tasks.md) | data | Schema validation for the ibm_database_tasks data source. |

## db2

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_db2_data](data_ibm_db2.md) | data | Reads ibm_db2. |
| [ibm_db2](ibm_db2.md) | resource | Manages ibm_db2. |
| [ibm_db2_allowlist_ip_data](data_ibm_db2_allowlist_ip.md) | data | Schema validation for the ibm_db2_allowlist_ip data source. |
| [ibm_db2_autoscale_data](data_ibm_db2_autoscale.md) | data | Schema validation for the ibm_db2_autoscale data source. |
| [ibm_db2_backup_data](data_ibm_db2_backup.md) | data | Schema validation for the ibm_db2_backup data source. |
| [ibm_db2_connection_info_data](data_ibm_db2_connection_info.md) | data | Schema validation for the ibm_db2_connection_info data source. |
| [ibm_db2_tuneable_param_data](data_ibm_db2_tuneable_param.md) | data | Schema validation for the ibm_db2_tuneable_param data source. |
| [ibm_db2_users_data](data_ibm_db2_users.md) | data | Schema validation for the ibm_db2_users data source. |
| [ibm_db2_whitelist_ip_data](data_ibm_db2_whitelist_ip.md) | data | Schema validation for the ibm_db2_whitelist_ip data source. |

## directlink

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_dl_export_route_filter_data](data_ibm_dl_export_route_filter.md) | data | Schema validation for the ibm_dl_export_route_filter data source. |
| [ibm_dl_export_route_filters_data](data_ibm_dl_export_route_filters.md) | data | Schema validation for the ibm_dl_export_route_filters data source. |
| [ibm_dl_gateway_data](data_ibm_dl_gateway.md) | data | Reads ibm_dl_gateway. |
| [ibm_dl_gateway](ibm_dl_gateway.md) | resource | Manages ibm_dl_gateway. |
| [ibm_dl_gateway_action](ibm_dl_gateway_action.md) | resource | Manages ibm_dl_gateway_action. |
| [ibm_dl_gateway_macsec_cak_data](data_ibm_dl_gateway_macsec_cak.md) | data | Reads ibm_dl_gateway_macsec_cak. |
| [ibm_dl_gateway_macsec_cak](ibm_dl_gateway_macsec_cak.md) | resource | Manages ibm_dl_gateway_macsec_cak. |
| [ibm_dl_gateway_macsec_caks_data](data_ibm_dl_gateway_macsec_caks.md) | data | Schema validation for the ibm_dl_gateway_macsec_caks data source. |
| [ibm_dl_gateway_macsec_config_data](data_ibm_dl_gateway_macsec_config.md) | data | Schema validation for the ibm_dl_gateway_macsec_config data source. |
| [ibm_dl_gateway_macsec_config](ibm_dl_gateway_macsec_config.md) | resource | Schema validation for the ibm_dl_gateway_macsec_config resource. |
| [ibm_dl_gateways_data](data_ibm_dl_gateways.md) | data | Schema validation for the ibm_dl_gateways data source. |
| [ibm_dl_import_route_filter_data](data_ibm_dl_import_route_filter.md) | data | Schema validation for the ibm_dl_import_route_filter data source. |
| [ibm_dl_import_route_filters_data](data_ibm_dl_import_route_filters.md) | data | Schema validation for the ibm_dl_import_route_filters data source. |
| [ibm_dl_locations_data](data_ibm_dl_locations.md) | data | Schema validation for the ibm_dl_locations data source. |
| [ibm_dl_offering_speeds_data](data_ibm_dl_offering_speeds.md) | data | Schema validation for the ibm_dl_offering_speeds data source. |
| [ibm_dl_port_data](data_ibm_dl_port.md) | data | Schema validation for the ibm_dl_port data source. |
| [ibm_dl_ports_data](data_ibm_dl_ports.md) | data | Schema validation for the ibm_dl_ports data source. |
| [ibm_dl_provider_gateway](ibm_dl_provider_gateway.md) | resource | Manages ibm_dl_provider_gateway. |
| [ibm_dl_provider_gateways_data](data_ibm_dl_provider_gateways.md) | data | Schema validation for the ibm_dl_provider_gateways data source. |
| [ibm_dl_provider_ports_data](data_ibm_dl_provider_ports.md) | data | Schema validation for the ibm_dl_provider_ports data source. |
| [ibm_dl_route_report_data](data_ibm_dl_route_report.md) | data | Schema validation for the ibm_dl_route_report data source. |
| [ibm_dl_route_report](ibm_dl_route_report.md) | resource | Schema validation for the ibm_dl_route_report resource. |
| [ibm_dl_route_reports_data](data_ibm_dl_route_reports.md) | data | Schema validation for the ibm_dl_route_reports data source. |
| [ibm_dl_routers_data](data_ibm_dl_routers.md) | data | Schema validation for the ibm_dl_routers data source. |
| [ibm_dl_virtual_connection](ibm_dl_virtual_connection.md) | resource | Manages ibm_dl_virtual_connection. |

## dnsservices

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_dns_custom_resolver](ibm_dns_custom_resolver.md) | resource | Manages ibm_dns_custom_resolver. |
| [ibm_dns_custom_resolver_forwarding_rule](ibm_dns_custom_resolver_forwarding_rule.md) | resource | Schema validation for the ibm_dns_custom_resolver_forwarding_rule resource. |
| [ibm_dns_custom_resolver_forwarding_rules_data](data_ibm_dns_custom_resolver_forwarding_rules.md) | data | Schema validation for the ibm_dns_custom_resolver_forwarding_rules data source. |
| [ibm_dns_custom_resolver_secondary_zone](ibm_dns_custom_resolver_secondary_zone.md) | resource | Schema validation for the ibm_dns_custom_resolver_secondary_zone resource. |
| [ibm_dns_custom_resolver_secondary_zones_data](data_ibm_dns_custom_resolver_secondary_zones.md) | data | Schema validation for the ibm_dns_custom_resolver_secondary_zones data source. |
| [ibm_dns_custom_resolvers_data](data_ibm_dns_custom_resolvers.md) | data | Schema validation for the ibm_dns_custom_resolvers data source. |
| [ibm_dns_glb](ibm_dns_glb.md) | resource | Manages ibm_dns_glb. |
| [ibm_dns_glb_monitor](ibm_dns_glb_monitor.md) | resource | Manages ibm_dns_glb_monitor. |
| [ibm_dns_glb_monitors_data](data_ibm_dns_glb_monitors.md) | data | Schema validation for the ibm_dns_glb_monitors data source. |
| [ibm_dns_glb_pool](ibm_dns_glb_pool.md) | resource | Manages ibm_dns_glb_pool. |
| [ibm_dns_glb_pools_data](data_ibm_dns_glb_pools.md) | data | Schema validation for the ibm_dns_glb_pools data source. |
| [ibm_dns_glbs_data](data_ibm_dns_glbs.md) | data | Schema validation for the ibm_dns_glbs data source. |
| [ibm_dns_linked_zone](ibm_dns_linked_zone.md) | resource | Manages ibm_dns_linked_zone. |
| [ibm_dns_permitted_network](ibm_dns_permitted_network.md) | resource | Schema validation for the ibm_dns_permitted_network resource. |
| [ibm_dns_permitted_networks_data](data_ibm_dns_permitted_networks.md) | data | Schema validation for the ibm_dns_permitted_networks data source. |
| [ibm_dns_resource_record](ibm_dns_resource_record.md) | resource | Manages ibm_dns_resource_record. |
| [ibm_dns_resource_records_data](data_ibm_dns_resource_records.md) | data | Schema validation for the ibm_dns_resource_records data source. |
| [ibm_dns_zone](ibm_dns_zone.md) | resource | Manages ibm_dns_zone. |
| [ibm_dns_zones_data](data_ibm_dns_zones.md) | data | Schema validation for the ibm_dns_zones data source. |

## drautomationservice

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_pdr_dr_locations_data](data_ibm_pdr_dr_locations.md) | data | Schema validation for the ibm_pdr_dr_locations data source. |
| [ibm_pdr_dr_summary_response_data](data_ibm_pdr_dr_summary_response.md) | data | Schema validation for the ibm_pdr_dr_summary_response data source. |
| [ibm_pdr_event_data](data_ibm_pdr_event.md) | data | Schema validation for the ibm_pdr_event data source. |
| [ibm_pdr_events_data](data_ibm_pdr_events.md) | data | Schema validation for the ibm_pdr_events data source. |
| [ibm_pdr_get_dr_locations_data](data_ibm_pdr_get_dr_locations.md) | data | Schema validation for the ibm_pdr_get_dr_locations data source. |
| [ibm_pdr_get_dr_summary_response_data](data_ibm_pdr_get_dr_summary_response.md) | data | Schema validation for the ibm_pdr_get_dr_summary_response data source. |
| [ibm_pdr_get_event_data](data_ibm_pdr_get_event.md) | data | Schema validation for the ibm_pdr_get_event data source. |
| [ibm_pdr_get_events_data](data_ibm_pdr_get_events.md) | data | Schema validation for the ibm_pdr_get_events data source. |
| [ibm_pdr_get_grs_location_pairs_data](data_ibm_pdr_get_grs_location_pairs.md) | data | Schema validation for the ibm_pdr_get_grs_location_pairs data source. |
| [ibm_pdr_get_machine_types_data](data_ibm_pdr_get_machine_types.md) | data | Schema validation for the ibm_pdr_get_machine_types data source. |
| [ibm_pdr_get_managed_vm_list_data](data_ibm_pdr_get_managed_vm_list.md) | data | Schema validation for the ibm_pdr_get_managed_vm_list data source. |
| [ibm_pdr_get_powervs_workspace_data](data_ibm_pdr_get_powervs_workspace.md) | data | Schema validation for the ibm_pdr_get_powervs_workspace data source. |
| [ibm_pdr_grs_location_pairs_data](data_ibm_pdr_grs_location_pairs.md) | data | Schema validation for the ibm_pdr_grs_location_pairs data source. |
| [ibm_pdr_last_operation_data](data_ibm_pdr_last_operation.md) | data | Schema validation for the ibm_pdr_last_operation data source. |
| [ibm_pdr_machine_types_data](data_ibm_pdr_machine_types.md) | data | Schema validation for the ibm_pdr_machine_types data source. |
| [ibm_pdr_managed_vm_list_data](data_ibm_pdr_managed_vm_list.md) | data | Schema validation for the ibm_pdr_managed_vm_list data source. |
| [ibm_pdr_managedr](ibm_pdr_managedr.md) | resource | Schema validation for the ibm_pdr_managedr resource. |
| [ibm_pdr_powervs_workspaces_data](data_ibm_pdr_powervs_workspaces.md) | data | Schema validation for the ibm_pdr_powervs_workspaces data source. |
| [ibm_pdr_validate_apikey](ibm_pdr_validate_apikey.md) | resource | Schema validation for the ibm_pdr_validate_apikey resource. |

## enterprise

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_enterprise](ibm_enterprise.md) | resource | Manages ibm_enterprise. |
| [ibm_enterprise_account](ibm_enterprise_account.md) | resource | Manages ibm_enterprise_account. |
| [ibm_enterprise_account_group](ibm_enterprise_account_group.md) | resource | Manages ibm_enterprise_account_group. |
| [ibm_enterprise_account_groups_data](data_ibm_enterprise_account_groups.md) | data | Reads ibm_enterprise_account_groups. |
| [ibm_enterprise_accounts_data](data_ibm_enterprise_accounts.md) | data | Reads ibm_enterprise_accounts. |
| [ibm_enterprises_data](data_ibm_enterprises.md) | data | Reads ibm_enterprises. |

## eventnotification

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_en_app_configuration_template_data](data_ibm_en_app_configuration_template.md) | data | Reads ibm_en_app_configuration_template. |
| [ibm_en_app_configuration_template](ibm_en_app_configuration_template.md) | resource | Manages ibm_en_app_configuration_template. |
| [ibm_en_bounce_metrics_data](data_ibm_en_bounce_metrics.md) | data | Schema validation for the ibm_en_bounce_metrics data source. |
| [ibm_en_code_engine_template_data](data_ibm_en_code_engine_template.md) | data | Reads ibm_en_code_engine_template. |
| [ibm_en_code_engine_template](ibm_en_code_engine_template.md) | resource | Manages ibm_en_code_engine_template. |
| [ibm_en_destination_android_data](data_ibm_en_destination_android.md) | data | Reads ibm_en_destination_android. |
| [ibm_en_destination_android](ibm_en_destination_android.md) | resource | Manages ibm_en_destination_android. |
| [ibm_en_destination_app_configuration_data](data_ibm_en_destination_app_configuration.md) | data | Reads ibm_en_destination_app_configuration. |
| [ibm_en_destination_app_configuration](ibm_en_destination_app_configuration.md) | resource | Manages ibm_en_destination_app_configuration. |
| [ibm_en_destination_ce_data](data_ibm_en_destination_ce.md) | data | Reads ibm_en_destination_ce. |
| [ibm_en_destination_ce](ibm_en_destination_ce.md) | resource | Manages ibm_en_destination_ce. |
| [ibm_en_destination_chrome_data](data_ibm_en_destination_chrome.md) | data | Reads ibm_en_destination_chrome. |
| [ibm_en_destination_chrome](ibm_en_destination_chrome.md) | resource | Manages ibm_en_destination_chrome. |
| [ibm_en_destination_cos_data](data_ibm_en_destination_cos.md) | data | Reads ibm_en_destination_cos. |
| [ibm_en_destination_cos](ibm_en_destination_cos.md) | resource | Manages ibm_en_destination_cos. |
| [ibm_en_destination_custom_email_data](data_ibm_en_destination_custom_email.md) | data | Reads ibm_en_destination_custom_email. |
| [ibm_en_destination_custom_email](ibm_en_destination_custom_email.md) | resource | Manages ibm_en_destination_custom_email. |
| [ibm_en_destination_custom_sms_data](data_ibm_en_destination_custom_sms.md) | data | Reads ibm_en_destination_custom_sms. |
| [ibm_en_destination_custom_sms](ibm_en_destination_custom_sms.md) | resource | Manages ibm_en_destination_custom_sms. |
| [ibm_en_destination_event_streams_data](data_ibm_en_destination_event_streams.md) | data | Reads ibm_en_destination_event_streams. |
| [ibm_en_destination_event_streams](ibm_en_destination_event_streams.md) | resource | Manages ibm_en_destination_event_streams. |
| [ibm_en_destination_firefox_data](data_ibm_en_destination_firefox.md) | data | Reads ibm_en_destination_firefox. |
| [ibm_en_destination_firefox](ibm_en_destination_firefox.md) | resource | Manages ibm_en_destination_firefox. |
| [ibm_en_destination_huawei_data](data_ibm_en_destination_huawei.md) | data | Reads ibm_en_destination_huawei. |
| [ibm_en_destination_huawei](ibm_en_destination_huawei.md) | resource | Manages ibm_en_destination_huawei. |
| [ibm_en_destination_ios_data](data_ibm_en_destination_ios.md) | data | Reads ibm_en_destination_ios. |
| [ibm_en_destination_ios](ibm_en_destination_ios.md) | resource | Manages ibm_en_destination_ios. |
| [ibm_en_destination_msteams_data](data_ibm_en_destination_msteams.md) | data | Reads ibm_en_destination_msteams. |
| [ibm_en_destination_msteams](ibm_en_destination_msteams.md) | resource | Manages ibm_en_destination_msteams. |
| [ibm_en_destination_pagerduty_data](data_ibm_en_destination_pagerduty.md) | data | Reads ibm_en_destination_pagerduty. |
| [ibm_en_destination_pagerduty](ibm_en_destination_pagerduty.md) | resource | Manages ibm_en_destination_pagerduty. |
| [ibm_en_destination_safari_data](data_ibm_en_destination_safari.md) | data | Reads ibm_en_destination_safari. |
| [ibm_en_destination_safari](ibm_en_destination_safari.md) | resource | Manages ibm_en_destination_safari. |
| [ibm_en_destination_slack_data](data_ibm_en_destination_slack.md) | data | Reads ibm_en_destination_slack. |
| [ibm_en_destination_slack](ibm_en_destination_slack.md) | resource | Manages ibm_en_destination_slack. |
| [ibm_en_destination_sn_data](data_ibm_en_destination_sn.md) | data | Reads ibm_en_destination_sn. |
| [ibm_en_destination_sn](ibm_en_destination_sn.md) | resource | Manages ibm_en_destination_sn. |
| [ibm_en_destination_webhook_data](data_ibm_en_destination_webhook.md) | data | Reads ibm_en_destination_webhook. |
| [ibm_en_destination_webhook](ibm_en_destination_webhook.md) | resource | Manages ibm_en_destination_webhook. |
| [ibm_en_destinations_data](data_ibm_en_destinations.md) | data | Schema validation for the ibm_en_destinations data source. |
| [ibm_en_email_template_data](data_ibm_en_email_template.md) | data | Reads ibm_en_email_template. |
| [ibm_en_email_template](ibm_en_email_template.md) | resource | Manages ibm_en_email_template. |
| [ibm_en_email_templates_data](data_ibm_en_email_templates.md) | data | Schema validation for the ibm_en_email_templates data source. |
| [ibm_en_event_streams_template_data](data_ibm_en_event_streams_template.md) | data | Reads ibm_en_event_streams_template. |
| [ibm_en_event_streams_template](ibm_en_event_streams_template.md) | resource | Manages ibm_en_event_streams_template. |
| [ibm_en_ibmsource](ibm_en_ibmsource.md) | resource | Schema validation for the ibm_en_ibmsource resource. |
| [ibm_en_integration_data](data_ibm_en_integration.md) | data | Schema validation for the ibm_en_integration data source. |
| [ibm_en_integration](ibm_en_integration.md) | resource | Schema validation for the ibm_en_integration resource. |
| [ibm_en_integration_cos_data](data_ibm_en_integration_cos.md) | data | Schema validation for the ibm_en_integration_cos data source. |
| [ibm_en_integration_cos](ibm_en_integration_cos.md) | resource | Schema validation for the ibm_en_integration_cos resource. |
| [ibm_en_integrations_data](data_ibm_en_integrations.md) | data | Schema validation for the ibm_en_integrations data source. |
| [ibm_en_metrics_data](data_ibm_en_metrics.md) | data | Schema validation for the ibm_en_metrics data source. |
| [ibm_en_pagerduty_template_data](data_ibm_en_pagerduty_template.md) | data | Reads ibm_en_pagerduty_template. |
| [ibm_en_pagerduty_template](ibm_en_pagerduty_template.md) | resource | Manages ibm_en_pagerduty_template. |
| [ibm_en_pre_defined_template_data](data_ibm_en_pre_defined_template.md) | data | Reads ibm_en_pre_defined_template. |
| [ibm_en_pre_defined_templates_data](data_ibm_en_pre_defined_templates.md) | data | Schema validation for the ibm_en_pre_defined_templates data source. |
| [ibm_en_slack_template_data](data_ibm_en_slack_template.md) | data | Reads ibm_en_slack_template. |
| [ibm_en_slack_template](ibm_en_slack_template.md) | resource | Manages ibm_en_slack_template. |
| [ibm_en_smtp_allowed_ips_data](data_ibm_en_smtp_allowed_ips.md) | data | Schema validation for the ibm_en_smtp_allowed_ips data source. |
| [ibm_en_smtp_configuration_data](data_ibm_en_smtp_configuration.md) | data | Reads ibm_en_smtp_configuration. |
| [ibm_en_smtp_configuration](ibm_en_smtp_configuration.md) | resource | Manages ibm_en_smtp_configuration. |
| [ibm_en_smtp_configurations_data](data_ibm_en_smtp_configurations.md) | data | Schema validation for the ibm_en_smtp_configurations data source. |
| [ibm_en_smtp_setting](ibm_en_smtp_setting.md) | resource | Schema validation for the ibm_en_smtp_setting resource. |
| [ibm_en_smtp_user_data](data_ibm_en_smtp_user.md) | data | Schema validation for the ibm_en_smtp_user data source. |
| [ibm_en_smtp_user](ibm_en_smtp_user.md) | resource | Schema validation for the ibm_en_smtp_user resource. |
| [ibm_en_smtp_users_data](data_ibm_en_smtp_users.md) | data | Schema validation for the ibm_en_smtp_users data source. |
| [ibm_en_source_data](data_ibm_en_source.md) | data | Reads ibm_en_source. |
| [ibm_en_source](ibm_en_source.md) | resource | Manages ibm_en_source. |
| [ibm_en_sources_data](data_ibm_en_sources.md) | data | Schema validation for the ibm_en_sources data source. |
| [ibm_en_subscription_android_data](data_ibm_en_subscription_android.md) | data | Reads ibm_en_subscription_android. |
| [ibm_en_subscription_android](ibm_en_subscription_android.md) | resource | Manages ibm_en_subscription_android. |
| [ibm_en_subscription_app_configuration_data](data_ibm_en_subscription_app_configuration.md) | data | Reads ibm_en_subscription_app_configuration. |
| [ibm_en_subscription_app_configuration](ibm_en_subscription_app_configuration.md) | resource | Manages ibm_en_subscription_app_configuration. |
| [ibm_en_subscription_ce_data](data_ibm_en_subscription_ce.md) | data | Reads ibm_en_subscription_ce. |
| [ibm_en_subscription_ce](ibm_en_subscription_ce.md) | resource | Manages ibm_en_subscription_ce. |
| [ibm_en_subscription_chrome_data](data_ibm_en_subscription_chrome.md) | data | Reads ibm_en_subscription_chrome. |
| [ibm_en_subscription_chrome](ibm_en_subscription_chrome.md) | resource | Manages ibm_en_subscription_chrome. |
| [ibm_en_subscription_cos_data](data_ibm_en_subscription_cos.md) | data | Reads ibm_en_subscription_cos. |
| [ibm_en_subscription_cos](ibm_en_subscription_cos.md) | resource | Manages ibm_en_subscription_cos. |
| [ibm_en_subscription_custom_email_data](data_ibm_en_subscription_custom_email.md) | data | Reads ibm_en_subscription_custom_email. |
| [ibm_en_subscription_custom_email](ibm_en_subscription_custom_email.md) | resource | Manages ibm_en_subscription_custom_email. |
| [ibm_en_subscription_custom_sms_data](data_ibm_en_subscription_custom_sms.md) | data | Reads ibm_en_subscription_custom_sms. |
| [ibm_en_subscription_custom_sms](ibm_en_subscription_custom_sms.md) | resource | Manages ibm_en_subscription_custom_sms. |
| [ibm_en_subscription_email_data](data_ibm_en_subscription_email.md) | data | Reads ibm_en_subscription_email. |
| [ibm_en_subscription_email](ibm_en_subscription_email.md) | resource | Manages ibm_en_subscription_email. |
| [ibm_en_subscription_event_streams_data](data_ibm_en_subscription_event_streams.md) | data | Reads ibm_en_subscription_event_streams. |
| [ibm_en_subscription_event_streams](ibm_en_subscription_event_streams.md) | resource | Manages ibm_en_subscription_event_streams. |
| [ibm_en_subscription_firefox_data](data_ibm_en_subscription_firefox.md) | data | Reads ibm_en_subscription_firefox. |
| [ibm_en_subscription_firefox](ibm_en_subscription_firefox.md) | resource | Manages ibm_en_subscription_firefox. |
| [ibm_en_subscription_huawei_data](data_ibm_en_subscription_huawei.md) | data | Reads ibm_en_subscription_huawei. |
| [ibm_en_subscription_huawei](ibm_en_subscription_huawei.md) | resource | Manages ibm_en_subscription_huawei. |
| [ibm_en_subscription_ios_data](data_ibm_en_subscription_ios.md) | data | Reads ibm_en_subscription_ios. |
| [ibm_en_subscription_ios](ibm_en_subscription_ios.md) | resource | Manages ibm_en_subscription_ios. |
| [ibm_en_subscription_msteams_data](data_ibm_en_subscription_msteams.md) | data | Reads ibm_en_subscription_msteams. |
| [ibm_en_subscription_msteams](ibm_en_subscription_msteams.md) | resource | Manages ibm_en_subscription_msteams. |
| [ibm_en_subscription_pagerduty_data](data_ibm_en_subscription_pagerduty.md) | data | Reads ibm_en_subscription_pagerduty. |
| [ibm_en_subscription_pagerduty](ibm_en_subscription_pagerduty.md) | resource | Manages ibm_en_subscription_pagerduty. |
| [ibm_en_subscription_safari_data](data_ibm_en_subscription_safari.md) | data | Reads ibm_en_subscription_safari. |
| [ibm_en_subscription_safari](ibm_en_subscription_safari.md) | resource | Manages ibm_en_subscription_safari. |
| [ibm_en_subscription_scheduler_data](data_ibm_en_subscription_scheduler.md) | data | Reads ibm_en_subscription_scheduler. |
| [ibm_en_subscription_scheduler](ibm_en_subscription_scheduler.md) | resource | Manages ibm_en_subscription_scheduler. |
| [ibm_en_subscription_slack_data](data_ibm_en_subscription_slack.md) | data | Reads ibm_en_subscription_slack. |
| [ibm_en_subscription_slack](ibm_en_subscription_slack.md) | resource | Manages ibm_en_subscription_slack. |
| [ibm_en_subscription_sms_data](data_ibm_en_subscription_sms.md) | data | Reads ibm_en_subscription_sms. |
| [ibm_en_subscription_sms](ibm_en_subscription_sms.md) | resource | Manages ibm_en_subscription_sms. |
| [ibm_en_subscription_sn_data](data_ibm_en_subscription_sn.md) | data | Reads ibm_en_subscription_sn. |
| [ibm_en_subscription_sn](ibm_en_subscription_sn.md) | resource | Manages ibm_en_subscription_sn. |
| [ibm_en_subscription_webhook_data](data_ibm_en_subscription_webhook.md) | data | Reads ibm_en_subscription_webhook. |
| [ibm_en_subscription_webhook](ibm_en_subscription_webhook.md) | resource | Manages ibm_en_subscription_webhook. |
| [ibm_en_subscriptions_data](data_ibm_en_subscriptions.md) | data | Schema validation for the ibm_en_subscriptions data source. |
| [ibm_en_topic_data](data_ibm_en_topic.md) | data | Reads ibm_en_topic. |
| [ibm_en_topic](ibm_en_topic.md) | resource | Manages ibm_en_topic. |
| [ibm_en_topics_data](data_ibm_en_topics.md) | data | Schema validation for the ibm_en_topics data source. |
| [ibm_en_webhook_template_data](data_ibm_en_webhook_template.md) | data | Reads ibm_en_webhook_template. |
| [ibm_en_webhook_template](ibm_en_webhook_template.md) | resource | Manages ibm_en_webhook_template. |

## eventstreams

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_event_streams_mirroring_config_data](data_ibm_event_streams_mirroring_config.md) | data | Schema validation for the ibm_event_streams_mirroring_config data source. |
| [ibm_event_streams_mirroring_config](ibm_event_streams_mirroring_config.md) | resource | Schema validation for the ibm_event_streams_mirroring_config resource. |
| [ibm_event_streams_quota_data](data_ibm_event_streams_quota.md) | data | Schema validation for the ibm_event_streams_quota data source. |
| [ibm_event_streams_quota](ibm_event_streams_quota.md) | resource | Schema validation for the ibm_event_streams_quota resource. |
| [ibm_event_streams_schema_data](data_ibm_event_streams_schema.md) | data | Schema validation for the ibm_event_streams_schema data source. |
| [ibm_event_streams_schema](ibm_event_streams_schema.md) | resource | Schema validation for the ibm_event_streams_schema resource. |
| [ibm_event_streams_schema_global_rule_data](data_ibm_event_streams_schema_global_rule.md) | data | Schema validation for the ibm_event_streams_schema_global_rule data source. |
| [ibm_event_streams_schema_global_rule](ibm_event_streams_schema_global_rule.md) | resource | Schema validation for the ibm_event_streams_schema_global_rule resource. |
| [ibm_event_streams_topic_data](data_ibm_event_streams_topic.md) | data | Reads ibm_event_streams_topic. |
| [ibm_event_streams_topic](ibm_event_streams_topic.md) | resource | Manages ibm_event_streams_topic. |

## functions

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_function_action_data](data_ibm_function_action.md) | data | Reads ibm_function_action. |
| [ibm_function_action](ibm_function_action.md) | resource | Manages ibm_function_action. |
| [ibm_function_namespace_data](data_ibm_function_namespace.md) | data | Reads ibm_function_namespace. |
| [ibm_function_namespace](ibm_function_namespace.md) | resource | Manages ibm_function_namespace. |
| [ibm_function_package_data](data_ibm_function_package.md) | data | Reads ibm_function_package. |
| [ibm_function_package](ibm_function_package.md) | resource | Manages ibm_function_package. |
| [ibm_function_rule_data](data_ibm_function_rule.md) | data | Reads ibm_function_rule. |
| [ibm_function_rule](ibm_function_rule.md) | resource | Manages ibm_function_rule. |
| [ibm_function_trigger_data](data_ibm_function_trigger.md) | data | Reads ibm_function_trigger. |
| [ibm_function_trigger](ibm_function_trigger.md) | resource | Manages ibm_function_trigger. |

## globaltagging

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_iam_access_tag_data](data_ibm_iam_access_tag.md) | data | Reads ibm_iam_access_tag. |
| [ibm_iam_access_tag](ibm_iam_access_tag.md) | resource | Manages ibm_iam_access_tag. |
| [ibm_resource_tag_data](data_ibm_resource_tag.md) | data | Schema validation for the ibm_resource_tag data source. |
| [ibm_resource_tag](ibm_resource_tag.md) | resource | Schema validation for the ibm_resource_tag resource. |

## hpcs

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_hpcs_data](data_ibm_hpcs.md) | data | Reads ibm_hpcs. |
| [ibm_hpcs](ibm_hpcs.md) | resource | Manages ibm_hpcs. |
| [ibm_hpcs_key_template_data](data_ibm_hpcs_key_template.md) | data | Reads ibm_hpcs_key_template. |
| [ibm_hpcs_key_template](ibm_hpcs_key_template.md) | resource | Manages ibm_hpcs_key_template. |
| [ibm_hpcs_keystore_data](data_ibm_hpcs_keystore.md) | data | Reads ibm_hpcs_keystore. |
| [ibm_hpcs_keystore](ibm_hpcs_keystore.md) | resource | Manages ibm_hpcs_keystore. |
| [ibm_hpcs_managed_key_data](data_ibm_hpcs_managed_key.md) | data | Schema validation for the ibm_hpcs_managed_key data source. |
| [ibm_hpcs_managed_key](ibm_hpcs_managed_key.md) | resource | Schema validation for the ibm_hpcs_managed_key resource. |
| [ibm_hpcs_vault_data](data_ibm_hpcs_vault.md) | data | Reads ibm_hpcs_vault. |
| [ibm_hpcs_vault](ibm_hpcs_vault.md) | resource | Manages ibm_hpcs_vault. |

## iamaccessgroup

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_iam_access_group_data](data_ibm_iam_access_group.md) | data | Schema validation for the ibm_iam_access_group data source. |
| [ibm_iam_access_group](ibm_iam_access_group.md) | resource | Manages ibm_iam_access_group. |
| [ibm_iam_access_group_account_settings](ibm_iam_access_group_account_settings.md) | resource | Schema validation for the ibm_iam_access_group_account_settings resource. |
| [ibm_iam_access_group_dynamic_rule](ibm_iam_access_group_dynamic_rule.md) | resource | Manages ibm_iam_access_group_dynamic_rule. |
| [ibm_iam_access_group_members](ibm_iam_access_group_members.md) | resource | Schema validation for the ibm_iam_access_group_members resource. |
| [ibm_iam_access_group_template](ibm_iam_access_group_template.md) | resource | Manages ibm_iam_access_group_template. |
| [ibm_iam_access_group_template_assignment_data](data_ibm_iam_access_group_template_assignment.md) | data | Schema validation for the ibm_iam_access_group_template_assignment data source. |
| [ibm_iam_access_group_template_assignment](ibm_iam_access_group_template_assignment.md) | resource | Schema validation for the ibm_iam_access_group_template_assignment resource. |
| [ibm_iam_access_group_template_version](ibm_iam_access_group_template_version.md) | resource | Manages ibm_iam_access_group_template_version. |
| [ibm_iam_access_group_template_versions_data](data_ibm_iam_access_group_template_versions.md) | data | Schema validation for the ibm_iam_access_group_template_versions data source. |

## iamidentity

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_iam_account_settings_data](data_ibm_iam_account_settings.md) | data | Schema validation for the ibm_iam_account_settings data source. |
| [ibm_iam_account_settings](ibm_iam_account_settings.md) | resource | Schema validation for the ibm_iam_account_settings resource. |
| [ibm_iam_account_settings_template_data](data_ibm_iam_account_settings_template.md) | data | Reads ibm_iam_account_settings_template. |
| [ibm_iam_account_settings_template](ibm_iam_account_settings_template.md) | resource | Manages ibm_iam_account_settings_template. |
| [ibm_iam_account_settings_template_assignment_data](data_ibm_iam_account_settings_template_assignment.md) | data | Schema validation for the ibm_iam_account_settings_template_assignment data source. |
| [ibm_iam_account_settings_template_assignment](ibm_iam_account_settings_template_assignment.md) | resource | Schema validation for the ibm_iam_account_settings_template_assignment resource. |
| [ibm_iam_api_key_data](data_ibm_iam_api_key.md) | data | Reads ibm_iam_api_key. |
| [ibm_iam_api_key](ibm_iam_api_key.md) | resource | Manages ibm_iam_api_key. |
| [ibm_iam_auth_token_data](data_ibm_iam_auth_token.md) | data | Schema validation for the ibm_iam_auth_token data source. |
| [ibm_iam_effective_account_settings_data](data_ibm_iam_effective_account_settings.md) | data | Schema validation for the ibm_iam_effective_account_settings data source. |
| [ibm_iam_identity_preference_data](data_ibm_iam_identity_preference.md) | data | Schema validation for the ibm_iam_identity_preference data source. |
| [ibm_iam_identity_preference](ibm_iam_identity_preference.md) | resource | Schema validation for the ibm_iam_identity_preference resource. |
| [ibm_iam_identity_preferences_data](data_ibm_iam_identity_preferences.md) | data | Schema validation for the ibm_iam_identity_preferences data source. |
| [ibm_iam_idp_data](data_ibm_iam_idp.md) | data | Reads ibm_iam_idp. |
| [ibm_iam_idp](ibm_iam_idp.md) | resource | Manages ibm_iam_idp. |
| [ibm_iam_idp_account_setting](ibm_iam_idp_account_setting.md) | resource | Schema validation for the ibm_iam_idp_account_setting resource. |
| [ibm_iam_idp_account_settings_data](data_ibm_iam_idp_account_settings.md) | data | Schema validation for the ibm_iam_idp_account_settings data source. |
| [ibm_iam_idps_data](data_ibm_iam_idps.md) | data | Schema validation for the ibm_iam_idps data source. |
| [ibm_iam_service_api_key](ibm_iam_service_api_key.md) | resource | Manages ibm_iam_service_api_key. |
| [ibm_iam_service_id_data](data_ibm_iam_service_id.md) | data | Reads ibm_iam_service_id. |
| [ibm_iam_service_id](ibm_iam_service_id.md) | resource | Manages ibm_iam_service_id. |
| [ibm_iam_serviceid_group_data](data_ibm_iam_serviceid_group.md) | data | Reads ibm_iam_serviceid_group. |
| [ibm_iam_serviceid_group](ibm_iam_serviceid_group.md) | resource | Manages ibm_iam_serviceid_group. |
| [ibm_iam_trusted_profile_data](data_ibm_iam_trusted_profile.md) | data | Reads ibm_iam_trusted_profile. |
| [ibm_iam_trusted_profile](ibm_iam_trusted_profile.md) | resource | Manages ibm_iam_trusted_profile. |
| [ibm_iam_trusted_profile_claim_rule_data](data_ibm_iam_trusted_profile_claim_rule.md) | data | Reads ibm_iam_trusted_profile_claim_rule. |
| [ibm_iam_trusted_profile_claim_rule](ibm_iam_trusted_profile_claim_rule.md) | resource | Manages ibm_iam_trusted_profile_claim_rule. |
| [ibm_iam_trusted_profile_claim_rules_data](data_ibm_iam_trusted_profile_claim_rules.md) | data | Schema validation for the ibm_iam_trusted_profile_claim_rules data source. |
| [ibm_iam_trusted_profile_identities_data](data_ibm_iam_trusted_profile_identities.md) | data | Schema validation for the ibm_iam_trusted_profile_identities data source. |
| [ibm_iam_trusted_profile_identities](ibm_iam_trusted_profile_identities.md) | resource | Schema validation for the ibm_iam_trusted_profile_identities resource. |
| [ibm_iam_trusted_profile_identity_data](data_ibm_iam_trusted_profile_identity.md) | data | Schema validation for the ibm_iam_trusted_profile_identity data source. |
| [ibm_iam_trusted_profile_identity](ibm_iam_trusted_profile_identity.md) | resource | Schema validation for the ibm_iam_trusted_profile_identity resource. |
| [ibm_iam_trusted_profile_link_data](data_ibm_iam_trusted_profile_link.md) | data | Reads ibm_iam_trusted_profile_link. |
| [ibm_iam_trusted_profile_link](ibm_iam_trusted_profile_link.md) | resource | Manages ibm_iam_trusted_profile_link. |
| [ibm_iam_trusted_profile_links_data](data_ibm_iam_trusted_profile_links.md) | data | Schema validation for the ibm_iam_trusted_profile_links data source. |
| [ibm_iam_trusted_profile_template_data](data_ibm_iam_trusted_profile_template.md) | data | Reads ibm_iam_trusted_profile_template. |
| [ibm_iam_trusted_profile_template](ibm_iam_trusted_profile_template.md) | resource | Manages ibm_iam_trusted_profile_template. |
| [ibm_iam_trusted_profile_template_assignment_data](data_ibm_iam_trusted_profile_template_assignment.md) | data | Schema validation for the ibm_iam_trusted_profile_template_assignment data source. |
| [ibm_iam_trusted_profile_template_assignment](ibm_iam_trusted_profile_template_assignment.md) | resource | Schema validation for the ibm_iam_trusted_profile_template_assignment resource. |
| [ibm_iam_trusted_profiles_data](data_ibm_iam_trusted_profiles.md) | data | Reads ibm_iam_trusted_profiles. |
| [ibm_iam_user_mfa_enrollments_data](data_ibm_iam_user_mfa_enrollments.md) | data | Schema validation for the ibm_iam_user_mfa_enrollments data source. |
| [ibm_iam_user_profile_data](data_ibm_iam_user_profile.md) | data | Schema validation for the ibm_iam_user_profile data source. |
| [ibm_iam_user_settings](ibm_iam_user_settings.md) | resource | Schema validation for the ibm_iam_user_settings resource. |
| [ibm_iam_users_data](data_ibm_iam_users.md) | data | Schema validation for the ibm_iam_users data source. |

## iampolicy

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_iam_access_group_policy_data](data_ibm_iam_access_group_policy.md) | data | Schema validation for the ibm_iam_access_group_policy data source. |
| [ibm_iam_access_group_policy](ibm_iam_access_group_policy.md) | resource | Schema validation for the ibm_iam_access_group_policy resource. |
| [ibm_iam_account_settings_external_interaction_data](data_ibm_iam_account_settings_external_interaction.md) | data | Schema validation for the ibm_iam_account_settings_external_interaction data source. |
| [ibm_iam_account_settings_external_interaction](ibm_iam_account_settings_external_interaction.md) | resource | Schema validation for the ibm_iam_account_settings_external_interaction resource. |
| [ibm_iam_action_control_assignment_data](data_ibm_iam_action_control_assignment.md) | data | Schema validation for the ibm_iam_action_control_assignment data source. |
| [ibm_iam_action_control_assignment](ibm_iam_action_control_assignment.md) | resource | Schema validation for the ibm_iam_action_control_assignment resource. |
| [ibm_iam_action_control_assignments_data](data_ibm_iam_action_control_assignments.md) | data | Schema validation for the ibm_iam_action_control_assignments data source. |
| [ibm_iam_action_control_template_data](data_ibm_iam_action_control_template.md) | data | Schema validation for the ibm_iam_action_control_template data source. |
| [ibm_iam_action_control_template](ibm_iam_action_control_template.md) | resource | Manages ibm_iam_action_control_template. |
| [ibm_iam_action_control_template_version_data](data_ibm_iam_action_control_template_version.md) | data | Reads ibm_iam_action_control_template_version. |
| [ibm_iam_action_control_template_version](ibm_iam_action_control_template_version.md) | resource | Manages ibm_iam_action_control_template_version. |
| [ibm_iam_authorization_policies_data](data_ibm_iam_authorization_policies.md) | data | Schema validation for the ibm_iam_authorization_policies data source. |
| [ibm_iam_authorization_policy](ibm_iam_authorization_policy.md) | resource | Schema validation for the ibm_iam_authorization_policy resource. |
| [ibm_iam_authorization_policy_detach](ibm_iam_authorization_policy_detach.md) | resource | Schema validation for the ibm_iam_authorization_policy_detach resource. |
| [ibm_iam_custom_role](ibm_iam_custom_role.md) | resource | Manages ibm_iam_custom_role. |
| [ibm_iam_policy_assignment_data](data_ibm_iam_policy_assignment.md) | data | Schema validation for the ibm_iam_policy_assignment data source. |
| [ibm_iam_policy_assignment](ibm_iam_policy_assignment.md) | resource | Schema validation for the ibm_iam_policy_assignment resource. |
| [ibm_iam_policy_assignments_data](data_ibm_iam_policy_assignments.md) | data | Schema validation for the ibm_iam_policy_assignments data source. |
| [ibm_iam_policy_template_data](data_ibm_iam_policy_template.md) | data | Schema validation for the ibm_iam_policy_template data source. |
| [ibm_iam_policy_template](ibm_iam_policy_template.md) | resource | Manages ibm_iam_policy_template. |
| [ibm_iam_policy_template_version_data](data_ibm_iam_policy_template_version.md) | data | Reads ibm_iam_policy_template_version. |
| [ibm_iam_policy_template_version](ibm_iam_policy_template_version.md) | resource | Schema validation for the ibm_iam_policy_template_version resource. |
| [ibm_iam_role_actions_data](data_ibm_iam_role_actions.md) | data | Schema validation for the ibm_iam_role_actions data source. |
| [ibm_iam_role_assignment_data](data_ibm_iam_role_assignment.md) | data | Schema validation for the ibm_iam_role_assignment data source. |
| [ibm_iam_role_assignment](ibm_iam_role_assignment.md) | resource | Schema validation for the ibm_iam_role_assignment resource. |
| [ibm_iam_role_assignments_data](data_ibm_iam_role_assignments.md) | data | Schema validation for the ibm_iam_role_assignments data source. |
| [ibm_iam_role_template_data](data_ibm_iam_role_template.md) | data | Schema validation for the ibm_iam_role_template data source. |
| [ibm_iam_role_template](ibm_iam_role_template.md) | resource | Manages ibm_iam_role_template. |
| [ibm_iam_role_template_version_data](data_ibm_iam_role_template_version.md) | data | Reads ibm_iam_role_template_version. |
| [ibm_iam_role_template_version](ibm_iam_role_template_version.md) | resource | Manages ibm_iam_role_template_version. |
| [ibm_iam_roles_data](data_ibm_iam_roles.md) | data | Schema validation for the ibm_iam_roles data source. |
| [ibm_iam_service_policy_data](data_ibm_iam_service_policy.md) | data | Schema validation for the ibm_iam_service_policy data source. |
| [ibm_iam_service_policy](ibm_iam_service_policy.md) | resource | Schema validation for the ibm_iam_service_policy resource. |
| [ibm_iam_trusted_profile_policy_data](data_ibm_iam_trusted_profile_policy.md) | data | Schema validation for the ibm_iam_trusted_profile_policy data source. |
| [ibm_iam_trusted_profile_policy](ibm_iam_trusted_profile_policy.md) | resource | Schema validation for the ibm_iam_trusted_profile_policy resource. |
| [ibm_iam_user_invite](ibm_iam_user_invite.md) | resource | Schema validation for the ibm_iam_user_invite resource. |
| [ibm_iam_user_policy_data](data_ibm_iam_user_policy.md) | data | Schema validation for the ibm_iam_user_policy data source. |
| [ibm_iam_user_policy](ibm_iam_user_policy.md) | resource | Schema validation for the ibm_iam_user_policy resource. |

## kms

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_kms_cryptounits](ibm_kms_cryptounits.md) | resource | Schema validation for the ibm_kms_cryptounits resource. |
| [ibm_kms_instance_policies_data](data_ibm_kms_instance_policies.md) | data | Schema validation for the ibm_kms_instance_policies data source. |
| [ibm_kms_instance_policies](ibm_kms_instance_policies.md) | resource | Schema validation for the ibm_kms_instance_policies resource. |
| [ibm_kms_key_data](data_ibm_kms_key.md) | data | Schema validation for the ibm_kms_key data source. |
| [ibm_kms_key](ibm_kms_key.md) | resource | Schema validation for the ibm_kms_key resource. |
| [ibm_kms_key_alias](ibm_kms_key_alias.md) | resource | Schema validation for the ibm_kms_key_alias resource. |
| [ibm_kms_key_policies_data](data_ibm_kms_key_policies.md) | data | Schema validation for the ibm_kms_key_policies data source. |
| [ibm_kms_key_policies](ibm_kms_key_policies.md) | resource | Schema validation for the ibm_kms_key_policies resource. |
| [ibm_kms_key_rings_data](data_ibm_kms_key_rings.md) | data | Schema validation for the ibm_kms_key_rings data source. |
| [ibm_kms_key_rings](ibm_kms_key_rings.md) | resource | Schema validation for the ibm_kms_key_rings resource. |
| [ibm_kms_key_with_policy_overrides](ibm_kms_key_with_policy_overrides.md) | resource | Schema validation for the ibm_kms_key_with_policy_overrides resource. |
| [ibm_kms_keys_data](data_ibm_kms_keys.md) | data | Schema validation for the ibm_kms_keys data source. |
| [ibm_kms_kmip_adapter_data](data_ibm_kms_kmip_adapter.md) | data | Reads ibm_kms_kmip_adapter. |
| [ibm_kms_kmip_adapter](ibm_kms_kmip_adapter.md) | resource | Manages ibm_kms_kmip_adapter. |
| [ibm_kms_kmip_adapters_data](data_ibm_kms_kmip_adapters.md) | data | Schema validation for the ibm_kms_kmip_adapters data source. |
| [ibm_kms_kmip_client_cert_data](data_ibm_kms_kmip_client_cert.md) | data | Reads ibm_kms_kmip_client_cert. |
| [ibm_kms_kmip_client_cert](ibm_kms_kmip_client_cert.md) | resource | Manages ibm_kms_kmip_client_cert. |
| [ibm_kms_kmip_client_certs_data](data_ibm_kms_kmip_client_certs.md) | data | Schema validation for the ibm_kms_kmip_client_certs data source. |
| [ibm_kms_kmip_object_data](data_ibm_kms_kmip_object.md) | data | Schema validation for the ibm_kms_kmip_object data source. |
| [ibm_kms_kmip_objects_data](data_ibm_kms_kmip_objects.md) | data | Schema validation for the ibm_kms_kmip_objects data source. |
| [ibm_kp_key_data](data_ibm_kp_key.md) | data | Schema validation for the ibm_kp_key data source. |
| [ibm_kp_key](ibm_kp_key.md) | resource | Schema validation for the ibm_kp_key resource. |

## kubernetes

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_container_addons_data](data_ibm_container_addons.md) | data | Schema validation for the ibm_container_addons data source. |
| [ibm_container_addons](ibm_container_addons.md) | resource | Schema validation for the ibm_container_addons resource. |
| [ibm_container_alb_data](data_ibm_container_alb.md) | data | Reads ibm_container_alb. |
| [ibm_container_alb](ibm_container_alb.md) | resource | Manages ibm_container_alb. |
| [ibm_container_alb_cert_data](data_ibm_container_alb_cert.md) | data | Schema validation for the ibm_container_alb_cert data source. |
| [ibm_container_alb_cert](ibm_container_alb_cert.md) | resource | Schema validation for the ibm_container_alb_cert resource. |
| [ibm_container_alb_create](ibm_container_alb_create.md) | resource | Manages ibm_container_alb_create. |
| [ibm_container_api_key_reset](ibm_container_api_key_reset.md) | resource | Schema validation for the ibm_container_api_key_reset resource. |
| [ibm_container_bind_service_data](data_ibm_container_bind_service.md) | data | Schema validation for the ibm_container_bind_service data source. |
| [ibm_container_bind_service](ibm_container_bind_service.md) | resource | Schema validation for the ibm_container_bind_service resource. |
| [ibm_container_cluster_data](data_ibm_container_cluster.md) | data | Reads ibm_container_cluster. |
| [ibm_container_cluster](ibm_container_cluster.md) | resource | Manages ibm_container_cluster. |
| [ibm_container_cluster_config_data](data_ibm_container_cluster_config.md) | data | Schema validation for the ibm_container_cluster_config data source. |
| [ibm_container_cluster_feature](ibm_container_cluster_feature.md) | resource | Schema validation for the ibm_container_cluster_feature resource. |
| [ibm_container_cluster_versions_data](data_ibm_container_cluster_versions.md) | data | Schema validation for the ibm_container_cluster_versions data source. |
| [ibm_container_cluster_worker_data](data_ibm_container_cluster_worker.md) | data | Schema validation for the ibm_container_cluster_worker data source. |
| [ibm_container_dedicated_host_data](data_ibm_container_dedicated_host.md) | data | Schema validation for the ibm_container_dedicated_host data source. |
| [ibm_container_dedicated_host](ibm_container_dedicated_host.md) | resource | Schema validation for the ibm_container_dedicated_host resource. |
| [ibm_container_dedicated_host_flavor_data](data_ibm_container_dedicated_host_flavor.md) | data | Schema validation for the ibm_container_dedicated_host_flavor data source. |
| [ibm_container_dedicated_host_flavors_data](data_ibm_container_dedicated_host_flavors.md) | data | Schema validation for the ibm_container_dedicated_host_flavors data source. |
| [ibm_container_dedicated_host_pool_data](data_ibm_container_dedicated_host_pool.md) | data | Reads ibm_container_dedicated_host_pool. |
| [ibm_container_dedicated_host_pool](ibm_container_dedicated_host_pool.md) | resource | Manages ibm_container_dedicated_host_pool. |
| [ibm_container_ingress_instance_data](data_ibm_container_ingress_instance.md) | data | Schema validation for the ibm_container_ingress_instance data source. |
| [ibm_container_ingress_instance](ibm_container_ingress_instance.md) | resource | Schema validation for the ibm_container_ingress_instance resource. |
| [ibm_container_ingress_secret_opaque_data](data_ibm_container_ingress_secret_opaque.md) | data | Schema validation for the ibm_container_ingress_secret_opaque data source. |
| [ibm_container_ingress_secret_opaque](ibm_container_ingress_secret_opaque.md) | resource | Schema validation for the ibm_container_ingress_secret_opaque resource. |
| [ibm_container_ingress_secret_tls_data](data_ibm_container_ingress_secret_tls.md) | data | Schema validation for the ibm_container_ingress_secret_tls data source. |
| [ibm_container_ingress_secret_tls](ibm_container_ingress_secret_tls.md) | resource | Schema validation for the ibm_container_ingress_secret_tls resource. |
| [ibm_container_nlb_dns_data](data_ibm_container_nlb_dns.md) | data | Schema validation for the ibm_container_nlb_dns data source. |
| [ibm_container_nlb_dns](ibm_container_nlb_dns.md) | resource | Schema validation for the ibm_container_nlb_dns resource. |
| [ibm_container_storage_attachment_data](data_ibm_container_storage_attachment.md) | data | Schema validation for the ibm_container_storage_attachment data source. |
| [ibm_container_storage_attachment](ibm_container_storage_attachment.md) | resource | Schema validation for the ibm_container_storage_attachment resource. |
| [ibm_container_vni_attachment_data](data_ibm_container_vni_attachment.md) | data | Schema validation for the ibm_container_vni_attachment data source. |
| [ibm_container_vni_attachments_data](data_ibm_container_vni_attachments.md) | data | Schema validation for the ibm_container_vni_attachments data source. |
| [ibm_container_vni_baremetal_attachment](ibm_container_vni_baremetal_attachment.md) | resource | Schema validation for the ibm_container_vni_baremetal_attachment resource. |
| [ibm_container_vpc_alb_data](data_ibm_container_vpc_alb.md) | data | Schema validation for the ibm_container_vpc_alb data source. |
| [ibm_container_vpc_alb](ibm_container_vpc_alb.md) | resource | Manages ibm_container_vpc_alb. |
| [ibm_container_vpc_alb_create](ibm_container_vpc_alb_create.md) | resource | Manages ibm_container_vpc_alb_create. |
| [ibm_container_vpc_cluster_data](data_ibm_container_vpc_cluster.md) | data | Reads ibm_container_vpc_cluster. |
| [ibm_container_vpc_cluster](ibm_container_vpc_cluster.md) | resource | Manages ibm_container_vpc_cluster. |
| [ibm_container_vpc_cluster_alb_data](data_ibm_container_vpc_cluster_alb.md) | data | Schema validation for the ibm_container_vpc_cluster_alb data source. |
| [ibm_container_vpc_cluster_worker_data](data_ibm_container_vpc_cluster_worker.md) | data | Schema validation for the ibm_container_vpc_cluster_worker data source. |
| [ibm_container_vpc_cluster_worker_pool_data](data_ibm_container_vpc_cluster_worker_pool.md) | data | Schema validation for the ibm_container_vpc_cluster_worker_pool data source. |
| [ibm_container_vpc_worker](ibm_container_vpc_worker.md) | resource | Schema validation for the ibm_container_vpc_worker resource. |
| [ibm_container_vpc_worker_pool_data](data_ibm_container_vpc_worker_pool.md) | data | Schema validation for the ibm_container_vpc_worker_pool data source. |
| [ibm_container_vpc_worker_pool](ibm_container_vpc_worker_pool.md) | resource | Schema validation for the ibm_container_vpc_worker_pool resource. |
| [ibm_container_worker_pool_data](data_ibm_container_worker_pool.md) | data | Schema validation for the ibm_container_worker_pool data source. |
| [ibm_container_worker_pool](ibm_container_worker_pool.md) | resource | Schema validation for the ibm_container_worker_pool resource. |
| [ibm_container_worker_pool_zone_attachment](ibm_container_worker_pool_zone_attachment.md) | resource | Schema validation for the ibm_container_worker_pool_zone_attachment resource. |

## logs

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_logs_alert_data](data_ibm_logs_alert.md) | data | Reads ibm_logs_alert. |
| [ibm_logs_alert](ibm_logs_alert.md) | resource | Manages ibm_logs_alert. |
| [ibm_logs_alert_definition_data](data_ibm_logs_alert_definition.md) | data | Reads ibm_logs_alert_definition. |
| [ibm_logs_alert_definition](ibm_logs_alert_definition.md) | resource | Manages ibm_logs_alert_definition. |
| [ibm_logs_alert_definitions_data](data_ibm_logs_alert_definitions.md) | data | Schema validation for the ibm_logs_alert_definitions data source. |
| [ibm_logs_alerts_data](data_ibm_logs_alerts.md) | data | Schema validation for the ibm_logs_alerts data source. |
| [ibm_logs_dashboard_data](data_ibm_logs_dashboard.md) | data | Reads ibm_logs_dashboard. |
| [ibm_logs_dashboard](ibm_logs_dashboard.md) | resource | Manages ibm_logs_dashboard. |
| [ibm_logs_dashboard_folder](ibm_logs_dashboard_folder.md) | resource | Manages ibm_logs_dashboard_folder. |
| [ibm_logs_dashboard_folders_data](data_ibm_logs_dashboard_folders.md) | data | Schema validation for the ibm_logs_dashboard_folders data source. |
| [ibm_logs_data_access_rule](ibm_logs_data_access_rule.md) | resource | Schema validation for the ibm_logs_data_access_rule resource. |
| [ibm_logs_data_access_rules_data](data_ibm_logs_data_access_rules.md) | data | Schema validation for the ibm_logs_data_access_rules data source. |
| [ibm_logs_data_usage_metrics_data](data_ibm_logs_data_usage_metrics.md) | data | Schema validation for the ibm_logs_data_usage_metrics data source. |
| [ibm_logs_data_usage_metrics](ibm_logs_data_usage_metrics.md) | resource | Schema validation for the ibm_logs_data_usage_metrics resource. |
| [ibm_logs_e2m_data](data_ibm_logs_e2m.md) | data | Reads ibm_logs_e2m. |
| [ibm_logs_e2m](ibm_logs_e2m.md) | resource | Manages ibm_logs_e2m. |
| [ibm_logs_e2ms_data](data_ibm_logs_e2ms.md) | data | Schema validation for the ibm_logs_e2ms data source. |
| [ibm_logs_enrichment](ibm_logs_enrichment.md) | resource | Schema validation for the ibm_logs_enrichment resource. |
| [ibm_logs_enrichments_data](data_ibm_logs_enrichments.md) | data | Schema validation for the ibm_logs_enrichments data source. |
| [ibm_logs_extension_data](data_ibm_logs_extension.md) | data | Reads ibm_logs_extension. |
| [ibm_logs_extension_deployment_data](data_ibm_logs_extension_deployment.md) | data | Schema validation for the ibm_logs_extension_deployment data source. |
| [ibm_logs_extension_deployment](ibm_logs_extension_deployment.md) | resource | Schema validation for the ibm_logs_extension_deployment resource. |
| [ibm_logs_extensions_data](data_ibm_logs_extensions.md) | data | Schema validation for the ibm_logs_extensions data source. |
| [ibm_logs_log_data_retention_tags_data](data_ibm_logs_log_data_retention_tags.md) | data | Schema validation for the ibm_logs_log_data_retention_tags data source. |
| [ibm_logs_log_data_retention_tags](ibm_logs_log_data_retention_tags.md) | resource | Schema validation for the ibm_logs_log_data_retention_tags resource. |
| [ibm_logs_outgoing_webhook_data](data_ibm_logs_outgoing_webhook.md) | data | Reads ibm_logs_outgoing_webhook. |
| [ibm_logs_outgoing_webhook](ibm_logs_outgoing_webhook.md) | resource | Manages ibm_logs_outgoing_webhook. |
| [ibm_logs_outgoing_webhooks_data](data_ibm_logs_outgoing_webhooks.md) | data | Schema validation for the ibm_logs_outgoing_webhooks data source. |
| [ibm_logs_policies_data](data_ibm_logs_policies.md) | data | Schema validation for the ibm_logs_policies data source. |
| [ibm_logs_policy_data](data_ibm_logs_policy.md) | data | Reads ibm_logs_policy. |
| [ibm_logs_policy](ibm_logs_policy.md) | resource | Manages ibm_logs_policy. |
| [ibm_logs_rule_group_data](data_ibm_logs_rule_group.md) | data | Reads ibm_logs_rule_group. |
| [ibm_logs_rule_group](ibm_logs_rule_group.md) | resource | Manages ibm_logs_rule_group. |
| [ibm_logs_rule_groups_data](data_ibm_logs_rule_groups.md) | data | Schema validation for the ibm_logs_rule_groups data source. |
| [ibm_logs_stream_data](data_ibm_logs_stream.md) | data | Reads ibm_logs_stream. |
| [ibm_logs_stream](ibm_logs_stream.md) | resource | Manages ibm_logs_stream. |
| [ibm_logs_streams_data](data_ibm_logs_streams.md) | data | Schema validation for the ibm_logs_streams data source. |
| [ibm_logs_view_data](data_ibm_logs_view.md) | data | Reads ibm_logs_view. |
| [ibm_logs_view](ibm_logs_view.md) | resource | Manages ibm_logs_view. |
| [ibm_logs_view_folder_data](data_ibm_logs_view_folder.md) | data | Reads ibm_logs_view_folder. |
| [ibm_logs_view_folder](ibm_logs_view_folder.md) | resource | Manages ibm_logs_view_folder. |
| [ibm_logs_view_folders_data](data_ibm_logs_view_folders.md) | data | Schema validation for the ibm_logs_view_folders data source. |
| [ibm_logs_views_data](data_ibm_logs_views.md) | data | Schema validation for the ibm_logs_views data source. |

## logsrouter

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_logs_router_route](ibm_logs_router_route.md) | resource | Manages ibm_logs_router_route. |
| [ibm_logs_router_routes_data](data_ibm_logs_router_routes.md) | data | Reads ibm_logs_router_routes. |
| [ibm_logs_router_settings](ibm_logs_router_settings.md) | resource | Schema validation for the ibm_logs_router_settings resource. |
| [ibm_logs_router_target](ibm_logs_router_target.md) | resource | Manages ibm_logs_router_target. |
| [ibm_logs_router_targets_data](data_ibm_logs_router_targets.md) | data | Reads ibm_logs_router_targets. |

## logsrouting

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_logs_router_tenant](ibm_logs_router_tenant.md) | resource | Manages ibm_logs_router_tenant. |
| [ibm_logs_router_tenants_data](data_ibm_logs_router_tenants.md) | data | Reads ibm_logs_router_tenants. |

## metricsrouter

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_metrics_router_route](ibm_metrics_router_route.md) | resource | Manages ibm_metrics_router_route. |
| [ibm_metrics_router_routes_data](data_ibm_metrics_router_routes.md) | data | Reads ibm_metrics_router_routes. |
| [ibm_metrics_router_settings](ibm_metrics_router_settings.md) | resource | Schema validation for the ibm_metrics_router_settings resource. |
| [ibm_metrics_router_target](ibm_metrics_router_target.md) | resource | Manages ibm_metrics_router_target. |
| [ibm_metrics_router_targets_data](data_ibm_metrics_router_targets.md) | data | Reads ibm_metrics_router_targets. |

## mqcloud

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_mqcloud_application_data](data_ibm_mqcloud_application.md) | data | Reads ibm_mqcloud_application. |
| [ibm_mqcloud_application](ibm_mqcloud_application.md) | resource | Manages ibm_mqcloud_application. |
| [ibm_mqcloud_keystore_certificate_data](data_ibm_mqcloud_keystore_certificate.md) | data | Schema validation for the ibm_mqcloud_keystore_certificate data source. |
| [ibm_mqcloud_keystore_certificate](ibm_mqcloud_keystore_certificate.md) | resource | Schema validation for the ibm_mqcloud_keystore_certificate resource. |
| [ibm_mqcloud_queue_manager_data](data_ibm_mqcloud_queue_manager.md) | data | Reads ibm_mqcloud_queue_manager. |
| [ibm_mqcloud_queue_manager](ibm_mqcloud_queue_manager.md) | resource | Manages ibm_mqcloud_queue_manager. |
| [ibm_mqcloud_queue_manager_options_data](data_ibm_mqcloud_queue_manager_options.md) | data | Schema validation for the ibm_mqcloud_queue_manager_options data source. |
| [ibm_mqcloud_queue_manager_status_data](data_ibm_mqcloud_queue_manager_status.md) | data | Schema validation for the ibm_mqcloud_queue_manager_status data source. |
| [ibm_mqcloud_truststore_certificate_data](data_ibm_mqcloud_truststore_certificate.md) | data | Schema validation for the ibm_mqcloud_truststore_certificate data source. |
| [ibm_mqcloud_truststore_certificate](ibm_mqcloud_truststore_certificate.md) | resource | Schema validation for the ibm_mqcloud_truststore_certificate resource. |
| [ibm_mqcloud_user_data](data_ibm_mqcloud_user.md) | data | Reads ibm_mqcloud_user. |
| [ibm_mqcloud_user](ibm_mqcloud_user.md) | resource | Manages ibm_mqcloud_user. |
| [ibm_mqcloud_virtual_private_endpoint_gateway_data](data_ibm_mqcloud_virtual_private_endpoint_gateway.md) | data | Reads ibm_mqcloud_virtual_private_endpoint_gateway. |
| [ibm_mqcloud_virtual_private_endpoint_gateway](ibm_mqcloud_virtual_private_endpoint_gateway.md) | resource | Manages ibm_mqcloud_virtual_private_endpoint_gateway. |
| [ibm_mqcloud_virtual_private_endpoint_gateways_data](data_ibm_mqcloud_virtual_private_endpoint_gateways.md) | data | Reads ibm_mqcloud_virtual_private_endpoint_gateways. |

## pag

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_pag_instance_data](data_ibm_pag_instance.md) | data | Reads ibm_pag_instance. |
| [ibm_pag_instance](ibm_pag_instance.md) | resource | Manages ibm_pag_instance. |

## partnercentersell

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_onboarding_catalog_deployment](ibm_onboarding_catalog_deployment.md) | resource | Manages ibm_onboarding_catalog_deployment. |
| [ibm_onboarding_catalog_plan](ibm_onboarding_catalog_plan.md) | resource | Manages ibm_onboarding_catalog_plan. |
| [ibm_onboarding_catalog_product](ibm_onboarding_catalog_product.md) | resource | Manages ibm_onboarding_catalog_product. |
| [ibm_onboarding_iam_registration](ibm_onboarding_iam_registration.md) | resource | Manages ibm_onboarding_iam_registration. |
| [ibm_onboarding_product](ibm_onboarding_product.md) | resource | Schema validation for the ibm_onboarding_product resource. |
| [ibm_onboarding_registration](ibm_onboarding_registration.md) | resource | Schema validation for the ibm_onboarding_registration resource. |
| [ibm_onboarding_resource_broker](ibm_onboarding_resource_broker.md) | resource | Manages ibm_onboarding_resource_broker. |

## platformnotifications

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_notification_distribution_list_destination](ibm_notification_distribution_list_destination.md) | resource | Schema validation for the ibm_notification_distribution_list_destination resource. |

## power

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_pi_available_hosts_data](data_ibm_pi_available_hosts.md) | data | Schema validation for the ibm_pi_available_hosts data source. |
| [ibm_pi_capture](ibm_pi_capture.md) | resource | Schema validation for the ibm_pi_capture resource. |
| [ibm_pi_catalog_images_data](data_ibm_pi_catalog_images.md) | data | Schema validation for the ibm_pi_catalog_images data source. |
| [ibm_pi_cloud_connection_data](data_ibm_pi_cloud_connection.md) | data | Schema validation for the ibm_pi_cloud_connection data source. |
| [ibm_pi_cloud_connection](ibm_pi_cloud_connection.md) | resource | Schema validation for the ibm_pi_cloud_connection resource. |
| [ibm_pi_cloud_connection_network_attach](ibm_pi_cloud_connection_network_attach.md) | resource | Schema validation for the ibm_pi_cloud_connection_network_attach resource. |
| [ibm_pi_cloud_connections_data](data_ibm_pi_cloud_connections.md) | data | Schema validation for the ibm_pi_cloud_connections data source. |
| [ibm_pi_cloud_instance_data](data_ibm_pi_cloud_instance.md) | data | Schema validation for the ibm_pi_cloud_instance data source. |
| [ibm_pi_console_language](ibm_pi_console_language.md) | resource | Schema validation for the ibm_pi_console_language resource. |
| [ibm_pi_console_languages_data](data_ibm_pi_console_languages.md) | data | Schema validation for the ibm_pi_console_languages data source. |
| [ibm_pi_datacenter_data](data_ibm_pi_datacenter.md) | data | Schema validation for the ibm_pi_datacenter data source. |
| [ibm_pi_datacenters_data](data_ibm_pi_datacenters.md) | data | Schema validation for the ibm_pi_datacenters data source. |
| [ibm_pi_dhcp_data](data_ibm_pi_dhcp.md) | data | Schema validation for the ibm_pi_dhcp data source. |
| [ibm_pi_dhcp](ibm_pi_dhcp.md) | resource | Schema validation for the ibm_pi_dhcp resource. |
| [ibm_pi_dhcps_data](data_ibm_pi_dhcps.md) | data | Schema validation for the ibm_pi_dhcps data source. |
| [ibm_pi_disaster_recovery_location_data](data_ibm_pi_disaster_recovery_location.md) | data | Schema validation for the ibm_pi_disaster_recovery_location data source. |
| [ibm_pi_disaster_recovery_locations_data](data_ibm_pi_disaster_recovery_locations.md) | data | Schema validation for the ibm_pi_disaster_recovery_locations data source. |
| [ibm_pi_host_data](data_ibm_pi_host.md) | data | Schema validation for the ibm_pi_host data source. |
| [ibm_pi_host](ibm_pi_host.md) | resource | Schema validation for the ibm_pi_host resource. |
| [ibm_pi_host_group_data](data_ibm_pi_host_group.md) | data | Reads ibm_pi_host_group. |
| [ibm_pi_host_group](ibm_pi_host_group.md) | resource | Manages ibm_pi_host_group. |
| [ibm_pi_host_groups_data](data_ibm_pi_host_groups.md) | data | Schema validation for the ibm_pi_host_groups data source. |
| [ibm_pi_hosts_data](data_ibm_pi_hosts.md) | data | Schema validation for the ibm_pi_hosts data source. |
| [ibm_pi_ike_policy](ibm_pi_ike_policy.md) | resource | Schema validation for the ibm_pi_ike_policy resource. |
| [ibm_pi_image_data](data_ibm_pi_image.md) | data | Reads ibm_pi_image. |
| [ibm_pi_image](ibm_pi_image.md) | resource | Schema validation for the ibm_pi_image resource. |
| [ibm_pi_image_export](ibm_pi_image_export.md) | resource | Schema validation for the ibm_pi_image_export resource. |
| [ibm_pi_images_data](data_ibm_pi_images.md) | data | Schema validation for the ibm_pi_images data source. |
| [ibm_pi_instance_data](data_ibm_pi_instance.md) | data | Schema validation for the ibm_pi_instance data source. |
| [ibm_pi_instance](ibm_pi_instance.md) | resource | Schema validation for the ibm_pi_instance resource. |
| [ibm_pi_instance_action](ibm_pi_instance_action.md) | resource | Schema validation for the ibm_pi_instance_action resource. |
| [ibm_pi_instance_ip_data](data_ibm_pi_instance_ip.md) | data | Schema validation for the ibm_pi_instance_ip data source. |
| [ibm_pi_instance_network_data](data_ibm_pi_instance_network.md) | data | Schema validation for the ibm_pi_instance_network data source. |
| [ibm_pi_instance_networks_data](data_ibm_pi_instance_networks.md) | data | Schema validation for the ibm_pi_instance_networks data source. |
| [ibm_pi_instance_snapshot_data](data_ibm_pi_instance_snapshot.md) | data | Reads ibm_pi_instance_snapshot. |
| [ibm_pi_instance_snapshot](ibm_pi_instance_snapshot.md) | resource | Schema validation for the ibm_pi_instance_snapshot resource. |
| [ibm_pi_instance_snapshots_data](data_ibm_pi_instance_snapshots.md) | data | Schema validation for the ibm_pi_instance_snapshots data source. |
| [ibm_pi_instance_volumes_data](data_ibm_pi_instance_volumes.md) | data | Schema validation for the ibm_pi_instance_volumes data source. |
| [ibm_pi_instance_vpmem_volume_data](data_ibm_pi_instance_vpmem_volume.md) | data | Reads ibm_pi_instance_vpmem_volume. |
| [ibm_pi_instance_vpmem_volumes_data](data_ibm_pi_instance_vpmem_volumes.md) | data | Schema validation for the ibm_pi_instance_vpmem_volumes data source. |
| [ibm_pi_instance_vpmem_volumes](ibm_pi_instance_vpmem_volumes.md) | resource | Schema validation for the ibm_pi_instance_vpmem_volumes resource. |
| [ibm_pi_instances_data](data_ibm_pi_instances.md) | data | Schema validation for the ibm_pi_instances data source. |
| [ibm_pi_ipsec_policy](ibm_pi_ipsec_policy.md) | resource | Schema validation for the ibm_pi_ipsec_policy resource. |
| [ibm_pi_key_data](data_ibm_pi_key.md) | data | Reads ibm_pi_key. |
| [ibm_pi_key](ibm_pi_key.md) | resource | Manages ibm_pi_key. |
| [ibm_pi_keys_data](data_ibm_pi_keys.md) | data | Schema validation for the ibm_pi_keys data source. |
| [ibm_pi_network_data](data_ibm_pi_network.md) | data | Reads ibm_pi_network. |
| [ibm_pi_network](ibm_pi_network.md) | resource | Schema validation for the ibm_pi_network resource. |
| [ibm_pi_network_address_group_data](data_ibm_pi_network_address_group.md) | data | Reads ibm_pi_network_address_group. |
| [ibm_pi_network_address_group](ibm_pi_network_address_group.md) | resource | Schema validation for the ibm_pi_network_address_group resource. |
| [ibm_pi_network_address_group_member](ibm_pi_network_address_group_member.md) | resource | Manages ibm_pi_network_address_group_member. |
| [ibm_pi_network_address_groups_data](data_ibm_pi_network_address_groups.md) | data | Schema validation for the ibm_pi_network_address_groups data source. |
| [ibm_pi_network_interface_data](data_ibm_pi_network_interface.md) | data | Reads ibm_pi_network_interface. |
| [ibm_pi_network_interface](ibm_pi_network_interface.md) | resource | Manages ibm_pi_network_interface. |
| [ibm_pi_network_interfaces_data](data_ibm_pi_network_interfaces.md) | data | Schema validation for the ibm_pi_network_interfaces data source. |
| [ibm_pi_network_peer_data](data_ibm_pi_network_peer.md) | data | Reads ibm_pi_network_peer. |
| [ibm_pi_network_peer](ibm_pi_network_peer.md) | resource | Schema validation for the ibm_pi_network_peer resource. |
| [ibm_pi_network_peer_interfaces_data](data_ibm_pi_network_peer_interfaces.md) | data | Schema validation for the ibm_pi_network_peer_interfaces data source. |
| [ibm_pi_network_peer_route_filter_data](data_ibm_pi_network_peer_route_filter.md) | data | Schema validation for the ibm_pi_network_peer_route_filter data source. |
| [ibm_pi_network_peer_route_filter](ibm_pi_network_peer_route_filter.md) | resource | Schema validation for the ibm_pi_network_peer_route_filter resource. |
| [ibm_pi_network_peers_data](data_ibm_pi_network_peers.md) | data | Schema validation for the ibm_pi_network_peers data source. |
| [ibm_pi_network_port_data](data_ibm_pi_network_port.md) | data | Schema validation for the ibm_pi_network_port data source. |
| [ibm_pi_network_port_attach](ibm_pi_network_port_attach.md) | resource | Schema validation for the ibm_pi_network_port_attach resource. |
| [ibm_pi_network_security_group_data](data_ibm_pi_network_security_group.md) | data | Reads ibm_pi_network_security_group. |
| [ibm_pi_network_security_group](ibm_pi_network_security_group.md) | resource | Schema validation for the ibm_pi_network_security_group resource. |
| [ibm_pi_network_security_group_action](ibm_pi_network_security_group_action.md) | resource | Schema validation for the ibm_pi_network_security_group_action resource. |
| [ibm_pi_network_security_group_member](ibm_pi_network_security_group_member.md) | resource | Manages ibm_pi_network_security_group_member. |
| [ibm_pi_network_security_group_rule](ibm_pi_network_security_group_rule.md) | resource | Manages ibm_pi_network_security_group_rule. |
| [ibm_pi_network_security_groups_data](data_ibm_pi_network_security_groups.md) | data | Schema validation for the ibm_pi_network_security_groups data source. |
| [ibm_pi_networks_data](data_ibm_pi_networks.md) | data | Schema validation for the ibm_pi_networks data source. |
| [ibm_pi_placement_group_data](data_ibm_pi_placement_group.md) | data | Reads ibm_pi_placement_group. |
| [ibm_pi_placement_group](ibm_pi_placement_group.md) | resource | Schema validation for the ibm_pi_placement_group resource. |
| [ibm_pi_placement_groups_data](data_ibm_pi_placement_groups.md) | data | Schema validation for the ibm_pi_placement_groups data source. |
| [ibm_pi_public_network_data](data_ibm_pi_public_network.md) | data | Reads ibm_pi_public_network. |
| [ibm_pi_pvm_snapshots_data](data_ibm_pi_pvm_snapshots.md) | data | Schema validation for the ibm_pi_pvm_snapshots data source. |
| [ibm_pi_route_data](data_ibm_pi_route.md) | data | Reads ibm_pi_route. |
| [ibm_pi_route](ibm_pi_route.md) | resource | Schema validation for the ibm_pi_route resource. |
| [ibm_pi_route_report_data](data_ibm_pi_route_report.md) | data | Schema validation for the ibm_pi_route_report data source. |
| [ibm_pi_routes_data](data_ibm_pi_routes.md) | data | Schema validation for the ibm_pi_routes data source. |
| [ibm_pi_sap_profile_data](data_ibm_pi_sap_profile.md) | data | Schema validation for the ibm_pi_sap_profile data source. |
| [ibm_pi_sap_profiles_data](data_ibm_pi_sap_profiles.md) | data | Schema validation for the ibm_pi_sap_profiles data source. |
| [ibm_pi_shared_processor_pool_data](data_ibm_pi_shared_processor_pool.md) | data | Reads ibm_pi_shared_processor_pool. |
| [ibm_pi_shared_processor_pool](ibm_pi_shared_processor_pool.md) | resource | Schema validation for the ibm_pi_shared_processor_pool resource. |
| [ibm_pi_shared_processor_pools_data](data_ibm_pi_shared_processor_pools.md) | data | Schema validation for the ibm_pi_shared_processor_pools data source. |
| [ibm_pi_snapshot](ibm_pi_snapshot.md) | resource | Schema validation for the ibm_pi_snapshot resource. |
| [ibm_pi_software_tiers_data](data_ibm_pi_software_tiers.md) | data | Schema validation for the ibm_pi_software_tiers data source. |
| [ibm_pi_spp_placement_group_data](data_ibm_pi_spp_placement_group.md) | data | Reads ibm_pi_spp_placement_group. |
| [ibm_pi_spp_placement_group](ibm_pi_spp_placement_group.md) | resource | Schema validation for the ibm_pi_spp_placement_group resource. |
| [ibm_pi_spp_placement_groups_data](data_ibm_pi_spp_placement_groups.md) | data | Schema validation for the ibm_pi_spp_placement_groups data source. |
| [ibm_pi_storage_pool_capacity_data](data_ibm_pi_storage_pool_capacity.md) | data | Schema validation for the ibm_pi_storage_pool_capacity data source. |
| [ibm_pi_storage_pools_capacity_data](data_ibm_pi_storage_pools_capacity.md) | data | Schema validation for the ibm_pi_storage_pools_capacity data source. |
| [ibm_pi_storage_tiers_data](data_ibm_pi_storage_tiers.md) | data | Schema validation for the ibm_pi_storage_tiers data source. |
| [ibm_pi_storage_type_capacity_data](data_ibm_pi_storage_type_capacity.md) | data | Schema validation for the ibm_pi_storage_type_capacity data source. |
| [ibm_pi_storage_types_capacity_data](data_ibm_pi_storage_types_capacity.md) | data | Schema validation for the ibm_pi_storage_types_capacity data source. |
| [ibm_pi_system_pools_data](data_ibm_pi_system_pools.md) | data | Schema validation for the ibm_pi_system_pools data source. |
| [ibm_pi_tenant_data](data_ibm_pi_tenant.md) | data | Schema validation for the ibm_pi_tenant data source. |
| [ibm_pi_virtual_serial_number_data](data_ibm_pi_virtual_serial_number.md) | data | Schema validation for the ibm_pi_virtual_serial_number data source. |
| [ibm_pi_virtual_serial_number](ibm_pi_virtual_serial_number.md) | resource | Schema validation for the ibm_pi_virtual_serial_number resource. |
| [ibm_pi_virtual_serial_numbers_data](data_ibm_pi_virtual_serial_numbers.md) | data | Schema validation for the ibm_pi_virtual_serial_numbers data source. |
| [ibm_pi_volume_data](data_ibm_pi_volume.md) | data | Reads ibm_pi_volume. |
| [ibm_pi_volume](ibm_pi_volume.md) | resource | Schema validation for the ibm_pi_volume resource. |
| [ibm_pi_volume_attach](ibm_pi_volume_attach.md) | resource | Schema validation for the ibm_pi_volume_attach resource. |
| [ibm_pi_volume_clone_data](data_ibm_pi_volume_clone.md) | data | Schema validation for the ibm_pi_volume_clone data source. |
| [ibm_pi_volume_clone](ibm_pi_volume_clone.md) | resource | Schema validation for the ibm_pi_volume_clone resource. |
| [ibm_pi_volume_flash_copy_mappings_data](data_ibm_pi_volume_flash_copy_mappings.md) | data | Schema validation for the ibm_pi_volume_flash_copy_mappings data source. |
| [ibm_pi_volume_group_data](data_ibm_pi_volume_group.md) | data | Schema validation for the ibm_pi_volume_group data source. |
| [ibm_pi_volume_group](ibm_pi_volume_group.md) | resource | Schema validation for the ibm_pi_volume_group resource. |
| [ibm_pi_volume_group_action](ibm_pi_volume_group_action.md) | resource | Schema validation for the ibm_pi_volume_group_action resource. |
| [ibm_pi_volume_group_details_data](data_ibm_pi_volume_group_details.md) | data | Schema validation for the ibm_pi_volume_group_details data source. |
| [ibm_pi_volume_group_remote_copy_relationships_data](data_ibm_pi_volume_group_remote_copy_relationships.md) | data | Schema validation for the ibm_pi_volume_group_remote_copy_relationships data source. |
| [ibm_pi_volume_group_storage_details_data](data_ibm_pi_volume_group_storage_details.md) | data | Schema validation for the ibm_pi_volume_group_storage_details data source. |
| [ibm_pi_volume_groups_data](data_ibm_pi_volume_groups.md) | data | Schema validation for the ibm_pi_volume_groups data source. |
| [ibm_pi_volume_groups_details_data](data_ibm_pi_volume_groups_details.md) | data | Schema validation for the ibm_pi_volume_groups_details data source. |
| [ibm_pi_volume_onboarding_data](data_ibm_pi_volume_onboarding.md) | data | Schema validation for the ibm_pi_volume_onboarding data source. |
| [ibm_pi_volume_onboarding](ibm_pi_volume_onboarding.md) | resource | Schema validation for the ibm_pi_volume_onboarding resource. |
| [ibm_pi_volume_onboardings_data](data_ibm_pi_volume_onboardings.md) | data | Schema validation for the ibm_pi_volume_onboardings data source. |
| [ibm_pi_volume_remote_copy_relationship_data](data_ibm_pi_volume_remote_copy_relationship.md) | data | Reads ibm_pi_volume_remote_copy_relationship. |
| [ibm_pi_volume_snapshot_data](data_ibm_pi_volume_snapshot.md) | data | Reads ibm_pi_volume_snapshot. |
| [ibm_pi_volume_snapshots_data](data_ibm_pi_volume_snapshots.md) | data | Schema validation for the ibm_pi_volume_snapshots data source. |
| [ibm_pi_volumes_data](data_ibm_pi_volumes.md) | data | Schema validation for the ibm_pi_volumes data source. |
| [ibm_pi_vpn_connection](ibm_pi_vpn_connection.md) | resource | Schema validation for the ibm_pi_vpn_connection resource. |
| [ibm_pi_workspace_data](data_ibm_pi_workspace.md) | data | Schema validation for the ibm_pi_workspace data source. |
| [ibm_pi_workspace](ibm_pi_workspace.md) | resource | Schema validation for the ibm_pi_workspace resource. |
| [ibm_pi_workspaces_data](data_ibm_pi_workspaces.md) | data | Schema validation for the ibm_pi_workspaces data source. |

## powerhaautomationservice

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_pha_api_key](ibm_pha_api_key.md) | resource | Schema validation for the ibm_pha_api_key resource. |
| [ibm_pha_cluster_nodes_data](data_ibm_pha_cluster_nodes.md) | data | Schema validation for the ibm_pha_cluster_nodes data source. |
| [ibm_pha_cluster_nodes](ibm_pha_cluster_nodes.md) | resource | Schema validation for the ibm_pha_cluster_nodes resource. |
| [ibm_pha_deployment_data](data_ibm_pha_deployment.md) | data | Schema validation for the ibm_pha_deployment data source. |
| [ibm_pha_deployment](ibm_pha_deployment.md) | resource | Schema validation for the ibm_pha_deployment resource. |
| [ibm_pha_last_operation_data](data_ibm_pha_last_operation.md) | data | Schema validation for the ibm_pha_last_operation data source. |
| [ibm_pha_powervs_workspaces_data](data_ibm_pha_powervs_workspaces.md) | data | Schema validation for the ibm_pha_powervs_workspaces data source. |
| [ibm_pha_supported_locations_data](data_ibm_pha_supported_locations.md) | data | Schema validation for the ibm_pha_supported_locations data source. |

## project

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_project_data](data_ibm_project.md) | data | Schema validation for the ibm_project data source. |
| [ibm_project](ibm_project.md) | resource | Schema validation for the ibm_project resource. |
| [ibm_project_config_data](data_ibm_project_config.md) | data | Schema validation for the ibm_project_config data source. |
| [ibm_project_config](ibm_project_config.md) | resource | Schema validation for the ibm_project_config resource. |
| [ibm_project_environment_data](data_ibm_project_environment.md) | data | Schema validation for the ibm_project_environment data source. |
| [ibm_project_environment](ibm_project_environment.md) | resource | Schema validation for the ibm_project_environment resource. |

## pushnotification

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_pn_application_chrome_data](data_ibm_pn_application_chrome.md) | data | Schema validation for the ibm_pn_application_chrome data source. |
| [ibm_pn_application_chrome](ibm_pn_application_chrome.md) | resource | Schema validation for the ibm_pn_application_chrome resource. |

## registry

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_cr_namespace](ibm_cr_namespace.md) | resource | Manages ibm_cr_namespace. |
| [ibm_cr_namespaces_data](data_ibm_cr_namespaces.md) | data | Schema validation for the ibm_cr_namespaces data source. |
| [ibm_cr_retention_policy](ibm_cr_retention_policy.md) | resource | Schema validation for the ibm_cr_retention_policy resource. |

## resourcecontroller

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_resource_instance_data](data_ibm_resource_instance.md) | data | Reads ibm_resource_instance. |
| [ibm_resource_instance](ibm_resource_instance.md) | resource | Manages ibm_resource_instance. |
| [ibm_resource_key_data](data_ibm_resource_key.md) | data | Reads ibm_resource_key. |
| [ibm_resource_key](ibm_resource_key.md) | resource | Manages ibm_resource_key. |
| [ibm_resource_quota_data](data_ibm_resource_quota.md) | data | Reads ibm_resource_quota. |
| [ibm_resource_reclamation_delete](ibm_resource_reclamation_delete.md) | resource | Schema validation for the ibm_resource_reclamation_delete resource. |
| [ibm_resource_reclamations_data](data_ibm_resource_reclamations.md) | data | Schema validation for the ibm_resource_reclamations data source. |

## resourcemanager

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_resource_group_data](data_ibm_resource_group.md) | data | Reads ibm_resource_group. |
| [ibm_resource_group](ibm_resource_group.md) | resource | Manages ibm_resource_group. |
| [ibm_resource_groups_data](data_ibm_resource_groups.md) | data | Reads ibm_resource_groups. |

## satellite

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_satellite_attach_host_script_data](data_ibm_satellite_attach_host_script.md) | data | Schema validation for the ibm_satellite_attach_host_script data source. |
| [ibm_satellite_cluster_data](data_ibm_satellite_cluster.md) | data | Reads ibm_satellite_cluster. |
| [ibm_satellite_cluster](ibm_satellite_cluster.md) | resource | Manages ibm_satellite_cluster. |
| [ibm_satellite_cluster_worker_pool_data](data_ibm_satellite_cluster_worker_pool.md) | data | Reads ibm_satellite_cluster_worker_pool. |
| [ibm_satellite_cluster_worker_pool](ibm_satellite_cluster_worker_pool.md) | resource | Manages ibm_satellite_cluster_worker_pool. |
| [ibm_satellite_cluster_worker_pool_zone_attachment_data](data_ibm_satellite_cluster_worker_pool_zone_attachment.md) | data | Schema validation for the ibm_satellite_cluster_worker_pool_zone_attachment data source. |
| [ibm_satellite_cluster_worker_pool_zone_attachment](ibm_satellite_cluster_worker_pool_zone_attachment.md) | resource | Schema validation for the ibm_satellite_cluster_worker_pool_zone_attachment resource. |
| [ibm_satellite_endpoint_data](data_ibm_satellite_endpoint.md) | data | Schema validation for the ibm_satellite_endpoint data source. |
| [ibm_satellite_endpoint](ibm_satellite_endpoint.md) | resource | Schema validation for the ibm_satellite_endpoint resource. |
| [ibm_satellite_host](ibm_satellite_host.md) | resource | Schema validation for the ibm_satellite_host resource. |
| [ibm_satellite_link_data](data_ibm_satellite_link.md) | data | Schema validation for the ibm_satellite_link data source. |
| [ibm_satellite_link](ibm_satellite_link.md) | resource | Schema validation for the ibm_satellite_link resource. |
| [ibm_satellite_location_data](data_ibm_satellite_location.md) | data | Schema validation for the ibm_satellite_location data source. |
| [ibm_satellite_location](ibm_satellite_location.md) | resource | Schema validation for the ibm_satellite_location resource. |
| [ibm_satellite_location_nlb_dns_data](data_ibm_satellite_location_nlb_dns.md) | data | Schema validation for the ibm_satellite_location_nlb_dns data source. |
| [ibm_satellite_location_nlb_dns](ibm_satellite_location_nlb_dns.md) | resource | Schema validation for the ibm_satellite_location_nlb_dns resource. |
| [ibm_satellite_storage_assignment_data](data_ibm_satellite_storage_assignment.md) | data | Schema validation for the ibm_satellite_storage_assignment data source. |
| [ibm_satellite_storage_assignment](ibm_satellite_storage_assignment.md) | resource | Schema validation for the ibm_satellite_storage_assignment resource. |
| [ibm_satellite_storage_configuration_data](data_ibm_satellite_storage_configuration.md) | data | Schema validation for the ibm_satellite_storage_configuration data source. |
| [ibm_satellite_storage_configuration](ibm_satellite_storage_configuration.md) | resource | Schema validation for the ibm_satellite_storage_configuration resource. |

## scc

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_scc_account_location_data](data_ibm_scc_account_location.md) | data | ibm_scc_account_location data source is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_account_location_settings_data](data_ibm_scc_account_location_settings.md) | data | ibm_scc_account_location_settings data source is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_account_locations_data](data_ibm_scc_account_locations.md) | data | ibm_scc_account_locations data source is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_account_notification_settings_data](data_ibm_scc_account_notification_settings.md) | data | ibm_scc_account_notification_settings data source is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_account_settings](ibm_scc_account_settings.md) | resource | ibm_scc_account_settings resource is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_control_libraries_data](data_ibm_scc_control_libraries.md) | data | Schema validation for the ibm_scc_control_libraries data source. |
| [ibm_scc_control_library_data](data_ibm_scc_control_library.md) | data | Schema validation for the ibm_scc_control_library data source. |
| [ibm_scc_control_library](ibm_scc_control_library.md) | resource | Schema validation for the ibm_scc_control_library resource. |
| [ibm_scc_instance_settings_data](data_ibm_scc_instance_settings.md) | data | Schema validation for the ibm_scc_instance_settings data source. |
| [ibm_scc_instance_settings](ibm_scc_instance_settings.md) | resource | Schema validation for the ibm_scc_instance_settings resource. |
| [ibm_scc_latest_reports_data](data_ibm_scc_latest_reports.md) | data | Schema validation for the ibm_scc_latest_reports data source. |
| [ibm_scc_profile_data](data_ibm_scc_profile.md) | data | Schema validation for the ibm_scc_profile data source. |
| [ibm_scc_profile](ibm_scc_profile.md) | resource | Schema validation for the ibm_scc_profile resource. |
| [ibm_scc_profile_attachment_data](data_ibm_scc_profile_attachment.md) | data | Reads ibm_scc_profile_attachment. |
| [ibm_scc_profile_attachment](ibm_scc_profile_attachment.md) | resource | Manages ibm_scc_profile_attachment. |
| [ibm_scc_profiles_data](data_ibm_scc_profiles.md) | data | Schema validation for the ibm_scc_profiles data source. |
| [ibm_scc_provider_type_data](data_ibm_scc_provider_type.md) | data | Reads ibm_scc_provider_type. |
| [ibm_scc_provider_type_collection_data](data_ibm_scc_provider_type_collection.md) | data | Schema validation for the ibm_scc_provider_type_collection data source. |
| [ibm_scc_provider_type_instance_data](data_ibm_scc_provider_type_instance.md) | data | Reads ibm_scc_provider_type_instance. |
| [ibm_scc_provider_type_instance](ibm_scc_provider_type_instance.md) | resource | Manages ibm_scc_provider_type_instance. |
| [ibm_scc_provider_types_data](data_ibm_scc_provider_types.md) | data | Schema validation for the ibm_scc_provider_types data source. |
| [ibm_scc_report_data](data_ibm_scc_report.md) | data | Schema validation for the ibm_scc_report data source. |
| [ibm_scc_report_controls_data](data_ibm_scc_report_controls.md) | data | Schema validation for the ibm_scc_report_controls data source. |
| [ibm_scc_report_evaluations_data](data_ibm_scc_report_evaluations.md) | data | Schema validation for the ibm_scc_report_evaluations data source. |
| [ibm_scc_report_resources_data](data_ibm_scc_report_resources.md) | data | Schema validation for the ibm_scc_report_resources data source. |
| [ibm_scc_report_rule_data](data_ibm_scc_report_rule.md) | data | Schema validation for the ibm_scc_report_rule data source. |
| [ibm_scc_report_summary_data](data_ibm_scc_report_summary.md) | data | Schema validation for the ibm_scc_report_summary data source. |
| [ibm_scc_report_tags_data](data_ibm_scc_report_tags.md) | data | Schema validation for the ibm_scc_report_tags data source. |
| [ibm_scc_report_violation_drift_data](data_ibm_scc_report_violation_drift.md) | data | Schema validation for the ibm_scc_report_violation_drift data source. |
| [ibm_scc_rule_data](data_ibm_scc_rule.md) | data | Schema validation for the ibm_scc_rule data source. |
| [ibm_scc_rule](ibm_scc_rule.md) | resource | Manages ibm_scc_rule. |
| [ibm_scc_rule_attachment](ibm_scc_rule_attachment.md) | resource | ibm_scc_rule_attachment resource is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_scope_data](data_ibm_scc_scope.md) | data | Reads ibm_scc_scope. |
| [ibm_scc_scope](ibm_scc_scope.md) | resource | Manages ibm_scc_scope. |
| [ibm_scc_scope_collection_data](data_ibm_scc_scope_collection.md) | data | Schema validation for the ibm_scc_scope_collection data source. |
| [ibm_scc_template](ibm_scc_template.md) | resource | ibm_scc_template resource is registered with an empty schema (deprecated or schema-less). |
| [ibm_scc_template_attachment](ibm_scc_template_attachment.md) | resource | ibm_scc_template_attachment resource is registered with an empty schema (deprecated or schema-less). |

## schematics

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_schematics_action_data](data_ibm_schematics_action.md) | data | Reads ibm_schematics_action. |
| [ibm_schematics_action](ibm_schematics_action.md) | resource | Manages ibm_schematics_action. |
| [ibm_schematics_agent_data](data_ibm_schematics_agent.md) | data | Reads ibm_schematics_agent. |
| [ibm_schematics_agent](ibm_schematics_agent.md) | resource | Manages ibm_schematics_agent. |
| [ibm_schematics_agent_deploy_data](data_ibm_schematics_agent_deploy.md) | data | Schema validation for the ibm_schematics_agent_deploy data source. |
| [ibm_schematics_agent_deploy](ibm_schematics_agent_deploy.md) | resource | Schema validation for the ibm_schematics_agent_deploy resource. |
| [ibm_schematics_agent_health_data](data_ibm_schematics_agent_health.md) | data | Schema validation for the ibm_schematics_agent_health data source. |
| [ibm_schematics_agent_health](ibm_schematics_agent_health.md) | resource | Schema validation for the ibm_schematics_agent_health resource. |
| [ibm_schematics_agent_prs_data](data_ibm_schematics_agent_prs.md) | data | Schema validation for the ibm_schematics_agent_prs data source. |
| [ibm_schematics_agent_prs](ibm_schematics_agent_prs.md) | resource | Schema validation for the ibm_schematics_agent_prs resource. |
| [ibm_schematics_agents_data](data_ibm_schematics_agents.md) | data | Reads ibm_schematics_agents. |
| [ibm_schematics_inventory_data](data_ibm_schematics_inventory.md) | data | Reads ibm_schematics_inventory. |
| [ibm_schematics_inventory](ibm_schematics_inventory.md) | resource | Manages ibm_schematics_inventory. |
| [ibm_schematics_job_data](data_ibm_schematics_job.md) | data | Reads ibm_schematics_job. |
| [ibm_schematics_job](ibm_schematics_job.md) | resource | Manages ibm_schematics_job. |
| [ibm_schematics_kms_data](data_ibm_schematics_kms.md) | data | Schema validation for the ibm_schematics_kms data source. |
| [ibm_schematics_output_data](data_ibm_schematics_output.md) | data | Schema validation for the ibm_schematics_output data source. |
| [ibm_schematics_policies_data](data_ibm_schematics_policies.md) | data | Schema validation for the ibm_schematics_policies data source. |
| [ibm_schematics_policy_data](data_ibm_schematics_policy.md) | data | Reads ibm_schematics_policy. |
| [ibm_schematics_policy](ibm_schematics_policy.md) | resource | Manages ibm_schematics_policy. |
| [ibm_schematics_resource_query_data](data_ibm_schematics_resource_query.md) | data | Reads ibm_schematics_resource_query. |
| [ibm_schematics_resource_query](ibm_schematics_resource_query.md) | resource | Manages ibm_schematics_resource_query. |
| [ibm_schematics_state_data](data_ibm_schematics_state.md) | data | Schema validation for the ibm_schematics_state data source. |
| [ibm_schematics_workspace_data](data_ibm_schematics_workspace.md) | data | Reads ibm_schematics_workspace. |
| [ibm_schematics_workspace](ibm_schematics_workspace.md) | resource | Manages ibm_schematics_workspace. |

## sdsaas

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_sds_host](ibm_sds_host.md) | resource | Manages ibm_sds_host. |
| [ibm_sds_volume](ibm_sds_volume.md) | resource | Manages ibm_sds_volume. |
| [ibm_sds_volume_mapping](ibm_sds_volume_mapping.md) | resource | Schema validation for the ibm_sds_volume_mapping resource. |

## secretsmanager

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_sm_arbitrary_secret_data](data_ibm_sm_arbitrary_secret.md) | data | Reads ibm_sm_arbitrary_secret. |
| [ibm_sm_arbitrary_secret](ibm_sm_arbitrary_secret.md) | resource | Manages ibm_sm_arbitrary_secret. |
| [ibm_sm_arbitrary_secret_metadata_data](data_ibm_sm_arbitrary_secret_metadata.md) | data | Reads ibm_sm_arbitrary_secret_metadata. |
| [ibm_sm_configurations_data](data_ibm_sm_configurations.md) | data | Schema validation for the ibm_sm_configurations data source. |
| [ibm_sm_custom_credentials_configuration_data](data_ibm_sm_custom_credentials_configuration.md) | data | Reads ibm_sm_custom_credentials_configuration. |
| [ibm_sm_custom_credentials_configuration](ibm_sm_custom_credentials_configuration.md) | resource | Manages ibm_sm_custom_credentials_configuration. |
| [ibm_sm_custom_credentials_secret_data](data_ibm_sm_custom_credentials_secret.md) | data | Reads ibm_sm_custom_credentials_secret. |
| [ibm_sm_custom_credentials_secret](ibm_sm_custom_credentials_secret.md) | resource | Manages ibm_sm_custom_credentials_secret. |
| [ibm_sm_custom_credentials_secret_metadata_data](data_ibm_sm_custom_credentials_secret_metadata.md) | data | Reads ibm_sm_custom_credentials_secret_metadata. |
| [ibm_sm_en_registration_data](data_ibm_sm_en_registration.md) | data | Schema validation for the ibm_sm_en_registration data source. |
| [ibm_sm_en_registration](ibm_sm_en_registration.md) | resource | Schema validation for the ibm_sm_en_registration resource. |
| [ibm_sm_iam_credentials_configuration_data](data_ibm_sm_iam_credentials_configuration.md) | data | Reads ibm_sm_iam_credentials_configuration. |
| [ibm_sm_iam_credentials_configuration](ibm_sm_iam_credentials_configuration.md) | resource | Manages ibm_sm_iam_credentials_configuration. |
| [ibm_sm_iam_credentials_secret_data](data_ibm_sm_iam_credentials_secret.md) | data | Reads ibm_sm_iam_credentials_secret. |
| [ibm_sm_iam_credentials_secret](ibm_sm_iam_credentials_secret.md) | resource | Manages ibm_sm_iam_credentials_secret. |
| [ibm_sm_iam_credentials_secret_metadata_data](data_ibm_sm_iam_credentials_secret_metadata.md) | data | Reads ibm_sm_iam_credentials_secret_metadata. |
| [ibm_sm_imported_certificate_data](data_ibm_sm_imported_certificate.md) | data | Reads ibm_sm_imported_certificate. |
| [ibm_sm_imported_certificate](ibm_sm_imported_certificate.md) | resource | Manages ibm_sm_imported_certificate. |
| [ibm_sm_imported_certificate_metadata_data](data_ibm_sm_imported_certificate_metadata.md) | data | Reads ibm_sm_imported_certificate_metadata. |
| [ibm_sm_kv_secret_data](data_ibm_sm_kv_secret.md) | data | Reads ibm_sm_kv_secret. |
| [ibm_sm_kv_secret](ibm_sm_kv_secret.md) | resource | Manages ibm_sm_kv_secret. |
| [ibm_sm_kv_secret_metadata_data](data_ibm_sm_kv_secret_metadata.md) | data | Reads ibm_sm_kv_secret_metadata. |
| [ibm_sm_private_certificate_data](data_ibm_sm_private_certificate.md) | data | Reads ibm_sm_private_certificate. |
| [ibm_sm_private_certificate](ibm_sm_private_certificate.md) | resource | Manages ibm_sm_private_certificate. |
| [ibm_sm_private_certificate_configuration_action_set_signed](ibm_sm_private_certificate_configuration_action_set_signed.md) | resource | Manages ibm_sm_private_certificate_configuration_action_set_signed. |
| [ibm_sm_private_certificate_configuration_action_sign_csr](ibm_sm_private_certificate_configuration_action_sign_csr.md) | resource | Manages ibm_sm_private_certificate_configuration_action_sign_csr. |
| [ibm_sm_private_certificate_configuration_intermediate_ca_data](data_ibm_sm_private_certificate_configuration_intermediate_ca.md) | data | Reads ibm_sm_private_certificate_configuration_intermediate_ca. |
| [ibm_sm_private_certificate_configuration_intermediate_ca](ibm_sm_private_certificate_configuration_intermediate_ca.md) | resource | Manages ibm_sm_private_certificate_configuration_intermediate_ca. |
| [ibm_sm_private_certificate_configuration_root_ca_data](data_ibm_sm_private_certificate_configuration_root_ca.md) | data | Reads ibm_sm_private_certificate_configuration_root_ca. |
| [ibm_sm_private_certificate_configuration_root_ca](ibm_sm_private_certificate_configuration_root_ca.md) | resource | Manages ibm_sm_private_certificate_configuration_root_ca. |
| [ibm_sm_private_certificate_configuration_template_data](data_ibm_sm_private_certificate_configuration_template.md) | data | Reads ibm_sm_private_certificate_configuration_template. |
| [ibm_sm_private_certificate_configuration_template](ibm_sm_private_certificate_configuration_template.md) | resource | Manages ibm_sm_private_certificate_configuration_template. |
| [ibm_sm_private_certificate_metadata_data](data_ibm_sm_private_certificate_metadata.md) | data | Reads ibm_sm_private_certificate_metadata. |
| [ibm_sm_public_certificate_data](data_ibm_sm_public_certificate.md) | data | Reads ibm_sm_public_certificate. |
| [ibm_sm_public_certificate](ibm_sm_public_certificate.md) | resource | Manages ibm_sm_public_certificate. |
| [ibm_sm_public_certificate_action_validate_manual_dns](ibm_sm_public_certificate_action_validate_manual_dns.md) | resource | Schema validation for the ibm_sm_public_certificate_action_validate_manual_dns resource. |
| [ibm_sm_public_certificate_configuration_ca_lets_encrypt_data](data_ibm_sm_public_certificate_configuration_ca_lets_encrypt.md) | data | Reads ibm_sm_public_certificate_configuration_ca_lets_encrypt. |
| [ibm_sm_public_certificate_configuration_ca_lets_encrypt](ibm_sm_public_certificate_configuration_ca_lets_encrypt.md) | resource | Manages ibm_sm_public_certificate_configuration_ca_lets_encrypt. |
| [ibm_sm_public_certificate_configuration_dns_cis_data](data_ibm_sm_public_certificate_configuration_dns_cis.md) | data | Reads ibm_sm_public_certificate_configuration_dns_cis. |
| [ibm_sm_public_certificate_configuration_dns_cis](ibm_sm_public_certificate_configuration_dns_cis.md) | resource | Manages ibm_sm_public_certificate_configuration_dns_cis. |
| [ibm_sm_public_certificate_configuration_dns_classic_infrastructure_data](data_ibm_sm_public_certificate_configuration_dns_classic_infrastructure.md) | data | Reads ibm_sm_public_certificate_configuration_dns_classic_infrastructure. |
| [ibm_sm_public_certificate_configuration_dns_classic_infrastructure](ibm_sm_public_certificate_configuration_dns_classic_infrastructure.md) | resource | Manages ibm_sm_public_certificate_configuration_dns_classic_infrastructure. |
| [ibm_sm_public_certificate_metadata_data](data_ibm_sm_public_certificate_metadata.md) | data | Reads ibm_sm_public_certificate_metadata. |
| [ibm_sm_secret_group_data](data_ibm_sm_secret_group.md) | data | Reads ibm_sm_secret_group. |
| [ibm_sm_secret_group](ibm_sm_secret_group.md) | resource | Manages ibm_sm_secret_group. |
| [ibm_sm_secret_groups_data](data_ibm_sm_secret_groups.md) | data | Schema validation for the ibm_sm_secret_groups data source. |
| [ibm_sm_secrets_data](data_ibm_sm_secrets.md) | data | Schema validation for the ibm_sm_secrets data source. |
| [ibm_sm_service_credentials_secret_data](data_ibm_sm_service_credentials_secret.md) | data | Reads ibm_sm_service_credentials_secret. |
| [ibm_sm_service_credentials_secret](ibm_sm_service_credentials_secret.md) | resource | Manages ibm_sm_service_credentials_secret. |
| [ibm_sm_service_credentials_secret_metadata_data](data_ibm_sm_service_credentials_secret_metadata.md) | data | Reads ibm_sm_service_credentials_secret_metadata. |
| [ibm_sm_username_password_secret_data](data_ibm_sm_username_password_secret.md) | data | Reads ibm_sm_username_password_secret. |
| [ibm_sm_username_password_secret](ibm_sm_username_password_secret.md) | resource | Manages ibm_sm_username_password_secret. |
| [ibm_sm_username_password_secret_metadata_data](data_ibm_sm_username_password_secret_metadata.md) | data | Reads ibm_sm_username_password_secret_metadata. |

## secretsmanagerinstancemanagement

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_sm_admin_token](ibm_sm_admin_token.md) | resource | Schema validation for the ibm_sm_admin_token resource. |
| [ibm_sm_instance_data](data_ibm_sm_instance.md) | data | Reads ibm_sm_instance. |

## transitgateway

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_tg_connection](ibm_tg_connection.md) | resource | Manages ibm_tg_connection. |
| [ibm_tg_connection_action](ibm_tg_connection_action.md) | resource | Schema validation for the ibm_tg_connection_action resource. |
| [ibm_tg_connection_prefix_filter_data](data_ibm_tg_connection_prefix_filter.md) | data | Schema validation for the ibm_tg_connection_prefix_filter data source. |
| [ibm_tg_connection_prefix_filter](ibm_tg_connection_prefix_filter.md) | resource | Schema validation for the ibm_tg_connection_prefix_filter resource. |
| [ibm_tg_connection_prefix_filters_data](data_ibm_tg_connection_prefix_filters.md) | data | Schema validation for the ibm_tg_connection_prefix_filters data source. |
| [ibm_tg_connection_rgre_tunnel](ibm_tg_connection_rgre_tunnel.md) | resource | Manages ibm_tg_connection_rgre_tunnel. |
| [ibm_tg_gateway_data](data_ibm_tg_gateway.md) | data | Reads ibm_tg_gateway. |
| [ibm_tg_gateway](ibm_tg_gateway.md) | resource | Manages ibm_tg_gateway. |
| [ibm_tg_gateways_data](data_ibm_tg_gateways.md) | data | Schema validation for the ibm_tg_gateways data source. |
| [ibm_tg_location_data](data_ibm_tg_location.md) | data | Reads ibm_tg_location. |
| [ibm_tg_locations_data](data_ibm_tg_locations.md) | data | Schema validation for the ibm_tg_locations data source. |
| [ibm_tg_route_report_data](data_ibm_tg_route_report.md) | data | Schema validation for the ibm_tg_route_report data source. |
| [ibm_tg_route_report](ibm_tg_route_report.md) | resource | Schema validation for the ibm_tg_route_report resource. |
| [ibm_tg_route_reports_data](data_ibm_tg_route_reports.md) | data | Schema validation for the ibm_tg_route_reports data source. |

## usagereports

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_billing_report_snapshot](ibm_billing_report_snapshot.md) | resource | Schema validation for the ibm_billing_report_snapshot resource. |
| [ibm_billing_snapshot_list_data](data_ibm_billing_snapshot_list.md) | data | Schema validation for the ibm_billing_snapshot_list data source. |

## vmware

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_vmaas_transit_gateway_connection_data](data_ibm_vmaas_transit_gateway_connection.md) | data | Schema validation for the ibm_vmaas_transit_gateway_connection data source. |
| [ibm_vmaas_transit_gateway_connection](ibm_vmaas_transit_gateway_connection.md) | resource | Schema validation for the ibm_vmaas_transit_gateway_connection resource. |
| [ibm_vmaas_vdc_data](data_ibm_vmaas_vdc.md) | data | Reads ibm_vmaas_vdc. |
| [ibm_vmaas_vdc](ibm_vmaas_vdc.md) | resource | Manages ibm_vmaas_vdc. |

## vpc

| Rule | Kind | Description |
| --- | --- | --- |
| [ibm_is_backup_policies_data](data_ibm_is_backup_policies.md) | data | Reads ibm_is_backup_policies. |
| [ibm_is_backup_policy_data](data_ibm_is_backup_policy.md) | data | Reads ibm_is_backup_policy. |
| [ibm_is_backup_policy](ibm_is_backup_policy.md) | resource | Manages ibm_is_backup_policy. |
| [ibm_is_backup_policy_job_data](data_ibm_is_backup_policy_job.md) | data | Schema validation for the ibm_is_backup_policy_job data source. |
| [ibm_is_backup_policy_jobs_data](data_ibm_is_backup_policy_jobs.md) | data | Schema validation for the ibm_is_backup_policy_jobs data source. |
| [ibm_is_backup_policy_plan_data](data_ibm_is_backup_policy_plan.md) | data | Reads ibm_is_backup_policy_plan. |
| [ibm_is_backup_policy_plan](ibm_is_backup_policy_plan.md) | resource | Manages ibm_is_backup_policy_plan. |
| [ibm_is_backup_policy_plans_data](data_ibm_is_backup_policy_plans.md) | data | Reads ibm_is_backup_policy_plans. |
| [ibm_is_bare_metal_server_data](data_ibm_is_bare_metal_server.md) | data | Reads ibm_is_bare_metal_server. |
| [ibm_is_bare_metal_server](ibm_is_bare_metal_server.md) | resource | Manages ibm_is_bare_metal_server. |
| [ibm_is_bare_metal_server_action](ibm_is_bare_metal_server_action.md) | resource | Schema validation for the ibm_is_bare_metal_server_action resource. |
| [ibm_is_bare_metal_server_capacities_data](data_ibm_is_bare_metal_server_capacities.md) | data | Schema validation for the ibm_is_bare_metal_server_capacities data source. |
| [ibm_is_bare_metal_server_disk_data](data_ibm_is_bare_metal_server_disk.md) | data | Reads ibm_is_bare_metal_server_disk. |
| [ibm_is_bare_metal_server_disk](ibm_is_bare_metal_server_disk.md) | resource | Manages ibm_is_bare_metal_server_disk. |
| [ibm_is_bare_metal_server_disks_data](data_ibm_is_bare_metal_server_disks.md) | data | Schema validation for the ibm_is_bare_metal_server_disks data source. |
| [ibm_is_bare_metal_server_initialization_data](data_ibm_is_bare_metal_server_initialization.md) | data | Schema validation for the ibm_is_bare_metal_server_initialization data source. |
| [ibm_is_bare_metal_server_initialization](ibm_is_bare_metal_server_initialization.md) | resource | Schema validation for the ibm_is_bare_metal_server_initialization resource. |
| [ibm_is_bare_metal_server_network_attachment_data](data_ibm_is_bare_metal_server_network_attachment.md) | data | Reads ibm_is_bare_metal_server_network_attachment. |
| [ibm_is_bare_metal_server_network_attachment](ibm_is_bare_metal_server_network_attachment.md) | resource | Manages ibm_is_bare_metal_server_network_attachment. |
| [ibm_is_bare_metal_server_network_attachments_data](data_ibm_is_bare_metal_server_network_attachments.md) | data | Schema validation for the ibm_is_bare_metal_server_network_attachments data source. |
| [ibm_is_bare_metal_server_network_interface_data](data_ibm_is_bare_metal_server_network_interface.md) | data | Reads ibm_is_bare_metal_server_network_interface. |
| [ibm_is_bare_metal_server_network_interface](ibm_is_bare_metal_server_network_interface.md) | resource | Manages ibm_is_bare_metal_server_network_interface. |
| [ibm_is_bare_metal_server_network_interface_allow_float](ibm_is_bare_metal_server_network_interface_allow_float.md) | resource | Manages ibm_is_bare_metal_server_network_interface_allow_float. |
| [ibm_is_bare_metal_server_network_interface_floating_ip_data](data_ibm_is_bare_metal_server_network_interface_floating_ip.md) | data | Reads ibm_is_bare_metal_server_network_interface_floating_ip. |
| [ibm_is_bare_metal_server_network_interface_floating_ip](ibm_is_bare_metal_server_network_interface_floating_ip.md) | resource | Manages ibm_is_bare_metal_server_network_interface_floating_ip. |
| [ibm_is_bare_metal_server_network_interface_floating_ips_data](data_ibm_is_bare_metal_server_network_interface_floating_ips.md) | data | Schema validation for the ibm_is_bare_metal_server_network_interface_floating_ips data source. |
| [ibm_is_bare_metal_server_network_interface_reserved_ip_data](data_ibm_is_bare_metal_server_network_interface_reserved_ip.md) | data | Reads ibm_is_bare_metal_server_network_interface_reserved_ip. |
| [ibm_is_bare_metal_server_network_interface_reserved_ips_data](data_ibm_is_bare_metal_server_network_interface_reserved_ips.md) | data | Schema validation for the ibm_is_bare_metal_server_network_interface_reserved_ips data source. |
| [ibm_is_bare_metal_server_network_interfaces_data](data_ibm_is_bare_metal_server_network_interfaces.md) | data | Schema validation for the ibm_is_bare_metal_server_network_interfaces data source. |
| [ibm_is_bare_metal_server_profile_data](data_ibm_is_bare_metal_server_profile.md) | data | Reads ibm_is_bare_metal_server_profile. |
| [ibm_is_bare_metal_server_profiles_data](data_ibm_is_bare_metal_server_profiles.md) | data | Schema validation for the ibm_is_bare_metal_server_profiles data source. |
| [ibm_is_bare_metal_servers_data](data_ibm_is_bare_metal_servers.md) | data | Reads ibm_is_bare_metal_servers. |
| [ibm_is_cluster_network_data](data_ibm_is_cluster_network.md) | data | Reads ibm_is_cluster_network. |
| [ibm_is_cluster_network](ibm_is_cluster_network.md) | resource | Manages ibm_is_cluster_network. |
| [ibm_is_cluster_network_interface_data](data_ibm_is_cluster_network_interface.md) | data | Reads ibm_is_cluster_network_interface. |
| [ibm_is_cluster_network_interface](ibm_is_cluster_network_interface.md) | resource | Manages ibm_is_cluster_network_interface. |
| [ibm_is_cluster_network_interfaces_data](data_ibm_is_cluster_network_interfaces.md) | data | Reads ibm_is_cluster_network_interfaces. |
| [ibm_is_cluster_network_profile_data](data_ibm_is_cluster_network_profile.md) | data | Reads ibm_is_cluster_network_profile. |
| [ibm_is_cluster_network_profiles_data](data_ibm_is_cluster_network_profiles.md) | data | Schema validation for the ibm_is_cluster_network_profiles data source. |
| [ibm_is_cluster_network_subnet_data](data_ibm_is_cluster_network_subnet.md) | data | Reads ibm_is_cluster_network_subnet. |
| [ibm_is_cluster_network_subnet](ibm_is_cluster_network_subnet.md) | resource | Manages ibm_is_cluster_network_subnet. |
| [ibm_is_cluster_network_subnet_reserved_ip_data](data_ibm_is_cluster_network_subnet_reserved_ip.md) | data | Reads ibm_is_cluster_network_subnet_reserved_ip. |
| [ibm_is_cluster_network_subnet_reserved_ip](ibm_is_cluster_network_subnet_reserved_ip.md) | resource | Manages ibm_is_cluster_network_subnet_reserved_ip. |
| [ibm_is_cluster_network_subnet_reserved_ips_data](data_ibm_is_cluster_network_subnet_reserved_ips.md) | data | Reads ibm_is_cluster_network_subnet_reserved_ips. |
| [ibm_is_cluster_network_subnets_data](data_ibm_is_cluster_network_subnets.md) | data | Reads ibm_is_cluster_network_subnets. |
| [ibm_is_cluster_networks_data](data_ibm_is_cluster_networks.md) | data | Reads ibm_is_cluster_networks. |
| [ibm_is_dedicated_host_data](data_ibm_is_dedicated_host.md) | data | Reads ibm_is_dedicated_host. |
| [ibm_is_dedicated_host](ibm_is_dedicated_host.md) | resource | Manages ibm_is_dedicated_host. |
| [ibm_is_dedicated_host_disk_data](data_ibm_is_dedicated_host_disk.md) | data | Reads ibm_is_dedicated_host_disk. |
| [ibm_is_dedicated_host_disk_management](ibm_is_dedicated_host_disk_management.md) | resource | Schema validation for the ibm_is_dedicated_host_disk_management resource. |
| [ibm_is_dedicated_host_disks_data](data_ibm_is_dedicated_host_disks.md) | data | Schema validation for the ibm_is_dedicated_host_disks data source. |
| [ibm_is_dedicated_host_group_data](data_ibm_is_dedicated_host_group.md) | data | Reads ibm_is_dedicated_host_group. |
| [ibm_is_dedicated_host_group](ibm_is_dedicated_host_group.md) | resource | Manages ibm_is_dedicated_host_group. |
| [ibm_is_dedicated_host_groups_data](data_ibm_is_dedicated_host_groups.md) | data | Reads ibm_is_dedicated_host_groups. |
| [ibm_is_dedicated_host_profile_data](data_ibm_is_dedicated_host_profile.md) | data | Reads ibm_is_dedicated_host_profile. |
| [ibm_is_dedicated_host_profiles_data](data_ibm_is_dedicated_host_profiles.md) | data | Schema validation for the ibm_is_dedicated_host_profiles data source. |
| [ibm_is_dedicated_hosts_data](data_ibm_is_dedicated_hosts.md) | data | Reads ibm_is_dedicated_hosts. |
| [ibm_is_endpoint_gateway_targets_data](data_ibm_is_endpoint_gateway_targets.md) | data | Schema validation for the ibm_is_endpoint_gateway_targets data source. |
| [ibm_is_floating_ip_data](data_ibm_is_floating_ip.md) | data | Reads ibm_is_floating_ip. |
| [ibm_is_floating_ip](ibm_is_floating_ip.md) | resource | Manages ibm_is_floating_ip. |
| [ibm_is_floating_ips_data](data_ibm_is_floating_ips.md) | data | Reads ibm_is_floating_ips. |
| [ibm_is_flow_log_data](data_ibm_is_flow_log.md) | data | Reads ibm_is_flow_log. |
| [ibm_is_flow_log](ibm_is_flow_log.md) | resource | Manages ibm_is_flow_log. |
| [ibm_is_flow_logs_data](data_ibm_is_flow_logs.md) | data | Reads ibm_is_flow_logs. |
| [ibm_is_ike_policies_data](data_ibm_is_ike_policies.md) | data | Schema validation for the ibm_is_ike_policies data source. |
| [ibm_is_ike_policy_data](data_ibm_is_ike_policy.md) | data | Reads ibm_is_ike_policy. |
| [ibm_is_ike_policy](ibm_is_ike_policy.md) | resource | Manages ibm_is_ike_policy. |
| [ibm_is_image_data](data_ibm_is_image.md) | data | Reads ibm_is_image. |
| [ibm_is_image](ibm_is_image.md) | resource | Manages ibm_is_image. |
| [ibm_is_image_bare_metal_server_profiles_data](data_ibm_is_image_bare_metal_server_profiles.md) | data | Schema validation for the ibm_is_image_bare_metal_server_profiles data source. |
| [ibm_is_image_deprecate](ibm_is_image_deprecate.md) | resource | Manages ibm_is_image_deprecate. |
| [ibm_is_image_export_job_data](data_ibm_is_image_export_job.md) | data | Reads ibm_is_image_export_job. |
| [ibm_is_image_export_job](ibm_is_image_export_job.md) | resource | Manages ibm_is_image_export_job. |
| [ibm_is_image_export_jobs_data](data_ibm_is_image_export_jobs.md) | data | Schema validation for the ibm_is_image_export_jobs data source. |
| [ibm_is_image_instance_profiles_data](data_ibm_is_image_instance_profiles.md) | data | Schema validation for the ibm_is_image_instance_profiles data source. |
| [ibm_is_image_obsolete](ibm_is_image_obsolete.md) | resource | Manages ibm_is_image_obsolete. |
| [ibm_is_images_data](data_ibm_is_images.md) | data | Reads ibm_is_images. |
| [ibm_is_instance_data](data_ibm_is_instance.md) | data | Reads ibm_is_instance. |
| [ibm_is_instance](ibm_is_instance.md) | resource | Manages ibm_is_instance. |
| [ibm_is_instance_action](ibm_is_instance_action.md) | resource | Schema validation for the ibm_is_instance_action resource. |
| [ibm_is_instance_cluster_network_attachment_data](data_ibm_is_instance_cluster_network_attachment.md) | data | Reads ibm_is_instance_cluster_network_attachment. |
| [ibm_is_instance_cluster_network_attachment](ibm_is_instance_cluster_network_attachment.md) | resource | Manages ibm_is_instance_cluster_network_attachment. |
| [ibm_is_instance_cluster_network_attachments_data](data_ibm_is_instance_cluster_network_attachments.md) | data | Schema validation for the ibm_is_instance_cluster_network_attachments data source. |
| [ibm_is_instance_disk_data](data_ibm_is_instance_disk.md) | data | Reads ibm_is_instance_disk. |
| [ibm_is_instance_disk_management](ibm_is_instance_disk_management.md) | resource | Schema validation for the ibm_is_instance_disk_management resource. |
| [ibm_is_instance_disks_data](data_ibm_is_instance_disks.md) | data | Schema validation for the ibm_is_instance_disks data source. |
| [ibm_is_instance_group_data](data_ibm_is_instance_group.md) | data | Reads ibm_is_instance_group. |
| [ibm_is_instance_group](ibm_is_instance_group.md) | resource | Manages ibm_is_instance_group. |
| [ibm_is_instance_group_manager_data](data_ibm_is_instance_group_manager.md) | data | Reads ibm_is_instance_group_manager. |
| [ibm_is_instance_group_manager](ibm_is_instance_group_manager.md) | resource | Manages ibm_is_instance_group_manager. |
| [ibm_is_instance_group_manager_action_data](data_ibm_is_instance_group_manager_action.md) | data | Reads ibm_is_instance_group_manager_action. |
| [ibm_is_instance_group_manager_action](ibm_is_instance_group_manager_action.md) | resource | Manages ibm_is_instance_group_manager_action. |
| [ibm_is_instance_group_manager_actions_data](data_ibm_is_instance_group_manager_actions.md) | data | Schema validation for the ibm_is_instance_group_manager_actions data source. |
| [ibm_is_instance_group_manager_policies_data](data_ibm_is_instance_group_manager_policies.md) | data | Schema validation for the ibm_is_instance_group_manager_policies data source. |
| [ibm_is_instance_group_manager_policy_data](data_ibm_is_instance_group_manager_policy.md) | data | Reads ibm_is_instance_group_manager_policy. |
| [ibm_is_instance_group_manager_policy](ibm_is_instance_group_manager_policy.md) | resource | Manages ibm_is_instance_group_manager_policy. |
| [ibm_is_instance_group_managers_data](data_ibm_is_instance_group_managers.md) | data | Schema validation for the ibm_is_instance_group_managers data source. |
| [ibm_is_instance_group_membership_data](data_ibm_is_instance_group_membership.md) | data | Reads ibm_is_instance_group_membership. |
| [ibm_is_instance_group_membership](ibm_is_instance_group_membership.md) | resource | Manages ibm_is_instance_group_membership. |
| [ibm_is_instance_group_memberships_data](data_ibm_is_instance_group_memberships.md) | data | Schema validation for the ibm_is_instance_group_memberships data source. |
| [ibm_is_instance_groups_data](data_ibm_is_instance_groups.md) | data | Schema validation for the ibm_is_instance_groups data source. |
| [ibm_is_instance_network_attachment_data](data_ibm_is_instance_network_attachment.md) | data | Reads ibm_is_instance_network_attachment. |
| [ibm_is_instance_network_attachment](ibm_is_instance_network_attachment.md) | resource | Manages ibm_is_instance_network_attachment. |
| [ibm_is_instance_network_attachments_data](data_ibm_is_instance_network_attachments.md) | data | Schema validation for the ibm_is_instance_network_attachments data source. |
| [ibm_is_instance_network_interface_data](data_ibm_is_instance_network_interface.md) | data | Reads ibm_is_instance_network_interface. |
| [ibm_is_instance_network_interface](ibm_is_instance_network_interface.md) | resource | Manages ibm_is_instance_network_interface. |
| [ibm_is_instance_network_interface_floating_ip](ibm_is_instance_network_interface_floating_ip.md) | resource | Manages ibm_is_instance_network_interface_floating_ip. |
| [ibm_is_instance_network_interface_reserved_ip_data](data_ibm_is_instance_network_interface_reserved_ip.md) | data | Reads ibm_is_instance_network_interface_reserved_ip. |
| [ibm_is_instance_network_interface_reserved_ips_data](data_ibm_is_instance_network_interface_reserved_ips.md) | data | Schema validation for the ibm_is_instance_network_interface_reserved_ips data source. |
| [ibm_is_instance_network_interfaces_data](data_ibm_is_instance_network_interfaces.md) | data | Schema validation for the ibm_is_instance_network_interfaces data source. |
| [ibm_is_instance_profile_data](data_ibm_is_instance_profile.md) | data | Schema validation for the ibm_is_instance_profile data source. |
| [ibm_is_instance_profiles_data](data_ibm_is_instance_profiles.md) | data | Schema validation for the ibm_is_instance_profiles data source. |
| [ibm_is_instance_reinitialize](ibm_is_instance_reinitialize.md) | resource | Schema validation for the ibm_is_instance_reinitialize resource. |
| [ibm_is_instance_software_attachment_data](data_ibm_is_instance_software_attachment.md) | data | Reads ibm_is_instance_software_attachment. |
| [ibm_is_instance_software_attachment](ibm_is_instance_software_attachment.md) | resource | Manages ibm_is_instance_software_attachment. |
| [ibm_is_instance_software_attachments_data](data_ibm_is_instance_software_attachments.md) | data | Schema validation for the ibm_is_instance_software_attachments data source. |
| [ibm_is_instance_template_data](data_ibm_is_instance_template.md) | data | Schema validation for the ibm_is_instance_template data source. |
| [ibm_is_instance_template](ibm_is_instance_template.md) | resource | Manages ibm_is_instance_template. |
| [ibm_is_instance_templates_data](data_ibm_is_instance_templates.md) | data | Schema validation for the ibm_is_instance_templates data source. |
| [ibm_is_instance_volume_attachment_data](data_ibm_is_instance_volume_attachment.md) | data | Reads ibm_is_instance_volume_attachment. |
| [ibm_is_instance_volume_attachment](ibm_is_instance_volume_attachment.md) | resource | Manages ibm_is_instance_volume_attachment. |
| [ibm_is_instance_volume_attachments_data](data_ibm_is_instance_volume_attachments.md) | data | Schema validation for the ibm_is_instance_volume_attachments data source. |
| [ibm_is_instances_data](data_ibm_is_instances.md) | data | Schema validation for the ibm_is_instances data source. |
| [ibm_is_ipsec_policies_data](data_ibm_is_ipsec_policies.md) | data | Schema validation for the ibm_is_ipsec_policies data source. |
| [ibm_is_ipsec_policy_data](data_ibm_is_ipsec_policy.md) | data | Reads ibm_is_ipsec_policy. |
| [ibm_is_ipsec_policy](ibm_is_ipsec_policy.md) | resource | Manages ibm_is_ipsec_policy. |
| [ibm_is_lb_data](data_ibm_is_lb.md) | data | Reads ibm_is_lb. |
| [ibm_is_lb](ibm_is_lb.md) | resource | Manages ibm_is_lb. |
| [ibm_is_lb_listener_data](data_ibm_is_lb_listener.md) | data | Schema validation for the ibm_is_lb_listener data source. |
| [ibm_is_lb_listener](ibm_is_lb_listener.md) | resource | Schema validation for the ibm_is_lb_listener resource. |
| [ibm_is_lb_listener_policies_data](data_ibm_is_lb_listener_policies.md) | data | Schema validation for the ibm_is_lb_listener_policies data source. |
| [ibm_is_lb_listener_policy_data](data_ibm_is_lb_listener_policy.md) | data | Reads ibm_is_lb_listener_policy. |
| [ibm_is_lb_listener_policy](ibm_is_lb_listener_policy.md) | resource | Manages ibm_is_lb_listener_policy. |
| [ibm_is_lb_listener_policy_rule_data](data_ibm_is_lb_listener_policy_rule.md) | data | Schema validation for the ibm_is_lb_listener_policy_rule data source. |
| [ibm_is_lb_listener_policy_rule](ibm_is_lb_listener_policy_rule.md) | resource | Schema validation for the ibm_is_lb_listener_policy_rule resource. |
| [ibm_is_lb_listener_policy_rules_data](data_ibm_is_lb_listener_policy_rules.md) | data | Schema validation for the ibm_is_lb_listener_policy_rules data source. |
| [ibm_is_lb_listeners_data](data_ibm_is_lb_listeners.md) | data | Schema validation for the ibm_is_lb_listeners data source. |
| [ibm_is_lb_pool_data](data_ibm_is_lb_pool.md) | data | Reads ibm_is_lb_pool. |
| [ibm_is_lb_pool](ibm_is_lb_pool.md) | resource | Manages ibm_is_lb_pool. |
| [ibm_is_lb_pool_member_data](data_ibm_is_lb_pool_member.md) | data | Schema validation for the ibm_is_lb_pool_member data source. |
| [ibm_is_lb_pool_member](ibm_is_lb_pool_member.md) | resource | Schema validation for the ibm_is_lb_pool_member resource. |
| [ibm_is_lb_pool_members_data](data_ibm_is_lb_pool_members.md) | data | Schema validation for the ibm_is_lb_pool_members data source. |
| [ibm_is_lb_pools_data](data_ibm_is_lb_pools.md) | data | Schema validation for the ibm_is_lb_pools data source. |
| [ibm_is_lb_profile_data](data_ibm_is_lb_profile.md) | data | Reads ibm_is_lb_profile. |
| [ibm_is_lb_profiles_data](data_ibm_is_lb_profiles.md) | data | Reads ibm_is_lb_profiles. |
| [ibm_is_lbs_data](data_ibm_is_lbs.md) | data | Schema validation for the ibm_is_lbs data source. |
| [ibm_is_network_acl_data](data_ibm_is_network_acl.md) | data | Reads ibm_is_network_acl. |
| [ibm_is_network_acl](ibm_is_network_acl.md) | resource | Manages ibm_is_network_acl. |
| [ibm_is_network_acl_rule_data](data_ibm_is_network_acl_rule.md) | data | Reads ibm_is_network_acl_rule. |
| [ibm_is_network_acl_rule](ibm_is_network_acl_rule.md) | resource | Manages ibm_is_network_acl_rule. |
| [ibm_is_network_acl_rules_data](data_ibm_is_network_acl_rules.md) | data | Schema validation for the ibm_is_network_acl_rules data source. |
| [ibm_is_network_acls_data](data_ibm_is_network_acls.md) | data | Schema validation for the ibm_is_network_acls data source. |
| [ibm_is_operating_system_data](data_ibm_is_operating_system.md) | data | Reads ibm_is_operating_system. |
| [ibm_is_operating_systems_data](data_ibm_is_operating_systems.md) | data | Schema validation for the ibm_is_operating_systems data source. |
| [ibm_is_placement_group_data](data_ibm_is_placement_group.md) | data | Reads ibm_is_placement_group. |
| [ibm_is_placement_group](ibm_is_placement_group.md) | resource | Manages ibm_is_placement_group. |
| [ibm_is_placement_groups_data](data_ibm_is_placement_groups.md) | data | Schema validation for the ibm_is_placement_groups data source. |
| [ibm_is_private_path_service_gateway_data](data_ibm_is_private_path_service_gateway.md) | data | Reads ibm_is_private_path_service_gateway. |
| [ibm_is_private_path_service_gateway](ibm_is_private_path_service_gateway.md) | resource | Manages ibm_is_private_path_service_gateway. |
| [ibm_is_private_path_service_gateway_account_policies_data](data_ibm_is_private_path_service_gateway_account_policies.md) | data | Schema validation for the ibm_is_private_path_service_gateway_account_policies data source. |
| [ibm_is_private_path_service_gateway_account_policy_data](data_ibm_is_private_path_service_gateway_account_policy.md) | data | Schema validation for the ibm_is_private_path_service_gateway_account_policy data source. |
| [ibm_is_private_path_service_gateway_account_policy](ibm_is_private_path_service_gateway_account_policy.md) | resource | Schema validation for the ibm_is_private_path_service_gateway_account_policy resource. |
| [ibm_is_private_path_service_gateway_endpoint_gateway_binding_data](data_ibm_is_private_path_service_gateway_endpoint_gateway_binding.md) | data | Schema validation for the ibm_is_private_path_service_gateway_endpoint_gateway_binding data source. |
| [ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations](ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations.md) | resource | Schema validation for the ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations resource. |
| [ibm_is_private_path_service_gateway_endpoint_gateway_bindings_data](data_ibm_is_private_path_service_gateway_endpoint_gateway_bindings.md) | data | Schema validation for the ibm_is_private_path_service_gateway_endpoint_gateway_bindings data source. |
| [ibm_is_private_path_service_gateway_operations](ibm_is_private_path_service_gateway_operations.md) | resource | Schema validation for the ibm_is_private_path_service_gateway_operations resource. |
| [ibm_is_private_path_service_gateway_revoke_account](ibm_is_private_path_service_gateway_revoke_account.md) | resource | Schema validation for the ibm_is_private_path_service_gateway_revoke_account resource. |
| [ibm_is_private_path_service_gateways_data](data_ibm_is_private_path_service_gateways.md) | data | Schema validation for the ibm_is_private_path_service_gateways data source. |
| [ibm_is_public_address_range_data](data_ibm_is_public_address_range.md) | data | Reads ibm_is_public_address_range. |
| [ibm_is_public_address_range](ibm_is_public_address_range.md) | resource | Manages ibm_is_public_address_range. |
| [ibm_is_public_address_ranges_data](data_ibm_is_public_address_ranges.md) | data | Schema validation for the ibm_is_public_address_ranges data source. |
| [ibm_is_public_gateway_data](data_ibm_is_public_gateway.md) | data | Reads ibm_is_public_gateway. |
| [ibm_is_public_gateway](ibm_is_public_gateway.md) | resource | Manages ibm_is_public_gateway. |
| [ibm_is_public_gateways_data](data_ibm_is_public_gateways.md) | data | Schema validation for the ibm_is_public_gateways data source. |
| [ibm_is_region_data](data_ibm_is_region.md) | data | Schema validation for the ibm_is_region data source. |
| [ibm_is_regions_data](data_ibm_is_regions.md) | data | Schema validation for the ibm_is_regions data source. |
| [ibm_is_reservation_data](data_ibm_is_reservation.md) | data | Reads ibm_is_reservation. |
| [ibm_is_reservation](ibm_is_reservation.md) | resource | Manages ibm_is_reservation. |
| [ibm_is_reservation_activate](ibm_is_reservation_activate.md) | resource | Manages ibm_is_reservation_activate. |
| [ibm_is_reservations_data](data_ibm_is_reservations.md) | data | Reads ibm_is_reservations. |
| [ibm_is_security_group_data](data_ibm_is_security_group.md) | data | Reads ibm_is_security_group. |
| [ibm_is_security_group](ibm_is_security_group.md) | resource | Manages ibm_is_security_group. |
| [ibm_is_security_group_rule_data](data_ibm_is_security_group_rule.md) | data | Reads ibm_is_security_group_rule. |
| [ibm_is_security_group_rule](ibm_is_security_group_rule.md) | resource | Manages ibm_is_security_group_rule. |
| [ibm_is_security_group_rules_data](data_ibm_is_security_group_rules.md) | data | Schema validation for the ibm_is_security_group_rules data source. |
| [ibm_is_security_group_target_data](data_ibm_is_security_group_target.md) | data | Reads ibm_is_security_group_target. |
| [ibm_is_security_group_target](ibm_is_security_group_target.md) | resource | Manages ibm_is_security_group_target. |
| [ibm_is_security_group_targets_data](data_ibm_is_security_group_targets.md) | data | Schema validation for the ibm_is_security_group_targets data source. |
| [ibm_is_security_groups_data](data_ibm_is_security_groups.md) | data | Schema validation for the ibm_is_security_groups data source. |
| [ibm_is_share_data](data_ibm_is_share.md) | data | Reads ibm_is_share. |
| [ibm_is_share](ibm_is_share.md) | resource | Manages ibm_is_share. |
| [ibm_is_share_accessor_binding_data](data_ibm_is_share_accessor_binding.md) | data | Schema validation for the ibm_is_share_accessor_binding data source. |
| [ibm_is_share_accessor_bindings_data](data_ibm_is_share_accessor_bindings.md) | data | Schema validation for the ibm_is_share_accessor_bindings data source. |
| [ibm_is_share_delete_accessor_binding](ibm_is_share_delete_accessor_binding.md) | resource | Schema validation for the ibm_is_share_delete_accessor_binding resource. |
| [ibm_is_share_mount_target_data](data_ibm_is_share_mount_target.md) | data | Reads ibm_is_share_mount_target. |
| [ibm_is_share_mount_target](ibm_is_share_mount_target.md) | resource | Manages ibm_is_share_mount_target. |
| [ibm_is_share_mount_targets_data](data_ibm_is_share_mount_targets.md) | data | Reads ibm_is_share_mount_targets. |
| [ibm_is_share_profile_data](data_ibm_is_share_profile.md) | data | Reads ibm_is_share_profile. |
| [ibm_is_share_profiles_data](data_ibm_is_share_profiles.md) | data | Schema validation for the ibm_is_share_profiles data source. |
| [ibm_is_share_replica_operations](ibm_is_share_replica_operations.md) | resource | Schema validation for the ibm_is_share_replica_operations resource. |
| [ibm_is_share_snapshot_data](data_ibm_is_share_snapshot.md) | data | Reads ibm_is_share_snapshot. |
| [ibm_is_share_snapshot](ibm_is_share_snapshot.md) | resource | Manages ibm_is_share_snapshot. |
| [ibm_is_share_snapshots_data](data_ibm_is_share_snapshots.md) | data | Reads ibm_is_share_snapshots. |
| [ibm_is_shares_data](data_ibm_is_shares.md) | data | Reads ibm_is_shares. |
| [ibm_is_snapshot_data](data_ibm_is_snapshot.md) | data | Reads ibm_is_snapshot. |
| [ibm_is_snapshot](ibm_is_snapshot.md) | resource | Manages ibm_is_snapshot. |
| [ibm_is_snapshot_clone_data](data_ibm_is_snapshot_clone.md) | data | Schema validation for the ibm_is_snapshot_clone data source. |
| [ibm_is_snapshot_clones_data](data_ibm_is_snapshot_clones.md) | data | Schema validation for the ibm_is_snapshot_clones data source. |
| [ibm_is_snapshot_consistency_group_data](data_ibm_is_snapshot_consistency_group.md) | data | Reads ibm_is_snapshot_consistency_group. |
| [ibm_is_snapshot_consistency_group](ibm_is_snapshot_consistency_group.md) | resource | Manages ibm_is_snapshot_consistency_group. |
| [ibm_is_snapshot_consistency_groups_data](data_ibm_is_snapshot_consistency_groups.md) | data | Reads ibm_is_snapshot_consistency_groups. |
| [ibm_is_snapshot_instance_profiles_data](data_ibm_is_snapshot_instance_profiles.md) | data | Schema validation for the ibm_is_snapshot_instance_profiles data source. |
| [ibm_is_snapshots_data](data_ibm_is_snapshots.md) | data | Reads ibm_is_snapshots. |
| [ibm_is_source_share_data](data_ibm_is_source_share.md) | data | Reads ibm_is_source_share. |
| [ibm_is_ssh_key_data](data_ibm_is_ssh_key.md) | data | Reads ibm_is_ssh_key. |
| [ibm_is_ssh_key](ibm_is_ssh_key.md) | resource | Manages ibm_is_ssh_key. |
| [ibm_is_ssh_keys_data](data_ibm_is_ssh_keys.md) | data | Schema validation for the ibm_is_ssh_keys data source. |
| [ibm_is_subnet_data](data_ibm_is_subnet.md) | data | Schema validation for the ibm_is_subnet data source. |
| [ibm_is_subnet](ibm_is_subnet.md) | resource | Manages ibm_is_subnet. |
| [ibm_is_subnet_network_acl_attachment](ibm_is_subnet_network_acl_attachment.md) | resource | Manages ibm_is_subnet_network_acl_attachment. |
| [ibm_is_subnet_public_gateway_attachment](ibm_is_subnet_public_gateway_attachment.md) | resource | Manages ibm_is_subnet_public_gateway_attachment. |
| [ibm_is_subnet_reserved_ip_data](data_ibm_is_subnet_reserved_ip.md) | data | Reads ibm_is_subnet_reserved_ip. |
| [ibm_is_subnet_reserved_ip](ibm_is_subnet_reserved_ip.md) | resource | Manages ibm_is_subnet_reserved_ip. |
| [ibm_is_subnet_reserved_ip_patch](ibm_is_subnet_reserved_ip_patch.md) | resource | Manages ibm_is_subnet_reserved_ip_patch. |
| [ibm_is_subnet_reserved_ips_data](data_ibm_is_subnet_reserved_ips.md) | data | Schema validation for the ibm_is_subnet_reserved_ips data source. |
| [ibm_is_subnet_routing_table_attachment](ibm_is_subnet_routing_table_attachment.md) | resource | Manages ibm_is_subnet_routing_table_attachment. |
| [ibm_is_subnets_data](data_ibm_is_subnets.md) | data | Schema validation for the ibm_is_subnets data source. |
| [ibm_is_virtual_endpoint_gateway_data](data_ibm_is_virtual_endpoint_gateway.md) | data | Reads ibm_is_virtual_endpoint_gateway. |
| [ibm_is_virtual_endpoint_gateway](ibm_is_virtual_endpoint_gateway.md) | resource | Manages ibm_is_virtual_endpoint_gateway. |
| [ibm_is_virtual_endpoint_gateway_ip](ibm_is_virtual_endpoint_gateway_ip.md) | resource | Manages ibm_is_virtual_endpoint_gateway_ip. |
| [ibm_is_virtual_endpoint_gateway_ips_data](data_ibm_is_virtual_endpoint_gateway_ips.md) | data | Schema validation for the ibm_is_virtual_endpoint_gateway_ips data source. |
| [ibm_is_virtual_endpoint_gateway_resource_binding_data](data_ibm_is_virtual_endpoint_gateway_resource_binding.md) | data | Reads ibm_is_virtual_endpoint_gateway_resource_binding. |
| [ibm_is_virtual_endpoint_gateway_resource_binding](ibm_is_virtual_endpoint_gateway_resource_binding.md) | resource | Manages ibm_is_virtual_endpoint_gateway_resource_binding. |
| [ibm_is_virtual_endpoint_gateway_resource_bindings_data](data_ibm_is_virtual_endpoint_gateway_resource_bindings.md) | data | Schema validation for the ibm_is_virtual_endpoint_gateway_resource_bindings data source. |
| [ibm_is_virtual_endpoint_gateways_data](data_ibm_is_virtual_endpoint_gateways.md) | data | Reads ibm_is_virtual_endpoint_gateways. |
| [ibm_is_virtual_network_interface_data](data_ibm_is_virtual_network_interface.md) | data | Reads ibm_is_virtual_network_interface. |
| [ibm_is_virtual_network_interface](ibm_is_virtual_network_interface.md) | resource | Manages ibm_is_virtual_network_interface. |
| [ibm_is_virtual_network_interface_floating_ip_data](data_ibm_is_virtual_network_interface_floating_ip.md) | data | Reads ibm_is_virtual_network_interface_floating_ip. |
| [ibm_is_virtual_network_interface_floating_ip](ibm_is_virtual_network_interface_floating_ip.md) | resource | Manages ibm_is_virtual_network_interface_floating_ip. |
| [ibm_is_virtual_network_interface_floating_ips_data](data_ibm_is_virtual_network_interface_floating_ips.md) | data | Schema validation for the ibm_is_virtual_network_interface_floating_ips data source. |
| [ibm_is_virtual_network_interface_ip_data](data_ibm_is_virtual_network_interface_ip.md) | data | Reads ibm_is_virtual_network_interface_ip. |
| [ibm_is_virtual_network_interface_ip](ibm_is_virtual_network_interface_ip.md) | resource | Manages ibm_is_virtual_network_interface_ip. |
| [ibm_is_virtual_network_interface_ips_data](data_ibm_is_virtual_network_interface_ips.md) | data | Schema validation for the ibm_is_virtual_network_interface_ips data source. |
| [ibm_is_virtual_network_interfaces_data](data_ibm_is_virtual_network_interfaces.md) | data | Schema validation for the ibm_is_virtual_network_interfaces data source. |
| [ibm_is_volume_data](data_ibm_is_volume.md) | data | Reads ibm_is_volume. |
| [ibm_is_volume](ibm_is_volume.md) | resource | Manages ibm_is_volume. |
| [ibm_is_volume_instance_profiles_data](data_ibm_is_volume_instance_profiles.md) | data | Schema validation for the ibm_is_volume_instance_profiles data source. |
| [ibm_is_volume_job_data](data_ibm_is_volume_job.md) | data | Reads ibm_is_volume_job. |
| [ibm_is_volume_job](ibm_is_volume_job.md) | resource | Manages ibm_is_volume_job. |
| [ibm_is_volume_job_cancel](ibm_is_volume_job_cancel.md) | resource | Manages ibm_is_volume_job_cancel. |
| [ibm_is_volume_jobs_data](data_ibm_is_volume_jobs.md) | data | Schema validation for the ibm_is_volume_jobs data source. |
| [ibm_is_volume_profile_data](data_ibm_is_volume_profile.md) | data | Reads ibm_is_volume_profile. |
| [ibm_is_volume_profiles_data](data_ibm_is_volume_profiles.md) | data | Schema validation for the ibm_is_volume_profiles data source. |
| [ibm_is_volumes_data](data_ibm_is_volumes.md) | data | Schema validation for the ibm_is_volumes data source. |
| [ibm_is_vpc_data](data_ibm_is_vpc.md) | data | Schema validation for the ibm_is_vpc data source. |
| [ibm_is_vpc](ibm_is_vpc.md) | resource | Manages ibm_is_vpc. |
| [ibm_is_vpc_address_prefix_data](data_ibm_is_vpc_address_prefix.md) | data | Reads ibm_is_vpc_address_prefix. |
| [ibm_is_vpc_address_prefix](ibm_is_vpc_address_prefix.md) | resource | Manages ibm_is_vpc_address_prefix. |
| [ibm_is_vpc_address_prefixes_data](data_ibm_is_vpc_address_prefixes.md) | data | Reads ibm_is_vpc_address_prefixes. |
| [ibm_is_vpc_default_routing_table_data](data_ibm_is_vpc_default_routing_table.md) | data | Reads ibm_is_vpc_default_routing_table. |
| [ibm_is_vpc_dns_resolution_binding_data](data_ibm_is_vpc_dns_resolution_binding.md) | data | Reads ibm_is_vpc_dns_resolution_binding. |
| [ibm_is_vpc_dns_resolution_binding](ibm_is_vpc_dns_resolution_binding.md) | resource | Manages ibm_is_vpc_dns_resolution_binding. |
| [ibm_is_vpc_dns_resolution_bindings_data](data_ibm_is_vpc_dns_resolution_bindings.md) | data | Schema validation for the ibm_is_vpc_dns_resolution_bindings data source. |
| [ibm_is_vpc_routing_table_data](data_ibm_is_vpc_routing_table.md) | data | Reads ibm_is_vpc_routing_table. |
| [ibm_is_vpc_routing_table](ibm_is_vpc_routing_table.md) | resource | Manages ibm_is_vpc_routing_table. |
| [ibm_is_vpc_routing_table_route_data](data_ibm_is_vpc_routing_table_route.md) | data | Reads ibm_is_vpc_routing_table_route. |
| [ibm_is_vpc_routing_table_route](ibm_is_vpc_routing_table_route.md) | resource | Manages ibm_is_vpc_routing_table_route. |
| [ibm_is_vpc_routing_table_routes_data](data_ibm_is_vpc_routing_table_routes.md) | data | Schema validation for the ibm_is_vpc_routing_table_routes data source. |
| [ibm_is_vpc_routing_tables_data](data_ibm_is_vpc_routing_tables.md) | data | Schema validation for the ibm_is_vpc_routing_tables data source. |
| [ibm_is_vpcs_data](data_ibm_is_vpcs.md) | data | Schema validation for the ibm_is_vpcs data source. |
| [ibm_is_vpn_gateway_data](data_ibm_is_vpn_gateway.md) | data | Reads ibm_is_vpn_gateway. |
| [ibm_is_vpn_gateway](ibm_is_vpn_gateway.md) | resource | Manages ibm_is_vpn_gateway. |
| [ibm_is_vpn_gateway_advertised_cidr](ibm_is_vpn_gateway_advertised_cidr.md) | resource | Schema validation for the ibm_is_vpn_gateway_advertised_cidr resource. |
| [ibm_is_vpn_gateway_advertised_cidrs_data](data_ibm_is_vpn_gateway_advertised_cidrs.md) | data | Schema validation for the ibm_is_vpn_gateway_advertised_cidrs data source. |
| [ibm_is_vpn_gateway_connection_data](data_ibm_is_vpn_gateway_connection.md) | data | Reads ibm_is_vpn_gateway_connection. |
| [ibm_is_vpn_gateway_connection](ibm_is_vpn_gateway_connection.md) | resource | Manages ibm_is_vpn_gateway_connection. |
| [ibm_is_vpn_gateway_connection_local_cidrs_data](data_ibm_is_vpn_gateway_connection_local_cidrs.md) | data | Schema validation for the ibm_is_vpn_gateway_connection_local_cidrs data source. |
| [ibm_is_vpn_gateway_connection_peer_cidrs_data](data_ibm_is_vpn_gateway_connection_peer_cidrs.md) | data | Schema validation for the ibm_is_vpn_gateway_connection_peer_cidrs data source. |
| [ibm_is_vpn_gateway_connections_data](data_ibm_is_vpn_gateway_connections.md) | data | Schema validation for the ibm_is_vpn_gateway_connections data source. |
| [ibm_is_vpn_gateway_service_connection_data](data_ibm_is_vpn_gateway_service_connection.md) | data | Schema validation for the ibm_is_vpn_gateway_service_connection data source. |
| [ibm_is_vpn_gateway_service_connections_data](data_ibm_is_vpn_gateway_service_connections.md) | data | Schema validation for the ibm_is_vpn_gateway_service_connections data source. |
| [ibm_is_vpn_gateways_data](data_ibm_is_vpn_gateways.md) | data | Schema validation for the ibm_is_vpn_gateways data source. |
| [ibm_is_vpn_server_data](data_ibm_is_vpn_server.md) | data | Reads ibm_is_vpn_server. |
| [ibm_is_vpn_server](ibm_is_vpn_server.md) | resource | Manages ibm_is_vpn_server. |
| [ibm_is_vpn_server_client_data](data_ibm_is_vpn_server_client.md) | data | Schema validation for the ibm_is_vpn_server_client data source. |
| [ibm_is_vpn_server_client](ibm_is_vpn_server_client.md) | resource | Schema validation for the ibm_is_vpn_server_client resource. |
| [ibm_is_vpn_server_client_configuration_data](data_ibm_is_vpn_server_client_configuration.md) | data | Schema validation for the ibm_is_vpn_server_client_configuration data source. |
| [ibm_is_vpn_server_clients_data](data_ibm_is_vpn_server_clients.md) | data | Schema validation for the ibm_is_vpn_server_clients data source. |
| [ibm_is_vpn_server_route_data](data_ibm_is_vpn_server_route.md) | data | Reads ibm_is_vpn_server_route. |
| [ibm_is_vpn_server_route](ibm_is_vpn_server_route.md) | resource | Manages ibm_is_vpn_server_route. |
| [ibm_is_vpn_server_routes_data](data_ibm_is_vpn_server_routes.md) | data | Schema validation for the ibm_is_vpn_server_routes data source. |
| [ibm_is_vpn_servers_data](data_ibm_is_vpn_servers.md) | data | Schema validation for the ibm_is_vpn_servers data source. |
| [ibm_is_zone_data](data_ibm_is_zone.md) | data | Schema validation for the ibm_is_zone data source. |
| [ibm_is_zones_data](data_ibm_is_zones.md) | data | Schema validation for the ibm_is_zones data source. |

