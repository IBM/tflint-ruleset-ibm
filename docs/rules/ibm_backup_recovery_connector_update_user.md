# ibm_backup_recovery_connector_update_user

Schema validation for the ibm_backup_recovery_connector_update_user resource.

This rule validates `ibm_backup_recovery_connector_update_user` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Patch Connector Request` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_connector_update_user" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `additional_group_names` | no | — |
| `allow_dso_modify` | no | — |
| `authentication_type` | no | one of: `kAuthAd`, `kAuthGoogle`, `kAuthLocal`, `kAuthSalesforce`, `kAuthSso` |
| `created_time_msecs` | no | — |
| `current_password` | no | — |
| `description` | no | — |
| `domain` | no | — |
| `effective_time_msecs` | no | — |
| `email_address` | no | — |
| `expired_time_msecs` | no | — |
| `force_password_change` | no | — |
| `group_roles` | no | — |
| `intercom_messenger_token` | no | — |
| `is_account_locked` | no | — |
| `is_account_mfa_enabled` | no | — |
| `is_active` | no | — |
| `is_cluster_mfa_enabled` | no | — |
| `last_successful_login_time_msecs` | no | — |
| `last_updated_time_msecs` | no | — |
| `mfa_methods` | no | — |
| `object_class` | no | — |
| `password` | no | — |
| `previous_login_time_msecs` | no | — |
| `primary_group_name` | no | — |
| `privilege_ids` | no | — |
| `restricted` | no | — |
| `roles` | no | — |
| `s3_access_key_id` | no | — |
| `s3_account_id` | no | — |
| `s3_secret_key` | no | — |
| `session_name_cookie` | yes | — |
| `sid` | no | — |
| `tenant_id` | no | — |
| `username` | no | — |
| `ad_user_info` (block) | no | nested block |
| `ad_user_info.group_sids` | no | — |
| `ad_user_info.groups` | no | — |
| `ad_user_info.is_floating_user` | no | — |
| `audit_log_settings` (block) | no | nested block |
| `audit_log_settings.read_logging` | no | — |
| `cluster_identifiers` (block) | no | nested block |
| `cluster_identifiers.cluster_id` | no | — |
| `cluster_identifiers.cluster_incarnation_id` | no | — |
| `google_account` (block) | no | nested block |
| `google_account.account_id` | no | — |
| `google_account.user_id` | no | — |
| `idp_user_info` (block) | no | nested block |
| `idp_user_info.group_sids` | no | — |
| `idp_user_info.groups` | no | — |
| `idp_user_info.idp_id` | no | — |
| `idp_user_info.is_floating_user` | no | — |
| `idp_user_info.issuer_id` | no | — |
| `idp_user_info.user_id` | no | — |
| `idp_user_info.vendor` | no | — |
| `mfa_info` (block) | no | nested block |
| `mfa_info.is_user_exempt_from_mfa` | no | — |
| `org_membership` (block) | no | nested block |
| `org_membership.bifrost_enabled` | no | — |
| `org_membership.is_managed_on_helios` | no | — |
| `org_membership.name` | no | — |
| `org_membership.restricted` | no | — |
| `org_membership.roles` | no | — |
| `org_membership.tenant_id` | no | — |
| `preferences` (block) | no | nested block |
| `preferences.locale` | no | — |
| `profiles` (block) | no | nested block |
| `profiles.is_active` | no | — |
| `profiles.is_deleted` | no | — |
| `profiles.region_ids` | no | — |
| `profiles.tenant_id` | no | — |
| `profiles.tenant_name` | no | — |
| `profiles.tenant_type` | no | — |
| `profiles.cluster_identifiers` (block) | no | nested block |
| `profiles.cluster_identifiers.cluster_id` | no | — |
| `profiles.cluster_identifiers.cluster_incarnation_id` | no | — |
| `salesforce_account` (block) | no | nested block |
| `salesforce_account.account_id` | no | — |
| `salesforce_account.helios_access_grant_status` | no | — |
| `salesforce_account.is_d_gaa_s_user` | no | — |
| `salesforce_account.is_d_maa_s_user` | no | — |
| `salesforce_account.is_d_raa_s_user` | no | — |
| `salesforce_account.is_r_paa_s_user` | no | — |
| `salesforce_account.is_sales_user` | no | — |
| `salesforce_account.is_support_user` | no | — |
| `salesforce_account.user_id` | no | — |
| `spog_context` (block) | no | nested block |
| `spog_context.primary_cluster_id` | no | — |
| `spog_context.primary_cluster_user_sid` | no | — |
| `spog_context.primary_cluster_username` | no | — |
| `subscription_info` (block) | no | nested block |
| `subscription_info.classification` (block) | no | nested block |
| `subscription_info.classification.end_date` | no | — |
| `subscription_info.classification.is_active` | no | — |
| `subscription_info.classification.is_free_trial` | no | — |
| `subscription_info.classification.start_date` | no | — |
| `subscription_info.data_protect` (block) | no | nested block |
| `subscription_info.data_protect.end_date` | no | — |
| `subscription_info.data_protect.is_active` | no | — |
| `subscription_info.data_protect.is_aws_subscription` | no | — |
| `subscription_info.data_protect.is_cohesity_subscription` | no | — |
| `subscription_info.data_protect.is_free_trial` | no | — |
| `subscription_info.data_protect.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.data_protect.start_date` | no | — |
| `subscription_info.data_protect.tiering` (block) | no | nested block |
| `subscription_info.data_protect.tiering.backend_tiering` | no | — |
| `subscription_info.data_protect.tiering.frontend_tiering` | no | — |
| `subscription_info.data_protect.tiering.max_retention` | no | — |
| `subscription_info.data_protect_azure` (block) | no | nested block |
| `subscription_info.data_protect_azure.end_date` | no | — |
| `subscription_info.data_protect_azure.is_active` | no | — |
| `subscription_info.data_protect_azure.is_free_trial` | no | — |
| `subscription_info.data_protect_azure.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.data_protect_azure.start_date` | no | — |
| `subscription_info.data_protect_azure.tiering` (block) | no | nested block |
| `subscription_info.data_protect_azure.tiering.backend_tiering` | no | — |
| `subscription_info.data_protect_azure.tiering.frontend_tiering` | no | — |
| `subscription_info.data_protect_azure.tiering.max_retention` | no | — |
| `subscription_info.fort_knox_azure_cool` (block) | no | nested block |
| `subscription_info.fort_knox_azure_cool.end_date` | no | — |
| `subscription_info.fort_knox_azure_cool.is_active` | no | — |
| `subscription_info.fort_knox_azure_cool.is_free_trial` | no | — |
| `subscription_info.fort_knox_azure_cool.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.fort_knox_azure_cool.start_date` | no | — |
| `subscription_info.fort_knox_azure_hot` (block) | no | nested block |
| `subscription_info.fort_knox_azure_hot.end_date` | no | — |
| `subscription_info.fort_knox_azure_hot.is_active` | no | — |
| `subscription_info.fort_knox_azure_hot.is_free_trial` | no | — |
| `subscription_info.fort_knox_azure_hot.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.fort_knox_azure_hot.start_date` | no | — |
| `subscription_info.fort_knox_cold` (block) | no | nested block |
| `subscription_info.fort_knox_cold.end_date` | no | — |
| `subscription_info.fort_knox_cold.is_active` | no | — |
| `subscription_info.fort_knox_cold.is_free_trial` | no | — |
| `subscription_info.fort_knox_cold.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.fort_knox_cold.start_date` | no | — |
| `subscription_info.ransomware` (block) | no | nested block |
| `subscription_info.ransomware.end_date` | no | — |
| `subscription_info.ransomware.is_active` | no | — |
| `subscription_info.ransomware.is_free_trial` | no | — |
| `subscription_info.ransomware.quantity` | no | API: Specifies the quantity of the subscription. |
| `subscription_info.ransomware.start_date` | no | — |
| `subscription_info.site_continuity` (block) | no | nested block |
| `subscription_info.site_continuity.end_date` | no | — |
| `subscription_info.site_continuity.is_active` | no | — |
| `subscription_info.site_continuity.is_free_trial` | no | — |
| `subscription_info.site_continuity.start_date` | no | — |
| `subscription_info.threat_protection` (block) | no | nested block |
| `subscription_info.threat_protection.end_date` | no | — |
| `subscription_info.threat_protection.is_active` | no | — |
| `subscription_info.threat_protection.is_free_trial` | no | — |
| `subscription_info.threat_protection.start_date` | no | — |
| `tenant_accesses` (block) | no | nested block |
| `tenant_accesses.created_time_msecs` | no | — |
| `tenant_accesses.effective_time_msecs` | no | — |
| `tenant_accesses.expired_time_msecs` | no | — |
| `tenant_accesses.is_access_active` | no | — |
| `tenant_accesses.is_active` | no | — |
| `tenant_accesses.is_deleted` | no | — |
| `tenant_accesses.last_updated_time_msecs` | no | — |
| `tenant_accesses.roles` | no | — |
| `tenant_accesses.tenant_id` | no | — |
| `tenant_accesses.tenant_name` | no | — |
| `tenant_accesses.tenant_type` | no | — |
| `tenant_accesses.cluster_identifiers` (block) | no | nested block |
| `tenant_accesses.cluster_identifiers.cluster_id` | no | — |
| `tenant_accesses.cluster_identifiers.cluster_incarnation_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_update_user_invalid_additional_group_names` | `additional_group_names` |
| `ibm_backup_recovery_connector_update_user_invalid_allow_dso_modify` | `allow_dso_modify` |
| `ibm_backup_recovery_connector_update_user_invalid_authentication_type` | `authentication_type` |
| `ibm_backup_recovery_connector_update_user_invalid_created_time_msecs` | `created_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_current_password` | `current_password` |
| `ibm_backup_recovery_connector_update_user_invalid_description` | `description` |
| `ibm_backup_recovery_connector_update_user_invalid_domain` | `domain` |
| `ibm_backup_recovery_connector_update_user_invalid_effective_time_msecs` | `effective_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_email_address` | `email_address` |
| `ibm_backup_recovery_connector_update_user_invalid_expired_time_msecs` | `expired_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_force_password_change` | `force_password_change` |
| `ibm_backup_recovery_connector_update_user_invalid_group_roles` | `group_roles` |
| `ibm_backup_recovery_connector_update_user_invalid_intercom_messenger_token` | `intercom_messenger_token` |
| `ibm_backup_recovery_connector_update_user_invalid_is_account_locked` | `is_account_locked` |
| `ibm_backup_recovery_connector_update_user_invalid_is_active` | `is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_last_successful_login_time_msecs` | `last_successful_login_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_last_updated_time_msecs` | `last_updated_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_mfa_methods` | `mfa_methods` |
| `ibm_backup_recovery_connector_update_user_invalid_object_class` | `object_class` |
| `ibm_backup_recovery_connector_update_user_invalid_password` | `password` |
| `ibm_backup_recovery_connector_update_user_invalid_previous_login_time_msecs` | `previous_login_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_primary_group_name` | `primary_group_name` |
| `ibm_backup_recovery_connector_update_user_invalid_privilege_ids` | `privilege_ids` |
| `ibm_backup_recovery_connector_update_user_invalid_restricted` | `restricted` |
| `ibm_backup_recovery_connector_update_user_invalid_roles` | `roles` |
| `ibm_backup_recovery_connector_update_user_invalid_s3_access_key_id` | `s3_access_key_id` |
| `ibm_backup_recovery_connector_update_user_invalid_s3_account_id` | `s3_account_id` |
| `ibm_backup_recovery_connector_update_user_invalid_s3_secret_key` | `s3_secret_key` |
| `ibm_backup_recovery_connector_update_user_invalid_session_name_cookie` | `session_name_cookie` |
| `ibm_backup_recovery_connector_update_user_invalid_sid` | `sid` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_id` | `tenant_id` |
| `ibm_backup_recovery_connector_update_user_invalid_username` | `username` |
| `ibm_backup_recovery_connector_update_user_invalid_ad_user_info__group_sids` | `ad_user_info.group_sids` |
| `ibm_backup_recovery_connector_update_user_invalid_ad_user_info__groups` | `ad_user_info.groups` |
| `ibm_backup_recovery_connector_update_user_invalid_ad_user_info__is_floating_user` | `ad_user_info.is_floating_user` |
| `ibm_backup_recovery_connector_update_user_invalid_audit_log_settings__read_logging` | `audit_log_settings.read_logging` |
| `ibm_backup_recovery_connector_update_user_invalid_cluster_identifiers__cluster_id` | `cluster_identifiers.cluster_id` |
| `ibm_backup_recovery_connector_update_user_invalid_cluster_identifiers__cluster_incarnation_id` | `cluster_identifiers.cluster_incarnation_id` |
| `ibm_backup_recovery_connector_update_user_invalid_google_account__account_id` | `google_account.account_id` |
| `ibm_backup_recovery_connector_update_user_invalid_google_account__user_id` | `google_account.user_id` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__group_sids` | `idp_user_info.group_sids` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__groups` | `idp_user_info.groups` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__idp_id` | `idp_user_info.idp_id` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__is_floating_user` | `idp_user_info.is_floating_user` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__issuer_id` | `idp_user_info.issuer_id` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__user_id` | `idp_user_info.user_id` |
| `ibm_backup_recovery_connector_update_user_invalid_idp_user_info__vendor` | `idp_user_info.vendor` |
| `ibm_backup_recovery_connector_update_user_invalid_mfa_info__is_user_exempt_from_mfa` | `mfa_info.is_user_exempt_from_mfa` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__bifrost_enabled` | `org_membership.bifrost_enabled` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__is_managed_on_helios` | `org_membership.is_managed_on_helios` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__name` | `org_membership.name` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__restricted` | `org_membership.restricted` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__roles` | `org_membership.roles` |
| `ibm_backup_recovery_connector_update_user_invalid_org_membership__tenant_id` | `org_membership.tenant_id` |
| `ibm_backup_recovery_connector_update_user_invalid_preferences__locale` | `preferences.locale` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__is_active` | `profiles.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__is_deleted` | `profiles.is_deleted` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__region_ids` | `profiles.region_ids` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__tenant_id` | `profiles.tenant_id` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__tenant_name` | `profiles.tenant_name` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__tenant_type` | `profiles.tenant_type` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__cluster_identifiers__cluster_id` | `profiles.cluster_identifiers.cluster_id` |
| `ibm_backup_recovery_connector_update_user_invalid_profiles__cluster_identifiers__cluster_incarnation_id` | `profiles.cluster_identifiers.cluster_incarnation_id` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__account_id` | `salesforce_account.account_id` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__helios_access_grant_status` | `salesforce_account.helios_access_grant_status` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_d_gaa_s_user` | `salesforce_account.is_d_gaa_s_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_d_maa_s_user` | `salesforce_account.is_d_maa_s_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_d_raa_s_user` | `salesforce_account.is_d_raa_s_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_r_paa_s_user` | `salesforce_account.is_r_paa_s_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_sales_user` | `salesforce_account.is_sales_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__is_support_user` | `salesforce_account.is_support_user` |
| `ibm_backup_recovery_connector_update_user_invalid_salesforce_account__user_id` | `salesforce_account.user_id` |
| `ibm_backup_recovery_connector_update_user_invalid_spog_context__primary_cluster_id` | `spog_context.primary_cluster_id` |
| `ibm_backup_recovery_connector_update_user_invalid_spog_context__primary_cluster_user_sid` | `spog_context.primary_cluster_user_sid` |
| `ibm_backup_recovery_connector_update_user_invalid_spog_context__primary_cluster_username` | `spog_context.primary_cluster_username` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__classification__end_date` | `subscription_info.classification.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__classification__is_active` | `subscription_info.classification.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__classification__is_free_trial` | `subscription_info.classification.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__classification__start_date` | `subscription_info.classification.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__end_date` | `subscription_info.data_protect.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__is_active` | `subscription_info.data_protect.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__is_aws_subscription` | `subscription_info.data_protect.is_aws_subscription` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__is_cohesity_subscription` | `subscription_info.data_protect.is_cohesity_subscription` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__is_free_trial` | `subscription_info.data_protect.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__quantity` | `subscription_info.data_protect.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__start_date` | `subscription_info.data_protect.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__tiering__backend_tiering` | `subscription_info.data_protect.tiering.backend_tiering` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__tiering__frontend_tiering` | `subscription_info.data_protect.tiering.frontend_tiering` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect__tiering__max_retention` | `subscription_info.data_protect.tiering.max_retention` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__end_date` | `subscription_info.data_protect_azure.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__is_active` | `subscription_info.data_protect_azure.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__is_free_trial` | `subscription_info.data_protect_azure.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__quantity` | `subscription_info.data_protect_azure.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__start_date` | `subscription_info.data_protect_azure.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__tiering__backend_tiering` | `subscription_info.data_protect_azure.tiering.backend_tiering` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__tiering__frontend_tiering` | `subscription_info.data_protect_azure.tiering.frontend_tiering` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__data_protect_azure__tiering__max_retention` | `subscription_info.data_protect_azure.tiering.max_retention` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_cool__end_date` | `subscription_info.fort_knox_azure_cool.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_cool__is_active` | `subscription_info.fort_knox_azure_cool.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_cool__is_free_trial` | `subscription_info.fort_knox_azure_cool.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_cool__quantity` | `subscription_info.fort_knox_azure_cool.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_cool__start_date` | `subscription_info.fort_knox_azure_cool.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_hot__end_date` | `subscription_info.fort_knox_azure_hot.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_hot__is_active` | `subscription_info.fort_knox_azure_hot.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_hot__is_free_trial` | `subscription_info.fort_knox_azure_hot.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_hot__quantity` | `subscription_info.fort_knox_azure_hot.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_azure_hot__start_date` | `subscription_info.fort_knox_azure_hot.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_cold__end_date` | `subscription_info.fort_knox_cold.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_cold__is_active` | `subscription_info.fort_knox_cold.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_cold__is_free_trial` | `subscription_info.fort_knox_cold.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_cold__quantity` | `subscription_info.fort_knox_cold.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__fort_knox_cold__start_date` | `subscription_info.fort_knox_cold.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__ransomware__end_date` | `subscription_info.ransomware.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__ransomware__is_active` | `subscription_info.ransomware.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__ransomware__is_free_trial` | `subscription_info.ransomware.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__ransomware__quantity` | `subscription_info.ransomware.quantity` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__ransomware__start_date` | `subscription_info.ransomware.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__site_continuity__end_date` | `subscription_info.site_continuity.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__site_continuity__is_active` | `subscription_info.site_continuity.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__site_continuity__is_free_trial` | `subscription_info.site_continuity.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__site_continuity__start_date` | `subscription_info.site_continuity.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__threat_protection__end_date` | `subscription_info.threat_protection.end_date` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__threat_protection__is_active` | `subscription_info.threat_protection.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__threat_protection__is_free_trial` | `subscription_info.threat_protection.is_free_trial` |
| `ibm_backup_recovery_connector_update_user_invalid_subscription_info__threat_protection__start_date` | `subscription_info.threat_protection.start_date` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__created_time_msecs` | `tenant_accesses.created_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__effective_time_msecs` | `tenant_accesses.effective_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__expired_time_msecs` | `tenant_accesses.expired_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__is_access_active` | `tenant_accesses.is_access_active` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__is_active` | `tenant_accesses.is_active` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__is_deleted` | `tenant_accesses.is_deleted` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__last_updated_time_msecs` | `tenant_accesses.last_updated_time_msecs` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__roles` | `tenant_accesses.roles` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__tenant_id` | `tenant_accesses.tenant_id` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__tenant_name` | `tenant_accesses.tenant_name` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__tenant_type` | `tenant_accesses.tenant_type` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__cluster_identifiers__cluster_id` | `tenant_accesses.cluster_identifiers.cluster_id` |
| `ibm_backup_recovery_connector_update_user_invalid_tenant_accesses__cluster_identifiers__cluster_incarnation_id` | `tenant_accesses.cluster_identifiers.cluster_incarnation_id` |

## Example

### Valid

```hcl
variable "session_name_cookie" {
  sensitive = true
}

resource "ibm_backup_recovery_connector_update_user" "test" {
  session_name_cookie = var.session_name_cookie
}
```
